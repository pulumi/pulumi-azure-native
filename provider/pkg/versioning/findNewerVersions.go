package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func FindNewerVersions(specVersions ModuleVersionResources, defaultVersions openapi.DefaultVersions) openapi.ModuleVersionList {
	olderModuleVersions := openapi.ModuleVersionList{}
	for moduleName, versions := range specVersions {
		newerVersions := collections.NewOrderableSet[openapi.ApiVersion]()
		defaultResourceVersions := defaultVersions[moduleName]
		maxCuratedVersion := findMaxDefaultVersion(defaultResourceVersions)
		for version := range versions {
			if version == "" || version <= maxCuratedVersion {
				continue
			}
			newerVersions.Add(version)
		}
		olderModuleVersions[moduleName] = newerVersions.SortedValues()
	}
	return olderModuleVersions
}

func findMaxDefaultVersion(resourceVersions map[openapi.DefinitionName]openapi.DefinitionVersion) openapi.ApiVersion {
	// We currently use empty string to represent when there is no version available which must be handled above.
	// This might be better being represented as a nil value.
	minVersion := openapi.ApiVersion("")
	for _, version := range resourceVersions {
		if minVersion == "" || version.ApiVersion < minVersion {
			minVersion = version.ApiVersion
		}
	}
	return minVersion
}
