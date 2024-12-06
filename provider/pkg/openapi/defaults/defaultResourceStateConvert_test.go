package defaults_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllDefaultStatesConvertable(t *testing.T) {
	metadataBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "..", "bin", "metadata-compact.json"))
	require.Nil(t, err)
	metadata, err := provider.LoadMetadata(metadataBytes)
	require.Nil(t, err)
	resourcesByNormalisedPath := map[string][]string{}
	for resourceToken, resource := range metadata.Resources {
		normalisedPath := paths.NormalizePath(resource.Path)
		resourcesByNormalisedPath[normalisedPath] = append(resourcesByNormalisedPath[normalisedPath], resourceToken)
	}

	resourceTokenVersionMatcher := regexp.MustCompile(`/(v[^:]+)`)

	for _, path := range defaults.ListPathsWithDefaults() {
		cleaned := strings.ReplaceAll(path, "/", "_")
		// Remove version query parameter
		pathWithoutVersion := path
		if idx := strings.Index(path, "?"); idx != -1 {
			pathWithoutVersion = path[:idx]
		}
		t.Run(cleaned, func(t *testing.T) {
			resourceTokens, found := resourcesByNormalisedPath[pathWithoutVersion]
			require.Truef(t, found, "Resource not found in test data: %s", pathWithoutVersion)
			for _, resourceToken := range resourceTokens {
				var apiVersion openapi.ApiVersion
				sdkVersionMatch := resourceTokenVersionMatcher.FindStringSubmatch(resourceToken)
				if len(sdkVersionMatch) > 1 {
					apiVersion, err = openapi.SdkToApiVersion(openapi.SdkVersion(sdkVersionMatch[1]))
					require.Nil(t, err, "Failed to convert SDK version to API version: %s", sdkVersionMatch[1])
				}
				defaultState := defaults.GetDefaultResourceState(path, string(apiVersion))
				if defaultState == nil || defaultState.SkipDelete {
					return
				}
				resource, found := metadata.Resources[resourceToken]
				require.Truef(t, found, "Resource not found in metadata: %s", resourceToken)
				converted, err := crud.PrepareAzureRESTBody("", resource.PutParameters, [][]string{}, defaultState.State, &convert.SdkShapeConverter{})
				assert.Nil(t, err, "Failed to prepare body for %s", resourceToken)
				assert.NotNil(t, converted, "No body returned for %s", resourceToken)
			}
		})
	}
}
