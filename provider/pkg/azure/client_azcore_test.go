// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitPipelineOpts(t *testing.T) {
	t.Run("retry delays", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.InDelta(t, 20*time.Second, opts.Retry.RetryDelay, 10.0)
		assert.InDelta(t, 120*time.Second, opts.Retry.MaxRetryDelay, 30.0)
	})

	t.Run("cloud is public", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.Equal(t, cloud.AzurePublic, opts.ClientOptions.Cloud)
	})

	t.Run("cloud is usgov", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzureGovernment, nil)
		assert.Equal(t, cloud.AzureGovernment, opts.ClientOptions.Cloud)
	})

	t.Run("cloud is china", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzureChina, nil)
		assert.Equal(t, cloud.AzureChina, opts.ClientOptions.Cloud)
	})

	t.Run("should retry", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.NotNil(t, opts.Retry.ShouldRetry)
	})

	t.Run("retries on 408 timeout", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusRequestTimeout}, nil))
	})

	t.Run("retries on 409 conflict when another operation is in progress", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		header := http.Header{}
		header.Add("x-ms-error-code", "AnotherOperationInProgress")
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{
			StatusCode: http.StatusConflict,
			Header:     header,
		}, nil))
	})

	t.Run("doesn't retry on 409 conflict when no other operation is in progress", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.False(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusConflict}, nil))
	})

	t.Run("retries on 429 too many requests", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusTooManyRequests}, nil))
	})

	t.Run("retries on 500 internal server error", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusInternalServerError}, nil))
	})

	t.Run("retries on 502 bad gateway", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusBadGateway}, nil))
	})

	t.Run("retries on 503 service unavailable", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusServiceUnavailable}, nil))
	})

	t.Run("retries on 504 gateway timeout", func(t *testing.T) {
		opts := initPipelineOpts(cloud.AzurePublic, nil)
		assert.True(t, opts.Retry.ShouldRetry(&http.Response{StatusCode: http.StatusGatewayTimeout}, nil))
	})
}

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

func TestInitRequestQueryParams(t *testing.T) {
	c, err := NewAzCoreClient(&fake.TokenCredential{}, "", cloud.AzurePublic, nil)
	require.NoError(t, err)
	client := c.(*azCoreClient)

	queryParams := map[string]any{
		"api-version":                 "2022-09-01",
		"string_with_unescaped_slash": "a/path",
		"string_with_escaped_slash":   "a%2Fpath",
		"bool":                        true,
		"int":                         42,
		"slice":                       []string{"a", "b"},
	}
	req, err := client.initRequest(context.Background(), http.MethodGet, "/subscriptions/123", queryParams)
	require.NoError(t, err)

	query := req.Raw().URL.Query()
	assert.Len(t, query, len(queryParams))
	assert.Equal(t, "2022-09-01", query.Get("api-version"))
	assert.Equal(t, "a/path", query.Get("string_with_unescaped_slash"))
	assert.Equal(t, "a/path", query.Get("string_with_escaped_slash"))
	assert.Equal(t, "true", query.Get("bool"))
	assert.Equal(t, "42", query.Get("int"))
	assert.Equal(t, "a", query.Get("slice"))
	assert.Equal(t, []string{"a", "b"}, query["slice"])
}

func TestInitRequestHeaders(t *testing.T) {
	c, err := NewAzCoreClient(&fake.TokenCredential{}, "extra", cloud.AzurePublic, nil)
	require.NoError(t, err)
	client := c.(*azCoreClient)

	queryParams := map[string]any{"api-version": "2022-09-01"}

	for method, contentType := range map[string]string{
		http.MethodGet:    "",
		http.MethodDelete: "",
		http.MethodPost:   "application/json; charset=utf-8",
		http.MethodPut:    "application/json; charset=utf-8",
		http.MethodPatch:  "application/json; charset=utf-8",
	} {
		req, err := client.initRequest(context.Background(), method, "/subscriptions/123", queryParams)
		require.NoError(t, err)

		headers := req.Raw().Header
		assert.Equal(t, "extra", headers.Get("User-Agent"))
		assert.Equal(t, "application/json", headers.Get("Accept"))
		assert.Equal(t, contentType, headers.Get("Content-Type"))
	}
}

func createTestClient(t *testing.T, handler func(t *testing.T, req *http.Request)) AzureClient {
	client, err := CreateTestClient(t, handler)
	require.NoError(t, err)
	return client
}

