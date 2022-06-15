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

	specVersions := versioning.FindSpecVersions(providers)

	err = emitJson(path.Join(outputDir, "spec.json"), specVersions)
	if err != nil {
		return err
	}

	err = writeResourceVersionSummary(specVersions, path.Join(outputDir, "spec-resources.json"))
	if err != nil {
		return err
	}

	v1, err := openapi.CalculateProviderDefaults(providers)
	if err != nil {
		return err
	}

	deprecated := versioning.FindOlderVersions(specVersions, v1)

	err = emitJson(path.Join(outputDir, "v1.json"), v1)
	if err != nil {
		return err
	}

	err = emitJson(path.Join(outputDir, "deprecated.json"), deprecated)
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

func writeResourceVersionSummary(providerVersions versioning.SpecVersions, outputPath string) error {
	formatted := versioning.FormatResourceVersions(providerVersions)
	return emitJson(outputPath, formatted)
}

func writeActiveVersions(outputPath string, activePathVersions providerlist.ProviderPathVersions) error {
	formatted := providerlist.FormatProviderPathVersionsJson(activePathVersions)
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
