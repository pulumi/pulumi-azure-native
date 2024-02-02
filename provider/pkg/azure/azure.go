package azure

import (
	"fmt"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
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
