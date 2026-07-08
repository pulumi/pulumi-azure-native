// Copyright 2026, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceControlDelete(t *testing.T) {
	id := "/subscriptions/123-456/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/sourcecontrols/789e0c1f-4a3d-43ad-809c-e713b677b04a"

	t.Run("Delete posts repository access to the /delete sub-path", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := securityInsightsSourceControl(&m)

		previousInputs := resource.PropertyMap{
			resource.PropertyKey("repositoryAccess"): resource.NewObjectProperty(resource.PropertyMap{
				resource.PropertyKey("kind"):     resource.NewStringProperty("OAuth"),
				resource.PropertyKey("code"):     resource.NewStringProperty("939fd7c6caf754f4f41f"),
				resource.PropertyKey("state"):    resource.NewStringProperty("state"),
				resource.PropertyKey("clientId"): resource.NewStringProperty("54b3c2c0-1f48-4a1c-af9f-6399c3240b73"),
			}),
		}

		err := custom.Delete(context.Background(), id, previousInputs, nil)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		assert.Equal(t, id+"/delete", m.PostIds[0])

		require.Len(t, m.PostBodies, 1)
		properties, ok := m.PostBodies[0]["properties"].(map[string]any)
		require.True(t, ok, "expected 'properties' to be a nested object")
		repositoryAccess, ok := properties["repositoryAccess"].(map[string]any)
		require.True(t, ok, "expected 'properties.repositoryAccess' to be a nested object")
		assert.Equal(t, "OAuth", repositoryAccess["kind"])
		assert.Equal(t, "939fd7c6caf754f4f41f", repositoryAccess["code"])
		assert.Equal(t, "state", repositoryAccess["state"])
		assert.Equal(t, "54b3c2c0-1f48-4a1c-af9f-6399c3240b73", repositoryAccess["clientId"])
	})

	t.Run("Delete without repositoryAccess fails", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := securityInsightsSourceControl(&m)

		err := custom.Delete(context.Background(), id, resource.PropertyMap{}, nil)
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "repositoryAccess")
		}
		assert.Len(t, m.PostIds, 0)
	})
}
