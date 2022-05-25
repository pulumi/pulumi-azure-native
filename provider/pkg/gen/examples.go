// Copyright 2022, Pulumi Corporation.  All rights reserved.

package gen

import (
	"fmt"
	"github.com/segmentio/encoding/json"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"

	"github.com/pulumi/pulumi-java/pkg/codegen/java"
	yaml "github.com/pulumi/pulumi-yaml/pkg/pulumiyaml/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	hcl2 "github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/schollz/progressbar/v3"
)

type description string
type programText string
type language string

type languageToExampleProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       description
	LanguageToExampleProgram languageToExampleProgram
}
type resourceExamplesRenderData struct {
	Data []exampleRenderData
}
type resourceImportRenderData struct {
	Token         string
	SampleResID   string
	SampleResName string
}

// Examples renders Azure API examples to the pkgSpec for the specified list of languages.
func Examples(pkgSpec *schema.PackageSpec, metadata *resources.AzureAPIMetadata,
	resExamples map[string][]resources.AzureAPIExample, languages []string) error {
	sortedKeys := codegen.SortedKeys(pkgSpec.Resources) // To generate in deterministic order

	// Use a progress bar to show progress since this can be a long running process
	bar := progressbar.Default(int64(len(sortedKeys)), "Resources processed:")

	// cache to speed up code generation
	hcl2Cache := hcl2.Cache(hcl2.NewPackageCache())
	pkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return err
	}
	loaderOption := hcl2.Loader(InMemoryPackageLoader(map[string]*schema.Package{
		"azure-native": pkg,
	}))
	for _, pulumiToken := range sortedKeys {
		err := bar.Add(1)
		if err != nil {
			log.Printf("error in the progress bar %v", err)
		}

		if ShouldExclude(pulumiToken) {
			log.Printf("Skipping '%s' since it matches exclusion pattern", pulumiToken)
			continue
		}

		debug.Log("Processing '%s'", pulumiToken)
		resource := metadata.Resources[pulumiToken]
		seen := codegen.NewStringSet()
		split := strings.Split(pulumiToken, ":")
		if len(split) == 0 {
			return fmt.Errorf("invalid resourcename: %s", pulumiToken)
		}
		resourceName := pcl.Camel(split[len(split)-1])

		examplesRenderData := resourceExamplesRenderData{}
		importRenderData := resourceImportRenderData{
			Token: pulumiToken,
			// The name and ID will be overridden later if we find an example that contains a sample response.
			SampleResName: "myresource1",
			SampleResID:   resource.Path,
		}
		if resourceExamples, ok := resExamples[pulumiToken]; ok {
			for _, example := range resourceExamples {
				var items []model.BodyItem
				if seen.Has(example.Location) {
					continue
				}

				seen.Add(example.Location)
				f, err := os.Open(example.Location)
				if err != nil {
					return err
				}
				var exampleJSON map[string]interface{}
				if err = json.NewDecoder(f).Decode(&exampleJSON); err != nil {
					return err
				}
				if err = f.Close(); err != nil {
					return err
				}
				if _, ok := exampleJSON["parameters"]; !ok {
					return fmt.Errorf("example %s missing expected key: 'parameters'", example.Location)
				}

				resourceParams := map[string]resources.AzureAPIParameter{}
				for _, param := range resource.PutParameters {
					resourceParams[param.Name] = param
				}

				if _, ok := exampleJSON["parameters"].(map[string]interface{}); !ok {
					fmt.Printf("Expect parameters to be a map, received: %T for resource: %s, skipping.",
						exampleJSON["parameters"], pulumiToken)
					continue
				}
				exampleParams := exampleJSON["parameters"].(map[string]interface{})

				// Fill in sample name and ID for the import section.
				if exampleResponses, ok := exampleJSON["responses"].(map[string]interface{}); ok {
					for _, response := range exampleResponses {
						if responseMap, ok := response.(map[string]interface{}); ok {
							if body, ok := responseMap["body"].(map[string]interface{}); ok {
								if exampleID, ok := body["id"].(string); ok {
									importRenderData.SampleResID = exampleID
								}
								if exampleName, ok := body["name"].(string); ok {
									importRenderData.SampleResName = exampleName
								}
							}
						}
					}
				}

				flattened, err := FlattenParams(exampleParams, resourceParams, metadata.Types)
				if err != nil {
					fmt.Printf("tranforming input for example %s for resource %s: %v", example.Description, pulumiToken, err)
					continue
				}

				keys := codegen.SortedKeys(flattened)
				for _, k := range keys {
					v := flattened[k]
					val, err := pcl.RenderValue(v)
					if err != nil {
						return err
					}
					items = append(items, &model.Attribute{
						Name:  k,
						Value: val,
					})
				}

				block := model.Block{
					Type:   "resource",
					Body:   &model.Body{Items: items},
					Labels: []string{resourceName, pulumiToken},
				}
				body := &model.Body{Items: []model.BodyItem{&block}}
				pcl.FormatBody(body)
				languageExample, err := generateExamplePrograms(example, body, languages, hcl2Cache, loaderOption)
				if err != nil {
					return err
				}

				examplesRenderData.Data = append(examplesRenderData.Data,
					exampleRenderData{
						ExampleDescription:       description(example.Description),
						LanguageToExampleProgram: languageExample,
					})
			}
			if len(examplesRenderData.Data) > 0 {
				err := renderExampleToSchema(pkgSpec, pulumiToken, &examplesRenderData)
				if err != nil {
					return err
				}
			}
		}

		err = renderImportToSchema(pkgSpec, pulumiToken, &importRenderData)
		if err != nil {
			return err
		}

		metadata.Resources[pulumiToken] = resource
	}

	return nil
}

