package versioning

import (
	"slices"
	"sort"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

type VersionResources = map[openapi.ApiVersion][]openapi.ResourceName

// Provider->version->[]resource
type ProvidersVersionResources = map[openapi.ProviderName]VersionResources

func FindAllResources(providerVersions openapi.AzureProviders) ProvidersVersionResources {
	formatted := ProvidersVersionResources{}
	for providerName, versions := range providerVersions {
		versionResources := VersionResources{}
		for _, resources := range versions {
			version := openapi.ApiVersion("")
			allResources := resources.All()
			specResources := make([]string, 0, len(allResources))
			for resourceName, spec := range allResources {
				specResources = append(specResources, resourceName)
				version = openapi.ApiVersion(spec.Swagger.Info.Version)
			}
			sort.Strings(specResources)
			versionResources[version] = specResources
		}
		formatted[providerName] = versionResources
	}
	return formatted
}

type ResourceVersions = map[openapi.DefinitionName][]openapi.ApiVersion
type ProviderResourceVersions = map[openapi.ProviderName]ResourceVersions

// FormatResourceVersions flips the hierarchy from Version->Resources to Resource->Versions
func FormatResourceVersions(providerVersions ProvidersVersionResources) ProviderResourceVersions {
	formatted := ProviderResourceVersions{}
	for providerName, versions := range providerVersions {
		resourceVersions := ResourceVersions{}

		for version, resources := range versions {
			for _, resourceName := range resources {
				resourceVersions[resourceName] = append(resourceVersions[resourceName], version)
			}
		}
		for _, apiVersions := range resourceVersions {
			slices.Sort(apiVersions)
		}
		formatted[providerName] = resourceVersions
	}
	return formatted
}
