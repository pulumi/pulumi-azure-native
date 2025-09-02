// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	_ "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// BuildUserAgent composes a User Agent string with the provided partner ID.
// see: https://azure.github.io/azure-sdk/general_azurecore.html#telemetry-policy
func BuildUserAgent(partnerID string) (userAgent string) {
	if !util.EnableAzcoreBackend() {
		userAgent = strings.TrimSpace(fmt.Sprintf("%s pulumi-azure-native/%s",
			autorest.UserAgent(), version.GetVersion()))
	}

	// azure-sdk-for-go sets a user agent string as per the telemetry policy, resembling:
	//   pulumi-azure-native/3.0.0 azsdk-go-azcore/1.0.0 go/1.16.5 (darwin; amd64)
	// Anything we add here will be appended to that.

	// append the CloudShell version to the user agent if it exists
	// https://github.com/Azure/azure-cli/issues/21808
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = strings.TrimSpace(fmt.Sprintf("%s %s", userAgent, azureAgent))
	}

	// Append partner ID, if it's defined.
	if partnerID != "" {
		userAgent = strings.TrimSpace(fmt.Sprintf("%s pid-%s", userAgent, partnerID))
	}

	logging.V(9).Infof("AzureNative User Agent: %s", userAgent)
	return
}

// IsNotFound returns true if the error is a HTTP 404.
func IsNotFound(err error) bool {
	if requestError, ok := err.(azure.RequestError); ok {
		return requestError.StatusCode == http.StatusNotFound
	}
	if requestError, ok := err.(*azure.RequestError); ok {
		return requestError.StatusCode == http.StatusNotFound
	}

	if responseError, ok := err.(*azcore.ResponseError); ok {
		return responseError.StatusCode == http.StatusNotFound
	}

	if responseError, ok := err.(*PulumiAzcoreResponseError); ok {
		return responseError.StatusCode == http.StatusNotFound
	}

	return false
}

// AzureError catches common errors and substitutes them with more user-friendly ones.
func AzureError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return errors.New("operation timed out")
	}
	if requestError, ok := err.(azure.RequestError); ok {
		if requestError.DetailedError.Message != "" {
			return fmt.Errorf("%w. %s", err, requestError.DetailedError.Message)
		}
	}
	return err
}

// Claims is used to unmarshall the claims from a JWT issued by the Microsoft Identity Platform.
type Claims struct {
	Audience          string   `json:"aud"`
	Issuer            string   `json:"iss"`
	IdentityProvider  string   `json:"idp"`
	ObjectId          string   `json:"oid"`
	Roles             []string `json:"roles"`
	Scopes            string   `json:"scp"`
	Subject           string   `json:"sub"`
	TenantRegionScope string   `json:"tenant_region_scope"`
	TenantId          string   `json:"tid"`
	Version           string   `json:"ver"`

	AppDisplayName string `json:"app_displayname,omitempty"`
	AppId          string `json:"appid,omitempty"`
	IdType         string `json:"idtyp,omitempty"`
}

// ParseClaims retrieves and parses the claims from a JWT issued by the Microsoft Identity Platform.
func ParseClaims(token azcore.AccessToken) (Claims, error) {
	jwt := strings.Split(token.Token, ".")
	if len(jwt) != 3 {
		return Claims{}, errors.New("unexpected token format: does not have 3 parts")
	}

	payload, err := base64.RawURLEncoding.DecodeString(jwt[1])
	if err != nil {
		return Claims{}, err
	}

	var claims Claims
	err = json.Unmarshal(payload, &claims)
	return claims, err
}
