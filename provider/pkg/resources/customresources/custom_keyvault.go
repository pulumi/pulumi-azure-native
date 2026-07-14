// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// Note: These are "hybrid" resources: schema, update, read use the default implementation, while CREATE and DELETE
// are overridden.
//
// DELETE is implemented via the KeyVault "data plane" API, because it isn't available via ARM.
//
// CREATE recovers a soft-deleted secret/key via the data plane (if one exists) before creating/updating it via the
// normal ARM PUT. Overriding Create means we're fully responsible for the ARM call ourselves: the provider does not
// fall back to its default creation logic once a custom Create is set. See #1174, #1211.
//
// Deletion (soft-delete) and the data plane's record of it becoming visible/recoverable are not instantaneous, so
// the ARM PUT can still race a soft-delete that "just happened" (e.g. destroy immediately followed by up). To
// tolerate that, creation is retried for a bounded time: each conflict re-checks for (and recovers) a soft-deleted
// resource before trying the PUT again. If several consecutive conflicts turn up no soft-deleted resource at all,
// the conflict is treated as a genuine, unrelated one (e.g. a real name collision) and surfaced immediately rather
// than retried for the full timeout.
const (
	keyVaultRecoveryMaxWait       = 3 * time.Minute
	keyVaultRecoveryPollInterval  = 5 * time.Second
	keyVaultRecoveryGraceAttempts = 6 // ~30s of tolerance for soft-delete visibility lag before giving up
)

// keyVaultSecret creates a custom resource for Azure KeyVault Secret.
func keyVaultSecret(cloud cloud.Configuration, tokenCred azcore.TokenCredential,
	crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc,
) (*CustomResource, error) {
	crudClient, err := keyVaultResourceCrudClient(crudClientFactory, lookupResource, "azure-native:keyvault:Secret")
	if err != nil {
		return nil, err
	}

	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}",
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

			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			// recoverIfSoftDeleted reports whether a soft-deleted secret was found and recovered. Any error here
			// (including an unexpected error while checking) is treated as "couldn't confirm", not fatal: the
			// caller retries the create rather than aborting, since the check itself is exploratory.
			recoverIfSoftDeleted := func(ctx context.Context) (recovered bool, err error) {
				deletedSecret, err := kvClient.GetDeletedSecret(ctx, secretName.StringValue(), nil)
				if err != nil {
					var respErr *azcore.ResponseError
					if !errors.As(err, &respErr) || respErr.StatusCode != http.StatusNotFound {
						logging.V(5).Infof("Warning: error checking for a soft-deleted secret %s: %v", secretName.StringValue(), err)
					}
					return false, nil
				}
				if deletedSecret.RecoveryID == nil {
					return false, nil
				}
				logging.V(5).Infof("Found soft-deleted secret %s, recovering it before creating", secretName.StringValue())
				if _, err := kvClient.RecoverDeletedSecret(ctx, secretName.StringValue(), nil); err != nil {
					return false, fmt.Errorf("failed to recover soft-deleted secret %s: %w", secretName.StringValue(), err)
				}
				logging.V(5).Infof("Successfully recovered soft-deleted secret %s", secretName.StringValue())
				return true, nil
			}

			return createOrUpdateArmResourceRecoveringSoftDeleted(ctx, crudClient, id, inputs, recoverIfSoftDeleted,
				fmt.Sprintf("creating KeyVault secret %s", secretName.StringValue()))
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
	}, nil
}

