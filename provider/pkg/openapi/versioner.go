// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
			normalizedPath := normalizePath(r.Path)
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

func (s RemovableResources) canBeRemoved(azureProvider, resource, version string) bool {
	fqn := ToFullyQualifiedName(azureProvider, resource, version)
	_, ok := s[fqn]
	return ok
}

func RemoveResources(providers AzureProviders, removable RemovableResources) AzureProviders {
	result := AzureProviders{}
	removedCount := 0
	for provider, versions := range providers {
		newVersions := ProviderVersions{}
		for version, resources := range versions {
			filteredResources := VersionResources{
				Resources: make(map[string]*ResourceSpec),
				Invokes:   resources.Invokes,
			}
			for resourceName, resource := range resources.Resources {
				if removable.canBeRemoved(provider, resourceName, version) {
					removedCount++
					continue
				}
				filteredResources.Resources[resourceName] = resource
			}
			newVersions[version] = filteredResources
		}
		result[provider] = newVersions
	}
	log.Printf("Removed %d resources from the spec", removedCount)
	return result
}