type programGenFn func(*hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)

func generateExamplePrograms(example resources.AzureAPIExample, body *model.Body, languages []string,
	bindOptions ...hcl2.BindOption) (languageToExampleProgram, error) {
	programBody := fmt.Sprintf("%v", body)
	debug.Log(programBody)
	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
		return nil, fmt.Errorf("failed to parse IR - file: %s: %v", example.Location, err)
	}
	if parser.Diagnostics.HasErrors() {
		debug.Log(programBody)
		err := parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	program, diags, err := hcl2.BindProgram(parser.Files, bindOptions...)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if diags.HasErrors() {
		log.Print(programBody)
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	languageExample := languageToExampleProgram{}

	for _, lang := range languages {
		var files map[string][]byte

		switch lang {
		case "dotnet":
			files, err = recoverableProgramGen(programBody, program, dotnet.GenerateProgram)
		case "go":
			files, err = recoverableProgramGen(programBody, program, gogen.GenerateProgram)
		case "nodejs":
			files, err = recoverableProgramGen(programBody, program, nodejs.GenerateProgram)
		case "python":
			files, err = recoverableProgramGen(programBody, program, python.GenerateProgram)
		case "yaml":
			files, err = recoverableProgramGen(programBody, program, yaml.GenerateProgram)
		case "java":
			files, err = recoverableProgramGen(programBody, program, java.GenerateProgram)
		default:
			continue
		}
		if err != nil {
			log.Printf("Program generation failed for language: %s for example %s, continuing", lang, example.Location)
			continue
		}

		buf := strings.Builder{}
		for _, f := range files {
			_, err := buf.Write(f)
			if err != nil {
				return nil, err
			}
		}
		languageExample[language(lang)] = programText(buf.String())
		debug.Log("Generated %s equivalent for %s", lang, example.Location)
		debug.Log("%s", buf.String())
	}

	return languageExample, nil
}

// Panic after 3 minutes, unless cancel is called.
func makeTimeout(timer string) (cancel func()) {
	timeout := make(chan struct{})
	go func() {
		for {
			select {
			case <-timeout:
				return
			case <-time.After(3 * time.Minute):
				log.Printf("%s timed out", timer)
				panic("timeout")
			}
		}
	}()
	return func() { timeout <- struct{}{} }
}

func recoverableProgramGen(name string, program *hcl2.Program, fn programGenFn) (files map[string][]byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during generation: %v", r)
		}
	}()

	var d hcl.Diagnostics
	cancel := makeTimeout(name)
	defer cancel()
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

func renderExampleToSchema(pkgSpec *schema.PackageSpec, resourceName string,
	examplesRenderData *resourceExamplesRenderData) error {
	const tmpl = `

{{"{{% examples %}}"}}
## Example Usage
{{- range .Data }}
{{ "{{% example %}}" }}
### {{ .ExampleDescription }}

{{- range $lang, $example := .LanguageToExampleProgram }}
{{ beginLanguage $lang }}
{{ $example }}
{{ endLanguage }}
{{ end }}
{{"{{% /example %}}"}}
{{- end }}
{{"{{% /examples %}}"}}
`
	res, ok := pkgSpec.Resources[resourceName]
	if !ok {
		return fmt.Errorf("missing resource from schema: %s", resourceName)
	}

	t, err := template.New("examples").Funcs(template.FuncMap{
		"beginLanguage": func(lang interface{}) string {
			l := fmt.Sprintf("%s", lang)
			switch l {
			case "nodejs":
				l = "typescript"
			case "dotnet":
				l = "csharp"
			}
			return fmt.Sprintf("```%s", l)
		},
		"endLanguage": func() string {
			return "```"
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}
	b := strings.Builder{}
	if err = t.Execute(&b, examplesRenderData); err != nil {
		return err
	}
	res.Description += b.String()
	pkgSpec.Resources[resourceName] = res
	return nil
}

func renderImportToSchema(pkgSpec *schema.PackageSpec, resourceName string,
	importRenderData *resourceImportRenderData) error {
	const tmpl = `
## Import

An existing resource can be imported using its type token, name, and identifier, e.g.

` + "```" + `sh
$ pulumi import {{ .Token }} {{ .SampleResName }} {{ .SampleResID }}
` + "```\n"
	res, ok := pkgSpec.Resources[resourceName]
	if !ok {
		return fmt.Errorf("missing resource from schema: %s", resourceName)
	}

	t, err := template.New("import").Parse(tmpl)
	if err != nil {
		return err
	}
	b := strings.Builder{}
	if err = t.Execute(&b, importRenderData); err != nil {
		return err
	}
	res.Description += b.String()
	pkgSpec.Resources[resourceName] = res
	return nil
}
