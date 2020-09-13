// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/debug"
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

var (
	languages        string
	generateSchema   *bool
	generateExamples *bool
	generateSDK      *bool
	version          *string
)

func init() {
	flag.StringVar(&languages, "languages", "", "comma seperated list of languages: nodejs,go,python,dotnet,schema")
	generateSchema = flag.Bool("schema", false, "emit schema as well")
	generateExamples = flag.Bool("examples", false, "emit API examples as well")
	generateSDK = flag.Bool("sdk", false, "emit SDK")
	debug.Debug = flag.Bool("debug", false, "enable debug logging")
	version = flag.String("version", "", "version to stamp in schema")
}

func main() {
	flag.Parse()
	if flag.NArg() == 1 {
		languages = flag.Arg(0)
	}

	azureProviders := openapi.Providers()

	result, err := gen.PulumiSchema(azureProviders)
	if err != nil {
		panic(err)
	}

	handleSchema := func() {
		outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azurerm")
		if err = emitSchema(result.PkgSpec, result.DocsSpec, *version, outdir); err != nil {
			panic(err)
		}
		// Also, emit the resource metadata and embeddable schema for the provider.
		if err = emitMetadata(result.Metadata, outdir); err != nil {
			panic(err)
		}
	}

	langs := strings.Split(languages, ",")

	if *generateExamples {
		// Note - Requires a provider executable in PATH
		err = gen.Examples(result.PkgSpec, result.Metadata, result.Examples, langs)
		if err != nil {
			panic(err)
		}
		// We must rerender the schema after the examples are generated.
		*generateSchema = true
	}

	// Generate SDK after examples so they get rendered as part of the SDK
	if *generateSDK {
		pkgSpec := result.PkgSpec
		for _, language := range langs {
			outdir := path.Join(".", "sdk", language)
			log.Printf("Generating SDK for language: %s", language)
			pkgSpec.Version = *version
			err = emitPackage(pkgSpec, language, outdir)
			if err != nil {
				panic(err)
			}
		}
	}

	if *generateSchema {
		handleSchema()
	}
}

// emitSchema writes the Pulumi schema JSON to the 'schema.json' file in the given directory.
func emitSchema(pkgSpec *schema.PackageSpec, docSpec *schema.PackageSpec, version, outDir string) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	docsJSON, err := json.MarshalIndent(docSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema for docs")
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

	if err := emitFile(outDir, "schema-docs.json", docsJSON); err != nil {
		return err
	}
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
