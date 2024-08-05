package defaults_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
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
	for _, path := range defaults.ListPathsWithDefaults() {
		cleaned := strings.ReplaceAll(path, "/", "_")
		t.Run(cleaned, func(t *testing.T) {
			resourceTokens, found := resourcesByNormalisedPath[path]
			require.Truef(t, found, "Resource not found in test data: %s", path)
			if !found {
				t.Errorf("Resource not found in mapping: %s", path)
			}
			defaultState := defaults.GetDefaultResourceState(path)
			if defaultState == nil || defaultState.SkipDelete {
				return
			}
			for _, resourceToken := range resourceTokens {
				resource, found := metadata.Resources[resourceToken]
				require.Truef(t, found, "Resource not found in metadata: %s", resourceToken)
				converted, err := crud.PrepareAzureRESTBody("", resource.PutParameters, [][]string{}, defaultState.State, &convert.SdkShapeConverter{})
				require.Nil(t, err, "Failed to prepare body for %s", resourceToken)
				assert.NotNil(t, converted, "No body returned for %s", resourceToken)
			}
		})
	}
}
