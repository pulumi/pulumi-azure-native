package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
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

	return writeProviderVersionSummary(providers, path.Join(outputDir, "spec.json"))
}

func writeVersion(curatedVersion openapi.CuratedVersion, outputDir string, version int) error {
	defaultVersions := formatCuratedVersion(curatedVersion)
	outputPath := path.Join(outputDir, fmt.Sprintf("v%v.json", version))
	return emitJson(outputPath, defaultVersions)
}

type ProviderResourceDefaultVersions = map[string]map[string]string

func formatCuratedVersion(curatedVersion openapi.CuratedVersion) ProviderResourceDefaultVersions {
	providerResourceDefaultVersions := make(ProviderResourceDefaultVersions)
	for providerName, resources := range curatedVersion {
		resourceVersions := make(map[string]string)
		for k, resource := range resources.All() {
			resourceVersions[k] = resource.Swagger.Info.Version
		}
		providerResourceDefaultVersions[providerName] = resourceVersions
	}
	return providerResourceDefaultVersions
}

func writeProviderVersionSummary(providerVersions openapi.AzureProviders, outputPath string) error {
	formatted := formatProviderVersions(providerVersions)
	return emitJson(outputPath, formatted)
}

func formatProviderVersions(providerVersions openapi.AzureProviders) map[string][]string {
	formatted := map[string][]string{}
	for name, versions := range providerVersions {
		formattedVersions := make([]string, 0, len(versions))
		for version, _ := range versions {
			formattedVersions = append(formattedVersions, version)
		}
		sort.Strings(formattedVersions)
		formatted[name] = formattedVersions
	}
	return formatted
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
