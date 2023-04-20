package versioning

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

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

type MajorVersion int

const (
	V1 MajorVersion = 1
	V2 MajorVersion = 2
)

func GenerateVersionMetadata(rootDir string, majorVersion MajorVersion, providers openapi.AzureProviders) (VersionMetadata, error) {
	versionSources, err := ReadVersionSources(rootDir)
	if err != nil {
		return VersionMetadata{}, err
	}
	return calculateVersionMetadata(majorVersion, versionSources, providers)
}

func calculateVersionMetadata(majorVersion MajorVersion, versionSources VersionSources, providers map[string]map[string]openapi.VersionResources) (VersionMetadata, error) {
	// map[LoweredProviderName]map[ResourcePath]ApiVersions
	activePathVersions := versionSources.activePathVersions

	// provider->version->[]resource
	specVersions := FindSpecVersions(providers)
	// provider->resource->[]version
	specResourceVersions := FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	// ProviderName->ProviderSpec (tracking + additions)
	v1Spec := versionSources.v1Spec
	v1Lock, err := DefaultConfigToDefaultVersionLock(specVersions, v1Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	deprecated := FindDeprecations(specVersions, v1Lock)
	specAfterRemovals := RemoveDeprecations(specVersions, deprecated)

	specAfterSqueezing := specAfterRemovals
	if majorVersion == V2 {
		specAfterSqueezing = SqueezeSpec(specAfterRemovals, versionSources.v1Squeeze)
	}

	v2Spec := versionSources.v2Spec

	v2Config := versionSources.v2Config
	v2Spec = BuildDefaultConfig(specAfterSqueezing, v2Config, v2Spec)

	violations := ValidateDefaultConfig(v2Spec, v2Config)
	PrintViolationsAsWarnings(versionSources.v2ConfigPath, violations)
	if err != nil {
		return VersionMetadata{}, err
	}
	v2Lock, err := DefaultConfigToDefaultVersionLock(specAfterSqueezing, v2Spec)
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

type VersionSources struct {
	activePathVersions providerlist.ProviderPathVersions
	v1Spec             DefaultConfig
	v2Spec             DefaultConfig
	v2Config           Curations
	v2ConfigPath       string
	v1Squeeze          Squeeze
}

func ReadVersionSources(rootDir string) (VersionSources, error) {
	activePathVersions, err := providerlist.ReadProviderList(filepath.Join(rootDir, "azure-provider-versions", "provider_list.json"))
	if err != nil {
		return VersionSources{}, err
	}

	v1Spec, err := ReadDefaultConfig(path.Join(rootDir, "versions", "v1-spec.yaml"))
	if err != nil {
		return VersionSources{}, err
	}

	v2Spec, err := ReadDefaultConfig(path.Join(rootDir, "versions", "v2-spec.yaml"))
	if err != nil {
		return VersionSources{}, err
	}

	v2ConfigPath := path.Join(rootDir, "versions", "v2-config.yaml")
	v2Config, err := ReadManualCurations(v2ConfigPath)
	if err != nil {
		return VersionSources{}, err
	}

	v1SqueezePath := path.Join(rootDir, "versions", "v1-squeeze.json")
	// TODO tkappler We read all version files for all schema versions, so for v1 there's no
	// squeeze file and v1 is required to generate the v2 squeeze file.
	v1Squeeze, err := ReadSqueeze(v1SqueezePath)
	if err != nil {
		log.Printf("Could not read v1-squeeze: %v, proceeding without", err)
	} else {
		log.Printf("Read %d squeeze entries from %s", len(v1Squeeze), v1SqueezePath)
	}

	return VersionSources{
		activePathVersions: activePathVersions,
		v1Spec:             v1Spec,
		v2Spec:             v2Spec,
		v2Config:           v2Config,
		v2ConfigPath:       v2ConfigPath,
		v1Squeeze:          v1Squeeze,
	}, nil
}

type Squeeze map[string]string

func ReadSqueeze(path string) (Squeeze, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	result := make(Squeeze)
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