// keyVaultKey creates a custom resource for Azure KeyVault Key.
func keyVaultKey(cloud cloud.Configuration, tokenCred azcore.TokenCredential,
	crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc,
) (*CustomResource, error) {
	crudClient, err := keyVaultResourceCrudClient(crudClientFactory, lookupResource, "azure-native:keyvault:Key")
	if err != nil {
		return nil, err
	}

	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}",
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

			// Related issues: https://github.com/pulumi/pulumi-azure-native/issues/1174
			//                 https://github.com/pulumi/pulumi-azure-native/issues/1211
			// recoverIfSoftDeleted reports whether a soft-deleted key was found and recovered. Any error here
			// (including an unexpected error while checking) is treated as "couldn't confirm", not fatal: the
			// caller retries the create rather than aborting, since the check itself is exploratory.
			recoverIfSoftDeleted := func(ctx context.Context) (recovered bool, err error) {
				deletedKey, err := kvClient.GetDeletedKey(ctx, keyName.StringValue(), nil)
				if err != nil {
					var respErr *azcore.ResponseError
					if !errors.As(err, &respErr) || respErr.StatusCode != http.StatusNotFound {
						logging.V(5).Infof("Warning: error checking for a soft-deleted key %s: %v", keyName.StringValue(), err)
					}
					return false, nil
				}
				if deletedKey.RecoveryID == nil {
					return false, nil
				}
				logging.V(5).Infof("Found soft-deleted key %s, recovering it before creating", keyName.StringValue())
				if _, err := kvClient.RecoverDeletedKey(ctx, keyName.StringValue(), nil); err != nil {
					return false, fmt.Errorf("failed to recover soft-deleted key %s: %w", keyName.StringValue(), err)
				}
				logging.V(5).Infof("Successfully recovered soft-deleted key %s", keyName.StringValue())
				return true, nil
			}

			return createOrUpdateArmResourceRecoveringSoftDeleted(ctx, crudClient, id, inputs, recoverIfSoftDeleted,
				fmt.Sprintf("creating KeyVault key %s", keyName.StringValue()))
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
	}, nil
}

// keyVaultResourceCrudClient builds the ARM CRUD client used to actually create/update KeyVault secrets/keys.
// crudClientFactory and lookupResource are nil when custom resources are built solely for schema/feature lookups
// (see the package-level featureLookup var), in which case no client is needed.
func keyVaultResourceCrudClient(crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc, tok string) (crud.ResourceCrudClient, error) {
	if crudClientFactory == nil || lookupResource == nil {
		return nil, nil
	}
	return createCrudClient(crudClientFactory, lookupResource, tok)
}

// createOrUpdateArmResourceRecoveringSoftDeleted issues the ARM PUT for a resource, mirroring the provider's
// default creation flow (custom Create implementations must do this themselves: the provider does not fall back to
// defaultCreate once a custom Create is set). If the PUT fails with a conflict, recoverIfSoftDeleted is given a
// chance to recover a same-named soft-deleted resource via the data plane, and the PUT is retried, for up to
// keyVaultRecoveryMaxWait: soft-delete and its visibility/recovery on the data plane are not instantaneous, so a
// create immediately following a delete can race them.
//
// If keyVaultRecoveryGraceAttempts consecutive conflicts turn up no soft-deleted resource at all, the conflict is
// assumed to be a genuine, unrelated one (e.g. a real name collision, throttling, or a lock) rather than a
// soft-delete race, and is returned immediately instead of being retried for the full keyVaultRecoveryMaxWait -
// otherwise every such permanent conflict would silently block for 3 minutes before failing with an uninformative
// timeout that discards the real ARM error.
func createOrUpdateArmResourceRecoveringSoftDeleted(
	ctx context.Context,
	crudClient crud.ResourceCrudClient,
	id string,
	inputs resource.PropertyMap,
	recoverIfSoftDeleted func(ctx context.Context) (recovered bool, err error),
	description string,
) (map[string]interface{}, error) {
	bodyParams, err := crudClient.PrepareAzureRESTBody(id, inputs, nil)
	if err != nil {
		return nil, err
	}
	_, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(inputs)
	if err != nil {
		return nil, err
	}

	var outputs map[string]interface{}
	unrecoverableConflicts := 0
	err = util.RetryOperation(keyVaultRecoveryMaxWait, keyVaultRecoveryPollInterval, description,
		func() (bool, error) {
			response, _, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err == nil {
				outputs = crudClient.ResponseBodyToSdkOutputs(response)
				return true, nil
			}

			var respErr *azure.PulumiAzcoreResponseError
			if !errors.As(err, &respErr) || respErr.StatusCode != http.StatusConflict {
				return false, err
			}
			recovered, recErr := recoverIfSoftDeleted(ctx)
			if recErr != nil {
				return false, recErr
			}
			if recovered {
				unrecoverableConflicts = 0
				return false, nil
			}
			unrecoverableConflicts++
			if unrecoverableConflicts >= keyVaultRecoveryGraceAttempts {
				return false, err
			}
			return false, nil
		})
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func reportDeletionError(err error) error {
	var respErr *azcore.ResponseError
	if errors.As(err, &respErr) && respErr.StatusCode == http.StatusNotFound {
		return nil
	}
	return err
}
