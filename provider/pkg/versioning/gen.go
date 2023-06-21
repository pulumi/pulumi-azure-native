package versioning

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	AllResourcesByVersion         ProvidersVersionResources
	AllResourceVersionsByResource ProviderResourceVersions
	Active                        providerlist.ProviderPathVersionsJson
	Pending                       openapi.ProviderVersionList
	V2Spec                        Spec
	V2Lock                        openapi.DefaultVersionLock
	V2RemovedInvokes              ResourceRemovals
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

	allResourcesByVersionWithoutDeprecations := RemoveDeprecations(allResourcesByVersion, versionSources.V2Removed)

	v2Spec := versionSources.v2Spec

	v2Config := versionSources.v2Config
	v2Spec = BuildSpec(allResourcesByVersionWithoutDeprecations, v2Config, v2Spec)

	violations := ValidateDefaultConfig(v2Spec, v2Config)
	PrintViolationsAsWarnings(versionSources.v2ConfigPath, violations)

	v2Lock, err := DefaultConfigToDefaultVersionLock(allResourcesByVersionWithoutDeprecations, v2Spec)
	if err != nil {
		return VersionMetadata{}, err
	}

	// provider->resource->[]version
	allResourceVersionsByResource := FormatResourceVersions(allResourcesByVersion)

	removedInvokes := ResourceRemovals(findRemovedInvokesFromResources(providers, openapi.RemovableResources(versionSources.v2ResourcesToRemove)))

	return VersionMetadata{
		AllResourcesByVersion:         allResourcesByVersion,
		AllResourceVersionsByResource: allResourceVersionsByResource,
		Active:                        activePathVersionsJson,
		Pending:                       FindNewerVersions(allResourcesByVersion, v2Lock),
		V2Spec:                        v2Spec,
		V2Lock:                        v2Lock,
		V2RemovedInvokes:              removedInvokes,
	}, nil
}

func (v VersionMetadata) WriteTo(outputDir string) error {
	return gen.EmitFiles(outputDir, gen.FileMap{
		"allResourcesByVersion.json":         v.AllResourcesByVersion,
		"allResourceVersionsByResource.json": v.AllResourceVersionsByResource,
		"active.json":                        v.Active,
		"pending.json":                       v.Pending,
		"v2-spec.yaml":                       v.V2Spec,
		"v2-lock.json":                       v.V2Lock,
		"v2-removed-invokes.yaml":            v.V2RemovedInvokes,
	})
}

type VersionSources struct {
	activePathVersions        providerlist.ProviderPathVersions
	requiredExplicitResources []string
	v1Lock                    openapi.DefaultVersionLock
	V2Removed                 openapi.ProviderVersionList
	v2Spec                    Spec
	v2Config                  Curations
	v2ConfigPath              string
	v2ResourcesToRemove       ResourceRemovals
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

	v1LockPath := path.Join(rootDir, "versions", "v1-lock.json")
	v1Lock, err := openapi.ReadDefaultVersionLock(v1LockPath)
	if err != nil {
		return VersionSources{}, err
	}

	v2Removed, err := ReadDeprecations(path.Join(rootDir, "versions", "v2-removed.json"))
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

	v2ResourcesToRemovePath := path.Join(rootDir, "versions", "v2-removed-resources.yaml")
	v2ResourcesToRemove, err := ReadResourceRemovals(v2ResourcesToRemovePath)
	if err != nil {
		return VersionSources{}, fmt.Errorf("could not read v2-removed-resources: %v", err)
	}

	return VersionSources{
		activePathVersions:        activePathVersions,
		requiredExplicitResources: knownExplicitResources,
		v1Lock:                    v1Lock,
		V2Removed:                 v2Removed,
		v2Spec:                    v2Spec,
		v2Config:                  v2Config,
		v2ConfigPath:              v2ConfigPath,
		v2ResourcesToRemove:       v2ResourcesToRemove,
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

func findRemovedInvokesFromResources(providers openapi.AzureProviders, removedResources openapi.RemovableResources) openapi.RemovableResources {
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
