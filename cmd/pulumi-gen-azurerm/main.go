package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/gen"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/pkg/provider"
	nodejsgen "github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"os"
	"path"
	"strings"
)

func main() {
	languages := os.Args[1]

	var specs []*openapi.Spec
	for _, location := range provider.SwaggerLocations() {
		spec, err := openapi.NewSpec(location)
		if err != nil {
			panic(err)
		}

		specs = append(specs, spec)
	}

	pkgSpec, err := gen.PulumiSchema(specs)
	if err != nil {
		panic(err)
	}

	for _, language := range strings.Split(languages, ",") {
		outdir := path.Join(".", "sdk", language)
		switch language {
		case "nodejs":
			err = emitPackage(pkgSpec, outdir)
		case "schema":
			err = emitSchema(pkgSpec, outdir)
		default:
			panic(fmt.Sprintf("Unrecognized language '%s'", language))
		}
	}

	if err != nil {
		panic(err)
	}
}

// emitSchema writes the Pulumi schema JSON to the 'schema.json' file in the given directory.
func emitSchema(pkgSpec *schema.PackageSpec, outDir string) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	return emitFile(outDir, "schema.json", schemaJSON)
}

// emitPackage emits an entire package pack into the configured output directory with the configured settings.
func emitPackage(pkgSpec *schema.PackageSpec, outDir string) error {
	toolDescription := "the Pulumi SDK Generator"
	extraFiles := map[string][]byte{}

	ppkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return errors.Wrap(err, "reading schema")
	}

	sdkGenerators := map[string]func() (map[string][]byte, error){
		"nodejs": func() (map[string][]byte, error) {
			return nodejsgen.GeneratePackage(toolDescription, ppkg, extraFiles)
		},
	}

	for sdkName, generator := range sdkGenerators {
		files, err := generator()
		if err != nil {
			return errors.Wrapf(err, "generating %s package", sdkName)
		}

		for f, contents := range files {
			if err := emitFile(outDir, f, contents); err != nil {
				return errors.Wrapf(err, "emitting file %v", f)
			}
		}
	}

	return nil
}

// emitFile creates a file in a given directory and writes the byte contents to it.
func emitFile(outDir, relPath string, contents []byte) error {
	p := path.Join(outDir, relPath)
	if err := tools.EnsureDir(path.Dir(p)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}
