// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	azureEnv "github.com/Azure/go-autorest/autorest/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"

	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// OriginalStateKey is a map key used to store the original state of a resource.
// It can be used to reset the resource to its original state when the resource is deleted, or to restore individual
// properties of a resource to their defaults when they are removed from the resource.
const OriginalStateKey = "__orig_state"

type CustomReadFunc func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error)

// CustomResource is a manual SDK-based implementation of a (part of) resource when Azure API is missing some
// crucial operations.
type CustomResource struct {
	path string
	// tok is the resource token for this custom resource, used both for looking up the resource in the schema and for
	// the resource name. To use the tok for schema lookup but give the resource a different name, set CustomResourceName.
	tok string
	// Types are net-new auxiliary types defined for this resource. Optional.
	// Deprecated: Use Schema instead.
	Types map[string]schema.ComplexTypeSpec
	// TypeOverrides define types that already exist in the auto-generated schema but we want to override
	// to our custom shape and behavior. Optional.
	// Deprecated: Use Schema instead.
	TypeOverrides map[string]schema.ComplexTypeSpec
	// Resource schema. Optional, by default the schema is assumed to be included in Azure Open API specs.
	// Deprecated: Use Schema instead.
	LegacySchema *schema.ResourceSpec
	// Resource & types, schema & metadata modifications function.
	// Optional, runs after main schema is generated allowing for modifications.
	// The resource and types (both schema and metadata) returned will be written back into the schema and metadata.
	// This can also include new types which were not present in the original schema.
	// Any dangling, unreferenced types will be removed from the schema automatically.
	// Returning nil will skip the resource and make no changes.
	Schema func(resource *ResourceDefinition) (*ResourceDefinition, error)
	// Resource metadata. Defines the parameters and properties that are used for diff calculation and validation.
	// Optional, by default the metadata is assumed to be derived from Azure Open API specs.
	// Deprecated: Use Schema instead.
	Meta *AzureAPIResource
	// MetaTypes are net-new auxiliary metadata types defined for this resource. Optional.
	// Deprecated: Use Schema instead.
	MetaTypes map[string]AzureAPIType
	// MetaTypeOverrides define meta types that already exist in the auto-generated metadata but we want to override
	// to our custom shape and behavior. Optional.
	// Deprecated: Use Schema instead.
	MetaTypeOverrides map[string]AzureAPIType
	// GetIdAndQuery returns the resource id and query parameters for the custom resource. Optional, by default the provider gets the
	// resource ID and query parameters from crud.ResourceCrudClient.PrepareAzureRESTIdAndQuery.
	GetIdAndQuery func(ctx context.Context, inputs resource.PropertyMap, crudClient crud.ResourceCrudClient) (string, map[string]any, error)
	// Create a new resource from a map of input values. Returns a map of resource outputs that match the schema shape.
	Create    func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]interface{}, error)
	CanCreate func(ctx context.Context, id string) error
	// Read the state of an existing resource. Constructs the resource ID based on input values. Returns a map of
	// resource outputs. If the requested resource does not exist, the second result is false. In that case, the
	// error must be nil.
	Read CustomReadFunc
	// Update an existing resource with a map of input values. Returns a map of resource outputs that match the schema shape.
	Update func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]interface{}, error)
	// Delete an existing resource. Constructs the resource ID based on input values.
	Delete func(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error

	// Only used when we want to specify a `tok` that exists in the spec so we can look up the resource, but we want
	// the custom resource's name to be different.
	CustomResourceName string

	// Optional. If specified, only include schema and metadata from the spec for this version. Has no effect for
	// resources that don't use the Azure spec. Format is 2020-01-30 or 2020-01-30-preview.
	apiVersion *string

	// IsSingleton is true if the resource is a singleton resource that cannot be created or deleted, only initialized
	// and reset to a default state. Normally, we infer this from whether the `Delete` property is set. In some cases
	// we need to set it explicitly if the resource is a singleton but does have a `Delete` property implementing a
	// custom reset to the default state.
	isSingleton bool
}

