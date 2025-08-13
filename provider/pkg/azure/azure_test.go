// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
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
		{name: "azureusgovernmentcloud", expected: cloud.AzureGovernment},
		{name: "AzureUSGovernmentCloud", expected: cloud.AzureGovernment},
	} {
		assert.Equal(t, tc.expected, GetCloudByName(tc.name), tc.name)
	}
}

func TestBuildUserAgent(t *testing.T) {
	tests := []struct {
		azcore    bool
		name      string
		partnerID string
		ExtraUA   string
		wantRegex string
	}{
		{
			azcore:    true,
			name:      "default",
			wantRegex: ``,
		},
		{
			azcore:    true,
			name:      "PartnerID",
			partnerID: "12345",
			wantRegex: `pid-12345`,
		},
		{
			azcore:    true,
			name:      "UserAgentPassthrough",
			ExtraUA:   "a/1.2.3 b-c",
			wantRegex: `a/(.+) b-c`,
		},
		{
			azcore:    false,
			name:      "legacy:default",
			wantRegex: `go-autorest/(.+) pulumi-azure-native/(.+)`,
		},
		{
			azcore:    false,
			name:      "legacy:PartnerID",
			partnerID: "12345",
			wantRegex: `go-autorest/(.+) pulumi-azure-native/(.+) pid-12345`,
		},
		{
			azcore:    false,
			name:      "legacy:UserAgentPassthrough",
			ExtraUA:   "a/1.2.3 b-c",
			wantRegex: `go-autorest/(.+) pulumi-azure-native/(.+) a/(.+) b-c`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", strconv.FormatBool(tc.azcore))
			os.Setenv("AZURE_HTTP_USER_AGENT", tc.ExtraUA)

			ua := BuildUserAgent(tc.partnerID)
			matched, err := regexp.MatchString(tc.wantRegex, ua)
			if err != nil || !matched {
				t.Errorf("user agent mismatch, expected %q: got %q", tc.wantRegex, ua)
			}
		})
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
