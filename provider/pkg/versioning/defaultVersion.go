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
func BuildDefaultConfig(spec SpecVersions, curations Curations) DefaultConfig {
	specs := DefaultConfig{}
	for providerName, versionResources := range spec {
		providerCurations := curations[providerName]
		specs[providerName] = buildSpec(providerName, versionResources, providerCurations)
	}
	return specs
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

func DefaultConfigToCuratedVersion(spec SpecVersions, defaultConfig DefaultConfig, curation Curations) (openapi.CuratedVersion, error) {
	var err error
	curatedVersion := openapi.CuratedVersion{}
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
				if !curation.IsExcluded(providerName, resourceName) {
					definitions[resourceName] = *providerSpec.Tracking
				}
			}
		}
		if providerSpec.Additions != nil {
			for resourceName, apiVersion := range *providerSpec.Additions {
				if existingVersion, ok := definitions[resourceName]; ok {
					err = multierror.Append(err, fmt.Errorf("duplicate resource %s:%s from %s and %s", providerName, resourceName, apiVersion, existingVersion))
				} else if !curation.IsExcluded(providerName, resourceName) {
					definitions[resourceName] = apiVersion
				}
			}
		}
		curatedVersion[providerName] = definitions
	}
	return curatedVersion, multierror.Flatten(err)
}

func buildSpec(providerName string, versions VersionResources, curations providerCuration) ProviderSpec {
	latestVersions := findLatestVersions(versions, curations)
	if len(latestVersions) == 0 {
		return ProviderSpec{}
	}
	if len(latestVersions) == 1 && !curations.Explicit {
		// If single apiVersion includes all resources, track it.
		for apiVersion := range latestVersions {
			return ProviderSpec{
				Tracking: &apiVersion,
			}
		}
	}

	// If multiple versions required, track the latest and include additional resources from previous versions
	additions := map[openapi.ResourceName]openapi.ApiVersion{}
	maxVersion := ""
	for apiVersion := range latestVersions {
		if apiVersion > maxVersion {
			maxVersion = apiVersion
		}
	}
	for apiVersion, resources := range latestVersions {
		if !curations.Explicit && apiVersion == maxVersion {
			continue
		}
		for _, resourceName := range resources {
			excludedMaxVersion, hasExclusion := curations.Exclusions[resourceName]
			if hasExclusion {
				if excludedMaxVersion == "*" || excludedMaxVersion >= apiVersion {
					continue
				}
				fmt.Printf("Exclusion %s/%s not applied, version %s is greater than %s", providerName, resourceName, apiVersion, excludedMaxVersion)
			}
			additions[resourceName] = apiVersion
		}
	}
	additionsPtr := &additions
	if len(additions) == 0 {
		additionsPtr = nil
	}
	trackingPtr := &maxVersion
	if curations.Explicit {
		trackingPtr = nil
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
