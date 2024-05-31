package crud

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathParamsErrorHandling(t *testing.T) {
	t.Run("No params, no error", func(t *testing.T) {
		_, _, err := PrepareAzureRESTIdAndQuery("/path", []resources.AzureAPIParameter{}, nil, nil)
		assert.NoError(t, err)
	})

	t.Run("String params, no error", func(t *testing.T) {
		_, _, err := PrepareAzureRESTIdAndQuery("/path",
			[]resources.AzureAPIParameter{
				{
					Name:     "p1",
					Location: "path",
					Value:    &resources.AzureAPIProperty{Type: "string"},
				},
			}, map[string]any{
				"p1": "yay",
			}, nil)
		assert.NoError(t, err)
	})

	t.Run("Non-string params, error", func(t *testing.T) {
		_, _, err := PrepareAzureRESTIdAndQuery("/path",
			[]resources.AzureAPIParameter{
				{
					Name:     "p1",
					Location: "path",
					Value:    &resources.AzureAPIProperty{Type: "string"}, // correct, but value is not
				},
			}, map[string]any{
				"p1": 42,
			}, nil)
		if assert.Error(t, err) {
			assert.Equal(t, "expected string value for path parameter 'p1', got int", err.Error())
		}
	})

	t.Run("Nested path param lookup from props", func(t *testing.T) {
		id, _, err := PrepareAzureRESTIdAndQuery("/path/{container.p1}",
			[]resources.AzureAPIParameter{
				{
					Name:     "container.p1",
					Location: "path",
					Value:    &resources.AzureAPIProperty{Type: "string"}, // correct, but value is not
				},
			}, map[string]any{
				"container": map[string]any{
					"p1": "innerVal",
				},
			}, nil)
		require.NoError(t, err)
		assert.Equal(t, "/path/innerVal", id)
	})
}
