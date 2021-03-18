// Copyright 2016-2020, Pulumi Corporation.

package openapi

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// AzureProviders maps provider names (e.g. Compute) to versions in that providers and resources therein.
type AzureProviders = map[string]ProviderVersions

// ProviderVersions maps API Versions (e.g. 2020-08-01) to resources and invokes in that version.
type ProviderVersions = map[string]VersionResources

// VersionResources contains all resources and invokes in a given API version.
type VersionResources struct {
	Resources map[string]*ResourceSpec
	Invokes   map[string]*ResourceSpec
}

// ResourceSpec contains a pointer in an Open API Spec that defines a resource and related metadata.
type ResourceSpec struct {
	Path               string
	PathItem           *spec.PathItem
	PathItemList       *spec.PathItem
	Swagger            *Spec
	CompatibleVersions []string
	DefaultBody        map[string]interface{}
}

// AllVersions finds all Azure Open API specs on disk, parses them, and creates in-memory representation of resources,
// collected per Azure Provider and API Version - for all API versions.
func AllVersions() AzureProviders {
	swaggerSpecLocations, err := swaggerLocations()
	if err != nil {
		panic(err)
	}

	// Collect all versions for each path in the API across all Swagger files.
	providers := AzureProviders{}
	for _, location := range swaggerSpecLocations {
		swagger, err := NewSpec(location)
		if err != nil {
			panic(errors.Wrapf(err, "failed to parse %q", location))
		}

		for path := range swagger.Paths.Paths {
			addAPIPath(providers, location, path, swagger)
		}
	}

	versionChecker, err := newVersioner()
	if err != nil {
		panic(errors.Wrapf(err, "reading provider versions"))
	}

	for providerName, versionMap := range providers {
		// Add a `latest` (stable) version for each resource and invoke.
		latestResources := versionChecker.calculateLatestVersions(providerName, versionMap, false /* invokes */, false /* preview */)
		latestInvokes := versionChecker.calculateLatestVersions(providerName, versionMap, true /* invokes */, false /* preview */)

		// Add a default version for each resource and invoke.
		defaultResources := versionChecker.calculateLatestVersions(providerName, versionMap, false /* invokes */, true /* preview */)
		defaultInvokes := versionChecker.calculateLatestVersions(providerName, versionMap, true /* invokes */, true /* preview */)

		versionMap["latest"] = VersionResources{
			Resources: latestResources,
			Invokes:   latestInvokes,
		}
		versionMap[""] = VersionResources{
			Resources: defaultResources,
			Invokes:   defaultInvokes,
		}

		// Set compatible versions to all other versions of the resource with the same normalized API path.
		pathVersions := versionChecker.calculatePathVersions(versionMap)
		for version, items := range versionMap {
			for _, r := range items.Resources {
				var otherVersions []string
				for _, otherVersion := range pathVersions[normalizePath(r.Path)].SortedValues() {
					if otherVersion != version {
						otherVersions = append(otherVersions, otherVersion)
					}
				}
				r.CompatibleVersions = otherVersions
			}
		}
	}

	return providers
}

// SingleVersion returns only a single (latest or preview) version of each resource from the full list of resource
// versions.
func SingleVersion(providers AzureProviders) AzureProviders {
	singleVersion := AzureProviders{}

	for providerName, allVersionMap := range providers {
		singleVersion[providerName] = ProviderVersions{ "": allVersionMap[""] }
	}

	return singleVersion
}

// IsPreview returns true for API versions that aren't considered stable.
func IsPreview(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "preview") || strings.Contains(lower, "beta")
}

// swaggerLocations returns a slice of URLs of all known Azure Resource Manager swagger files.
func swaggerLocations() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pattern := filepath.Join(dir, "/azure-rest-api-specs/specification/**/resource-manager/Microsoft.*/**/20*/*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	sort.Strings(files)

	return files, nil
}

var excludeRegexPatterns = []string{
	// This preview version defines two types with the same name (one enum, one object) which fails to pass our codegen.
	// It's old, preview, and not important - so exclude the files of this version.
	".*frontdoor/resource-manager/Microsoft.Network/preview/2018-08-01-preview.*",
}
var excludeRegexes []*regexp.Regexp
func init() {
	for _, pattern := range excludeRegexPatterns {
		excludeRegexes = append(excludeRegexes, regexp.MustCompile(pattern))
	}
}

