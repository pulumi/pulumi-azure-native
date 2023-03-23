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
	V1Lock        openapi.CuratedVersion
	Deprecated    openapi.ProviderVersionList
	Active        providerlist.ProviderPathVersionsJson
	Pending       openapi.ProviderVersionList
	V2Spec        DefaultConfig
	V2Lock        openapi.CuratedVersion
}

func GenerateVersionMetadata(providers openapi.AzureProviders) (VersionMetadata, error) {
	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return VersionMetadata{}, err
	}

	specVersions := FindSpecVersions(providers)
	specResourceVersions := FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1Config, err := ReadDefaultConfig(path.Join("versions", "v1-spec.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}
	v1, err := DefaultConfigToCuratedVersion(specVersions, v1Config)
	if err != nil {
		return VersionMetadata{}, err
	}

	deprecated := FindDeprecations(specVersions, v1)
	specAfterRemovals := RemoveDeprecations(specVersions, deprecated)

	v2Config, err := ReadDefaultConfig(path.Join("versions", "v2-spec.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}

	v2curations, err := ReadManualCurations(path.Join("versions", "config.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}
	v2Config = BuildDefaultConfig(specAfterRemovals, v2curations, v2Config)

	violations := ValidateDefaultConfig(v2Config, v2curations)
	if len(violations) > 0 {
		fmt.Printf("Warning: %d curation violations found:\n", len(violations))
		for _, v := range violations {
			fmt.Printf("  %s: %s\n", v.Provider, v.Detail)
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
		V1Lock:        v1,
		Deprecated:    deprecated,
		Active:        activePathVersionsJson,
		Pending:       FindNewerVersions(specVersions, v1),
		V2Spec:        v2Config,
		V2Lock:        v2,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) error {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"spec.json":           v.Spec,
		"spec-resources.json": v.SpecResources,
		"v1-lock.json":        v.V1Lock,
		"deprecated.json":     v.Deprecated,
		"active.json":         v.Active,
		"pending.json":        v.Pending,
		"v2-spec.yaml":        v.V2Spec,
		"v2-lock.json":        v.V2Lock,
	})
}
