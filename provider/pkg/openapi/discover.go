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
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources/customresources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

type ModuleName = resources.ModuleName

// ApiVersion e.g. 2020-01-30
// Occasionally we use empty string to represent the default version or no version.
// Sometimes this is also used as a query and can include a wildcard.
type ApiVersion string

func (v ApiVersion) IsDefault() bool {
	return v == ""
}

func (v ApiVersion) IsWildcard() bool {
	return strings.Contains(string(v), "*")
}

func (v ApiVersion) ToSdkVersion() SdkVersion {
	return ApiToSdkVersion(v)
}

// DefinitionName is the name of either an 'invoke' or a resource (e.g. listBuckets or Bucket)
type DefinitionName = string

// ResourceName e.g. Bucket
type ResourceName = string

// InvokeName e.g. listBuckets
type InvokeName = string

// SdkVersion e.g. v20200130
type SdkVersion string

// AzureModules maps module names (e.g. Compute) to versions in that module and resources therein.
type AzureModules map[ModuleName]ModuleVersions

// ModuleVersions maps API Versions (e.g. v20200801) to resources and invokes in that version.
type ModuleVersions = map[ApiVersion]VersionResources

// Represents a failed attempt to determine the module name for a given path within a spec.
// This results in the path being skipped and not considered for resource or invoke generation.
type ModuleNameError struct {
	FilePath string
	Path     string
	Error    string
}

type DiscoveryDiagnostics struct {
	NamingDisambiguations []resources.NameDisambiguation
	// POST endpoints defined in the Azure spec that we don't include because they don't belong to a resource.
	// Map is module -> operation id -> path.
	SkippedPOSTEndpoints map[ModuleName]map[string]string
	// module -> resource/type name -> path -> Endpoints
	Endpoints Endpoints
	// Errors where we can't determine the module name for a given path within a spec.
	ModuleNameErrors []ModuleNameError
}

// module -> resource/type name -> path -> Endpoint
type Endpoints map[ModuleName]map[ResourceName]map[string]*Endpoint

// merge combines e2 into e, which is modified in-place.
func (e Endpoints) merge(e2 Endpoints) {
	for moduleName, moduleEndpoints := range e2 {
		if _, ok := e[moduleName]; !ok {
			e[moduleName] = map[string]map[string]*Endpoint{}
		}

		for typeName, byPath := range moduleEndpoints {
			if _, ok := e[moduleName][typeName]; !ok {
				e[moduleName][typeName] = map[string]*Endpoint{}
			}
			for path, things := range byPath {
				if existing, ok := e[moduleName][typeName][path]; ok {
					// the POST endpoints are unique per version, but we don't track versions here, so check for duplicates
					ops := codegen.NewStringSet(existing.PostOperations...)
					for _, op := range things.PostOperations {
						ops.Add(op)
					}
					existing.PostOperations = ops.SortedValues()

					verbs := codegen.NewStringSet(existing.HttpVerbs...)
					for _, verb := range things.HttpVerbs {
						verbs.Add(verb)
					}
					existing.HttpVerbs = verbs.SortedValues()
				} else {
					e[moduleName][typeName][path] = things
				}
			}
		}
	}
}

func (e Endpoints) add(pathItem spec.PathItem, moduleName ModuleName, typeName, path, filePath string, addedResourceOrInvoke bool) {
	if _, ok := e[moduleName]; !ok {
		e[moduleName] = map[string]map[string]*Endpoint{}
	}

	endpoint := &Endpoint{
		Path:     path,
		FilePath: filePath,
		Added:    addedResourceOrInvoke,
	}
	if pathItem.Delete != nil && !pathItem.Delete.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "DELETE")
	}
	if pathItem.Get != nil && !pathItem.Get.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "GET")
	}
	if pathItem.Head != nil && !pathItem.Head.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "HEAD")
	}
	if pathItem.Patch != nil && !pathItem.Patch.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "PATCH")
	}
	if pathItem.Post != nil && !pathItem.Post.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "POST")
		lastSlash := strings.LastIndex(path, "/")
		endpoint.PostOperations = append(endpoint.PostOperations, path[lastSlash+1:])
		endpoint.Path = path[:lastSlash] // normalize path to not include the POST operation, to match the resource path
	}
	if pathItem.Put != nil && !pathItem.Put.Deprecated {
		endpoint.HttpVerbs = append(endpoint.HttpVerbs, "PUT")
	}

	if _, ok := e[moduleName][typeName]; !ok {
		e[moduleName][typeName] = map[string]*Endpoint{}
	}
	e[moduleName][typeName][endpoint.Path] = endpoint
}

