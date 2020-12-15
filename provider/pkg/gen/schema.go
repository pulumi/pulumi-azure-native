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

package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/blang/semver"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-azure-nextgen/sdk/go/azure"

// PulumiSchema will generate a Pulumi schema for the given Azure providers and resources map.
func PulumiSchema(providerMap openapi.AzureProviders) (*pschema.PackageSpec, *resources.AzureAPIMetadata, map[string][]resources.AzureAPIExample, error) {
	pkg := pschema.PackageSpec{
		Name:        "azure-nextgen",
		Description: "A Next Generation Pulumi package for creating and managing Azure resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "azure", "azure-nextgen"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-azure-nextgen",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"subscriptionId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Subscription ID which should be used.",
				},
				"clientId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client ID which should be used.",
				},
				"clientSecret": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
				},
				"tenantId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Tenant ID which should be used.",
				},
				"auxiliaryTenantIds": {
					TypeSpec: pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
				},
				"environment": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.",
				},
				"clientCertificatePath": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.",
				},
				"clientCertificatePassword": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
				},
				"useMsi": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "Allowed Managed Service Identity be used for Authentication.",
				},
				"msiEndpoint": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically. ",
				},
				// Managed Tracking GUID for User-Agent.
				"partnerId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.",
				},
				"disablePulumiPartnerId": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified.",
				},
			},
		},
		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the Azure NextGen package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"subscriptionId": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_SUBSCRIPTION_ID",
						},
					},
					Description: "The Subscription ID which should be used.",
				},
				"clientId": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_CLIENT_ID",
						},
					},
					Description: "The Client ID which should be used.",
				},
				"clientSecret": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_CLIENT_SECRET",
						},
					},
					Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
				},
				"tenantId": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_TENANT_ID",
						},
					},
					Description: "The Tenant ID which should be used.",
				},
				"auxiliaryTenantIds": {
					TypeSpec: pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_AUXILIARY_TENANT_IDS",
						},
					},
				},
				"environment": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					Default:  "public",
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_ENVIRONMENT",
						},
					},
					Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.",
				},
				"clientCertificatePath": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_CLIENT_CERTIFICATE_PATH",
						},
					},
					Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.",
				},
				"clientCertificatePassword": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_CLIENT_CERTIFICATE_PASSWORD",
						},
					},
					Description: "The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
				},
				"useMsi": {
					TypeSpec: pschema.TypeSpec{Type: "boolean"},
					Default:  false,
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_USE_MSI",
						},
					},
					Description: "Allowed Managed Service Identity be used for Authentication.",
				},
				"msiEndpoint": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_MSI_ENDPOINT",
						},
					},
					Description: "The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically. ",
				},
				// Managed Tracking GUID for User-Agent.
				"partnerId": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_PARTNER_ID",
						},
					},
					Description: "A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.",
				},
				"disablePulumiPartnerId": {
					TypeSpec: pschema.TypeSpec{Type: "boolean"},
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"ARM_DISABLE_PULUMI_PARTNER_ID",
						},
					},
					Description: "This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified.",
				},
			},
		},
		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}
	metadata := resources.AzureAPIMetadata{
		Types:     map[string]resources.AzureAPIType{},
		Resources: map[string]resources.AzureAPIResource{},
		Invokes:   map[string]resources.AzureAPIInvoke{},
	}

	csharpVersionReplacer := strings.NewReplacer("privatepreview", "PrivatePreview", "preview", "Preview", "beta", "Beta")
	csharpNamespaces := map[string]string{
		"azure-nextgen": "AzureNextGen",
	}
	pythonModuleNames := map[string]string{}
	golangImportAliases := map[string]string{}

	var providers []string
	for prov := range providerMap {
		providers = append(providers, prov)
	}
	sort.Strings(providers)

	exampleMap := make(map[string][]resources.AzureAPIExample)
	for _, providerName := range providers {
		versionMap := providerMap[providerName]
		var versions []string
		for version := range versionMap {
			versions = append(versions, version)
		}
		sort.Strings(versions)

		for _, version := range versions {
			gen := packageGenerator{
				pkg:        &pkg,
				metadata:   &metadata,
				apiVersion: version,
				examples:   exampleMap,
			}

			// Populate C#, Python and Go module mapping.
			module := gen.providerToModule(providerName)
			csVersion := strings.Title(csharpVersionReplacer.Replace(version))
			csharpNamespaces[strings.ToLower(providerName)] = providerName
			csharpNamespaces[module] = fmt.Sprintf("%s.%s", providerName, csVersion)
			pythonModuleNames[module] = module
			golangImportAliases[filepath.Join(goBasePath, module)] = strings.ToLower(providerName)

			// Populate resources and get invokes.
			items := versionMap[version]
			var resources []string
			for resource := range items.Resources {
				resources = append(resources, resource)
			}
			sort.Strings(resources)

			for _, typeName := range resources {
				resource := items.Resources[typeName]
				gen.genResources(providerName, typeName, resource)
			}

			// Populate POST invokes.
			var invokes []string
			for invoke := range items.Invokes {
				invokes = append(invokes, invoke)
			}
			sort.Strings(invokes)

			for _, typeName := range invokes {
				invoke := items.Invokes[typeName]
				gen.genPostFunctions(providerName, typeName, invoke.Path, invoke.PathItem, invoke.Swagger)
			}
		}
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":       goBasePath,
		"packageImportAliases": golangImportAliases,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^2.0.0",
		},
		"readme": `The Azure NextGen provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=2.11.2,<3.0.0",
		},
		"usesIOClasses": true,
		"readme": `The Azure NextGen provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, exampleMap, nil
}

