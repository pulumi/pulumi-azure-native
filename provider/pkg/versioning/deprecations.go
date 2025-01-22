package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func FindDeprecations(specVersions ModuleVersionResources, defaultVersions openapi.DefaultVersions) openapi.ModuleVersionList {
	deprecationCutoff := openapi.ApiVersion("2021-01-01")
	olderVersions := findOlderVersions(specVersions, defaultVersions)
	for name, versions := range olderVersions {
		filteredVersions := []openapi.ApiVersion{}
		for _, version := range versions {
			if version < deprecationCutoff {
				filteredVersions = append(filteredVersions, version)
			}
		}
		olderVersions[name] = filteredVersions
	}
	return olderVersions
}

func RemoveDeprecations(versions ModuleVersionResources, deprecations openapi.ModuleVersionList) ModuleVersionResources {
	filteredSpec := ModuleVersionResources{}
	for moduleName, versionResources := range versions {
		filteredVersions := VersionResources{}
		versionsToRemove := collections.NewOrderableSet[openapi.ApiVersion](deprecations[moduleName]...)
		for apiVersion, resourceNames := range versionResources {
			if versionsToRemove.Has(apiVersion) {
				continue
			}
			filteredVersions[apiVersion] = resourceNames
		}
		filteredSpec[moduleName] = filteredVersions
	}
	return filteredSpec
}

func ReadDeprecations(deprecationPath string) (openapi.ModuleVersionList, error) {
	return openapi.ReadModuleVersionList(deprecationPath)
}
