package versioning

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"gopkg.in/yaml.v3"
)

type ProviderSpec struct {
	// Tracking is a single API version from which updates will be included
	Tracking *openapi.ApiVersion `yaml:"tracking,omitempty"`
	// Additions are specific resource versions to be included. These *must not* overlap with any resources from the tracking version
	Additions *map[openapi.ResourceName]openapi.ApiVersion `yaml:"additions,omitempty"`
}

type DefaultConfig map[openapi.ProviderName]ProviderSpec

// BuildDefaultConfig calculates a config from available API versions
func BuildDefaultConfig(spec SpecVersions, curations Curations, existingConfig DefaultConfig) DefaultConfig {
	specs := DefaultConfig{}
	for providerName, versionResources := range spec {
		existing := existingConfig[providerName]
		specs[providerName] = buildSpec(providerName, versionResources, curations, existing)
	}
	return specs
}

type CurationViolation struct {
	Provider openapi.ProviderName
	Detail   string
}

func ValidateDefaultConfig(config DefaultConfig, curations Curations) []CurationViolation {
	var violations []CurationViolation
	for provider, providerSpec := range config {
		providerCuration := curations[provider]
		expectedTracking := providerCuration.ExpectTracking
		if expectedTracking == "" {
			expectedTracking = "stable"
		}
		var actualTracking string
		if providerSpec.Tracking != nil && openapi.IsPreview(*providerSpec.Tracking) {
			actualTracking = "preview"
		} else {
			actualTracking = "stable"
		}
		if expectedTracking != actualTracking {
			violations = append(violations, CurationViolation{
				Provider: provider,
				Detail:   fmt.Sprintf("expected tracking %s but found %s", expectedTracking, actualTracking),
			})
		}
		hasAdditions := providerSpec.Additions != nil && len(*providerSpec.Additions) > 0
		if providerCuration.ExpectAdditions != hasAdditions && !providerCuration.Explicit {
			var detail string
			if providerCuration.ExpectAdditions {
				detail = "expected additions but found none"
			} else {
				detail = "expected no additions but found some"
			}
			violations = append(violations, CurationViolation{
				Provider: provider,
				Detail:   detail,
			})
		}
	}
	// sort violations by provider name
	sort.Slice(violations, func(i, j int) bool {
		return strings.Compare(violations[i].Provider, violations[j].Provider) < 0
	})
	return violations
}

func PrintViolationsAsWarnings(curationsPath string, violations []CurationViolation) {
	if len(violations) > 0 {
		// Set colour to yellow
		fmt.Printf("\x1b[33m")
		fmt.Printf("Warning: %d curation violations found in %s:\n", len(violations), curationsPath)
		for _, v := range violations {
			fmt.Printf("  %s: %s\n", v.Provider, v.Detail)
		}
		// Reset colour
		fmt.Printf("\x1b[0m")
	}
}

// ReadDefaultConfig parses a default config from a YAML file
func ReadDefaultConfig(path string) (DefaultConfig, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var curatedVersion DefaultConfig
	err = yaml.Unmarshal(byteValue, &curatedVersion)
	return curatedVersion, err
}

func DefaultConfigToDefaultVersionLock(spec SpecVersions, defaultConfig DefaultConfig) (openapi.DefaultVersionLock, error) {
	var err error
	defaultVersionLock := openapi.DefaultVersionLock{}
	for providerName, versionResources := range spec {
		definitions := map[openapi.DefinitionName]openapi.ApiVersion{}
		providerSpec, ok := defaultConfig[providerName]
		if !ok {
			if len(versionResources) > 0 {
				var versions []string
				for version := range versionResources {
					if version != "" {
						versions = append(versions, version)
					}
				}
				if len(versions) > 0 {
					err = multierror.Append(err, fmt.Errorf("no version specified for %s, available versions: %s", providerName, strings.Join(versions, ", ")))
				}
				continue
			}
		}

		if providerSpec.Tracking != nil {
			for _, resourceName := range versionResources[*providerSpec.Tracking] {
				definitions[resourceName] = *providerSpec.Tracking
			}
		}
		if providerSpec.Additions != nil {
			for resourceName, apiVersion := range *providerSpec.Additions {
				if existingVersion, ok := definitions[resourceName]; ok {
					err = multierror.Append(err, fmt.Errorf("duplicate resource %s:%s from %s and %s", providerName, resourceName, apiVersion, existingVersion))
				} else {
					definitions[resourceName] = apiVersion
				}
			}
		}
		defaultVersionLock[providerName] = definitions
	}
	return defaultVersionLock, multierror.Flatten(err)
}

