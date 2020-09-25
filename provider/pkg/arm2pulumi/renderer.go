package arm2pulumi

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/gen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/pcl"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/sourcegraph/jsonx"
)

// RenderIR generates an intermediate representation from template
// files passed as a map of file name to content.
// If metadata or pkgSpec are nil, they are loaded from the serialized
// versions generated at schema generation time.
// The return type is a *model.Body along with diagnostic information
// grouped by template element name.
func RenderIR(files map[string]string) (*model.Body, map[string][]Diagnostic, error) {
	jsonxMap := map[string]*jsonx.Node{}

	for fileName, content := range files {
		var err error
		jsonxMap[fileName], err = parseJsonxTree(content)
		if err != nil {
			return nil, nil, err
		}
	}
	return renderTemplate(jsonxMap)
}

// RenderIRFromReader generates an intermediate representation from template
// body read through reader.
//
// Metadata and pkgSpec are loaded from the serialized
// versions generated at schema generation time.
//
// The return type is a *model.Body along with diagnostic information
// grouped by template element name.
func RenderIRFromReader(reader io.Reader) (*model.Body, map[string][]Diagnostic, error) {

	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, nil, err
	}
	root, err := parseJsonxTree(buf.String())
	if err != nil {
		return nil, nil, err
	}
	return renderTemplate(map[string]*jsonx.Node{"azure-deploy.json": root})
}

// RenderFileIR generates an intermediate representation from template body
// from an ARM template file or directory at the specified path.
// Metadata and pkgSpec are loaded from the serialized
// versions generated at schema generation time.
// The return type is a *model.Body along with diagnostic information
// grouped by template element name. If there are fatal errors processing
// the template, an error is returned.
func RenderFileIR(path string) (*model.Body, map[string][]Diagnostic, error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}
	var files []string
	if fileInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, nil, err
		}
		for _, finfo := range fileInfos {
			if filepath.Ext(finfo.Name()) != ".json" {
				continue
			}
			files = append(files, filepath.Join(path, finfo.Name()))
		}
	} else {
		files = append(files, path)
	}

	parseTrees := map[string]*jsonx.Node{}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, nil, err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, nil, err
		}
		root, err := parseJsonxTree(string(b))
		if err != nil {
			return nil, nil, err
		}
		parseTrees[file] = root

	}

	return renderTemplate(parseTrees)
}

// renderTemplate renders a parsed ARM template to a PCL program body.
// If there are errors in the template, the function returns an error.
// Otherwise partial failures are returned through a map of diagnostics
// keyed by the element name.
func renderTemplate(templates map[string]*jsonx.Node) (*model.Body, map[string][]Diagnostic, error) {
	switch len(templates) {
	case 0:
		return &model.Body{}, nil, nil
	}

	if _, err := loadSchema(); err != nil {
		return nil, nil, err
	}
	if _, err := loadMetadata(); err != nil {
		return nil, nil, err
	}

	templ := NewTemplateElements()

	for _, root := range templates {
		rootValue := jsonx.NodeValue(*root).(map[string]interface{})
		if _, ok := rootValue["$schema"]; !ok {
			// simple validation for templates.. Skip if no "$schema" field exists
			continue
		}
		if f, ok := rootValue["parameters"]; ok {
			const key = "parameters"
			fObj, ok := f.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for param, f := range fObj {
				fMap, ok := f.(map[string]interface{})
				if !ok {
					return nil, nil, fmt.Errorf("expect param %s to be defined with a map, got %T", param, f)
				}

				if err := templ.addParameter(param, fMap, true); err != nil {
					return nil, nil, err
				}
			}
		}

		if f, ok := rootValue["variables"]; ok {
			const key = "variables"
			fObj, ok := f.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for varName, f := range fObj {
				if _, err := templ.addVariable(varName, f, true); err != nil {
					return nil, nil, err
				}
			}
		}

		if f, ok := rootValue["resources"]; ok {
			const key = "resources"
			fObj, ok := f.([]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("expect %s block to be a list, got: %T", key, f)
			}

			for _, res := range fObj {
				resMap, ok := res.(map[string]interface{})
				if !ok {
					return nil, nil, fmt.Errorf("expect resource body to be a map, got: %T", res)
				}

				if err := templ.addResource(resMap); err != nil {
					return nil, nil, err
				}
			}
		}

		if f, ok := rootValue["outputs"]; ok {
			const key = "outputs"
			fObj, ok := f.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for outputName, args := range fObj {
				if err := templ.addOutput(outputName, args.(map[string]interface{}), true); err != nil {
					return nil, nil, err
				}
			}
		}
	}

	if err := templ.EvaluateExpressions(true); err != nil {
		return nil, nil, err
	}

	if err := templ.EvaluateExpressions(false); err != nil {
		return nil, nil, err
	}

	if err := templ.Validate(); err != nil {
		return nil, nil, err
	}

	// We may not have a version when called from arm2pulumi website.
	// Set it here so go program gen doesn't barf
	if pkgSpec.Version == "" {
		pkgSpec.Version = "0.1.0"
	}

	body, err := templ.RenderPCL(metadata, pkgSpec)
	if err != nil {
		return nil, nil, err
	}
	pcl.FormatBody(body)
	return body, templ.GetDiagnostics(), nil
}

func parseJsonxTree(text string) (*jsonx.Node, error) {
	root, errs := jsonx.ParseTree(text, jsonx.ParseOptions{Comments: true, TrailingCommas: true})
	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to parse JSON: %v", errs)
	}
	return root, nil
}

// RenderPrograms takes a program model.Body and generates program text for each of the languages
// provided in the 'languages' argument. Result is a mapping from language name to program code.
// Partial failure rendering in a particular language is returned in error. If any language failed,
// the second return value is set to 'false' and an error is returned but partial successes are
// returned in the map.
// If all languages succeeded, true is returned for success and error is nil.
func RenderPrograms(body *model.Body, languages []string) (map[string]string, bool, error) {
	programBody := fmt.Sprintf("%v", body)
	debug.Log("%s\n", programBody)

	pkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return nil, false, fmt.Errorf("importing package spec: %w", err)
	}
	loaderOption := hcl2.Loader(gen.InMemoryPackageLoader(map[string]*schema.Package{
		"azure-nextgen": pkg,
	}))

	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
		return nil, false, fmt.Errorf("parse IR: %v", err)
	}
	if parser.Diagnostics.HasErrors() {
		err := parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	program, diags, err := hcl2.BindProgram(parser.Files, loaderOption)
	if err != nil {
		return nil, false, fmt.Errorf("bind program: %v", err)
	}
	if diags.HasErrors() {
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	languageToProgram := map[string]string{}

	var errs []error
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
			log.Printf("Program generation failed for language: %s, %+v", lang, err)
			err = fmt.Errorf("generating program for language %s: %w", lang, err)
			errs = append(errs, err)
			continue
		}

		buf := strings.Builder{}
		for _, f := range files {
			_, err := buf.Write(f)
			if err != nil {
				// unlikely to ever happen.
				return nil, false, err
			}
		}
		languageToProgram[lang] = buf.String()
	}

	if len(errs) > 0 {
		var errStrs []string
		for _, e := range errs {
			errStrs = append(errStrs, e.Error())
		}
		// kinda gross - returning partial success
		return languageToProgram, false, fmt.Errorf("program generation errors" + strings.Join(errStrs, "\n"))
	}

	return languageToProgram, true, nil
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
