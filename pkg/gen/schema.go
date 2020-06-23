package gen

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"log"
	"strings"
)

// PulumiSchema will generate a Pulumi schema for the given swagger specs.
func PulumiSchema(swaggers []*openapi.Spec) (*pschema.PackageSpec, error) {
	pkg := pschema.PackageSpec{
		Name:        "azurerm",
		Version:     "0.1.0", // TODO
		Description: "A Pulumi package for creating and managing Azure resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "azure"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-azurerm",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{},
		},
		Types:     map[string]pschema.ObjectTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}

	for _, swagger := range swaggers {
		for key, path := range swagger.Paths.Paths {
			if path.Put == nil || path.Get == nil || path.Delete == nil {
				continue
			}

			module, typeName := provider.ResourceQualifiedName(key)
			if typeName == "" {
				continue
			}

			tok := fmt.Sprintf(`%s:%s:%s`, pkg.Name, module, typeName)
			gen := moduleGenerator{
				pkg:           &pkg,
				module:        module,
				resourceToken: tok,
				visitedTypes:  make(map[string]bool),
			}

			requestProperties, err := gen.genMethodParameters(path.Put.Parameters, swagger.ReferenceContext)
			if err != nil {
				log.Printf("failed to generate '%s': request type: %s", tok, err.Error())
				continue
			}

			responseProperties, err := gen.genResponse(path.Put.Responses.StatusCodeResponses, swagger.ReferenceContext)
			if err != nil {
				log.Printf("failed to generate '%s': response type: %s", tok, err.Error())
				continue
			}

			err = gen.normalizeName(key, requestProperties, responseProperties)
			if err != nil {
				log.Printf("failed to assign name for '%s'", tok, err.Error())
				continue
			}

			objectSpec := pschema.ObjectTypeSpec{
				Description: responseProperties.description,
				Type:        "object",
				Properties:  responseProperties.all,
				Required:    responseProperties.required.SortedValues(),
			}
			pkg.Types[tok] = objectSpec

			resourceSpec := pschema.ResourceSpec{
				ObjectTypeSpec:  objectSpec,
				InputProperties: requestProperties.all,
				RequiredInputs:  requestProperties.required.SortedValues(),
				// TODO: this is probably wrong, state inputs are set here because codegen fails without them.
				StateInputs: &objectSpec,
			}
			pkg.Resources[tok] = resourceSpec
		}
	}

	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^2.0.0",
		},
		"readme": pkg.Description, // TODO: add a proper readme.
	})

	return &pkg, nil
}

type moduleGenerator struct {
	pkg           *pschema.PackageSpec
	module        string
	resourceToken string
	visitedTypes  map[string]bool
}

// normalizeName replaces a custom name input property like `accountName` or `resourceGroupName` with the standard
// `name` property.
func (m *moduleGenerator) normalizeName(path string, requestProperties *propertyBag, responseProperties *propertyBag) error {
	// Do nothing if there's no `name` in response properties - we always expect it for any resource.
	if _, ok := responseProperties.all["name"]; !ok {
		return nil
	}

	// Do nothing if there's already a `name` property.
	if _, ok := requestProperties.all["name"]; ok {
		return nil
	}

	parts := strings.Split(path, "/")
	for i := len(parts)-1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name := part[1:len(part)-1]
			if nameProp, ok := requestProperties.all[name]; ok {
				// We expect names to be always a required string.
				if nameProp.Type != "string" {
					return errors.Errorf("name property '%s' is not a string but %v", name, nameProp.Type)
				}
				if !requestProperties.required.Has(name) {
					return errors.Errorf("name property '%s' is not required", name)
				}

				delete(requestProperties.all, name)
				requestProperties.all["name"] = nameProp
				requestProperties.required.Delete(name)
				requestProperties.required.Add("name")
				break
			} else {
				return errors.Errorf("name property '%s' not found", name)
			}
		}
	}

	return nil
}

func (m *moduleGenerator) genMethodParameters(parameters []spec.Parameter, ctx *openapi.ReferenceContext) (*propertyBag, error) {
	result := newPropertyBag()

	for _, param := range parameters {
		param, err := ctx.ResolveParameter(param)
		if err != nil {
			return nil, err
		}

		switch {
		case param.Name == "api-version":
		case param.Name == "subscriptionId":
		case strings.Contains(param.Name, "-"): // If-Match, Accept-Header, x-ms-foobar, ...
			// TODO: Find a more principled criteria to skip those.

		// The body parameter is flattened, so that all its properties become the properties of the type.
		case param.In == "body":
			if param.Schema == nil {
				return nil, errors.Errorf("no schema for body parameter '%s'", param.Name)
			}

			resolvedSchema, err := param.ResolveSchema(param.Schema)
			if err != nil {
				return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
			}

			props, err := m.genProperties(resolvedSchema, false)
			if err != nil {
				return nil, err
			}

			result.merge(props)

		default:
			propertySpec := pschema.PropertySpec{
				Description: param.Description,
				TypeSpec: pschema.TypeSpec{
					Type: param.Type,
				},
			}
			result.all[param.Name] = propertySpec
			if param.Required {
				result.required.Add(param.Name)
			}
		}
	}

	return result, nil
}

