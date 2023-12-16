// Copyright 2016-2020, Pulumi Corporation.

package openapi

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/defaults"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// ProviderName e.g. aad
type ProviderName = string

// ApiVersion e.g. 2020-01-30
type ApiVersion = string

// DefinitionName is the name of either an 'invoke' or a resource (e.g. listBuckets or Bucket)
type DefinitionName = string

// ResourceName e.g. Bucket
type ResourceName = string

// InvokeName e.g. listBuckets
type InvokeName = string

// SdkVersion e.g. v20200130
type SdkVersion = string

// AzureProviders maps provider names (e.g. Compute) to versions in that providers and resources therein.
type AzureProviders map[ProviderName]ProviderVersions

// ProviderVersions maps API Versions (e.g. v20200801) to resources and invokes in that version.
type ProviderVersions = map[SdkVersion]VersionResources

type DiscoveryDiagnostics struct {
	// Naming disambiguations
	NamingDisambiguations []resources.NameDisambiguation
}

// VersionResources contains all resources and invokes in a given API version.
type VersionResources struct {
	Resources map[ResourceName]*ResourceSpec
	Invokes   map[InvokeName]*ResourceSpec
}

func NewVersionResources() VersionResources {
	return VersionResources{
		Resources: map[ResourceName]*ResourceSpec{},
		Invokes:   map[InvokeName]*ResourceSpec{},
	}
}

type ProviderVersionList = map[ProviderName][]ApiVersion

// DefaultVersionLock is an amalgamation of multiple API versions
type DefaultVersionLock = map[ProviderName]map[DefinitionName]ApiVersion

func (v VersionResources) All() map[string]*ResourceSpec {
	specs := map[string]*ResourceSpec{}
	for s, resourceSpec := range v.Invokes {
		specs[s] = resourceSpec
	}
	for s, resourceSpec := range v.Resources {
		specs[s] = resourceSpec
	}
	return specs
}

// ResourceSpec contains a pointer in an Open API Spec that defines a resource and related metadata.
type ResourceSpec struct {
	Path               string
	PathItem           *spec.PathItem
	PathItemList       *spec.PathItem
	Swagger            *Spec
	CompatibleVersions []string
	DefaultBody        map[string]interface{}
	DeprecationMessage string
	PreviousVersion    string
}

// ApplyProvidersTransformations adds the default version for each provider and deprecates and removes specified API versions.
func ApplyProvidersTransformations(providers AzureProviders, defaultVersion DefaultVersionLock, previousVersion DefaultVersionLock, deprecated, removed ProviderVersionList) AzureProviders {
	ApplyRemovals(providers, removed)
	AddDefaultVersion(providers, defaultVersion, previousVersion)
	ApplyDeprecations(providers, deprecated)

	return providers
}

func ApplyRemovals(providers map[string]map[string]VersionResources, removed map[string][]string) {
	for _, providerName := range codegen.SortedKeys(providers) {
		versionMap := providers[providerName]
		if removedVersion, ok := removed[providerName]; ok {
			for _, versionToRemove := range removedVersion {
				sdkVersionToRemove := ApiToSdkVersion(versionToRemove)
				delete(versionMap, sdkVersionToRemove)
			}
		}
	}
}

func AddDefaultVersion(providers map[string]map[string]VersionResources, defaultVersion DefaultVersionLock, previousVersion DefaultVersionLock) {
	for _, providerName := range codegen.SortedKeys(providers) {
		versionMap := providers[providerName]
		// Add a default version for each resource and invoke.
		defaultResourceVersions := defaultVersion[providerName]
		versionMap[""] = buildDefaultVersion(versionMap, defaultResourceVersions, previousVersion[providerName])

		// Set compatible versions to all other versions of the resource with the same normalized API path.
		pathVersions := calculatePathVersions(versionMap)
		versions := []string{}
		for version := range versionMap {
			versions = append(versions, version)
		}
		sort.Strings(versions)
		for _, version := range versions {
			items := versionMap[version]
			for _, resourceName := range codegen.SortedKeys(items.Resources) {
				r := items.Resources[resourceName]
				var otherVersions []string
				normalisedPath := paths.NormalizePath(r.Path)
				otherVersionsSorted := pathVersions[normalisedPath].SortedValues()
				for _, otherVersion := range otherVersionsSorted {
					if otherVersion != version {
						otherVersions = append(otherVersions, otherVersion)
					}
				}
				r.CompatibleVersions = otherVersions
			}
		}
	}
}

