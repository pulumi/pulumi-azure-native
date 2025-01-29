package versioning

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"gopkg.in/yaml.v3"
)

type VersionMetadata struct {
	VersionSources
	// module->resource->[]version
	AllResourceVersionsByResource ModuleResourceVersions
	// map[ModuleName][]ApiVersion
	Pending                 openapi.ModuleVersionList
	Spec                    Spec
	DefaultVersions         openapi.DefaultVersions
	CurationViolations      []CurationViolation
	InactiveDefaultVersions map[openapi.ModuleName][]openapi.ApiVersion
}

// Ensure our VersionMetadata type implements the gen.Versioning interface
// The compiler will raise an error here if the interface isn't implemented
var _ gen.Versioning = (*VersionMetadata)(nil)

// Determine if an explicit resource version is being included in the SDK.
// Version being nil indicates the default version of the resource which should always be included.
func (v VersionMetadata) ShouldInclude(moduleName openapi.ModuleName, version *openapi.ApiVersion, typeName, token string) bool {
	// Nil version indicates this is in the default version.
	if version == nil {
		return true
	}
	// Keep any resources in the default version lock
	if v.DefaultVersions.IsAtVersion(moduleName, typeName, *version) {
		return true
	}
	// Keep any resources in the previous version lock for easier migration
	if v.MajorVersion >= 3 && v.PreviousDefaultVersions.IsAtVersion(moduleName, typeName, *version) {
		// We're making an exception for these types because we want to remove their previous default version since it
		// was broken and had no Pulumi Cloud users according to Metabase. #3817
		if moduleName == "MachineLearningServices" && (typeName == "ConnectionRaiBlocklist" || typeName == "ConnectionRaiBlocklistItem") {
			return false
		}
		return true
	}
	// Exclude versions from removed versions
	if versions, ok := v.RemovedVersions[moduleName]; ok {
		for _, removedVersion := range versions {
			if removedVersion == *version {
				return false
			}
		}
	}
	// Exclude removed resources
	if _, ok := v.ResourcesToRemove[token]; ok {
		return false
	}
	// Exclude removed invokes
	if _, ok := v.RemovedInvokes[token]; ok {
		return false
	}
	return true
}

func (v VersionMetadata) GetDeprecation(token string) (gen.ResourceDeprecation, bool) {
	if newToken, ok := v.NextResourcesToRemove[token]; ok {
		return gen.ResourceDeprecation{
			ReplacementToken: newToken,
		}, true
	}
	return gen.ResourceDeprecation{}, false
}

func (v VersionMetadata) GetAllVersions(moduleName openapi.ModuleName, resource openapi.ResourceName) []openapi.ApiVersion {
	return v.AllResourceVersionsByResource[moduleName][resource]
}

// Collect, for each module, the API versions that are older than any _previous_ default version and that weren't
// explicitly removed.
func (v VersionMetadata) GetOldApiVersionsPerModule() map[openapi.ModuleName][]openapi.ApiVersion {
	result := map[openapi.ModuleName][]openapi.ApiVersion{}
	for moduleName, resourceVersions := range v.AllResourceVersionsByResource {
		// Collect the explicitly removed versions in a set for easy lookup. We don't want to list them as removable
		// when they're already removed.
		removedFromModule := map[openapi.ApiVersion]struct{}{}
		for _, version := range v.RemovedVersions[moduleName] {
			removedFromModule[version] = struct{}{}
		}

		// API versions can be removed by module, not by resource. Determine the oldest previous default version for
		// this module. It's the cut-off below which we can safely remove versions.
		previousDefaultVersions := []openapi.ApiVersion{}
		for _, version := range v.PreviousDefaultVersions[moduleName] {
			previousDefaultVersions = append(previousDefaultVersions, version.ApiVersion)
		}
		if len(previousDefaultVersions) == 0 {
			continue
		}
		oldestPreviousDefaultVersion := slices.MinFunc(previousDefaultVersions, openapi.CompareApiVersions)

		// With oldestPreviousDefaultVersion as cut-off, iterate over all resources and their versions and collect the
		// versions that are older than the cut-off.
		oldVersionsSet := map[openapi.ApiVersion]struct{}{}
		for _, versions := range resourceVersions {
			for _, version := range versions {
				if _, ok := removedFromModule[version]; !ok && openapi.CompareApiVersions(version, oldestPreviousDefaultVersion) < 0 {
					oldVersionsSet[version] = struct{}{}
				}
			}
		}

		// If there are any old versions, sort them and add them to the result.
		if len(oldVersionsSet) > 0 {
			oldVersions := util.UnsortedKeys(oldVersionsSet)
			slices.SortStableFunc(oldVersions, openapi.CompareApiVersions)
			result[moduleName] = oldVersions
		}
	}
	return result
}