func buildSpec(providerName string, versions VersionResources, curations Curations, existing ProviderSpec) ProviderSpec {
	var additionsPtr *map[string]string
	var trackingPtr *string

	providerCuration := curations[providerName]
	latestVersions := findLatestVersions(versions, providerCuration)

	if len(latestVersions) == 0 {
		return ProviderSpec{}
	}

	// If multiple versions required, track the latest and include additional resources from previous versions
	var trackingResources codegen.StringSet
	if existing.Tracking != nil {
		trackingPtr = existing.Tracking
		trackingResources = codegen.NewStringSet(versions[*trackingPtr]...)
	} else if !providerCuration.Explicit {
		maxVersion := maxKey(latestVersions)
		if maxVersion != nil {
			trackingPtr = maxVersion
			trackingResources = codegen.NewStringSet(latestVersions[*trackingPtr]...)
		}
	}

	existingAdditions := map[openapi.ResourceName]openapi.ApiVersion{}
	if existing.Additions != nil {
		existingAdditions = *existing.Additions
	}
	additions := map[openapi.ResourceName]openapi.ApiVersion{}
	for apiVersion, resources := range latestVersions {
		if !providerCuration.Explicit && (trackingPtr == nil || apiVersion == *trackingPtr) {
			continue
		}
		for _, resourceName := range resources {
			isExcluded, exclusionErr := providerCuration.IsExcluded(resourceName, apiVersion)
			if exclusionErr != nil {
				fmt.Printf("Error checking exclusion for %s/%s: %s", providerName, resourceName, exclusionErr)
			}
			if isExcluded {
				continue
			} else if version, ok := existingAdditions[resourceName]; ok {
				additions[resourceName] = version
			} else if !trackingResources.Has(resourceName) {
				additions[resourceName] = apiVersion
			}
		}
	}

	if len(additions) > 0 {
		additionsPtr = &additions
	}

	return ProviderSpec{
		Tracking:  trackingPtr,
		Additions: additionsPtr,
	}
}

func findLatestVersions(versions VersionResources, curations providerCuration) map[openapi.ApiVersion][]openapi.ResourceName {
	latestResourceVersions := findLatestResourceVersions(versions, curations)
	latestVersions := map[openapi.ApiVersion][]openapi.ResourceName{}
	for resourceName, version := range latestResourceVersions {
		latestVersions[version] = append(latestVersions[version], resourceName)
	}
	for _, resourceNames := range latestVersions {
		sort.Strings(resourceNames)
	}
	return latestVersions
}

func findLatestResourceVersions(versions VersionResources, curations providerCuration) map[openapi.ResourceName]openapi.ApiVersion {
	candidateVersions := filterCandidateVersions(versions, curations.Preview)
	candidates := filterVersionResources(versions, candidateVersions)
	minimalVersionSet := findMinimalVersionSet(candidates)
	minimalVersions := filterVersionResources(candidates, minimalVersionSet)

	var orderedVersions []openapi.ApiVersion
	for apiVersion := range minimalVersions {
		orderedVersions = append(orderedVersions, apiVersion)
	}

	// Descending order - newest first
	sort.Sort(sort.Reverse(sort.StringSlice(orderedVersions)))

	latestResourceVersions := map[openapi.ResourceName]openapi.ApiVersion{}
	for _, version := range orderedVersions {
		resources := minimalVersions[version]
		for _, resourceName := range resources {
			// Only add if not already exists
			if _, ok := latestResourceVersions[resourceName]; !ok {
				latestResourceVersions[resourceName] = version
			}
		}
	}

	return latestResourceVersions
}

