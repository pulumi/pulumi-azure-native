package customresources

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockResourceLookup resources.ResourceLookupFunc = func(resourceType string) (resources.AzureAPIResource, bool, error) {
	return resources.AzureAPIResource{}, true, nil
}

type WebAppMockAzureClient struct {
	queryParamsOfLastDelete map[string]any
	getIds                  []string
}

func (m *WebAppMockAzureClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	m.queryParamsOfLastDelete = queryParams
	return nil
}
func (m *WebAppMockAzureClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	return nil
}
func (m *WebAppMockAzureClient) Get(ctx context.Context, id string, apiVersion string, queryParams map[string]any) (any, error) {
	m.getIds = append(m.getIds, id)
	return map[string]any{}, nil
}
func (m *WebAppMockAzureClient) Head(ctx context.Context, id string, apiVersion string) error {
	return nil
}
func (m *WebAppMockAzureClient) Patch(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	return nil, false, nil
}
func (m *WebAppMockAzureClient) Post(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (any, error) {
	return nil, nil
}
func (m *WebAppMockAzureClient) Put(ctx context.Context, id string, bodyProps map[string]interface{}, queryParameters map[string]interface{}, asyncStyle string) (map[string]interface{}, bool, error) {
	return nil, false, nil
}
func (m *WebAppMockAzureClient) IsNotFound(err error) bool {
	return false
}

func TestSetsDeleteParam(t *testing.T) {
	azureClient := &WebAppMockAzureClient{}

	custom, err := webApp(nil, azureClient, mockResourceLookup)
	require.NoError(t, err)
	custom.Delete(context.Background(), "id", resource.PropertyMap{}, resource.PropertyMap{})
	assert.Len(t, azureClient.queryParamsOfLastDelete, 1)
	assert.Contains(t, azureClient.queryParamsOfLastDelete, "deleteEmptyServerFarm")
}

func TestReadsSiteConfig(t *testing.T) {
	appId := "/subscriptions/123/resourceGroups/rg123/providers/Microsoft.Web/sites/app123"
	azureClient := &WebAppMockAzureClient{}

	var f crud.ResourceCrudClientFactory = func(res *resources.AzureAPIResource) crud.ResourceCrudClient {
		return crud.NewResourceCrudClient(azureClient, nil, nil, "123", res)
	}

	custom, err := webApp(f, azureClient, mockResourceLookup)
	require.NoError(t, err)

	// Returns an error because the responses for both GETs are empty, but we only care about the requests.
	_, _, _ = custom.Read(context.Background(), appId, resource.PropertyMap{})
	require.Contains(t, azureClient.getIds, appId)
	require.Contains(t, azureClient.getIds, appId+"/config/web")
}

func TestMergeWebAppSiteConfig(t *testing.T) {
	webApp := map[string]any{
		"id": "/subscriptions/123/resourceGroups/rg123/providers/Microsoft.Web/sites/app123",
		"properties": map[string]any{
			"enabled": true,
			"siteConfig": map[string]any{
				"defaultDocuments":       nil,
				"ipSecurityRestrictions": nil,
			},
		},
	}

	siteConfig := map[string]any{
		"id": "/subscriptions/123/resourceGroups/rg123/providers/Microsoft.Web/sites/app123/config/web",
		"properties": map[string]any{
			"defaultDocuments": []any{
				"pulumi.html",
			},
			"ipSecurityRestrictions": []any{
				map[string]any{
					"action":    "Allow",
					"ipAddress": "198.51.100.0/22",
					"name":      "pulumi",
					"priority":  float64(100),
					"tag":       "Default",
				},
			},
		},
	}

	err := mergeWebAppSiteConfig(webApp, siteConfig)
	require.NoError(t, err)
	assert.Equal(t, webApp["properties"].(map[string]any)["siteConfig"], siteConfig["properties"])
}
