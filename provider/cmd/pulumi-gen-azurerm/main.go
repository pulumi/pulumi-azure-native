// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-azurerm/provider/pkg/gen"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	dotnetgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func main() {
	languages := os.Args[1]

	version := ""
	if len(os.Args) == 3 {
		version = os.Args[2]
	}

	azureProviders := openapi.AllVersions()

	pkgSpec, meta, _, err := gen.PulumiSchema(azureProviders)
	if err != nil {
		panic(err)
	}

	for _, language := range strings.Split(languages, ",") {
		switch language {
		case "schema":
			outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azurerm")
			if err = emitSchema(*pkgSpec, version, outdir); err != nil {
				break
			}
			// Also, emit the resource metadata for the provider.
			err = emitMetadata(meta, outdir)
		case "docs":
			outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azurerm")
			docsProviders := openapi.SingleVersion(azureProviders)
			docsPkgSpec, _, resExamples, err := gen.PulumiSchema(docsProviders)
			if err != nil {
				break
			}
			err = gen.Examples(docsPkgSpec, meta, resExamples, []string{"nodejs", "dotnet", "python", "go"})
			if err != nil {
				break
			}
			err = emitDocsSchema(docsPkgSpec, version, outdir)
		default:
			outdir := path.Join(".", "sdk", language)
			pkgSpec.Version = version
			err = emitPackage(pkgSpec, language, outdir)
		}
		if err != nil {
			panic(err)
		}
	}
}

// emitSchema writes the Pulumi schema JSON to the 'schema.json' file in the given directory.
func emitSchema(pkgSpec schema.PackageSpec, version, outDir string) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	// Ensure the spec is stamped with a version.
	pkgSpec.Version = version

	compressedSchema := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedSchema)
	err = json.NewEncoder(compressedWriter).Encode(pkgSpec)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}
	if err = compressedWriter.Close(); err != nil {
		return err
	}

	err = emitFile(outDir, "schema.go", []byte(fmt.Sprintf(`package main
var pulumiSchema = %#v
`, compressedSchema.Bytes())))
	if err != nil {
		return errors.Wrap(err, "saving metadata")
	}

	return emitFile(outDir, "schema-full.json", schemaJSON)
}

// emitDocsSchema writes the Pulumi schema JSON to the 'schema-docs.json' file in the given directory.
func emitDocsSchema(pkgSpec *schema.PackageSpec, version, outDir string) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	// Ensure the spec is stamped with a version.
	pkgSpec.Version = version

	return emitFile(outDir, "schema.json", schemaJSON)
}

func emitMetadata(metadata *provider.AzureAPIMetadata, outDir string) error {
	compressedMeta := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedMeta)
	err := json.NewEncoder(compressedWriter).Encode(metadata)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	if err = compressedWriter.Close(); err != nil {
		return err
	}

	formatted, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	err = emitFile(outDir, "metadata.go", []byte(fmt.Sprintf(`package main
var azureApiResources = %#v
`, compressedMeta.Bytes())))
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