type Endpoint struct {
	Path                          string
	FilePath                      string
	HttpVerbs                     []string
	Get, Put, Delete, Patch, Head string   `json:",omitempty"` // operation id
	PostOperations                []string `json:",omitempty"` // path suffixes
	Added                         bool     `json:",omitempty"`
}

func (d *DiscoveryDiagnostics) addSkippedPOSTEndpoint(moduleName ModuleName, operation, path string) {
	cur, ok := d.SkippedPOSTEndpoints[moduleName]
	if !ok {
		cur = map[string]string{}
		d.SkippedPOSTEndpoints[moduleName] = cur
	}
	cur[operation] = path
}

func (d *DiscoveryDiagnostics) addPathItem(pathItem spec.PathItem, moduleName ModuleName, typeName, path, filePath string, addedResourceOrInvoke bool) {
	if d.Endpoints == nil {
		d.Endpoints = map[ModuleName]map[string]map[string]*Endpoint{}
	}
	d.Endpoints.add(pathItem, moduleName, typeName, path, filePath, addedResourceOrInvoke)
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

type ModuleVersionList = map[ModuleName][]ApiVersion

// A definition version is a resource or invoke version and its source information.
type DefinitionVersion struct {
	ApiVersion   ApiVersion `yaml:"ApiVersion,omitempty"`
	SpecFilePath string     `yaml:"SpecFilePath,omitempty"`
	ResourceUri  string     `yaml:"ResourceUri,omitempty"`
	RpNamespace  string     `yaml:"RpNamespace,omitempty"`
}

// DefaultVersions is an amalgamation of multiple API versions, generated from a specification.
type DefaultVersions map[ModuleName]map[DefinitionName]DefinitionVersion

func (defaultVersions DefaultVersions) IsAtVersion(moduleName ModuleName, typeName DefinitionName, version ApiVersion) bool {
	if resources, ok := defaultVersions[moduleName]; ok {
		if resourceVersion, ok := resources[typeName]; ok {
			if resourceVersion.ApiVersion == version {
				return true
			}
		}
	}
	return false
}

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
	// FileLocation is the path to the Open API Spec file.
	FileLocation string
	// API version of the resource
	ApiVersion ApiVersion
	// SDK version of the resource
	SdkVersion SdkVersion
	// Path is the API path in the Open API Spec.
	Path         string
	PathItem     *spec.PathItem
	PathItemList *spec.PathItem
	Swagger      *Spec
	// CompatibleVersions is a list of other API versions that are compatible with this resource.
	// These versions can be the default version, indicated by an empty string.
	CompatibleVersions []ApiVersion
	DefaultBody        map[string]interface{}
	DeprecationMessage string
	PreviousVersion    ApiVersion
	ModuleNaming       resources.ModuleNaming
}

// ApplyTransformations adds the default version for each module and deprecates and removes specified API versions.
func ApplyTransformations(modules AzureModules, defaultVersions DefaultVersions, previousDefaultVersions DefaultVersions, deprecated, removed ModuleVersionList) AzureModules {
	ApplyRemovals(modules, removed)
	AddDefaultVersion(modules, defaultVersions, previousDefaultVersions)
	ApplyDeprecations(modules, deprecated)

	return modules
}

func ApplyRemovals(modules map[ModuleName]ModuleVersions, removed ModuleVersionList) {
	for moduleName, versionMap := range util.MapOrdered(modules) {
		if removedVersion, ok := removed[moduleName]; ok {
			for _, versionToRemove := range removedVersion {
				delete(versionMap, versionToRemove)
			}
		}
	}
}

