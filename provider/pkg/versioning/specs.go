package versioning

import (
	"slices"
	"sort"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

type VersionResources = map[openapi.ApiVersion][]openapi.ResourceName

// Module->version->[]resource
type ModuleVersionResources = map[openapi.ModuleName]VersionResources

func FindAllResources(moduleVersions openapi.AzureModules) ModuleVersionResources {
	formatted := ModuleVersionResources{}
	for moduleName, versions := range moduleVersions {
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
		formatted[moduleName] = versionResources
	}
	return formatted
}

type ResourceVersions = map[openapi.DefinitionName][]openapi.ApiVersion
type ModuleResourceVersions = map[openapi.ModuleName]ResourceVersions

// FormatResourceVersions flips the hierarchy from Version->Resources to Resource->Versions
func FormatResourceVersions(moduleVersions ModuleVersionResources) ModuleResourceVersions {
	formatted := ModuleResourceVersions{}
	for moduleName, versions := range moduleVersions {
		resourceVersions := ResourceVersions{}

		for version, resources := range versions {
			for _, resourceName := range resources {
				resourceVersions[resourceName] = append(resourceVersions[resourceName], version)
			}
		}
		for _, apiVersions := range resourceVersions {
			slices.Sort(apiVersions)
		}
		formatted[moduleName] = resourceVersions
	}
	return formatted
}
