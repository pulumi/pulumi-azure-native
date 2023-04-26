package versioning

import (
	"sort"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
)

func FindNewerVersions(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		//goland:noinspection GoPreferNilSlice
		newerVersions := []openapi.ApiVersion{}
		defaultResourceVersions := defaultVersion[providerName]
		maxCuratedVersion := findMaxDefaultVersion(defaultResourceVersions)
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

func findMaxDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	minVersion := ""
	for _, version := range resourceVersions {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
