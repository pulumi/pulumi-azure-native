// Copyright 2026, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const (
	sourceControlPath       = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/sourcecontrols/{sourceControlId}"
	sourceControlApiVersion = "2025-09-01"
)

// securityInsightsSourceControl customizes the standard generated SourceControl resource so it can be deleted.
// Azure moved this resource's delete operation from a plain HTTP DELETE on the resource's own path to a POST
// against a `.../delete` sub-path that additionally requires re-submitting the repository access credentials
// (Azure uses them to remove the connected repository's webhook). Generic codegen resource discovery only
// recognizes a DELETE verb on the resource's own path, so it can no longer see this resource at all. Registering
// it here - with no Schema, Create, Read, or Update override - restores standard generated behavior for
// everything except Delete.
func securityInsightsSourceControl(azureClient azure.AzureClient) *CustomResource {
	return &CustomResource{
		tok:  "azure-native:securityinsights:SourceControl",
		path: sourceControlPath,
		Delete: func(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
			repositoryAccess, ok := previousInputs["repositoryAccess"]
			if !ok || !repositoryAccess.IsObject() {
				return fmt.Errorf("deleting a SecurityInsights SourceControl requires the 'repositoryAccess' " +
					"credentials it was created with, since Azure needs them to remove the connected repository's webhook")
			}

			body := map[string]any{
				"properties": map[string]any{
					"repositoryAccess": repositoryAccess.Mappable(),
				},
			}
			queryParams := map[string]any{"api-version": sourceControlApiVersion}

			_, err := azureClient.Post(ctx, id+"/delete", body, queryParams)
			return err
		},
	}
}
