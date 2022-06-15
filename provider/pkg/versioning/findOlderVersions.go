package versioning

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"sort"
)

type ProviderName = string
// ApiVersion e.g. 2020-01-30
type ApiVersion = string
type ResourceName = string

type ProviderVersions = map[ProviderName][]ApiVersion

func FindOlderVersions(specVersions SpecVersions, curatedVersion openapi.CuratedVersion) ProviderVersions {
	olderProviderVersions := ProviderVersions{}
	for providerName, versions := range specVersions {
		//goland:noinspection GoPreferNilSlice
		olderVersions := []ApiVersion{}
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

func findMinDefaultVersion(versionResources map[ResourceName]ApiVersion) ApiVersion {
	minVersion := ""
	for _, version := range versionResources {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
