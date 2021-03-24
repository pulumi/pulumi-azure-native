package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// versioner checks whether the given combination of a provider, a resource path, and a versions
// are known in version metadata returned by Azure (retrieved via the Azure CLI, stored in a local JSON file).
type versioner struct {
	lookup map[string]map[string]codegen.StringSet
}

func newVersioner() (*versioner, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(filepath.Join(dir, "/azure-provider-versions/provider_list.json"))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var provs []prov
	err = json.Unmarshal(byteValue, &provs)
	if err != nil {
		return nil, err
	}

	result := map[string]map[string]codegen.StringSet{}

	for _, prov := range provs {
		namespace := strings.ToLower(prov.Namespace)
		if !strings.HasPrefix(namespace, "microsoft.") {
			continue
		}
		providerName := strings.TrimPrefix(namespace, "microsoft.")

		versions := map[string]codegen.StringSet{}
		allVersions := codegen.NewStringSet()
		for _, rt := range prov.ResourceTypes {
			set := codegen.NewStringSet()
			for _, v := range rt.ApiVersions {
				name := "v" + strings.ReplaceAll(v, "-", "")
				allVersions.Add(name)
				set.Add(name)
			}
			versions[strings.ToLower(rt.ResourceType)] = set
		}
		versions[""] = allVersions
		result[providerName] = versions
	}

	return &versioner{lookup: result}, nil
}

// A manually-maintained list of stable versions that we want to promote a later preview version to be used for
// the top-level resource. These versions are "known" to be outdated but no newer stable versions were released yet.
// However, there's no formal way to derive this from Open API specs, so we have to maintain this manual map.
var deprecatedProviderVersions = map[string][]string{
	"Authorization": {"v20150701"},
	"Sql": {"v20140401", "v20150501"},
}

// A manually-maintained list of API versions to ignore while calculating the top-level resources.
var ignoredProviderVersions = map[string][]string{
	// This preview version introduces new resources but also changes the shape of some items, so merging it with
	// the previous stable version isn't easy. Ignore this preview for now. We should be able to remove this entry
	// once the next stable version ships.
	"EventGrid": {"v20201015preview"},
	"StorSimple": {"v20161001"},
	// Not supported yet (2020-03-23)? Check back later.
	"ApiManagement": {"v20201201"},
}

// A manually-maintained list of versions where we want to use for the top-level resource. The primary goal is to
// avoid breaking changes within a single major version of the provider that could come with new API versions.
// We reset this map every time we release a new major version (or a new 0.x minor version).
// Currently populated for 0.7.* series.
var lockedTypeVersions = map[string]string{
}

// A manually-maintained list of resources that were deprecated and have no direct successor. They "pollute"
// the top-level resources, including bringing old incompatible types, so we'd rather exclude them. Note that
// they aren't officially deprecated in the APIs. We want to keep this list as small as possible - mostly to prevent
// type clashing problems.
var deprecatedResources = codegen.NewStringSet(
	"apimanagement:TenantPolicy",
	"consumption:BudgetByResourceGroupName",
	"containerregistry:BuildStep",
	"containerservice:ContainerService",
	"costmanagement:Budget",
	"costmanagement:ReportConfig",
	"costmanagement:ReportConfigByResourceGroupName",
	"datamigration:ServiceTask",
	"synapse:SqlDatabase",
	"web:CertificateCsr",
	"web:SiteInstanceDeployment",
	"web:SiteInstanceDeploymentSlot",
)

