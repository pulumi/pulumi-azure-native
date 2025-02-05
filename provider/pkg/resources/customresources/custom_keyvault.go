// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(keyVaultDNSSuffix string, tokenCred azcore.TokenCredential) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			vaultName := properties["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			secretName := properties["secretName"]
			if !secretName.HasValue() || !secretName.IsString() {
				return errors.New("secretName not found in resource state")
			}

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azsecrets.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return err
			}
			_, err = kvClient.DeleteSecret(ctx, secretName.StringValue(), nil)
			return reportDeletionError(err)
		},
	}
}

// keyVaultKey creates a custom resource for Azure KeyVault Key.
func keyVaultKey(keyVaultDNSSuffix string, tokenCred azcore.TokenCredential) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}",
		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			vaultName := properties["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return errors.New("vaultName not found in resource state")
			}
			keyName := properties["keyName"]
			if !keyName.HasValue() || !keyName.IsString() {
				return errors.New("keyName not found in resource state")
			}

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azkeys.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return err
			}
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
