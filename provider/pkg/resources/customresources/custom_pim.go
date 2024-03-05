package customresources

import (
	"context"

	armauth "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func pimRoleManagementPolicyDirect(azureClient azure.AzureClient) *CustomResource {
	return &CustomResource{
		path: "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}",

		// PIM Role Management Policies cannot be created. Instead, each resource has a default policy.
		// But we need to allow the user to create the Pulumi resource, so we return true here and then
		// let Create just return the existing policy.
		CanCreate: func(ctx context.Context, id string) error {
			// TODO should we validate the id parts here?
			return nil
		},

		// PIM Role Management Policies cannot be created. Instead, each resource has a default policy.
		// Simply look up the default policy and return it.
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			result, err := azureClient.Get(ctx, id, "2020-10-01" /* TODO from version lock */, nil)
			return result.(map[string]any), err
		},
	}
}

func pimRoleManagementPolicy(policiesClient *armauth.RoleManagementPoliciesClient) *CustomResource {
	client := &pimRoleManagementPolicyClient{
		armClient: *policiesClient,
	}

	return &CustomResource{
		path: "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}",

		// PIM Role Management Policies cannot be created. Instead, each resource has a default policy.
		// Simply look up the default policy and return it.
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]interface{}, error) {
			_, err := client.Get(ctx, id)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
		// Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]interface{}, bool, error) {
		// 	policy, err := client.Get(ctx, id)
		// 	if err != nil {
		// 		return nil, false, err
		// 	}
		// 	return nil, true, nil
		// },
		// Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]interface{}, error) {
		// 	return nil, nil
		// },
		// Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
		// 	return nil
		// },
	}
}

type pimRoleManagementPolicyClient struct {
	armClient armauth.RoleManagementPoliciesClient
}

func (c *pimRoleManagementPolicyClient) Get(ctx context.Context, id string) (*armauth.RoleManagementPolicy, error) {
	scope := "TODO"
	policyName := "TODO"
	resp, err := c.armClient.Get(ctx, scope, policyName, nil)
	if err != nil {
		return nil, err
	}
	return &resp.RoleManagementPolicy, nil
}
