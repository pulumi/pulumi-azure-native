// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
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

func TestPOST(t *testing.T) {
	t.Run("Simple POST 200", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{{StatusCode: 200}})
		resp, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
		require.NoError(t, err)
		require.Empty(t, resp)
	})

	t.Run("POST 201 Created is ok", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{{StatusCode: 201}})
		_, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
		require.NoError(t, err)
	})

	t.Run("POST response != 200 is an error", func(t *testing.T) {
		client := newClientWithPreparedResponses([]*http.Response{{StatusCode: 400}})
		_, err := client.Post(context.Background(), "/subscriptions/123", nil, map[string]any{"api-version": "2022-09-01"})
		require.Error(t, err)
	})
}

// Implements azcore's policy.Transporter by simpling returning the responses in order.
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
		},
	}

	return NewAzCoreClient(&fake.TokenCredential{}, "pulumi", cloud.AzurePublic, &opts).(*azCoreClient)
}
