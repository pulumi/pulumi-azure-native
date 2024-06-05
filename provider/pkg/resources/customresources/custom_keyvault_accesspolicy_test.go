package customresources

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePathParams(t *testing.T) {
	t.Run("Valid legacy path", func(t *testing.T) {
		const accessPolicyAzureId = "/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault/vaults/pulumi-testing/accessPolicy"
		params, err := parseKeyVaultPathParams(accessPolicyAzureId)
		require.NoError(t, err)
		assert.Equal(t, "pulumi-testing", params.VaultName)
		assert.Equal(t, "pulumi-dev-shared", params.ResourceGroup)
		assert.Nil(t, params.ObjectId)
	})

	t.Run("Valid path", func(t *testing.T) {
		const accessPolicyAzureId = "/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault/vaults/pulumi-testing/accessPolicy/a5907b7f-6627-4a74-b831-b40bc5ceefdc"
		params, err := parseKeyVaultPathParams(accessPolicyAzureId)
		require.NoError(t, err)
		assert.Equal(t, "pulumi-testing", params.VaultName)
		assert.Equal(t, "pulumi-dev-shared", params.ResourceGroup)
		assert.Equal(t, "a5907b7f-6627-4a74-b831-b40bc5ceefdc", *params.ObjectId)
	})

	t.Run("Invalid paths", func(t *testing.T) {
		for _, invalidId := range []string{
			"",
			"/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared",
			"/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault",
			"/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault/vaults/pulumi-testing",
			"/subscriptions//resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault/vaults/pulumi-testing/accessPolicy",
			"/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups//providers/Microsoft.KeyVault/vaults/pulumi-testing/accessPolicy",
			"/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/pulumi-dev-shared/providers/Microsoft.KeyVault/vaults//accessPolicy",
		} {
			_, err := parseKeyVaultPathParams(invalidId)
			assert.Error(t, err, invalidId)
		}
	})
}

func TestAzureIdFromProperties(t *testing.T) {
	for _, incomplete := range []resource.PropertyMap{
		{
			resource.PropertyKey(resourceGroupName): resource.NewStringProperty("rg"),
		},
		{
			resource.PropertyKey(vaultName): resource.NewStringProperty("vault"),
		},
	} {
		_, err := azureIdFromProperties(incomplete)
		assert.Error(t, err)
	}

	id, err := azureIdFromProperties(resource.PropertyMap{
		resource.PropertyKey(resourceGroupName): resource.NewStringProperty("rg"),
		resource.PropertyKey(vaultName):         resource.NewStringProperty("vault"),
	})
	require.NoError(t, err)
	assert.Equal(t, "/subscriptions/{subscription}/resourceGroups/rg/providers/Microsoft.KeyVault/vaults/vault/accessPolicy", id)
}

func TestSdkParamsFromProperties(t *testing.T) {
	for _, incomplete := range []resource.PropertyMap{
		{},
		{
			policy: resource.NewObjectProperty(resource.PropertyMap{
				"objectId": resource.NewStringProperty("objectId"),
			}),
		},
		{
			policy: resource.NewObjectProperty(resource.PropertyMap{
				"tenantId": resource.NewStringProperty("tenantId"),
			}),
		},
	} {
		_, err := sdkPolicyParamsFromProperties(incomplete)
		assert.Error(t, err)
	}

	complete, err := sdkPolicyParamsFromProperties(resource.PropertyMap{
		policy: resource.NewObjectProperty(resource.PropertyMap{
			"objectId":      resource.NewStringProperty("objectId"),
			"tenantId":      resource.NewStringProperty("tenantId"),
			"applicationId": resource.NewStringProperty("applicationId"),
		}),
	})
	require.NoError(t, err)
	assert.Len(t, complete.Properties.AccessPolicies, 1)
	assert.Equal(t, "objectId", *complete.Properties.AccessPolicies[0].ObjectID)
	assert.Equal(t, "tenantId", *complete.Properties.AccessPolicies[0].TenantID)
	assert.Equal(t, "applicationId", *complete.Properties.AccessPolicies[0].ApplicationID)
}

func TestPropertyPermissionsToSdk(t *testing.T) {
	permissions := resource.PropertyMap{
		"keys": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("get"),
		}),
		"certificates": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("get"),
			resource.NewStringProperty("list"),
		}),
		"secrets": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("purge"),
		}),
		"storage": resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewStringProperty("deletesas"),
		}),
	}

	sdkPermissions := propertyPermissionsToSdk(permissions)
	assert.Len(t, sdkPermissions.Keys, 1)
	assert.Equal(t, armkeyvault.KeyPermissionsGet, *sdkPermissions.Keys[0])
	assert.Len(t, sdkPermissions.Certificates, 2)
	assert.Equal(t, armkeyvault.CertificatePermissionsGet, *sdkPermissions.Certificates[0])
	assert.Equal(t, armkeyvault.CertificatePermissionsList, *sdkPermissions.Certificates[1])
	assert.Len(t, sdkPermissions.Secrets, 1)
	assert.Equal(t, armkeyvault.SecretPermissionsPurge, *sdkPermissions.Secrets[0])
	assert.Len(t, sdkPermissions.Storage, 1)
	assert.Equal(t, armkeyvault.StoragePermissionsDeletesas, *sdkPermissions.Storage[0])
}
