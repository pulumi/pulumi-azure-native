package versioning

import (
	"slices"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
)

type VersionResources = map[openapi.ApiVersion]map[openapi.DefinitionName]openapi.DefinitionVersion

// Module->version->[]resource
type ModuleVersionResources = map[openapi.ModuleName]VersionResources

func FindAllResources(moduleVersions openapi.AzureModules) ModuleVersionResources {
	formatted := ModuleVersionResources{}
	for moduleName, versions := range moduleVersions {
		versionResources := VersionResources{}
		for version, resources := range versions {
			specResources := map[openapi.DefinitionName]openapi.DefinitionVersion{}
			definitions := resources.All()
			for _, definitionName := range util.SortedKeys(definitions) {
				spec := definitions[definitionName]
				definitionType := openapi.DefinitionTypeResource
				if _, isInvoke := resources.Invokes[definitionName]; isInvoke {
					definitionType = openapi.DefinitionTypeInvoke
				}
				specResources[definitionName] = openapi.DefinitionVersion{
					Type:         definitionType,
					Name:         definitionName,
					ApiVersion:   spec.ApiVersion,
					SpecFilePath: spec.FileLocation,
					ResourceUri:  spec.Path,
					RpNamespace:  spec.ModuleNaming.RpNamespace,
				}
			}
			if len(specResources) > 0 {
				versionResources[version] = specResources
			}
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
			for _, definitionName := range util.SortedKeys(resources) {
				resourceVersions[definitionName] = append(resourceVersions[definitionName], version)
			}
		}
		for _, apiVersions := range resourceVersions {
			slices.Sort(apiVersions)
		}
		formatted[moduleName] = resourceVersions
	}
	return formatted
}
