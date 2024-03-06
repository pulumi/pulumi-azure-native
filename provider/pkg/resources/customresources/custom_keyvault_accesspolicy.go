// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Frequently used resource property names.
const (
	vaultName = "vaultName"
	policy    = "policy"
)

// This used to be "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicy"
// This was updated to include the objectId in the path so it accurately represents the logical resource so we can perform imports by ID.
// However, it's not possible to update an existing id, so we have to ensure we can still work with the old id format - where we fetch the objectId from the properties.
// The objectId is prefixed with `policy.` because it's filled in from the `objectId` within the `policy` property.
const path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicy/{policy.objectId}"

var keyVaultAccessPolicyProperties = map[string]schema.PropertySpec{
	resourceGroupName: {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the resource group that contains the vault.",
	},
	vaultName: {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the Key Vault.",
	},
	policy: {
		// This type is generated from the Azure spec because the Vault resource references it.
		TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:keyvault:AccessPolicyEntry"},
		Description: "The definition of the access policy.",
	},
}

func keyVaultAccessPolicy(client *armkeyvault.VaultsClient) *CustomResource {
	c := &accessPolicyClient{client: client}
	return &CustomResource{
		tok:  "azure-native:keyvault:AccessPolicy",
		path: path,
		LegacySchema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Key Vault Access Policy for managing policies on existing vaults.",
				Type:        "object",
				Properties:  keyVaultAccessPolicyProperties,
			},
			InputProperties: keyVaultAccessPolicyProperties,
			RequiredInputs:  []string{resourceGroupName, vaultName, policy},
		},
		Meta: &AzureAPIResource{
			Path: path,
			PutParameters: []AzureAPIParameter{
				{Name: subscriptionId, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: vaultName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: "policy.objectId", Location: "path", Value: &AzureAPIProperty{Type: "string"}},
				{
					Name:     "properties",
					Location: "body",
					Body: &AzureAPIType{
						Properties: map[string]AzureAPIProperty{
							policy: {Type: "#/types/azure-native:keyvault:AccessPolicyEntry"},
						},
						RequiredProperties: []string{resourceGroupName, vaultName, policy},
					},
				},
			},
		},
		Read: c.read,
		Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
			return c.write(ctx, properties, false /* shouldExist */)
		},
		Update: func(ctx context.Context, id string, properties, olds resource.PropertyMap) (map[string]interface{}, error) {
			return c.write(ctx, properties, true /* shouldExist */)
		},
		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			return c.modify(ctx, properties, armkeyvault.AccessPolicyUpdateKindRemove)
		},
	}
}

type accessPolicyClient struct {
	client *armkeyvault.VaultsClient
}

func (c *accessPolicyClient) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	// input from path
	parsedId, err := parseKeyVaultPathParams(id)
	if err != nil {
		return nil, false, err
	}

	// The old id format doesn't have the objectId in the path, so we fetch it from the properties instead.
	var objectId string
	if parsedId.ObjectId != nil {
		objectId = *parsedId.ObjectId
	} else {
		// input from body
		policyObj := properties[policy].ObjectValue()
		objectId = policyObj["objectId"].StringValue()
	}

	vaultResult, err := c.client.Get(ctx, parsedId.ResourceGroup, parsedId.VaultName, &armkeyvault.VaultsClientGetOptions{})
	if err != nil {
		var respErr *azcore.ResponseError
		if errors.As(err, &respErr) && respErr.StatusCode == 404 {
			return nil, false, nil
		}
		return nil, false, err
	}

	for _, ap := range vaultResult.Properties.AccessPolicies {
		if *ap.ObjectID == objectId {
			ape := map[string]interface{}{
				"tenantId":      ap.TenantID,
				"objectId":      ap.ObjectID,
				"applicationId": ap.ApplicationID,
				"permissions":   sdkPermissionsToMap(ap.Permissions),
			}

			return map[string]interface{}{
				resourceGroupName: parsedId.ResourceGroup,
				vaultName:         vaultResult.Name,
				policy:            ape,
			}, true, nil
		}
	}

	return properties.Mappable(), false, nil
}

