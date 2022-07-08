// Copyright 2016-2021, Pulumi Corporation.
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
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/segmentio/encoding/json"

	"github.com/blang/semver"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-azure-native/sdk/go/azure"

// PulumiSchema will generate a Pulumi schema for the given Azure providers and resources map.
func PulumiSchema(providerMap openapi.AzureProviders) (*pschema.PackageSpec, *resources.AzureAPIMetadata, map[string][]resources.AzureAPIExample, error) {
	pkg := pschema.PackageSpec{
		Name:        "azure-native",
		Description: "A native Pulumi package for creating and managing Azure resources.",
		DisplayName: "Azure Native",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "azure", "azure-native", "category/cloud", "kind/native"},
		Homepage:    "https://pulumi.com",
		Publisher:   "Pulumi",
		Repository:  "https://github.com/pulumi/pulumi-azure-native",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"subscriptionId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Subscription ID which should be used.",
				},
				"clientId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client ID which should be used.",
					Secret:      true,
				},
				"clientSecret": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
					Secret:      true,
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
					Secret:      true,
				},
				"useMsi": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "Allowed Managed Service Identity be used for Authentication.",
					Default:     false,
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
				Description: "The provider type for the native Azure package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"subscriptionId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Subscription ID which should be used.",
				},
				"clientId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client ID which should be used.",
					Secret:      true,
				},
				"clientSecret": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
					Secret:      true,
				},
				"tenantId": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Tenant ID which should be used.",
				},
				"auxiliaryTenantIds": {
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					Description: "Any additional Tenant IDs which should be used for authentication.",
				},
				"environment": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Default:     "public",
					Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.",
				},
				"clientCertificatePath": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.",
				},
				"clientCertificatePassword": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
					Secret:      true,
				},
				"useMsi": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "Allowed Managed Service Identity be used for Authentication.",
					Default:     false,
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
		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]pschema.RawMessage{},
	}
	metadata := resources.AzureAPIMetadata{
		Types:     map[string]resources.AzureAPIType{},
		Resources: map[string]resources.AzureAPIResource{},
		Invokes:   map[string]resources.AzureAPIInvoke{},
	}

	csharpVersionReplacer := strings.NewReplacer("privatepreview", "PrivatePreview", "preview", "Preview", "beta", "Beta")
	csharpNamespaces := map[string]string{
		"azure-native": "AzureNative",
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
			csharpNamespaces[strings.ToLower(providerName)] = providerName
			if version != "" {
				csVersion := strings.Title(csharpVersionReplacer.Replace(version))
				csharpNamespaces[module] = fmt.Sprintf("%s.%s", providerName, csVersion)
			}
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
				err := gen.genResources(providerName, typeName, resource)
				if err != nil {
					return nil, nil, nil, err
				}
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

	err := genMixins(&pkg, &metadata)
	if err != nil {
		return nil, nil, nil, err
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":                 goBasePath,
		"packageImportAliases":           golangImportAliases,
		"generateResourceContainerTypes": false,
		"disableInputTypeRegistrations":  true,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^3.0.0",
		},
		"readme": `The native Azure provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		"usesIOClasses": true,
		"readme": `The native Azure provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "3.*",
			"System.Collections.Immutable": "5.0.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, exampleMap, nil
}

