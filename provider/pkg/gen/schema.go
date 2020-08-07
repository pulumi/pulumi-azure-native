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
		Types:     map[string]provider.AzureApiType{},
		Resources: map[string]provider.AzureApiResource{},
		Invokes:   map[string]provider.AzureApiInvoke{},
	}

	csharpNamespaces := map[string]string{
		"azurerm": "AzureRM",
	}
	pythonModuleNames := map[string]string{}

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

			prov := provider.ResourceProvider(key)
			if prov != "" {
				module := gen.providerToModule(prov)
				version := strings.Replace(gen.apiVersion(), "preview", "Preview", 1)
				csharpNamespaces[module] = fmt.Sprintf("%s.V%s", prov, version)
				pythonModuleNames[module] = module
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
		"moduleNameOverrides": pythonModuleNames,
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

	prov := provider.ResourceProvider(key)
	typeName := provider.ResourceName(path.Get.ID)
	if prov == "" || typeName == "" {
		return
	}

	module := g.providerToModule(prov)
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		resourceToken: resourceTok,
		visitedTypes:  make(map[string]bool),
	}

	parameters := g.mergeParameters(path.Put.Parameters, path.Parameters)
	resourceRequest, err := gen.genMethodParameters(parameters, g.swagger.ReferenceContext)
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
		log.Printf("failed to assign name for '%s': %s", resourceTok, err.Error())
		return
	}

	gen.escapeCSharpNames(typeName, resourceResponse)

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec:  pschema.ObjectTypeSpec{
			Description: resourceResponse.description,
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties: resourceRequest.specs,
		RequiredInputs:  resourceRequest.requiredSpecs.SortedValues(),
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, typeName)

	parameters = g.mergeParameters(path.Get.Parameters, path.Parameters)
	requestFunction, err := gen.genMethodParameters(parameters, g.swagger.ReferenceContext)
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
			Properties:  requestFunction.specs,
			Required:    requestFunction.requiredSpecs.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: responseFunction.description,
			Type:        "object",
			Properties:  responseFunction.specs,
			Required:    responseFunction.requiredSpecs.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	r := provider.AzureApiResource{
		ApiVersion:    g.swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.parameters,
		PutParameters: resourceRequest.parameters,
		Response:      resourceResponse.properties,
	}
	g.metadata.Resources[resourceTok] = r

	f := provider.AzureApiInvoke{
		ApiVersion:    g.swagger.Info.Version,
		Path:          key,
		GetParameters: requestFunction.parameters,
		Response:      responseFunction.properties,
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

	baseUrl := strings.TrimSuffix(key, "/"+listOperation)
	prov := provider.ResourceProvider(baseUrl)
	typeName := provider.ResourceName(path.Post.ID)
	if prov == "" || typeName == "" {
		return
	}

	module := g.providerToModule(prov)
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		resourceToken: fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		visitedTypes:  make(map[string]bool),
	}

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:list%s`, g.pkg.Name, module, typeName)

	parameters := g.mergeParameters(path.Post.Parameters, path.Parameters)
	request, err := gen.genMethodParameters(parameters, g.swagger.ReferenceContext)
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
			Properties:  request.specs,
			Required:    request.requiredSpecs.SortedValues(),
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: response.description,
			Type:        "object",
			Properties:  response.specs,
			Required:    response.requiredSpecs.SortedValues(),
		},
	}
	g.pkg.Functions[functionTok] = functionSpec

	f := provider.AzureApiInvoke{
		ApiVersion:     g.swagger.Info.Version,
		Path:           key,
		PostParameters: request.parameters,
		Response:       response.properties,
	}
	g.metadata.Invokes[functionTok] = f
}

// apiVersion returns the module version from the Azure API spec version (e.g. `2020-07-01` => `v20200701`).
func (g *packageGenerator) apiVersion() string {
	return strings.ReplaceAll(g.swagger.Info.Version, "-", "")
}

// providerToModule produces the module name from the provider name and the API version (e.g. (`Compute`, `2020-07-01` => `compute/v20200701`).
func (g *packageGenerator) providerToModule(prov string) string {
	return fmt.Sprintf("%s/v%s", strings.ToLower(prov), g.apiVersion())
}

// mergeParameters combines the Path Item parameters with Operation parameters.
func (g *packageGenerator) mergeParameters (operation []spec.Parameter, pathItem []spec.Parameter) []spec.Parameter {
	// Open API spec for operations:
	// > If a parameter is already defined at the Path Item, the new definition will override it.
	// > A unique parameter is defined by a combination of a name and location.
	var result []spec.Parameter
	seen := map[string]bool{}
	for _, p := range operation {
		key := fmt.Sprintf("%s@%s", p.Name, p.In)
		seen[key] = true
		result = append(result, p)
	}
	for _, p := range pathItem {
		key := fmt.Sprintf("%s@%s", p.Name, p.In)
		if _, ok := seen[key]; !ok {
			result = append(result, p)
		}
	}
	return result
}

type moduleGenerator struct {
	pkg           *pschema.PackageSpec
	metadata      *provider.AzureApiMetadata
	module        string
	resourceToken string
	visitedTypes  map[string]bool
}

// normalizeName replaces a custom name input property like `accountName` or `resourceGroupName` with the standard
// `name` property.
func (m *moduleGenerator) normalizeName(path string, requestProperties *parameterBag, responseProperties *propertyBag) error {
	// Do nothing if there's no `name` in response properties - we always expect it for any resource.
	if _, ok := responseProperties.specs["name"]; !ok {
		return nil
	}

	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name := part[1 : len(part)-1]
			sdkName := name
			for _, v := range requestProperties.parameters {
				if v.Name == name {
					prop := v.Value
					if prop.SdkName != "" {
						sdkName = prop.SdkName
					}
					// We expect names to be always a required string.
					if prop.Type != "string" {
						return errors.Errorf("name property '%s' is not a string", name)
					}
					prop.SdkName = "name"
					break
				}
			}
			if !requestProperties.requiredSpecs.Has(sdkName) {
				return errors.Errorf("name property '%s' is not required", name)
			}
			if nameProp, ok := requestProperties.specs[sdkName]; ok {
				delete(requestProperties.specs, sdkName)
				requestProperties.specs["name"] = nameProp
				requestProperties.requiredSpecs.Delete(sdkName)
				requestProperties.requiredSpecs.Add("name")
				break
			} else {
				return errors.Errorf("name property '%s' not found", sdkName)
			}
		}
	}

	return nil
}

func (m *moduleGenerator) escapeCSharpNames(typeName string, resourceResponse *propertyBag) {
	for name, spec := range resourceResponse.specs {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			spec.Language = map[string]json.RawMessage{
				"csharp": rawMessage(map[string]interface{}{
					"name": fmt.Sprintf("%sValue", typeName),
				}),
			}
			resourceResponse.specs[name] = spec
		}
	}
}

func (m *moduleGenerator) genMethodParameters(parameters []spec.Parameter, ctx *openapi.ReferenceContext) (*parameterBag, error) {
	result := newParameterBag()

	for _, param := range parameters {
		param, err := ctx.ResolveParameter(param)
		if err != nil {
			return nil, err
		}

		location, _ := param.Extensions.GetString("x-ms-parameter-location")
		apiParameter := provider.AzureApiParameter{
			Name:       param.Name,
			Location:   param.In,
			Source:     location,
			IsRequired: param.Required,
			Value: &provider.AzureApiProperty{
				Type:      param.Type,
				MinLength: param.MinLength,
				MaxLength: param.MaxLength,
				Pattern:   param.Pattern,
			},
		}

		switch {
		case param.In == "header":
			continue // Header parameters aren't mapped to the SDK.
		case param.Name == "subscriptionId":
		case param.Name == "api-version":
			continue // No need to include these in the schema, they are added automatically by the provider.
		case param.In == "body":
			// The body parameter is flattened, so that all its properties become the properties of the type.
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
			apiParameter.Body = &provider.AzureApiType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}

		default:
			name := param.Name
			if clientName, ok := param.Extensions.GetString("x-ms-client-name"); ok {
				name = clientName
			}

			// Change the name to lowerCamelCase.
			name = toLowerCamel(name)
			if name != param.Name {
				apiParameter.Value.SdkName = name
			}

			propertySpec := pschema.PropertySpec{
				Description: param.Description,
				TypeSpec: pschema.TypeSpec{
					Type: param.Type,
				},
			}
			result.specs[name] = propertySpec
			if param.Required {
				result.requiredSpecs.Add(name)
			}
		}

		result.parameters = append(result.parameters, apiParameter)
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

		// Skip empty result objects.
		if len(result.specs) == 0 {
			continue
		}

		// Id is a property of the base Custom Resource, we don't want to introduce it on derived resources.
		delete(result.specs, "id")
		result.requiredSpecs.Delete("id")

		// Special case the 'properties' output property as required.
		// This should be gone when we apply flattening.
		if _, ok := result.specs["properties"]; ok {
			result.requiredSpecs.Add("properties")
		}

		result.description = responseSchema.Description
		return result, nil
	}

	return nil, errors.New("no matching 2xx response found")
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
		sdkName := name
		if clientName, ok := property.Extensions.GetString("x-ms-client-name"); ok {
			sdkName = firstToLower(clientName)
		}
		// Change the name to lowerCamelCase.
		sdkName = toLowerCamel(sdkName)

		// Flattened properties aren't modelled in the SDK explicitly: their sub-properties are merged directly to the parent.
		if flatten, ok := property.Extensions.GetBool("x-ms-client-flatten"); ok && flatten {
			resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
			if err != nil {
				return nil, err
			}

			bag, err := m.genProperties(resolvedProperty, isOutput)
			if err != nil {
				return nil, err
			}

			// Adjust every property to mark them as flattened.
			newProperties := map[string]provider.AzureApiProperty{}
			for n, value := range bag.properties {
				value.Container = name
				newProperties[n] = value
			}
			bag.properties = newProperties

			result.merge(bag)
			continue
		}

		// Skip read-only properties for input types.
		if property.ReadOnly && !isOutput {
			continue
		}

		propertySpec, err := m.genProperty(name, &property, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, err
		}

		// Skip properties that yield degenerate types (e.g., when an input type has only read-only properties).
		if propertySpec == nil {
			continue
		}

		var apiProperty provider.AzureApiProperty
		if isOutput {
			if property.ReadOnly {
				result.requiredSpecs.Add(sdkName)
			}
			apiProperty = provider.AzureApiProperty{
				Ref: propertySpec.Ref,
			}
		} else {
			resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
			if err != nil {
				return nil, err
			}
			apiProperty = provider.AzureApiProperty{
				Type:      propertySpec.Type,
				Enum:      m.getEnumValues(&property),
				Ref:       propertySpec.Ref,
				Minimum:   resolvedProperty.Minimum,
				Maximum:   resolvedProperty.Maximum,
				MinLength: resolvedProperty.MinLength,
				MaxLength: resolvedProperty.MaxLength,
				Pattern:   resolvedProperty.Pattern,
			}
		}

		if sdkName != name {
			apiProperty.SdkName = sdkName
		}
		if propertySpec.Items != nil {
			apiProperty.Items = &provider.AzureApiProperty{
				Type: propertySpec.Items.Type,
				Ref:  propertySpec.Items.Ref,
			}
		}
		result.properties[name] = apiProperty
		result.specs[sdkName] = *propertySpec
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
		if prop, ok := result.properties[name]; ok {
			if prop.SdkName != "" {
				result.requiredSpecs.Add(prop.SdkName)
			} else {
				result.requiredSpecs.Add(name)
			}
			result.requiredProperties.Add(name)
		}
	}
	return result, nil
}

func (m *moduleGenerator) getEnumValues(property *spec.Schema) (enum []string) {
	if property.Enum == nil {
		return
	}

	restrictive := true
	// If x-ms-enum is present and modelAsString is set to false, the enum is not strict, so we don't want to enforce it.
	if extension, ok := property.Extensions["x-ms-enum"]; ok {
		if modelAsString, ok := extension.(map[string]interface{})["modelAsString"]; ok {
			if v, ok := modelAsString.(bool); ok {
				restrictive = !v
			}
		}
	}
	if restrictive {
		for _, value := range property.Enum {
			if s, ok := value.(string); ok {
				enum = append(enum, s)
			}
		}
	}
	return
}

func (m *moduleGenerator) genProperty(name string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.PropertySpec, error) {
	description := schema.Description
	if description == "" {
		resolvedSchema, err := context.ResolveSchema(schema)
		if err != nil {
			return nil, err
		}

		description = resolvedSchema.Description
	}

	typeSpec, err := m.genTypeSpec(name, schema, context, isOutput)
	if err != nil {
		return nil, err
	}

	// Nil type means empty: e.g., when an input type has only read-only properties. Bubble the nil up.
	if typeSpec == nil {
		return nil, nil
	}

	propertySpec := pschema.PropertySpec{
		Description: description,
		TypeSpec:    *typeSpec,
	}

	return &propertySpec, nil
}

func (m *moduleGenerator) genTypeSpec(propertyName string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, err
	}

	// Type specification specifies either a primitive type (e.g. 'string') or a reference to a separately defined
	// strongly-typed object. Either `primitiveTypeName` or `referencedTypeName` should be filled, but not both.
	var tok, primitiveTypeName, referencedTypeName string
	if len(resolvedSchema.Properties) > 0 || len(resolvedSchema.AllOf) > 0 {
		ptr := schema.Ref.GetPointer()
		if ptr != nil && !ptr.IsEmpty() {
			tok = m.typeName(resolvedSchema.ReferenceContext, isOutput)
		} else {
			// Inline properties have no type in the Open API schema, so we use parent type's name + property name.
			tok = m.typeName(context, isOutput) + strings.Title(propertyName)
		}
	} else if len(resolvedSchema.Type) > 0 {
		primitiveTypeName = resolvedSchema.Type[0]
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

			// Don't generate a type definition for a typed object with zero properties.
			if len(props.specs) == 0 {
				delete(m.visitedTypes, tok)
				return nil, nil
			}

			m.pkg.Types[tok] = pschema.ObjectTypeSpec{
				Description: resolvedSchema.Description,
				Type:        "object",
				Properties:  props.specs,
				Required:    props.requiredSpecs.SortedValues(),
			}

			m.metadata.Types[tok] = provider.AzureApiType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}
		}
	}

	// Define the type of maps (untyped objects).
	var additionalProperties *pschema.TypeSpec
	if primitiveTypeName == "object" {
		if resolvedSchema.AdditionalProperties != nil && resolvedSchema.AdditionalProperties.Schema != nil {
			additionalProperties, err = m.genTypeSpec(propertyName, resolvedSchema.AdditionalProperties.Schema, resolvedSchema.ReferenceContext, isOutput)
			if err != nil {
				return nil, err
			}

			// Don't generate a type definition for a typed dictionary with empty value type.
			if additionalProperties == nil {
				return nil, nil
			}
		} else {
			additionalProperties = &pschema.TypeSpec{
				Ref: "pulumi.json#/Any",
			}
		}
	}

	result := pschema.TypeSpec{
		Type:                 primitiveTypeName,
		AdditionalProperties: additionalProperties,
		Ref:                  referencedTypeName,
	}

	// Resolve the element type for array types.
	if resolvedSchema.Items != nil && resolvedSchema.Items.Schema != nil {
		itemsSpec, err := m.genProperty(propertyName, resolvedSchema.Items.Schema, context, isOutput)
		if err != nil {
			return nil, err
		}

		// Don't generate a type definition for a typed array with empty item type.
		if itemsSpec == nil {
			return nil, nil
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

// isPrimitiveType returns false for object and array, and true for all other types defined in
// https://swagger.io/docs/specification/data-models/data-types/
func isPrimitiveType(typeName string) bool {
	switch typeName {
	case "string", "number", "integer", "boolean":
		return true
	case "object", "array":
		return false
	default:
		panic(fmt.Sprintf("Unknown OpenAPI type %s", typeName))
	}
}

// parameterBag keeps the schema and metadata parameters for a single resource or invocation.
type parameterBag struct {
	description   string
	specs         map[string]pschema.PropertySpec
	parameters    []provider.AzureApiParameter
	requiredSpecs codegen.StringSet
}

func newParameterBag() *parameterBag {
	return &parameterBag{
		specs:         map[string]pschema.PropertySpec{},
		requiredSpecs: codegen.NewStringSet(),
	}
}

func (bag *parameterBag) merge(other *propertyBag) {
	for key, value := range other.specs {
		bag.specs[key] = value
	}
	for key, _ := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
}

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	properties         map[string]provider.AzureApiProperty
	requiredSpecs      codegen.StringSet
	requiredProperties codegen.StringSet
}

func newPropertyBag() *propertyBag {
	return &propertyBag{
		specs:              map[string]pschema.PropertySpec{},
		properties:         map[string]provider.AzureApiProperty{},
		requiredSpecs:      codegen.NewStringSet(),
		requiredProperties: codegen.NewStringSet(),
	}
}

func (bag *propertyBag) merge(other *propertyBag) {
	for key, value := range other.specs {
		bag.specs[key] = value
	}
	for key, value := range other.properties {
		bag.properties[key] = value
	}
	for key, _ := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
	for key, _ := range other.requiredProperties {
		bag.requiredProperties.Add(key)
	}
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
