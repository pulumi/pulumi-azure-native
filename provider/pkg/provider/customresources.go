package provider

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"net/http"
)

// customResource is a manual SDK-based implementation of a (part of) resource when Azure API is missing some
// crucial operations.
type customResource struct {
	path   string
	delete func(context.Context, resource.PropertyMap) error
}

func reportDeletionError(err error) error {
	if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
		return nil
	}
	return err
}

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(keyVaultDNSSuffix string, kvClient *keyvault.BaseClient) *customResource {
	return &customResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
		delete: func(ctx context.Context, properties resource.PropertyMap) error {
			vaultName := properties["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			secretName := properties["secretName"]
			if !secretName.HasValue() || !secretName.IsString() {
				return errors.New("secretName not found in resource state")
			}

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			_, err := kvClient.DeleteSecret(ctx, vaultUrl, secretName.StringValue())
			return reportDeletionError(err)
		},
	}
}

// keyVaultKey creates a custom resource for Azure KeyVault Key.
func keyVaultKey(keyVaultDNSSuffix string, kvClient *keyvault.BaseClient) *customResource {
	return &customResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}",
		delete: func(ctx context.Context, properties resource.PropertyMap) error {
			vaultName := properties["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			keyName := properties["keyName"]
			if !keyName.HasValue() || !keyName.IsString() {
				return errors.New("keyName not found in resource state")
			}

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			_, err := kvClient.DeleteKey(ctx, vaultUrl, keyName.StringValue())
			return reportDeletionError(err)
		},
	}
}

// buildCustomResources creates a map of custom resources for given environment parameters.
func buildCustomResources(env *azure.Environment, bearerAuth autorest.Authorizer,
	userAgent string) map[string]*customResource {

	kvClient := keyvault.New()
	kvClient.Authorizer = bearerAuth
	kvClient.UserAgent = userAgent

	resources := []*customResource{
		// Azure KeyVault resources
		keyVaultSecret(env.KeyVaultDNSSuffix, &kvClient),
		// Doesn't work yet in our tests. We need to figure this out with Azure before publishing.
		keyVaultKey(env.KeyVaultDNSSuffix, &kvClient),
	}

	result := map[string]*customResource{}
	for _, r := range resources {
		result[r.path] = r
	}
	return result
}

// featureLookup is a map of custom resource to lookup their capabilities.
var featureLookup = buildCustomResources(&azure.Environment{}, nil, "")

// HasCustomDelete returns true if a custom DELETE operation is defined for a given API path.
func HasCustomDelete(path string) bool {
	if res, ok := featureLookup[path]; ok {
		return res.delete != nil
	}
	return false
}