// Microsoft-specific API extension constants.
const (
	extensionClientFlatten      = "x-ms-client-flatten"
	extensionClientName         = "x-ms-client-name"
	extensionEnum               = "x-ms-enum"
	extensionLongRunning        = "x-ms-long-running-operation"
	extensionLongRunningDefault = "azure-async-operation"
	extensionLongRunningOpts    = "x-ms-long-running-operation-options"
	extensionLongRunningVia     = "final-state-via"
	extensionMutability         = "x-ms-mutability"
	extensionMutabilityCreate   = "create"
	extensionMutabilityUpdate   = "update"
)

type packageGenerator struct {
	pkg        *pschema.PackageSpec
	metadata   *resources.AzureAPIMetadata
	examples   map[string][]resources.AzureAPIExample
	apiVersion string
}

func (g *packageGenerator) genResources(prov, typeName string, resource *openapi.ResourceSpec) {
	module := g.providerToModule(prov)
	swagger := resource.Swagger
	path := resource.PathItem
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		prov:          prov,
		resourceName:  typeName,
		resourceToken: resourceTok,
		visitedTypes:  make(map[string]bool),
	}

	parameters := resource.Swagger.MergeParameters(path.Put.Parameters, path.Parameters)
	resourceRequest, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", resourceTok, err.Error())
		return
	}

	resourceResponse, err := gen.genResponse(path.Put.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", resourceTok, err.Error())
		return
	}

	if len(resourceResponse.specs) == 0 {
		// Response is specified empty, do not generate a resource for it.
		return
	}

	gen.escapeCSharpNames(typeName, resourceResponse)

	// Id is a property of the base Custom Resource, we don't want to introduce it on derived resources.
	delete(resourceResponse.specs, "id")
	resourceResponse.requiredSpecs.Delete("id")

	// Add an alias for each API version that has the same path in it.
	var aliases []pschema.AliasSpec
	for _, version := range resource.CompatibleVersions {
		alias := fmt.Sprintf("%s:%s/%s:%s", g.pkg.Name, strings.ToLower(prov), version, typeName)
		aliases = append(aliases, pschema.AliasSpec{Type: &alias})
	}

	resourceDescription := resourceResponse.description
	if g.apiVersion == "latest" {
		resourceDescription = fmt.Sprintf("%s\nLatest API Version: %s.", resourceDescription, swagger.Info.Version)
	}

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: resourceDescription,
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties: resourceRequest.specs,
		RequiredInputs:  resourceRequest.requiredSpecs.SortedValues(),
		Aliases:         aliases,
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, typeName)

	parameters = swagger.MergeParameters(path.Get.Parameters, path.Parameters)
	requestFunction, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	responseFunction, err := gen.genResponse(path.Get.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

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

	r := resources.AzureAPIResource{
		APIVersion:       swagger.Info.Version,
		Path:             resource.Path,
		GetParameters:    requestFunction.parameters,
		PutParameters:    resourceRequest.parameters,
		Response:         resourceResponse.properties,
		DefaultBody:      resource.DefaultBody,
		Singleton:        resource.PathItem.Delete == nil,
		PutAsyncStyle:    g.getAsyncStyle(resource.PathItem.Put),
		DeleteAsyncStyle: g.getAsyncStyle(resource.PathItem.Delete),
	}
	g.metadata.Resources[resourceTok] = r

	f := resources.AzureAPIInvoke{
		APIVersion:    swagger.Info.Version,
		Path:          resource.Path,
		GetParameters: requestFunction.parameters,
		Response:      responseFunction.properties,
	}
	g.metadata.Invokes[functionTok] = f

	g.generateExampleReferences(resourceTok, path, swagger)
}

