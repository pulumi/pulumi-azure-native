package versioning

import (
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

// providerCuration contains manual  edits to the automatically determined API versions for a resource provider
type providerCuration struct {
	// Exclude these resources from the provider. Used when generating the final vN.json from vN-config.yaml.
	Exclusions []string
	// Don't use a tracking version, list all resources with their API version instead. Used when generating vN-config.yaml.
	Explicit bool
	// Either "exclude" or "prefer"
	Preview string
}

// Curations contains manual edits to the automatically determined API versions
type Curations map[openapi.ProviderName]providerCuration

func (c *providerCuration) IsExcluded(resource string) bool {
	if c == nil || c.Exclusions == nil {
		return false
	}
	for _, ex := range c.Exclusions {
		if ex == resource {
			return true
		}
	}
	return false
}

func (c Curations) IsExcluded(provider openapi.ProviderName, resource string) bool {
	curation, ok := c[provider]
	return ok && curation.IsExcluded(resource)
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
