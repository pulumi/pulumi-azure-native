// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// keyVaultSecret_autorest creates a custom resource for Azure KeyVault Secrets, based on the
// deprecated autorest and go-azure-helpers backend.
func keyVaultSecret_autorest(keyVaultDNSSuffix string, kvClient *keyvault.BaseClient) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
		Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
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
			return reportDeletionError_autorest(err)
		},
	}
}

// keyVaultKey_autorest creates a custom resource for Azure KeyVault Keys, based on the
// deprecated autorest and go-azure-helpers backend.
func keyVaultKey_autorest(keyVaultDNSSuffix string, kvClient *keyvault.BaseClient) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}",
		Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
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
			return reportDeletionError_autorest(err)
		},
	}
}

func reportDeletionError_autorest(err error) error {
	if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
		return nil
	}
	return err
}