func (g *packageGenerator) generateExampleReferences(resourceTok string, path *spec.PathItem, swagger *openapi.Spec) error {
	raw, ok := path.Put.Extensions["x-ms-examples"]
	if !ok {
		return nil
	}

	examples := raw.(map[string]interface{})

	result := make([]resources.AzureAPIExample, 0, len(examples))
	for k, v := range examples {
		resolved := v.(map[string]interface{})
		if _, ok := resolved["$ref"]; !ok {
			continue
		}

		relativePath := resolved["$ref"].(string)
		relativeURL, err := url.Parse(relativePath)
		if err != nil {
			return err
		}
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		url, err := swagger.ResolveReference(relativeURL.String())
		if err != nil {
			return err
		}

		url, err = filepath.Rel(cwd, url)
		if err != nil {
			return err
		}
		if _, err := os.Stat(url); err != nil {
			return err
		}
		result = append(result, resources.AzureAPIExample{
			Description: k,
			Location:    url,
		})
	}

	// Deterministic ordering of the examples.
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Description < result[j].Description
	})

	g.examples[resourceTok] = result
	return nil
}

// genPostFunctions defines functions for list* (listKeys, listSecrets, etc.)
// and get* (getFullUrl, getBastionShareableLink, etc.) POST endpoints.
func (g *packageGenerator) genPostFunctions(prov, typeName, path string, pathItem *spec.PathItem, swagger *openapi.Spec) {
	module := g.providerToModule(prov)
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		resourceToken: fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		prov:          prov,
		resourceName:  typeName,
		visitedTypes:  make(map[string]bool),
	}

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	parameters := swagger.MergeParameters(pathItem.Post.Parameters, pathItem.Parameters)
	request, err := gen.genMethodParameters(parameters, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(pathItem.Post.Responses.StatusCodeResponses, swagger.ReferenceContext)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	if len(response.specs) == 0 {
		// Response is specified empty, do not generate an invoke for it.
		return
	}

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

	f := resources.AzureAPIInvoke{
		APIVersion:     swagger.Info.Version,
		Path:           path,
		PostParameters: request.parameters,
		Response:       response.properties,
	}
	g.metadata.Invokes[functionTok] = f
}

// providerToModule produces the module name from the provider name and the API version (e.g. (`Compute`, `2020-07-01` => `compute/v20200701`).
func (g *packageGenerator) providerToModule(prov string) string {
	return fmt.Sprintf("%s/%s", strings.ToLower(prov), g.apiVersion)
}

func (g *packageGenerator) getAsyncStyle(op *spec.Operation) string {
	if op == nil {
		return ""
	}

	enabled, ok := op.Extensions.GetBool(extensionLongRunning)
	if !ok || !enabled {
		return ""
	}

	if options, ok := op.Extensions[extensionLongRunningOpts]; ok {
		optionsMap := options.(map[string]interface{})
		if finalStateVia, ok := optionsMap[extensionLongRunningVia].(string); ok {
			return finalStateVia
		}
	}
	return extensionLongRunningDefault
}

