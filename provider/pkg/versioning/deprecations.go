package versioning

import (
	"log"
	"strings"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

func FindDeprecations(specVersions ProvidersVersionResources, defaultVersion openapi.DefaultVersionLock) openapi.ProviderVersionList {
	deprecationCutoff := "2021-01-01"
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
		versionsToRemove := codegen.NewStringSet(deprecations[providerName]...)
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

func SqueezeSpec(versions ProvidersVersionResources, squeeze Squeeze) ProvidersVersionResources {
	if squeeze == nil {
		return versions
	}

	filteredSpec := ProvidersVersionResources{}
	for providerName, versionResources := range versions {
		lowerProvider := strings.ToLower(providerName)
		filteredVersionResources := VersionResources{}
		for apiVersion, resourceNames := range versionResources {
			sdkVersion := openapi.ApiToSdkVersion(apiVersion)
			filteredResourceNames := make([]string, 0, len(resourceNames))
			for _, resourceName := range resourceNames {
				// construct fully qualified name like azure-native:aad/v20210301:DomainService
				fqn := "azure-native:" + lowerProvider + "/" + sdkVersion + ":" + resourceName
				if _, ok := squeeze[fqn]; !ok {
					filteredResourceNames = append(filteredResourceNames, resourceName)
				} else {
					log.Printf("Removed %s from the spec based on squeezing", fqn)
				}
			}
			filteredVersionResources[apiVersion] = filteredResourceNames
		}
		filteredSpec[providerName] = filteredVersionResources
	}
	return filteredSpec
}
