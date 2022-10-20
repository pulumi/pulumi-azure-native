// Copyright 2021, Pulumi Corporation.  All rights reserved.

package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/manicminer/hamilton/environments"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// TODO,tkappler move to a new auth package
type AuthorizerFactory func(api environments.Api) (autorest.Authorizer, error)

// CustomResource is a manual SDK-based implementation of a (part of) resource when Azure API is missing some
// crucial operations.
type CustomResource struct {
	path string
	tok  string
	// Auxiliary types defined for this resource. Optional.
	Types map[string]schema.ComplexTypeSpec
	// Resource schema. Optional, by default the schema is assumed to be included in Azure Open API specs.
	Schema *schema.ResourceSpec
	// Resource metadata. Defines the parameters and properties that are used for diff calculation and validation.
	// Optional, by default the metadata is assumed to be derived from Azure Open API specs.
	Meta *AzureAPIResource
	// Create a new resource from a map of input values. Returns a map of resource outputs that match the schema shape.
	Create func(context.Context, resource.PropertyMap) (map[string]interface{}, error)
	// Read the state of an existing resource. Constructs the resource ID based on input values. Returns a map of
	// resource outputs. If the requested resource does not exist, the second result is false.
	Read func(context.Context, resource.PropertyMap) (map[string]interface{}, bool, error)
	// Update an existing resource with a map of input values. Returns a map of resource outputs that match the schema shape.
	Update func(context.Context, resource.PropertyMap) (map[string]interface{}, error)
	// Delete an existing resource. Constructs the resource ID based on input values.
	Delete func(context.Context, resource.PropertyMap) error
}

// BuildCustomResources creates a map of custom resources for given environment parameters.
func BuildCustomResources(env *azure.Environment,
	subscriptionID string,
	bearerAuth autorest.Authorizer,
	tokenAuth autorest.Authorizer,
	kvBearerAuth autorest.Authorizer,
	bearerAuthFactory AuthorizerFactory,
	userAgent string) map[string]*CustomResource {

	kvClient := keyvault.New()
	kvClient.Authorizer = kvBearerAuth
	kvClient.UserAgent = userAgent

	storageAccountsClient := storage.NewAccountsClientWithBaseURI(env.ResourceManagerEndpoint, subscriptionID)
	storageAccountsClient.Authorizer = tokenAuth
	storageAccountsClient.UserAgent = userAgent

	resources := []*CustomResource{
		// Azure KeyVault resources.
		keyVaultSecret(env.KeyVaultDNSSuffix, &kvClient),
		keyVaultKey(env.KeyVaultDNSSuffix, &kvClient),
		// Storage resources.
		newStorageAccountStaticWebsite(env, &storageAccountsClient),
		newBlob(subscriptionID, env, &storageAccountsClient, bearerAuthFactory),
	}

	result := map[string]*CustomResource{}
	for _, r := range resources {
		result[r.path] = r
	}
	return result
}

// featureLookup is a map of custom resource to lookup their capabilities.
var featureLookup = BuildCustomResources(&azure.Environment{}, "", nil, nil, nil, nil, "")

// HasCustomDelete returns true if a custom DELETE operation is defined for a given API path.
func HasCustomDelete(path string) bool {
	if res, ok := featureLookup[path]; ok {
		return res.Delete != nil
	}
	return false
}

// SchemaMixins returns the map of custom resource schema definitions per resource token.
func SchemaMixins() map[string]schema.ResourceSpec {
	specs := map[string]schema.ResourceSpec{}
	for _, r := range featureLookup {
		if r.tok != "" && r.Schema != nil {
			specs[r.tok] = *r.Schema
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

// SchemaMixins returns the map of custom resource metadata definitions per resource token.
func MetaMixins() map[string]AzureAPIResource {
	meta := map[string]AzureAPIResource{}
	for _, r := range featureLookup {
		if r.tok != "" && r.Meta != nil {
			meta[r.tok] = *r.Meta
		}
	}
	return meta
}
