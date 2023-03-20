package versioning

import (
	"fmt"
	"path"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
)

func GenerateVersionFiles(namespace string, refreshConfig bool) (gen.FileMap, error) {
	providers, err := openapi.SpecVersions("*")
	if err != nil {
		return nil, err
	}

	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return nil, err
	}

	specVersions := FindSpecVersions(providers)
	specResourceVersions := FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1Config, err := ReadDefaultConfig(path.Join("versions", "v1-config.yaml"))
	if err != nil {
		return nil, err
	}
	v1, err := DefaultConfigToCuratedVersion(specVersions, v1Config)
	if err != nil {
		return nil, err
	}

	deprecated := FindDeprecations(specVersions, v1)
	specAfterRemovals := RemoveDeprecations(specVersions, deprecated)

	v2Config, err := ReadDefaultConfig(path.Join("versions", "v2-config.yaml"))

	if refreshConfig {
		curations, err := ReadManualCurations(path.Join("versions", "curations.yaml"))
		if err != nil {
			return nil, err
		}
		v2Config = BuildDefaultConfig(specAfterRemovals, curations, v2Config)

		violations := ValidateDefaultConfig(v2Config, curations)
		if len(violations) > 0 {
			fmt.Printf("Warning: %d curation violations found:\n", len(violations))
			for _, v := range violations {
				fmt.Printf("  %s: %s\n", v.Provider, v.Detail)
			}
		}
	}
	if err != nil {
		return nil, err
	}
	v2, err := DefaultConfigToCuratedVersion(specAfterRemovals, v2Config)
	if err != nil {
		return nil, err
	}

	return gen.FileMap{
		"spec.json":           specVersions,
		"spec-resources.json": specResourceVersions,
		"v1.json":             v1,
		"deprecated.json":     deprecated,
		"active.json":         activePathVersionsJson,
		"pending.json":        FindNewerVersions(specVersions, v1),
		"v2-config.yaml":      v2Config,
		"v2.json":             v2,
	}, nil
}
