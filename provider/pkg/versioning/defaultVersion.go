package versioning

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"gopkg.in/yaml.v3"
)

type ModuleSpec struct {
	// Tracking is a single API version from which updates will be included
	Tracking *openapi.ApiVersion `yaml:"tracking,omitempty"`
	// Additions are specific resource versions to be included. These *must not* overlap with any resources from the tracking version
	Additions *map[openapi.ResourceName]openapi.ApiVersion `yaml:"additions,omitempty"`
	// Validation warnings for this module
	ExclusionErrors []ExclusionError `yaml:"exclusionErrors,omitempty"`
}

type ExclusionError struct {
	ModuleName   openapi.ModuleName
	ResourceName string
	Detail       string
}

// A Spec describes what versions of what resources should be included in the module.
// It is the input to schema generation.
type Spec map[openapi.ModuleName]ModuleSpec

// BuildSpec calculates a Spec from available API versions and manual curations (config).
func BuildSpec(spec ModuleVersionResources, curations Curations, existingConfig Spec, activeVersionChecker providerlist.ActiveVersionChecker) Spec {
	specs := Spec{}
	for moduleName, versionResources := range spec {
		existing := existingConfig[moduleName]
		builder := moduleSpecBuilder{moduleName: moduleName, activeVersionChecker: activeVersionChecker}
		specs[moduleName] = builder.buildSpec(versionResources, curations, existing)
	}
	return specs
}

type CurationViolation struct {
	ModuleName openapi.ModuleName
	Detail     string
}

func ValidateDefaultConfig(config Spec, curations Curations) []CurationViolation {
	var violations []CurationViolation
	for _, moduleName := range util.SortedKeys(config) {
		moduleSpec := config[moduleName]
		moduleCuration := curations[moduleName]
		expectedTracking := moduleCuration.ExpectTracking
		if expectedTracking == "" {
			expectedTracking = "stable"
		}
		var actualTracking string
		if moduleSpec.Tracking != nil && openapi.IsPreview(string(*moduleSpec.Tracking)) {
			actualTracking = "preview"
		} else {
			actualTracking = "stable"
		}
		if expectedTracking != actualTracking {
			violations = append(violations, CurationViolation{
				ModuleName: moduleName,
				Detail:     fmt.Sprintf("expected tracking %s but found %s", expectedTracking, actualTracking),
			})
		}
		hasAdditions := moduleSpec.Additions != nil && len(*moduleSpec.Additions) > 0
		if moduleCuration.ExpectAdditions != hasAdditions && !moduleCuration.Explicit {
			var detail string
			if moduleCuration.ExpectAdditions {
				detail = "expected additions but found none"
			} else {
				detail = "expected no additions but found some"
			}
			violations = append(violations, CurationViolation{
				ModuleName: moduleName,
				Detail:     detail,
			})
		}
	}
	return violations
}

// ReadSpec parses a spec from a YAML file
func ReadSpec(path string) (Spec, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion Spec
	err = yaml.Unmarshal(byteValue, &curatedVersion)
	return curatedVersion, err
}

func FindInactiveDefaultVersions(defaultVersions openapi.DefaultVersions, activeVersions providerlist.ActiveVersionChecker) map[openapi.ModuleName][]openapi.ApiVersion {
	result := map[openapi.ModuleName][]openapi.ApiVersion{}
	for moduleName, versions := range defaultVersions {
		inactiveVersions := collections.NewOrderableSet[openapi.ApiVersion]()
		for _, version := range versions {
			// TODO: Use the original AZ namespace rather that our adjusted module name
			if !activeVersions.HasProviderVersion(string(moduleName), version.ApiVersion) {
				inactiveVersions.Add(version.ApiVersion)
			}
		}
		if inactiveVersions.Count() > 0 {
			result[moduleName] = inactiveVersions.SortedValues()
		}
	}
	return result
}

