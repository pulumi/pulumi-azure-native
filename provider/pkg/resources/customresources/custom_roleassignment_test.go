// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRoleAssignmentClient mocks the roleAssignmentClient interface for testing
type mockRoleAssignmentClient struct {
	tenantID          string
	tenantIDError     error
	capturedSubsID    string
	responseOutputs   map[string]any
}

func (m *mockRoleAssignmentClient) getTenantId(ctx context.Context, subscriptionId string) (string, error) {
	m.capturedSubsID = subscriptionId
	if m.tenantIDError != nil {
		return "", m.tenantIDError
	}
	return m.tenantID, nil
}

func (m *mockRoleAssignmentClient) convertResponseToOutputs(response map[string]any) map[string]any {
	if m.responseOutputs != nil {
		return m.responseOutputs
	}
	return response
}

func TestExtractSubscriptionId(t *testing.T) {
	t.Parallel()

	t.Run("valid resource ID", func(t *testing.T) {
		t.Parallel()
		resourceId := "/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"
		subscriptionId, err := extractSubscriptionId(resourceId)
		require.NoError(t, err)
		assert.Equal(t, "12345678-1234-1234-1234-123456789abc", subscriptionId)
	})

	t.Run("resource ID at subscription level", func(t *testing.T) {
		t.Parallel()
		resourceId := "/subscriptions/abcd-efgh-ijkl/providers/Microsoft.Authorization/roleDefinitions/role1"
		subscriptionId, err := extractSubscriptionId(resourceId)
		require.NoError(t, err)
		assert.Equal(t, "abcd-efgh-ijkl", subscriptionId)
	})

	t.Run("invalid resource ID - no subscriptions segment", func(t *testing.T) {
		t.Parallel()
		resourceId := "/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"
		_, err := extractSubscriptionId(resourceId)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not extract subscription ID")
	})

	t.Run("empty resource ID", func(t *testing.T) {
		t.Parallel()
		_, err := extractSubscriptionId("")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not extract subscription ID")
	})
}

func TestReadRoleAssignment_CrossTenant(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"
	delegatedIdentityResourceID := "/subscriptions/delegated-sub-789/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"
	tenantID := "tenant-abc"

	// Mock Azure API response for role assignment
	roleAssignmentResponse := map[string]any{
		"id":   roleAssignmentID,
		"name": "role-456",
		"type": "Microsoft.Authorization/roleAssignments",
		"properties": map[string]any{
			"roleDefinitionId":                    "/subscriptions/sub-123/providers/Microsoft.Authorization/roleDefinitions/role-def-1",
			"principalId":                         "principal-123",
			"principalType":                       "ServicePrincipal",
			"delegatedManagedIdentityResourceId": delegatedIdentityResourceID,
		},
	}

	mockAzureClient := &azure.MockAzureClient{
		GetResponse: roleAssignmentResponse,
	}

	mockClient := &mockRoleAssignmentClient{
		tenantID: tenantID,
		responseOutputs: map[string]any{
			"id":                                  roleAssignmentID,
			"name":                                "role-456",
			"roleDefinitionId":                    "/subscriptions/sub-123/providers/Microsoft.Authorization/roleDefinitions/role-def-1",
			"principalId":                         "principal-123",
			"delegatedManagedIdentityResourceId": delegatedIdentityResourceID,
		},
	}

	res := &resources.AzureAPIResource{
		APIVersion:      "2022-04-01",
		ReadPath:        "",
		ReadQueryParams: map[string]any{},
	}

	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewStringProperty(delegatedIdentityResourceID),
	}

	outputs, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.NoError(t, err)
	assert.True(t, found)
	assert.NotNil(t, outputs)

	// Verify subscription ID was extracted and looked up
	assert.Equal(t, "delegated-sub-789", mockClient.capturedSubsID)

	// Verify the GET request was made to the correct URL
	require.Len(t, mockAzureClient.GetIds, 1)
	assert.Equal(t, roleAssignmentID, mockAzureClient.GetIds[0])

	// Verify API version
	require.Len(t, mockAzureClient.GetApiVersions, 1)
	assert.Equal(t, "2022-04-01", mockAzureClient.GetApiVersions[0])
}

func TestReadRoleAssignment_CrossTenant_WithExistingQueryParams(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"
	delegatedIdentityResourceID := "/subscriptions/delegated-sub-789/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"
	tenantID := "tenant-abc"

	mockAzureClient := &azure.MockAzureClient{
		GetResponse: map[string]any{"id": roleAssignmentID},
	}

	mockClient := &mockRoleAssignmentClient{
		tenantID:        tenantID,
		responseOutputs: map[string]any{"id": roleAssignmentID},
	}

	// Resource with existing query parameters
	res := &resources.AzureAPIResource{
		APIVersion: "2022-04-01",
		ReadPath:   "",
		ReadQueryParams: map[string]any{
			"api-version": "2022-04-01",
		},
	}

	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewStringProperty(delegatedIdentityResourceID),
	}

	_, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.NoError(t, err)
	assert.True(t, found)

	// Note: We can't directly verify query params with the current MockAzureClient implementation,
	// but we verify the call was made successfully
	require.Len(t, mockAzureClient.GetIds, 1)
}

