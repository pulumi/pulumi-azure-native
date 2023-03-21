package versioning

import (
	"fmt"
	"path"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
)

type VersionMetadata struct {
	Spec          SpecVersions
	SpecResources ProviderResourceVersions
	V1            openapi.CuratedVersion
	Deprecated    openapi.ProviderVersionList
	Active        providerlist.ProviderPathVersionsJson
	Pending       openapi.ProviderVersionList
	V2Config      DefaultConfig
	V2            openapi.CuratedVersion
}

func GenerateVersionMetadata(providers openapi.AzureProviders, refreshConfig bool) (VersionMetadata, error) {
	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return VersionMetadata{}, err
	}

	specVersions := FindSpecVersions(providers)
	specResourceVersions := FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1Config, err := ReadDefaultConfig(path.Join("versions", "v1-config.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}
	v1, err := DefaultConfigToCuratedVersion(specVersions, v1Config)
	if err != nil {
		return VersionMetadata{}, err
	}

	deprecated := FindDeprecations(specVersions, v1)
	specAfterRemovals := RemoveDeprecations(specVersions, deprecated)

	v2Config, err := ReadDefaultConfig(path.Join("versions", "v2-config.yaml"))

	if refreshConfig {
		curations, err := ReadManualCurations(path.Join("versions", "curations.yaml"))
		if err != nil {
			return VersionMetadata{}, err
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
		return VersionMetadata{}, err
	}
	v2, err := DefaultConfigToCuratedVersion(specAfterRemovals, v2Config)
	if err != nil {
		return VersionMetadata{}, err
	}

	return VersionMetadata{
		Spec:          specVersions,
		SpecResources: specResourceVersions,
		V1:            v1,
		Deprecated:    deprecated,
		Active:        activePathVersionsJson,
		Pending:       FindNewerVersions(specVersions, v1),
		V2Config:      v2Config,
		V2:            v2,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) error {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"spec.json":           v.Spec,
		"spec-resources.json": v.SpecResources,
		"v1.json":             v.V1,
		"deprecated.json":     v.Deprecated,
		"active.json":         v.Active,
		"pending.json":        v.Pending,
		"v2-config.yaml":      v.V2Config,
		"v2.json":             v.V2,
	})
}