func genMixins(pkg *pschema.PackageSpec, metadata *resources.AzureAPIMetadata) error {
	// Mixin 'getClientConfig' to read current configuration values.
	if _, has := pkg.Functions["azure-native:authorization:getClientConfig"]; has {
		return errors.New("Invoke 'azure-native:authorization:getClientConfig' is already defined")
	}
	pkg.Functions["azure-native:authorization:getClientConfig"] = schema.FunctionSpec{
		Description: "Use this function to access the current configuration of the native Azure provider.",
		Outputs: &schema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientConfig.",
			Properties: map[string]schema.PropertySpec{
				"clientId": {
					Description: "Azure Client ID (Application Object ID).",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"objectId": {
					Description: "Azure Object ID of the current user or service principal.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"subscriptionId": {
					Description: "Azure Subscription ID",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"tenantId": {
					Description: "Azure Tenant ID",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			Type:     "object",
			Required: []string{"clientId", "objectId", "subscriptionId", "tenantId"},
		},
	}

	// Mixin 'getClientToken' to obtain an Azure auth token.
	if _, has := pkg.Functions["azure-native:authorization:getClientToken"]; has {
		return errors.New("Invoke 'azure-native:authorization:getClientToken' is already defined")
	}
	pkg.Functions["azure-native:authorization:getClientToken"] = schema.FunctionSpec{
		Description: "Use this function to get an Azure authentication token for the current login context.",
		Inputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"endpoint": {
					Description: "Optional authentication endpoint. Defaults to the endpoint of Azure Resource Manager.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			Type: "object",
		},
		Outputs: &schema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientToken.",
			Properties: map[string]schema.PropertySpec{
				"token": {
					Description: "OAuth token for Azure Management API and SDK authentication.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			Type:     "object",
			Required: []string{"token"},
		},
	}

	// Mixin all the custom resources that define schema and/or metadata.
	for tok, r := range resources.SchemaMixins() {
		if _, has := pkg.Resources[tok]; has {
			return errors.Errorf("Resource %q is already defined", tok)
		}
		pkg.Resources[tok] = r
	}
	for tok, t := range resources.SchemaTypeMixins() {
		if _, has := pkg.Types[tok]; has {
			return errors.Errorf("Type %q is already defined", tok)
		}
		pkg.Types[tok] = t
	}
	for tok, r := range resources.MetaMixins() {
		if _, has := metadata.Resources[tok]; has {
			return errors.Errorf("Metadata %q is already defined", tok)
		}
		metadata.Resources[tok] = r
	}

	// Add a note regarding WorkspaceSqlAadAdmin creation.
	workspaceSqlAadAdmin := pkg.Resources["azure-native:synapse:WorkspaceSqlAadAdmin"]
	workspaceSqlAadAdmin.Description += "\n\nNote: SQL AAD Admin is configured automatically during workspace creation and assigned to the current user. One can't add more admins with this resource unless you manually delete the current SQL AAD Admin."
	pkg.Resources["azure-native:synapse:WorkspaceSqlAadAdmin"] = workspaceSqlAadAdmin

	return nil
}

// Microsoft-specific API extension constants.
const (
	extensionClientFlatten      = "x-ms-client-flatten"
	extensionClientName         = "x-ms-client-name"
	extensionDiscriminatorValue = "x-ms-discriminator-value"
	extensionEnum               = "x-ms-enum"
	extensionLongRunning        = "x-ms-long-running-operation"
	extensionLongRunningDefault = "azure-async-operation"
	extensionLongRunningOpts    = "x-ms-long-running-operation-options"
	extensionLongRunningVia     = "final-state-via"
	extensionMutability         = "x-ms-mutability"
	extensionMutabilityCreate   = "create"
	extensionMutabilityRead     = "read"
	extensionMutabilityUpdate   = "update"
	extensionParameterLocation  = "x-ms-parameter-location"
)

type packageGenerator struct {
	pkg        *pschema.PackageSpec
	metadata   *resources.AzureAPIMetadata
	examples   map[string][]resources.AzureAPIExample
	apiVersion string
}

func (g *packageGenerator) genResources(prov, typeName string, resource *openapi.ResourceSpec) error {
	// A single API path can be modelled as several resources if it accepts a polymorphic payload:
	// i.e., when the request body is a discriminated union type of several object types. Pulumi
	// schema doesn't support polymorphic (OneOf) resources, so the provider creates a separate resource
	// per each union case. We call them "variants" in the code below.
	variants, err := g.findResourceVariants(resource)
	if err != nil {
		return errors.Wrapf(err, "resource %s.%s", prov, typeName)
	}
	var variantNames []string
	for _, d := range variants {
		err = g.genResourceVariant(prov, d)
		if err != nil {
			return err
		}
		variantNames = append(variantNames, d.typeName)
	}
	mainResource := &resourceVariant{
		ResourceSpec: resource,
		typeName:     typeName,
	}
	if resource.DeprecationMessage != "" {
		mainResource.deprecationMessage = resource.DeprecationMessage
	} else if len(variants) > 0 {
		// If variants are found, we still want to generate the "base" resource but only for compatibility
		// reason. We should remove those resources in 2.0, as most of them simply don't work.
		mainResource.deprecationMessage =
			fmt.Sprintf("Please use one of the variants: %s.", strings.Join(variantNames, ", "))
	}
	return g.genResourceVariant(prov, mainResource)
}

// resourceVariant points to request body's and response's schemas of a resource which is one of the variants
// of a resource related to an API path.
type resourceVariant struct {
	*openapi.ResourceSpec
	typeName           string
	body               *openapi.Schema
	response           *openapi.Schema
	deprecationMessage string
}

// findResourceVariants searches for discriminated unions in the resource's API specs and returns
// a list of those variants. An empty list is returned for non-discriminated types.
func (g *packageGenerator) findResourceVariants(resource *openapi.ResourceSpec) ([]*resourceVariant, error) {
	updateOp := resource.PathItem.Put
	if resource.PathItem.Put == nil {
		updateOp = resource.PathItem.Patch
	}

	// Check if the body schema has a discriminator, return otherwise.
	bodySchema, err := getRequestBodySchema(resource.Swagger.ReferenceContext, updateOp.Parameters)
	if bodySchema == nil || err != nil {
		return nil, err
	}

	if bodySchema.Discriminator == "" {
		return nil, nil
	}

	// Find the base schema of the response. Variants would all derive from this base schema.
	responseSchema, err := getResponseSchema(resource.Swagger.ReferenceContext, updateOp.Responses.StatusCodeResponses)
	if responseSchema == nil || err != nil {
		return nil, err
	}

	// Built the map of response schemas per discriminator value or the default reponse schema if no discriminator.
	responses := map[string]*openapi.Schema{}
	var defaultResponse *openapi.Schema
	if responseSchema.Discriminator != "" {
		responseSubtypes, err := responseSchema.FindSubtypes()
		if err != nil {
			return nil, err
		}
		for _, subtype := range responseSubtypes {
			resolvedSubSchema, err := responseSchema.ResolveSchema(subtype)
			if err != nil {
				return nil, err
			}

			discriminatorValue := strings.ToLower(getDiscriminatorValue(resolvedSubSchema))
			responses[discriminatorValue] = resolvedSubSchema
		}
	} else {
		defaultResponse = responseSchema
	}

	bodySubtypes, err := bodySchema.FindSubtypes()
	if err != nil {
		return nil, err
	}

	// For each body schema subtype, find the corresponding response schema subtype and return the pair
	// as a resource variant.
	var result []*resourceVariant
	for _, subtype := range bodySubtypes {
		resolvedSubSchema, err := bodySchema.ResolveSchema(subtype)
		if err != nil {
			return nil, err
		}

		discriminatorValue := strings.ToLower(getDiscriminatorValue(resolvedSubSchema))

		response := defaultResponse
		if v, ok := responses[discriminatorValue]; ok {
			response = v
		}

		result = append(result, &resourceVariant{
			ResourceSpec: resource,
			typeName:     resources.DiscriminatedResourceName(resolvedSubSchema.ReferenceName),
			body:         resolvedSubSchema,
			response:     response,
		})
	}
	return result, nil
}

func (g *packageGenerator) genResourceVariant(prov string, resource *resourceVariant) error {
	module := g.providerToModule(prov)
	swagger := resource.Swagger
	path := resource.PathItem
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, resource.typeName)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:           g.pkg,
		metadata:      g.metadata,
		module:        module,
		prov:          prov,
		resourceName:  resource.typeName,
		resourceToken: resourceTok,
		visitedTypes:  make(map[string]bool),
		inlineTypes:   map[*openapi.ReferenceContext]codegen.StringSet{},
	}

	updateOp := path.Put
	updateMethod := ""
	if path.Put == nil {
		updateOp = path.Patch
		updateMethod = "PATCH"
	}

	resourceResponse, err := gen.genResponse(updateOp.Responses.StatusCodeResponses, swagger.ReferenceContext, resource.response)
	if err != nil {
		return errors.Wrapf(err, "failed to generate '%s': response type", resourceTok)
	}

	if resourceResponse == nil || len(resourceResponse.specs) == 0 {
		// Response is specified empty, do not generate a resource for it.
		return nil
	}

	parameters := resource.Swagger.MergeParameters(updateOp.Parameters, path.Parameters)
	autoNamer := resources.NewAutoNamer(resource.Path)

	resourceRequest, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, &autoNamer, resource.body)
	if err != nil {
		return errors.Wrapf(err, "failed to generate '%s': request type", resourceTok)
	}

	gen.escapeCSharpNames(resource.typeName, resourceResponse)

	// Id is a property of the base Custom Resource, we don't want to introduce it on derived resources.
	delete(resourceResponse.specs, "id")
	resourceResponse.requiredSpecs.Delete("id")

	// Add an alias for each API version that has the same path in it.
	var aliases []pschema.AliasSpec
	for _, version := range resource.CompatibleVersions {
		moduleName := providerApiToModule(prov, version)
		alias := fmt.Sprintf("%s:%s:%s", g.pkg.Name, moduleName, resource.typeName)
		aliases = append(aliases, pschema.AliasSpec{Type: &alias})
	}

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: g.formatDescription(resourceResponse, swagger.Info),
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties:    resourceRequest.specs,
		RequiredInputs:     resourceRequest.requiredSpecs.SortedValues(),
		Aliases:            aliases,
		DeprecationMessage: resource.deprecationMessage,
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, resource.typeName)

	var readOp *spec.Operation
	switch {
	case resource.PathItemList != nil:
		if resource.PathItemList.Post != nil {
			readOp = resource.PathItemList.Post
		} else {
			readOp = resource.PathItemList.Get
		}
	case path.Get == nil:
		readOp = path.Head
	default:
		readOp = path.Get
	}

	parameters = swagger.MergeParameters(readOp.Parameters, path.Parameters)
	requestFunction, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, nil, resource.body)
	if err != nil {
		return errors.Wrapf(err, "failed to generate '%s': request type", functionTok)
	}
	responseFunction, err := gen.genResponse(readOp.Responses.StatusCodeResponses, swagger.ReferenceContext, resource.response)
	if err != nil {
		return errors.Wrapf(err, "failed to generate '%s': response type", functionTok)
	}

	if path.Get != nil && responseFunction != nil {
		functionSpec := pschema.FunctionSpec{
			Description:        g.formatDescription(resourceResponse, swagger.Info),
			DeprecationMessage: resource.deprecationMessage,
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

		f := resources.AzureAPIInvoke{
			APIVersion:    swagger.Info.Version,
			Path:          resource.Path,
			GetParameters: requestFunction.parameters,
			Response:      responseFunction.properties,
		}
		g.metadata.Invokes[functionTok] = f
	}

	var readMethod, readPath string
	switch {
	case resource.PathItemList != nil:
		readPath = "/list"
		if resource.PathItemList.Post != nil {
			readMethod = "POST"
		}
	case resource.PathItem.Get == nil:
		readMethod = "HEAD"
	}

	r := resources.AzureAPIResource{
		APIVersion:           swagger.Info.Version,
		Path:                 resource.Path,
		UpdateMethod:         updateMethod,
		PutParameters:        resourceRequest.parameters,
		Response:             resourceResponse.properties,
		DefaultBody:          resource.DefaultBody,
		Singleton:            resource.PathItem.Delete == nil,
		PutAsyncStyle:        g.getAsyncStyle(updateOp),
		DeleteAsyncStyle:     g.getAsyncStyle(resource.PathItem.Delete),
		ReadMethod:           readMethod,
		ReadPath:             readPath,
		AutoLocationDisabled: resources.AutoLocationDisabled(resource.Path),
	}
	g.metadata.Resources[resourceTok] = r

	g.generateExampleReferences(resourceTok, path, swagger)
	return nil
}

