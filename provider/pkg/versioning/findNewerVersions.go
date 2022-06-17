package versioning

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"sort"
)

func FindNewerVersions(specVersions SpecVersions, curatedVersion openapi.CuratedVersion) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		//goland:noinspection GoPreferNilSlice
		newerVersions := []openapi.ApiVersion{}
		curated := curatedVersion[providerName]
		maxCuratedVersion := findMaxDefaultVersion(curated)
		for version, _ := range versions {
			if version == "" || version <= maxCuratedVersion {
				continue
			}
			newerVersions = append(newerVersions, version)
		}
		sort.Strings(newerVersions)
		olderProviderVersions[providerName] = newerVersions
	}
	return olderProviderVersions
}

func findMaxDefaultVersion(versionResources map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	minVersion := ""
	for _, version := range versionResources {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