func AddDefaultVersion(modules AzureModules, defaultVersions DefaultVersions, previousDefaultVersions DefaultVersions) {
	for moduleName, versionMap := range util.MapOrdered(modules) {
		// Add a default version for each resource and invoke.
		defaultResourceVersions := defaultVersions[moduleName]
		versionMap[""] = buildDefaultVersion(versionMap, defaultResourceVersions, previousDefaultVersions[moduleName])

		// Set compatible versions to all other versions of the resource with the same normalized API path.
		pathVersions := calculatePathVersions(versionMap)
		for version, items := range util.MapOrdered(versionMap) {
			for _, resource := range util.MapOrdered(items.Resources) {
				var otherVersions []ApiVersion
				normalisedPath := paths.NormalizePath(resource.Path)
				otherVersionsSorted := pathVersions[normalisedPath].SortedValues()
				for _, otherVersion := range otherVersionsSorted {
					if otherVersion != version {
						otherVersions = append(otherVersions, otherVersion)
					}
				}
				resource.CompatibleVersions = otherVersions
			}
		}
	}
}

func ApplyDeprecations(modules AzureModules, deprecated ModuleVersionList) AzureModules {
	for moduleName, versionMap := range util.MapOrdered(modules) {
		if deprecatedVersions, ok := deprecated[moduleName]; ok {
			for _, apiVersion := range deprecatedVersions {
				resources := versionMap[apiVersion]
				deprecateAll(resources.All(), apiVersion)
			}
		}
	}

	return modules
}

func buildDefaultVersion(versionMap ModuleVersions, defaultResourceVersions map[ResourceName]DefinitionVersion, previousResourceVersions map[ResourceName]DefinitionVersion) VersionResources {
	resources := map[string]*ResourceSpec{}
	invokes := map[string]*ResourceSpec{}
	for resourceName, apiVersion := range util.MapOrdered(defaultResourceVersions) {
		if versionResources, ok := versionMap[apiVersion.ApiVersion]; ok {
			if resource, ok := versionResources.Resources[resourceName]; ok {
				resourceCopy := *resource
				if previousVersion, hasPreviousVersion := previousResourceVersions[resourceName]; hasPreviousVersion {
					resourceCopy.PreviousVersion = previousVersion.ApiVersion
				}
				resources[resourceName] = &resourceCopy
			} else if invoke, ok := versionResources.Invokes[resourceName]; ok {
				invokeCopy := *invoke
				if previousVersion, hasPreviousVersion := previousResourceVersions[resourceName]; hasPreviousVersion {
					invokeCopy.PreviousVersion = previousVersion.ApiVersion
				}
				invokes[resourceName] = &invokeCopy
			}
		}
	}
	return VersionResources{
		Resources: resources,
		Invokes:   invokes,
	}
}

// ReadAzureModules finds Azure Open API specs on disk, parses them, and creates in-memory representation of resources,
// collected per Azure Module and API Version - for all API versions.
// Use the namespace "*" to load all available namespaces, or a specific namespace to filter, e.g. "Compute".
// Use apiVersions with a wildcard to filter versions, e.g. "2022*preview", or leave it blank to use the default of "20*".
func ReadAzureModules(specsDir, namespace, apiVersions string) (AzureModules, DiscoveryDiagnostics, error) {
	diagnostics := DiscoveryDiagnostics{
		SkippedPOSTEndpoints: map[ModuleName]map[string]string{},
		Endpoints:            map[ModuleName]map[string]map[string]*Endpoint{},
	}
	swaggerSpecLocations, err := swaggerLocations(specsDir, namespace, apiVersions)
	if err != nil {
		return nil, diagnostics, err
	}

	// Collect all versions for each path in the API across all Swagger files.
	modules := AzureModules{}
	for _, location := range swaggerSpecLocations {
		relLocation, err := filepath.Rel(specsDir, location)
		if err != nil {
			return nil, diagnostics, errors.Wrapf(err, "failed to get relative path for %q", location)
		}

		if exclude(relLocation) {
			continue
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
			moduleDiagnostics := modules.addAPIPath(specsDir, relLocation, path, swagger)

			// Update reports
			diagnostics.NamingDisambiguations = append(diagnostics.NamingDisambiguations, moduleDiagnostics.NamingDisambiguations...)
			for moduleName, operations := range moduleDiagnostics.SkippedPOSTEndpoints {
				for op, path := range operations {
					diagnostics.addSkippedPOSTEndpoint(moduleName, op, path)
				}
			}
			diagnostics.Endpoints.merge(moduleDiagnostics.Endpoints)
			diagnostics.ModuleNameErrors = append(diagnostics.ModuleNameErrors, moduleDiagnostics.ModuleNameErrors...)
		}
	}
	return modules, diagnostics, nil
}

