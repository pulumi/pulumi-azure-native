// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const (
	roleAssignmentTok  = "azure-native:authorization:RoleAssignment"
	roleAssignmentPath = "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
)

var subscriptionIdRegex = regexp.MustCompile(`/subscriptions/([^/]+)`)

// roleAssignmentClient defines the interface for role assignment operations needed for testing.
type roleAssignmentClient interface {
	// getTenantId retrieves the tenant ID for a given subscription.
	getTenantId(ctx context.Context, subscriptionId string) (string, error)
	// convertResponseToOutputs converts an Azure API response to SDK outputs.
	convertResponseToOutputs(response map[string]any) map[string]any
}

// roleAssignmentClientImpl implements roleAssignmentClient using real Azure SDK clients.
type roleAssignmentClientImpl struct {
	subsClient *armsubscriptions.Client
	crudClient crud.ResourceCrudClient
}

func (c *roleAssignmentClientImpl) getTenantId(ctx context.Context, subscriptionId string) (string, error) {
	if c.subsClient == nil {
		return "", fmt.Errorf("subscriptions client is nil")
	}

	resp, err := c.subsClient.Get(ctx, subscriptionId, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get subscription %s: %w", subscriptionId, err)
	}

	if resp.Subscription.TenantID == nil {
		return "", fmt.Errorf("subscription %s has no tenant ID", subscriptionId)
	}

	return *resp.Subscription.TenantID, nil
}

func (c *roleAssignmentClientImpl) convertResponseToOutputs(response map[string]any) map[string]any {
	return c.crudClient.ResponseBodyToSdkOutputs(response)
}

// roleAssignment returns a custom resource for RoleAssignment that handles cross-tenant scenarios.
//
// Cross-tenant role assignments (those with delegatedManagedIdentityResourceId) require a special
// tenantId query parameter during Read operations. This custom resource extracts the tenant ID from
// the subscription that owns the delegated managed identity and includes it when reading the assignment.
func roleAssignment(
	lookupResource resources.ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	azureClient azure.AzureClient,
	tokenCred azcore.TokenCredential,
) (*CustomResource, error) {
	// This func's parameters are all nil when the function is called for the first time, for
	// `customresources.featureLookup`, so we initialize the objects we need conditionally
	var client roleAssignmentClient
	var res resources.AzureAPIResource

	if lookupResource != nil && crudClientFactory != nil && tokenCred != nil {
		var found bool
		var err error
		res, found, err = lookupResource(roleAssignmentTok)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", roleAssignmentTok)
		}
		crudClient := crudClientFactory(&res)

		subsClient, err := armsubscriptions.NewClient(tokenCred, nil)
		if err != nil {
			return nil, err
		}

		client = &roleAssignmentClientImpl{
			subsClient: subsClient,
			crudClient: crudClient,
		}
	}

	return &CustomResource{
		tok:  roleAssignmentTok,
		path: roleAssignmentPath,
		Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
			return readRoleAssignment(ctx, id, inputs, client, &res, azureClient)
		},
	}, nil
}

// readRoleAssignment implements the Read operation for RoleAssignment, handling cross-tenant scenarios.
func readRoleAssignment(
	ctx context.Context,
	id string,
	inputs resource.PropertyMap,
	client roleAssignmentClient,
	res *resources.AzureAPIResource,
	azureClient azure.AzureClient,
) (map[string]any, bool, error) {
	// Check if this is a cross-tenant role assignment by looking for delegatedManagedIdentityResourceId
	delegatedIdentityResourceId := ""
	if prop, ok := inputs["delegatedManagedIdentityResourceId"]; ok && !prop.IsNull() {
		delegatedIdentityResourceId = prop.StringValue()
	}

	// Determine the tenantId to use for the read operation
	var tenantId string
	if delegatedIdentityResourceId != "" {
		logging.V(9).Infof("Cross-tenant role assignment detected with delegatedManagedIdentityResourceId: %s", delegatedIdentityResourceId)

		// Extract subscription ID from the delegated identity resource ID
		subscriptionId, err := extractSubscriptionId(delegatedIdentityResourceId)
		if err != nil {
			return nil, false, fmt.Errorf("failed to extract subscription ID from delegatedManagedIdentityResourceId: %w", err)
		}

		// Get tenant ID for the subscription
		tenantId, err = client.getTenantId(ctx, subscriptionId)
		if err != nil {
			return nil, false, fmt.Errorf("failed to get tenant ID for subscription %s: %w", subscriptionId, err)
		}

		logging.V(9).Infof("Using tenant ID %s for cross-tenant role assignment read", tenantId)
	}

	// Prepare query parameters, adding tenantId if needed
	queryParams := make(map[string]any)
	for k, v := range res.ReadQueryParams {
		queryParams[k] = v
	}
	if tenantId != "" {
		queryParams["tenantId"] = tenantId
	}

	// Construct the read URL
	url := id + res.ReadPath
	apiVersion := res.APIVersion

	// Call Azure API directly with the modified query parameters
	response, err := azureClient.Get(ctx, url, apiVersion, queryParams)
	if err != nil {
		return nil, false, err
	}

	// Convert response to SDK outputs
	outputs := client.convertResponseToOutputs(response)
	return outputs, true, nil
}

// extractSubscriptionId extracts the subscription ID from an Azure resource ID.
// Format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/...
func extractSubscriptionId(resourceId string) (string, error) {
	matches := subscriptionIdRegex.FindStringSubmatch(resourceId)
	if len(matches) < 2 {
		return "", fmt.Errorf("could not extract subscription ID from resource ID: %s", resourceId)
	}
	return matches[1], nil
}
