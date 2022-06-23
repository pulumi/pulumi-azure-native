package versioning

import (
	"encoding/json"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type ProviderSpec struct {
	NoVersion     *bool                                          `json:"NoVersion,omitempty"`
	SingleVersion *openapi.ApiVersion                            `json:"SingleVersion,omitempty"`
	Rollup        *map[openapi.ApiVersion][]openapi.ResourceName `json:"Rollup,omitempty"`
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
	err = json.Unmarshal(byteValue, &curatedVersion)
	return curatedVersion, err
}

func DefaultConfigToCuratedVersion(spec SpecVersions, defaultConfig DefaultConfig) openapi.CuratedVersion {
	curatedVersion := openapi.CuratedVersion{}
	for providerName, providerSpec := range defaultConfig {
		if providerSpec.NoVersion != nil {
			continue
		}
		definitions := map[openapi.DefinitionName]openapi.ApiVersion{}
		if providerSpec.SingleVersion != nil {
			for _, resourceName := range spec[providerName][*providerSpec.SingleVersion] {
				definitions[resourceName] = *providerSpec.SingleVersion
			}
		} else if providerSpec.Rollup != nil {
			maxVersion := ""
			for apiVersion := range *providerSpec.Rollup {
				if apiVersion > maxVersion {
					maxVersion = apiVersion
				}
			}
			for apiVersion, resourceNames := range *providerSpec.Rollup {
				if apiVersion == maxVersion {
					// Use all resource from the latest version
					for _, resourceName := range spec[providerName][apiVersion] {
						definitions[resourceName] = apiVersion
					}
				} else {
					for _, resourceName := range resourceNames {
						// Don't overwrite if resource is now included in latest version
						if _, ok := definitions[resourceName]; !ok {
							definitions[resourceName] = apiVersion
						}
					}
				}
			}
		}
		curatedVersion[providerName] = definitions
	}
	return curatedVersion
}

func buildSpec(versions VersionResources) ProviderSpec {
	noVersion := true
	latestVersions := findLatestVersions(versions)
	switch len(latestVersions) {
	case 0:
		return ProviderSpec{
			NoVersion: &noVersion,
		}
	case 1:
		for apiVersion := range latestVersions {
			return ProviderSpec{
				SingleVersion: &apiVersion,
			}
		}
	}
	return ProviderSpec{
		Rollup: &latestVersions,
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
			if latestVersion > version {
				continue
			}
			if latestVersion == "" {
				latestResourceVersions[resourceName] = version
				continue
			}

			isToStable := !openapi.IsPreview(version)
			isFromPreview := openapi.IsPreview(latestVersion)
			timeSinceLastVersion, err := timeBetweenVersions(latestVersion, version)
			if err != nil {
				println("unable to parse version as time", latestVersion, version, err)
				continue
			}
			isLatestOld := timeSinceLastVersion.Hours() > (2 * 365 * 24)
			if isToStable || isFromPreview || isLatestOld {
				latestResourceVersions[resourceName] = version
			}
		}
	}
	return latestResourceVersions
}

func timeBetweenVersions(from, to openapi.ApiVersion) (diff time.Duration, err error) {
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