func deprecateAll(resourceSpecs map[string]*ResourceSpec, version ApiVersion) {
	for _, resourceSpec := range resourceSpecs {
		deprecationMessage := fmt.Sprintf(
			"Version %s will be removed in v2 of the provider.",
			version)
		resourceSpec.DeprecationMessage = deprecationMessage
	}
}

// SingleVersion returns only a single (latest or preview) version of each resource from the full list of resource
// versions.
func SingleVersion(modules AzureModules) AzureModules {
	singleVersion := AzureModules{}

	for moduleName, allVersionMap := range modules {
		singleVersion[moduleName] = ModuleVersions{"": allVersionMap[""]}
	}

	return singleVersion
}

// IsPreview returns true for API versions that aren't considered stable.
func IsPreview(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "preview") || strings.Contains(lower, "beta")
}

func ApiVersionToDate(apiVersion ApiVersion) (time.Time, error) {
	if len(apiVersion) < 10 {
		return time.Time{}, fmt.Errorf("invalid API version %q", apiVersion)
	}
	// The API version is in the format YYYY-MM-DD - ignore suffixes like "-preview".
	return time.Parse("2006-01-02", string(apiVersion)[:10])
}

// Check if the string contains the word "private", ignoring case.
func IsPrivate(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "private")
}

// Attempt to convert both versions to dates and compare them.
// Fall back to string comparison if either version is not a date.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func CompareApiVersions(a, b ApiVersion) int {
	timeA, err := ApiVersionToDate(a)
	if err != nil {
		return strings.Compare(string(a), string(b))
	}
	timeB, err := ApiVersionToDate(b)
	if err != nil {
		return strings.Compare(string(a), string(b))
	}
	timeDiff := timeA.Compare(timeB)
	if timeDiff != 0 {
		return timeDiff
	}

	// Sort private first, preview second, stable last.
	aPrivate := IsPrivate(string(a))
	bPrivate := IsPrivate(string(b))
	if aPrivate != bPrivate {
		if aPrivate {
			return -1
		}
		return 1
	}
	aPreview := IsPreview(string(a))
	bPreview := IsPreview(string(b))
	if aPreview != bPreview {
		if aPreview {
			return -1
		}
		return 1
	}
	return 0
}

func SortApiVersions(versions []ApiVersion) {
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

	pattern3 := filepath.Join(specsDir, "specification", "*", "resource-manager", "PaloAltoNetworks."+namespace, "*", apiVersions, "*.json")
	files3, err := filepath.Glob(pattern3)
	if err != nil {
		return nil, err
	}

	fileSet := codegen.NewStringSet()
	for _, file := range append(append(files, files2...), files3...) {
		// In December 2022, Azure started authoring some API specs in https://github.com/microsoft/cadl.
		// pattern2 above matches some of these folders, like
		// voiceservices/resource-manager/Microsoft.VoiceServices/cadl/examples/2023-01-31, so we exclude them.
		if strings.Contains(file, "/cadl/") {
			continue
		}
		if strings.Contains(file, "/examples/") {
			continue
		}
		fileSet.Add(file)
	}

	// Sorting alphabetically means the schemas with the latest API version are the last.
	return fileSet.SortedValues(), nil
}

var excludeRegexes = []*regexp.Regexp{
	// This preview version defines two types with the same name (one enum, one object) which fails to pass our codegen.
	// It's old, preview, and not important - so exclude the files of this version.
	regexp.MustCompile(".*frontdoor/resource-manager/Microsoft.Network/preview/2018-08-01-preview.*"),
	// This version conflicts with the managed folder version:
	// servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview
	// This causes a conflict in the version-specific folder, not the default version folder, so we have to completely exclude it.
	regexp.MustCompile(".*servicefabric/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview.*"),
	// This preview version is invalid OpenAPI JSON, reading it fails with encoding/json.UnmarshalTypeError in field "definitions".
	regexp.MustCompile(".*network/resource-manager/Microsoft.Network/preview/2023-03-01-preview.*"),
}

func exclude(filePath string) bool {
	for _, re := range excludeRegexes {
		if re.MatchString(filePath) {
			return true
		}
	}
	return false
}