func TestReadRoleAssignment_SameTenant(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"

	roleAssignmentResponse := map[string]any{
		"id":   roleAssignmentID,
		"name": "role-456",
		"type": "Microsoft.Authorization/roleAssignments",
		"properties": map[string]any{
			"roleDefinitionId": "/subscriptions/sub-123/providers/Microsoft.Authorization/roleDefinitions/role-def-1",
			"principalId":      "principal-123",
			"principalType":    "User",
		},
	}

	mockAzureClient := &azure.MockAzureClient{
		GetResponse: roleAssignmentResponse,
	}

	// Client that should NOT be called for tenant ID lookup
	mockClient := &mockRoleAssignmentClient{
		tenantIDError: fmt.Errorf("should not be called"),
		responseOutputs: map[string]any{
			"id":               roleAssignmentID,
			"name":             "role-456",
			"roleDefinitionId": "/subscriptions/sub-123/providers/Microsoft.Authorization/roleDefinitions/role-def-1",
			"principalId":      "principal-123",
		},
	}

	res := &resources.AzureAPIResource{
		APIVersion:      "2022-04-01",
		ReadPath:        "",
		ReadQueryParams: map[string]any{},
	}

	// No delegatedManagedIdentityResourceId in inputs
	inputs := resource.PropertyMap{}

	outputs, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.NoError(t, err)
	assert.True(t, found)
	assert.NotNil(t, outputs)

	// Verify tenant ID lookup was NOT called (capturedSubsID should be empty)
	assert.Empty(t, mockClient.capturedSubsID)

	// Verify the GET request was made
	require.Len(t, mockAzureClient.GetIds, 1)
	assert.Equal(t, roleAssignmentID, mockAzureClient.GetIds[0])
}

func TestReadRoleAssignment_InvalidDelegatedIdentityResourceId(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"
	invalidResourceID := "/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"

	mockAzureClient := &azure.MockAzureClient{}
	mockClient := &mockRoleAssignmentClient{}
	res := &resources.AzureAPIResource{
		APIVersion: "2022-04-01",
	}

	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewStringProperty(invalidResourceID),
	}

	_, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.Error(t, err)
	assert.False(t, found)
	assert.Contains(t, err.Error(), "failed to extract subscription ID")
}

func TestReadRoleAssignment_SubscriptionLookupFails(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"
	delegatedIdentityResourceID := "/subscriptions/delegated-sub-789/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"

	mockAzureClient := &azure.MockAzureClient{}
	mockClient := &mockRoleAssignmentClient{
		tenantIDError: fmt.Errorf("subscription lookup failed"),
	}
	res := &resources.AzureAPIResource{
		APIVersion: "2022-04-01",
	}

	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewStringProperty(delegatedIdentityResourceID),
	}

	_, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.Error(t, err)
	assert.False(t, found)
	assert.Contains(t, err.Error(), "failed to get tenant ID for subscription")
}

func TestReadRoleAssignment_AzureGetFails(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"
	delegatedIdentityResourceID := "/subscriptions/delegated-sub-789/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity"

	mockAzureClient := &azure.MockAzureClient{
		GetResponseErr: fmt.Errorf("role assignment not found"),
	}
	mockClient := &mockRoleAssignmentClient{
		tenantID: "tenant-123",
	}
	res := &resources.AzureAPIResource{
		APIVersion: "2022-04-01",
	}

	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewStringProperty(delegatedIdentityResourceID),
	}

	_, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.Error(t, err)
	assert.False(t, found)
	assert.Contains(t, err.Error(), "role assignment not found")
}

func TestReadRoleAssignment_NullDelegatedIdentityResourceId(t *testing.T) {
	t.Parallel()

	roleAssignmentID := "/subscriptions/sub-123/providers/Microsoft.Authorization/roleAssignments/role-456"

	mockAzureClient := &azure.MockAzureClient{
		GetResponse: map[string]any{"id": roleAssignmentID},
	}
	mockClient := &mockRoleAssignmentClient{
		tenantIDError:   fmt.Errorf("should not be called"),
		responseOutputs: map[string]any{"id": roleAssignmentID},
	}
	res := &resources.AzureAPIResource{
		APIVersion: "2022-04-01",
	}

	// delegatedManagedIdentityResourceId is explicitly null
	inputs := resource.PropertyMap{
		"delegatedManagedIdentityResourceId": resource.NewNullProperty(),
	}

	outputs, found, err := readRoleAssignment(context.Background(), roleAssignmentID, inputs, mockClient, res, mockAzureClient)

	require.NoError(t, err)
	assert.True(t, found)
	assert.NotNil(t, outputs)

	// Verify tenant ID lookup was NOT called
	assert.Empty(t, mockClient.capturedSubsID)
}

// Example response format from Azure API for documentation
const exampleRoleAssignmentResponse = `
{
  "id": "/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747",
  "type": "Microsoft.Authorization/roleAssignments",
  "name": "b0f43c54-e787-4862-89b1-a653fa9cf747",
  "properties": {
    "roleDefinitionId": "/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d",
    "principalId": "ce2ce14e-85d7-4629-bdbc-454d0519d987",
    "principalType": "User",
    "scope": "/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"
  }
}
`

func TestResponseParsing(t *testing.T) {
	t.Parallel()

	// Verify we can parse the example response
	var response map[string]any
	err := json.Unmarshal([]byte(exampleRoleAssignmentResponse), &response)
	require.NoError(t, err)
	assert.Contains(t, response, "properties")

	props := response["properties"].(map[string]any)
	assert.Contains(t, props, "roleDefinitionId")
	assert.Contains(t, props, "principalId")
}
