package versioning

import (
	"path"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
)

type VersionMetadata struct {
	Spec          SpecVersions
	SpecResources ProviderResourceVersions
	V1Lock        openapi.DefaultVersionLock
	Deprecated    openapi.ProviderVersionList
	Active        providerlist.ProviderPathVersionsJson
	Pending       openapi.ProviderVersionList
	V2Spec        DefaultConfig
	V2Lock        openapi.DefaultVersionLock
}

func GenerateVersionMetadata(providers openapi.AzureProviders) (VersionMetadata, error) {
	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return VersionMetadata{}, err
	}

	specVersions := FindSpecVersions(providers)
	specResourceVersions := FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1Spec, err := ReadDefaultConfig(path.Join("versions", "v1-spec.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}
	v1Lock, err := DefaultConfigToDefaultVersionLock(specVersions, v1Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	deprecated := FindDeprecations(specVersions, v1Lock)
	specAfterRemovals := RemoveDeprecations(specVersions, deprecated)

	v2Spec, err := ReadDefaultConfig(path.Join("versions", "v2-spec.yaml"))
	if err != nil {
		return VersionMetadata{}, err
	}

	v2ConfigPath := path.Join("versions", "v2-config.yaml")
	v2Config, err := ReadManualCurations(v2ConfigPath)
	if err != nil {
		return VersionMetadata{}, err
	}
	v2Spec = BuildDefaultConfig(specAfterRemovals, v2Config, v2Spec)

	violations := ValidateDefaultConfig(v2Spec, v2Config)
	PrintViolationsAsWarnings(v2ConfigPath, violations)
	if err != nil {
		return VersionMetadata{}, err
	}
	v2Lock, err := DefaultConfigToDefaultVersionLock(specAfterRemovals, v2Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	return VersionMetadata{
		Spec:          specVersions,
		SpecResources: specResourceVersions,
		V1Lock:        v1Lock,
		Deprecated:    deprecated,
		Active:        activePathVersionsJson,
		Pending:       FindNewerVersions(specVersions, v1Lock),
		V2Spec:        v2Spec,
		V2Lock:        v2Lock,
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
