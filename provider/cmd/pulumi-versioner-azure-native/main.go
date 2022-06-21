package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/versioning"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"os"
	"path"
	"sort"
	"time"
)

func main() {
	outputDir := flag.String("o", "versions", "Output directory")
	flag.Parse()

	err := writeAll(*outputDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeAll(outputDir string) error {
	providers, err := openapi.SpecVersions()
	if err != nil {
		return err
	}

	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return err
	}

	specVersions := versioning.FindSpecVersions(providers)
	specResourceVersions := versioning.FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)
	v1 := openapi.CalculateProviderDefaults(activePathVersions, providers)
	deprecated := versioning.FindDeprecations(specVersions, v1)
	pending := versioning.FindNewerVersions(specVersions, v1)

	return emitJsonFiles(outputDir, map[Filename]Json{
		"spec.json":           specVersions,
		"spec-resources.json": specResourceVersions,
		"v1.json":             v1,
		"deprecated.json":     deprecated,
		"active.json":         activePathVersionsJson,
		"pending.json":        pending,
		"v2-config.json":      buildSpecs(specVersions),
	})
}

type Filename = string
type Json = interface{}

func emitJsonFiles(outDir string, files map[Filename]Json) error {
	for filename, data := range files {
		outPath := path.Join(outDir, filename)
		err := emitJson(outPath, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func emitJson(outputPath string, data Json) error {
	formatted, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshaling JSON")
	}

	return emitFile(outputPath, formatted)
}

func emitFile(outPath string, contents []byte) error {
	if err := tools.EnsureDir(path.Dir(outPath)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(outPath)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}

type ProviderSpec struct {
	NoVersion     *bool                                          `json:"NoVersion,omitempty"`
	SingleVersion *openapi.ApiVersion                            `json:"SingleVersion,omitempty"`
	Rollup        *map[openapi.ApiVersion][]openapi.ResourceName `json:"Rollup,omitempty"`
}

func buildSpecs(spec versioning.SpecVersions) map[openapi.ProviderName]ProviderSpec {
	specs := map[openapi.ProviderName]ProviderSpec{}
	for providerName, versionResources := range spec {
		specs[providerName] = buildSpec(versionResources)
	}
	return specs
}

func buildSpec(versions versioning.VersionResources) ProviderSpec {
	noVersion := true
	latestVersions := findLatestVersions(versions)
	switch len(latestVersions) {
	case 0:
		return ProviderSpec{
			NoVersion: &noVersion,
		}
	case 1:
		for apiVersion, _ := range latestVersions {
			return ProviderSpec{
				SingleVersion: &apiVersion,
			}
		}
	}
	return ProviderSpec{
		Rollup: &latestVersions,
	}
}

func findLatestVersions(versions versioning.VersionResources) map[openapi.ApiVersion][]openapi.ResourceName {
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

func findLatestResourceVersions(versions versioning.VersionResources) map[openapi.ResourceName]openapi.ApiVersion {
	orderedVersions := make([]openapi.ApiVersion, 0, len(versions))
	for version, _ := range versions {
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
