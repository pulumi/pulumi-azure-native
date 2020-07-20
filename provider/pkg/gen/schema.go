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
	"sort"
	"strings"
)

// PulumiSchema will generate a Pulumi schema for the given swagger specs.
func PulumiSchema(swaggers []*openapi.Spec) (*pschema.PackageSpec, *provider.AzureApiMetadata, error) {
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
	metadata := provider.AzureApiMetadata{
		Resources: map[string]provider.AzureApiResource{},
		Invokes:   map[string]provider.AzureApiInvoke{},
	}

	csharpNamespaces := map[string]string{
		"azurerm": "AzureRM",
	}

	for _, swagger := range swaggers {
		gen := packageGenerator{pkg: &pkg, metadata: &metadata, swagger: swagger}

		// Sort paths to make codegen deterministic.
		var paths []string
		for key, _ := range swagger.Paths.Paths {
			paths = append(paths, key)
		}
		sort.Strings(paths)

		for _, key := range paths {
			path := swagger.Paths.Paths[key]

			prov, _ := provider.ResourceQualifiedName(key)
			if prov != "" {
				module := strings.ToLower(prov)
				csharpNamespaces[module] = prov
			}

			gen.genResources(key, &path)
			gen.genListFunctions(key, &path)
		}
	}

	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^2.0.0",
		},
		"readme": pkg.Description, // TODO: add a proper readme.
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"requires": map[string]string{
			"pulumi": ">=2.0.0,<3.0.0",
		},
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, nil
}

type packageGenerator struct {
	pkg      *pschema.PackageSpec
	swagger  *openapi.Spec
	metadata *provider.AzureApiMetadata
}

func (g *packageGenerator) genResources(key string, path *spec.PathItem) {
	if path.Put == nil || path.Get == nil || path.Delete == nil {
		return
	}

	prov, typeName := provider.ResourceQualifiedName(key)
	if typeName == "" {
		return
	}

	module := strings.ToLower(prov)
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:           g.pkg,
		module:        module,
		resourceToken: resourceTok,
		visitedTypes:  make(map[string]bool),
	}

	resourceRequest, err := gen.genMethodParameters(path.Put.Parameters, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", resourceTok, err.Error())
		return
	}

	resourceResponse, err := gen.genResponse(path.Put.Responses.StatusCodeResponses, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", resourceTok, err.Error())
		return
	}

	err = gen.normalizeName(key, resourceRequest, resourceResponse)
	if err != nil {
		log.Printf("failed to assign name for '%s'", resourceTok, err.Error())
		return
	}

	objectSpec := pschema.ObjectTypeSpec{
		Description: resourceResponse.description,
		Type:        "object",
		Properties:  resourceResponse.all,
		Required:    resourceResponse.required.SortedValues(),
	}
	g.pkg.Types[resourceTok] = objectSpec

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec:  objectSpec,
		InputProperties: resourceRequest.all,
		RequiredInputs:  resourceRequest.required.SortedValues(),
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, typeName)

	requestFunction, err := gen.genMethodParameters(path.Get.Parameters, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	responseFunction, err := gen.genResponse(path.Get.Responses.StatusCodeResponses, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	gen.normalizeName(key, requestFunction, responseFunction)

	functionSpec := pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Description: requestFunction.description,
			Type:        "object",
			Properties:  requestFunction.all,
			Required:    requestFunction.required.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: responseFunction.description,
			Type:        "object",
			Properties:  responseFunction.all,
			Required:    responseFunction.required.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	r := provider.AzureApiResource{
		ApiVersion:    g.swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.metadata,
		PutParameters: resourceRequest.metadata,
	}
	g.metadata.Resources[resourceTok] = r

	f := provider.AzureApiInvoke{
		ApiVersion:    g.swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.metadata,
	}
	g.metadata.Invokes[functionTok] = f
}

// genListFunctions defines functions for list* (listKeys, listSecrets, etc.) POST endpoints.
func (g *packageGenerator) genListFunctions(key string, path *spec.PathItem) {
	if path.Post == nil || !strings.Contains(key, "/list") {
		return
	}

	parts := strings.Split(key, "/")
	listOperation := parts[len(parts)-1]
	if !strings.HasPrefix(listOperation, "list") {
		return
	}

	baseUrl := strings.TrimSuffix(key, "/" + listOperation)
	prov, typeName := provider.ResourceQualifiedName(baseUrl)
	if typeName == "" {
		return
	}

	module := strings.ToLower(prov)

	gen := moduleGenerator{
		pkg:           g.pkg,
		module:        module,
		resourceToken: fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		visitedTypes:  make(map[string]bool),
	}

	// Generate the function to get this resource.
	subject := strings.Title(strings.TrimPrefix(listOperation, "list"))
	functionTok := fmt.Sprintf(`%s:%s:list%s%s`, g.pkg.Name, module, typeName, subject)

	request, err := gen.genMethodParameters(path.Post.Parameters, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(path.Post.Responses.StatusCodeResponses, g.swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	gen.normalizeName(key, request, response)

	functionSpec := pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Description: request.description,
			Type:        "object",
			Properties:  request.all,
			Required:    request.required.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: response.description,
			Type:        "object",
			Properties:  response.all,
			Required:    response.required.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	f := provider.AzureApiInvoke{
		ApiVersion:     g.swagger.Info.Version,
		Path:           key,
		PostParameters: request.metadata,
	}
	g.metadata.Invokes[functionTok] = f
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

	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name := part[1 : len(part)-1]
			if nameProp, ok := requestProperties.all[name]; ok {
				// We expect names to be always a required string.
				if nameProp.Type != "string" {
					return errors.Errorf("name property '%s' is not a string", name)
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

		var bodyProperties []string
		var bodyRequiredProperties []string

		switch {
		case param.Name == "subscriptionId":
		case !isLegalIdentifier(param.Name): // If-Match, Accept-Header, x-ms-foobar, ...
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

			for k, _ := range props.all {
				bodyProperties = append(bodyProperties, k)
			}
			sort.Strings(bodyProperties)
			for _, k := range props.required.SortedValues() {
				bodyRequiredProperties = append(bodyRequiredProperties, k)
			}

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

		location, _ := param.Extensions.GetString("x-ms-parameter-location")
		p := provider.AzureApiParameter{
			Name:               param.Name,
			Location:           param.In,
			Source:             location,
			IsRequired:         param.Required,
			Properties:         bodyProperties,
			RequiredProperties: bodyRequiredProperties,
		}
		result.metadata = append(result.metadata, p)
	}

	return result, nil
}

func (m *moduleGenerator) genResponse(statusCodeResponses map[int]spec.Response, ctx *openapi.ReferenceContext) (*propertyBag, error) {
	var responseSchema *openapi.Schema

	// Find all 2xx codes and sort them to make codegen deterministic.
	var codes []int
	for code, _ := range statusCodeResponses {
		if code >= 300 {
			continue
		}

		codes = append(codes, code)
	}
	sort.Ints(codes)

	// Find the lowest 2xx response with a schema definition and derive response properties from it.
	for _, code := range codes {
		resp := statusCodeResponses[code]
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

	// Sort properties to make codegen deterministic.
	var names []string
	for name, _ := range resolvedSchema.Properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		property := resolvedSchema.Properties[name]
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
	metadata    []provider.AzureApiParameter
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
