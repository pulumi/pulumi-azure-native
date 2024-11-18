package versioning

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
	"gopkg.in/yaml.v3"
)

type VersionMetadata struct {
	VersionSources
	AllResourcesByVersion         ProvidersVersionResources
	AllResourceVersionsByResource ProviderResourceVersions
	Active                        providerlist.ProviderPathVersionsJson
	Pending                       openapi.ProviderVersionList
	Spec                          Spec
	Lock                          openapi.DefaultVersionLock
	RemovedInvokes                ResourceRemovals
	CurationViolations            []CurationViolation
}

// Ensure our VersionMetadata type implements the gen.Versioning interface
// The compiler will raise an error here if the interface isn't implemented
var _ gen.Versioning = (*VersionMetadata)(nil)

func (v VersionMetadata) ShouldInclude(provider string, version string, typeName, token string) bool {
	// Keep any resources in the default version lock
	if resources, ok := v.Lock[provider]; ok {
		if defaultResourceVersion, ok := resources[typeName]; ok {
			if openapi.ApiToSdkVersion(defaultResourceVersion) == version {
				return true
			}
		}
	}
	// Exclude versions from removed versions
	if versions, ok := v.RemovedVersions[provider]; ok {
		for _, removedVersion := range versions {
			if openapi.ApiToSdkVersion(removedVersion) == version {
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

func (v VersionMetadata) GetAllVersions(provider openapi.ProviderName, resource openapi.ResourceName) []openapi.ApiVersion {
	return v.AllResourceVersionsByResource[provider][resource]
}

func LoadVersionMetadata(rootDir string, providers openapi.AzureProviders, majorVersion int) (VersionMetadata, error) {
	versionSources, err := ReadVersionSources(rootDir, majorVersion)
	if err != nil {
		return VersionMetadata{}, err
	}
	return calculateVersionMetadata(versionSources, providers, majorVersion)
}

func calculateVersionMetadata(versionSources VersionSources, providers openapi.AzureProviders, majorVersion int) (VersionMetadata, error) {
	// map[LoweredProviderName]map[ResourcePath]ApiVersions
	activePathVersions := versionSources.activePathVersions
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	// provider->version->[]resource
	allResourcesByVersion := FindAllResources(providers)

	allResourcesByVersionWithoutDeprecations := RemoveDeprecations(allResourcesByVersion, versionSources.RemovedVersions)

	spec := versionSources.Spec

	config := versionSources.Config
	spec = BuildSpec(allResourcesByVersionWithoutDeprecations, config, spec)

	violations := ValidateDefaultConfig(spec, config)

	v2Lock, err := DefaultConfigToDefaultVersionLock(allResourcesByVersionWithoutDeprecations, spec)
	if err != nil {
		// Format updated spec to YAML and print for context
		specYaml, yamlErr := yaml.Marshal(spec)
		if yamlErr != nil {
			log.Printf("error marshalling spec to YAML: %v", err)
		}
		wrapped := fmt.Errorf("generating default version lock from spec\n%s\n%w", string(specYaml), err)
		return VersionMetadata{}, wrapped
	}

	// provider->resource->[]version
	allResourceVersionsByResource := FormatResourceVersions(allResourcesByVersion)

	removedInvokes := ResourceRemovals(FindRemovedInvokesFromResources(providers, openapi.RemovableResources(versionSources.ResourcesToRemove)))

	return VersionMetadata{
		VersionSources:                versionSources,
		AllResourcesByVersion:         allResourcesByVersion,
		AllResourceVersionsByResource: allResourceVersionsByResource,
		Active:                        activePathVersionsJson,
		Pending:                       FindNewerVersions(allResourcesByVersion, v2Lock),
		Spec:                          spec,
		Lock:                          v2Lock,
		RemovedInvokes:                removedInvokes,
		CurationViolations:            violations,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) ([]string, error) {
	filePrefix := fmt.Sprintf("v%d-", v.MajorVersion)
	specPath := filePrefix + "spec.yaml"
	lockPath := filePrefix + "lock.json"
	removedInvokesPath := filePrefix + "removed-invokes.yaml"
	return gen.EmitFiles(outputDir, gen.FileMap{
		specPath:           v.Spec,
		lockPath:           v.Lock,
		removedInvokesPath: v.RemovedInvokes,
	})
}

type VersionSources struct {
	MajorVersion              int
	activePathVersions        providerlist.ProviderPathVersions
	requiredExplicitResources []string
	PreviousLock              openapi.DefaultVersionLock
	RemovedVersions           openapi.ProviderVersionList
	Spec                      Spec
	Config                    Curations
	ConfigPath                string
	ResourcesToRemove         ResourceRemovals
	NextResourcesToRemove     ResourceRemovals
}

func ReadVersionSources(rootDir string, majorVersion int) (VersionSources, error) {
	activePathVersions, err := providerlist.ReadProviderList(filepath.Join(rootDir, "azure-provider-versions", "provider_list.json"))
	if err != nil {
		return VersionSources{}, err
	}

	knownExplicitResources, err := ReadRequiredExplicitResources(path.Join(rootDir, "versions", "required-explicit-resources.txt"))
	if err != nil {
		return VersionSources{}, err
	}

	previousFilePrefix := fmt.Sprintf("v%d-", majorVersion-1)
	filePrefix := fmt.Sprintf("v%d-", majorVersion)

	previousLockPath := path.Join(rootDir, "versions", previousFilePrefix+"lock.json")
	previousLock, err := openapi.ReadDefaultVersionLock(previousLockPath)
	if err != nil {
		return VersionSources{}, err
	}

	removed, err := ReadDeprecations(path.Join(rootDir, "versions", filePrefix+"removed.json"))
	if err != nil {
		return VersionSources{}, err
	}

	spec, err := ReadSpec(path.Join(rootDir, "versions", filePrefix+"spec.yaml"))
	if err != nil {
		return VersionSources{}, err
	}

	configPath := path.Join(rootDir, "versions", filePrefix+"config.yaml")
	config, err := ReadManualCurations(configPath)
	if err != nil {
		return VersionSources{}, err
	}

	resourcesToRemovePath := path.Join(rootDir, "versions", filePrefix+"removed-resources.json")
	resourcesToRemove, err := ReadResourceRemovals(resourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", resourcesToRemovePath, err)
	}

	nextVersionFilePrefix := fmt.Sprintf("v%d-", majorVersion+1)
	nextResourcesToRemovePath := path.Join(rootDir, "versions", nextVersionFilePrefix+"removed-resources.yaml")
	nextResourcesToRemove, err := ReadResourceRemovals(nextResourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read %s: %v", nextResourcesToRemovePath, err)
	}

	return VersionSources{
		MajorVersion:              majorVersion,
		activePathVersions:        activePathVersions,
		requiredExplicitResources: knownExplicitResources,
		PreviousLock:              previousLock,
		RemovedVersions:           removed,
		Spec:                      spec,
		Config:                    config,
		ConfigPath:                configPath,
		ResourcesToRemove:         resourcesToRemove,
		NextResourcesToRemove:     nextResourcesToRemove,
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

	byteValue, err := io.ReadAll(sourceFile)
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
	bytes, err := io.ReadAll(txtFile)
	if err != nil {
		return nil, err
	}

	// Split on new line
	lines := strings.Split(string(bytes), "\r")
	return lines, nil
}

func FindRemovedInvokesFromResources(providers openapi.AzureProviders, removedResources openapi.RemovableResources) openapi.RemovableResources {
	removableInvokes := openapi.RemovableResources{}
	for provider, versions := range providers {
		for version, resources := range versions {
			removedResourcePaths := []string{}
			for resourceName, resource := range resources.Resources {
				if removedResources.CanBeRemoved(provider, resourceName, version) {
					removedResourcePaths = append(removedResourcePaths, paths.NormalizePath(resource.Path))
					continue
				}
			}
			for invokeName, invoke := range resources.Invokes {
				fullyQualifiedName := openapi.ToFullyQualifiedName(provider, invokeName, version)
				// Check if the "resource" removal is actually an invoke.
				if removedResources.CanBeRemoved(provider, invokeName, version) {
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
