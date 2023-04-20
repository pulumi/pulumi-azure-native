package versioning

import (
	"sort"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

func findOlderVersions(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		//goland:noinspection GoPreferNilSlice
		olderVersions := []openapi.ApiVersion{}
		defaultResourceVersions := defaultVersion[providerName]
		minCuratedVersion := findMinDefaultVersion(defaultResourceVersions)
		for version := range versions {
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

func findMinDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	minVersion := ""
	for _, version := range resourceVersions {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
