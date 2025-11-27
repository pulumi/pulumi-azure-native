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
// CREATE is also customized to recover soft-deleted secrets/keys before creation.

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(cloud cloud.Configuration, tokenCred azcore.TokenCredential) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
		CanCreate: func(ctx context.Context, id string) error {
			// We override CanCreate to handle soft-deleted secrets.
			// If a soft-deleted secret exists, we'll recover it in the Create function.
			// This prevents the default CanCreate from failing on soft-deleted resources.
			return nil
		},
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]interface{}, error) {
			vaultName := inputs["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return nil, errors.New("vaultName not found in inputs")
			}
			secretName := inputs["secretName"]
			if !secretName.HasValue() || !secretName.IsString() {
				return nil, errors.New("secretName not found in inputs")
			}

			keyVaultDNSSuffix := strings.TrimPrefix(cloud.Suffixes.KeyVaultDNS, ".")
			if keyVaultDNSSuffix == "" {
				return nil, errors.New("The provider configuration must include a value for keyVaultDNSSuffix")
			}
			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azsecrets.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return nil, err
			}

			// Check if a soft-deleted secret exists and recover it if so
			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			deletedSecret, err := kvClient.GetDeletedSecret(ctx, secretName.StringValue(), nil)
			if err == nil && deletedSecret.RecoveryID != nil {
				logging.V(5).Infof("Found soft-deleted secret %s, recovering it before creating", secretName.StringValue())
				_, err = kvClient.RecoverDeletedSecret(ctx, secretName.StringValue(), nil)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to recover soft-deleted secret %s", secretName.StringValue())
				}
				logging.V(5).Infof("Successfully recovered soft-deleted secret %s", secretName.StringValue())
			} else if err != nil {
				// Check if the error is a 404 (secret not in deleted state), which is fine
				if respErr, ok := err.(*azcore.ResponseError); ok && respErr.StatusCode == http.StatusNotFound {
					// Secret doesn't exist in deleted state, continue with normal creation
					logging.V(9).Infof("Secret %s not found in deleted state, proceeding with normal creation", secretName.StringValue())
				} else {
					// Some other error occurred while checking for deleted secret
					logging.V(5).Infof("Warning: error checking for deleted secret %s: %v. Continuing with creation attempt.", secretName.StringValue(), err)
				}
			}

			// Return nil to let the default ARM-based creation flow handle the actual creation
			return nil, nil
		},
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
		CanCreate: func(ctx context.Context, id string) error {
			// We override CanCreate to handle soft-deleted keys.
			// If a soft-deleted key exists, we'll recover it in the Create function.
			// This prevents the default CanCreate from failing on soft-deleted resources.
			return nil
		},
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]interface{}, error) {
			vaultName := inputs["vaultName"]
			if !vaultName.HasValue() || !vaultName.IsString() {
				return nil, errors.New("vaultName not found in inputs")
			}
			keyName := inputs["keyName"]
			if !keyName.HasValue() || !keyName.IsString() {
				return nil, errors.New("keyName not found in inputs")
			}

			keyVaultDNSSuffix := strings.TrimPrefix(cloud.Suffixes.KeyVaultDNS, ".")
			if keyVaultDNSSuffix == "" {
				return nil, errors.New("The provider configuration must include a value for keyVaultDNSSuffix")
			}
			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)
			kvClient, err := azkeys.NewClient(vaultUrl, tokenCred, nil)
			if err != nil {
				return nil, err
			}

			// Check if a soft-deleted key exists and recover it if so
			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			deletedKey, err := kvClient.GetDeletedKey(ctx, keyName.StringValue(), nil)
			if err == nil && deletedKey.RecoveryID != nil {
				logging.V(5).Infof("Found soft-deleted key %s, recovering it before creating", keyName.StringValue())
				_, err = kvClient.RecoverDeletedKey(ctx, keyName.StringValue(), nil)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to recover soft-deleted key %s", keyName.StringValue())
				}
				logging.V(5).Infof("Successfully recovered soft-deleted key %s", keyName.StringValue())
			} else if err != nil {
				// Check if the error is a 404 (key not in deleted state), which is fine
				if respErr, ok := err.(*azcore.ResponseError); ok && respErr.StatusCode == http.StatusNotFound {
					// Key doesn't exist in deleted state, continue with normal creation
					logging.V(9).Infof("Key %s not found in deleted state, proceeding with normal creation", keyName.StringValue())
				} else {
					// Some other error occurred while checking for deleted key
					logging.V(5).Infof("Warning: error checking for deleted key %s: %v. Continuing with creation attempt.", keyName.StringValue(), err)
				}
			}

			// Return nil to let the default ARM-based creation flow handle the actual creation
			return nil, nil
		},
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
