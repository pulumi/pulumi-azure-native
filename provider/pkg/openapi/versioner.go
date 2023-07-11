// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/segmentio/encoding/json"
)

func ReadProviderVersionList(path string) (ProviderVersionList, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion ProviderVersionList
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

// calculatePathVersions builds a map of all versions defined for an API paths from a map of all versions of a resource
// provider. The result is a map from a normalized path to a set of versions for that path.
func calculatePathVersions(versionMap ProviderVersions) map[string]codegen.StringSet {
	pathVersions := map[string]codegen.StringSet{}
	for version, items := range versionMap {
		for _, r := range items.Resources {
			normalizedPath := paths.NormalizePath(r.Path)
			versions, ok := pathVersions[normalizedPath]
			if !ok {
				versions = codegen.StringSet{}
				pathVersions[normalizedPath] = versions
			}
			versions.Add(version)
		}
	}
	return pathVersions
}

func ApiToSdkVersion(apiVersion ApiVersion) SdkVersion {
	return "v" + strings.ReplaceAll(apiVersion, "-", "")
}

// RemovableResources represents removable resources mapped to the resource that can replace them since the two are
// schema-compatible. Both are represented as fully qualified names like azure-native:azuread/v20210301:DomainService.
type RemovableResources map[string]string

// Returns azure-native:azureProvider/version:resource
// TODO tkappler version should be optional
func ToFullyQualifiedName(azureProvider, resource, version string) string {
	// construct fully qualified name like azure-native:aad/v20210301:DomainService
	const fqnFmt = "azure-native:%s/%s:%s"
	if !strings.HasPrefix(version, "v") {
		version = ApiToSdkVersion(version)
	}
	return fmt.Sprintf(fqnFmt, strings.ToLower(azureProvider), version, resource)
}

func (s RemovableResources) CanBeRemoved(azureProvider, resource, version string) bool {
	fqn := ToFullyQualifiedName(azureProvider, resource, version)
	_, ok := s[fqn]
	return ok
}

func RemoveResources(providers AzureProviders, removable RemovableResources) AzureProviders {
	result := AzureProviders{}
	removedResourceCount := 0
	removedInvokeCount := 0
	for provider, versions := range providers {
		newVersions := ProviderVersions{}
		for version, resources := range versions {
			filteredResources := VersionResources{
				Resources: map[string]*ResourceSpec{},
				Invokes:   map[string]*ResourceSpec{},
			}
			removedResourcePaths := []string{}
			for resourceName, resource := range resources.Resources {
				if removable.CanBeRemoved(provider, resourceName, version) {
					removedResourceCount++
					removedResourcePaths = append(removedResourcePaths, paths.NormalizePath(resource.Path))
					continue
				}
				filteredResources.Resources[resourceName] = resource
			}
			for invokeName, invoke := range resources.Invokes {
				if removable.CanBeRemoved(provider, invokeName, version) {
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
					log.Printf("Removable invoke: azure-native:%s/%s:%s", strings.ToLower(provider), version, invokeName)
				}
				continue
			}
			newVersions[version] = filteredResources
		}
		result[provider] = newVersions
	}
	log.Printf("Removed %d resources and %d invokes from the spec", removedResourceCount, removedInvokeCount)
	return result
}
