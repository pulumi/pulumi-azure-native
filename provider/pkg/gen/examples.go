package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/pulumi/pulumi-azurerm/provider/pkg/debug"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/schollz/progressbar/v3"
)

type description string
type programText string
type language string

type languageToProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       description
	LanguageToExampleProgram languageToProgram
}
type resourceExamplesRenderData struct {
	Data []exampleRenderData
}

// Examples renders Azure API examples to the pkgSpec for the specified list of languages.
func Examples(pkgSpec *schema.PackageSpec, metadata *provider.AzureApiMetadata, languages []string) error {
	sortedKeys := codegen.SortedKeys(metadata.Resources) // To generate in deterministic order

	pulumiOptions := []hcl2.BindOption{hcl2.Cache(hcl2.NewPackageCache())}
	bar := progressbar.Default(int64(len(metadata.Resources)), "Resources processed:")
	for _, pulumiToken := range sortedKeys {
		bar.Add(1)
		if excludeResources.Has(pulumiToken) {
			debug.Log("Skipping '%s'", pulumiToken)
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
		for _, example := range resource.Examples {
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

			resourceParams := map[string]provider.AzureApiParameter{}
			for _, param := range resource.PutParameters {
				resourceParams[param.Name] = param
			}

			if _, ok := exampleJSON["parameters"].(map[string]interface{}); !ok {
				return fmt.Errorf("expect parameters to be a map, received: %T", exampleJSON["parameters"])
			}
			exampleParams := exampleJSON["parameters"].(map[string]interface{})

			flattened, err := flattenInput(exampleParams, resourceParams, metadata.Types)
			if err != nil {
				return err
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
			programBody := fmt.Sprintf("%v", body)
			debug.Log(example.Location)
			debug.Log(programBody)

			parser := syntax.NewParser()
			if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
				log.Fatalf("failed to parse IR - file: %s: %v", example, err)
			}
			if parser.Diagnostics.HasErrors() {
				fmt.Print(programBody)
				parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
				continue
			}

			program, diags, err := hcl2.BindProgram(parser.Files, pulumiOptions...)
			if err != nil {
				log.Fatalf("failed to bind program: %v", err)
			}
			if diags.HasErrors() {
				log.Printf(programBody)
				program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
				continue
			}

			languageExample := languageToProgram{}
			for _, lang := range languages {
				var files map[string][]byte

				switch lang {
				case "dotnet":
					files, diags, err = dotnet.GenerateProgram(program)
				case "go":
					files, diags, err = gogen.GenerateProgram(program)
				case "nodejs":
					files, diags, err = nodejs.GenerateProgram(program)
				case "python":
					files, diags, err = python.GenerateProgram(program)
				case "schema":
					continue
				}
				if err != nil {
					log.Fatalf("failed to generate program: %v", err)
				}
				if diags.HasErrors() {
					log.Printf(programBody)
					program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
					os.Exit(-1)
				}

				buf := strings.Builder{}
				for _, f := range files {
					_, err := buf.Write(f)
					if err != nil {
						return err
					}
				}
				languageExample[language(lang)] = programText(buf.String())
				debug.Log("Generated %s equivalent for %s", lang, example.Location)
				debug.Log("%s", buf.String())
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

		// Don't need the example information in metadata anymore
		resource.Examples = nil
		metadata.Resources[pulumiToken] = resource
	}

	return nil
}

func renderExampleToSchema(pkgSpec *schema.PackageSpec, resourceName string, examplesRenderData *resourceExamplesRenderData) error {
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