func ApplyDeprecations(providers AzureProviders, deprecated ProviderVersionList) AzureProviders {
	for _, providerName := range codegen.SortedKeys(providers) {
		versionMap := providers[providerName]
		for _, apiVersion := range deprecated[providerName] {
			sdkVersion := ApiToSdkVersion(apiVersion)
			resources := versionMap[sdkVersion]
			deprecateAll(resources.All(), apiVersion)
		}
	}

	return providers
}

func buildDefaultVersion(versionMap ProviderVersions, defaultResourceVersions map[ResourceName]ApiVersion, previousResourceVersions map[ResourceName]ApiVersion) VersionResources {
	resources := map[string]*ResourceSpec{}
	invokes := map[string]*ResourceSpec{}
	for _, resourceName := range codegen.SortedKeys(defaultResourceVersions) {
		apiVersion := defaultResourceVersions[resourceName]
		if versionResources, ok := versionMap[ApiToSdkVersion(apiVersion)]; ok {
			if resource, ok := versionResources.Resources[resourceName]; ok {
				resourceCopy := *resource
				resourceCopy.PreviousVersion = previousResourceVersions[resourceName]
				resources[resourceName] = &resourceCopy
			} else if invoke, ok := versionResources.Invokes[resourceName]; ok {
				invokeCopy := *invoke
				invokeCopy.PreviousVersion = previousResourceVersions[resourceName]
				invokes[resourceName] = &invokeCopy
			}
		}
	}
	return VersionResources{
		Resources: resources,
		Invokes:   invokes,
	}
}

// ReadAzureProviders finds Azure Open API specs on disk, parses them, and creates in-memory representation of resources,
// collected per Azure Provider and API Version - for all API versions.
// Use the namespace "*" to load all available namespaces, or a specific namespace to filter, e.g. "Compute".
// Use apiVersions with a wildcard to filter versions, e.g. "2022*preview", or leave it blank to use the default of "20*".
func ReadAzureProviders(specsDir, namespace, apiVersions string) (AzureProviders, DiscoveryDiagnostics, error) {
	diagnostics := DiscoveryDiagnostics{}
	swaggerSpecLocations, err := swaggerLocations(specsDir, namespace, apiVersions)
	if err != nil {
		return nil, diagnostics, err
	}

	// Collect all versions for each path in the API across all Swagger files.
	providers := AzureProviders{}
	for _, location := range swaggerSpecLocations {
		relLocation, err := filepath.Rel(specsDir, location)
		if err != nil {
			return nil, diagnostics, errors.Wrapf(err, "failed to get relative path for %q", location)
		}
		swagger, err := NewSpec(location)
		if err != nil {
			return nil, diagnostics, errors.Wrapf(err, "failed to parse %q", location)
		}

		orderedPaths := make([]string, 0, len(swagger.Paths.Paths))
		for path := range swagger.Paths.Paths {
			orderedPaths = append(orderedPaths, path)
		}
		sort.Strings(orderedPaths)
		for _, path := range orderedPaths {
			namingDisambiguations := providers.addAPIPath(specsDir, relLocation, path, swagger)
			diagnostics.NamingDisambiguations = append(diagnostics.NamingDisambiguations, namingDisambiguations...)
		}
	}
	return providers, diagnostics, nil
}

func deprecateAll(resourceSpecs map[string]*ResourceSpec, version string) {
	for _, resourceSpec := range resourceSpecs {
		deprecationMessage := fmt.Sprintf(
			"Version %s will be removed in v2 of the provider.",
			version)
		resourceSpec.DeprecationMessage = deprecationMessage
	}
}