func TestRequestQueryParams(t *testing.T) {
	const resourceId = "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm"

	t.Run("GET adds API version to query", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
		})
		client.Get(context.Background(), resourceId, "2022-02-02", nil)
	})

	t.Run("GET with query params adds API version to query", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
			assert.Equal(t, "bar", q.Get("foo"))
		})
		client.Get(context.Background(), resourceId, "2022-02-02", map[string]any{"foo": "bar"})
	})

	t.Run("DELETE adds API version to query", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
		})
		client.Delete(context.Background(), resourceId, "2022-02-02", "", nil)
	})

	t.Run("DELETE combines API version with other query params", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			q := req.URL.Query()
			assert.Equal(t, "2022-02-02", q.Get("api-version"))
			assert.Equal(t, "bar", q.Get("foo"))
		})
		client.Delete(context.Background(), resourceId, "2022-02-02", "", map[string]any{"foo": "bar"})
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

func TestRequestUserAgent(t *testing.T) {
	fake := &fakeTransporter{
		responses: []*http.Response{{StatusCode: 200}},
	}
	client := newClientWithFakeTransport(fake)
	_, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
	require.NoError(t, err)

	require.Regexp(t, `^pulumi-azure-native/(.+) azsdk-go-azcore/(.+) \(.+\) pid-12345$`, fake.requests[0].Header.Get("User-Agent"))
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
		resp, err := client.Get(context.Background(), "/subscriptions/123", "2022-09-01", nil)
		require.NoError(t, err)
		require.Empty(t, resp)
	})

	t.Run("GET error", func(t *testing.T) {
		for _, statusCode := range []int{201, 202, 204, 400, 401, 403, 404, 409, 410, 500, 502, 503, 504} {
			client := newClientWithPreparedResponses([]*http.Response{{StatusCode: statusCode}})
			_, err := client.Get(context.Background(), "/subscriptions/123", "2022-09-01", nil)
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

	t.Run("PUT polling success after 202 response", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 202,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 502, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "Bad Gateway"}`)),
			},
			{
				StatusCode: 503, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "Unavailable"}`)),
			},
			{
				StatusCode: 202,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"status": "Succeeded"}`)),
			},
		})
		_, found, err := client.Put(context.Background(), "/subscriptions/123/rg/rg", nil, map[string]any{"api-version": "2022-09-01"}, "")
		require.NoError(t, err)
		assert.True(t, found)
	})

	t.Run("PUT polling success when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 201, // created
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 502, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "Bad Gateway"}`)),
			},
			{
				StatusCode: 503, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "Unavailable"}`)),
			},
			{
				StatusCode: 201,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"status": "Succeeded"}`)),
			},
		})
		_, found, err := client.Put(context.Background(), "/subscriptions/123/rg/rg", nil, map[string]any{"api-version": "2022-09-01"},
			"the actual value doesn't matter!")
		require.NoError(t, err)
		assert.True(t, found)
	})

	t.Run("DELETE polling success when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 202,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 503, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "Unavailable"}}`)),
			},
			{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"status": "Succeeded"}`)),
			},
		})
		err := client.Delete(context.Background(), "/subscriptions/123/rg/rg", "2022-09-01", "the actual value doesn't matter!", nil)
		require.NoError(t, err)
	})

	t.Run("DELETE initial success on 404 when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "ResourceNotFound"}}`)),
			},
		})
		err := client.Delete(context.Background(), "/subscriptions/123/rg/rg", "2022-09-01", "the actual value doesn't matter!", nil)
		require.NoError(t, err)
	})

	t.Run("DELETE polling success on 404 when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 202,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "ResourceNotFound"}}`)),
			},
		})
		err := client.Delete(context.Background(), "/subscriptions/123/rg/rg", "2022-09-01", "the actual value doesn't matter!", nil)
		require.NoError(t, err)
	})

	t.Run("DELETE initial failure when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "BadRequest"}}`)),
			},
		})
		err := client.Delete(context.Background(), "/subscriptions/123/rg/rg", "2022-09-01", "the actual value doesn't matter!", nil)
		require.Error(t, err)
		require.IsType(t, &PulumiAzcoreResponseError{}, err)
		require.Equal(t, "BadRequest", err.(*PulumiAzcoreResponseError).ErrorCode)
	})

	t.Run("DELETE polling failure when asyncStyle is given", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{
			{
				StatusCode: 202,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"status": "InProgress"}`)),
			},
			{
				StatusCode: 503, // temporary failure
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "Unavailable"}}`)),
			},
			{
				StatusCode: 400,
				Header:     http.Header{"Location": []string{"https://management.azure.com/operation"}},
				Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "BadRequest"}}`)),
			},
		})
		err := client.Delete(context.Background(), "/subscriptions/123/rg/rg", "2022-09-01", "the actual value doesn't matter!", nil)
		require.Error(t, err)
		require.IsType(t, &PulumiAzcoreResponseError{}, err)
		require.Equal(t, "BadRequest", err.(*PulumiAzcoreResponseError).ErrorCode)
	})
}

func TestCanCreateUsesResourcePath(t *testing.T) {
	canCreate := func(client AzureClient, path string) {
		client.CanCreate(context.Background(),
			"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
			path,
			"2022-09-01", http.MethodGet, false, false, nil)
	}
	t.Run("no path", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm", req.URL.Path)
		})
		canCreate(client, "")
	})

	t.Run("path", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm/extraPath", req.URL.Path)
		})
		canCreate(client, "/extraPath")
	})

}

func TestCanCreate_Responses(t *testing.T) {
	id := "/subscriptions/123/resourceGroups/rg"

	type testCase struct {
		responseStatus      int
		response            string
		id                  string
		readMethod          string
		isSingletonResource bool // if true and 200 OK -> true
		hasDefaultBody      bool // if true and 200 OK -> isDefaultResponse
		isDefaultResponse   bool
	}

	t.Run("200 OK", func(t *testing.T) {
		// HTTP method makes no difference when the response is 200 OK.
		for _, method := range []string{http.MethodGet, http.MethodHead} {
			// Success cases
			for _, tc := range []testCase{
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: false, hasDefaultBody: false, isDefaultResponse: false},
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: true, hasDefaultBody: false, isDefaultResponse: false},
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: true, hasDefaultBody: false, isDefaultResponse: true},
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: false, hasDefaultBody: true, isDefaultResponse: true},
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: false, hasDefaultBody: false, isDefaultResponse: false},
			} {
				client := newClientWithPreparedResponses([]*http.Response{{
					StatusCode: tc.responseStatus,
					Body:       io.NopCloser(strings.NewReader(tc.response)),
				}})
				err := client.CanCreate(context.Background(), tc.id, "" /* path */, "2022-09-01", tc.readMethod,
					tc.isSingletonResource, tc.hasDefaultBody, func(map[string]any) bool { return tc.isDefaultResponse })
				require.NoError(t, err)
			}

			// Error cases
			for _, tc := range []testCase{
				// This resource has a default state but the response we got is not that state, i.e., it's been modified already.
				{responseStatus: 200, id: id, readMethod: method, isSingletonResource: false, hasDefaultBody: true, isDefaultResponse: false},
				// 200 OK with a reponse body means a resource exists, so we can't create it.
				{responseStatus: 200, response: `{"foo": 1}`, id: id, readMethod: method, isSingletonResource: false, hasDefaultBody: false, isDefaultResponse: false},
			} {
				client := newClientWithPreparedResponses([]*http.Response{{
					StatusCode: tc.responseStatus,
					Body:       io.NopCloser(strings.NewReader(tc.response)),
				}})
				err := client.CanCreate(context.Background(), tc.id, "" /* path */, "2022-09-01", tc.readMethod,
					tc.isSingletonResource, tc.hasDefaultBody, func(map[string]any) bool { return tc.isDefaultResponse })
				require.Error(t, err)
			}
		}
	})

	// 404 always means "can create"
	t.Run("404 Not Found", func(t *testing.T) {
		for _, method := range []string{http.MethodGet, http.MethodHead} {
			for _, resp := range []string{``, `{"foo": 1}`} {
				for _, isSingleton := range []bool{true, false} {
					for _, hasDefault := range []bool{true, false} {
						for _, isDefault := range []bool{true, false} {
							client := newClientWithPreparedResponses([]*http.Response{{
								StatusCode: 404,
								Body:       io.NopCloser(strings.NewReader(resp)),
							}})
							err := client.CanCreate(context.Background(), id, "" /* path */, "2022-09-01", method,
								isSingleton, hasDefault, func(map[string]any) bool { return isDefault })
							require.NoError(t, err)
						}
					}
				}
			}
		}
	})

	// 204 No Content always means "can create" when read via GET, but never when read via HEAD
	t.Run("204 No Content", func(t *testing.T) {
		for _, resp := range []string{``, `{"foo": 1}`} {
			for _, isSingleton := range []bool{true, false} {
				for _, hasDefault := range []bool{true, false} {
					for _, isDefault := range []bool{true, false} {
						client := newClientWithPreparedResponses([]*http.Response{{
							StatusCode: 204,
							Body:       io.NopCloser(strings.NewReader(resp)),
						}})
						err := client.CanCreate(context.Background(), id, "" /* path */, "2022-09-01", "GET",
							isSingleton, hasDefault, func(map[string]any) bool { return isDefault })
						require.NoError(t, err)
					}
				}
			}
		}

		for _, resp := range []string{``, `{"foo": 1}`} {
			for _, isSingleton := range []bool{true, false} {
				for _, hasDefault := range []bool{true, false} {
					for _, isDefault := range []bool{true, false} {
						client := newClientWithPreparedResponses([]*http.Response{{
							StatusCode: 204,
							Body:       io.NopCloser(strings.NewReader(resp)),
						}})
						err := client.CanCreate(context.Background(), id, "" /* path */, "2022-09-01", "HEAD",
							isSingleton, hasDefault, func(map[string]any) bool { return isDefault })
						require.Error(t, err)
					}
				}
			}
		}
	})

	// Other HTTP status codes are always an error.
	t.Run("other HTTP statuses", func(t *testing.T) {
		for _, status := range []int{201, 202, 400, 401, 403, 409, 410, 500, 502, 503, 504} {
			for _, method := range []string{http.MethodGet, http.MethodHead} {
				for _, resp := range []string{``, `{"foo": 1}`} {
					for _, isSingleton := range []bool{true, false} {
						for _, hasDefault := range []bool{true, false} {
							for _, isDefault := range []bool{true, false} {
								client := newClientWithPreparedResponses([]*http.Response{{
									StatusCode: status,
									Body:       io.NopCloser(strings.NewReader(resp)),
								}})
								err := client.CanCreate(context.Background(), id, "" /* path */, "2022-09-01", method,
									isSingleton, hasDefault, func(map[string]any) bool { return isDefault })
								require.Error(t, err)
							}
						}
					}
				}
			}
		}
	})
}

// Implements azcore's policy.Transporter by returning the given responses in order.
type fakeTransporter struct {
	requests  []*http.Request
	responses []*http.Response
	index     int
}

func (f *fakeTransporter) Do(req *http.Request) (*http.Response, error) {
	f.requests = append(f.requests, req)
	cur := f.responses[f.index]
	f.index++
	return cur, nil
}

func newClientWithFakeTransport(transport *fakeTransporter) *azCoreClient {
	opts := arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Transport: transport,
			Retry:     policy.RetryOptions{MaxRetries: -1},
			Cloud:     cloud.AzurePublic,
			Telemetry: policy.TelemetryOptions{
				ApplicationID: fmt.Sprintf("pulumi-azure-native/%s", version.GetVersion()),
			},
		},
	}

	client, _ := NewAzCoreClient(&fake.TokenCredential{}, "pid-12345", cloud.AzurePublic, &opts)
	azCoreClient := client.(*azCoreClient)
	azCoreClient.updatePollingIntervalSeconds = 1
	azCoreClient.deletePollingIntervalSeconds = 1
	return azCoreClient
}

func newClientWithPreparedResponses(responses []*http.Response) *azCoreClient {
	return newClientWithFakeTransport(&fakeTransporter{
		responses: responses,
	})
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

func TestPutPatchRequestBody(t *testing.T) {
	qp := map[string]any{"api-version": "2022-09-01"}

	t.Run("PUT", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, `{"foo":"bar"}`, string(body))
		})
		client.Put(context.Background(), "/subscriptions/123", map[string]any{"foo": "bar"}, qp, "")
	})

	t.Run("PATCH", func(t *testing.T) {
		client := createTestClient(t, func(t *testing.T, req *http.Request) {
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, `{"foo":{"bar":11}}`, string(body))
		})
		client.Patch(context.Background(), "/subscriptions/123",
			map[string]any{"foo": map[string]any{"bar": 11}},
			qp, "")
	})
}

func TestShouldRetryConflict(t *testing.T) {
	shouldRetry := func(status int, body string) bool {
		resp := http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		}
		return shouldRetryConflict(&resp)
	}
	t.Run("Not a conflict", func(t *testing.T) {
		assert.False(t, shouldRetry(200, ""))
	})
	t.Run("Blank body", func(t *testing.T) {
		assert.False(t, shouldRetry(409, ""))
	})
	t.Run("Unreadable body", func(t *testing.T) {
		resp := http.Response{
			StatusCode: 409,
			Body:       errorReadCloser{},
		}
		assert.False(t, shouldRetryConflict(&resp))
	})
	t.Run("empty json", func(t *testing.T) {
		assert.False(t, shouldRetry(409, "{}"))
	})
	t.Run("wrong error type", func(t *testing.T) {
		assert.False(t, shouldRetry(409, `{"error": "Conflict"}`))
	})
	t.Run("no error code", func(t *testing.T) {
		assert.False(t, shouldRetry(409, `{"error": {"message": "Conflict"}}`))
	})
	t.Run("other error code", func(t *testing.T) {
		assert.False(t, shouldRetry(409, `{"error": {"code": "Conflict"}}`))
	})
	t.Run("AnotherOperationInProgress", func(t *testing.T) {
		assert.True(t, shouldRetry(409, `{"error": {"code": "AnotherOperationInProgress"}}`))
	})
	t.Run("ConcurrentFederatedIdentityCredentialsWritesForSingleManagedIdentity", func(t *testing.T) {
		assert.True(t, shouldRetry(409, `{"error": {"code": "ConcurrentFederatedIdentityCredentialsWritesForSingleManagedIdentity"}}`))
	})
}

type errorReadCloser struct{}

func (errorReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}
func (errorReadCloser) Close() error {
	return nil
}

func TestPostResponsesCanBeAnything(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		resp := &http.Response{
			Body: io.NopCloser(strings.NewReader(`"hello"`)),
		}
		val, err := readResponse(resp)
		require.NoError(t, err)
		assert.Equal(t, "hello", val)
	})

	t.Run("object", func(t *testing.T) {
		resp := &http.Response{
			Body: io.NopCloser(strings.NewReader(`{"k": 1}`)),
		}
		val, err := readResponse(resp)
		require.NoError(t, err)
		assert.Equal(t, map[string]any{"k": 1.0}, val)
	})
}

func TestNewResponseError(t *testing.T) {
	t.Run("standard azcore error", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 409,
			Body:       io.NopCloser(strings.NewReader(`{"error": {"message": "Foo already exists"}}`)),
			Header:     http.Header{"X-Ms-Error-Code": []string{"Conflict"}},
		}
		err := newResponseError(resp)
		require.Error(t, err)
		assert.Equal(t, "Status=409 Code=\"Conflict\" Message=\"Foo already exists\"", err.Error())
	})

	t.Run("no message", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 409,
			Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "Conflict"}}`)),
			Header:     http.Header{"X-Ms-Error-Code": []string{"Conflict"}},
		}
		err := newResponseError(resp)
		require.Error(t, err)
		assert.Equal(t, `Status=409 Code="Conflict" Message="{"error": {"code": "Conflict"}}"`, err.Error())
	})

	t.Run("unknown error", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 409,
			Body:       io.NopCloser(strings.NewReader(`{"foo": "bar"}`)),
		}
		err := newResponseError(resp)
		require.Error(t, err)
		assert.Equal(t, `Status=409 Message="{"foo": "bar"}"`, err.Error())
	})

	t.Run("404 with valid error code is recognized by IsNotFound", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(strings.NewReader(`{"error": {"code": "ResourceNotFound", "message": "not found"}}`)),
			Header:     http.Header{"X-Ms-Error-Code": []string{"ResourceNotFound"}},
		}
		assert.True(t, IsNotFound(newResponseError(resp)))
	})

	t.Run("404 without error code is NOT recognized by IsNotFound", func(t *testing.T) {
		// This simulates a proxy/WAF 404 response without Azure error codes
		resp := &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(strings.NewReader(`not found`)),
		}
		assert.False(t, IsNotFound(newResponseError(resp)))
	})
}
