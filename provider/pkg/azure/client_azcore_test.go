// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeLocationHeader(t *testing.T) {
	const host = "https://management.azure.com"
	const apiVersion = "2022-09-01"

	for testName, loc := range map[string]string{
		"RelativeUrl":               "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
		"RelativeUrlWithApiVersion": "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version=" + apiVersion,
		"FullUrl":                   "https://management.azure.com/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
		"FullUrlWithApiVersion":     "https://management.azure.com/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version=" + apiVersion,
	} {
		t.Run(testName, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Location", loc)
			normalizeLocationHeader(host, apiVersion, headers)
			assert.Equal(t,
				host+"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm?api-version="+apiVersion,
				headers.Get("Location"))
		})
	}

	for testName, loc := range map[string]string{
		"MalformedUrl": "foo[:",
		"EmptyUrl":     "",
	} {
		t.Run(testName, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Location", loc)
			normalizeLocationHeader(host, apiVersion, headers)
			assert.Equal(t, loc, headers.Get("Location"))
		})
	}

	t.Run("NoLocationHeader", func(t *testing.T) {
		headers := http.Header{}
		normalizeLocationHeader(host, apiVersion, headers)
		assert.Empty(t, headers.Get("Location"))
	})
}

func TestInitRequestQueryParamsHandling(t *testing.T) {
	c, err := NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, nil)
	require.NoError(t, err)
	client := c.(*azCoreClient)

	queryParams := map[string]any{
		"api-version": "2022-09-01",
		"bool":        true,
		"int":         42,
		"slice":       []string{"a", "b"},
	}
	req, err := client.initRequest(context.Background(), http.MethodGet, "/subscriptions/123", queryParams)
	require.NoError(t, err)

	query := req.Raw().URL.Query()
	assert.Len(t, query, len(queryParams))
	assert.Equal(t, "2022-09-01", query.Get("api-version"))
	assert.Equal(t, "true", query.Get("bool"))
	assert.Equal(t, "42", query.Get("int"))
	assert.Equal(t, "a", query.Get("slice"))
	assert.Equal(t, []string{"a", "b"}, query["slice"])
}

func TestRequestQueryParams(t *testing.T) {
	const resourceId = "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm"

	createTestClient := func(t *testing.T, assertions func(t *testing.T, req *http.Request)) AzureClient {
		transp := &requestAssertingTransporter{
			t:          t,
			assertions: assertions,
		}
		opts := arm.ClientOptions{
			ClientOptions: policy.ClientOptions{
				Retry:     policy.RetryOptions{MaxRetries: -1}, // speeds up the tests
				Telemetry: policy.TelemetryOptions{Disabled: true},
				Transport: transp,
			},
			DisableRPRegistration: true,
		}
		client, _ := NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, &opts)
		return client
	}

	t.Run("GET adds API version to query", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
		})
		client.Get(context.Background(), resourceId, "2022-02-02")
	})

	t.Run("DELETE adds API version to query", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
		})
		client.Delete(context.Background(), resourceId, "2022-02-02", "", nil)
	})

	t.Run("POST adds all query params", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
			assert.Equal(t, "11", q.Get("foo"))
		})
		client.Post(context.Background(), resourceId, nil,
			map[string]any{
				"api-version": "2022-02-02",
				"foo":         11,
			})
	})

	t.Run("PUT adds all query params", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
			assert.Equal(t, "11", q.Get("foo"))
		})
		client.Put(context.Background(), resourceId, nil,
			map[string]any{
				"api-version": "2022-02-02",
				"foo":         11,
			}, "")
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