// SingleVersion returns only a single (latest or preview) version of each resource from the full list of resource
// versions.
func SingleVersion(providers AzureProviders) AzureProviders {
	singleVersion := AzureProviders{}

	for providerName, allVersionMap := range providers {
		singleVersion[providerName] = ProviderVersions{"": allVersionMap[""]}
	}

	return singleVersion
}

// IsPreview returns true for API versions that aren't considered stable.
func IsPreview(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "preview") || strings.Contains(lower, "beta")
}

func ApiVersionToDate(apiVersion string) (time.Time, error) {
	if len(apiVersion) < 10 {
		return time.Time{}, fmt.Errorf("invalid API version %q", apiVersion)
	}
	// The API version is in the format YYYY-MM-DD - ignore suffixes like "-preview".
	return time.Parse("2006-01-02", apiVersion[:10])
}

func IsPrivate(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "private")
}

func CompareApiVersions(a, b string) int {
	timeA, err := ApiVersionToDate(a)
	if err != nil {
		return strings.Compare(a, b)
	}
	timeB, err := ApiVersionToDate(b)
	if err != nil {
		return strings.Compare(a, b)
	}
	timeDiff := timeA.Compare(timeB)
	if timeDiff != 0 {
		return timeDiff
	}

	// Sort private first, preview second, stable last.
	aPrivate := IsPrivate(a)
	bPrivate := IsPrivate(b)
	if aPrivate != bPrivate {
		if aPrivate {
			return -1
		}
		return 1
	}
	aPreview := IsPreview(a)
	bPreview := IsPreview(b)
	if aPreview != bPreview {
		if aPreview {
			return -1
		}
		return 1
	}
	return 0
}

func SortApiVersions(versions []string) {
	sort.SliceStable(versions, func(i, j int) bool {
		return CompareApiVersions(versions[i], versions[j]) < 0
	})
}

