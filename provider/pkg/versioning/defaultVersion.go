package versioning

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type ProviderSpec struct {
	// Tracking is a single API version from which updates will be included
	Tracking *openapi.ApiVersion `yaml:"Tracking,omitempty"`
	// Additions are specific resource versions to be included. These *must not* overlap with any resources from the tracking version
	Additions *map[openapi.ResourceName]openapi.ApiVersion `yaml:"Additions,omitempty"`
}

type DefaultConfig map[openapi.ProviderName]ProviderSpec

// BuildDefaultConfig calculates a config from available API versions
func BuildDefaultConfig(spec SpecVersions) DefaultConfig {
	specs := DefaultConfig{}
	for providerName, versionResources := range spec {
		specs[providerName] = buildSpec(versionResources)
	}
	return specs
}

// ReadDefaultConfig parses a default config from a JSON file
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

func DefaultConfigToCuratedVersion(spec SpecVersions, defaultConfig DefaultConfig) (openapi.CuratedVersion, error) {
	curatedVersion := openapi.CuratedVersion{}
	for providerName, providerSpec := range defaultConfig {
		definitions := map[openapi.DefinitionName]openapi.ApiVersion{}
		if providerSpec.Tracking != nil {
			for _, resourceName := range spec[providerName][*providerSpec.Tracking] {
				definitions[resourceName] = *providerSpec.Tracking
			}
		}
		if providerSpec.Additions != nil {
			for resourceName, apiVersion := range *providerSpec.Additions {
				if existingVersion, ok := definitions[resourceName]; ok {
					fmt.Printf("duplicate resource %s from %s and %s\n", resourceName, apiVersion, existingVersion)
				}
				definitions[resourceName] = apiVersion
			}
		}
		curatedVersion[providerName] = definitions
	}
	return curatedVersion, nil
}

func buildSpec(versions VersionResources) ProviderSpec {
	latestVersions := findLatestVersions(versions)
	switch len(latestVersions) {
	case 0:
		return ProviderSpec{}
	case 1:
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
		if apiVersion == maxVersion {
			continue
		}
		for _, resourceName := range resources {
			additions[resourceName] = apiVersion
		}
	}

	return ProviderSpec{
		Tracking:  &maxVersion,
		Additions: &additions,
	}
}

func findLatestVersions(versions VersionResources) map[openapi.ApiVersion][]openapi.ResourceName {
	latestResourceVersions := findLatestResourceVersions(versions)
	latestVersions := map[openapi.ApiVersion][]openapi.ResourceName{}
	for resourceName, version := range latestResourceVersions {
		latestVersions[version] = append(latestVersions[version], resourceName)
	}
	for _, resourceNames := range latestVersions {
		sort.Strings(resourceNames)
	}
	return latestVersions
}

func findLatestResourceVersions(versions VersionResources) map[openapi.ResourceName]openapi.ApiVersion {
	candidateVersions := filterCandidateVersions(versions)
	candidates := filterVersionResources(versions, candidateVersions)
	minimalVersionSet := findMinimalVersionSet(candidates)
	minimalVersions := filterVersionResources(candidates, minimalVersionSet)

	var orderedVersions []openapi.ApiVersion
	for apiVersion := range minimalVersions {
		orderedVersions = append(orderedVersions, apiVersion)
	}
	sort.Strings(orderedVersions)

	latestResourceVersions := map[openapi.ResourceName]openapi.ApiVersion{}
	// Descending order - newest first
	for i := len(orderedVersions) - 1; i >= 0; i-- {
		version := orderedVersions[i]
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

func filterCandidateVersions(versions VersionResources) []openapi.ApiVersion {
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
		if !openapi.IsPreview(version) {
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
