// Copyright 2021, Pulumi Corporation.  All rights reserved.

package arm2pulumi

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-java/pkg/codegen/java"
	yaml "github.com/pulumi/pulumi-yaml/pkg/pulumiyaml/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	hcl2 "github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"
	"github.com/sourcegraph/jsonx"
)

var SupportedLanguages = map[string]interface{}{
	"nodejs": true,
	"python": true,
	"dotnet": true,
	"go":     true,
	"java":   true,
	"yaml":   true,
}

func IsSupportedLanguage(lang string) bool {
	_, ok := SupportedLanguages[lang]
	return ok
}

type Renderer struct {
	pkgSpec  *schema.PackageSpec
	metadata *resources.AzureAPIMetadata
}

type renderOptions struct {
	disableResourceLinking bool
}

// RenderOption applies a specific optional render option
type RenderOption interface {
	apply(*renderOptions)
}

type optFunc func(*renderOptions)

func (o optFunc) apply(options *renderOptions) {
	o(options)
}

// DisableResourceLinking instructs the renderer to avoid dereference linking between resources. By default,
// the renderer will perform link analysis. However, some ARM templates result in dependency cycles which
// can't be resolved. Enabling this RenderOption allows partially valid output to be generated in templates
// that contain reference cycles.
func DisableResourceLinking() RenderOption {
	return optFunc(func(o *renderOptions) {
		o.disableResourceLinking = true
	})
}

// NewRenderer creates a new Renderer which can be used to
// generate Pulumi programs from ARM templates.
func NewRenderer(pkgSpec *schema.PackageSpec, metadata *resources.AzureAPIMetadata) *Renderer {
	return &Renderer{
		pkgSpec:  pkgSpec,
		metadata: metadata,
	}
}

// RenderIR generates an intermediate representation from template
// files passed as a map of file name to content.
// If metadata or pkgSpec are nil, they are loaded from the serialized
// versions generated at schema generation time.
// The return type is a *model.Body along with diagnostic information
// grouped by template element name.
func (r *Renderer) RenderIR(files map[string]string) (*model.Body, map[string][]Diagnostic, error) {
	jsonxMap := map[string]*jsonx.Node{}

	for fileName, content := range files {
		var err error
		jsonxMap[fileName], err = parseJsonxTree(content)
		if err != nil {
			return nil, nil, err
		}
	}
	return renderTemplate(r.pkgSpec, r.metadata, jsonxMap)
}

// RenderIRFromReader generates an intermediate representation from template
// body read through reader.
//
// Metadata and pkgSpec are loaded from the serialized
// versions generated at schema generation time.
//
// The return type is a *model.Body along with diagnostic information
// grouped by template element name.
func (r *Renderer) RenderIRFromReader(reader io.Reader, options ...RenderOption) (*model.Body, map[string][]Diagnostic, error) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, nil, err
	}
	root, err := parseJsonxTree(buf.String())
	if err != nil {
		return nil, nil, err
	}
	return renderTemplate(r.pkgSpec, r.metadata, map[string]*jsonx.Node{"azure-deploy.json": root}, options...)
}

// RenderFileIR generates an intermediate representation from template body
// from an ARM template file or directory at the specified path.
// Metadata and pkgSpec are loaded from the serialized
// versions generated at schema generation time.
// The return type is a *model.Body along with diagnostic information
// grouped by template element name. If there are fatal errors processing
// the template, an error is returned.
func (r *Renderer) RenderFileIR(path string, options ...RenderOption) (*model.Body, map[string][]Diagnostic, error) {

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

	return renderTemplate(r.pkgSpec, r.metadata, parseTrees, options...)
}

