// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"bytes"
	"fmt"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

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

	var azureProviders *openapi.AzureProviders
	var pkgSpec *schema.PackageSpec
	var meta *resources.AzureAPIMetadata
	var err error

	languageSet := codegen.NewStringSet(strings.Split(languages, ",")...)
	basicLanguages := []string{"dotnet", "nodejs", "python"}
	supportedOutputs := codegen.NewStringSet("schema", "docs", "go")
	for _, language := range basicLanguages {
		supportedOutputs.Add(language)
	}

	if unsupported := supportedOutputs.Subtract(languageSet); len(unsupported) > 0 {
		err = fmt.Errorf("unsupported outputs: %v", unsupported.SortedValues())
	}

	if languageSet.Has("schema") {
		versions := openapi.ReadVersions(namespaces)
		azureProviders = &versions
		pkgSpec, meta, _, err = gen.PulumiSchema(*azureProviders)
		if err != nil {
			panic(err)
		}
		outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azure-native")
		if err = emitSchema(*pkgSpec, version, outdir, "main", true); err != nil {
			panic(err)
		}
		if languages == "schema" {
			// We can't generate schema.json every time because it's slow and isn't reproducible.
			// So we warn in case someone's expecting to see changes to schema.json after running this.
			fmt.Println("Emitted `schema-full.json`. `schema.json` is generated as part of the docs.")
		}
		// Also, emit the resource metadata for the provider.
		if err = emitMetadata(meta, outdir, "main"); err != nil {
			panic(err)
		}

		// Now emit schema and metadata as byte encoded files for arm2pulumi
		arm2pulumiDir := path.Join(".", "provider", "cmd", "arm2pulumi")
		if err = emitSchema(*pkgSpec, version, arm2pulumiDir, "main", false); err != nil {
			panic(err)
		}
		// Also, emit the resource metadata for the provider.
		err = emitMetadata(meta, arm2pulumiDir, "main")
	} else if languageSet.Subtract(codegen.NewStringSet("docs")).Any() {
		// Just read existing schema if we're not re-generating
		schemaPath := path.Join("provider", "cmd", "pulumi-resource-azure-native", "schema-full.json")
		schemaBytes, err := os.ReadFile(schemaPath)
		if err != nil {
			panic(err)
		}
		var schema schema.PackageSpec
		err = json.Unmarshal([]byte(schemaBytes), &schema)
		if err != nil {
			panic(err)
		}
		pkgSpec = &schema
	}

	if languageSet.Has("docs") {
		if azureProviders == nil {
			versions := openapi.ReadVersions(namespaces)
			azureProviders = &versions
		}
		outdir := path.Join(".", "provider", "cmd", "pulumi-resource-azure-native")
		docsProviders := openapi.SingleVersion(*azureProviders)
		var docsPkgSpec *schema.PackageSpec
		var resExamples map[string][]resources.AzureAPIExample
		var docsMeta *resources.AzureAPIMetadata
		docsPkgSpec, docsMeta, resExamples, err = gen.PulumiSchema(docsProviders)
		if err != nil {
			panic(err)
		}
		// Ensure the spec is stamped with a version - Go gen needs it.
		docsPkgSpec.Version = version
		err = gen.Examples(docsPkgSpec, docsMeta, resExamples, []string{"nodejs", "dotnet", "python", "go", "java", "yaml"})
		if err != nil {
			panic(err)
		}
		// Remove the version again.
		docsPkgSpec.Version = ""
		err = emitDocsSchema(docsPkgSpec, outdir)
		if err != nil {
			panic(err)
		}
	}

	for _, language := range basicLanguages {
		if languageSet.Has(language) {
			outdir := path.Join(".", "sdk", language)
			pkgSpec.Version = version
			err = emitPackage(pkgSpec, language, outdir)
			if err != nil {
				panic(err)
			}
		}
	}

	if languageSet.Has("go") {
		outdir := path.Join(".", "sdk")
		pkgSpec.Version = version
		err = emitSplitPackage(pkgSpec, "go", outdir)
		if err != nil {
			panic(err)
		}
	}
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

func emitSplitPackage(pkgSpec *schema.PackageSpec, language, outDir string) error {
	pkgCopy := gen.SetGoBasePath(*pkgSpec, "github.com/pulumi/pulumi-azure-native-sdk")

	ppkg, err := schema.ImportSpec(*pkgCopy, nil)
	if err != nil {
		return errors.Wrap(err, "reading schema")
	}

	files, err := generate(ppkg, language)
	if err != nil {
		return errors.Wrapf(err, "generating %s package", language)
	}

	version := gen.GoModVersion(ppkg.Version)
	files["pulumi-azure-native-sdk/version.txt"] = []byte(version)
	files["pulumi-azure-native-sdk/go.mod"] = []byte(goModTemplate(GoMod{}))

	for f, contents := range files {
		if err := emitFile(outDir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
		}

		matched, err := filepath.Match("pulumi-azure-native-sdk/*/init.go", f)
		if err != nil {
			return err
		}
		if matched {
			dir := filepath.Dir(f)
			module := filepath.Base(dir)
			modPath := filepath.Join(dir, "go.mod")
			modContent := goModTemplate(GoMod{
				Version:       version,
				SubmoduleName: module,
			})

			if err := emitFile(outDir, modPath, []byte(modContent)); err != nil {
				return errors.Wrapf(err, "emitting file %v", modPath)
			}
		}
	}

	return nil
}

type GoMod struct {
	Version       string
	SubmoduleName string
}

var goModTemplateCache *template.Template

func goModTemplate(goMod GoMod) string {
	var err error
	if goModTemplateCache == nil {
		goModTemplateCache, err = template.New("go-mod").Parse(`
{{ if eq .SubmoduleName "" }}
module github.com/pulumi/pulumi-azure-native-sdk
{{ else }}
module github.com/pulumi/pulumi-azure-native-sdk/{{ .SubmoduleName }}
{{ end }}

go 1.17

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/pkg/errors v0.9.1
{{ if ne .SubmoduleName "" }}
	github.com/pulumi/pulumi-azure-native-sdk {{ .Version }}
{{ end }}
	github.com/pulumi/pulumi/sdk/v3 v3.37.2
)

{{ if ne .SubmoduleName "" }}
replace github.com/pulumi/pulumi-azure-native-sdk {{ .Version }} => ../
{{ end }}
`)
		if err != nil {
			panic(err)
		}
	}

	var result bytes.Buffer
	err = goModTemplateCache.Execute(&result, goMod)
	if err != nil {
		panic(err)
	}
	return result.String()
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
