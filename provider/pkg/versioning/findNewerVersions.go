package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func FindNewerVersions(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	olderProviderVersions := openapi.ProviderVersionList{}
	for providerName, versions := range specVersions {
		newerVersions := collections.NewOrderableSet[openapi.ApiVersion]()
		defaultResourceVersions := defaultVersion[providerName]
		maxCuratedVersion := findMaxDefaultVersion(defaultResourceVersions)
		for version := range versions {
			if version == "" || version <= maxCuratedVersion {
				continue
			}
			newerVersions.Add(version)
		}
		olderProviderVersions[providerName] = newerVersions.SortedValues()
	}
	return olderProviderVersions
}

func findMaxDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.ApiVersion) openapi.ApiVersion {
	// We currently use empty string to represent when there is no version available which must be handled above.
	// This might be better being represented as a nil value.
	minVersion := openapi.ApiVersion("")
	for _, version := range resourceVersions {
		if minVersion == "" || version < minVersion {
			minVersion = version
		}
	}
	return minVersion
}
