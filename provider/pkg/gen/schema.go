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
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/segmentio/encoding/json"

	"github.com/blang/semver"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-azure-native-sdk/v2"
const goModuleRepoPath = "github.com/pulumi/pulumi-azure-native-sdk"
const goModuleVersion = "/v2"

// pulumiProviderName is the name of the provider as used in all tokens.
const pulumiProviderName = "azure-native"

type ResourceDeprecation struct {
	ReplacementToken string
}

type Versioning interface {
	ShouldInclude(moduleName openapi.ModuleName, version *openapi.ApiVersion, typeName, token string) bool
	GetDeprecation(token string) (ResourceDeprecation, bool)
	GetAllVersions(openapi.ModuleName, openapi.ResourceName) []openapi.ApiVersion
	GetPreviousCompatibleTokensLookup() (*CompatibleTokensLookup, error)
}

type GenerationResult struct {
	Schema                     *pschema.PackageSpec
	Metadata                   *resources.AzureAPIMetadata
	Examples                   map[string][]resources.AzureAPIExample
	ForceNewTypes              []ForceNewType
	TypeCaseConflicts          CaseConflicts
	FlattenedPropertyConflicts map[string]map[string]any
	// A map of module -> resource -> set of paths, to record resources that have conflicts where the same resource
	// maps to more than one API path.
	PathConflicts map[openapi.ModuleName]map[openapi.ResourceName]map[string][]openapi.ApiVersion
	TokenPaths    map[string]string
}

