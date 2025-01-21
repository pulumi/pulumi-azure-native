// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/segmentio/encoding/json"
)

func ReadModuleVersionList(path string) (ModuleVersionList, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion ModuleVersionList
	err = json.Unmarshal(byteValue, &curatedVersion)
	if err != nil {
		return nil, err
	}

	return curatedVersion, nil
}

func ReadDefaultVersionLock(path string) (DefaultVersionLock, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion DefaultVersionLock
	err = json.Unmarshal(byteValue, &curatedVersion)
	if err != nil {
		return nil, err
	}

	return curatedVersion, nil
}

// calculatePathVersions builds a map of all versions defined for an API paths from a map of all versions of a module.
// The result is a map from a normalized path to a set of versions for that path.
func calculatePathVersions(versionMap ModuleVersions) map[string]*collections.OrderableSet[SdkVersion] {
	pathVersions := map[string]*collections.OrderableSet[SdkVersion]{}
	for version, items := range versionMap {
		for _, r := range items.Resources {
			normalizedPath := paths.NormalizePath(r.Path)
			versions, ok := pathVersions[normalizedPath]
			if !ok {
				versions = collections.NewOrderableSet[SdkVersion]()
				pathVersions[normalizedPath] = versions
			}
			versions.Add(version)
		}
	}
	return pathVersions
}

// 2022-02-02-preview -> v20220202preview
func ApiToSdkVersion(apiVersion ApiVersion) SdkVersion {
	return SdkVersion("v" + strings.ReplaceAll(string(apiVersion), "-", ""))
}

// v20220202preview -> 2022-02-02-preview
func SdkToApiVersion(v SdkVersion) (ApiVersion, error) {
	if !strings.HasPrefix(string(v), "v") || len(v) < len("v20220202") || len(v) > len("v20220202privatepreview") {
		return "", fmt.Errorf("invalid sdk version: %s", v)
	}
	res := v[1:5] + "-" + v[5:7] + "-" + v[7:9]
	suffix := v[9:]
	switch suffix {
	case "preview", "privatepreview", "beta":
		res += "-" + suffix
	case "":
	default:
		return "", fmt.Errorf("invalid sdk version suffix: %s", v)
	}
	return ApiVersion(res), nil
}

// RemovableResources represents removable resources mapped to the resource that can replace them since the two are
// schema-compatible. Both are represented as fully qualified names like azure-native:azuread/v20210301:DomainService.
type RemovableResources map[string]string

// Returns azure-native:azureModule/version:resource.
// Version can be either ApiVersion or SdkVersion.
// TODO tkappler version should be optional
func ToFullyQualifiedName(moduleName ModuleName, resource, version string) string {
	// construct fully qualified name like azure-native:aad/v20210301:DomainService
	const fqnFmt = "azure-native:%s/%s:%s"
	if !strings.HasPrefix(version, "v") {
		version = string(ApiToSdkVersion(ApiVersion(version)))
	}
	return fmt.Sprintf(fqnFmt, moduleName.Lowered(), version, resource)
}

// Version can be either ApiVersion or SdkVersion
func (s RemovableResources) CanBeRemoved(moduleName ModuleName, resource, version string) bool {
	fqn := ToFullyQualifiedName(moduleName, resource, version)
	_, ok := s[fqn]
	return ok
}

func RemoveResources(modules AzureModules, removable RemovableResources) AzureModules {
	result := AzureModules{}
	removedResourceCount := 0
	removedInvokeCount := 0
	for moduleName, versions := range modules {
		newVersions := ModuleVersions{}
		for version, resources := range versions {
			filteredResources := NewVersionResources()
			removedResourcePaths := []string{}
			for resourceName, resource := range resources.Resources {
				if removable.CanBeRemoved(moduleName, resourceName, string(version)) {
					removedResourceCount++
					removedResourcePaths = append(removedResourcePaths, paths.NormalizePath(resource.Path))
					continue
				}
				filteredResources.Resources[resourceName] = resource
			}
			for invokeName, invoke := range resources.Invokes {
				if removable.CanBeRemoved(moduleName, invokeName, string(version)) {
					removedInvokeCount++
					continue
				}
				invokePath := paths.NormalizePath(invoke.Path)
				// If we can't match on name, we try to match on the path.
				found := false
				for _, resourcePath := range removedResourcePaths {
					if strings.HasPrefix(invokePath, resourcePath) {
						found = true
						break
					}
				}
				if found {
					removedInvokeCount++
					continue
				}
				filteredResources.Invokes[invokeName] = invoke
			}
			// If there are no resources left, we can remove the version entirely.
			if version != "" && len(filteredResources.Resources) == 0 && len(filteredResources.Invokes) > 0 {
				removedInvokeCount += len(filteredResources.Invokes)
				for invokeName := range filteredResources.Invokes {
					log.Printf("Removable invoke: azure-native:%s/%s:%s", moduleName.Lowered(), version, invokeName)
				}
				continue
			}
			newVersions[version] = filteredResources
		}
		result[moduleName] = newVersions
	}
	log.Printf("Removed %d resources and %d invokes from the spec", removedResourceCount, removedInvokeCount)
	return result
}