// calculateLatestVersions builds a map of latest versions per API paths from a map of all versions of a resource
// provider. The result is a map from a resource type name to resource specs.
func (c *versioner) calculateLatestVersions(provider string, versionMap ProviderVersions, invokes,
	preview bool) (latestResources map[string]*ResourceSpec) {
	deprecatedVersions := codegen.NewStringSet()
	if v, ok := deprecatedProviderVersions[provider]; ok {
		deprecatedVersions = codegen.NewStringSet(v...)
	}
	ignoredVersions := codegen.NewStringSet()
	if v, ok := ignoredProviderVersions[provider]; ok {
		ignoredVersions = codegen.NewStringSet(v...)
	}

	var stables, previews []string
	for version := range versionMap {
		switch {
		case strings.Contains(version, "private"):
			// skip
		case ignoredVersions.Has(version):
			// skip
		case deprecatedVersions.Has(version):
			// Treat deprecated versions as preview for sorting purpose.
			previews = append(previews, version)
		case !IsPreview(version):
			stables = append(stables, version)
		case preview:
			previews = append(previews, version)
		}
	}
	// Sort the versions from earliest to latest previews, then from earliest to latest stable.
	sort.Strings(previews)
	sort.Strings(stables)
	versions := append(previews, stables...)

	pathTypeNames := map[string]string{}
	latestResources = map[string]*ResourceSpec{}
	knownResources := codegen.NewStringSet()
	for _, version := range versions {
		items := versionMap[version]
		res := items.Resources
		if invokes {
			res = items.Invokes
		}

		for typeName, r := range res {
			fullTypeName := fmt.Sprintf("%s:%s", strings.ToLower(provider), typeName)
			if deprecatedResources.Has(fullTypeName) {
				continue
			}
			if lockedVersion, ok := lockedTypeVersions[fullTypeName]; ok && lockedVersion != version {
				// If we have a locked version for this resource, ignore all other versions.
				continue
			}

			normalizedPath := normalizePath(r.Path)
			isKnown := c.isKnown(provider, r.Path, version)
			if !isKnown && knownResources.Has(normalizedPath) {
				continue
			}
			if isKnown {
				knownResources.Add(normalizedPath)
			}

			previousTypeName := pathTypeNames[normalizedPath]
			if previousTypeName != "" && previousTypeName != typeName {
				delete(latestResources, previousTypeName)
			}

			pathTypeNames[normalizedPath] = typeName
			copyResource := *r
			latestResources[typeName] = &copyResource
		}
	}
	return latestResources
}

// isKnown returns true when a given provider, a resource path, and a version exist in the lookup map.
func (c *versioner) isKnown(provider, path, version string) bool {
	providerLookup, ok := c.lookup[strings.ToLower(provider)]
	if !ok {
		return false
	}

	resourceTypes := c.armResourceTypes(path)
	for _, resourceType := range resourceTypes {
		if resourceVersions, ok := providerLookup[resourceType]; ok {
			return resourceVersions.Has(version)
		}
	}
	return false
}

// armResourceTypes returns a list of ARM resource type identifiers starting from the most specific one and ending
// with an empty name. We use this helper function to lookup resource types starting from the most specific one
// and then generalizing. We don't know at which level Microsoft defined the versions for this resource, so we should
// check them all until we find one.
// Example of a result: { "servers/databases/transparentDataEncryption", "servers/databases", "servers", "" }.
func (c *versioner) armResourceTypes(path string) []string {
	parts := strings.Split(strings.ToLower(path), "microsoft.")
	resourcePath := parts[len(parts)-1]
	segments := strings.Split(resourcePath, "/")
	var result, builder []string
	for _, segment := range segments[1:] {
		if !strings.HasPrefix(segment, "{") {
			builder = append(builder, segment)
			result = append([]string{strings.Join(builder, "/")}, result...)
		}
	}
	result = append(result, "")
	return result
}

// calculatePathVersions builds a map of all versions defined for an API paths from a map of all versions of a resource
// provider. The result is a map from a normalized path to a set of versions for that path.
func (c *versioner) calculatePathVersions(versionMap ProviderVersions) (pathVersions map[string]codegen.StringSet) {
	pathVersions = map[string]codegen.StringSet{}
	for version, items := range versionMap {
		for _, r := range items.Resources {
			normalizedPath := normalizePath(r.Path)
			versions, ok := pathVersions[normalizedPath]
			if !ok {
				versions = codegen.StringSet{}
				pathVersions[normalizedPath] = versions
			}
			versions.Add(version)
		}
	}
	return pathVersions
}

type prov struct {
	Namespace string `json:"namespace"`
	ResourceTypes []provRes `json:"resourceTypes"`
}

type provRes struct {
	ResourceType string `json:"resourceType"`
	ApiVersions []string `json:"apiVersions"`
}
