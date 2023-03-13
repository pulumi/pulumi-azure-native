// Copyright 2021, Pulumi Corporation.  All rights reserved.

package openapi

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/segmentio/encoding/json"
)

func ReadV1Version() (CuratedVersion, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	v1Path := filepath.Join(dir, "versions", "v1.json")
	return ReadCuratedVersion(v1Path)
}

func ReadDeprecated() (ProviderVersionList, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	deprecatedPath := filepath.Join(dir, "versions", "deprecated.json")
	return readProviderVersionList(deprecatedPath)
}

func ReadRemoved() (ProviderVersionList, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	removedPath := filepath.Join(dir, "versions", "removed.json")
	return readProviderVersionList(removedPath)
}

func readProviderVersionList(path string) (ProviderVersionList, error) {
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

func ReadCuratedVersion(path string) (CuratedVersion, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion CuratedVersion
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
