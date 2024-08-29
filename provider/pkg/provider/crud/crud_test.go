package crud

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
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

	t.Run("Path param from props", func(t *testing.T) {
		id, _, err := PrepareAzureRESTIdAndQuery("/path/{p1}",
			[]resources.AzureAPIParameter{
				{
					Name:     "p1",
					Location: "path",
					Value:    &resources.AzureAPIProperty{Type: "string"},
				},
			}, map[string]any{
				"p1": "val",
			}, nil)
		require.NoError(t, err)
		assert.Equal(t, "/path/val", id)
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

	t.Run("Deeply nested path param", func(t *testing.T) {
		id, _, err := PrepareAzureRESTIdAndQuery("/path/{container.inner.p1}",
			[]resources.AzureAPIParameter{
				{
					Name:     "container.inner.p1",
					Location: "path",
					Value:    &resources.AzureAPIProperty{Type: "string"}, // correct, but value is not
				},
			}, map[string]any{
				"container": map[string]any{
					"inner": map[string]any{
						"p1": "deepVal",
					},
				},
			}, nil)
		require.NoError(t, err)
		assert.Equal(t, "/path/deepVal", id)
	})
}

func TestCanCreate_RequestUrls(t *testing.T) {
	const resourceId = "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm"

	runTest := func(t *testing.T, res *resources.AzureAPIResource, assertions func(t *testing.T, req *http.Request)) {
		transp := &requestAssertingTransporter{
			t:          t,
			assertions: assertions,
		}
		// TODO,tkappler despite disabled telemetry and retries, the pipeline is still slow at ~9s per sub-test.
		// Several policies are added by default with no way to disable them.
		opts := arm.ClientOptions{
			ClientOptions: policy.ClientOptions{
				Retry:     policy.RetryOptions{MaxRetries: 0}, // speeds up the tests
				Telemetry: policy.TelemetryOptions{Disabled: true},
				Transport: transp,
			},
			DisableRPRegistration: true,
		}
		client := azure.NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, &opts)

		crudClient := NewResourceCrudClient(client, nil, nil, "123", res)
		// Runs the assertions as part of HTTP transport
		crudClient.CanCreate(context.Background(), resourceId)
	}

	t.Run("explicit GET, no read path", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadMethod: "GET",
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, resourceId, req.URL.Path)
		}
		runTest(t, &res, assertions)
	})

	t.Run("explicit GET, read path", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadMethod: "GET",
			ReadPath:   "/read/me",
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, resourceId+"/read/me", req.URL.Path)
		}
		runTest(t, &res, assertions)
	})

	t.Run("implicit GET, no read path", func(t *testing.T) {
		res := resources.AzureAPIResource{}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, resourceId, req.URL.Path)
		}
		runTest(t, &res, assertions)
	})

	t.Run("implicit GET, read path", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadPath: "/read/me",
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, resourceId+"/read/me", req.URL.Path)
		}
		runTest(t, &res, assertions)
	})

	t.Run("POST, no read path", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadMethod: "POST",
			ReadPath:   "/read/me",
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "POST", req.Method)
			assert.Equal(t, resourceId+"/read/me", req.URL.Path)
		}
		runTest(t, &res, assertions)
	})

	t.Run("POST, read path", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadMethod: "POST",
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "POST", req.Method)
			assert.Equal(t, resourceId, req.URL.Path)
		}
		runTest(t, &res, assertions)
	})
}

type requestAssertingTransporter struct {
	t          *testing.T
	assertions func(*testing.T, *http.Request)
}

func (r *requestAssertingTransporter) Do(req *http.Request) (*http.Response, error) {
	r.assertions(r.t, req)
	return nil, nil
}
