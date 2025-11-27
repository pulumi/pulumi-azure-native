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
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// keyVaultSecret_autorest creates a custom resource for Azure KeyVault Secrets, based on the
// deprecated autorest and go-azure-helpers backend.
func keyVaultSecret_autorest(keyVaultDNSSuffix string, kvClient *keyvault.BaseClient) *CustomResource {
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

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)

			// Check if a soft-deleted secret exists and recover it if so
			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			deletedSecret, err := kvClient.GetDeletedSecret(ctx, vaultUrl, secretName.StringValue())
			if err == nil && deletedSecret.RecoveryID != nil {
				logging.V(5).Infof("Found soft-deleted secret %s, recovering it before creating", secretName.StringValue())
				_, err = kvClient.RecoverDeletedSecret(ctx, vaultUrl, secretName.StringValue())
				if err != nil {
					return nil, errors.Wrapf(err, "failed to recover soft-deleted secret %s", secretName.StringValue())
				}
				logging.V(5).Infof("Successfully recovered soft-deleted secret %s", secretName.StringValue())
			} else if err != nil {
				// Check if the error is a 404 (secret not in deleted state), which is fine
				if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
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

			vaultUrl := fmt.Sprintf("https://%s.%s", vaultName.StringValue(), keyVaultDNSSuffix)

			// Check if a soft-deleted key exists and recover it if so
			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			deletedKey, err := kvClient.GetDeletedKey(ctx, vaultUrl, keyName.StringValue())
			if err == nil && deletedKey.RecoveryID != nil {
				logging.V(5).Infof("Found soft-deleted key %s, recovering it before creating", keyName.StringValue())
				_, err = kvClient.RecoverDeletedKey(ctx, vaultUrl, keyName.StringValue())
				if err != nil {
					return nil, errors.Wrapf(err, "failed to recover soft-deleted key %s", keyName.StringValue())
				}
				logging.V(5).Infof("Successfully recovered soft-deleted key %s", keyName.StringValue())
			} else if err != nil {
				// Check if the error is a 404 (key not in deleted state), which is fine
				if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
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