// swaggerLocations returns a slice of URLs of all known Azure Resource Manager swagger files.
// namespace and apiVersion can be blank to return all files, or can be used to filter the results.
func swaggerLocations(specsDir, namespace, apiVersions string) ([]string, error) {
	if namespace == "" {
		namespace = "*"
	}
	if apiVersions == "" {
		apiVersions = "20*"
	}

	pattern := filepath.Join(specsDir, "specification", "*", "resource-manager", "Microsoft."+namespace, "*", apiVersions, "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	pattern2 := filepath.Join(specsDir, "specification", "*", "resource-manager", "Microsoft."+namespace, "*", "*", apiVersions, "*.json")
	files2, err := filepath.Glob(pattern2)
	if err != nil {
		return nil, err
	}

	fileSet := codegen.NewStringSet()
	for _, file := range append(files, files2...) {
		// In December 2022, Azure started authoring some API specs in https://github.com/microsoft/cadl.
		// pattern2 above matches some of these folders, like
		// voiceservices/resource-manager/Microsoft.VoiceServices/cadl/examples/2023-01-31, so we exclude them.
		if !strings.Contains(file, "/cadl/") {
			fileSet.Add(file)
		}
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	return fileSet.SortedValues(), nil
}

var excludeRegexes = []*regexp.Regexp{
	// This preview version defines two types with the same name (one enum, one object) which fails to pass our codegen.
	// It's old, preview, and not important - so exclude the files of this version.
	regexp.MustCompile(".*frontdoor/resource-manager/Microsoft.Network/preview/2018-08-01-preview.*"),
}

// addAPIPath considers whether an API path contains resources and/or invokes and adds corresponding entries to the
// provider map. `providers` are mutated in-place.
func (providers AzureProviders) addAPIPath(specsDir, fileLocation, path string, swagger *Spec) []resources.NameDisambiguation {
	for _, re := range excludeRegexes {
		if re.MatchString(fileLocation) {
			return nil
		}
	}

	prov := resources.ResourceProvider(filepath.Join(specsDir, fileLocation), path)
	if prov == "" {
		return nil
	}

	// Find (or create) the version map with this name.
	versionMap, ok := providers[prov]
	if !ok {
		versionMap = map[string]VersionResources{}
		providers[prov] = versionMap
	}

	// Find (or create) the resource map with this name.
	apiVersion := ApiToSdkVersion(swagger.Info.Version)
	version, ok := versionMap[apiVersion]
	if !ok {
		version = NewVersionResources()
		versionMap[apiVersion] = version
	}

	return addResourcesAndInvokes(version, fileLocation, path, swagger)
}

func addResourcesAndInvokes(version VersionResources, fileLocation, path string, swagger *Spec) []resources.NameDisambiguation {
	apiVersion := ApiToSdkVersion(swagger.Info.Version)

	pathItem := swagger.Paths.Paths[path]
	pathItemList, hasList := swagger.Paths.Paths[path+"/list"]

	if ok := resources.HasCustomDelete(path); ok {
		pathItem.Delete = &spec.Operation{
			OperationProps: spec.OperationProps{
				Description: "Custom implementation of this operation is available",
			},
		}
	}

	var nameDisambiguations []resources.NameDisambiguation
	recordDisambiguation := func(disambiguation *resources.NameDisambiguation) []resources.NameDisambiguation {
		if disambiguation != nil {
			disambiguation.FileLocation = fileLocation
			nameDisambiguations = append(nameDisambiguations, *disambiguation)
		}
		return nameDisambiguations
	}

	// Add a resource entry, if appropriate HTTP endpoints are defined.
	foundResourceOrInvoke := false
	addResource := func(typeName string, defaultBody map[string]interface{}, pathItemList *spec.PathItem) {
		version.Resources[typeName] = &ResourceSpec{
			Path:         path,
			PathItem:     &pathItem,
			Swagger:      swagger,
			DefaultBody:  defaultBody,
			PathItemList: pathItemList,
		}
		foundResourceOrInvoke = true
	}
	addInvoke := func(typeName string) {
		version.Invokes[typeName] = &ResourceSpec{
			Path:     path,
			PathItem: &pathItem,
			Swagger:  swagger,
		}
		foundResourceOrInvoke = true
	}

	if pathItem.Put != nil && !pathItem.Put.Deprecated {
		hasDelete := pathItem.Delete != nil && !pathItem.Delete.Deprecated

		switch {
		case pathItem.Get != nil && !pathItem.Get.Deprecated:
			defaultState := defaults.GetDefaultResourceState(path)
			if defaultState != nil && hasDelete && defaultState.HasNonEmptyCollections {
				// See the limitation in `SdkShapeConverter.isDefaultResponse()`
				panic(fmt.Sprintf("invalid defaultResourcesState '%s': non-empty collections aren't supported for deletable resources", path))
			}

			typeName, disambiguation := resources.ResourceName(pathItem.Get.ID, path)
			nameDisambiguations = recordDisambiguation(disambiguation)

			if typeName != "" && (hasDelete || defaultState != nil) {
				if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", apiVersion, typeName, path, version.Resources[typeName].Path)
				}
				defaultBody := map[string]interface{}{}
				if defaultState != nil {
					defaultBody = defaultState.State
				}
				addResource(typeName, defaultBody, nil /* pathItemList */)
			}
		case pathItem.Head != nil && !pathItem.Head.Deprecated:
			typeName, disambiguation := resources.ResourceName(pathItem.Head.ID, path)
			nameDisambiguations = recordDisambiguation(disambiguation)

			if typeName != "" && hasDelete {
				if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", apiVersion, typeName, path, version.Resources[typeName].Path)
				}
				addResource(typeName, nil /* defaultBody */, nil /* pathItemList */)
			}
		case hasList:
			var typeName string
			var disambiguation *resources.NameDisambiguation
			switch {
			case pathItemList.Get != nil && !pathItemList.Get.Deprecated:
				typeName, disambiguation = resources.ResourceName(pathItemList.Get.ID, path)
			case pathItemList.Post != nil && !pathItemList.Post.Deprecated:
				typeName, disambiguation = resources.ResourceName(pathItemList.Post.ID, path)
			}
			nameDisambiguations = recordDisambiguation(disambiguation)

			if typeName != "" {
				var defaultBody map[string]interface{}
				defaultState := defaults.GetDefaultResourceState(path)
				if defaultState != nil {
					defaultBody = defaultState.State
				} else if !hasDelete {
					// The /list pattern that we handle here seems to (almost) universally have this shape of the default body.
					// Instead of maintaining the resources in defaultResourcesState, we can hard-code it here.
					defaultBody = map[string]interface{}{
						"properties": map[string]interface{}{},
					}
				}
				if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", apiVersion, typeName, path, version.Resources[typeName].Path)
				}
				addResource(typeName, defaultBody, &pathItemList)
			}
		}
	}

	// Add an entry for PATCH-based resources.
	if pathItem.Patch != nil && !pathItem.Patch.Deprecated && pathItem.Get != nil && !pathItem.Get.Deprecated {
		defaultState := defaults.GetDefaultResourceState(path)
		typeName, disambiguation := resources.ResourceName(pathItem.Get.ID, path)
		nameDisambiguations = recordDisambiguation(disambiguation)

		if typeName != "" && defaultState != nil {
			if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
				fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", apiVersion, typeName, path, version.Resources[typeName].Path)
			}
			addResource(typeName, defaultState.State, nil /* pathItemList */)
		}
	}

	// Add a POST invoke entry.
	if pathItem.Post != nil && !pathItem.Post.Deprecated {
		parts := strings.Split(path, "/")
		operationName := strings.ToLower(parts[len(parts)-1])
		operationId := operationFromOperationID(pathItem.Post.OperationProps.ID)
		prefix := ""
		switch {
		case strings.HasPrefix(operationName, "list"):
			prefix = "list"
		case strings.HasPrefix(operationName, "get"):
			prefix = "get"
		case strings.HasPrefix(operationId, "get") && pathItem.Put == nil &&
			(strings.Contains(operationName, "key") ||
				strings.Contains(operationName, "token") ||
				strings.Contains(operationName, "credential")):
			// Operation ID-based selection is a bit tricky, so we apply the following heuristic:
			// - Called according to the pattern `Foo_GetBar`,
			// - It's not a resource (ensured by lack of a PUT operation),
			// - It's about a key, a token, or credentials.
			prefix = "get"
		default:
			if !foundResourceOrInvoke && shouldIncludeInvoke(path, pathItem.Post) {
				prefix = operationId
			} else {
				return nameDisambiguations
			}
		}

		typeName, disambiguation := resources.ResourceName(pathItem.Post.ID, path)
		if typeName != "" {
			addInvoke(prefix + typeName)
			nameDisambiguations = recordDisambiguation(disambiguation)
		}
	}

	// Add an invoke if a GET endpoint is defined, but only if we haven't added a resource for this path yet.
	// Resources can be read through the Pulumi resource model without a dedicated invoke.
	if !foundResourceOrInvoke && pathItem.Get != nil && shouldIncludeInvoke(path, pathItem.Get) {
		typeName, disambiguation := resources.ResourceName(pathItem.Get.ID, path)
		if typeName != "" {
			operation := operationFromOperationID(pathItem.Get.OperationProps.ID)
			prefix := "get"
			if operation == "list" {
				prefix = "list"
			}

			addInvoke(prefix + typeName)
			nameDisambiguations = recordDisambiguation(disambiguation)
		}
	}

	return nameDisambiguations
}

// DiagnosticSettingsCategory_List -> list
func operationFromOperationID(opID string) string {
	parts := strings.Split(opID, "_")
	return strings.ToLower(parts[len(parts)-1])
}
