// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/go-autorest/autorest"
	autorestAzure "github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
)

func TestGetCloudByName(t *testing.T) {
	for _, tc := range []struct {
		name     string
		expected cloud.Configuration
	}{
		{name: "", expected: cloud.AzurePublic},
		{name: "public", expected: cloud.AzurePublic},
		{name: "AzureCloud", expected: cloud.AzurePublic},
		{name: "china", expected: cloud.AzureChina},
		{name: "azurechinacloud", expected: cloud.AzureChina},
		{name: "AzureChinaCloud", expected: cloud.AzureChina},
		{name: "usgov", expected: cloud.AzureGovernment},
		{name: "usgovernment", expected: cloud.AzureGovernment},
		{name: "azureusgovernment", expected: cloud.AzureGovernment},
		{name: "AzureUSGovernment", expected: cloud.AzureGovernment},
	} {
		assert.Equal(t, tc.expected, GetCloudByName(tc.name), tc.name)
	}
}

func TestIsNotFound(t *testing.T) {
	t.Run("autorest", func(t *testing.T) {
		assert.True(t, IsNotFound(&autorestAzure.RequestError{
			DetailedError: autorest.DetailedError{
				StatusCode: http.StatusNotFound,
			},
		}))
		assert.False(t, IsNotFound(&autorestAzure.RequestError{
			DetailedError: autorest.DetailedError{
				StatusCode: http.StatusForbidden,
			},
		}))
	})

	t.Run("azcore", func(t *testing.T) {
		assert.True(t, IsNotFound(&azcore.ResponseError{
			StatusCode: http.StatusNotFound,
		}))
		assert.False(t, IsNotFound(&azcore.ResponseError{
			StatusCode: http.StatusForbidden,
		}))
	})

	t.Run("provider", func(t *testing.T) {
		assert.True(t, IsNotFound(&PulumiAzcoreResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "ResourceNotFound",
		}))
		assert.False(t, IsNotFound(&PulumiAzcoreResponseError{
			StatusCode: http.StatusForbidden,
			ErrorCode:  "Unauthorized",
		}))
	})
}
