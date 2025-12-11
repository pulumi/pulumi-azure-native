// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// Note: These are "hybrid" resources: schema, create, update, read use default implementation, while DELETE is overridden.
// DELETE is implemented via the KeyVault "data plane" API, because it isn't available via ARM.

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(cloud cloud.Configuration, tokenCred azcore.TokenCredential) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
		Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
			vaultName := inputs["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			secretName := inputs["secretName"]
			if !secretName.HasValue() || !secretName.IsString() {
				return errors.New("secretName not found in resource state")
			}

			keyVaultDNSSuffix := strings.TrimPrefix(cloud.Suffixes.KeyVaultDNS, ".")
			if keyVaultDNSSuffix == "" {
				return errors.New("The provider configuration must include a value for keyVaultDNSSuffix")
			}
			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azsecrets.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return err
			}
			logging.V(9).Infof("connecting to vault: %s", vaultUrl)

			_, err = kvClient.DeleteSecret(ctx, secretName.StringValue(), nil)
			return reportDeletionError(err)
		},
	}
}

// keyVaultKey creates a custom resource for Azure KeyVault Key.
func keyVaultKey(cloud cloud.Configuration, tokenCred azcore.TokenCredential) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}",
		Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
			vaultName := inputs["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			keyName := inputs["keyName"]
			if !keyName.HasValue() || !keyName.IsString() {
				return errors.New("keyName not found in resource state")
			}

			keyVaultDNSSuffix := strings.TrimPrefix(cloud.Suffixes.KeyVaultDNS, ".")
			if keyVaultDNSSuffix == "" {
				return errors.New("The provider configuration must include a value for keyVaultDNSSuffix")
			}
			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azkeys.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return err
			}
			logging.V(9).Infof("connecting to vault: %s", vaultUrl)

			_, err = kvClient.DeleteKey(ctx, keyName.StringValue(), nil)
			return reportDeletionError(err)
		},
	}
}

func reportDeletionError(err error) error {
	if detailed, ok := err.(*azcore.ResponseError); ok && detailed.StatusCode == http.StatusNotFound {
		return nil
	}
	return err
}
