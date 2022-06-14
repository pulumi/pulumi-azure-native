package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/providerlist"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"os"
	"path"
	"sort"
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

	err = writeProviderVersionSummary(providers, path.Join(outputDir, "spec-providers.json"))
	if err != nil {
		return err
	}

	err = writeResourceVersionSummary(providers, path.Join(outputDir, "spec-resources.json"))
	if err != nil {
		return err
	}

	v1, err := openapi.CalculateProviderDefaults(providers)
	if err != nil {
		return err
	}

	deprecated := openapi.FindOlderVersions(providers, v1)

	err = writeVersion(v1, outputDir, 1)
	if err != nil {
		return err
	}

	err = writeProviderVersionSummary(deprecated, path.Join(outputDir, "deprecated.json"))
	if err != nil {
		return err
	}

	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return err
	}

	err = writeActiveVersions(path.Join(outputDir, "active.json"), activePathVersions)
	if err != nil {
		return err
	}

	return nil
}

func writeVersion(curatedVersion openapi.CuratedVersion, outputDir string, version int) error {
	outputPath := path.Join(outputDir, fmt.Sprintf("v%v.json", version))
	return emitJson(outputPath, curatedVersion)
}

func writeProviderVersionSummary(providerVersions openapi.AzureProviders, outputPath string) error {
	formatted := formatProviderVersions(providerVersions)
	return emitJson(outputPath, formatted)
}

func formatProviderVersions(providerVersions openapi.AzureProviders) map[string][]string {
	formatted := map[string][]string{}
	for name, versions := range providerVersions {
		formattedVersions := make([]string, 0, len(versions))
		for _, resources := range versions {
			for _, spec := range resources.All() {
				formattedVersions = append(formattedVersions, spec.Swagger.Info.Version)
				break
			}
		}
		sort.Strings(formattedVersions)
		formatted[name] = formattedVersions
	}
	return formatted
}

func writeResourceVersionSummary(providerVersions openapi.AzureProviders, outputPath string) error {
	formatted := formatResourceVersions(providerVersions)
	return emitJson(outputPath, formatted)
}

func formatResourceVersions(providerVersions openapi.AzureProviders) map[string]map[string][]string {
	formatted := map[string]map[string][]string{}
	for providerName, versions := range providerVersions {
		resourceVersions := map[string]codegen.StringSet{}

		for _, resources := range versions {
			for resourceName, spec := range resources.All() {
				var versionSet codegen.StringSet
				var ok bool
				if versionSet, ok = resourceVersions[resourceName]; !ok {
					versionSet = codegen.NewStringSet()
					resourceVersions[resourceName] = versionSet
				}
				versionSet.Add(spec.Swagger.Info.Version)
			}
		}

		formattedResources := map[string][]string{}
		for resourceName, versions := range resourceVersions {
			formattedResources[resourceName] = versions.SortedValues()
		}
		formatted[providerName] = formattedResources
	}
	return formatted
}

func writeActiveVersions(outputPath string, activePathVersions providerlist.ProviderPathVersions) error {
	formatted := map[string]map[string][]string{}
	for providerName, paths := range activePathVersions {
		formattedProvider := map[string][]string{}
		for resourcePath, versions := range paths {
			formattedProvider[resourcePath] = versions.SortedValues()
		}
		formatted[providerName] = formattedProvider
	}
	return emitJson(outputPath, formatted)
}

func emitJson(outputPath string, data interface{}) error {
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