// addAPIPath considers whether an API path contains resources and/or invokes and adds corresponding entries to the
// module map. Modules are mutated in-place.
func (modules AzureModules) addAPIPath(specsDir, fileLocation, path string, swagger *Spec) DiscoveryDiagnostics {
	moduleNaming, err := resources.GetModuleName(version.GetVersion().Major, filepath.Join(specsDir, fileLocation), path)
	if err != nil {
		return DiscoveryDiagnostics{
			ModuleNameErrors: []ModuleNameError{
				{
					FilePath: fileLocation,
					Path:     path,
					Error:    err.Error(),
				},
			},
		}
	}
	moduleName := moduleNaming.ResolvedName

	// Find (or create) the version map with this name.
	versionMap, ok := modules[moduleName]
	if !ok {
		versionMap = map[ApiVersion]VersionResources{}
		modules[moduleName] = versionMap
	}

	// Find (or create) the resource map with this name.
	apiVersion := ApiVersion(swagger.Info.Version)
	version, ok := versionMap[apiVersion]
	if !ok {
		version = NewVersionResources()
		versionMap[apiVersion] = version
	}

	return addResourcesAndInvokes(version, fileLocation, path, moduleNaming, swagger)
}

// getTypeName returns the type name for a given operation and path. The path is used to check for custom resource
// names. In the standard case, it's unused and the name is based on the operation id.
func getTypeName(op *spec.Operation, path string) (string, *resources.NameDisambiguation) {
	if typeName, found := customresources.GetCustomResourceName(path); found {
		return typeName, nil
	}
	return resources.ResourceName(op.ID, path)
}

