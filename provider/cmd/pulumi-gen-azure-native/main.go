// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/segmentio/encoding/json"

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

	// Use DEBUG_CODEGEN to just generate a single namespace (e.g. "Compute") for quick testing
	namespaces := os.Getenv("DEBUG_CODEGEN_NAMESPACES")
	if namespaces == "" {
		namespaces = "*"
	}

	azureProviders := openapi.ReadVersions(namespaces)

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
			if languages == "schema" {
				// We can't generate schema.json every time because it's slow and isn't reproducible.
				// So we warn in case someone's expecting to see changes to schema.json after running this.
				fmt.Println("Emitted `schema-full.json`. `schema.json` is generated as part of the docs.")
			}
			// Also, emit the resource metadata for the provider.
			if err = emitMetadata(meta, outdir, "main"); err != nil {
				break
			}

			// Now emit schema and metadata as byte encoded files for arm2pulumi
			arm2pulumiDir := path.Join(".", "provider", "cmd", "arm2pulumi")
			if err = emitSchema(*pkgSpec, version, arm2pulumiDir, "main", false); err != nil {
				break
			}
			// Also, emit the resource metadata for the provider.
			err = emitMetadata(meta, arm2pulumiDir, "main")

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
			err = gen.Examples(docsPkgSpec, docsMeta, resExamples, []string{"nodejs", "dotnet", "python", "go", "java", "yaml"})
			if err != nil {
				break
			}
			// Remove the version again.
			docsPkgSpec.Version = ""
			err = emitDocsSchema(docsPkgSpec, outdir)
		case "go":
			outdir := path.Join(".", "sdk", language)
			pkgSpec.Version = version
			// Work around the SDK size exceeding 512 MB by removing comments from versioned modules.
			// Roll this back when we have a better fix for the SDK size.
			specNoComments := removeAllComments(*pkgSpec)
			err = emitPackage(specNoComments, language, outdir)
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

func removeAllComments(pkgSpec schema.PackageSpec) *schema.PackageSpec {
	types := map[string]schema.ComplexTypeSpec{}
	for n, t := range pkgSpec.Types {
		props := map[string]schema.PropertySpec{}
		for pn, p := range t.Properties {
			p.Description = ""
			props[pn] = p
		}
		t.Description = ""
		t.Properties = props
		types[n] = t
	}
	pkgSpec.Types = types
	functions := map[string]schema.FunctionSpec{}
	for n, f := range pkgSpec.Functions {
		f.Description = ""
		if f.Inputs != nil {
			inputs := map[string]schema.PropertySpec{}
			for pn, p := range f.Inputs.Properties {
				p.Description = ""
				inputs[pn] = p
			}
			f.Inputs.Properties = inputs
		}
		if f.Outputs != nil {
			outputs := map[string]schema.PropertySpec{}
			for pn, p := range f.Outputs.Properties {
				p.Description = ""
				outputs[pn] = p
			}
			f.Outputs.Properties = outputs
		}
		functions[n] = f
	}
	pkgSpec.Functions = functions
	resources := map[string]schema.ResourceSpec{}
	for n, r := range pkgSpec.Resources {
		inputProperties := map[string]schema.PropertySpec{}
		for pn, p := range r.InputProperties {
			p.Description = ""
			inputProperties[pn] = p
		}
		properties := map[string]schema.PropertySpec{}
		for pn, p := range r.Properties {
			p.Description = ""
			properties[pn] = p
		}
		r.Description = ""
		r.InputProperties = inputProperties
		r.Properties = properties
		resources[n] = r
	}
	pkgSpec.Resources = resources
	return &pkgSpec
}

// emitSchema writes the Pulumi schema JSON to the 'schema.json' file in the given directory.
func emitSchema(pkgSpec schema.PackageSpec, version, outDir string, goPackageName string, emitJSON bool) error {
	pkgSpec.Version = version
	schemaJSON, err := json.Marshal(pkgSpec)
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	if err := emitFile(outDir, "schema-full.json", schemaJSON); err != nil {
		return err
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

func emitMetadata(metadata *resources.AzureAPIMetadata, outDir string, goPackageName string) error {
	meta := bytes.Buffer{}
	err := json.NewEncoder(&meta).Encode(metadata)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	err = emitFile(outDir, "metadata-compact.json", meta.Bytes())
	if err != nil {
		return err
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