func LoadVersionMetadata(rootDir string, modules openapi.AzureModules, majorVersion int) (VersionMetadata, error) {
	versionSources, err := ReadVersionSources(rootDir, modules, majorVersion)
	if err != nil {
		return VersionMetadata{}, err
	}
	return calculateVersionMetadata(versionSources)
}

func calculateVersionMetadata(versionSources VersionSources) (VersionMetadata, error) {
	indexedProviderList := versionSources.ProviderList.Index()
	allResourcesByVersionWithoutDeprecations := RemoveDeprecations(versionSources.AllResourcesByVersion, versionSources.RemovedVersions)

	spec := versionSources.Spec

	config := versionSources.Config
	spec = BuildSpec(allResourcesByVersionWithoutDeprecations, config, spec, indexedProviderList)

	violations := ValidateDefaultConfig(spec, config)

	v2Lock, err := DefaultVersionsFromConfig(allResourcesByVersionWithoutDeprecations, spec)
	if err != nil {
		// Format updated spec to YAML and print for context
		specYaml, yamlErr := yaml.Marshal(spec)
		if yamlErr != nil {
			log.Printf("error marshalling spec to YAML: %v", err)
		}
		wrapped := fmt.Errorf("generating default version lock from spec\n%s\n%w", string(specYaml), err)
		return VersionMetadata{}, wrapped
	}
	inactiveVersions := FindInactiveDefaultVersions(v2Lock, indexedProviderList)

	return VersionMetadata{
		VersionSources:                versionSources,
		AllResourceVersionsByResource: FormatResourceVersions(versionSources.AllResourcesByVersion),
		Pending:                       FindNewerVersions(versionSources.AllResourcesByVersion, v2Lock),
		Spec:                          spec,
		DefaultVersions:               v2Lock,
		CurationViolations:            violations,
		InactiveDefaultVersions:       inactiveVersions,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) ([]string, error) {
	filePrefix := fmt.Sprintf("v%d", v.MajorVersion)
	specPath := filePrefix + "-spec.yaml"
	lockPath := filePrefix + ".yaml"
	return gen.EmitFiles(outputDir, gen.FileMap{
		specPath: v.Spec,
		lockPath: v.DefaultVersions,
	})
}

type VersionSources struct {
	MajorVersion              int
	ProviderList              providerlist.ProviderList
	requiredExplicitResources []string
	// map[ModuleName]map[DefinitionName]ApiVersion
	PreviousDefaultVersions openapi.DefaultVersions
	RemovedVersions         openapi.ModuleVersionList
	Spec                    Spec
	Config                  Curations
	ConfigPath              string
	// Module->version->[]resource
	AllResourcesByVersion ModuleVersionResources
	// map[TokenToRemove]TokenReplacedWith
	ResourcesToRemove     ResourceRemovals
	RemovedInvokes        ResourceRemovals
	NextResourcesToRemove ResourceRemovals
	// This is not required and will be empty if the file does not exist.
	PreviousTokenPaths map[string]string
}

func ReadVersionSources(rootDir string, modules openapi.AzureModules, majorVersion int) (VersionSources, error) {
	providerList, err := providerlist.ReadProviderList(filepath.Join(rootDir, "versions", "az-provider-list.json"))
	if err != nil {
		return VersionSources{}, err
	}

	knownExplicitResources, err := ReadRequiredExplicitResources(filepath.Join(rootDir, "versions", "required-explicit-resources.txt"))
	if err != nil {
		return VersionSources{}, err
	}

	previousPrefix := fmt.Sprintf("v%d", majorVersion-1)
	previousLockPath := filepath.Join(rootDir, "versions", previousPrefix+".yaml")
	previousLock, err := openapi.ReadDefaultVersions(previousLockPath)
	if err != nil {
		return VersionSources{}, err
	}

	filePrefix := fmt.Sprintf("v%d-", majorVersion)
	removed, err := ReadDeprecations(filepath.Join(rootDir, "versions", filePrefix+"removed.json"))
	if err != nil {
		return VersionSources{}, err
	}

	spec, err := ReadSpec(filepath.Join(rootDir, "versions", filePrefix+"spec.yaml"))
	if err != nil {
		return VersionSources{}, err
	}

	configPath := filepath.Join(rootDir, "versions", filePrefix+"config.yaml")
	config, err := ReadManualCurations(configPath)
	if err != nil {
		return VersionSources{}, err
	}

	resourcesToRemovePath := filepath.Join(rootDir, "versions", filePrefix+"removed-resources.json")
	resourcesToRemove, err := ReadResourceRemovals(resourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", resourcesToRemovePath, err)
	}

	removedInvokesPath := filepath.Join(rootDir, "versions", filePrefix+"removed-invokes.yaml")
	removedInvokes, err := ReadResourceRemovals(removedInvokesPath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", removedInvokesPath, err)
	}

	nextVersionFilePrefix := fmt.Sprintf("v%d-", majorVersion+1)
	nextResourcesToRemovePath := filepath.Join(rootDir, "versions", nextVersionFilePrefix+"removed-resources.json")
	nextResourcesToRemove, err := ReadResourceRemovals(nextResourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", nextResourcesToRemovePath, err)
	}

	previousTokenPathsPath := filepath.Join(rootDir, "versions", previousPrefix+"-token-paths.json")
	previousTokenPaths, err := ReadTokenPaths(previousTokenPathsPath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", previousTokenPathsPath, err)
	}

	return VersionSources{
		MajorVersion:              majorVersion,
		ProviderList:              *providerList,
		requiredExplicitResources: knownExplicitResources,
		PreviousDefaultVersions:   previousLock,
		RemovedVersions:           removed,
		Spec:                      spec,
		Config:                    config,
		ConfigPath:                configPath,
		AllResourcesByVersion:     FindAllResources(modules),
		ResourcesToRemove:         resourcesToRemove,
		RemovedInvokes:            removedInvokes,
		NextResourcesToRemove:     nextResourcesToRemove,
		PreviousTokenPaths:        previousTokenPaths,
	}, nil
}

// ReadRequiredExplicitResources reads a list of resource tokens which map to an equivalent resource which can be migrated to.
type ResourceRemovals map[string]string

func ReadResourceRemovals(path string) (ResourceRemovals, error) {
	resourceRemovals, err := ReadSource[ResourceRemovals](path)
	if err != nil {
		return nil, err
	}
	return *resourceRemovals, nil
}

func ReadTokenPaths(path string) (map[string]string, error) {
	tokenPaths, err := ReadSource[map[string]string](path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, err
	}
	return *tokenPaths, nil
}

func ReadSource[T any](path string) (*T, error) {
	isYaml := filepath.Ext(path) == ".yaml"
	sourceFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()

	byteValue, err := io.ReadAll(sourceFile)
	if err != nil {
		return nil, err
	}

	var result T
	if isYaml {
		err = yaml.Unmarshal(byteValue, &result)
	} else {
		err = json.Unmarshal(byteValue, &result)
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
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
	bytes, err := io.ReadAll(txtFile)
	if err != nil {
		return nil, err
	}

	// Split on new line
	lines := strings.Split(string(bytes), "\r")
	return lines, nil
}

func FindRemovedInvokesFromResources(modules openapi.AzureModules, removedResources openapi.RemovableResources) openapi.RemovableResources {
	removableInvokes := openapi.RemovableResources{}
	for moduleName, versions := range modules {
		for version, resources := range versions {
			removedResourcePaths := []string{}
			for resourceName, resource := range resources.Resources {
				if removedResources.CanBeRemoved(moduleName, resourceName, string(version)) {
					removedResourcePaths = append(removedResourcePaths, paths.NormalizePath(resource.Path))
					continue
				}
			}
			for invokeName, invoke := range resources.Invokes {
				fullyQualifiedName := openapi.ToFullyQualifiedName(moduleName, invokeName, string(version))
				// Check if the "resource" removal is actually an invoke.
				if removedResources.CanBeRemoved(moduleName, invokeName, string(version)) {
					removableInvokes[fullyQualifiedName] = ""
					continue
				}
				invokePath := paths.NormalizePath(invoke.Path)
				// Try to match on the path - if the invoke sits below a.
				found := false
				for _, resourcePath := range removedResourcePaths {
					if strings.HasPrefix(invokePath, resourcePath) {
						found = true
						break
					}
				}
				if found {
					removableInvokes[fullyQualifiedName] = ""
					continue
				}
			}
		}
	}
	return removableInvokes
}
