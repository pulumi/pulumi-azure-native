package resources

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

type TestAzureClient struct {
	queryParamsOfLastDelete map[string]any
}

func (t *TestAzureClient) LookupResource(id string) (AzureAPIResource, bool, error) {
	return AzureAPIResource{}, true, nil
}

func (t *TestAzureClient) AzureDelete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	t.queryParamsOfLastDelete = queryParams
	return nil
}

func TestSetsDeleteParam(t *testing.T) {
	client := TestAzureClient{}
	custom := customWebAppDelete(&client)
	custom.Delete(context.TODO(), "id", resource.PropertyMap{})
	assert.Len(t, client.queryParamsOfLastDelete, 1)
	assert.Contains(t, client.queryParamsOfLastDelete, "deleteEmptyServerFarm")
}