func TestErrorStatusCodes(t *testing.T) {
	t.Run("POST ok", func(t *testing.T) {
		for _, statusCode := range []int{200, 201} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			resp, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
			require.NoError(t, err, statusCode)
			require.Empty(t, resp, statusCode)
		}
	})

	t.Run("POST error", func(t *testing.T) {
		for _, statusCode := range []int{202, 204, 400, 401, 403, 404, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			_, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
			require.Error(t, err, statusCode)
		}
	})

	t.Run("GET ok", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{{StatusCode: 200}})
		resp, err := client.Get(context.Background(), "/subscriptions/123", "2022-09-01")
		require.NoError(t, err)
		require.Empty(t, resp)
	})

	t.Run("GET error", func(t *testing.T) {
		for _, statusCode := range []int{201, 202, 204, 400, 401, 403, 404, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			_, err := client.Get(context.Background(), "/subscriptions/123", "2022-09-01")
			require.Error(t, err, statusCode)
		}
	})

	t.Run("PUT, PATCH ok", func(t *testing.T) {
		for _, statusCode := range []int{200, 201} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			resp, created, err := client.Put(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"}, "")
			require.NoError(t, err, statusCode)
			assert.True(t, created)
			require.Empty(t, resp, statusCode)
		}
	})

	// This test does not cover long-running operations with polling, only the handling of the first response.
	t.Run("PUT, PATCH error", func(t *testing.T) {
		for _, statusCode := range []int{202, 204, 400, 401, 403, 404, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(``)),
				Request:    &http.Request{Method: http.MethodPut, URL: &url.URL{Path: "/subscriptions/123"}},
			}})
			_, created, err := client.Put(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"}, "")
			assert.Equal(t, statusCode < 203, created, statusCode)
			require.Error(t, err, statusCode)
		}
	})

	t.Run("DELETE ok", func(t *testing.T) {
		for _, statusCode := range []int{200, 202, 204, 404} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			err := client.Delete(context.Background(), "/subscriptions/123", "2022-09-01", "", nil)
			require.NoError(t, err, statusCode)
		}
	})

	t.Run("DELETE error", func(t *testing.T) {
		for _, statusCode := range []int{201, 400, 401, 403, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			err := client.Delete(context.Background(), "/subscriptions/123", "2022-09-01", "", nil)
			require.Error(t, err, statusCode)
		}
	})

	t.Run("HEAD ok", func(t *testing.T) {
		for _, statusCode := range []int{200, 204} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			err := client.Head(context.Background(), "/subscriptions/123", "2022-09-01")
			require.NoError(t, err, statusCode)
		}
	})

	t.Run("HEAD error", func(t *testing.T) {
		for _, statusCode := range []int{201, 202, 400, 401, 403, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			err := client.Head(context.Background(), "/subscriptions/123", "2022-09-01")
			require.Error(t, err, statusCode)
		}
	})
}

// Implements azcore's policy.Transporter by returning the given responses in order.
type fakeTransporter struct {
	responses []*http.Response
	index     int
}

func (f *fakeTransporter) Do(req *http.Request) (*http.Response, error) {
	cur := f.responses[f.index]
	f.index++
	return cur, nil
}

func newClientWithPreparedResponses(responses []*http.Response) *azCoreClient {
	opts := arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Transport: &fakeTransporter{
				responses: responses,
			},
			Retry: policy.RetryOptions{MaxRetries: -1},
		},
	}

	client, _ := NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, &opts)
	return client.(*azCoreClient)
}

func TestHandleResponseError(t *testing.T) {
	t.Run("Cannot unmarshal JSON", func(t *testing.T) {
		resp := http.Response{
			Body:   io.NopCloser(strings.NewReader("not JSON")),
			Status: "400 Bad Request",
		}

		var outputs map[string]any
		err := runtime.UnmarshalAsJSON(&resp, &outputs)

		handledErr := handleAzCoreResponseError(err, &resp)
		require.Error(t, handledErr)
		assert.Contains(t, handledErr.Error(), "not JSON")
		assert.Contains(t, handledErr.Error(), "400 Bad Request")
	})

	t.Run("no changes to error", func(t *testing.T) {
		resp := http.Response{
			Body: io.NopCloser(strings.NewReader(`{"status": "ok"}`)),
		}

		err := errors.New("some error unrelated to unmarshaling")

		handledErr := handleAzCoreResponseError(err, &resp)
		assert.Same(t, err, handledErr)
	})
}