// ResourceDefinition is a combination of the external schema and runtime metadata
// for a resource and its associated types.
type ResourceDefinition struct {
	Resource     schema.ResourceSpec
	Types        map[string]schema.ComplexTypeSpec
	MetaResource AzureAPIResource
	MetaTypes    map[string]AzureAPIType
}

// makePropertyNotRequired removes a property from the various lists of required properties, but leaves the property itself.
func (r *ResourceDefinition) makePropertyNotRequired(property string) {
	r.Resource.RequiredInputs = util.RemoveFromSlice(r.Resource.RequiredInputs, property)
	r.Resource.ObjectTypeSpec.Required = util.RemoveFromSlice(r.Resource.ObjectTypeSpec.Required, property)
	if body, hasBody := r.MetaResource.BodyParameter(); hasBody {
		body.Body.RequiredProperties = util.RemoveFromSlice(body.Body.RequiredProperties, property)
	}
}

// ApplySchemas applies custom schema modifications to the given package.
// These modifications should never overlap with each other, but we apply in a deterministic order to ensure
// that the end result of the modifications is consistent.
func ApplySchemas(pkg *schema.PackageSpec, meta *AzureAPIMetadata) error {
	for _, r := range util.MapOrdered(featureLookup) {
		if err := r.ApplySchema(pkg, meta); err != nil {
			return fmt.Errorf("failed to apply schema modifications for resource %s: %w", r.tok, err)
		}
	}
	return nil
}

func (r *CustomResource) ApplySchema(pkg *schema.PackageSpec, meta *AzureAPIMetadata) error {
	if r.tok == "" || r.Schema == nil {
		return nil
	}

	tok := r.tok
	if r.CustomResourceName != "" {
		module, version, _, err := resources.ParseToken(tok)
		if err != nil {
			return fmt.Errorf("failed to parse token %s: %w", tok, err)
		}
		tok = resources.BuildToken(module, version, r.CustomResourceName)
	}

	existingResource, resourceAlreadyExists := pkg.Resources[tok]
	var originalResource *ResourceDefinition
	types := map[string]schema.ComplexTypeSpec{} // Hoist scope for easy lookup when checking if the type is new.
	if resourceAlreadyExists {
		resourceMeta, resourceMetadataFound := meta.Resources[tok]
		if !resourceMetadataFound {
			return fmt.Errorf("metadata for resource %s not found", tok)
		}
		VisitResourceTypes(pkg, tok, func(tok string, t schema.ComplexTypeSpec) {
			// Capture referenced schema types
			types[tok] = t
		})
		metaTypes := map[string]AzureAPIType{}
		for n := range types {
			if mt, metaTypeFound := meta.Types[n]; metaTypeFound {
				metaTypes[n] = mt
			} else {
				logging.V(9).Infof("metadata for type %s not found, assuming it's an enum", n)
			}
		}
		originalResource = &ResourceDefinition{
			Resource:     existingResource,
			Types:        types,
			MetaResource: resourceMeta,
			MetaTypes:    metaTypes,
		}
	}

	customResource, err := r.Schema(originalResource)
	if err != nil {
		if !resourceAlreadyExists && len(pkg.Resources) < 100 {
			logging.V(5).Infof("warning: error applying custom resource schema modifications for %s: '%s'; continuing since this is likely a test run with only %d resources",
				tok, err, len(pkg.Resources))
			return nil
		}
		return fmt.Errorf("failed to apply custom resource schema modifications for %s: %w", tok, err)
	}
	if customResource == nil {
		return nil
	}

	pkg.Resources[tok] = customResource.Resource
	for typeToken, typeSchema := range customResource.Types {
		if _, isExistingType := types[typeToken]; !isExistingType {
			// New type, ensure it doesn't conflict
			if _, existsInPackage := pkg.Types[typeToken]; existsInPackage {
				return fmt.Errorf("new type from custom resource modification %s already exists in package", typeToken)
			}
		}
		pkg.Types[typeToken] = typeSchema
	}
	meta.Resources[tok] = customResource.MetaResource
	for n, t := range customResource.MetaTypes {
		// Skip validation of meta types as they should be 1:1 with schema types
		meta.Types[n] = t
	}
	return nil
}

