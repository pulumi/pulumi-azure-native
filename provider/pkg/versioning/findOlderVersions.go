package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func findOlderVersions(specVersions ModuleVersionResources, defaultVersions openapi.DefaultVersions) openapi.ModuleVersionList {
	olderModuleVersions := openapi.ModuleVersionList{}
	for moduleName, versions := range specVersions {
		olderVersions := collections.NewOrderableSet[openapi.ApiVersion]()
		defaultResourceVersions := defaultVersions[moduleName]
		minCuratedVersion := findMinDefaultVersion(defaultResourceVersions)
		for version := range versions {
			if version == "" || version >= minCuratedVersion {
				continue
			}
			olderVersions.Add(version)
		}
		olderModuleVersions[moduleName] = olderVersions.SortedValues()
	}
	return olderModuleVersions
}

func findMinDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.DefinitionVersion) openapi.ApiVersion {
	// TODO: Consider using a pointer instead of empty string for when there is no version
	minVersion := openapi.ApiVersion("")
	for _, version := range resourceVersions {
		if minVersion == "" || version.ApiVersion < minVersion {
			minVersion = version.ApiVersion
		}
	}
	return minVersion
}