// addAPIPath considers whether an API path contains resources and/or invokes and adds corresponding entries to the
// provider map. `providers` are mutated in-place.
func addAPIPath(providers AzureProviders, fileLocation, path string, swagger *Spec) {
	for _, re := range excludeRegexes {
		if re.MatchString(fileLocation) {
			return
		}
	}

	prov := resources.ResourceProvider(fileLocation, path)
	if prov == "" {
		return
	}

	// Find (or create) the version map with this name.
	versionMap, ok := providers[prov]
	if !ok {
		versionMap = map[string]VersionResources{}
		providers[prov] = versionMap
	}

	// Find (or create) the resource map with this name.
	apiVersion := "v" + strings.ReplaceAll(swagger.Info.Version, "-", "")
	version, ok := versionMap[apiVersion]
	if !ok {
		version = VersionResources{
			Resources: map[string]*ResourceSpec{},
			Invokes:   map[string]*ResourceSpec{},
		}
		versionMap[apiVersion] = version
	}

	pathItem := swagger.Paths.Paths[path]

	if ok := resources.HasCustomDelete(path); ok {
		pathItem.Delete = &spec.Operation{
			OperationProps: spec.OperationProps{
				Description: "Custom implementation of this operation is available",
			},
		}
	}

	pathItemList, hasList := swagger.Paths.Paths[path+"/list"]

	// Add a resource entry.
	if pathItem.Put != nil && !pathItem.Put.Deprecated {
		hasDelete := pathItem.Delete != nil && !pathItem.Delete.Deprecated

		switch {
		case pathItem.Get != nil && !pathItem.Get.Deprecated:
			defaultBody, hasDefault := defaultResourcesStateNormalized[normalizePath(path)]
			if hasDefault && hasDelete && containsNonEmptyCollections(defaultBody) {
				// See the limitation in `SdkShapeConverter.isDefaultResponse()`
				panic(fmt.Sprintf("invalid defaultResourcesState '%s': non-empty collections aren't supported for deletable resources", path))
			}

			typeName := resources.ResourceName(pathItem.Get.ID)
			if typeName != "" && (hasDelete || hasDefault) {
				version.Resources[typeName] = &ResourceSpec{
					Path:        path,
					PathItem:    &pathItem,
					Swagger:     swagger,
					DefaultBody: defaultBody,
				}
			}
		case pathItem.Head != nil && !pathItem.Head.Deprecated:
			typeName := resources.ResourceName(pathItem.Head.ID)
			if typeName != "" && hasDelete {
				version.Resources[typeName] = &ResourceSpec{
					Path:     path,
					PathItem: &pathItem,
					Swagger:  swagger,
				}
			}
		case hasList:
			var typeName string
			switch {
			case pathItemList.Get != nil && !pathItemList.Get.Deprecated:
				typeName = resources.ResourceName(pathItemList.Get.ID)
			case pathItemList.Post != nil && !pathItemList.Post.Deprecated:
				typeName = resources.ResourceName(pathItemList.Post.ID)
			}
			if typeName != "" {
				var defaultBody map[string]interface{}
				if !hasDelete {
					// The /list pattern that we handle here seems to universally have this shape of the default body.
					// Instead of maintaining the resources in defaultResourcesState, we can hard-code it here.
					defaultBody = map[string]interface{}{
						"properties": map[string]interface{}{},
					}
				}
				version.Resources[typeName] = &ResourceSpec{
					Path:         path,
					PathItem:     &pathItem,
					PathItemList: &pathItemList,
					Swagger:      swagger,
					DefaultBody:  defaultBody,
				}
			}
		}
	}

	// Add a POST invoke entry.
	if pathItem.Post != nil && !pathItem.Post.Deprecated {
		parts := strings.Split(path, "/")
		operationName := strings.ToLower(parts[len(parts)-1])
		prefix := ""
		switch {
		case strings.HasPrefix(operationName, "list"):
			prefix = "list"
		case strings.HasPrefix(operationName, "get"):
			prefix = "get"
		default:
			return
		}

		typeName := resources.ResourceName(pathItem.Post.ID)
		if typeName != "" {
			version.Invokes[prefix+typeName] = &ResourceSpec{
				Path:     path,
				PathItem: &pathItem,
				Swagger:  swagger,
			}
		}
	}
}

// Sometimes, Azure resources change the API paths between API versions. Most of the time, we can detect that based
// on operation names. However, in a number of cases, the operation names change at the same time.
// legacyPathMappings provides a manual map to help our codegen discover the old "aliases" of new resource paths
// and group them under the same Pulumi resource, including proper top-level resource calculation and aliasing.
var legacyPathMappings = map[string]string{
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}",
	"/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.CostManagement/connectors/{connectorName}": "/providers/Microsoft.CostManagement/cloudConnectors/{connectorName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}":"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/interfaceEndpoints/{interfaceEndpointName}": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/appliances/{applianceName}": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}",
}
var legacyPathMappingNormalized = map[string]string{}
func init() {
	for key, value := range legacyPathMappings {
		legacyPathMappingNormalized[normalizePath(key)] = normalizePath(value)
	}
}

// normalizePath converts an API path to its canonical form (lowercase, with all placeholders removed). The paths that
// convert to the same canonical path are considered to represent the same resource.
func normalizePath(path string) string {
	lowerPath := strings.ReplaceAll(strings.ToLower(strings.TrimSuffix(path, "/")), "-", "")
	parts := strings.Split(lowerPath, "/")
	newParts := make([]string, len(parts))
	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			newParts[i] = "{}"
		} else {
			newParts[i] = part
		}
	}
	normalized := strings.Join(newParts, "/")
	if override, ok := legacyPathMappingNormalized[normalized]; ok {
		return override
	}
	return normalized
}
