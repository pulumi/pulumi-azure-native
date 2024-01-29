package customresources

import (
	"context"
	"testing"

	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

type MockResourceLookupper struct{}

func (t *MockResourceLookupper) LookupResource(id string) (AzureAPIResource, bool, error) {
	return AzureAPIResource{}, true, nil
}

type MockAzureDeleter struct {
	queryParamsOfLastDelete map[string]any
}

func (t *MockAzureDeleter) AzureDelete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	t.queryParamsOfLastDelete = queryParams
	return nil
}

func TestSetsDeleteParam(t *testing.T) {
	deleter := MockAzureDeleter{}
	custom := customWebAppDelete(&MockResourceLookupper{}, &deleter)
	custom.Delete(context.Background(), "id", resource.PropertyMap{})
	assert.Len(t, deleter.queryParamsOfLastDelete, 1)
	assert.Contains(t, deleter.queryParamsOfLastDelete, "deleteEmptyServerFarm")
}