func (g *packageGenerator) generateExampleReferences(resourceTok string, path *spec.PathItem, swagger *openapi.Spec) error {
	op := path.Put
	if path.Put == nil {
		op = path.Patch
	}

	raw, ok := op.Extensions["x-ms-examples"]
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
		inlineTypes:   map[*openapi.ReferenceContext]codegen.StringSet{},
	}

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName)

	parameters := swagger.MergeParameters(pathItem.Post.Parameters, pathItem.Parameters)
	request, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, nil, nil)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(pathItem.Post.Responses.StatusCodeResponses, swagger.ReferenceContext, nil)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	if response == nil || len(response.specs) == 0 {
		// Response is specified empty, do not generate an invoke for it.
		return
	}

	functionSpec := pschema.FunctionSpec{
		Description: g.formatDescription(response, swagger.Info),
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
	return providerApiToModule(prov, g.apiVersion)
}
func providerApiToModule(prov, apiVersion string) string {
	if apiVersion == "" {
		return strings.ToLower(prov)
	}
	return fmt.Sprintf("%s/%s", strings.ToLower(prov), apiVersion)
}

func (g *packageGenerator) formatDescription(response *propertyBag, info *spec.Info) string {
	if g.apiVersion == "" {
		return fmt.Sprintf("%s\nAPI Version: %s.", response.description, info.Version)
	}
	return response.description
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

func getRequestBodySchema(ctx *openapi.ReferenceContext, parameters []spec.Parameter) (*openapi.Schema, error) {
	for _, p := range parameters {
		param, err := ctx.ResolveParameter(p)
		if err != nil {
			return nil, err
		}
		if param.In != "body" {
			continue
		}

		resolvedSchema, err := param.ResolveSchema(param.Schema)
		if err != nil {
			return nil, err
		}

		return resolvedSchema, nil
	}

	return nil, nil
}

func getResponseSchema(ctx *openapi.ReferenceContext, statusCodeResponses map[int]spec.Response) (*openapi.Schema, error) {
	var codes []int
	for code := range statusCodeResponses {
		if code >= 300 || code < 200 {
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

		return response.ResolveSchema(response.Schema)

	}
	return nil, nil
}

type moduleGenerator struct {
	pkg           *pschema.PackageSpec
	metadata      *resources.AzureAPIMetadata
	module        string
	prov          string
	resourceName  string
	resourceToken string
	visitedTypes  map[string]bool
	inlineTypes   map[*openapi.ReferenceContext]codegen.StringSet
}

func (m *moduleGenerator) escapeCSharpNames(typeName string, resourceResponse *propertyBag) {
	for name, swagger := range resourceResponse.specs {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			swagger.Language = map[string]pschema.RawMessage{
				"csharp": rawMessage(map[string]interface{}{
					"name": fmt.Sprintf("%sValue", typeName),
				}),
			}
			resourceResponse.specs[name] = swagger
		}
	}
}

func (m *moduleGenerator) genMethodParameters(parameters []spec.Parameter, ctx *openapi.ReferenceContext,
	namer *resources.AutoNamer, bodySchema *openapi.Schema) (*parameterBag, error) {
	result := newParameterBag()
	var autoNamedSpec string

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

		// Provider has a value defined for subscription ID, so, by default, don't include it into SDKs unless
		// the API requires that explicitly.
		providerHasDefaultValue := param.Name == "subscriptionId"

		switch {
		case param.In == "header":
			continue // Header parameters aren't mapped to the SDK.
		case providerHasDefaultValue && !isMethodParameter(param):
			// Don't include values with a provider-configured value to the schema unless it's a method parameter.
		case param.Name == "api-version":
			continue // No need to include API version in the schema or meta, it is added automatically by the provider.
		case param.In == "body":
			if param.Schema == nil {
				return nil, errors.Errorf("no schema for body parameter '%s'", param.Name)
			}

			if bodySchema == nil {
				bodySchema, err = param.ResolveSchema(param.Schema)
				if err != nil {
					return nil, errors.Wrapf(err, "body parameter '%s'", param.Name)
				}
			}

			// The body parameter is flattened, so that all its properties become the properties of the type.
			props, err := m.genProperties(bodySchema, false /* isOutput */, false /* isType */)
			if err != nil {
				return nil, err
			}

			// Top-level location is never required: it can be derived from a config value or the parent resource group.
			props.requiredSpecs.Delete("location")

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

			// Check each parameter for auto-naming.
			if namer != nil {
				if kind, ok := namer.AutoName(param.Name, param.Format); ok {
					apiParameter.Value.AutoName = kind
					autoNamedSpec = name
				}
			}

			result.specs[name] = propertySpec
			if param.Required && !providerHasDefaultValue {
				result.requiredSpecs.Add(name)
			}
		}

		result.parameters = append(result.parameters, apiParameter)
	}

	// Remove the auto-name property from the list of required specs. This happens after we processed all parameters
	// to account for several cases where the same name is both a method parameter and a body property.
	result.requiredSpecs.Delete(autoNamedSpec)

	return result, nil
}

