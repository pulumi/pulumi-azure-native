package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func findOlderVersions(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		olderVersions := collections.NewOrderableSet[openapi.ApiVersion]()
		defaultResourceVersions := defaultVersion[providerName]
		minCuratedVersion := findMinDefaultVersion(defaultResourceVersions)
		for version := range versions {
			if version == "" || version >= minCuratedVersion {
				continue
			}
			olderVersions.Add(version)
		}
		olderProviderVersions[providerName] = olderVersions.SortedValues()
	}
	return olderProviderVersions
}

func findMinDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	// TODO: Consider using a pointer instead of empty string for when there is no version
	minVersion := openapi.ApiVersion("")
	for _, version := range resourceVersions {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