// BuildCustomResources creates a map of custom resources for given environment parameters.
func BuildCustomResources(env *azureEnv.Environment,
	azureClient azure.AzureClient,
	lookupResource ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	subscriptionID string,
	bearerAuth autorest.Authorizer,
	tokenAuth autorest.Authorizer,
	kvBearerAuth autorest.Authorizer,
	userAgent string,
	tokenCred azcore.TokenCredential) (map[string]*CustomResource, error) {

	armKVClient, err := armkeyvault.NewVaultsClient(subscriptionID, tokenCred, &arm.ClientOptions{})
	if err != nil {
		return nil, err
	}

	customWebApp, err := webApp(crudClientFactory, azureClient, lookupResource)
	if err != nil {
		return nil, err
	}
	customWebAppSlot, err := webAppSlot(crudClientFactory, azureClient, lookupResource)
	if err != nil {
		return nil, err
	}

	postgresConf, err := postgresFlexibleServerConfiguration(crudClientFactory, lookupResource)
	if err != nil {
		return nil, err
	}

	protectedItem, err := recoveryServicesProtectedItem(subscriptionID, tokenCred)
	if err != nil {
		return nil, err
	}

	pimRoleManagementPolicy, err := pimRoleManagementPolicy(lookupResource, crudClientFactory)
	if err != nil {
		return nil, err
	}

	pimRoleEligibilitySchedule, err := pimRoleEligibilitySchedule(lookupResource, crudClientFactory, tokenCred)
	if err != nil {
		return nil, err
	}

	resources := []*CustomResource{
		keyVaultAccessPolicy(armKVClient),

		// Customization of regular resources
		blobContainerLegalHold(azureClient),
		portalDashboard(),
		customWebApp,
		customWebAppSlot,
		postgresConf,
		protectedItem,
		pimRoleManagementPolicy,
		pimRoleEligibilitySchedule,
		cdnCustomDomainHttps(azureClient),
	}

	// For Key Vault, we need to use separate token sources for azidentity and for the legacy auth. The
	// `azCoreTokenCredential` adapter that we use elsewhere to translate legacy token sources to azidentity doesn't
	// work here because KV needs a different token source for the KV endpoint.
	if util.EnableAzcoreBackend() {
		resources = append(resources, keyVaultSecret(env.KeyVaultDNSSuffix, tokenCred))
		resources = append(resources, keyVaultKey(env.KeyVaultDNSSuffix, tokenCred))

		cloud := azure.GetCloudByName(env.Name)
		resources = append(resources, storageAccountStaticWebsite_azidentity(cloud, tokenCred))
		resources = append(resources, newBlob_azidentity(cloud, tokenCred))
	} else {
		kvClient := keyvault.New()
		kvClient.Authorizer = kvBearerAuth
		kvClient.UserAgent = userAgent
		resources = append(resources, keyVaultSecret_autorest(env.KeyVaultDNSSuffix, &kvClient))
		resources = append(resources, keyVaultKey_autorest(env.KeyVaultDNSSuffix, &kvClient))

		// Storage resources.
		storageAccountsClient := storage.NewAccountsClientWithBaseURI(env.ResourceManagerEndpoint, subscriptionID)
		storageAccountsClient.Authorizer = tokenAuth
		storageAccountsClient.UserAgent = userAgent
		resources = append(resources, newStorageAccountStaticWebsite(env, &storageAccountsClient))
		resources = append(resources, newBlob(env, &storageAccountsClient))
	}

	result := map[string]*CustomResource{}
	for _, r := range resources {
		result[r.path] = r
	}
	return result, nil
}