func addResourcesAndInvokes(version VersionResources, fileLocation, path string, moduleNaming resources.ModuleNaming, swagger *Spec) DiscoveryDiagnostics {
	apiVersion := ApiVersion(swagger.Info.Version)
	sdkVersion := ApiToSdkVersion(apiVersion)

	pathItem := swagger.Paths.Paths[path]
	pathItemList, hasList := swagger.Paths.Paths[path+"/list"]

	if ok := customresources.HasCustomDelete(path); ok {
		pathItem.Delete = &spec.Operation{
			OperationProps: spec.OperationProps{
				Description: "Custom implementation of this operation is available",
			},
		}
	}

	diagnostics := DiscoveryDiagnostics{
		SkippedPOSTEndpoints: map[ModuleName]map[string]string{},
	}

	recordDisambiguation := func(disambiguation *resources.NameDisambiguation) {
		if disambiguation != nil {
			disambiguation.FileLocation = fileLocation
			diagnostics.NamingDisambiguations = append(diagnostics.NamingDisambiguations, *disambiguation)
		}
	}

	// Add a resource entry, if appropriate HTTP endpoints are defined.
	foundResourceOrInvoke := false
	addResource := func(typeName string, defaultBody map[string]interface{}, pathItemList *spec.PathItem) {
		version.Resources[typeName] = &ResourceSpec{
			FileLocation: fileLocation,
			ApiVersion:   apiVersion,
			SdkVersion:   sdkVersion,
			Path:         path,
			PathItem:     &pathItem,
			Swagger:      swagger,
			DefaultBody:  defaultBody,
			PathItemList: pathItemList,
			ModuleNaming: moduleNaming,
		}
		foundResourceOrInvoke = true
	}
	addInvoke := func(typeName string) {
		version.Invokes[typeName] = &ResourceSpec{
			FileLocation: fileLocation,
			ApiVersion:   apiVersion,
			SdkVersion:   sdkVersion,
			Path:         path,
			PathItem:     &pathItem,
			Swagger:      swagger,
			ModuleNaming: moduleNaming,
		}
		foundResourceOrInvoke = true
	}

	var resourceBaseName string
	for _, spec := range []*spec.Operation{pathItem.Post, pathItem.Put, pathItem.Get} {
		if spec != nil && !spec.Deprecated {
			resourceBaseName = strings.Split(spec.ID, "_")[0]
			break
		}
	}

	if pathItem.Put != nil && !pathItem.Put.Deprecated {
		hasDelete := pathItem.Delete != nil && !pathItem.Delete.Deprecated

		switch {
		case pathItem.Get != nil && !pathItem.Get.Deprecated:
			defaultState := defaults.GetDefaultResourceState(path, swagger.Info.Version)
			if defaultState != nil && hasDelete && defaultState.HasNonEmptyCollections {
				// See the limitation in `SdkShapeConverter.isDefaultResponse()`
				panic(fmt.Sprintf("invalid defaultResourcesState '%s': non-empty collections aren't supported for deletable resources", path))
			}
			typeName, disambiguation := getTypeName(pathItem.Get, path)
			recordDisambiguation(disambiguation)

			isCustom, includeCustom := customresources.IncludeCustomResource(path, string(apiVersion))
			canBeDeleted := hasDelete || defaultState != nil
			if typeName != "" && ((isCustom && includeCustom) || (!isCustom && canBeDeleted)) {
				if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", sdkVersion, typeName, path, version.Resources[typeName].Path)
				}
				var defaultBody map[string]any
				if defaultState != nil {
					defaultBody = defaultState.State
				}
				addResource(typeName, defaultBody, nil /* pathItemList */)
			}
		case pathItem.Head != nil && !pathItem.Head.Deprecated:
			typeName, disambiguation := getTypeName(pathItem.Head, path)
			recordDisambiguation(disambiguation)

			if typeName != "" && hasDelete {
				if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", sdkVersion, typeName, path, version.Resources[typeName].Path)
				}
				addResource(typeName, nil /* defaultBody */, nil /* pathItemList */)
			}
		case hasList:
			var typeName string
			var disambiguation *resources.NameDisambiguation
			switch {
			case pathItemList.Get != nil && !pathItemList.Get.Deprecated:
				typeName, disambiguation = getTypeName(pathItemList.Get, path)
			case pathItemList.Post != nil && !pathItemList.Post.Deprecated:
				typeName, disambiguation = getTypeName(pathItemList.Post, path)
			}
			recordDisambiguation(disambiguation)

			if typeName != "" {
				var defaultBody map[string]interface{}
				defaultState := defaults.GetDefaultResourceState(path, swagger.Info.Version)
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
					fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", sdkVersion, typeName, path, version.Resources[typeName].Path)
				}
				addResource(typeName, defaultBody, &pathItemList)
			}
		}
	}

	// Add an entry for PATCH-based resources.
	if pathItem.Patch != nil && !pathItem.Patch.Deprecated && pathItem.Get != nil && !pathItem.Get.Deprecated {
		defaultState := defaults.GetDefaultResourceState(path, swagger.Info.Version)
		typeName, disambiguation := getTypeName(pathItem.Get, path)
		recordDisambiguation(disambiguation)

		isCustom, includeCustom := customresources.IncludeCustomResource(path, string(apiVersion))
		if typeName != "" && ((isCustom && includeCustom) || (!isCustom && defaultState != nil)) {
			if _, ok := version.Resources[typeName]; ok && version.Resources[typeName].Path != path {
				fmt.Printf("warning: duplicate resource %s/%s at paths:\n  - %s\n  - %s\n", sdkVersion, typeName, path, version.Resources[typeName].Path)
			}
			var defaultBody map[string]any
			if defaultState != nil {
				defaultBody = defaultState.State
			}
			addResource(typeName, defaultBody, nil /* pathItemList */)
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
		case (strings.HasPrefix(operationId, "get") || strings.HasPrefix(operationId, "retrieve")) &&
			pathItem.Put == nil &&
			(strings.Contains(operationName, "key") ||
				strings.Contains(operationName, "token") ||
				strings.Contains(operationName, "credential")):
			// Operation ID-based selection is a bit tricky, so we apply the following heuristic:
			// - Called according to the pattern `Foo_GetBar`,
			// - It's not a resource (ensured by lack of a PUT operation),
			// - It's about a key, a token, or credentials.
			prefix = "get"
		default:
			diagnostics.addSkippedPOSTEndpoint(moduleNaming.ResolvedName, pathItem.Post.ID, path)
		}

		if prefix != "" {
			typeName, disambiguation := getTypeName(pathItem.Post, path)
			if typeName != "" {
				addInvoke(prefix + typeName)
				recordDisambiguation(disambiguation)
			}
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
			recordDisambiguation(disambiguation)
		}
	}

	diagnostics.addPathItem(pathItem, moduleNaming.ResolvedName, resourceBaseName, path, fileLocation, foundResourceOrInvoke)
	return diagnostics
}

// DiagnosticSettingsCategory_List -> list
func operationFromOperationID(opID string) string {
	parts := strings.Split(opID, "_")
	return strings.ToLower(parts[len(parts)-1])
}
