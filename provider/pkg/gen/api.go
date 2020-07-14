package gen

import (
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/pkg/provider"
	"strings"
)

// apiGenerator creates metadata that the provider needs at runtime to create requests to Azure API endpoints.
type apiGenerator struct {
}

// buildApiMethods returns a map of HTTP methods (get, put, etc.) and parameter definitions for each one of them
// based on the provided Open API Spec.
func (g *apiGenerator) buildApiMethods(spec *openapi.Spec, path string, endpoint *spec.PathItem) (map[string][]provider.AzureApiParameter, error)  {
	nameParam := g.nameParameter(path)
	gets, err := g.buildApiParameters(spec, endpoint.Get.Parameters, nameParam)
	if err != nil {
		return nil, err
	}

	puts, err := g.buildApiParameters(spec, endpoint.Put.Parameters, nameParam)
	if err != nil {
		return nil, err
	}

	result := map[string][]provider.AzureApiParameter{
		"get": gets,
		"put": puts,
	}
	return result, nil
}

func (g *apiGenerator) buildApiParameters(spec *openapi.Spec, parameters []spec.Parameter, nameParam string) ([]provider.AzureApiParameter, error) {
	var puts []provider.AzureApiParameter
	for _, param := range parameters {
		param, err := spec.ResolveParameter(param)
		if err != nil {
			return nil, err
		}

		var properties, required []string
		if param.In == "body" {
			if param.Schema == nil {
				return nil, errors.Errorf("no schema for body parameter '%s'", param.Name)
			}

			schema, err := param.ResolveSchema(param.Schema)
			if err != nil {
				return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}

			properties, required, err = g.resolveProperties(*schema)
			if err != nil {
				return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}
		}

		location, _ := param.Extensions.GetString("x-ms-parameter-location")
		p := provider.AzureApiParameter{
			Name:               param.Name,
			Location:           param.In,
			Source:             location,
			IsRequired:         param.Required,
			IsResourceName:     nameParam == param.Name,
			Properties:         properties,
			RequiredProperties: required,
		}
		puts = append(puts, p)
	}
	return puts, nil
}

// resolveProperties returns the slice of schema's property names and the slice of schema's required properties.
func (g *apiGenerator) resolveProperties(schema openapi.Schema) ([]string, []string, error) {
	var properties []string
	var required []string

	for k, _ := range schema.Properties {
		properties = append(properties, k)
	}
	for _, k := range schema.Required {
		required = append(required, k)
	}

	for _, s := range schema.AllOf {
		allOfSchema, err := schema.ResolveSchema(&s)
		if err != nil {
			return nil, nil, err
		}

		ps, rs, err := g.resolveProperties(*allOfSchema)
		if err != nil {
			return nil, nil, err
		}

		properties = append(properties, ps...)
		required = append(required, rs...)
	}

	return properties, required, nil
}

// nameParameter parses the given URL path to find the name of the last template parameter.
func (g *apiGenerator) nameParameter(path string) string {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			return part[1 : len(part)-1]
		}
	}
	return ""
}
