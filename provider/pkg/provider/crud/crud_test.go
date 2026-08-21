package crud

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

func TestPathParamEncoding(t *testing.T) {

	t.Run("path encoding", func(t *testing.T) {
		tests := []struct {
			p1       string
			expected string
		}{
			{
				p1:       "my-value",
				expected: "/path/my-value",
			},
			{
				p1:       "my-value with space and /",
				expected: "/path/my-value%20with%20space%20and%20%2F",
			},
		}

		for _, tt := range tests {
			t.Run(tt.p1, func(t *testing.T) {
				id, _, err := PrepareAzureRESTIdAndQuery("/path/{p1}",
					[]resources.AzureAPIParameter{
						{
							Name:            "p1",
							Location:        "path",
							Value:           &resources.AzureAPIProperty{Type: "string"},
							SkipUrlEncoding: false,
						},
					}, map[string]any{
						"p1": tt.p1,
					}, nil)
				require.NoError(t, err)
				assert.Equal(t, tt.expected, id)
			})
		}
	})

	t.Run("skip url encoding", func(t *testing.T) {
		tests := []struct {
			scope    string
			expected string
		}{
			{
				scope:    "subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group",
				expected: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/path",
			},
			{
				scope:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group",
				expected: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/path",
			},
		}

		for _, tt := range tests {
			t.Run(tt.scope, func(t *testing.T) {
				id, _, err := PrepareAzureRESTIdAndQuery("/{scope}/path",
					[]resources.AzureAPIParameter{
						{
							Name:            "scope",
							Location:        "path",
							Value:           &resources.AzureAPIProperty{Type: "string"},
							SkipUrlEncoding: true,
						},
					}, map[string]any{
						"scope": tt.scope,
					}, nil)
				require.NoError(t, err)
				assert.Equal(t, tt.expected, id)
			})
		}
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

// TestCanCreateSkipsNetworkCallForSingleton is a regression test for issue #4738: CanCreate must
// not make a network call for singleton resources, since a singleton always exists once its
// parent does and the network round trip only exposes a window for transient/unexpected ARM
// errors (e.g. a 400 AuthenticationFailed while auth context propagates for a just-recreated
// parent) to wrongly abort creation.
func TestCanCreateSkipsNetworkCallForSingleton(t *testing.T) {
	res := &resources.AzureAPIResource{Singleton: true}

	called := false
	client, err := azure.CreateTestClient(t, func(t *testing.T, req *http.Request) {
		called = true
	})
	require.NoError(t, err)

	crudClient := NewResourceCrudClient(client, nil, nil, "123", res)
	err = crudClient.CanCreate(context.Background(), "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/sa/fileServices/default")

	assert.NoError(t, err)
	assert.False(t, called, "CanCreate should not make a network call for singleton resources")
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

func TestContainsNilValues(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		assert.False(t, containsNilValues(map[string]any{}))
	})
	t.Run("no nils", func(t *testing.T) {
		assert.False(t, containsNilValues(map[string]any{
			"a": "hello",
			"b": map[string]any{"c": 42},
		}))
	})
	t.Run("top-level nil", func(t *testing.T) {
		assert.True(t, containsNilValues(map[string]any{
			"a": "hello",
			"b": nil,
		}))
	})
	t.Run("nested nil", func(t *testing.T) {
		assert.True(t, containsNilValues(map[string]any{
			"identity": map[string]any{
				"userAssignedIdentities": map[string]any{
					"msi1": struct{}{},
					"msi2": nil,
				},
			},
		}))
	})
	t.Run("deeply nested no nil", func(t *testing.T) {
		assert.False(t, containsNilValues(map[string]any{
			"identity": map[string]any{
				"userAssignedIdentities": map[string]any{
					"msi1": struct{}{},
					"msi2": struct{}{},
				},
			},
		}))
	})
}

// TestHandleErrorWithCheckpoint_NotFound is a regression test for issue #2816 and related issues
// (#867, #1126, #1138, #1681, #2633, #2916): when a create/update's long-running operation
// ultimately fails after an initial 202 Accepted response, the resource never actually came into
// existence. The subsequent checkpoint read confirms this with a 404. In that case, the original
// creation/update error should be returned directly instead of being wrapped into a confusing
// compound "resource created but read failed 404 ...: <original error>" message.
func TestHandleErrorWithCheckpoint_NotFound(t *testing.T) {
	originalErr := errors.New(`Code="BadRequest" Message="The provided principal ID was not found in the AAD tenant(s)"`)

	client := &azure.MockAzureClient{
		GetResponseErr: &azure.PulumiAzcoreResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "NotFound",
			Message:    `Unable to find a SQL Role Assignment with ID [...]`,
		},
	}
	res := &resources.AzureAPIResource{APIVersion: "2023-04-15"}
	crudClient := NewResourceCrudClient(client, nil, nil, "123", res)

	err := crudClient.HandleErrorWithCheckpoint(context.Background(), originalErr,
		"/subscriptions/123/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/acct/sqlRoleAssignments/abc",
		resource.PropertyMap{}, nil)

	require.Error(t, err)
	assert.Equal(t, originalErr.Error(), err.Error())
	assert.NotContains(t, err.Error(), "read failed")
}

// TestHandleErrorWithCheckpoint_OtherReadError verifies the pre-existing behavior is preserved
// when the checkpoint read fails for a reason other than a confirmed 404: since we can't tell
// whether the resource exists, we still combine both errors into a partial error so a subsequent
// operation can attempt to reconcile the state.
func TestHandleErrorWithCheckpoint_OtherReadError(t *testing.T) {
	originalErr := errors.New("original creation error")

	client := &azure.MockAzureClient{
		GetResponseErr: &azure.PulumiAzcoreResponseError{
			StatusCode: http.StatusTooManyRequests,
			ErrorCode:  "TooManyRequests",
			Message:    "throttled",
		},
	}
	res := &resources.AzureAPIResource{APIVersion: "2023-04-15"}
	crudClient := NewResourceCrudClient(client, nil, nil, "123", res)

	err := crudClient.HandleErrorWithCheckpoint(context.Background(), originalErr,
		"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/sa",
		resource.PropertyMap{}, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "read failed")
	assert.Contains(t, err.Error(), "original creation error")
}