func DefaultVersionsFromConfig(spec ModuleVersionResources, defaultConfig Spec) (openapi.DefaultVersions, error) {
	var err error
	defaultVersions := openapi.DefaultVersions{}
	for moduleName, versionResources := range spec {
		definitions := map[openapi.DefinitionName]openapi.DefinitionVersion{}
		moduleConfig, ok := defaultConfig[moduleName]
		if !ok {
			if len(versionResources) > 0 {
				var versions []string
				for version := range versionResources {
					if version != "" {
						versions = append(versions, string(version))
					}
				}
				if len(versions) > 0 {
					err = multierror.Append(err, fmt.Errorf("no version specified for %s, available versions: %s", moduleName, strings.Join(versions, ", ")))
				}
				continue
			}
		}

		if moduleConfig.Tracking != nil {
			trackingDefinitions := versionResources[*moduleConfig.Tracking]
			for _, definitionName := range util.SortedKeys(trackingDefinitions) {
				definition := trackingDefinitions[definitionName]
				definitions[definitionName] = openapi.DefinitionVersion{
					ApiVersion:   definition.ApiVersion,
					SpecFilePath: definition.SpecFilePath,
					ResourceUri:  definition.ResourceUri,
					RpNamespace:  definition.RpNamespace,
				}
			}
		}
		if moduleConfig.Additions != nil {
			for definitionName, apiVersion := range *moduleConfig.Additions {
				if existingVersion, ok := definitions[definitionName]; ok {
					err = multierror.Append(err, fmt.Errorf("duplicate resource %s:%s from %s and %s", moduleName, definitionName, apiVersion, existingVersion))
				} else {
					versionDefinitions := versionResources[apiVersion]
					versionDefinition := versionDefinitions[definitionName]
					definitions[definitionName] = openapi.DefinitionVersion{
						ApiVersion:   apiVersion,
						SpecFilePath: versionDefinition.SpecFilePath,
						ResourceUri:  versionDefinition.ResourceUri,
						RpNamespace:  versionDefinition.RpNamespace,
					}
				}
			}
		}
		defaultVersions[moduleName] = definitions
	}
	return defaultVersions, multierror.Flatten(err)
}

type moduleSpecBuilder struct {
	moduleName           openapi.ModuleName
	activeVersionChecker providerlist.ActiveVersionChecker
}

func (b moduleSpecBuilder) buildSpec(versions VersionResources, curations Curations, existing ModuleSpec) ModuleSpec {
	var additionsPtr *map[string]openapi.ApiVersion
	var trackingPtr *openapi.ApiVersion
	var exclusionErrors []ExclusionError

	moduleCuration := curations[b.moduleName]
	latestVersions := b.findLatestVersions(versions, moduleCuration)

	if len(latestVersions) == 0 {
		return ModuleSpec{}
	}

	// If multiple versions required, track the latest and include additional resources from previous versions
	var trackingResources codegen.StringSet
	if existing.Tracking != nil {
		trackingPtr = existing.Tracking
		trackingResources = codegen.NewStringSet(util.SortedKeys(versions[*trackingPtr])...)
	} else if !moduleCuration.Explicit {
		// Take the latest version as the new tracking version
		maxVersion := maxKey(latestVersions)
		if maxVersion != nil {
			trackingPtr = maxVersion
			// Capture all the resources in the version we're tracking ready to find any resources which aren't in this version
			trackingResources = codegen.NewStringSet(latestVersions[*trackingPtr]...)
		}
	}

	sortedVersions := keys(versions)
	openapi.SortApiVersions(sortedVersions)

	existingAdditions := map[openapi.ResourceName]openapi.ApiVersion{}
	if existing.Additions != nil {
		existingAdditions = *existing.Additions
	}
	additions := map[openapi.ResourceName]openapi.ApiVersion{}
	// Loop through every version in order, skipping excluded and private versions
	// and overwriting additions from previous versions.
	for _, apiVersion := range sortedVersions {
		definitions := versions[apiVersion]
		if !moduleCuration.Explicit && (trackingPtr == nil || apiVersion == *trackingPtr) {
			continue
		}
		for _, definitionName := range util.SortedKeys(definitions) {
			isExcluded, exclusionErr := moduleCuration.IsExcluded(definitionName, apiVersion)
			if exclusionErr != nil {
				exclusionErrors = append(exclusionErrors, ExclusionError{
					ModuleName:   b.moduleName,
					ResourceName: definitionName,
					Detail:       exclusionErr.Error(),
				})
			}
			if isExcluded || openapi.IsPrivate(string(apiVersion)) {
				continue
			}
			if existingVersion, ok := existingAdditions[definitionName]; ok {
				additions[definitionName] = existingVersion
				continue
			}
			// TODO: Use the original AZ namespace rather that our adjusted module name
			if !b.activeVersionChecker.HasProviderVersion(string(b.moduleName), apiVersion) {
				// Don't add if not marked as live
				continue
			}
			if !trackingResources.Has(definitionName) {
				additions[definitionName] = apiVersion
			}
		}
	}

	if len(additions) > 0 {
		additionsPtr = &additions
	}

	return ModuleSpec{
		Tracking:        trackingPtr,
		Additions:       additionsPtr,
		ExclusionErrors: exclusionErrors,
	}
}

