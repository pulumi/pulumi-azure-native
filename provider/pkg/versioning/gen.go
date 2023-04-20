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
	AllResourcesByVersion         ProvidersVersionResources
	AllResourceVersionsByResource ProviderResourceVersions
	V1Lock                        openapi.DefaultVersionLock
	Deprecated                    openapi.ProviderVersionList
	Active                        providerlist.ProviderPathVersionsJson
	Pending                       openapi.ProviderVersionList
	V2Spec                        Spec
	V2Lock                        openapi.DefaultVersionLock
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
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	// provider->version->[]resource
	allResourcesByVersion := FindAllResources(providers)

	// ProviderName->ProviderSpec (tracking + additions)
	v1Spec := versionSources.v1Spec
	v1Lock, err := DefaultConfigToDefaultVersionLock(allResourcesByVersion, v1Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	deprecated := FindDeprecations(allResourcesByVersion, v1Lock)
	allResourcesByVersionWithoutDeprecations := RemoveDeprecations(allResourcesByVersion, deprecated)

	v2Spec := versionSources.v2Spec

	v2Config := versionSources.v2Config
	v2Spec = BuildSpec(allResourcesByVersionWithoutDeprecations, v2Config, v2Spec)

	violations := ValidateDefaultConfig(v2Spec, v2Config)
	PrintViolationsAsWarnings(versionSources.v2ConfigPath, violations)
	if err != nil {
		return VersionMetadata{}, err
	}

	squeezedResources := allResourcesByVersionWithoutDeprecations
	if majorVersion == V2 {
		squeezedResources = SqueezeSpec(allResourcesByVersionWithoutDeprecations, versionSources.v1Squeeze)
	}

	v2Lock, err := DefaultConfigToDefaultVersionLock(squeezedResources, v2Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	// provider->resource->[]version
	allResourceVersionsByResource := FormatResourceVersions(allResourcesByVersion)

	return VersionMetadata{
		AllResourcesByVersion:         allResourcesByVersion,
		AllResourceVersionsByResource: allResourceVersionsByResource,
		V1Lock:                        v1Lock,
		Deprecated:                    deprecated,
		Active:                        activePathVersionsJson,
		Pending:                       FindNewerVersions(allResourcesByVersion, v1Lock),
		V2Spec:                        v2Spec,
		V2Lock:                        v2Lock,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) error {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"allResourcesByVersion.json":         v.AllResourcesByVersion,
		"allResourceVersionsByResource.json": v.AllResourceVersionsByResource,
		"v1-lock.json":                       v.V1Lock,
		"deprecated.json":                    v.Deprecated,
		"active.json":                        v.Active,
		"pending.json":                       v.Pending,
		"v2-spec.yaml":                       v.V2Spec,
		"v2-lock.json":                       v.V2Lock,
	})
}

type VersionSources struct {
	activePathVersions providerlist.ProviderPathVersions
	v1Spec             Spec
	v2Spec             Spec
	v2Config           Curations
	v2ConfigPath       string
	v1Squeeze          Squeeze
}

func ReadVersionSources(rootDir string) (VersionSources, error) {
	activePathVersions, err := providerlist.ReadProviderList(filepath.Join(rootDir, "azure-provider-versions", "provider_list.json"))
	if err != nil {
		return VersionSources{}, err
	}

	v1Spec, err := ReadSpec(path.Join(rootDir, "versions", "v1-spec.yaml"))
	if err != nil {
		return VersionSources{}, err
	}

	v2Spec, err := ReadSpec(path.Join(rootDir, "versions", "v2-spec.yaml"))
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
