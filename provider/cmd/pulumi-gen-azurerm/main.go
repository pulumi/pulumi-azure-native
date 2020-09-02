// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

	pkgSpec, meta, err := gen.PulumiSchema(azureProviders)
	if err != nil {
		panic(err)
	}

	handleSchema := func() {
		outdir := path.Join(".", "sdk", "schema")
		if err = emitSchema(*pkgSpec, *version, outdir); err != nil {
			panic(err)
		}
		// Also, emit the resource metadata and embeddable schema for the provider.
		if err = emitMetadata(meta, outdir); err != nil {
			panic(err)
		}
	}

	langs := strings.Split(languages, ",")

	if *generateExamples {
		// Note - Requires a provider executable in PATH
		err = gen.Examples(pkgSpec, meta, langs)
		if err != nil {
			panic(err)
		}
		// We must rerender the schema after the examples are generated.
		*generateSchema = true
	}

	// Generate SDK after examples so they get rendered as part of the SDK
	if *generateSDK {
		for _, language := range langs {
			outdir := path.Join(".", "sdk", language)
			switch language {
			case "schema":
				// back-compat, should use -schema for this instead
				*generateSchema = true
			default:
				log.Printf("Generating SDK for language: %s", language)
				err = emitPackage(pkgSpec, language, outdir)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if *generateSchema {
		handleSchema()
	}
}

func emitMetadata(metadata *provider.AzureApiMetadata, outDir string) error {
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

	err = ioutil.WriteFile("./provider/cmd/pulumi-resource-azurerm/metadata.go", []byte(fmt.Sprintf(`package main
var azureApiResources = %#v
`, compressedMeta.Bytes())), 0600)
	if err != nil {
		return err
	}

	return emitFile(outDir, "metadata.json", formatted)
}

func emitSchema(pkgSpec schema.PackageSpec, version string, outdir string) error {
	// Ensure the spec is stamped with a version.
	pkgSpec.Version = version
	log.Printf("Schema version: %s", pkgSpec.Version)

	compressedSchema := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedSchema)
	formatted, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return err
	}

	enc := json.NewEncoder(compressedWriter)
	if err := enc.Encode(pkgSpec); err != nil {
		return err
	}
	if err = compressedWriter.Close(); err != nil {
		return err
	}

	if err = ioutil.WriteFile("./provider/cmd/pulumi-resource-azurerm/schema.go", []byte(fmt.Sprintf(`package main
var pulumiSchema   = %#v
`,
		compressedSchema.Bytes(),
	)), 0600); err != nil {
		return err
	}
	return emitFile(outdir, "schema.json", formatted)
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