// Returns each resource's latest version grouped by version
func (b moduleSpecBuilder) findLatestVersions(versions VersionResources, curations moduleCuration) map[openapi.ApiVersion][]openapi.ResourceName {
	latestResourceVersions := b.findLatestResourceVersions(versions, curations)
	latestVersions := map[openapi.ApiVersion][]openapi.ResourceName{}
	for resourceName, version := range latestResourceVersions {
		latestVersions[version] = append(latestVersions[version], resourceName)
	}
	for _, resourceNames := range latestVersions {
		sort.Strings(resourceNames)
	}
	return latestVersions
}

func (b moduleSpecBuilder) findLatestResourceVersions(versions VersionResources, curations moduleCuration) map[openapi.ResourceName]openapi.ApiVersion {
	candidateVersions := b.filterCandidateVersions(versions, curations.Preview)
	candidates := filterVersionResources(versions, candidateVersions)
	minimalVersionSet := findMinimalVersionSet(candidates)
	minimalVersions := filterVersionResources(candidates, minimalVersionSet)

	var orderedVersions []openapi.ApiVersion
	for apiVersion := range minimalVersions {
		orderedVersions = append(orderedVersions, apiVersion)
	}

	openapi.SortApiVersions(orderedVersions)

	latestResourceVersions := map[openapi.ResourceName]openapi.ApiVersion{}
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		version := orderedVersions[i]
		definitions := minimalVersions[version]
		for _, definitionName := range util.SortedKeys(definitions) {
			// Only add if not already exists
			if _, ok := latestResourceVersions[definitionName]; !ok {
				latestResourceVersions[definitionName] = version
			}
		}
	}

	return latestResourceVersions
}

func filterVersionResources(versions VersionResources, filter []openapi.ApiVersion) VersionResources {
	filtered := VersionResources{}
	filterMap := collections.NewOrderableSet(filter...)

	for apiVersion, resourceNames := range versions {
		if filterMap.Has(apiVersion) {
			filtered[apiVersion] = resourceNames
		}
	}
	return filtered
}

