package customresources

import (
	"context"
	"testing"

	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockResourceLookup ResourceLookupFunc = func(resourceType string) (AzureAPIResource, bool, error) {
	return AzureAPIResource{}, true, nil
}

type WebAppMockAzureClient struct {
	queryParamsOfLastDelete map[string]any
}

func (m *WebAppMockAzureClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	m.queryParamsOfLastDelete = queryParams
	return nil
}
func (m *WebAppMockAzureClient) CanCreate(ctx context.Context, id, path, apiVersion, readMethod string, isSingletonResource, hasDefaultBody bool, isDefaultResponse func(map[string]any) bool) error {
	return nil
}
func (m *WebAppMockAzureClient) Get(ctx context.Context, id string, apiVersion string) (any, error) {
	return nil, nil
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
	custom.Delete(context.Background(), "id", resource.PropertyMap{})
	assert.Len(t, azureClient.queryParamsOfLastDelete, 1)
	assert.Contains(t, azureClient.queryParamsOfLastDelete, "deleteEmptyServerFarm")
}