// DecodeTemplate parses an ARM Templates to an array of resource property maps. Each map contains four keys:
// 'name' for resource name, 'token' for resource type token, 'args' for resource inputs, and 'dependsOn' for
// the names of resources that this resource depends on.
func DecodeTemplate(pkgSpec *schema.PackageSpec, metadata *resources.AzureAPIMetadata,
	templateJson string) ([]map[string]interface{}, error) {
	root, err := parseJsonxTree(templateJson)
	if err != nil {
		return nil, err
	}

	templ, err := renderTemplateElements(pkgSpec, metadata, map[string]*jsonx.Node{"default": root})
	if err != nil {
		return nil, err
	}

	var items []map[string]interface{}
	for _, e := range templ.elements {
		if r, ok := e.TemplateElement.(*resource); ok {
			var dependsOn []string
			for d := range e.dependencies {
				if dr, ok := d.TemplateElement.(*resource); ok {
					dependsOn = append(dependsOn, dr.resourceName)
				}
			}
			result := map[string]interface{}{
				"name":      r.resourceName,
				"token":     r.resourceToken,
				"args":      r.resourceParams,
				"dependsOn": dependsOn,
			}
			items = append(items, result)
		}
	}

	return items, nil
}

// renderTemplate renders a parsed ARM template to a PCL program body.
// If there are errors in the template, the function returns an error.
// Otherwise partial failures are returned through a map of diagnostics
// keyed by the element name.
func renderTemplate(
	pkgSpec *schema.PackageSpec,
	metadata *resources.AzureAPIMetadata,
	templates map[string]*jsonx.Node,
	options ...RenderOption,
) (*model.Body, map[string][]Diagnostic, error) {
	if len(templates) == 0 {
		return &model.Body{}, nil, nil
	}

	templ, err := renderTemplateElements(pkgSpec, metadata, templates, options...)
	if err != nil {
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

func renderTemplateElements(
	pkgSpec *schema.PackageSpec,
	metadata *resources.AzureAPIMetadata,
	templates map[string]*jsonx.Node,
	options ...RenderOption,
) (*TemplateElements, error) {
	templ := NewTemplateElements(options...)
	if len(templates) == 0 {
		return templ, nil
	}

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
				return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for param, f := range fObj {
				fMap, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect param %s to be defined with a map, got %T", param, f)
				}

				if err := templ.addParameter(param, fMap, true); err != nil {
					return nil, err
				}
			}
		}

		if f, ok := rootValue["variables"]; ok {
			const key = "variables"
			fObj, ok := f.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for varName, f := range fObj {
				if _, err := templ.addVariable(varName, f, true); err != nil {
					return nil, err
				}
			}
		}

		if f, ok := rootValue["resources"]; ok {
			const key = "resources"
			fObj, ok := f.([]interface{})
			if !ok {
				return nil, fmt.Errorf("expect %s block to be a list, got: %T", key, f)
			}

			for _, res := range fObj {
				resMap, ok := res.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect resource body to be a map, got: %T", res)
				}

				if err := templ.addResource(resMap); err != nil {
					return nil, err
				}
			}
		}

		if f, ok := rootValue["outputs"]; ok {
			const key = "outputs"
			fObj, ok := f.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
			}

			for outputName, args := range fObj {
				if err := templ.addOutput(outputName, args.(map[string]interface{}), true); err != nil {
					return nil, err
				}
			}
		}
	}

	if err := templ.EvaluateExpressions(); err != nil {
		return nil, err
	}

	if err := templ.Validate(pkgSpec, metadata); err != nil {
		return nil, err
	}

	return templ, nil
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
func (r *Renderer) RenderPrograms(body *model.Body, languages []string) (map[string]string, bool, error) {
	programBody := fmt.Sprintf("%v", body)
	debug.Log("%s\n", programBody)

	pkg, err := schema.ImportSpec(*r.pkgSpec, nil)
	if err != nil {
		return nil, false, fmt.Errorf("importing package spec: %w", err)
	}
	loaderOption := hcl2.Loader(gen.InMemoryPackageLoader(map[string]*schema.Package{
		"azure-native": pkg,
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

	program, diags, err := hcl2.BindProgram(parser.Files, loaderOption, hcl2.AllowMissingVariables, hcl2.AllowMissingProperties)
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
		var gen programGenFn

		switch lang {
		case "dotnet":
			gen = dotnet.GenerateProgram
		case "go":
			gen = gogen.GenerateProgram
		case "nodejs":
			gen = nodejs.GenerateProgram
		case "python":
			gen = python.GenerateProgram
		case "java":
			gen = java.GenerateProgram
		case "yaml":
			gen = yaml.GenerateProgram
		default:
			continue
		}
		files, err := recoverableProgramGen(program, gen)
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
