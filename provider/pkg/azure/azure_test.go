// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/assert"
)

func TestBuildUserAgent(t *testing.T) {
	tests := []struct {
		name      string
		partnerID string
		ExtraUA   string
		wantRegex string
	}{
		{
			name:      "default",
			wantRegex: ``,
		},
		{
			name:      "PartnerID",
			partnerID: "12345",
			wantRegex: `pid-12345`,
		},
		{
			name:      "UserAgentPassthrough",
			ExtraUA:   "a/1.2.3 b-c",
			wantRegex: `a/(.+) b-c`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Run("azcore with valid error code", func(t *testing.T) {
		// Test all valid "not found" error codes
		validCodes := []string{"NotFound", "ResourceNotFound", "ResourceGroupNotFound"}
		for _, code := range validCodes {
			assert.True(t, IsNotFound(&azcore.ResponseError{
				StatusCode: http.StatusNotFound,
				ErrorCode:  code,
			}), "Should return true for error code: "+code)
		}
	})

	t.Run("azcore with 404 but no error code", func(t *testing.T) {
		// 404 without ErrorCode (e.g., proxy/WAF response)
		assert.False(t, IsNotFound(&azcore.ResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "",
		}))
	})

	t.Run("azcore with 404 but invalid error code", func(t *testing.T) {
		// 404 with non-matching error code
		assert.False(t, IsNotFound(&azcore.ResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "SomeOtherError",
		}))
	})

	t.Run("azcore with non-404 status", func(t *testing.T) {
		assert.False(t, IsNotFound(&azcore.ResponseError{
			StatusCode: http.StatusForbidden,
			ErrorCode:  "ResourceNotFound",
		}))
	})

	t.Run("provider with valid error code", func(t *testing.T) {
		// Test all valid "not found" error codes
		validCodes := []string{"NotFound", "ResourceNotFound", "ResourceGroupNotFound"}
		for _, code := range validCodes {
			assert.True(t, IsNotFound(&PulumiAzcoreResponseError{
				StatusCode: http.StatusNotFound,
				ErrorCode:  code,
			}), "Should return true for error code: "+code)
		}
	})

	t.Run("provider with 404 but no error code", func(t *testing.T) {
		// 404 without ErrorCode (e.g., proxy/WAF response)
		assert.False(t, IsNotFound(&PulumiAzcoreResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "",
		}))
	})

	t.Run("provider with 404 but invalid error code", func(t *testing.T) {
		// 404 with non-matching error code
		assert.False(t, IsNotFound(&PulumiAzcoreResponseError{
			StatusCode: http.StatusNotFound,
			ErrorCode:  "SomeOtherError",
		}))
	})

	t.Run("provider with non-404 status", func(t *testing.T) {
		assert.False(t, IsNotFound(&PulumiAzcoreResponseError{
			StatusCode: http.StatusForbidden,
			ErrorCode:  "Unauthorized",
		}))
	})
}

func TestParseClaims(t *testing.T) {
	// Create a Claims struct and marshal it to JSON
	expectedClaims := Claims{
		Audience:          "audience",
		Issuer:            "issuer",
		IdentityProvider:  "idp",
		ObjectId:          "objectid",
		Roles:             []string{"role1", "role2"},
		Scopes:            "scope",
		Subject:           "subject",
		TenantRegionScope: "region",
		TenantId:          "tenantid",
		Version:           "1.0",
		AppDisplayName:    "appdisplayname",
		AppId:             "appid",
		IdType:            "idtype",
	}
	payload, err := json.Marshal(expectedClaims)
	assert.NoError(t, err)

	// JWT: header.payload.signature (all base64url, but only payload matters)
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none"}`))
	payloadEnc := base64.RawURLEncoding.EncodeToString(payload)
	tokenStr := header + "." + payloadEnc + ".signature"
	token := azcore.AccessToken{Token: tokenStr}

	claims, err := ParseClaims(token)
	assert.NoError(t, err)
	assert.Equal(t, expectedClaims, claims)
}

func TestParseClaims_InvalidToken(t *testing.T) {
	// Token with not enough segments
	token := azcore.AccessToken{Token: "invalidtoken"}
	_, err := ParseClaims(token)
	assert.Error(t, err)

	// Token with invalid base64 payload
	token = azcore.AccessToken{Token: "a.b@d!.c"}
	_, err = ParseClaims(token)
	assert.Error(t, err)

	// Token with invalid JSON in payload
	badPayload := base64.RawURLEncoding.EncodeToString([]byte("notjson"))
	token = azcore.AccessToken{Token: "a." + badPayload + ".c"}
	_, err = ParseClaims(token)
	assert.Error(t, err)
}