func azureIdFromProperties(properties resource.PropertyMap) (string, error) {
	if !properties.HasValue(resourceGroupName) || !properties.HasValue(vaultName) {
		return "", fmt.Errorf("missing required property %s or %s", resourceGroupName, vaultName)
	}
	rg := properties[resourceGroupName].StringValue()
	vaultName := properties[vaultName].StringValue()

	return fmt.Sprintf("/subscriptions/{subscription}/resourceGroups/%s/providers/Microsoft.KeyVault/vaults/%s/accessPolicy",
		rg, vaultName), nil
}

func (c *accessPolicyClient) readFromProperties(ctx context.Context, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	id, err := azureIdFromProperties(properties)
	if err != nil {
		return nil, false, err
	}
	return c.read(ctx, id, properties)
}

func (c *accessPolicyClient) write(ctx context.Context, properties resource.PropertyMap, shouldExist bool) (map[string]interface{}, error) {
	_, found, err := c.readFromProperties(ctx, properties)
	if err != nil {
		return nil, err
	}

	if (found && !shouldExist) || (!found && shouldExist) {
		policyObj := properties[policy].ObjectValue()
		objectId := policyObj["objectId"].StringValue()
		msg := "access policy for %s already exists"
		if shouldExist {
			msg = "access policy for %s does not exist"
		}
		return nil, fmt.Errorf(msg, objectId)
	}

	err = c.modify(ctx, properties, armkeyvault.AccessPolicyUpdateKindReplace)
	if err != nil {
		return nil, err
	}

	// Read it back.
	state, found, err := c.readFromProperties(ctx, properties)
	if !found {
		return nil, errors.New("newly written access policy not found")
	}
	return state, err
}

func sdkPolicyParamsFromProperties(properties resource.PropertyMap) (*armkeyvault.VaultAccessPolicyParameters, error) {
	if !properties.HasValue(policy) {
		return nil, fmt.Errorf("missing required property %s", policy)
	}
	policyObj := properties[policy].ObjectValue()
	if !policyObj.HasValue("objectId") || !policyObj.HasValue("tenantId") {
		return nil, fmt.Errorf("missing required property objectId or tenantId")
	}
	objectId := policyObj["objectId"].StringValue()
	tenantId := policyObj["tenantId"].StringValue()

	var applicationId string
	if policyObj.HasValue("applicationId") {
		applicationId = policyObj["applicationId"].StringValue()
	}

	var permissions armkeyvault.Permissions
	if permissionMap, ok := policyObj["permissions"]; ok {
		permissionVals := permissionMap.ObjectValue()
		permissions = propertyPermissionsToSdk(permissionVals)
	}

	return &armkeyvault.VaultAccessPolicyParameters{
		Properties: &armkeyvault.VaultAccessPolicyProperties{
			AccessPolicies: []*armkeyvault.AccessPolicyEntry{
				{
					ObjectID:      &objectId,
					Permissions:   &permissions,
					TenantID:      &tenantId,
					ApplicationID: &applicationId,
				},
			},
		},
	}, nil
}

// modify creates, updates, or deletes depending on the op parameter.
func (c *accessPolicyClient) modify(ctx context.Context, properties resource.PropertyMap, op armkeyvault.AccessPolicyUpdateKind) error {
	rg := properties[resourceGroupName].StringValue()
	vaultName := properties[vaultName].StringValue()

	params, err := sdkPolicyParamsFromProperties(properties)
	if err != nil {
		return err
	}

	_, err = c.client.UpdateAccessPolicy(ctx, rg, vaultName,
		op,
		*params,
		&armkeyvault.VaultsClientUpdateAccessPolicyOptions{})
	return err
}

type vaultPathParams struct {
	ResourceGroup string
	VaultName     string
	ObjectId      *string
}

