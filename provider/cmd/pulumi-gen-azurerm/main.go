package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/gen"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/pkg/provider"
	dotnetgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	languages := os.Args[1]

	swaggerSpecLocations, err := provider.SwaggerLocations()
	if err != nil {
		panic(err)
	}

	var specs []*openapi.Spec
	for _, location := range swaggerSpecLocations {
		spec, err := openapi.NewSpec(location)
		if err != nil {
			panic(err)
		}

		specs = append(specs, spec)
	}

	pkgSpec, meta, err := gen.PulumiSchema(specs)
	if err != nil {
		panic(err)
	}

	for _, language := range strings.Split(languages, ",") {
		outdir := path.Join(".", "sdk", language)
		switch language {
		case "schema":
			err = emitSchema(pkgSpec, outdir)
			if err == nil {
				// Also, emit the resource metadata for the provider.
				err = emitMetadata(meta, outdir)
			}
		default:
			err = emitPackage(pkgSpec, language, outdir)
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

func emitMetadata(metadata *provider.AzureApiMetadata, outDir string) error {
	raw, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	formatted, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	err = ioutil.WriteFile("./provider/pkg/provider/metadata.go", []byte(fmt.Sprintf(`package provider
var azureApiResources = %#v
`, raw)), 0600)
	if err != nil {
		return err
	}

	return emitFile(outDir, "metadata.json", formatted)
}

func generate(ppkg *schema.Package, language string) (map[string][]byte, error) {
	toolDescription := "the Pulumi SDK Generator"
	extraFiles := map[string][]byte{}
	switch language {
	case "nodejs":
		return nodejsgen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "python":
		return pythongen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "go":
		return gogen.GeneratePackage(toolDescription, ppkg)
	case "dotnet":
		return dotnetgen.GeneratePackage(toolDescription, ppkg, extraFiles)
	}

	return nil, errors.Errorf("unknown language '%s'", language)
}

// emitPackage emits an entire package pack into the configured output directory with the configured settings.
func emitPackage(pkgSpec *schema.PackageSpec, language, outDir string) error {
	ppkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return errors.Wrap(err, "reading schema")
	}

	files, err := generate(ppkg, language)
	if err != nil {
		return errors.Wrapf(err, "generating %s package", language)
	}

	for f, contents := range files {
		if err := emitFile(outDir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
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