// filterCandidateVersions returns a sorted list of versions which are the best upgrade options,
// removing versions which have now been superseded by a newer or more stable version.
func (b moduleSpecBuilder) filterCandidateVersions(versions VersionResources, previewPolicy string) []openapi.ApiVersion {
	orderedVersions := make([]openapi.ApiVersion, 0, len(versions))
	for _, version := range keys(versions) {
		if version != "" { // Ignore default version placeholders
			orderedVersions = append(orderedVersions, version)
		}
	}
	openapi.SortApiVersions(orderedVersions)
	liveOrderedVersions := []openapi.ApiVersion{}
	for _, version := range orderedVersions {
		// TODO: Use the original AZ namespace rather that our adjusted module name
		if b.activeVersionChecker.HasProviderVersion(string(b.moduleName), version) {
			liveOrderedVersions = append(liveOrderedVersions, version)
		}
	}
	// If there are live versions, consider only those.
	if len(liveOrderedVersions) > 0 {
		orderedVersions = liveOrderedVersions
	}

	candidateVersions := collections.NewOrderableSet[openapi.ApiVersion]()
	hasFutureStableVersion := false
	// Start with newest and work backwards
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		version := orderedVersions[i]
		if openapi.IsPrivate(string(version)) {
			continue
		}
		if previewPolicy == PreviewExclude && openapi.IsPreview(string(version)) {
			continue
		}
		if !openapi.IsPreview(string(version)) || previewPolicy == PreviewPrefer {
			candidateVersions.Add(version)
			hasFutureStableVersion = true
			continue
		}
		if hasFutureStableVersion {
			continue
		}

		// If no more to iterate through, just add
		if i == 0 || !b.containsRecentStable(orderedVersions[:i], version) {
			candidateVersions.Add(version)
			continue
		}

		isOld, err := isMoreThanOneYearOld(version)
		if err != nil {
			panic(fmt.Errorf("error checking version age: %s", err))
		}
		if isOld {
			candidateVersions.Add(version)
		}
	}

	return candidateVersions.SortedValues()
}

func isMoreThanOneYearOld(version openapi.ApiVersion) (bool, error) {
	asTime, err := openapi.ApiVersionToDate(version)
	if err != nil {
		return false, fmt.Errorf("failed parsing version as date: %s", version)
	}
	diff := time.Since(asTime)
	// Round up to 366 to account for leap years
	return diff.Hours() > 24*366, nil
}

func (b moduleSpecBuilder) containsRecentStable(orderedVersions []openapi.ApiVersion, comparisonVersion openapi.ApiVersion) bool {
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		previousVersion := orderedVersions[i]
		// Only looking for recent stable versions
		if openapi.IsPreview(string(previousVersion)) {
			continue
		}
		// Check if stable version is also recent
		timeDiff, err := timeBetweenVersions(previousVersion, comparisonVersion)
		if err != nil {
			panic(errors.Wrapf(err, "failed parsing version as date for %s: %s or %s", b.moduleName, comparisonVersion, previousVersion))
		}
		const maxRecentVersionTimeInHours = 24 * 366
		isRecent := timeDiff.Hours() <= maxRecentVersionTimeInHours
		return isRecent
	}
	return false
}

// findMinimalVersionSet returns the minimum set of versions required to produce all resources
func findMinimalVersionSet(versions VersionResources) []openapi.ApiVersion {
	// TODO: Consider using openapi.SortApiVersions
	orderedVersions := util.SortedKeys(versions)

	latestResourceVersions := map[openapi.ResourceName]openapi.ApiVersion{}
	for _, version := range orderedVersions {
		definitions := versions[version]
		for _, definitionName := range util.SortedKeys(definitions) {
			latestVersion := latestResourceVersions[definitionName]
			// TODO: Consider using openapi.CompareApiVersions
			if version > latestVersion {
				latestResourceVersions[definitionName] = version
			}
		}
	}

	minimalVersions := collections.NewOrderableSet[openapi.ApiVersion]()
	for _, apiVersion := range latestResourceVersions {
		minimalVersions.Add(apiVersion)
	}

	result := minimalVersions.SortedValues()
	return result
}

func timeBetweenVersions(from, to openapi.ApiVersion) (diff time.Duration, err error) {
	if len(from) < 10 || len(to) < 10 {
		return diff, fmt.Errorf("invalid version string")
	}
	fromTime, err := time.Parse("2006-01-02", string(from)[:10])
	if err != nil {
		return diff, err
	}
	toTime, err := time.Parse("2006-01-02", string(to)[:10])
	if err != nil {
		return diff, err
	}
	return toTime.Sub(fromTime), err
}

func maxKey[K cmp.Ordered, V any](m map[K]V) *K {
	keys := keys(m)
	if len(keys) == 0 {
		return nil
	}
	max := slices.Max(keys)
	return &max
}

// Returns an unordered slice of keys from a map
func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
