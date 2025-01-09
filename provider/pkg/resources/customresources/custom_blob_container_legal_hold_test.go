package customresources

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Verbatim from https://learn.microsoft.com/en-us/rest/api/storagerp/blob-containers/get?view=rest-storagerp-2023-01-01&tabs=HTTP
const legalHoldPropertiesForContainer = `
{
	"properties": {
		"legalHold": {
			"hasLegalHold": true,
			"protectedAppendWritesHistory": {
				"allowProtectedAppendWritesAll": true,
				"timestamp": "2022-09-01T01:58:44.5044483Z"
			},
			"tags": [
				{
					"tag": "tag1",
					"timestamp": "2018-03-26T05:06:09.6964643Z",
					"objectIdentifier": "ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b",
					"tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47"
				},
				{
					"tag": "tag2",
					"timestamp": "2018-03-26T05:06:09.6964643Z",
					"objectIdentifier": "ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b",
					"tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47"
				},
				{
					"tag": "tag3",
					"timestamp": "2018-03-26T05:06:09.6964643Z",
					"objectIdentifier": "ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b",
					"tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47"
				}
			]
		}
	}
}`

func TestCreate(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Basic Create", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		inputs := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
		}
		_, err := custom.Create(context.Background(), containerId+"/legalHold", inputs)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		assert.Equal(t, m.PostIds[0], containerId+"/setLegalHold")
		require.Len(t, m.PostBodies, 1)
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag1")
		assert.NotContains(t, m.PostBodies[0], "allowProtectedAppendWritesAll")
	})

	t.Run("Create with allowProtectedAppendWritesAll", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		inputs := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(true),
		}
		_, err := custom.Create(context.Background(), containerId+"/legalHold", inputs)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		assert.Equal(t, m.PostIds[0], containerId+"/setLegalHold")
		require.Len(t, m.PostBodies, 1)
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[0], "allowProtectedAppendWritesAll")
		assert.True(t, m.PostBodies[0]["allowProtectedAppendWritesAll"].(bool))
	})

	t.Run("Create without tags fails", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		inputs := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{}),
		}
		_, err := custom.Create(context.Background(), containerId+"/legalHold", inputs)
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "'tags'")
		}
	})
}

func TestRead(t *testing.T) {
	id := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName/legalHold"

	azureResponse := map[string]any{}
	err := json.Unmarshal([]byte(legalHoldPropertiesForContainer), &azureResponse)
	require.NoError(t, err)
	m := azure.MockAzureClient{
		GetResponse: azureResponse,
	}

	custom := blobContainerLegalHold(&m)

	res, found, err := custom.Read(context.Background(), id, nil)
	require.NoError(t, err)
	require.True(t, found)
	require.NotNil(t, res)

	require.Len(t, res, 2)
	require.Contains(t, res, "tags")
	tags := res["tags"].([]string)
	assert.Equal(t, []string{"tag1", "tag2", "tag3"}, tags)

	require.Contains(t, res, "allowProtectedAppendWritesAll")
	assert.Equal(t, res["allowProtectedAppendWritesAll"], true)
}

func TestUpdate(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Basic Update: add a tag", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
		}
		news := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
		}

		_, err := custom.Update(context.Background(), containerId+"/legalHold", news, olds)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		assert.Equal(t, m.PostIds[0], containerId+"/setLegalHold")
		require.Len(t, m.PostBodies, 1)
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.NotContains(t, m.PostBodies[0]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag2")
	})

	t.Run("Basic Update: remove a tag", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
		}
		news := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
		}

		_, err := custom.Update(context.Background(), containerId+"/legalHold", news, olds)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		require.Len(t, m.PostBodies, 1)
		assert.Equal(t, m.PostIds[0], containerId+"/clearLegalHold")
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.NotContains(t, m.PostBodies[0]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag2")
	})

	t.Run("Full update", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(true),
		}
		news := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag7"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(false),
		}

		_, err := custom.Update(context.Background(), containerId+"/legalHold", news, olds)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 2)
		require.Len(t, m.PostBodies, 2)

		assert.Equal(t, m.PostIds[0], containerId+"/setLegalHold")
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.NotContains(t, m.PostBodies[0]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag7")
		assert.Contains(t, m.PostBodies[0], "allowProtectedAppendWritesAll")
		assert.Equal(t, m.PostBodies[0]["allowProtectedAppendWritesAll"], false)

		assert.Equal(t, m.PostIds[1], containerId+"/clearLegalHold")
		assert.Contains(t, m.PostBodies[1], "tags")
		assert.Contains(t, m.PostBodies[1]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[1]["tags"], "tag2")
	})

	t.Run("No-op update for changed tag order", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(true),
		}
		news := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag2"),
				resource.NewStringProperty("tag1"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(true),
		}

		_, err := custom.Update(context.Background(), containerId+"/legalHold", news, olds)
		require.NoError(t, err)

		assert.Len(t, m.PostIds, 0)
	})
}

func TestDelete(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Delete removes all tags", func(t *testing.T) {
		m := azure.MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
		}

		err := custom.Delete(context.Background(), containerId+"/legalHold", olds)
		require.NoError(t, err)

		require.Len(t, m.PostIds, 1)
		require.Len(t, m.PostBodies, 1)
		assert.Equal(t, m.PostIds[0], containerId+"/clearLegalHold")
		assert.Contains(t, m.PostBodies[0], "tags")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag1")
		assert.Contains(t, m.PostBodies[0]["tags"], "tag2")
	})
}
