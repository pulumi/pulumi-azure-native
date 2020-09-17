package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/debug"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

func main() {
	var readFrom io.Reader
	var langs string

	var debugFlag = true
	debug.Debug = &debugFlag

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [arm-template path] lang\n", os.Args[0])
		fmt.Print("arm-template   Path to arm-template or assumed to be stdin if omitted\n")
		fmt.Print("lang           Comma separated list of languages - e.g. nodejs,python,dotnet,go\n")
		fmt.Println()
		return
	}
	if len(os.Args) == 2 {
		readFrom = os.Stdin
		langs = os.Args[1]
	} else {
		var err error
		filePath := os.Args[1]
		readFrom, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to read from file: %s: %+v", filePath, err)
		}
		langs = os.Args[2]
	}

	metadata, err := provider.LoadMetadata(azureApiResources)
	if err != nil {
		log.Fatalf("Failed to load metadata: %+v", err)
	}

	uncompressed, err := gzip.NewReader(bytes.NewReader(pulumiSchema))
	if err != nil {
		log.Fatalf("Failed to decompress schema: %+v", err)
	}

	var pkgSpec schema.PackageSpec
	if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
		log.Fatalf("Failure unmarshalling resource map: %+v", err)
	}
	if err = uncompressed.Close(); err != nil {
		log.Fatalf("Failure Closing uncompress stream for resource map: %+v", err)
	}

	body, diagnostics, err := gen.Render(readFrom, metadata, &pkgSpec)
	if err != nil {
		log.Fatalf("Failure rendering IR from template: %+v", err)
	}
	languages := strings.Split(langs, ",")
	programsMap, err := renderPrograms(body, &pkgSpec, languages)
	if err != nil {
		log.Fatalf("Failure rendering programs: %+v", err)
	}
	for k, v := range programsMap {
		fmt.Printf("Language: %s\n", k)
		fmt.Println()
		fmt.Printf("%s\n", v)
		fmt.Println()
		fmt.Println()
	}
	for k, diags := range diagnostics {
		fmt.Printf("Diagnostics for %s\n", k)
		for _, diag := range diags {
			fmt.Printf("WARN: [%s] at '%s' - '%s'\n", diag.Severity, diag.SourceToken, diag.Description)
		}
		fmt.Println()
	}
	return

}

func renderPrograms(body *model.Body, pkgSpec *schema.PackageSpec, languages []string) (map[string]string, error) {
	programBody := fmt.Sprintf("%v", body)
	debug.Log("%s\n", programBody)
	pkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return nil, fmt.Errorf("importing package spec: %w", err)
	}
	loaderOption := hcl2.Loader(gen.InMemoryPackageLoader(map[string]*schema.Package{
		"azure-nextgen": pkg,
	}))

	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
		return nil, fmt.Errorf("parse IR: %v", err)
	}
	if parser.Diagnostics.HasErrors() {
		err := parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	program, diags, err := hcl2.BindProgram(parser.Files, loaderOption)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if diags.HasErrors() {
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	languageToProgram := map[string]string{}

	for _, lang := range languages {
		var files map[string][]byte

		switch lang {
		case "dotnet":
			files, err = recoverableProgramGen(program, dotnet.GenerateProgram)
		case "go":
			files, err = recoverableProgramGen(program, gogen.GenerateProgram)
		case "nodejs":
			files, err = recoverableProgramGen(program, nodejs.GenerateProgram)
		case "python":
			files, err = recoverableProgramGen(program, python.GenerateProgram)
		default:
			continue
		}
		if err != nil {
			log.Printf("Program generation failed for language: %s, continuing", lang)
			continue
		}

		buf := strings.Builder{}
		for _, f := range files {
			_, err := buf.Write(f)
			if err != nil {
				return nil, err
			}
		}
		languageToProgram[lang] = buf.String()
	}

	return languageToProgram, nil
}

type programGenFn func(*hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)

func recoverableProgramGen(program *hcl2.Program, fn programGenFn) (files map[string][]byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during generation: %v", r)
		}
	}()

	var d hcl.Diagnostics
	files, d, err = fn(program)

	if err != nil {
		return nil, err
	}
	if d.HasErrors() {
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(d)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}
	return
}