type moduleGenerator struct {
	pkg           *pschema.PackageSpec
	metadata      *resources.AzureAPIMetadata
	module        string
	prov          string
	resourceName  string
	resourceToken string
	visitedTypes  map[string]bool
}

func (m *moduleGenerator) escapeCSharpNames(typeName string, resourceResponse *propertyBag) {
	for name, swagger := range resourceResponse.specs {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			swagger.Language = map[string]json.RawMessage{
				"csharp": rawMessage(map[string]interface{}{
					"name": fmt.Sprintf("%sValue", typeName),
				}),
			}
			resourceResponse.specs[name] = swagger
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

		apiParameter := resources.AzureAPIParameter{
			Name:       param.Name,
			Location:   param.In,
			IsRequired: param.Required,
			Value: &resources.AzureAPIProperty{
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

			props, err := m.genProperties(resolvedSchema, false /* isOutput */, false /* isType */)
			if err != nil {
				return nil, err
			}

			result.merge(props)
			apiParameter.Body = &resources.AzureAPIType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}

		default:
			name := param.Name
			if clientName, ok := param.Extensions.GetString(extensionClientName); ok {
				name = clientName
			}

			// Change the name to lowerCamelCase.
			name = ToLowerCamel(name)
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
	for code := range statusCodeResponses {
		if code >= 300 {
			continue
		}

		codes = append(codes, code)
	}
	sort.Ints(codes)

	if len(codes) == 0 {
		return nil, errors.New("no 2xx response found")
	}

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

		if len(responseSchema.Type) > 0 && responseSchema.Type[0] == "array" {
			return nil, errors.New("array responses are not implemented yet (see issue #120)")
		}

		result, err := m.genProperties(responseSchema, true /* isOutput */, false /* isType */)
		if err != nil {
			return nil, err
		}

		// Skip empty result objects.
		if len(result.specs) == 0 {
			continue
		}

		// Special case the 'properties' output property as required.
		// This should be gone when we apply flattening.
		if _, ok := result.specs["properties"]; ok {
			result.requiredSpecs.Add("properties")
		}

		result.description = responseSchema.Description
		return result, nil
	}

	// There was at least one 2xx response defined, but it has no schema. This is not a valid resource for us,
	// skip its processing.
	return &propertyBag{}, nil
}

func (m *moduleGenerator) genProperties(resolvedSchema *openapi.Schema, isOutput, isType bool) (*propertyBag, error) {
	result := newPropertyBag()

	// Sort properties to make codegen deterministic.
	var names []string
	for name := range resolvedSchema.Properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		property := resolvedSchema.Properties[name]
		resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
		if err != nil {
			return nil, err
		}

		if name == "clusterID" && property.Description == "The deprecated identity" {
			// TODO: Get rid of this in https://github.com/pulumi/pulumi-azure-nextgen-provider/issues/331
			continue
		}

		sdkName := name
		if clientName, ok := resolvedProperty.Extensions.GetString(extensionClientName); ok {
			sdkName = firstToLower(clientName)
		}
		// Change the name to lowerCamelCase.
		sdkName = ToLowerCamel(sdkName)

		// Flattened properties aren't modelled in the SDK explicitly: their sub-properties are merged directly to the parent.
		// If the type is marked as a dictionary, ignore the extension and proceed with modeling this property explicitly.
		// We can't flatten dictionaries in a type-safe manner.
		isDict := resolvedProperty.AdditionalProperties != nil
		if flatten, ok := resolvedProperty.Extensions.GetBool(extensionClientFlatten); ok && flatten && !isDict {
			bag, err := m.genProperties(resolvedProperty, isOutput, isType)
			if err != nil {
				return nil, err
			}

			// Adjust every property to mark them as flattened.
			newProperties := map[string]resources.AzureAPIProperty{}
			for n, value := range bag.properties {
				// The order of containers is important, so we prepend the outermost name.
				value.Containers = append([]string{name}, value.Containers...)
				newProperties[n] = value
			}
			bag.properties = newProperties

			result.merge(bag)
			continue
		}

		// Skip read-only properties for input types.
		if resolvedProperty.ReadOnly && !isOutput {
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

		var apiProperty resources.AzureAPIProperty
		if isOutput {
			if resolvedProperty.ReadOnly {
				result.requiredSpecs.Add(sdkName)
			}
			apiProperty = resources.AzureAPIProperty{
				OneOf: m.getOneOfValues(propertySpec),
				Ref:   propertySpec.Ref,
				Items: m.itemTypeToProperty(propertySpec.Items),
			}
		} else {
			if m.isEnum(&propertySpec.TypeSpec) {
				apiProperty = resources.AzureAPIProperty{
					Type: "string",
					Enum: m.getEnumValues(resolvedProperty.Schema),
				}
			} else {
				apiProperty = resources.AzureAPIProperty{
					Type:      propertySpec.Type,
					OneOf:     m.getOneOfValues(propertySpec),
					Ref:       propertySpec.Ref,
					Minimum:   resolvedProperty.Minimum,
					Maximum:   resolvedProperty.Maximum,
					MinLength: resolvedProperty.MinLength,
					MaxLength: resolvedProperty.MaxLength,
					Pattern:   resolvedProperty.Pattern,
					Items:     m.itemTypeToProperty(propertySpec.Items),
				}
			}

			// Mutability extension signals whether a property can be updated in-place. Lack of the extension means
			// updatable by default.
			// Note: a non-updatable property at a subtype level (a property of a property of a resource) does not
			// mandate the replacement of the whole resource. Anyway, it's used very seldom (2 places at the time of writing).
			// Example: `StorageAccount.encryption.services.blob.keyType` is non-updatable, but a user can remove `blob`
			// and then re-add it with the new `keyType` without replacing the whole storage account (which would be
			// very disruptive).
			if mutability, ok := resolvedProperty.Extensions.GetStringSlice(extensionMutability); ok && !isType {
				operations := codegen.NewStringSet(mutability...)
				apiProperty.ForceNew = operations.Has(extensionMutabilityCreate) && !operations.Has(extensionMutabilityUpdate)
			}

			// Apply manual metadata about Force New properties.
			if m.forceNew(name) {
				apiProperty.ForceNew = true
			}
		}

		if sdkName != name {
			apiProperty.SdkName = sdkName
		}

		result.properties[name] = apiProperty
		result.specs[sdkName] = *propertySpec
	}

	for _, s := range resolvedSchema.AllOf {
		allOfSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return nil, err
		}

		allOfProperties, err := m.genProperties(allOfSchema, isOutput, isType)
		if err != nil {
			return nil, err
		}

		// For a derived type, set the discriminator property to the const value, if any.
		if allOfSchema.Discriminator != "" {
			discriminator := allOfSchema.Discriminator
			prop := allOfProperties.properties[discriminator]
			sdkDiscriminator := discriminator
			if prop.SdkName != "" {
				sdkDiscriminator = prop.SdkName
			}

			propSpec := allOfProperties.specs[sdkDiscriminator]
			discriminatorValue := resolvedSchema.ReferenceContext.ReferenceName
			if v, ok := resolvedSchema.Extensions.GetString("x-ms-discriminator-value"); ok {
				discriminatorValue = v
			}
			propSpec.Const = discriminatorValue
			propSpec.Type = "string"
			propSpec.Ref = ""
			propSpec.OneOf = nil

			// Add the discriminator value to the property description to help users fill it.
			propSpec.Description = fmt.Sprintf("%s\nExpected value is '%s'.", propSpec.Description, discriminatorValue)

			allOfProperties.specs[sdkDiscriminator] = propSpec
			prop.Const = discriminatorValue
			allOfProperties.properties[discriminator] = prop
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

// forceNew return true if a property with a given name requires a replacement in the resource
// that is currently being generated, based on forceNewMap.
func (m *moduleGenerator) forceNew(propertyName string) bool {
	if resourceMap, ok := forceNewMap[m.prov]; ok {
		if properties, ok := resourceMap[m.resourceName]; ok {
			if properties.Has(propertyName) {
				return true
			}
		}
	}

	return false
}

func (m *moduleGenerator) getEnumValues(property *spec.Schema) (enum []string) {
	if property.Enum == nil {
		return
	}

	restrictive := false
	// If x-ms-enum is present and modelAsString is set to false, the enum is strict, so we want to enforce it.
	// Otherwise, treat the enum as a string.
	if extension, ok := property.Extensions[extensionEnum]; ok {
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

func (m *moduleGenerator) getOneOfValues(property *pschema.PropertySpec) (values []string) {
	for _, value := range property.OneOf {
		values = append(values, value.Ref)
	}
	return
}

// itemTypeToProperty converts a type of an element in an array to a corresponding API property
// definition of that element type. It only converts the relevant subset of properties, and does
// so recursively.
func (m *moduleGenerator) itemTypeToProperty(typ *schema.TypeSpec) *resources.AzureAPIProperty {
	if typ == nil {
		return nil
	}

	if m.isEnum(typ) {
		return &resources.AzureAPIProperty{Type: "string"}
	}

	var oneOf []string
	for _, subType := range typ.OneOf {
		if subType.Ref != "" {
			oneOf = append(oneOf, subType.Ref)
		} else {
			oneOf = append(oneOf, subType.Type)
		}
	}

	return &resources.AzureAPIProperty{
		Type:  typ.Type,
		Ref:   typ.Ref,
		OneOf: oneOf,
		Items: m.itemTypeToProperty(typ.Items),
	}
}

func (m *moduleGenerator) isEnum(typ *schema.TypeSpec) bool {
	for _, subType := range typ.OneOf {
		if m.isEnum(&subType) {
			return true
		}
	}

	if typ.Ref == "" {
		return false
	}

	typeName := strings.TrimPrefix(typ.Ref, "#/types/")
	refTyp := m.pkg.Types[typeName]
	return len(refTyp.Enum) > 0
}

func getPropertyDescription(schema *spec.Schema, context *openapi.ReferenceContext) (string, error) {
	description := schema.Description
	if description == "" {
		resolvedSchema, err := context.ResolveSchema(schema)
		if err != nil {
			return "", err
		}

		description = resolvedSchema.Description
	}
	return description, nil
}

func (m *moduleGenerator) genProperty(name string, schema *spec.Schema, context *openapi.ReferenceContext, isOutput bool) (*pschema.PropertySpec, error) {
	description, err := getPropertyDescription(schema, context)
	if err != nil {
		return nil, err
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
	var primitiveTypeName string
	if len(resolvedSchema.Type) > 0 {
		primitiveTypeName = resolvedSchema.Type[0]
	}

	// If this is an enum and it's an input type, generate the enum type definition.
	enumExtension, isEnum := resolvedSchema.Extensions[extensionEnum].(map[string]interface{})
	if !isOutput && isEnum && resolvedSchema.Enum != nil &&
		// There are some boolean- and number- based definitions but they aren't allowed by the Azure specs.
		// Ignore both to preserve over-the-wire format even if they say it's an enum to be modelled as a string.
		primitiveTypeName != "boolean" && primitiveTypeName != "number" {
		return m.genEnumType(schema, context, enumExtension)
	}

	switch {
	case len(resolvedSchema.Properties) > 0 || len(resolvedSchema.AllOf) > 0:
		ptr := schema.Ref.GetPointer()
		var tok string
		if ptr != nil && !ptr.IsEmpty() {
			tok = m.typeName(resolvedSchema.ReferenceContext, isOutput)
		} else {
			// Inline properties have no type in the Open API schema, so we use parent type's name + property name.
			tok = m.typeName(context, isOutput) + strings.Title(propertyName)
		}

		// If an object type is referenced, add its definition to the type map.
		discriminatedType, ok, err := m.genDiscriminatedType(resolvedSchema, isOutput)
		if err != nil {
			return nil, err
		}
		if ok {
			return discriminatedType, nil
		}

		referencedTypeName := fmt.Sprintf("#/types/%s", tok)

		if _, ok := m.visitedTypes[tok]; !ok {
			m.visitedTypes[tok] = true

			props, err := m.genProperties(resolvedSchema, isOutput, true /* isType */)
			if err != nil {
				return nil, err
			}

			// Don't generate a type definition for a typed object with zero properties.
			if len(props.specs) == 0 {
				delete(m.visitedTypes, tok)
				return nil, nil
			}

			m.pkg.Types[tok] = pschema.ComplexTypeSpec{
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Description: resolvedSchema.Description,
					Type:        "object",
					Properties:  props.specs,
					Required:    props.requiredSpecs.SortedValues(),
				},
			}

			m.metadata.Types[tok] = resources.AzureAPIType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}
		}
		return &pschema.TypeSpec{
			Type: "object",
			Ref:  referencedTypeName,
		}, nil

	case len(resolvedSchema.Enum) > 0:
		if primitiveTypeName == "" || primitiveTypeName == "object" {
			// Default Enum properties to strings if the type isn't specified.
			primitiveTypeName = "string"
		}
		return &pschema.TypeSpec{Type: primitiveTypeName}, nil

	case resolvedSchema.Items != nil && resolvedSchema.Items.Schema != nil:
		// Resolve the element type for array types.
		itemsSpec, err := m.genProperty(propertyName, resolvedSchema.Items.Schema, context, isOutput)
		if err != nil {
			return nil, err
		}

		// Don't generate a type definition for a typed array with empty item type.
		if itemsSpec == nil {
			return nil, nil
		}

		return &pschema.TypeSpec{
			Type:  "array",
			Items: &itemsSpec.TypeSpec,
		}, nil

	case resolvedSchema.AdditionalProperties != nil && resolvedSchema.AdditionalProperties.Schema != nil:
		// Define the type of maps (untyped objects).
		additionalProperties, err := m.genTypeSpec(propertyName, resolvedSchema.AdditionalProperties.Schema, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, err
		}

		// Use a generic 'object' value type for a dictionary with empty value type.
		if additionalProperties == nil {
			additionalProperties = &pschema.TypeSpec{
				Ref: "pulumi.json#/Any",
			}
		}

		return &pschema.TypeSpec{
			Type:                 "object",
			AdditionalProperties: additionalProperties,
		}, nil

	case primitiveTypeName == "" || primitiveTypeName == "object":
		// Open API v2:
		// > A schema without a type matches any data type – numbers, strings, objects, and so on.
		// Azure uses a 'naked' object type for the same purpose: to specify 'any' type.
		return &pschema.TypeSpec{
			Ref: "pulumi.json#/Any",
		}, nil

	default:
		// Primitive type ('string', 'integer', etc.)
		return &pschema.TypeSpec{Type: primitiveTypeName}, nil
	}
}

// genEnumType generates the enum type.
func (m *moduleGenerator) genEnumType(schema *spec.Schema, context *openapi.ReferenceContext, enumExtension map[string]interface{}) (*pschema.TypeSpec, error) {
	resolvedSchema, err := context.ResolveSchema(schema)
	if err != nil {
		return nil, err
	}

	description, err := getPropertyDescription(schema, context)
	if err != nil {
		return nil, err
	}

	var enumName string
	if name, ok := enumExtension["name"].(string); ok {
		enumName = ToUpperCamel(name)
	} else {
		return nil, fmt.Errorf("name key missing from enum metadata")
	}

	tok := fmt.Sprintf("%s:%s:%s", m.pkg.Name, m.module, enumName)

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []*pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: description,
			Type:        "string", // This provider only has string enums
		},
	}
	if values, ok := enumExtension["values"].([]interface{}); ok {
		for _, val := range values {
			if val, ok := val.(map[string]interface{}); ok {
				enumVal := &pschema.EnumValueSpec{
					Value: fmt.Sprintf("%v", val["value"]),
				}
				if name, ok := val["name"].(string); ok {
					enumVal.Name = name
				}
				if description, ok := val["description"].(string); ok {
					enumVal.Description = description
				}
				enumSpec.Enum = append(enumSpec.Enum, enumVal)
			}
		}
	} else {
		for _, val := range resolvedSchema.Enum {
			enumVal := &pschema.EnumValueSpec{Value: fmt.Sprintf("%v", val)}
			enumSpec.Enum = append(enumSpec.Enum, enumVal)
		}
	}

	m.pkg.Types[tok] = *enumSpec

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)

	modelAsString, ok := enumExtension["modelAsString"].(bool)
	if !ok || modelAsString {
		return &pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{
				{Type: "string"},
				{Ref: referencedTypeName},
			},
		}, nil
	}

	return &pschema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
}

