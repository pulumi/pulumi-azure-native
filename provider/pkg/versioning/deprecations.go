package versioning

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func FindDeprecations(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	deprecationCutoff := openapi.ApiVersion("2021-01-01")
	olderVersions := findOlderVersions(specVersions, defaultVersion)
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

func RemoveDeprecations(versions ProvidersVersionResources, deprecations openapi.ProviderVersionList) ProvidersVersionResources {
	filteredSpec := ProvidersVersionResources{}
	for providerName, versionResources := range versions {
		filteredVersions := VersionResources{}
		versionsToRemove := collections.NewOrderableSet[openapi.ApiVersion](deprecations[providerName]...)
		for apiVersion, resourceNames := range versionResources {
			if versionsToRemove.Has(apiVersion) {
				continue
			}
			filteredVersions[apiVersion] = resourceNames
		}
		filteredSpec[providerName] = filteredVersions
	}
	return filteredSpec
}

func ReadDeprecations(deprecationPath string) (openapi.ProviderVersionList, error) {
	return openapi.ReadProviderVersionList(deprecationPath)
}
