// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"encoding/json"
	"fmt"
	"github.com/gedex/inflector"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// AzureApiParameter represents a parameter of a Azure REST API endpoint.
type AzureApiParameter struct {
	Name string
	// Location defines the parameter's place the HTTP request: "path", "query", or "body".
	Location string
	// Source defines the value source: "method" (resource arguments) or "client" (provider configuration).
	Source string
	// Requires is true for mandatory parameters.
	IsRequired bool
	// IsResourceName is true for parameters that contain resource name (e.g. `accountName`).
	IsResourceName bool
	// Properties contains the names of properties for a body-placed parameter.
	Properties []string
	// RequiredProperties contains the names of required properties for a body-placed parameter.
	RequiredProperties []string
}

// AzureApiResource is a resource in Azure REST API.
type AzureApiResource struct {
	ApiVersion    string
	Path          string
	GetParameters []AzureApiParameter
	PutParameters []AzureApiParameter
}

type AzureApiMetadata struct {
	Resources map[string]AzureApiResource
	Invokes   map[string]AzureApiResource
}

func GenerateResourceMap(specs []*openapi.Spec) error {
	resourceMap, err := buildResourceMap(specs)
	if err != nil {
		return err
	}

	contents, err := json.Marshal(resourceMap)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./provider/pkg/provider/metadata.go", []byte(fmt.Sprintf(`package provider
var azureApiResources = %#v
`, contents)), 0600)
	if err != nil {
		return err
	}

	return nil
}

func buildResourceMap(specs []*openapi.Spec) (*AzureApiMetadata, error) {
	resources := map[string]AzureApiResource{}
	invokes := map[string]AzureApiResource{}

	for _, spec := range specs {
		for key, path := range spec.Paths.Paths {
			if path.Put == nil || path.Get == nil || path.Delete == nil {
				continue
			}

			provider, resourceName := ResourceQualifiedName(key)
			if resourceName == "" {
				continue
			}
			module := strings.ToLower(provider)

			path := spec.Swagger.Paths.Paths[key]
			if path.Put == nil {
				continue
			}

			nameParam := nameParameter(key)

			gets, err := buildParameters(spec, path.Get.Parameters, nameParam)
			if err != nil {
				return nil, err
			}

			puts, err := buildParameters(spec, path.Put.Parameters, nameParam)
			if err != nil {
				return nil, err
			}

			r := AzureApiResource{
				ApiVersion:    spec.Info.Version,
				Path:          key,
				GetParameters: gets,
				PutParameters: puts,
			}
			resourceTypeName := fmt.Sprintf("azurerm:%s:%s", module, resourceName)
			resources[resourceTypeName] = r

			f := AzureApiResource{
				ApiVersion:    spec.Info.Version,
				Path:          key,
				GetParameters: gets,
			}
			functionTypeName := fmt.Sprintf("azurerm:%s:get%s", module, resourceName)
			invokes[functionTypeName] = f
		}
	}

	return &AzureApiMetadata{Resources: resources, Invokes: invokes}, nil
}

func buildParameters(spec *openapi.Spec, parameters []spec.Parameter, nameParam string) ([]AzureApiParameter, error) {
	var puts []AzureApiParameter
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

			properties, required, err = resolveProperties(*schema)
			if err != nil {
				return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}
		}

		location, _ := param.Extensions.GetString("x-ms-parameter-location")
		p := AzureApiParameter{
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
func resolveProperties(schema openapi.Schema) ([]string, []string, error) {
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

		ps, rs, err := resolveProperties(*allOfSchema)
		if err != nil {
			return nil, nil, err
		}

		properties = append(properties, ps...)
		required = append(required, rs...)
	}

	return properties, required, nil
}

// nameParameter parses the given URL path to find the name of the last template parameter.
func nameParameter(path string) string {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			return part[1 : len(part)-1]
		}
	}
	return ""
}

// blessedVersions contains the preferred versions per resource provider. If a resource provider is not specified,
// the latest stable version is used.
var blessedVersions map[string]string

func init() {
	blessedVersions = map[string]string{
		"Microsoft.Cdn": "2020-03-31", // the later version throws runtime errors
		"Microsoft.Web": "2019-08-01", // an earlier version contains extra files that spoil generation
	}
}

// SwaggerLocations returns a slice of URLs of all known Azure Resource Manager swagger files.
func SwaggerLocations() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(dir + "/azure-rest-api-specs/specification/**/resource-manager/Microsoft.*/stable/*/*.json")
	if err != nil {
		return nil, err
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	sort.Strings(files)

	// Deduplicate files of the same provider having the same name.
	latest := map[string]string{}
	for _, file := range files {
		parts := strings.Split(file, "/")
		length := len(parts)
		provider := parts[length-4]
		apiVersion := parts[length-2]
		filename := parts[length-1]

		if v, ok := blessedVersions[provider]; ok && v != apiVersion {
			continue
		}

		key := provider + "/" + filename
		latest[key] = file
	}

	var locations []string
	for _, file := range latest {
		locations = append(locations, file)
	}
	sort.Strings(locations)

	return locations, nil
}

// ResourceQualifiedName returns a tuple of (module, resource name) for a given PUT path.
func ResourceQualifiedName(path string) (string, string) {
	if path == "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}" {
		return "core", "ResourceGroup"
	}

	parts := strings.Split(path, "/")

	// The path is /subscriptions/id/resoursegroups/id/providers/Microsoft.SomeProvider/foos/id/bars/id/...
	if len(parts) < 8 || len(parts)%2 != 1 {
		return "", ""
	}

	armProvider := parts[6]
	// TODO: handle non-Microsoft providers.
	if !strings.HasPrefix(armProvider, "Microsoft.") {
		return "", ""
	}

	// Build a name like SomeProvider:FooBar.
	provider := armProvider[10:]
	resource := ""
	for i := 7; i < len(parts); i += 2 {
		if strings.Contains(parts[i], "{") {
			// We don't support this shape of URLs yet. Example: dnszones/{zoneName}/{recordType}/{relativeRecordSetName}.
			return "", ""
		}

		// TODO: generalize this case to a map of well-known aliases.
		switch strings.ToLower(parts[i]) {
		case "redis":
			resource += "Redis"
		case "sites":
			resource += "AppService"
		case "serverfarms":
			resource += "AppServicePlan"
		default:
			// TODO: we may get better singular names from some metadata.
			resource += strings.ReplaceAll(strings.Title(inflector.Singularize(parts[i])), "-", "")
		}
	}
	return provider, resource
}
