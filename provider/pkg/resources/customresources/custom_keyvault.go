// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Note: These are "hybrid" resources: schema, create, update, read use default implementation, while DELETE is overridden.
// DELETE is implemented via the KeyVault "data plane" API, because it isn't available via ARM.

// from: https://github.com/Azure/go-autorest/blob/autorest/v0.11.29/autorest/azure/environments.go
func getKeyVaultSuffix(cloud azcloud.Configuration) (string, error) {
	suffix := "vault.azure.net"
	if cloud.ActiveDirectoryAuthorityHost == azcloud.AzureChina.ActiveDirectoryAuthorityHost {
		suffix = "vault.azure.cn"
	} else if cloud.ActiveDirectoryAuthorityHost == azcloud.AzureGovernment.ActiveDirectoryAuthorityHost {
		suffix = "vault.usgovcloudapi.net"
	}
	return suffix, nil
}

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(cloud azcloud.Configuration, tokenCred azcore.TokenCredential) *CustomResource {
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

			keyVaultDNSSuffix, err := getKeyVaultSuffix(cloud)
			if err != nil {
				return err
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
func keyVaultKey(cloud azcloud.Configuration, tokenCred azcore.TokenCredential) *CustomResource {
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

			keyVaultDNSSuffix, err := getKeyVaultSuffix(cloud)
			if err != nil {
				return err
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
