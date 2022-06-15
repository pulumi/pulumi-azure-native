package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hashicorp/go-multierror"
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

	activePathVersions, err := providerlist.ReadProviderList()
	if err != nil {
		return err
	}

	specVersions := versioning.FindSpecVersions(providers)
	specResourceVersions := versioning.FormatResourceVersions(specVersions)
	activePathVersionsJson := providerlist.FormatProviderPathVersionsJson(activePathVersions)

	v1, err := openapi.CalculateProviderDefaults(providers)
	if err != nil {
		return err
	}

	deprecated := versioning.FindOlderVersions(specVersions, v1)

	return emitJsonFiles(outputDir, map[Filename]Json{
		"spec.json":           specVersions,
		"spec-resources.json": specResourceVersions,
		"v1.json":             v1,
		"deprecated.json":     deprecated,
		"active.json":         activePathVersionsJson,
	})
}

type Filename = string
type Json = interface{}

func emitJsonFiles(outDir string, files map[Filename]Json) error {
	var result error
	for filename, data := range files {
		outPath := path.Join(outDir, filename)
		err := emitJson(outPath, data)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
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
