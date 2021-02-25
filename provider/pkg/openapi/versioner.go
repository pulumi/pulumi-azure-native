package openapi

import (
	"encoding/json"
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

// A manually-maintained list of stable versions where we want to promote a later preview version to be used for
// the top-level resource. These versions are "known" to be outdated but no newer stable versions were released yet.
// However, there's no formal way to derive this from Open API specs, so we have to maintain this manual map.
var deprecatedProviderVersions = map[string][]string{
	"Authorization": {"v20150701"},
	"Sql": {"v20140401", "v20150501"},
}

// calculateLatestVersions builds a map of latest versions per API paths from a map of all versions of a resource
// provider. The result is a map from a resource type name to resource specs.
func (c *versioner) calculateLatestVersions(provider string, versionMap ProviderVersions, invokes,
	preview bool) (latestResources map[string]*ResourceSpec) {
	deprecatedVersions := codegen.NewStringSet()
	if v, ok := deprecatedProviderVersions[provider]; ok {
		deprecatedVersions = codegen.NewStringSet(v...)
	}

	var stables, previews []string
	for version := range versionMap {
		switch {
		case strings.Contains(version, "private"):
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