// isMethodParameter returns true if a parameter is marked with an extension location=method.
func isMethodParameter(param *openapi.Parameter) bool {
	if value, ok := param.Extensions.GetString(extensionParameterLocation); ok {
		return value == "method"
	}

	return false
}

func (m *moduleGenerator) genResponse(statusCodeResponses map[int]spec.Response, ctx *openapi.ReferenceContext,
	responseSchema *openapi.Schema) (*propertyBag, error) {

	if responseSchema == nil {
		v, err := getResponseSchema(ctx, statusCodeResponses)
		if v == nil || err != nil {
			return nil, err
		}
		responseSchema = v
	}

	if len(responseSchema.Type) > 0 && responseSchema.Type[0] == "array" {
		// Array responses are not implemented yet (see issue #120).
		return nil, nil
	}

	result, err := m.genProperties(responseSchema, true /* isOutput */, false /* isType */)
	if err != nil {
		return nil, err
	}

	// Special case the 'properties' output property as required.
	// This should be gone when we apply flattening.
	if _, ok := result.specs["properties"]; ok {
		result.requiredSpecs.Add("properties")
	}

	result.description = responseSchema.Description
	return result, nil
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

		if name == "etag" && !isType && !isOutput {
			// ETags may be defined as inputs to PUT endpoints, but we should never model them as
			// resource inputs because they are a protocol implementation detail rather than
			// a meaningful desired-state property.
			continue
		}
		if name == "clusterID" && property.Description == "The deprecated identity" {
			// TODO: Get rid of this in https://github.com/pulumi/pulumi-azure-native/issues/331
			continue
		}
		// Workaround for https://github.com/pulumi/pulumi/issues/9883 - remove once merged
		if name == "conditionSets" && property.Items.Schema.Ref.String() == "#/definitions/GovernanceRuleConditionSets" {
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
		//TODO: Remove when https://github.com/Azure/azure-rest-api-specs/pull/14550 is rolled back
		workaroundDelegatedNetworkBreakingChange := property.Ref.String() == "#/definitions/OrchestratorResourceProperties" ||
			property.Ref.String() == "#/definitions/DelegatedSubnetProperties" ||
			property.Ref.String() == "#/definitions/DelegatedControllerProperties"
		if flatten, ok := property.Extensions.GetBool(extensionClientFlatten); (ok && flatten && !isDict) || workaroundDelegatedNetworkBreakingChange {
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

		// Skip read-only properties for input types and write-only properties for output types.
		if resolvedProperty.ReadOnly && !isOutput {
			continue
		}
		if isOutput && isWriteOnly(resolvedProperty) {
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
				OneOf:                m.getOneOfValues(propertySpec),
				Ref:                  propertySpec.Ref,
				Items:                m.itemTypeToProperty(propertySpec.Items),
				AdditionalProperties: m.itemTypeToProperty(propertySpec.AdditionalProperties),
			}
		} else {
			if m.isEnum(&propertySpec.TypeSpec) {
				apiProperty = resources.AzureAPIProperty{Type: "string"}
			} else {
				apiProperty = resources.AzureAPIProperty{
					Type:                 propertySpec.Type,
					OneOf:                m.getOneOfValues(propertySpec),
					Ref:                  propertySpec.Ref,
					Minimum:              resolvedProperty.Minimum,
					Maximum:              resolvedProperty.Maximum,
					MinLength:            resolvedProperty.MinLength,
					MaxLength:            resolvedProperty.MaxLength,
					Pattern:              resolvedProperty.Pattern,
					Items:                m.itemTypeToProperty(propertySpec.Items),
					AdditionalProperties: m.itemTypeToProperty(propertySpec.AdditionalProperties),
				}
			}

			// Apply manual metadata about Force New properties.
			apiProperty.ForceNew = m.forceNew(resolvedProperty, name, isType)
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
		discriminator, discriminatorDesc, isDU, err := m.getDiscriminator(allOfSchema)
		if err != nil {
			return nil, err
		}
		if isDU {
			prop := allOfProperties.properties[discriminator]
			sdkDiscriminator := discriminator
			if prop.SdkName != "" {
				sdkDiscriminator = prop.SdkName
			}

			propSpec := allOfProperties.specs[sdkDiscriminator]
			discriminatorValue := getDiscriminatorValue(resolvedSchema)
			propSpec.Const = discriminatorValue
			propSpec.Type = "string"
			propSpec.Ref = ""
			propSpec.OneOf = nil

			// Add the discriminator value to the property description to help users fill it.
			propSpec.Description = fmt.Sprintf("%s\nExpected value is '%s'.", discriminatorDesc, discriminatorValue)

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

// getDiscriminator returns a property name and description for a discriminator if it's defined on the schema.
// The boolean return flag is true when a discriminator is found.
func (m *moduleGenerator) getDiscriminator(resolvedSchema *openapi.Schema) (string, string, bool, error) {
	if resolvedSchema.Discriminator != "" {
		property := resolvedSchema.Properties[resolvedSchema.Discriminator]
		resolvedProperty, err := resolvedSchema.ResolveSchema(&property)
		if err != nil {
			return "", "", false, err
		}
		return resolvedSchema.Discriminator, resolvedProperty.Description, true, nil
	}
	for _, s := range resolvedSchema.AllOf {
		parentSchema, err := resolvedSchema.ResolveSchema(&s)
		if err != nil {
			return "", "", false, err
		}
		parentDiscriminator, parentDescription, has, err := m.getDiscriminator(parentSchema)
		if err != nil {
			return "", "", false, err
		}
		if has {
			return parentDiscriminator, parentDescription, true, nil
		}
	}
	return "", "", false, nil
}

// getDiscriminatorValue return the value of the discriminator value extension, or the schema name as the default.
func getDiscriminatorValue(resolvedSchema *openapi.Schema) string {
	if v, ok := resolvedSchema.Extensions.GetString(extensionDiscriminatorValue); ok {
		return v
	}
	return resolvedSchema.ReferenceContext.ReferenceName
}

// isWriteOnly return true for properties which are annotated with mutability extension that contain no 'read' value.
func isWriteOnly(schema *openapi.Schema) bool {
	mutability, has := schema.Extensions.GetStringSlice(extensionMutability)
	if !has {
		return false
	}
	for _, v := range mutability {
		if v == extensionMutabilityRead {
			return false
		}
	}
	return true
}

// forceNew return true if a property with a given name requires a replacement in the resource
// that is currently being generated, based on forceNewMap.
func (m *moduleGenerator) forceNew(schema *openapi.Schema, propertyName string, isType bool) bool {
	// Mutability extension signals whether a property can be updated in-place. Lack of the extension means
	// updatable by default.
	// Note: a non-updatable property at a subtype level (a property of a property of a resource) does not
	// mandate the replacement of the whole resource. Anyway, it's used very seldom (2 places at the time of writing).
	// Example: `StorageAccount.encryption.services.blob.keyType` is non-updatable, but a user can remove `blob`
	// and then re-add it with the new `keyType` without replacing the whole storage account (which would be
	// very disruptive).
	if mutability, ok := schema.Extensions.GetStringSlice(extensionMutability); ok && !isType {
		for _, v := range mutability {
			if v == extensionMutabilityUpdate {
				return false
			}
		}
		return true
	}

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

// itemTypeToProperty converts a type of an element in an array or a dictionary to a corresponding
// API property definition of that element type. It only converts the relevant subset of properties,
// and does so recursively.
func (m *moduleGenerator) itemTypeToProperty(typ *schema.TypeSpec) *resources.AzureAPIProperty {
	if typ == nil || typ.Ref == resources.TypeAny {
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
		Type:                 typ.Type,
		Ref:                  typ.Ref,
		OneOf:                oneOf,
		Items:                m.itemTypeToProperty(typ.Items),
		AdditionalProperties: m.itemTypeToProperty(typ.AdditionalProperties),
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

	// TODO: Remove this switch if https://github.com/Azure/azure-rest-api-specs/issues/13167 is fixed
	var defaultValue interface{}
	switch typeSpec.Type {
	case "object":
		if schema.Default != nil {
			fmt.Printf("Default value '%v' can't be specified for an object property %q\n", schema.Default, name)
		}
	default:
		defaultValue = schema.Default
	}

	propertySpec := pschema.PropertySpec{
		Description: description,
		Default:     defaultValue,
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
		if primitiveTypeName == "integer" && resolvedSchema.Format == "int64" {
			// Pulumi schema's integer is limited to 32 bits. Model a 64-bit integer as a number.
			primitiveTypeName = "number"
		}
	}

	// If this is an enum and it's an input type, generate the enum type definition.
	enumExtension, isEnum := resolvedSchema.Extensions[extensionEnum].(map[string]interface{})
	if !isOutput && isEnum && resolvedSchema.Enum != nil &&
		// There are some boolean-, integer-, and number- based definitions but they aren't allowed by the Azure specs.
		// Ignore both to preserve over-the-wire format even if they say it's an enum to be modelled as a string.
		primitiveTypeName != "boolean" && primitiveTypeName != "integer" && primitiveTypeName != "number" {
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
			tok = m.typeName(context, isOutput) + m.inlineTypeName(context, propertyName)
		}

		// If an object type is referenced, add its definition to the type map.
		discriminatedType, ok, err := m.genDiscriminatedType(resolvedSchema, isOutput)
		if err != nil {
			return nil, err
		}
		if ok {
			return discriminatedType, nil
		}

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

			spec := pschema.ComplexTypeSpec{
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Description: resolvedSchema.Description,
					Type:        "object",
					Properties:  props.specs,
					Required:    props.requiredSpecs.SortedValues(),
				},
			}

			if v, has := m.pkg.Types[tok]; has {
				// TODO: @stack72 handle this as part of https://github.com/pulumi/pulumi-azure-native/issues/1606
				if strings.HasPrefix(tok, "azure-native:eventgrid:") ||
					tok == "azure-native:authorization:PrincipalResponse" && len(v.Properties) == 2 ||
					tok == "azure-native:netapp:ExportPolicyRuleResponse" && len(v.Properties) == 14 ||
					tok == "azure-native:netapp:ReplicationObject" && len(v.Properties) == 5 ||
					tok == "azure-native:netapp:ExportPolicyRule" && len(v.Properties) == 14 ||
					tok == "azure-native:securityinsights/v20211001:IncidentOwnerInfoResponse" && len(v.Properties) == 5 ||
					tok == "azure-native:securityinsights/v20211001:IncidentOwnerInfo" && len(v.Properties) == 5 ||
					tok == "azure-native:automanage:ConfigurationProfileAssignmentProperties" && len(v.Properties) == 4 ||
					tok == "azure-native:automanage:ConfigurationProfileAssignmentPropertiesResponse" && len(v.Properties) == 6 {
					// TODO: this was needed to unblock nightly generation: generalize this case.
					v = spec
				}
				err := compatibleTypes(spec, v, isOutput)
				if err != nil {
					return nil, errors.Wrapf(err, "incompatible type %q for resource %q (%q)", tok, m.resourceName,
						m.resourceToken)
				}
			}

			m.pkg.Types[tok] = spec

			m.metadata.Types[tok] = resources.AzureAPIType{
				Properties:         props.properties,
				RequiredProperties: props.requiredProperties.SortedValues(),
			}
		}
		return &pschema.TypeSpec{
			Type: "object",
			Ref:  fmt.Sprintf("#/types/%s", tok),
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
				Ref: resources.TypeAny,
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
			Ref: resources.TypeAny,
		}, nil

	default:
		// Primitive type ('string', 'integer', etc.)
		return &pschema.TypeSpec{Type: primitiveTypeName}, nil
	}
}

// inlineTypeName returns a type name suffix to be used as a type name for properties defined inline in the spec.
// It defaults to the TitleCased property name but it also keeps a map of potential collisions. If a collision occurs,
// (i.e. an inline property `foo` has inline properties and one of them is also `foo`, that can also be down several
// levels), then the property name is duplicated (to get `FooFoo` in this example).
func (m *moduleGenerator) inlineTypeName(ctx *openapi.ReferenceContext, propertyName string) string {
	result := strings.Title(propertyName)
	if ex, ok := m.inlineTypes[ctx]; ok {
		for {
			if !ex.Has(result) {
				break
			}
			result += strings.Title(propertyName)
		}
	} else {
		m.inlineTypes[ctx] = codegen.NewStringSet()
	}
	m.inlineTypes[ctx].Add(result)
	return result
}

// compatibleTypes checks that two type specs are allowed to be represented as a single schema type.
// If you face an error coming from this function for a new API change, consider changing the typeNameOverride map
// or adjusting the way the top-level resources are projected in versioner.go.
func compatibleTypes(t1 schema.ComplexTypeSpec, t2 schema.ComplexTypeSpec, isOutput bool) error {
	if t1.Type != t2.Type {
		return errors.Errorf("types do not match: %s vs %s", t1.Type, t2.Type)
	}
	if len(t1.Properties) != len(t2.Properties) {
		return errors.Errorf("property count do not match: %d vs %d", len(t1.Properties), len(t2.Properties))
	}
	if !isOutput && len(t1.Required) != len(t2.Required) {
		return errors.Errorf("required property count do not match: %d vs %d", len(t1.Required), len(t2.Required))
	}

	// Check that every property of T2 exists in T1 and has the same type.
	t1props := map[string]pschema.PropertySpec{}
	for name, p := range t1.Properties {
		t1props[name] = p
	}
	for name, p2 := range t2.Properties {
		if p1, ok := t1props[name]; ok {
			if p1.Type != p2.Type {
				return errors.Errorf("property %q types do not match: %s vs %s", name, p1.Type, p2.Type)
			}
			if p1.Ref != p2.Ref {
				return errors.Errorf("property %q refs do not match: %s vs %s", name, p1.Ref, p2.Ref)
			}
		} else {
			return errors.Errorf("property %q exists in one but not the other", name)
		}
	}

	// Check that required properties are the same. Output types aren't that important in this regards.
	if !isOutput {
		t1reqs := codegen.NewStringSet(t1.Required...)
		for _, name := range t2.Required {
			if !t1reqs.Has(name) {
				return errors.Errorf("property %q is required in one but not the other", name)
			}
		}
	}

	return nil
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
		enumName = m.typeNameOverride(ToUpperCamel(name))
	} else {
		return nil, fmt.Errorf("name key missing from enum metadata")
	}

	tok := fmt.Sprintf("%s:%s:%s", m.pkg.Name, m.module, enumName)

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: description,
			Type:        "string", // This provider only has string enums
		},
	}
	if values, ok := enumExtension["values"].([]interface{}); ok {
		for _, val := range values {
			if val, ok := val.(map[string]interface{}); ok {
				enumVal := pschema.EnumValueSpec{
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
			enumVal := pschema.EnumValueSpec{Value: fmt.Sprintf("%v", val)}
			// Override the name for the values for this Enum since it contains unfortunately
			// named values like `Input` and `Output` which collide especially for Go Codegen.
			if strings.HasPrefix(m.module, "datafactory") && enumName == "ScriptActivityParameterDirection" {
				enumVal.Name = fmt.Sprintf("Value%s", val)
			}
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
	mapping := map[string]string{}
	subtypes, err := resolvedSchema.FindSubtypes()
	if err != nil {
		return nil, false, err
	}
	for _, subtype := range subtypes {
		typ, err := m.genTypeSpec("", subtype, resolvedSchema.ReferenceContext, isOutput)
		if err != nil {
			return nil, false, err
		}
		oneOf = append(oneOf, *typ)
		subtypeSchema, err := resolvedSchema.ResolveSchema(subtype)
		if err != nil {
			return nil, false, err
		}
		discriminatorValue := getDiscriminatorValue(subtypeSchema)
		mapping[discriminatorValue] = typ.Ref
	}

	switch len(oneOf) {
	case 0:
		// Type specifies a discriminator but doesn't have actual subtypes. Ignore the discriminator in this case.
		return nil, false, nil
	case 1:
		// There is just one subtype specified: use it as a definite type.
		return &oneOf[0], true, nil
	default:
		sdkDiscriminator := discriminator
		if clientName, ok := prop.Extensions.GetString(extensionClientName); ok {
			sdkDiscriminator = clientName
		}
		sdkDiscriminator = ToLowerCamel(sdkDiscriminator)

		// Union type for two or more types.
		return &pschema.TypeSpec{
			OneOf: oneOf,
			Discriminator: &pschema.DiscriminatorSpec{
				PropertyName: sdkDiscriminator,
				Mapping:      mapping,
			},
		}, true, nil
	}
}

// typeNameOverrides is a manually-maintained map of alternative names when a type name represents two or more
// distinct types in the same resource provider. This can happen if there are multiple Open API spec files
// for the same RP and version, and each of those files has its own definition of the type under the same name.
// That happens a lot (there are many files for several RPs) but most of the time the definitions are similar
// enough to treat them as same. The following map lists all exceptions from this rule.
// If a new mismatch is introduced in a newly published spec, our codegen will catch the difference
// (see compatibleTypes) and fail. We will have to extend the list below with a new exception to unblock codegen.
var typeNameOverrides = map[string]string{
	// SKU for Redis Enterprise is different from SKU for Redis. Keep them as separate types.
	"Cache.RedisEnterprise.Sku": "EnterpriseSku",
	// Devices RP comes from "deviceprovisioningservices" and "iothub" which are similar but slightly different.
	// In particular, the IP Filter Rule has more properties in the DPS version.
	"Devices.IotDpsResource.IpFilterRule": "TargetIpFilterRule",
	// Workbook vs. MyWorkbook types are slightly different. Probably, a bug in the spec, but we have to disambiguate.
	"Insights.MyWorkbook.ManagedIdentity":        "MyManagedIdentity",
	"Insights.MyWorkbook.UserAssignedIdentities": "MyUserAssignedIdentities",
	// Experiment's endpoint is a much narrower type compared to endpoints in other network resources.
	"Network.Experiment.Endpoint": "ExperimentEndpoint",
	// These are all FrontDoor types. FrontDoor shares a bunch of type names with generate Network provider,
	// but defines them in its own way.
	"Network.Policy.ManagedRuleGroupOverride": "FrontDoorManagedRuleGroupOverride",
	"Network.Policy.ManagedRuleOverride":      "FrontDoorManagedRuleOverride",
	"Network.Policy.ManagedRuleSet":           "FrontDoorManagedRuleSet",
	"Network.Policy.MatchCondition":           "FrontDoorMatchCondition",
	"Network.Policy.MatchVariable":            "FrontDoorMatchVariable",
	"Network.Policy.PolicySettings":           "FrontDoorPolicySettings",
	// IpConfigurationResponse conflicts with IPConfigurationResponse used for existing networking resources.
	// This avoids Dotnet SDK failing to build since codegen currently elides the IPConfigurationResponse
	// output types since it assumes the existing one is sufficient.
	"Network.InboundEndpoint.IpConfiguration": "InboundEndpointIPConfiguration",
	// The following two types are read-only, while the same types in another spec are writable.
	"RecoveryServices.Vault.PrivateEndpointConnection":         "VaultPrivateEndpointConnection",
	"RecoveryServices.Vault.PrivateLinkServiceConnectionState": "VaultPrivateLinkServiceConnectionState",
	// Watchlist resources only appear in a preview spec and not in stable specs. Anyway, the shapes of their
	// types are slightly different from later specs, so we have to disambiguate for top-level resources.
	"SecurityInsights.Watchlist.UserInfo":     "WatchlistUserInfo",
	"SecurityInsights.WatchlistItem.UserInfo": "WatchlistUserInfo",
}

func (m *moduleGenerator) typeNameOverride(typeName string) string {
	key := fmt.Sprintf("%s.%s.%s", m.prov, m.resourceName, typeName)
	if v, ok := typeNameOverrides[key]; ok {
		return v
	}
	return typeName
}

func (m *moduleGenerator) typeName(ctx *openapi.ReferenceContext, isOutput bool) string {
	suffix := ""
	if isOutput {
		suffix = "Response"
	}
	referenceName := m.typeNameOverride(ToUpperCamel(MakeLegalIdentifier(ctx.ReferenceName)))
	return fmt.Sprintf("azure-native:%s:%s%s", m.module, referenceName, suffix)
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

func rawMessage(v interface{}) pschema.RawMessage {
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
