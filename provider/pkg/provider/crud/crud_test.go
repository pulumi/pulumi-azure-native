package crud

import (
	"context"
	"net/http"
	"testing"

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
		client, err := azure.CreateTestClient(t, assertions)
		require.NoError(t, err)

		crudClient := NewResourceCrudClient(client, nil, nil, "123", res)
		// Runs the assertions as part of HTTP transport
		crudClient.Read(context.Background(), resourceId, "")
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

	t.Run("explicit GET, additional query params", func(t *testing.T) {
		res := resources.AzureAPIResource{
			ReadMethod:      "GET",
			ReadQueryParams: map[string]any{"$expand": "*"},
		}
		assertions := func(t *testing.T, req *http.Request) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, resourceId, req.URL.Path)
			assert.Equal(t, "*", req.URL.Query().Get("$expand"))
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

func TestSqlVirtualMachineUsesReadQueryParams(t *testing.T) {
	sqlVmResource := resources.AzureAPIResource{
		Path:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}",
		ReadQueryParams: map[string]any{"$expand": "*"},
	}
	sqlVmId := "/subscriptions/123/resourceGroups/rg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/vm"

	runTest := func(t *testing.T, res *resources.AzureAPIResource, assertions func(t *testing.T, req *http.Request)) {
		client, err := azure.CreateTestClient(t, assertions)
		require.NoError(t, err)

		crudClient := NewResourceCrudClient(client, nil, nil, "123", res)
		// Runs the assertions as part of HTTP transport
		crudClient.Read(context.Background(), sqlVmId, "")
	}

	runTest(t, &sqlVmResource, func(t *testing.T, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, sqlVmId, req.URL.Path)
		assert.Equal(t, "*", req.URL.Query().Get("$expand"))
	})

	// Sanity check
	sqlVmResource.ReadQueryParams = nil
	runTest(t, &sqlVmResource, func(t *testing.T, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, sqlVmId, req.URL.Path)
		assert.Empty(t, req.URL.Query().Get("$expand"))
	})
}

func TestNestedFieldNoCopy(t *testing.T) {
	target := map[string]any{"foo": "bar"}

	obj := map[string]any{
		"a": map[string]any{
			"b": target,
			"c": nil,
			"d": []any{"foo"},
			"e": []any{
				map[string]any{
					"f": "bar",
				},
			},
		},
	}

	// case 1: field exists and is non-nil
	res, exists, err := nestedFieldNoCopy(obj, "a", "b")
	assert.True(t, exists)
	assert.NoError(t, err)
	assert.Equal(t, target, res)
	target["foo"] = "baz"
	assert.Equal(t, target["foo"], res.(map[string]any)["foo"], "result should be a reference to the expected item")

	// case 2: field exists and is nil
	res, exists, err = nestedFieldNoCopy(obj, "a", "c")
	assert.True(t, exists)
	assert.NoError(t, err)
	assert.Nil(t, res)

	// case 3: error traversing obj
	res, exists, err = nestedFieldNoCopy(obj, "a", "d", "foo")
	assert.False(t, exists)
	assert.Error(t, err)
	assert.Nil(t, res)

	// case 4: field does not exist
	res, exists, err = nestedFieldNoCopy(obj, "a", "g")
	assert.False(t, exists)
	assert.NoError(t, err)
	assert.Nil(t, res)

	// case 5: intermediate field does not exist
	res, exists, err = nestedFieldNoCopy(obj, "a", "g", "f")
	assert.False(t, exists)
	assert.NoError(t, err)
	assert.Nil(t, res)

	// case 6: intermediate field is null
	//         (background: happens easily in YAML)
	res, exists, err = nestedFieldNoCopy(obj, "a", "c", "f")
	assert.False(t, exists)
	assert.NoError(t, err)
	assert.Nil(t, res)

	// case 7: array/slice syntax is not supported
	//         (background: users may expect this to be supported)
	res, exists, err = nestedFieldNoCopy(obj, "a", "e[0]")
	assert.False(t, exists)
	assert.NoError(t, err)
	assert.Nil(t, res)
}

func TestSetNestedFieldNoCopy(t *testing.T) {
	obj := map[string]any{
		"x": map[string]any{
			"y": 1,
			"a": "foo",
		},
	}

	// setting into a new container
	err := setNestedFieldNoCopy(obj, []any{"bar"}, "z")
	assert.NoError(t, err)
	assert.Len(t, obj, 2)
	assert.Equal(t, "bar", obj["z"].([]interface{})[0])

	// setting into an existing container
	err = setNestedFieldNoCopy(obj, []any{"bar"}, "x", "z")
	assert.NoError(t, err)
	assert.Len(t, obj["x"], 3)
	assert.Len(t, obj["x"].(map[string]interface{})["z"], 1)
	assert.Equal(t, "bar", obj["x"].(map[string]interface{})["z"].([]interface{})[0])

	// error traversing a non-container
	err = setNestedFieldNoCopy(obj, []any{}, "x", "y", "z")
	assert.Error(t, err, `value cannot be set because x.y is not a map[string]interface{}`)

}