func parseKeyVaultPathParams(id string) (vaultPathParams, error) {
	idMatcher := regexp.MustCompile(`(?i)^/subscriptions/.+?/resourceGroups/(.+?)/providers/Microsoft.KeyVault/vaults/(.+?)/accessPolicy/?(.*)$`)
	matches := idMatcher.FindStringSubmatch(id)
	if len(matches) < 3 {
		return vaultPathParams{}, fmt.Errorf("unable to parse key vault access policy id in the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicy/{objectId}: %s", id)
	}
	// Old ids don't have the objectId in the path, so it's optional.
	// If it's not present, we'll fetch it from the properties.
	var objectId *string
	if len(matches) == 4 && matches[3] != "" {
		objectId = &matches[3]
	}
	return vaultPathParams{
		ResourceGroup: matches[1],
		VaultName:     matches[2],
		ObjectId:      objectId,
	}, nil
}

func sdkPolicyToMap(ap *armkeyvault.AccessPolicyEntry) map[string]interface{} {
	return map[string]interface{}{
		"tenantId":      ap.TenantID,
		"objectId":      ap.ObjectID,
		"applicationId": ap.ApplicationID,
		"permissions": map[string]interface{}{
			"certificatePermissions": ap.Permissions.Certificates,
			"keyPermissions":         ap.Permissions.Keys,
			"secretPermissions":      ap.Permissions.Secrets,
			"storagePermissions":     ap.Permissions.Storage,
		},
	}
}

func sdkPermissionsToMap(permissions *armkeyvault.Permissions) map[string]interface{} {
	result := map[string]interface{}{}

	if len(permissions.Certificates) > 0 {
		certPermissions := make([]string, 0, len(permissions.Certificates))
		for _, p := range permissions.Certificates {
			certPermissions = append(certPermissions, string(*p))
		}
		result["certificates"] = certPermissions
	}
	if len(permissions.Keys) > 0 {
		keyPermissions := make([]string, 0, len(permissions.Keys))
		for _, p := range permissions.Keys {
			keyPermissions = append(keyPermissions, string(*p))
		}
		result["keys"] = keyPermissions
	}
	if len(permissions.Secrets) > 0 {
		secretPermissions := make([]string, 0, len(permissions.Secrets))
		for _, p := range permissions.Secrets {
			secretPermissions = append(secretPermissions, string(*p))
		}
		result["secrets"] = secretPermissions
	}
	if len(permissions.Storage) > 0 {
		storagePermissions := make([]string, 0, len(permissions.Storage))
		for _, p := range permissions.Storage {
			storagePermissions = append(storagePermissions, string(*p))
		}
		result["storage"] = storagePermissions
	}
	return result
}

func propertyPermissionsToSdk(permissions resource.PropertyMap) armkeyvault.Permissions {
	result := armkeyvault.Permissions{}

	keyPermissions, ok := permissions["keys"]
	if ok {
		result.Keys = []*armkeyvault.KeyPermissions{}
		for _, v := range keyPermissions.ArrayValue() {
			perm := armkeyvault.KeyPermissions(v.StringValue())
			result.Keys = append(result.Keys, &perm)
		}
	}
	certPermissions, ok := permissions["certificates"]
	if ok {
		result.Certificates = []*armkeyvault.CertificatePermissions{}
		for _, v := range certPermissions.ArrayValue() {
			perm := armkeyvault.CertificatePermissions(v.StringValue())
			result.Certificates = append(result.Certificates, &perm)
		}
	}
	secretPermissions, ok := permissions["secrets"]
	if ok {
		result.Secrets = []*armkeyvault.SecretPermissions{}
		for _, v := range secretPermissions.ArrayValue() {
			perm := armkeyvault.SecretPermissions(v.StringValue())
			result.Secrets = append(result.Secrets, &perm)
		}
	}
	storagePermissions, ok := permissions["storage"]
	if ok {
		result.Storage = []*armkeyvault.StoragePermissions{}
		for _, v := range storagePermissions.ArrayValue() {
			perm := armkeyvault.StoragePermissions(v.StringValue())
			result.Storage = append(result.Storage, &perm)
		}
	}

	return result
}
