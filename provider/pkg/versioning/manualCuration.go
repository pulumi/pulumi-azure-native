package versioning

import (
	"io"
	"os"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"gopkg.in/yaml.v3"
)

type providerCuration struct {
	Exclusions    []string
	Explicit      bool
	PreferPreview bool
}

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
