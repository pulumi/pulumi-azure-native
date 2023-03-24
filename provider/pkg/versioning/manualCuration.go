package versioning

import (
	"fmt"
	"io"
	"os"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"gopkg.in/yaml.v3"
)

const (
	// PreviewExclude means that the preview versions are excluded from consideration for the default version
	PreviewExclude = "exclude"
	// PreviewPrefer means that the preview versions are preferred over stable versions if newer
	PreviewPrefer = "prefer"
)

const (
	ExpectTrackingAny     = ""
	ExpectTrackingStable  = "stable"
	ExpectTrackingPreview = "preview"
)

const ExclusionAllVersions = "*"

// providerCuration contains manual  edits to the automatically determined API versions for a resource provider
type providerCuration struct {
	// Exclude these resources from the provider. Used when generating the final vN.json from vN-config.yaml.
	// The value is the API version to exclude, or `*` to exclude all versions.
	Exclusions map[openapi.ResourceName]openapi.ApiVersion `yaml:"exclusions,omitempty"`
	// Don't use a tracking version, list all resources with their API version instead. Used when generating vN-config.yaml.
	Explicit bool `yaml:"explicit"`
	// Either "exclude" or "prefer"
	Preview string `yaml:"preview"`
	// Either "stable" or "preview" - defaults to expecting stable tracking version
	ExpectTracking string `yaml:"expectTracking"`
	// Whether to expect additions to the provider. Defaults to false.
	ExpectAdditions bool `yaml:"expectAdditions"`
}

// Curations contains manual edits to the automatically determined API versions
type Curations map[openapi.ProviderName]providerCuration

func (curation providerCuration) IsExcluded(resourceName openapi.ResourceName, apiVersion openapi.ApiVersion) (bool, error) {
	if curation.Exclusions == nil {
		return false, nil
	}
	excludedMaxVersion, ok := curation.Exclusions[resourceName]
	if !ok {
		return false, nil
	}
	if excludedMaxVersion == ExclusionAllVersions {
		return true, nil
	}
	if apiVersion > excludedMaxVersion {
		return false, fmt.Errorf("version %s is greater than %s", apiVersion, excludedMaxVersion)
	}
	return true, nil
}

// The error is returned when the exclusion is specified but the range doesn't include the requested version
func (c Curations) IsExcluded(providerName openapi.ProviderName, resourceName openapi.ResourceName, apiVersion openapi.ApiVersion) (bool, error) {
	return c[providerName].IsExcluded(resourceName, apiVersion)
}

func ReadManualCurations(path string) (Curations, error) {
	yamlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	var cur Curations
	err = yaml.Unmarshal(byteValue, &cur)
	return cur, err
}
