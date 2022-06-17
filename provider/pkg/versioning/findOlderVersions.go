package versioning

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"sort"
)

func findOlderVersions(specVersions SpecVersions, curatedVersion openapi.CuratedVersion) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		//goland:noinspection GoPreferNilSlice
		olderVersions := []openapi.ApiVersion{}
		curated := curatedVersion[providerName]
		minCuratedVersion := findMinDefaultVersion(curated)
		for version, _ := range versions {
			if version == "" || version >= minCuratedVersion {
				continue
			}
			olderVersions = append(olderVersions, version)
		}
		sort.Strings(olderVersions)
		olderProviderVersions[providerName] = olderVersions
	}
	return olderProviderVersions
}

func findMinDefaultVersion(versionResources map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	minVersion := ""
	for _, version := range versionResources {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