func filterVersionResources(versions VersionResources, filter []openapi.ApiVersion) VersionResources {
	filtered := VersionResources{}
	filterSet := codegen.NewStringSet(filter...)
	for apiVersion, resourceNames := range versions {
		if filterSet.Has(apiVersion) {
			filtered[apiVersion] = resourceNames
		}
	}
	return filtered
}

func filterCandidateVersions(versions VersionResources, previewPolicy string) []openapi.ApiVersion {
	orderedVersions := make([]openapi.ApiVersion, 0, len(versions))
	for version := range versions {
		if version != "" {
			orderedVersions = append(orderedVersions, version)
		}
	}
	sort.Strings(orderedVersions)

	candidateVersions := codegen.NewStringSet()
	hasFutureStableVersion := false
	// Start with newest and work backwards
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		version := orderedVersions[i]
		if openapi.IsPrivate(version) {
			continue
		}
		if previewPolicy == PreviewExclude && openapi.IsPreview(version) {
			continue
		}
		if !openapi.IsPreview(version) || previewPolicy == PreviewPrefer {
			candidateVersions.Add(version)
			hasFutureStableVersion = true
			continue
		}
		if hasFutureStableVersion {
			continue
		}

		// If no more to iterate through, just add
		if i == 0 || !containsRecentStable(orderedVersions[:i], version) {
			candidateVersions.Add(version)
		}
	}

	return candidateVersions.SortedValues()
}

func containsRecentStable(orderedVersions []openapi.ApiVersion, comparisonVersion openapi.ApiVersion) bool {
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		previousVersion := orderedVersions[i]
		// Only looking for recent stable versions
		if openapi.IsPreview(previousVersion) {
			continue
		}
		// Check if stable version is also recent
		timeDiff, err := timeBetweenVersions(previousVersion, comparisonVersion)
		if err != nil {
			panic(errors.Wrapf(err, "failed parsing version as date: %s or %s", comparisonVersion, previousVersion))
		}
		const maxRecentVersionTimeInHours = 24 * 366
		isRecent := timeDiff.Hours() <= maxRecentVersionTimeInHours
		return isRecent
	}
	return false
}

// findMinimalVersionSet returns the minimum set of versions required to produce all resources
func findMinimalVersionSet(versions VersionResources) []openapi.ApiVersion {
	orderedVersions := make([]openapi.ApiVersion, 0, len(versions))
	for version := range versions {
		orderedVersions = append(orderedVersions, version)
	}
	sort.Strings(orderedVersions)

	latestResourceVersions := map[openapi.ResourceName]openapi.ApiVersion{}
	for _, version := range orderedVersions {
		resourceNames := versions[version]
		for _, resourceName := range resourceNames {
			latestVersion := latestResourceVersions[resourceName]
			if version > latestVersion {
				latestResourceVersions[resourceName] = version
			}
		}
	}

	minimalVersions := codegen.NewStringSet()
	for _, apiVersion := range latestResourceVersions {
		minimalVersions.Add(apiVersion)
	}

	return minimalVersions.SortedValues()
}

func timeBetweenVersions(from, to openapi.ApiVersion) (diff time.Duration, err error) {
	if len(from) < 10 || len(to) < 10 {
		return diff, fmt.Errorf("invalid version string")
	}
	fromTime, err := time.Parse("2006-01-02", from[:10])
	if err != nil {
		return diff, err
	}
	toTime, err := time.Parse("2006-01-02", to[:10])
	if err != nil {
		return diff, err
	}
	return toTime.Sub(fromTime), err
}

func maxKey[K comparable, V any](m map[K]V) *K {
	keys := keys(m)
	if len(keys) == 0 {
		return nil
	}
	last := keys[len(keys)-1]
	return &last
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
