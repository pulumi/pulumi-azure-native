// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/segmentio/encoding/json"
)

func ReadV1DefaultVersionLock() (DefaultVersionLock, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	v1Path := filepath.Join(dir, "versions", "v1-lock.json")
	return ReadDefaultVersionLock(v1Path)
}

func ReadV2DefaultVersionLock() (DefaultVersionLock, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	v1Path := filepath.Join(dir, "versions", "v2-lock.json")
	return ReadDefaultVersionLock(v1Path)
}

func ReadDeprecated() (ProviderVersionList, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	deprecatedPath := filepath.Join(dir, "versions", "deprecated.json")
	return ReadProviderVersionList(deprecatedPath)
}

func ReadRemoved() (ProviderVersionList, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	removedPath := filepath.Join(dir, "versions", "v1-removed.json")
	return ReadProviderVersionList(removedPath)
}

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

type Squeeze map[string]string

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

func (s Squeeze) canBeUpgraded(azureProvider, resource, version string) bool {
	fqn := ToFullyQualifiedName(azureProvider, resource, version)
	_, ok := s[fqn]
	return ok
}

func (s Squeeze) Preserve(azureProvider, resource, version string) bool {
	fqn := ToFullyQualifiedName(azureProvider, resource, version)
	_, ok := s[fqn]
	if ok {
		delete(s, fqn)
	}
	return ok
}

func SqueezeResources(providers AzureProviders, squeeze Squeeze) AzureProviders {
	result := AzureProviders{}
	squeezeCount := 0
	for provider, versions := range providers {
		newVersions := ProviderVersions{}
		for version, resources := range versions {
			filteredResources := VersionResources{
				Resources: make(map[string]*ResourceSpec),
				Invokes:   resources.Invokes,
			}
			for resourceName, resource := range resources.Resources {
				if squeeze.canBeUpgraded(provider, resourceName, version) {
					// log.Printf("Squeezing %s:%s from the spec", resourceName, version)
					squeezeCount++
					continue
				}
				filteredResources.Resources[resourceName] = resource
			}
			newVersions[version] = filteredResources
		}
		result[provider] = newVersions
	}
	log.Printf("Squeezed %d resources from the spec", squeezeCount)
	return result
}

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
