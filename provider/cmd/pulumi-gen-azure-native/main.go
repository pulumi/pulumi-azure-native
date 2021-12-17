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

	"github.com/pulumi/pulumi-azure-native/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	var debugEnabled bool
	debugEnv := os.Getenv("DEBUG_CODEGEN")
	if debugEnabled = debugEnv == "true"; debugEnabled {
		debug.Debug = &debugEnabled
	}

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
			outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azure-native")
			if err = emitSchema(*pkgSpec, version, outdir, "main", true); err != nil {
				break
			}
			// Also, emit the resource metadata for the provider.
			if err = emitMetadata(meta, outdir, "main", true); err != nil {
				break
			}

			// Now emit schema and metadata as byte encoded files for arm2pulumi
			arm2pulumiDir := path.Join(".", "provider", "cmd", "arm2pulumi")
			if err = emitSchema(*pkgSpec, version, arm2pulumiDir, "main", false); err != nil {
				break
			}
			// Also, emit the resource metadata for the provider.
			err = emitMetadata(meta, arm2pulumiDir, "main", false)
		case "docs":
			outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azure-native")
			docsProviders := openapi.SingleVersion(azureProviders)
			var docsPkgSpec *schema.PackageSpec
			var resExamples map[string][]resources.AzureAPIExample
			var docsMeta *resources.AzureAPIMetadata
			docsPkgSpec, docsMeta, resExamples, err = gen.PulumiSchema(docsProviders)
			if err != nil {
				break
			}
			// Ensure the spec is stamped with a version - Go gen needs it.
			docsPkgSpec.Version = version
			err = gen.Examples(docsPkgSpec, docsMeta, resExamples, []string{"nodejs", "dotnet", "python", "go"})
			if err != nil {
				break
			}
			// Remove the version again.
			docsPkgSpec.Version = ""
			err = emitDocsSchema(docsPkgSpec, outdir)
		case "go":
			outdir := path.Join(".", "sdk", language)
			pkgSpec.Version = version
			err = emitPackage(pkgSpec, language, outdir)
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
func emitSchema(pkgSpec schema.PackageSpec, version, outDir string, goPackageName string, emitJSON bool) error {
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

	err = emitFile(outDir, "schema.go", []byte(fmt.Sprintf(`package %s
var pulumiSchema = %#v
`, goPackageName, compressedSchema.Bytes())))
	if err != nil {
		return errors.Wrap(err, "saving metadata")
	}

	if emitJSON {
		if err := emitFile(outDir, "schema-full.json", schemaJSON); err != nil {
			return err
		}
	}
	return nil
}

// emitDocsSchema writes the Pulumi schema JSON to the 'schema-docs.json' file in the given directory.
func emitDocsSchema(pkgSpec *schema.PackageSpec, outDir string) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	return emitFile(outDir, "schema.json", schemaJSON)
}

func emitMetadata(metadata *resources.AzureAPIMetadata, outDir string, goPackageName string, emitJSON bool) error {
	compressedMeta := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedMeta)
	err := json.NewEncoder(compressedWriter).Encode(metadata)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	if err = compressedWriter.Close(); err != nil {
		return err
	}

	err = emitFile(outDir, "metadata.go", []byte(fmt.Sprintf(`package %s
var azureApiResources = %#v
`, goPackageName, compressedMeta.Bytes())))
	if err != nil {
		return err
	}

	if emitJSON {
		// To reduce the size, blank out
		formatted, err := json.MarshalIndent(metadata, "", "  ")
		if err != nil {
			return errors.Wrap(err, "marshaling metadata")
		}

		err = emitFile(outDir, "metadata.json", formatted)
		if err != nil {
			return err
		}
	}
	return nil
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