func (m *moduleGenerator) genResponse(statusCodeResponses map[int]spec.Response, ctx *openapi.ReferenceContext) (*propertyBag, error) {
	var responseSchema *openapi.Schema
	for code, resp := range statusCodeResponses {
		if code >= 300 {
			continue
		}

		response, err := ctx.ResolveResponse(resp)
		if err != nil {
			return nil, err
		}

		if response.Schema == nil {
			continue
		}

		responseSchema, err = response.ResolveSchema(response.Schema)
		if err != nil {
			return nil, err
		}

		result, err := m.genProperties(responseSchema, true)
		if err != nil {
			return nil, err
		}

		// Id is a property of the base Custom Resource, we don't want to introduce it on derived resources.
		delete(result.all, "id")
		result.required.Delete("id")

		// Special case the 'properties' output property as required.
		// This should be gone when we apply flattening.
		if _, ok := result.all["properties"]; ok {
			result.required.Add("properties")
		}

		result.description = responseSchema.Description
		return result, nil
	}

	return nil, errors.New("no 2xx response found")
}

func (m *moduleGenerator) genProperties(resolvedSchema *openapi.Schema, isOutput bool) (*propertyBag, error) {
	result := newPropertyBag()

	for name, property := range resolvedSchema.Properties {
		if !isLegalIdentifier(name) {
			// TODO: Support mapping to a legal name, or make the schema codegen do so?
			continue
		}

		propertySpec, err := m.genProperty(&property, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, err
		}

		if isOutput {
			result.all[name] = *propertySpec
			if property.ReadOnly {
				result.required.Add(name)
			}
		} else if !property.ReadOnly {
			result.all[name] = *propertySpec
		}
	}

	for _, s := range resolvedSchema.AllOf {
		allOfSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return nil, err
		}

		allOfProperties, err := m.genProperties(allOfSchema, isOutput)
		if err != nil {
			return nil, err
		}

		result.merge(allOfProperties)
	}

	for _, name := range resolvedSchema.Required {
		if _, ok := result.all[name]; ok {
			result.required.Add(name)
		}
	}
	return result, nil
}

func (m *moduleGenerator) genProperty(schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.PropertySpec, error) {
	description := schema.Description
	if description == "" {
		resolvedSchema, err := context.ResolveSchema(schema)
		if err != nil {
			return nil, err
		}

		description = resolvedSchema.Description
	}

	typeSpec, err := m.genTypeSpec(schema, context, isOutput)
	if err != nil {
		return nil, err
	}

	propertySpec := pschema.PropertySpec{
		Description: description,
		TypeSpec:    *typeSpec,
	}

	return &propertySpec, nil
}

func (m *moduleGenerator) genTypeSpec(schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, err
	}

	// Type is either a primitive (e.g. 'string'), or an 'object' defined by a reference to its type.
	var tok, primitiveTypeName, referencedTypeName string
	ptr := schema.Ref.GetPointer()
	if ptr != nil && !ptr.IsEmpty() {
		// Erase type information about enums. TODO: implement proper enums.
		if len(resolvedSchema.Type) > 0 && resolvedSchema.Type[0] == "string" {
			primitiveTypeName = "string"
		} else {
			tok = m.typeName(resolvedSchema.ReferenceContext, isOutput)
			// Avoid collision of resource name vs. property type name (example: azurerm:web:StaticSite).
			if tok == m.resourceToken {
				tok = tok + "Definition"
			}
		}
	} else if len(schema.Properties) > 0 {
		// Inline properties have no type in the Open API schema but we need a type in the Pulumi schema.
		tok = m.typeName(context, isOutput) + "Properties"
	} else if len(schema.Type) > 0 {
		primitiveTypeName = schema.Type[0]
	} else {
		primitiveTypeName = "object"
	}

	// If an object type is referenced, add its definition to the type map.
	if tok != "" {
		referencedTypeName = fmt.Sprintf("#/types/%s", tok)

		if _, ok := m.visitedTypes[tok]; !ok {
			m.visitedTypes[tok] = true

			props, err := m.genProperties(resolvedSchema, isOutput)
			if err != nil {
				return nil, err
			}

			m.pkg.Types[tok] = pschema.ObjectTypeSpec{
				Description: resolvedSchema.Description,
				Type:        "object",
				Properties:  props.all,
				Required:    props.required.SortedValues(),
			}
		}
	}

	result := pschema.TypeSpec{
		Type: primitiveTypeName,
		Ref:  referencedTypeName,
	}

	// Resolve the element type for array types.
	if schema.Items != nil && schema.Items.Schema != nil {
		itemsSpec, err := m.genProperty(schema.Items.Schema, context, isOutput)
		if err != nil {
			return nil, err
		}

		result.Items = &itemsSpec.TypeSpec
	}

	return &result, nil
}

func (m *moduleGenerator) typeName(ctx *openapi.ReferenceContext, isOutput bool) string {
	suffix := ""
	if isOutput {
		suffix = "Response"
	}
	return fmt.Sprintf("azurerm:%s:%s%s", m.module, makeLegalIdentifier(ctx.ReferenceName), suffix)
}

type propertyBag struct {
	description string
	all         map[string]pschema.PropertySpec
	required    codegen.StringSet
}

func newPropertyBag() *propertyBag {
	return &propertyBag{
		all:      map[string]pschema.PropertySpec{},
		required: codegen.NewStringSet(),
	}
}

func (bag *propertyBag) merge(other *propertyBag) {
	for key, value := range other.all {
		bag.all[key] = value
	}
	for key, _ := range other.required {
		bag.required.Add(key)
	}
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
