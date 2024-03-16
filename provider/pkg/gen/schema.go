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

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-azure-native-sdk/v2"

type ResourceDeprecation struct {
	ReplacementToken string
}

type Versioning interface {
	ShouldInclude(provider string, version string, typeName, token string) bool
	GetDeprecation(token string) (ResourceDeprecation, bool)
	GetAllVersions(openapi.ProviderName, openapi.ResourceName) []openapi.ApiVersion
}

type GenerationResult struct {
	Schema                     *pschema.PackageSpec
	Metadata                   *resources.AzureAPIMetadata
	Examples                   map[string][]resources.AzureAPIExample
	ForceNewTypes              []ForceNewType
	TypeCaseConflicts          CaseConflicts
	FlattenedPropertyConflicts map[string]map[string]struct{}
	Endpoints                  openapi.Endpoints
}

// PulumiSchema will generate a Pulumi schema for the given Azure providers and resources map.
func PulumiSchema(rootDir string, providerMap openapi.AzureProviders, versioning Versioning) (*GenerationResult, error) {
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
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					Description: "Any additional Tenant IDs which should be used for authentication.",
				},
				"environment": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.",
				},
				"location": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The location to use. ResourceGroups will consult this property for a default location, if one was not supplied explicitly when defining the resource.",
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
					Description: "Allow Managed Service Identity be used for Authentication.",
					Default:     false,
				},
				"msiEndpoint": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.",
				},
				"metadataHost": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Hostname of the Azure Metadata Service.",
				},

				"useOidc": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "Allow OpenID Connect (OIDC) to be used for Authentication.",
					Default:     false,
				},
				"oidcToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The OIDC token to exchange for an Azure token.",
				},
				"oidcRequestToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "Your cloud service or provider’s bearer token to exchange for an OIDC ID token.",
				},
				"oidcRequestUrl": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The URL to initiate the OIDC token exchange. ",
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
					Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.",
				},
				"location": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The location to use. ResourceGroups will consult this property for a default location, if one was not supplied explicitly when defining the resource.",
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
					Description: "Allow Managed Service Identity to be used for Authentication.",
					Default:     false,
				},
				"msiEndpoint": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.",
				},
				"metadataHost": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The Hostname of the Azure Metadata Service.",
				},

				"useOidc": {
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					Description: "Allow OpenID Connect (OIDC) to be used for Authentication.",
					Default:     false,
				},
				"oidcToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The OIDC token to exchange for an Azure token.",
				},
				"oidcRequestToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "Your cloud service or provider’s bearer token to exchange for an OIDC ID token.",
				},
				"oidcRequestUrl": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The URL to initiate the `oidcRequestToken` OIDC token exchange.",
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
	javaPackages := map[string]string{
		"azure-native": "azurenative",
	}
	pythonModuleNames := map[string]string{}
	golangImportAliases := map[string]string{}

	var providers []string
	for prov := range providerMap {
		providers = append(providers, prov)
	}
	sort.Strings(providers)

	// Track some global data
	warnings := []string{}
	var forceNewTypes []ForceNewType
	caseSensitiveTypes := newCaseSensitiveTokens()
	flattenedPropertyConflicts := map[string]map[string]struct{}{}
	endpoints := openapi.Endpoints{}

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
				pkg:                        &pkg,
				metadata:                   &metadata,
				provider:                   providerName,
				apiVersion:                 version,
				allApiVersions:             versions,
				examples:                   exampleMap,
				versioning:                 versioning,
				caseSensitiveTypes:         caseSensitiveTypes,
				rootDir:                    rootDir,
				flattenedPropertyConflicts: flattenedPropertyConflicts,
				allEndpoints:               endpoints,
			}

			// Populate C#, Java, Python and Go module mapping.
			module := gen.moduleName()
			csharpNamespaces[strings.ToLower(providerName)] = providerName
			javaPackages[module] = strings.ToLower(providerName)
			if version != "" {
				csVersion := strings.Title(csharpVersionReplacer.Replace(version))
				csharpNamespaces[module] = fmt.Sprintf("%s.%s", providerName, csVersion)
				javaPackages[module] = fmt.Sprintf("%s.%s", strings.ToLower(providerName), version)
			}
			pythonModuleNames[module] = module
			golangImportAliases[filepath.Join(goBasePath, module, "v2")] = strings.ToLower(providerName)

			// Populate resources and get invokes.
			items := versionMap[version]
			var resources []string
			for resource := range items.Resources {
				resources = append(resources, resource)
			}
			sort.Strings(resources)

			for _, typeName := range resources {
				resource := items.Resources[typeName]
				nestedResourceBodyRefs := findNestedResourceBodyRefs(resource, items.Resources)
				err := gen.genResources(typeName, resource, nestedResourceBodyRefs)
				if err != nil {
					return nil, err
				}
			}

			// Populate invokes.
			gen.genInvokes(items.Invokes)
			warnings = append(warnings, gen.warnings...)
			forceNewTypes = append(forceNewTypes, gen.forceNewTypes...)
		}
	}

	err := genMixins(&pkg, &metadata)
	if err != nil {
		return nil, err
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":                 goBasePath,
		"packageImportAliases":           golangImportAliases,
		"rootPackageName":                "pulumiazurenativesdk",
		"generateResourceContainerTypes": false,
		"disableInputTypeRegistrations":  true,
		"internalModuleName":             "utilities",
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
			"pulumi": ">=3.35.0,<4.0.0",
		},
		"usesIOClasses": true,
		"readme": `The native Azure provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
		"pyproject": map[string]bool{
			"enabled": true,
		},
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "3.*",
			"System.Collections.Immutable": "5.0.0",
		},
		"namespaces": csharpNamespaces,
	})

	pkg.Language["java"] = rawMessage(map[string]interface{}{
		"packages": javaPackages,
	})

	return &GenerationResult{
		Schema:                     &pkg,
		Metadata:                   &metadata,
		Examples:                   exampleMap,
		ForceNewTypes:              forceNewTypes,
		TypeCaseConflicts:          caseSensitiveTypes.findCaseConflicts(),
		FlattenedPropertyConflicts: flattenedPropertyConflicts,
		Endpoints:                  endpoints,
	}, nil
}

func (g *packageGenerator) genInvokes(invokes map[string]*openapi.ResourceSpec) {
	var invokeNames []string
	for invokeName := range invokes {
		invokeNames = append(invokeNames, invokeName)
	}
	sort.Strings(invokeNames)

	for _, invokeName := range invokeNames {
		invoke := invokes[invokeName]

		op := invoke.PathItem.Post
		if op == nil {
			op = invoke.PathItem.Get
		}
		if op == nil {
			panic(fmt.Sprintf("Invoke %s has no POST or GET operation", invokeName))
		}

		g.genFunctions(invokeName, invoke.Path, invoke.PathItem.Parameters, op, invoke.Swagger)
	}
}

func findNestedResourceBodyRefs(resource *openapi.ResourceSpec, resourceSpecs map[string]*openapi.ResourceSpec) []string {
	nestedResources := findNestedResources(resource, resourceSpecs)
	return findResourcesBodyRefs(nestedResources)
}

func findResourcesBodyRefs(nestedResources []*openapi.ResourceSpec) []string {
	var nestedResourcePostBodyTypeRefs []string
	for _, nestedResource := range nestedResources {
		nestedResourcePostBodyTypeRef, ok := findResourcePutBodyTypeRef(nestedResource)
		if ok {
			nestedResourcePostBodyTypeRefs = append(nestedResourcePostBodyTypeRefs, nestedResourcePostBodyTypeRef)
		}
	}
	return nestedResourcePostBodyTypeRefs
}

func findResourcePutBodyTypeRef(resource *openapi.ResourceSpec) (string, bool) {
	if resource.PathItem != nil && resource.PathItem.Put != nil {
		for _, parameter := range resource.PathItem.Put.Parameters {
			if parameter.In != "body" || parameter.Schema == nil {
				continue
			}
			return parameter.Schema.Ref.String(), true
		}
	}
	return "", false
}

func findNestedResources(resource *openapi.ResourceSpec, resourceSpecs map[string]*openapi.ResourceSpec) []*openapi.ResourceSpec {
	orderedNames := codegen.SortedKeys(resourceSpecs)
	var nestedResourceSpecs []*openapi.ResourceSpec
	for _, otherResourceName := range orderedNames {
		otherResource := resourceSpecs[otherResourceName]
		// Ignore self
		if otherResource.Path == resource.Path {
			continue
		}
		if strings.HasPrefix(otherResource.Path, resource.Path) {
			nestedResourceSpecs = append(nestedResourceSpecs, otherResource)
		}
	}
	return nestedResourceSpecs
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
	for tok, r := range customresources.SchemaMixins() {
		if _, has := pkg.Resources[tok]; has {
			return errors.Errorf("Resource %q is already defined", tok)
		}
		pkg.Resources[tok] = r
	}
	for tok, t := range customresources.SchemaTypeMixins() {
		if _, has := pkg.Types[tok]; has {
			return errors.Errorf("Type %q is already defined", tok)
		}
		pkg.Types[tok] = t
	}
	for tok, t := range customresources.SchemaTypeOverrides() {
		pkg.Types[tok] = t
	}
	for tok, r := range customresources.MetaMixins() {
		if _, has := metadata.Resources[tok]; has {
			return errors.Errorf("Metadata %q is already defined", tok)
		}
		metadata.Resources[tok] = r
	}
	for tok, r := range customresources.MetaTypeMixins() {
		if _, has := metadata.Types[tok]; has {
			return errors.Errorf("Metadata type %q is already defined", tok)
		}
		metadata.Types[tok] = r
	}
	for tok, r := range customresources.MetaTypeOverrides() {
		metadata.Types[tok] = r
	}

	// Add a note regarding WorkspaceSqlAadAdmin creation.
	workspaceSqlAadAdmin := pkg.Resources["azure-native:synapse:WorkspaceSqlAadAdmin"]
	workspaceSqlAadAdmin.Description += "\n\nNote: SQL AAD Admin is configured automatically during workspace creation and assigned to the current user. One can't add more admins with this resource unless you manually delete the current SQL AAD Admin."
	pkg.Resources["azure-native:synapse:WorkspaceSqlAadAdmin"] = workspaceSqlAadAdmin

	// Remove unused types.
	normalizePackage(pkg, metadata)

	return nil
}

func normalizePackage(pkg *pschema.PackageSpec, metadata *resources.AzureAPIMetadata) {
	// Record all type tokens referenced from resources and functions.
	usedTypes := map[string]bool{}
	visitor := func(t string, _ pschema.ComplexTypeSpec) {
		usedTypes[t] = true
	}
	VisitPackageSpecTypes(pkg, visitor)

	// Elide unused types.
	allTypeNames := codegen.SortedKeys(pkg.Types)
	for _, typeName := range allTypeNames {
		if !usedTypes[typeName] {
			t := pkg.Types[typeName]
			if len(t.Enum) > 0 {
				continue
			}
			delete(pkg.Types, typeName)
			delete(metadata.Types, typeName)
		}
	}
}

// Microsoft-specific API extension constants.
const (
	extensionClientFlatten      = "x-ms-client-flatten"
	extensionClientName         = "x-ms-client-name"
	extensionDiscriminatorValue = "x-ms-discriminator-value"
	extensionEnum               = "x-ms-enum"
	extensionIdentifiers        = "x-ms-identifiers" // ids in keyed arrays
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
	pkg                *pschema.PackageSpec
	metadata           *resources.AzureAPIMetadata
	provider           openapi.ProviderName
	examples           map[string][]resources.AzureAPIExample
	apiVersion         string
	allApiVersions     []openapi.ApiVersion
	versioning         Versioning
	caseSensitiveTypes caseSensitiveTokens
	warnings           []string
	// rootDir is used to resolve relative paths in the examples.
	rootDir                    string
	forceNewTypes              []ForceNewType
	flattenedPropertyConflicts map[string]map[string]struct{}
	allEndpoints               openapi.Endpoints
}

func (g *packageGenerator) genResources(typeName string, resource *openapi.ResourceSpec, nestedResourceBodyRefs []string) error {
	// Resource names should consistently start with upper case.
	// We need to alias the previous, lowercase name so users can upgrade to v2 without replacement.
	// These aliases will not be required anymore with v3; their removal is tracked by #2411.
	typeNameAliases := []string{}
	titleCasedTypeName := strings.Title(typeName)
	if titleCasedTypeName != typeName {
		typeNameAliases = append(typeNameAliases, typeName)
	}

	// A single API path can be modelled as several resources if it accepts a polymorphic payload:
	// i.e., when the request body is a discriminated union type of several object types. Pulumi
	// schema doesn't support polymorphic (OneOf) resources, so the provider creates a separate resource
	// per each union case. We call them "variants" in the code below.
	variants, err := g.findResourceVariants(resource)
	if err != nil {
		return errors.Wrapf(err, "resource %s.%s", g.provider, titleCasedTypeName)
	}

	for _, v := range variants {
		err = g.genResourceVariant(resource, v, nestedResourceBodyRefs, typeNameAliases...)
		if err != nil {
			return err
		}
	}

	// If variants are found, we don't need the main or base resource.
	if len(variants) > 0 {
		return nil
	}

	mainResource := &resourceVariant{
		ResourceSpec: resource,
		typeName:     titleCasedTypeName,
	}
	if resource.DeprecationMessage != "" {
		mainResource.deprecationMessage = resource.DeprecationMessage
	}
	return g.genResourceVariant(resource, mainResource, nestedResourceBodyRefs, typeNameAliases...)
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

func (g *packageGenerator) makeTypeAlias(alias, apiVersion string) pschema.AliasSpec {
	fqAlias := fmt.Sprintf("%s:%s:%s", g.pkg.Name, g.providerApiToModule(apiVersion), alias)
	return pschema.AliasSpec{Type: &fqAlias}
}

func (g *packageGenerator) genResourceVariant(apiSpec *openapi.ResourceSpec, resource *resourceVariant, nestedResourceBodyRefs []string, typeNameAliases ...string) error {
	module := g.moduleName()
	swagger := resource.Swagger
	path := resource.PathItem

	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, resource.typeName)
	if !g.versioning.ShouldInclude(g.provider, g.apiVersion, resource.typeName, resourceTok) {
		return nil
	}

	// Generate the resource.
	gen := moduleGenerator{
		pkg:                        g.pkg,
		metadata:                   g.metadata,
		module:                     module,
		prov:                       g.provider,
		resourceName:               resource.typeName,
		resourceToken:              resourceTok,
		visitedTypes:               make(map[string]bool),
		caseSensitiveTypes:         g.caseSensitiveTypes,
		inlineTypes:                map[*openapi.ReferenceContext]codegen.StringSet{},
		nestedResourceBodyRefs:     nestedResourceBodyRefs,
		flattenedPropertyConflicts: map[string]struct{}{},
		allEndpoints:               g.allEndpoints,
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

	if resourceDeprecation, ok := g.versioning.GetDeprecation(resourceTok); ok {
		deprecationMessage := fmt.Sprintf(
			"%s is being removed in the next major version of this provider. "+
				"Upgrade to at least %s to guarantee forwards compatibility.", resourceTok, resourceDeprecation.ReplacementToken,
		)
		if resource.deprecationMessage == "" {
			resource.deprecationMessage = deprecationMessage
		} else {
			resource.deprecationMessage = resource.deprecationMessage + "\n" + deprecationMessage
		}
	}

	additionalDocs := getAdditionalDocs(g.rootDir, resourceTok)

	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: g.formatDescription(resourceResponse.description, resource.typeName, swagger.Info.Version, apiSpec.PreviousVersion, additionalDocs),
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties:    resourceRequest.specs,
		RequiredInputs:     resourceRequest.requiredSpecs.SortedValues(),
		Aliases:            g.generateAliases(resource, typeNameAliases...),
		DeprecationMessage: resource.deprecationMessage,
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, resource.typeName)
	if g.versioning.ShouldInclude(g.provider, g.apiVersion, resource.typeName, functionTok) {
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
				Description:        g.formatFunctionDescription(readOp, resource.typeName, resourceResponse, swagger.Info),
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

	requiredContainers := mergeRequiredContainers(resourceRequest.requiredContainers, additionalRequiredContainers(resourceTok))

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
		RequiredContainers:   requiredContainers,
		DefaultProperties:    propertyDefaults(module, resource.typeName),
	}

	g.metadata.Resources[resourceTok] = r

	g.generateExampleReferences(resourceTok, path, swagger)
	g.forceNewTypes = append(g.forceNewTypes, gen.forceNewTypes...)
	if len(gen.flattenedPropertyConflicts) > 0 {
		g.flattenedPropertyConflicts[resourceTok] = gen.flattenedPropertyConflicts
	}
	return nil
}

func (g *packageGenerator) generateAliases(resource *resourceVariant, typeNameAliases ...string) []pschema.AliasSpec {
	var aliases []pschema.AliasSpec

	for _, alias := range typeNameAliases {
		aliases = append(aliases, g.makeTypeAlias(alias, g.apiVersion))
	}

	// Add an alias for each API version that has the same path in it.
	for _, version := range resource.CompatibleVersions {
		aliases = append(aliases, g.makeTypeAlias(resource.typeName, version))

		for _, alias := range typeNameAliases {
			aliases = append(aliases, g.makeTypeAlias(alias, version))
		}
	}

	return aliases
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
	sortedExampleKeys := codegen.SortedKeys(examples)

	result := make([]resources.AzureAPIExample, 0, len(examples))
	for _, k := range sortedExampleKeys {
		v := examples[k]
		resolved := v.(map[string]interface{})
		if _, ok := resolved["$ref"]; !ok {
			continue
		}

		relativePath := resolved["$ref"].(string)
		relativeURL, err := url.Parse(relativePath)
		if err != nil {
			return err
		}

		url, err := swagger.ResolveReference(relativeURL.String())
		if err != nil {
			return err
		}

		url, err = filepath.Rel(g.rootDir, url)
		if err != nil {
			return err
		}
		if _, err := os.Stat(filepath.Join(g.rootDir, url)); err != nil {
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

// genFunctions defines functions for list* (listKeys, listSecrets, etc.)
// and get* (getFullUrl, getBastionShareableLink, etc.) POST and GET endpoints.
func (g *packageGenerator) genFunctions(typeName, path string, specParams []spec.Parameter, operation *spec.Operation, swagger *openapi.Spec) {
	module := g.moduleName()
	gen := moduleGenerator{
		pkg:                        g.pkg,
		metadata:                   g.metadata,
		module:                     module,
		resourceToken:              fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		prov:                       g.provider,
		resourceName:               typeName,
		visitedTypes:               make(map[string]bool),
		caseSensitiveTypes:         g.caseSensitiveTypes,
		inlineTypes:                map[*openapi.ReferenceContext]codegen.StringSet{},
		flattenedPropertyConflicts: map[string]struct{}{},
	}

	// Generate the function to get this resource.
	functionTok := g.generateTok(typeName, g.apiVersion)
	if !g.shouldInclude(typeName, functionTok, g.apiVersion) {
		return
	}

	parameters := swagger.MergeParameters(operation.Parameters, specParams)
	request, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, nil, nil)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(operation.Responses.StatusCodeResponses, swagger.ReferenceContext, nil)
	if err != nil {
		log.Printf("failed to generate '%s': response type: %s", functionTok, err.Error())
		return
	}

	if response == nil || len(response.specs) == 0 {
		// Response is specified empty, do not generate an invoke for it.
		return
	}

	functionSpec := pschema.FunctionSpec{
		Description: g.formatFunctionDescription(operation, typeName, response, swagger.Info),
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

// moduleName produces the module name from the provider name and the API version (e.g. (`Compute`, `2020-07-01` => `compute/v20200701`).
func (g *packageGenerator) moduleName() string {
	return g.providerApiToModule(g.apiVersion)
}

func (g *packageGenerator) providerApiToModule(apiVersion string) string {
	if apiVersion == "" {
		return strings.ToLower(g.provider)
	}
	return fmt.Sprintf("%s/%s", strings.ToLower(g.provider), apiVersion)
}

func (g *packageGenerator) generateTok(typeName string, apiVersion string) string {
	return fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.providerApiToModule(apiVersion), typeName)
}

func (g *packageGenerator) shouldInclude(typeName, tok, version string) bool {
	return g.versioning.ShouldInclude(g.provider, version, typeName, tok)
}

func (g *packageGenerator) formatFunctionDescription(op *spec.Operation, typeName string, response *propertyBag, info *spec.Info) string {
	desc := response.description
	if op.Description != "" {
		desc = op.Description
	}
	return g.formatDescription(desc, typeName, info.Version, "", nil)
}

func (g *packageGenerator) formatDescription(desc string, typeName string, defaultVersion, previousDefaultVersion string, additionalDocs *string) string {
	var b strings.Builder
	b.WriteString(desc)

	if g.apiVersion == "" {
		fmt.Fprintf(&b, "\nAzure REST API version: %s.", defaultVersion)
		if previousDefaultVersion != "" {
			fmt.Fprintf(&b, " Prior API version in Azure Native 1.x: %s.", previousDefaultVersion)
		}

		// List other available API versions, if any.
		allVersions := g.versioning.GetAllVersions(g.provider, typeName)
		includedVersions := []openapi.ApiVersion{}
		for _, v := range allVersions {
			// Don't list the default version twice.
			if v == defaultVersion {
				continue
			}
			tok := g.generateTok(typeName, openapi.ApiToSdkVersion(v))
			if g.shouldInclude(typeName, tok, v) {
				includedVersions = append(includedVersions, v)
			}
		}

		if len(includedVersions) > 0 {
			fmt.Fprintf(&b, "\n\nOther available API versions: %s.", strings.Join(includedVersions, ", "))
		}
	}

	if additionalDocs != nil {
		fmt.Fprintf(&b, "\n\n%s", *additionalDocs)
	}

	return b.String()
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

type caseSensitiveTokens struct {
	// Map of the lowered tokens to the list of all case variants. The first item in the list is the
	// first seen casing which will be used everywhere.
	tokensLowered map[string][]string
}

func newCaseSensitiveTokens() caseSensitiveTokens {
	return caseSensitiveTokens{tokensLowered: make(map[string][]string)}
}

// normalizeTokenCase returns the normalized token.
// This normalizes all casing to the first seen casing. For this to be consistent,
// we rely on iterating all specs in the same order each time.
func (t *caseSensitiveTokens) normalizeTokenCase(token string) string {
	tokenLowered := strings.ToLower(token)
	caseVariants, alreadySeen := t.tokensLowered[tokenLowered]
	if !alreadySeen {
		t.tokensLowered[tokenLowered] = []string{token}
		return token
	}
	foundExactCasing := false
	for _, v := range caseVariants {
		if v == token {
			foundExactCasing = true
			break
		}
	}
	if !foundExactCasing {
		caseVariants = append(caseVariants, token)
		t.tokensLowered[tokenLowered] = caseVariants
	}
	return caseVariants[0]
}

// Map of the resolved type and a list of all its case variants.
type CaseConflicts map[string][]string

func (t *caseSensitiveTokens) findCaseConflicts() CaseConflicts {
	conflicts := make(map[string][]string)
	for _, caseVariants := range t.tokensLowered {
		if len(caseVariants) > 1 {
			conflicts[caseVariants[0]] = caseVariants[1:]
		}
	}
	return conflicts
}

type ForceNewType struct {
	Module        string
	Provider      string
	ResourceName  string
	ReferenceName string
	Property      string
}

type moduleGenerator struct {
	pkg                        *pschema.PackageSpec
	metadata                   *resources.AzureAPIMetadata
	module                     string
	prov                       string
	resourceName               string
	resourceToken              string
	visitedTypes               map[string]bool
	caseSensitiveTypes         caseSensitiveTokens
	inlineTypes                map[*openapi.ReferenceContext]codegen.StringSet
	nestedResourceBodyRefs     []string
	forceNewTypes              []ForceNewType
	flattenedPropertyConflicts map[string]struct{}
	allEndpoints               openapi.Endpoints
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
				// All path parameters are part of resource ID, so they always cause replacement.
				WillReplaceOnChanges: param.In == "path",
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

// getDiscriminatorValue return the value of the discriminator value extension, or the schema name as the default.
func getDiscriminatorValue(resolvedSchema *openapi.Schema) string {
	if v, ok := resolvedSchema.Extensions.GetString(extensionDiscriminatorValue); ok {
		return v
	}
	return resolvedSchema.ReferenceContext.ReferenceName
}

func getPropertyDescription(schema *spec.Schema, context *openapi.ReferenceContext, maintainSubResourceIfUnset bool) (string, error) {
	description := schema.Description
	if description == "" {
		resolvedSchema, err := context.ResolveSchema(schema)
		if err != nil {
			return "", err
		}

		description = resolvedSchema.Description
	}
	if maintainSubResourceIfUnset {
		if description != "" {
			description = description + "\n"
		}
		description = description + "These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion."
	}
	return description, nil
}

// parameterBag keeps the schema and metadata parameters for a single resource or invocation.
type parameterBag struct {
	description        string
	specs              map[string]pschema.PropertySpec
	parameters         []resources.AzureAPIParameter
	requiredSpecs      codegen.StringSet
	requiredContainers RequiredContainers
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
	bag.requiredContainers = mergeRequiredContainers(bag.requiredContainers, other.requiredContainers)
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

// GoModVersion Creates a valid go mod version from our pulumictl version.
// Essentially, this removes any '+xxx' additions. See tests for examples.
func GoModVersion(packageVersion *semver.Version) string {
	if packageVersion == nil {
		return "latest"
	}
	buildVersion := *packageVersion
	// If the version has a prerelease with a build number like 0+9fa804e8, make if Go-compatible by simplifying to 9fa804e8.
	if buildVersion.Pre != nil && buildVersion.Build != nil {
		buildVersion.Pre = buildVersion.Pre[:1]
		for _, build := range buildVersion.Build {
			buildVersion.Pre = append(buildVersion.Pre, semver.PRVersion{VersionStr: build})
		}
	}
	// Ignore build versions
	buildVersion.Build = nil
	return "v" + buildVersion.String()
}

func GoModulePathVersion(packageVersion string) string {
	// Take up to the first dot
	vString, _, found := strings.Cut(packageVersion, ".")
	if !found {
		return ""
	}
	switch vString {
	case "0", "1":
		return ""
	}
	return "/v" + vString
}