// genDiscriminatedType generates polymorphic types (base type and subtypes) if the schema specifies a discriminator property.
// If no error occurs, the bool result indicates whether a discriminated (union) type is detected. If true, the TypeSpec
// result points to the specification of the union type.
func (m *moduleGenerator) genDiscriminatedType(resolvedSchema *openapi.Schema, isOutput bool) (*pschema.TypeSpec, bool, error) {
	if resolvedSchema.Discriminator == "" {
		return nil, false, nil
	}

	discriminator := resolvedSchema.Discriminator
	if _, ok := resolvedSchema.Properties[discriminator]; !ok {
		return nil, false, nil
	}

	prop := resolvedSchema.Properties[discriminator]
	if prop.ReadOnly && !isOutput {
		return nil, false, nil
	}

	var oneOf []pschema.TypeSpec
	subtypes := resolvedSchema.FindSubtypes()
	for _, subtype := range subtypes {
		typ, err := m.genTypeSpec("", subtype, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, false, err
		}
		oneOf = append(oneOf, *typ)
	}

	switch len(oneOf) {
	case 0:
		// Type specifies a discriminator but doesn't have actual subtypes. Ignore the discriminator in this case.
		return nil, false, nil
	case 1:
		// There is just one subtype specified: use it as a definite type.
		return &oneOf[0], true, nil
	default:
		// Union type for two or more types.
		return &pschema.TypeSpec{OneOf: oneOf}, true, nil
	}
}

