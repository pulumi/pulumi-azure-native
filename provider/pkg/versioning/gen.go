package versioning

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
	"gopkg.in/yaml.v3"
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

func GenerateVersionMetadata(rootDir string, providers openapi.AzureProviders) (VersionMetadata, error) {
	versionSources, err := ReadVersionSources(rootDir)
	if err != nil {
		return VersionMetadata{}, err
	}
	return calculateVersionMetadata(versionSources, providers)
}

func calculateVersionMetadata(versionSources VersionSources, providers map[string]map[string]openapi.VersionResources) (VersionMetadata, error) {
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

	v2Lock, err := DefaultConfigToDefaultVersionLock(allResourcesByVersionWithoutDeprecations, v2Spec)
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
	activePathVersions        providerlist.ProviderPathVersions
	requiredExplicitResources []string
	v1Spec                    Spec
	v2Spec                    Spec
	v2Config                  Curations
	v2ConfigPath              string
	v1Squeeze                 ResourceRemovals
	V2ResourcesToRemove       ResourceRemovals
}

func ReadVersionSources(rootDir string) (VersionSources, error) {
	activePathVersions, err := providerlist.ReadProviderList(filepath.Join(rootDir, "azure-provider-versions", "provider_list.json"))
	if err != nil {
		return VersionSources{}, err
	}

	knownExplicitResources, err := ReadRequiredExplicitResources(path.Join(rootDir, "versions", "required-explicit-resources.txt"))
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
	v1Squeeze, err := ReadResourceRemovals(v1SqueezePath)
	if err != nil {
		log.Printf("Could not read v1-squeeze: %v, proceeding without", err)
	} else {
		log.Printf("Read %d squeeze entries from %s", len(v1Squeeze), v1SqueezePath)
	}

	v2ResourcesToRemovePath := path.Join(rootDir, "versions", "v2-removed-resources.yaml")
	v2ResourcesToRemove, err := ReadResourceRemovals(v2ResourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read v2-removed-resources: %v", err)
	}

	return VersionSources{
		activePathVersions:        activePathVersions,
		requiredExplicitResources: knownExplicitResources,
		v1Spec:                    v1Spec,
		v2Spec:                    v2Spec,
		v2Config:                  v2Config,
		v2ConfigPath:              v2ConfigPath,
		v1Squeeze:                 v1Squeeze,
		V2ResourcesToRemove:       v2ResourcesToRemove,
	}, nil
}

// ReadRequiredExplicitResources reads a list of resource tokens which map to an equivalent resource which can be migrated to.
type ResourceRemovals map[string]string

func ReadResourceRemovals(path string) (ResourceRemovals, error) {
	isYaml := strings.HasSuffix(path, ".yaml")
	sourceFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()

	byteValue, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		return nil, err
	}

	result := make(ResourceRemovals)
	if isYaml {
		err = yaml.Unmarshal(byteValue, &result)
	} else {
		err = json.Unmarshal(byteValue, &result)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s ResourceRemovals) PreserveResources(tokens []string) {
	for _, token := range tokens {
		_, ok := s[token]
		if ok {
			delete(s, token)
		}
	}
}

func ReadRequiredExplicitResources(path string) ([]string, error) {
	txtFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer txtFile.Close()
	// Read each line into an array
	bytes, err := ioutil.ReadAll(txtFile)
	if err != nil {
		return nil, err
	}

	// Split on new line
	lines := strings.Split(string(bytes), "\r")
	return lines, nil
}