// PulumiSchema will generate a Pulumi schema for the given Azure modules and resources map.
func PulumiSchema(rootDir string, modules openapi.AzureModules, versioning Versioning, providerVersion semver.Version) (*GenerationResult, error) {
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
					Description: "The Client Secret which should be used. For use when authenticating as a Service Principal using a Client Secret.",
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
				},
				"oidcToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The OIDC token to exchange for an Azure token.",
				},
				"oidcTokenFilePath": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "The path to a file containing an OIDC token to exchange for an Azure token.",
				},
				"oidcRequestToken": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "Your cloud service or provider's bearer token to exchange for an OIDC ID token.",
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

	// Track some global data
	var forceNewTypes []ForceNewType
	caseSensitiveTypes := newCaseSensitiveTokens()
	flattenedPropertyConflicts := map[string]map[string]any{}
	exampleMap := make(map[string][]resources.AzureAPIExample)
	resourcesPathTracker := newResourcesPathConflictsTracker()
	previousCompatibleTokensLookup, err := versioning.GetPreviousCompatibleTokensLookup()
	if err != nil {
		return nil, err
	}

	if !previousCompatibleTokensLookup.IsPopulated() && providerVersion.Major >= 3 {
		return nil, fmt.Errorf("GetPreviousCompatibleTokensLookup is not populated. This is likely due to v%d-token-paths.json being missing or empty", providerVersion.Major-1)
	}

	for moduleName, moduleVersions := range util.MapOrdered(modules) {
		resourcePaths := map[openapi.ResourceName]map[string][]openapi.ApiVersion{}

		allVersions := util.SortedKeys(moduleVersions)
		// The version in the parsed module is "" for the default version.
		for moduleVersion, versionResources := range util.MapOrdered(moduleVersions) {
			var sdkVersion openapi.SdkVersion
			var apiVersion *openapi.ApiVersion
			if moduleVersion.IsDefault() {
				apiVersion = nil
				sdkVersion = ""
			} else {
				apiVersion = &moduleVersion
				sdkVersion = openapi.ApiToSdkVersion(moduleVersion)
			}

			gen := packageGenerator{
				pkg:                            &pkg,
				metadata:                       &metadata,
				moduleName:                     moduleName,
				apiVersion:                     apiVersion,
				sdkVersion:                     sdkVersion,
				allVersions:                    allVersions,
				examples:                       exampleMap,
				versioning:                     versioning,
				caseSensitiveTypes:             caseSensitiveTypes,
				rootDir:                        rootDir,
				flattenedPropertyConflicts:     flattenedPropertyConflicts,
				majorVersion:                   int(providerVersion.Major),
				resourcePaths:                  map[openapi.ResourceName]map[string]openapi.ApiVersion{},
				previousCompatibleTokensLookup: previousCompatibleTokensLookup,
			}

			// Populate C#, Java, Python and Go module mapping.
			module := gen.moduleWithVersion()
			csharpNamespaces[moduleName.Lowered()] = string(moduleName)
			javaPackages[string(module)] = moduleName.Lowered()
			if sdkVersion != "" {
				csVersion := strings.Title(csharpVersionReplacer.Replace(string(sdkVersion)))
				csharpNamespaces[string(module)] = fmt.Sprintf("%s.%s", moduleName, csVersion)
				javaPackages[string(module)] = fmt.Sprintf("%s.%s", moduleName.Lowered(), sdkVersion)
			}
			pythonModuleNames[string(module)] = string(module)
			golangImportAliases[goModuleName(gen.moduleName, gen.sdkVersion)] = moduleName.Lowered()

			// Populate resources and get invokes.
			for typeName, resource := range util.MapOrdered(versionResources.Resources) {
				nestedResourceBodyRefs := findNestedResourceBodyRefs(resource, versionResources.Resources)
				err := gen.genResources(typeName, resource, nestedResourceBodyRefs)
				if err != nil {
					return nil, err
				}
			}

			// Populate invokes.
			gen.genInvokes(versionResources.Invokes)

			forceNewTypes = append(forceNewTypes, gen.forceNewTypes...)
			gen.mergeResourcePathsInto(resourcePaths)
		}

		resourcesPathTracker.addPathConflictsForModule(moduleName, resourcePaths)
	}

	// When a resource maps to more than one API path, it's a conflict and we need to detect and report it. #2495
	isReleaseBuild := len(providerVersion.Build) == 0
	if providerVersion.Major >= 3 && isReleaseBuild && resourcesPathTracker.hasConflicts() {
		return nil, fmt.Errorf("path conflicts detected. You probably need to add a case to schema.go/dedupResourceNameByPath.\n%+v", resourcesPathTracker.pathConflicts)
	}

	err = genMixins(&pkg, &metadata)
	if err != nil {
		return nil, err
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":                 goBasePath,
		"packageImportAliases":           golangImportAliases,
		"importPathPattern":              "github.com/pulumi/pulumi-azure-native-sdk/{module}/v2",
		"rootPackageName":                "pulumiazurenativesdk",
		"generateResourceContainerTypes": false,
		"disableInputTypeRegistrations":  true,
		"internalModuleName":             "utilities",
		"respectSchemaVersion":           true,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"readme": `The native Azure provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
		"respectSchemaVersion": true,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"usesIOClasses":       true,
		"inputTypes":          "classes-and-dicts",
		"readme": `The native Azure provider package offers support for all Azure Resource Manager (ARM)
resources and their properties. Resources are exposed as types from modules based on Azure Resource
Providers such as 'compute', 'network', 'storage', and 'web', among many others. Using this package
allows you to programmatically declare instances of any Azure resource and any supported resource
version using infrastructure as code, which Pulumi then uses to drive the ARM API.`,
		"pyproject": map[string]bool{
			"enabled": true,
		},
		"respectSchemaVersion": true,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "3.*",
			"System.Collections.Immutable": "5.0.0",
		},
		"namespaces":           csharpNamespaces,
		"respectSchemaVersion": true,
	})

	pkg.Language["java"] = rawMessage(map[string]interface{}{
		"packages": javaPackages,
	})

	tokenPaths := map[string]string{}
	for token, resource := range metadata.Resources {
		tokenPaths[token] = resource.Path
	}
	return &GenerationResult{
		Schema:                     &pkg,
		Metadata:                   &metadata,
		Examples:                   exampleMap,
		ForceNewTypes:              forceNewTypes,
		TypeCaseConflicts:          caseSensitiveTypes.findCaseConflicts(),
		FlattenedPropertyConflicts: flattenedPropertyConflicts,
		PathConflicts:              resourcesPathTracker.pathConflicts,
		TokenPaths:                 tokenPaths,
	}, nil
}

// resourcesPathConflictsTracker tracks resource path conflicts by module. Use newResourcesPathTracker to instantiate.
type resourcesPathConflictsTracker struct {
	pathConflicts map[openapi.ModuleName]map[openapi.ResourceName]map[string][]openapi.ApiVersion
}

func newResourcesPathConflictsTracker() *resourcesPathConflictsTracker {
	return &resourcesPathConflictsTracker{pathConflicts: map[openapi.ModuleName]map[openapi.ResourceName]map[string][]openapi.ApiVersion{}}
}

func (rpt *resourcesPathConflictsTracker) addPathConflictsForModule(moduleName openapi.ModuleName, resourcePaths map[openapi.ResourceName]map[string][]openapi.ApiVersion) {
	modulePathConflicts := map[openapi.ResourceName]map[string][]openapi.ApiVersion{}
	for resource, paths := range resourcePaths {
		if len(paths) > 1 {
			modulePathConflicts[resource] = paths
		}
	}
	if len(modulePathConflicts) > 0 {
		rpt.pathConflicts[moduleName] = modulePathConflicts
	}
}

func (rpt *resourcesPathConflictsTracker) hasConflicts() bool {
	return len(rpt.pathConflicts) > 0
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
	var nestedResourceSpecs []*openapi.ResourceSpec
	for _, otherResource := range util.MapOrdered(resourceSpecs) {
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
	pkg.Functions["azure-native:authorization:getClientConfig"] = pschema.FunctionSpec{
		Description: "Use this function to access the current configuration of the native Azure provider.",
		Outputs: &pschema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientConfig.",
			Properties: map[string]pschema.PropertySpec{
				"clientId": {
					Description: "Azure Client ID (Application Object ID).",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"objectId": {
					Description: "Azure Object ID of the current user or service principal.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"subscriptionId": {
					Description: "Azure Subscription ID",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"tenantId": {
					Description: "Azure Tenant ID",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
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
	pkg.Functions["azure-native:authorization:getClientToken"] = pschema.FunctionSpec{
		Description: "Use this function to get an Azure authentication token for the current login context.",
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"endpoint": {
					Description: "Optional authentication endpoint. Defaults to the endpoint of Azure Resource Manager.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
			},
			Type: "object",
		},
		Outputs: &pschema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientToken.",
			Properties: map[string]pschema.PropertySpec{
				"token": {
					Description: "OAuth token for Azure Management API and SDK authentication.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
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

	// Apply custom resource modifications to the schema and metadata.
	err := customresources.ApplySchemas(pkg, metadata)
	if err != nil {
		return err
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
	resources.VisitPackageSpecTypes(pkg, visitor)

	// Elide unused types.
	for typeName, t := range util.MapOrdered(pkg.Types) {
		if !usedTypes[typeName] {
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
	moduleName         openapi.ModuleName
	examples           map[string][]resources.AzureAPIExample
	apiVersion         *openapi.ApiVersion
	sdkVersion         openapi.SdkVersion
	allVersions        []openapi.ApiVersion
	versioning         Versioning
	caseSensitiveTypes caseSensitiveTokens
	// rootDir is used to resolve relative paths in the examples.
	rootDir       string
	forceNewTypes []ForceNewType
	// Token -> set<property>
	flattenedPropertyConflicts map[string]map[string]any
	majorVersion               int
	// A resource -> path -> API version map to record API paths per resource and later detect conflicts.
	// Each packageGenerator instance is only used for a single API version, so there won't be conflicting paths here.
	resourcePaths                  map[openapi.ResourceName]map[string]openapi.ApiVersion
	previousCompatibleTokensLookup *CompatibleTokensLookup
}

func (g *packageGenerator) genResources(typeName string, resource *openapi.ResourceSpec, nestedResourceBodyRefs []string) error {
	// Resource names should consistently start with upper case.
	typeNameAliases := g.genAliases(typeName)
	titleCasedTypeName := strings.Title(typeName)

	// A single API path can be modelled as several resources if it accepts a polymorphic payload:
	// i.e., when the request body is a discriminated union type of several object types. Pulumi
	// schema doesn't support polymorphic (OneOf) resources, so the provider creates a separate resource
	// per each union case. We call them "variants" in the code below.
	variants, err := g.findResourceVariants(resource)
	if err != nil {
		return errors.Wrapf(err, "resource %s.%s", g.moduleName, titleCasedTypeName)
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

func (g *packageGenerator) genAliases(typeName string) []string {
	switch g.majorVersion {
	case 2:
		// We need to alias the previous, lowercase name so users can upgrade to v2 without replacement.
		// These aliases aren't required anymore with v3.
		if strings.Title(typeName) != typeName {
			return []string{typeName}
		}
	case 3:
		// TODO: Add additional aliases for v3 case changes: https://github.com/pulumi/pulumi-azure-native/issues/3848
	}
	return nil
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

// recordPath adds path to keep track of all API paths a resource is mapped to.
func (g *packageGenerator) recordPath(typeName openapi.ResourceName, canonPath string, apiVersion openapi.ApiVersion) {
	// Some resources have a /default path, e.g., azure-native:azurestackhci:GuestAgent has conflicting paths
	// /subscriptions/{}/resourcegroups/{}/providers/microsoft.azurestackhci/virtualmachines/{}/guestagents/{},
	// /{}/providers/microsoft.azurestackhci/virtualmachineinstances/default/guestagents/default,
	// also azure-native:hybridcontainerservice:HybridIdentityMetadatum
	if strings.HasSuffix(canonPath, "/default") {
		return
	}

	// We use the map here only as a tuple of (path, apiVersion), it will only have a single key.
	g.resourcePaths[typeName] = map[string]openapi.ApiVersion{canonPath: apiVersion}
}

// mergeResourcePathsInto merges this packageGenerator's resource paths into the given map. This happens for each API
// version, so that in the end `resourcePaths` contains all paths for each resource and API version.
func (g *packageGenerator) mergeResourcePathsInto(resourcePaths map[openapi.ResourceName]map[string][]openapi.ApiVersion) {
	for resource, path := range g.resourcePaths {
		if _, ok := resourcePaths[resource]; !ok {
			resourcePaths[resource] = map[string][]openapi.ApiVersion{}
		}
		for path, apiVersion := range path {
			if _, ok := resourcePaths[resource][path]; !ok {
				resourcePaths[resource][path] = []openapi.ApiVersion{}
			}
			apiVersions := append(resourcePaths[resource][path], apiVersion)
			slices.Sort(apiVersions)
			resourcePaths[resource][path] = apiVersions
		}
	}
}

func (g *packageGenerator) genResourceVariant(apiSpec *openapi.ResourceSpec, resource *resourceVariant, nestedResourceBodyRefs []string, typeNameAliases ...string) error {
	module := g.moduleWithVersion()
	swagger := resource.Swagger
	path := resource.PathItem
	canonPath := paths.NormalizePath(resource.Path)

	typeName := resource.typeName

	resourceTok := generateTok(g.moduleName, typeName, g.sdkVersion)
	if !g.versioning.ShouldInclude(g.moduleName, g.apiVersion, typeName, resourceTok) {
		return nil
	}

	apiVersion := openapi.ApiVersion("default")
	if g.apiVersion != nil {
		apiVersion = *g.apiVersion
	}
	g.recordPath(typeName, canonPath, apiVersion)

	// Generate the resource.
	gen := moduleGenerator{
		pkg:                        g.pkg,
		metadata:                   g.metadata,
		module:                     module,
		moduleName:                 g.moduleName,
		resourceName:               resource.typeName,
		resourceToken:              resourceTok,
		visitedTypes:               make(map[string]bool),
		caseSensitiveTypes:         g.caseSensitiveTypes,
		inlineTypes:                map[*openapi.ReferenceContext]codegen.StringSet{},
		nestedResourceBodyRefs:     nestedResourceBodyRefs,
		flattenedPropertyConflicts: map[string]any{},
	}

	updateOp := path.Put
	updateMethod := ""
	if path.Put == nil {
		updateOp = path.Patch
		updateMethod = "PATCH"
	}

	resourceResponse, err := gen.genResponse(updateOp.Responses.StatusCodeResponses, swagger.ReferenceContext, resource.response, true)
	if err != nil {
		return errors.Wrapf(err, "failed to generate '%s': response type", resourceTok)
	}

	if resourceResponse == nil || len(resourceResponse.specs) == 0 {
		// Response is specified empty, do not generate a resource for it.
		return nil
	}

	parameters := resource.Swagger.MergeParameters(updateOp.Parameters, path.Parameters)
	autoNamer := resources.NewAutoNamer(resource.Path)

	resourceRequest, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, &autoNamer, resource.body, true)
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
			Description: g.formatDescription(resourceResponse.description, resource.typeName, openapi.ApiVersion(swagger.Info.Version), apiSpec.PreviousVersion, additionalDocs),
			Type:        "object",
			Properties:  resourceResponse.specs,
			Required:    resourceResponse.requiredSpecs.SortedValues(),
		},
		InputProperties:    resourceRequest.specs,
		RequiredInputs:     resourceRequest.requiredSpecs.SortedValues(),
		Aliases:            g.generateAliases(resourceTok, resource, typeNameAliases...),
		DeprecationMessage: resource.deprecationMessage,
	}
	g.pkg.Resources[resourceTok] = resourceSpec

	// Generate the function to get this resource.
	functionTok := fmt.Sprintf(`%s:%s:get%s`, g.pkg.Name, module, resource.typeName)
	if g.versioning.ShouldInclude(g.moduleName, g.apiVersion, resource.typeName, functionTok) {
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
		requestFunction, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, nil, resource.body, true)
		if err != nil {
			return errors.Wrapf(err, "failed to generate '%s': request type", functionTok)
		}
		responseFunction, err := gen.genResponse(readOp.Responses.StatusCodeResponses, swagger.ReferenceContext, resource.response, true)
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
				APIVersion:       swagger.Info.Version,
				Path:             resource.Path,
				GetParameters:    requestFunction.parameters,
				Response:         responseFunction.properties,
				IsResourceGetter: true,
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

	readUrlParams := defaultReadQueryParams[resourceTok]

	requiredContainers := mergeRequiredContainers(resourceRequest.requiredContainers, additionalRequiredContainers(resourceTok))

	r := resources.AzureAPIResource{
		APIVersion:           swagger.Info.Version,
		Path:                 resource.Path,
		UpdateMethod:         updateMethod,
		PutParameters:        resourceRequest.parameters,
		Response:             resourceResponse.properties,
		DefaultBody:          resource.DefaultBody,
		Singleton:            isSingleton(resource),
		PutAsyncStyle:        g.getAsyncStyle(updateOp),
		DeleteAsyncStyle:     g.getAsyncStyle(resource.PathItem.Delete),
		ReadMethod:           readMethod,
		ReadPath:             readPath,
		ReadQueryParams:      readUrlParams,
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

func isSingleton(resource *resourceVariant) bool {
	return resource.PathItem.Delete == nil || customresources.IsSingleton(resource.Path)
}

func (g *packageGenerator) generateAliases(resourceTok string, resource *resourceVariant, typeNameAliases ...string) []pschema.AliasSpec {
	typeAliases := collections.NewOrderableSet[string]()

	previousCompatibleTokens := g.previousCompatibleTokensLookup.FindCompatibleTokens(resource.ModuleNaming.ResolvedName, resource.typeName, resource.Path)
	for _, token := range previousCompatibleTokens {
		typeAliases.Add(token)
	}

	for _, alias := range typeNameAliases {
		typeAliases.Add(generateTok(g.moduleName, alias, g.sdkVersion))
	}

	// Add an alias for each API version that has the same path in it.
	for _, version := range resource.CompatibleVersions {
		typeAliases.Add(generateTok(g.moduleName, resource.typeName, version.ToSdkVersion()))

		// Add type name aliases for each compatible version.
		for _, alias := range typeNameAliases {
			typeAliases.Add(generateTok(g.moduleName, alias, version.ToSdkVersion()))
		}
	}

	var aliasSpecs []pschema.AliasSpec
	for _, v := range typeAliases.SortedValues() {
		// Skip aliasing to itself.
		if v != resourceTok {
			aliasSpecs = append(aliasSpecs, pschema.AliasSpec{Type: &v})
		}
	}
	return aliasSpecs
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
	for k, v := range util.MapOrdered(examples) {
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
	module := g.moduleWithVersion()
	gen := moduleGenerator{
		pkg:                        g.pkg,
		metadata:                   g.metadata,
		module:                     module,
		resourceToken:              fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, module, typeName),
		moduleName:                 g.moduleName,
		resourceName:               typeName,
		visitedTypes:               make(map[string]bool),
		caseSensitiveTypes:         g.caseSensitiveTypes,
		inlineTypes:                map[*openapi.ReferenceContext]codegen.StringSet{},
		flattenedPropertyConflicts: map[string]any{},
	}

	// Generate the function to get this resource.
	functionTok := generateTok(g.moduleName, typeName, g.sdkVersion)
	if !g.shouldInclude(typeName, functionTok, g.apiVersion) {
		return
	}

	parameters := swagger.MergeParameters(operation.Parameters, specParams)
	request, err := gen.genMethodParameters(parameters, swagger.ReferenceContext, nil, nil, false)
	if err != nil {
		log.Printf("failed to generate '%s': request type: %s", functionTok, err.Error())
		return
	}
	response, err := gen.genResponse(operation.Responses.StatusCodeResponses, swagger.ReferenceContext, nil, false)
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
		APIVersion:       swagger.Info.Version,
		Path:             path,
		PostParameters:   request.parameters,
		Response:         response.properties,
		IsResourceGetter: false,
	}
	g.metadata.Invokes[functionTok] = f
}

// TokenModule is the module as appears in the token e.g. `compute/v20200701` or `network`.
type TokenModule string

// moduleWithVersion produces the token's module part from the module name and the version e.g. `compute/v20200701` or `network`.
func (g *packageGenerator) moduleWithVersion() TokenModule {
	return TokenModule(versionedModule(g.moduleName, g.sdkVersion))
}

// goModuleName produces the *Go* module name from the provider name and the version e.g. `compute/v20200701`.
// or just the module name if the version is empty (default version) e.g. `compute`.
func goModuleName(moduleName openapi.ModuleName, sdkVersion openapi.SdkVersion) string {
	return filepath.Join(goModuleRepoPath, moduleName.Lowered(), goModuleVersion, string(sdkVersion))
}

// versionedModule produces the module name from the module name, and the version if not empty, e.g. `compute/v20200701`.
func versionedModule(moduleName openapi.ModuleName, sdkVersion openapi.SdkVersion) string {
	if sdkVersion == "" {
		return moduleName.Lowered()
	}
	return fmt.Sprintf("%s/%s", moduleName.Lowered(), sdkVersion)
}

func generateTok(moduleName openapi.ModuleName, typeName string, apiVersion openapi.SdkVersion) string {
	return fmt.Sprintf(`%s:%s:%s`, pulumiProviderName, versionedModule(moduleName, apiVersion), typeName)
}

func (g *packageGenerator) shouldInclude(typeName, tok string, version *openapi.ApiVersion) bool {
	return g.versioning.ShouldInclude(g.moduleName, version, typeName, tok)
}

func (g *packageGenerator) formatFunctionDescription(op *spec.Operation, typeName string, response *propertyBag, info *spec.Info) string {
	desc := response.description
	if op.Description != "" {
		desc = op.Description
	}
	return g.formatDescription(desc, typeName, openapi.ApiVersion(info.Version), "", nil)
}

func (g *packageGenerator) formatDescription(desc string, typeName string, defaultVersion, previousDefaultVersion openapi.ApiVersion, additionalDocs *string) string {
	var b strings.Builder
	b.WriteString(desc)

	if g.sdkVersion == "" {
		fmt.Fprintf(&b, "\nAzure REST API version: %s.", defaultVersion)
		if previousDefaultVersion != "" {
			fmt.Fprintf(&b, " Prior API version in Azure Native %d.x: %s.", g.majorVersion-1, previousDefaultVersion)
		}

		if g.majorVersion <= 2 {
			// List other available API versions, if any.
			allVersions := g.versioning.GetAllVersions(g.moduleName, typeName)
			includedVersions := []string{}
			for _, v := range allVersions {
				// Don't list the default version twice.
				if v == defaultVersion {
					continue
				}
				tok := generateTok(g.moduleName, typeName, openapi.ApiToSdkVersion(v))
				if g.shouldInclude(typeName, tok, &v) {
					includedVersions = append(includedVersions, string(v))
				}
			}

			if len(includedVersions) > 0 {
				fmt.Fprintf(&b, "\n\nOther available API versions: %s.", strings.Join(includedVersions, ", "))
			}
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

	// These operations from Microsoft.RecoveryServices/stable/2023-04-01/bms.json are incorrectly not marked as
	// long-running. https://github.com/Azure/azure-rest-api-specs/issues/31943
	if op.ID == "ProtectedItems_CreateOrUpdate" || op.ID == "ProtectedItems_Delete" {
		return extensionLongRunningDefault
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
	VersionedModule TokenModule
	Module          openapi.ModuleName
	ResourceName    string
	ReferenceName   string
	Property        string
}

type moduleGenerator struct {
	pkg                        *pschema.PackageSpec
	metadata                   *resources.AzureAPIMetadata
	module                     TokenModule
	moduleName                 openapi.ModuleName
	resourceName               string
	resourceToken              string
	visitedTypes               map[string]bool
	caseSensitiveTypes         caseSensitiveTokens
	inlineTypes                map[*openapi.ReferenceContext]codegen.StringSet
	nestedResourceBodyRefs     []string
	forceNewTypes              []ForceNewType
	flattenedPropertyConflicts map[string]any
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

func normalizeParamPattern(param *openapi.Parameter) string {
	// #3560: the regex has the wrong max length in the spec.
	// See portal/resource-manager/Microsoft.Portal/preview/2022-12-01-preview/portal.json#L185-L192
	if param.Name == "dashboardName" && *param.MaxLength > 24 && param.Pattern == "^[a-zA-Z0-9-]{3,24}$" {
		return fmt.Sprintf("^[a-zA-Z0-9-]{3,%d}$", *param.MaxLength)
	}
	return param.Pattern
}

func (m *moduleGenerator) genMethodParameters(parameters []spec.Parameter, ctx *openapi.ReferenceContext,
	namer *resources.AutoNamer, bodySchema *openapi.Schema, isResourceGetter bool) (*parameterBag, error) {
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
				Pattern:   normalizeParamPattern(param),
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
		case m.paramIsSetByProvider(param, bodySchema):
			continue // No need to include params like API version in the schema or meta, they are added automatically by the provider.
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
			props, err := m.genProperties(bodySchema, genPropertiesVariant{
				isTopLevel: isResourceGetter,
				isOutput:   false,
				isType:     false,
				isResponse: false,
			})
			if err != nil {
				return nil, err
			}

			// Top-level location is never required: it can be derived from a config value or the parent resource group.
			props.requiredSpecs.Delete("location")

			// #3757 StaticSiteLinkedBackend has a non-required property that is required in the API.
			// https://github.com/Azure/azure-rest-api-specs/issues/31807
			if m.resourceToken == "azure-native:web:StaticSiteLinkedBackend" {
				if _, ok := props.properties["region"]; ok {
					props.requiredProperties.Add("region")
					props.requiredSpecs.Add("region")
				}
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

// paramIsSetByProvider returns true if this parameter is set _only_ by the provider, so it shouldn't be visible to
// users. Currently, that's `api-version`.
func (m *moduleGenerator) paramIsSetByProvider(param *openapi.Parameter, bodySchema *openapi.Schema) bool {
	if param.Name != "api-version" || param.In != "query" {
		return false
	}

	// #3524: for this resource, the api-version is passed through to another Azure service, so the user needs to specify it.
	if m.moduleName == "Resources" && m.resourceName == "Resource" {
		return false
	}

	// #3524: for this resource, the api-version is passed through to another Azure service, so the user needs to specify it.
	if m.moduleName == "Authorization" && m.resourceName == "ManagementLockAtResourceLevel" {
		return false
	}

	return true
}

func (m *moduleGenerator) genResponse(statusCodeResponses map[int]spec.Response, ctx *openapi.ReferenceContext,
	responseSchema *openapi.Schema, isResourceGetter bool) (*propertyBag, error) {

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

	result, err := m.genProperties(responseSchema, genPropertiesVariant{
		isTopLevel: isResourceGetter,
		isOutput:   true,
		isType:     false,
		isResponse: true,
	})
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

	description = strings.Map(func(r rune) rune {
		// Remove Line Separator and Paragraph Separator characters, they break dotnet build.
		if r == 0x2028 || r == 0x2029 {
			return -1 // negative value means drop the character
		}
		return r
	}, description)

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
	contract.AssertNoErrorf(err, "failed to marshal '%v'", v)
	return bytes
}

// InMemoryPackageLoader prevents having to fetch the schema from
// the provider every time which significantly speeds up codegen.
func InMemoryPackageLoader(pkgs map[string]*pschema.Package) pschema.Loader {
	return &inMemoryLoader{pkgs: pkgs}
}

type inMemoryLoader struct {
	pkgs map[string]*pschema.Package
}

func (l *inMemoryLoader) LoadPackage(pkg string, _ *semver.Version) (*pschema.Package, error) {
	if p, ok := l.pkgs[pkg]; ok {
		return p, nil
	}

	return nil, errors.Errorf("package %s not found in the in-memory map", pkg)
}

func (l *inMemoryLoader) LoadPackageV2(ctx context.Context, descriptor *pschema.PackageDescriptor) (*pschema.Package, error) {
	if p, ok := l.pkgs[descriptor.Name]; ok {
		return p, nil
	}

	return nil, errors.Errorf("package %s not found in the in-memory map", descriptor.Name)
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
