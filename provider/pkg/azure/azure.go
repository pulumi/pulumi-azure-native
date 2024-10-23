// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// BuildUserAgent composes a User Agent string with the provided partner ID.
func BuildUserAgent(partnerID string) (userAgent string) {
	userAgent = strings.TrimSpace(fmt.Sprintf("%s pulumi-azure-native/%s",
		autorest.UserAgent(), version.Version))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	// Append partner ID, if it's defined.
	if partnerID != "" {
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, partnerID)
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

// GetCloudByName returns the azure-sdk-for-go/sdk/azcore/cloud configuration for the given cloud.
// Valid names are as documented in the provider's installation & configuration guide, currently
// public, china, usgovernment, or the empty value for public.
// NOTE: this method doesn't do any validation. If an unknown cloud is given, it falls through to
// the default public cloud. It's assumed that validation of cloud name in the provider's config
// has been done earlier.
func GetCloudByName(cloudName string) azcloud.Configuration {
	switch cloudName {
	case "china":
		return azcloud.AzureChina
	case "usgov":
		return azcloud.AzureGovernment
	case "usgovernment":
		return azcloud.AzureGovernment
	}
	return azcloud.AzurePublic
}
