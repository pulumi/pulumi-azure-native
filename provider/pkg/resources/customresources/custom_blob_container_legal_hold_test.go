package customresources

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

type MockAzureClient struct {
	postIds    []string
	postBodies []map[string]any
}

func (m *MockAzureClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	return nil
}
func (m *MockAzureClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	return nil
}
func (m *MockAzureClient) Get(ctx context.Context, id string, apiVersion string) (map[string]interface{}, error) {
	return nil, nil
}
func (m *MockAzureClient) Head(ctx context.Context, id string, apiVersion string) error {
	return nil
}
func (m *MockAzureClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	return nil, false, nil
}
func (m *MockAzureClient) Post(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (map[string]interface{}, error) {
	m.postIds = append(m.postIds, id)
	m.postBodies = append(m.postBodies, bodyProps)
	return nil, nil
}
func (m *MockAzureClient) Put(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	return nil, false, nil
}
func (m *MockAzureClient) IsNotFound(err error) bool {
	return false
}

func TestCreate(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Basic Create", func(t *testing.T) {
		m := MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		inputs := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
		}
		_, err := custom.Create(context.Background(), containerId+"/legalHold", inputs)
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 1)
		assert.Equal(t, m.postIds[0], containerId+"/setLegalHold")
		assert.Len(t, m.postBodies, 1)
		assert.Contains(t, m.postBodies[0], "tags")
		assert.Contains(t, m.postBodies[0]["tags"], "tag1")
		assert.NotContains(t, m.postBodies[0], "allowProtectedAppendWritesAll")
	})

	t.Run("Create with allowProtectedAppendWritesAll", func(t *testing.T) {
		m := MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		inputs := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
			}),
			resource.PropertyKey("allowProtectedAppendWritesAll"): resource.NewBoolProperty(true),
		}
		_, err := custom.Create(context.Background(), containerId+"/legalHold", inputs)
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 1)
		assert.Equal(t, m.postIds[0], containerId+"/setLegalHold")
		assert.Len(t, m.postBodies, 1)
		assert.Contains(t, m.postBodies[0], "tags")
		assert.Contains(t, m.postBodies[0]["tags"], "tag1")
		assert.Contains(t, m.postBodies[0], "allowProtectedAppendWritesAll")
		assert.True(t, m.postBodies[0]["allowProtectedAppendWritesAll"].(bool))
	})

	t.Run("Create without tags fails", func(t *testing.T) {
		m := MockAzureClient{}
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

func TestUpdate(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Basic Update: add a tag", func(t *testing.T) {
		m := MockAzureClient{}
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
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 1)
		assert.Equal(t, m.postIds[0], containerId+"/setLegalHold")
		assert.Len(t, m.postBodies, 1)
		assert.Contains(t, m.postBodies[0], "tags")
		assert.NotContains(t, m.postBodies[0]["tags"], "tag1")
		assert.Contains(t, m.postBodies[0]["tags"], "tag2")
	})

	t.Run("Basic Update: remove a tag", func(t *testing.T) {
		m := MockAzureClient{}
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
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 1)
		assert.Len(t, m.postBodies, 1)
		assert.Equal(t, m.postIds[0], containerId+"/clearLegalHold")
		assert.Contains(t, m.postBodies[0], "tags")
		assert.NotContains(t, m.postBodies[0]["tags"], "tag1")
		assert.Contains(t, m.postBodies[0]["tags"], "tag2")
	})

	t.Run("Full update", func(t *testing.T) {
		m := MockAzureClient{}
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
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 2)
		assert.Len(t, m.postBodies, 2)

		assert.Equal(t, m.postIds[0], containerId+"/setLegalHold")
		assert.Contains(t, m.postBodies[0], "tags")
		assert.NotContains(t, m.postBodies[0]["tags"], "tag1")
		assert.Contains(t, m.postBodies[0]["tags"], "tag7")
		assert.Contains(t, m.postBodies[0], "allowProtectedAppendWritesAll")
		assert.Equal(t, m.postBodies[0]["allowProtectedAppendWritesAll"], false)

		assert.Equal(t, m.postIds[1], containerId+"/clearLegalHold")
		assert.Contains(t, m.postBodies[1], "tags")
		assert.Contains(t, m.postBodies[1]["tags"], "tag1")
		assert.Contains(t, m.postBodies[1]["tags"], "tag2")
	})

	t.Run("No-op update for changed tag order", func(t *testing.T) {
		m := MockAzureClient{}
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
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 0)
	})
}

func TestDelete(t *testing.T) {
	containerId := "/subscriptions/123-456/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/accountName/blobServices/default/containers/containerName"

	t.Run("Delete removes all tags", func(t *testing.T) {
		m := MockAzureClient{}
		custom := blobContainerLegalHold(&m)

		olds := resource.PropertyMap{
			resource.PropertyKey("tags"): resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("tag1"),
				resource.NewStringProperty("tag2"),
			}),
		}

		err := custom.Delete(context.Background(), containerId+"/legalHold", olds)
		assert.NoError(t, err)

		assert.Len(t, m.postIds, 1)
		assert.Len(t, m.postBodies, 1)
		assert.Equal(t, m.postIds[0], containerId+"/clearLegalHold")
		assert.Contains(t, m.postBodies[0], "tags")
		assert.Contains(t, m.postBodies[0]["tags"], "tag1")
		assert.Contains(t, m.postBodies[0]["tags"], "tag2")
	})
}
