// Copyright 2021, Pulumi Corporation.  All rights reserved.

package resources

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/mgmt/keyvault"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func keyVaultAccessPolicy(client keyvault.VaultsClient) *CustomResource {
	return &CustomResource{
		tok:  "azure-native:keyvault:AccessPolicy",
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/objectId/{objectId}/applicationId/{applicationId}",
		Schema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Key Vault Access Policy for managing policies on existing vaults.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"resourceGroupName": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "Name of the resource group that contains the vault.",
					},
					"keyVaultName": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "Name of the Key Vault.",
					},
					"tenantId": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
					},
					"objectId": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault.",
					},
					"applicationId": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault.",
					},
					"certificatePermissions": {
						TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
						Description: "Permissions the identity has for certificates.",
					},
					"keyPermissions": {
						TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
						Description: "Permissions the identity has for keys.",
					},
					"secretPermissions": {
						TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
						Description: "Permissions the identity has for secrets.",
					},
					"storagePermissions": {
						TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
						Description: "Permissions the identity has for storage accounts.",
					},
				},
			},
			InputProperties: map[string]schema.PropertySpec{
				"resourceGroupName": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "Name of the resource group that contains the vault.",
				},
				"keyVaultName": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "Name of the Key Vault.",
				},
				"tenantId": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
				},
				"objectId": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault.",
				},
				"applicationId": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault.",
				},
				"certificatePermissions": {
					TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
					Description: "Permissions the identity has for certificates.",
				},
				"keyPermissions": {
					TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
					Description: "Permissions the identity has for keys.",
				},
				"secretPermissions": {
					TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
					Description: "Permissions the identity has for secrets.",
				},
				"storagePermissions": {
					TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
					Description: "Permissions the identity has for storage accounts.",
				},
			},
		},
		Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
			parsedId, err := parseKeyVaultAccessPolicyId(id)
			if err != nil {
				return nil, false, err
			}
			vaultResult, err := client.Get(ctx, parsedId.ResourceGroup, parsedId.VaultName)
			if err != nil {
				return nil, false, err
			}
			for _, ap := range *vaultResult.Properties.AccessPolicies {
				if *ap.ObjectID == parsedId.ObjectId && ap.ApplicationID.String() == parsedId.ApplicationId {
					certPermissions := []string{}
					for _, p := range *ap.Permissions.Certificates {
						certPermissions = append(certPermissions, string(p))
					}
					keyPermissions := []string{}
					for _, p := range *ap.Permissions.Keys {
						keyPermissions = append(keyPermissions, string(p))
					}
					secretPermissions := []string{}
					for _, p := range *ap.Permissions.Secrets {
						secretPermissions = append(secretPermissions, string(p))
					}
					storagePermissions := []string{}
					for _, p := range *ap.Permissions.Storage {
						storagePermissions = append(storagePermissions, string(p))
					}
					return map[string]interface{}{
						"resourceGroupName":      parsedId.ResourceGroup,
						"keyVaultName":           *vaultResult.Name,
						"tenantId":               ap.TenantID.String(),
						"objectId":               *ap.ObjectID,
						"applicationId":          ap.ApplicationID.String(),
						"certificatePermissions": certPermissions,
						"keyPermissions":         keyPermissions,
						"secretPermissions":      secretPermissions,
						"storagePermissions":     storagePermissions,
					}, true, nil
				}
			}

			return properties.Mappable(), true, nil
		},
	}
}

type KeyVaultId struct {
	SubscriptionId string
	ResourceGroup  string
	VaultName      string
	ObjectId       string
	ApplicationId  string
}

func parseKeyVaultAccessPolicyId(id string) (KeyVaultId, error) {
	idMatcher := regexp.MustCompile(`(?i)^/subscriptions/(.*)/resourceGroups/(.*)/providers/Microsoft.KeyVault/vaults/(.*)/objectId/(.*)/applicationId/(.*)$`)
	matches := idMatcher.FindStringSubmatch(id)
	if len(matches) != 6 {
		return KeyVaultId{}, fmt.Errorf("unable to parse key vault access policy id in the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/objectId/{objectId}/applicationId/{applicationId}: %s", id)
	}
	return KeyVaultId{
		SubscriptionId: matches[1],
		ResourceGroup:  matches[2],
		VaultName:      matches[3],
		ObjectId:       matches[4],
		ApplicationId:  matches[5],
	}, nil
}