// featureLookup is a map of custom resource to lookup their capabilities.
var featureLookup, _ = BuildCustomResources(&azureEnv.Environment{}, nil, nil, nil, "", nil, nil, nil, "", nil)

// IncludeCustomResource returns isCustom=true if a custom resource is defined for the given path, and include=true if
// the given API version should be included.
// Note on the `apiVersion` parameter: using the `openapi.ApiVersion` type would create an import cycle.
func IncludeCustomResource(path, apiVersion string) (isCustom, include bool) {
	var resource *CustomResource
	resource, isCustom = featureLookup[path]
	if !isCustom {
		return
	}

	include = resource.apiVersion == nil || *resource.apiVersion == apiVersion
	return
}

// HasCustomDelete returns true if a custom DELETE operation is defined for a given API path.
func HasCustomDelete(path string) bool {
	if res, ok := featureLookup[path]; ok {
		return res.Delete != nil
	}
	return false
}

func IsSingleton(path string) bool {
	if res, ok := featureLookup[path]; ok {
		return res.isSingleton
	}
	return false
}

func GetCustomResourceName(path string) (string, bool) {
	if res, ok := featureLookup[path]; ok {
		return res.CustomResourceName, res.CustomResourceName != ""
	}
	return "", false
}

// SchemaMixins returns the map of custom resource schema definitions per resource token.
func SchemaMixins() map[string]schema.ResourceSpec {
	specs := map[string]schema.ResourceSpec{}
	for _, r := range featureLookup {
		if r.tok != "" && r.LegacySchema != nil {
			specs[r.tok] = *r.LegacySchema
		}
	}
	return specs
}

// SchemaTypeMixins returns the map of custom resource schema definitions per resource token.
func SchemaTypeMixins() map[string]schema.ComplexTypeSpec {
	types := map[string]schema.ComplexTypeSpec{}
	for _, r := range featureLookup {
		for n, v := range r.Types {
			types[n] = v
		}
	}
	return types
}

// SchemaTypeOverrides returns the map of custom resource schema overrides per resource token.
func SchemaTypeOverrides() map[string]schema.ComplexTypeSpec {
	types := map[string]schema.ComplexTypeSpec{}
	for _, r := range featureLookup {
		for n, v := range r.TypeOverrides {
			types[n] = v
		}
	}
	return types
}

// MetaMixins returns the map of custom resource metadata definitions per resource token.
func MetaMixins() map[string]AzureAPIResource {
	meta := map[string]AzureAPIResource{}
	for _, r := range featureLookup {
		if r.tok != "" && r.Meta != nil {
			meta[r.tok] = *r.Meta
		}
	}
	return meta
}

// MetaMixins returns the map of custom resource metadata definitions per resource token.
func MetaTypeMixins() map[string]AzureAPIType {
	types := map[string]AzureAPIType{}
	for _, r := range featureLookup {
		for n, v := range r.MetaTypes {
			types[n] = v
		}
	}
	return types
}

// MetaTypeOverrides returns the map of custom resource metadata overrides per resource token.
func MetaTypeOverrides() map[string]AzureAPIType {
	types := map[string]AzureAPIType{}
	for _, r := range featureLookup {
		for n, v := range r.MetaTypeOverrides {
			types[n] = v
		}
	}
	return types
}

// createCrudClient creates a CRUD client for the given resource type, looking up the fully
// qualified `resourceToken` like "azure-native:web:WebApp" via `lookupResource`.
func createCrudClient(
	crudClientFactory crud.ResourceCrudClientFactory, lookupResource ResourceLookupFunc, resourceToken string,
) (crud.ResourceCrudClient, error) {
	res, ok, err := lookupResource(resourceToken)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("resource %s not found", resourceToken)
	}

	return crudClientFactory(&res), nil
}