func (m *moduleGenerator) typeName(ctx *openapi.ReferenceContext, isOutput bool) string {
	suffix := ""
	if isOutput {
		suffix = "Response"
	}
	referenceName := ToUpperCamel(MakeLegalIdentifier(ctx.ReferenceName))
	return fmt.Sprintf("azure-nextgen:%s:%s%s", m.module, referenceName, suffix)
}

// parameterBag keeps the schema and metadata parameters for a single resource or invocation.
type parameterBag struct {
	description   string
	specs         map[string]pschema.PropertySpec
	parameters    []resources.AzureAPIParameter
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
	for key := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
}

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	properties         map[string]resources.AzureAPIProperty
	requiredSpecs      codegen.StringSet
	requiredProperties codegen.StringSet
}

func newPropertyBag() *propertyBag {
	return &propertyBag{
		specs:              map[string]pschema.PropertySpec{},
		properties:         map[string]resources.AzureAPIProperty{},
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
	for key := range other.requiredSpecs {
		bag.requiredSpecs.Add(key)
	}
	for key := range other.requiredProperties {
		bag.requiredProperties.Add(key)
	}
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}

// InMemoryPackageLoader prevents having to fetch the schema from
// the provider every time which significantly speeds up codegen.
func InMemoryPackageLoader(pkgs map[string]*schema.Package) schema.Loader {
	return &inMemoryLoader{pkgs: pkgs}
}

type inMemoryLoader struct {
	pkgs map[string]*schema.Package
}

func (l *inMemoryLoader) LoadPackage(pkg string, _ *semver.Version) (*schema.Package, error) {
	if p, ok := l.pkgs[pkg]; ok {
		return p, nil
	}

	return nil, errors.Errorf("package %s not found in the in-memory map", pkg)
}
