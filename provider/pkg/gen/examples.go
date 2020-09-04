package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
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

type languageToExampleProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       description
	LanguageToExampleProgram languageToExampleProgram
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

			flattened, err := FlattenInput(exampleParams, resourceParams, metadata.Types)
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

			languageExample := languageToExampleProgram{}
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

// FlattenInput takes the parameters specified in Azure API specs and flattens them
// to match the desired format for the Pulumi schema. resourceParams is a mapping
// of API parameter names to provider.AzureApiParameter and types is a mapping for
// the API type names to provider.AzureApiType. The latter two can be derived from
// the metadata generated during schema generation.
func FlattenInput(
	input map[string]interface{},
	resourceParams map[string]provider.AzureApiParameter,
	types map[string]provider.AzureApiType,
) (map[string]interface{}, error) {
	containers := map[string]codegen.StringSet{}
	for k, v := range resourceParams {
		if v.Value != nil && v.Value.Container != "" {
			if _, ok := containers[v.Value.Container]; !ok {
				containers[v.Value.Container] = codegen.NewStringSet()
			}
			containers[v.Value.Container].Add(k)
		}
	}

	out := map[string]interface{}{}
	for k, v := range input {
		switch k {
		case "If-Match": // TODO: Not handled in schema
			continue
		case "api-version", "subscriptionId":
			continue // No need to emit these since we auto inject them
		}

		if _, ok := containers[k]; ok {
			contained, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("property %s is expected to be a map, recieved: %T", k, v)
			}
			flattened, err := FlattenInput(contained, resourceParams, types)
			if err != nil {
				return nil, err
			}

			mergeMap(out, flattened)
			continue
		}

		paramMetadata, ok := resourceParams[k]
		if !ok {
			debug.Log("missing item '%s' from resource metadata for resource, skipping", k)
			continue
		}
		// body parameters are folded in
		if paramMetadata.Body != nil {
			inBody, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expect body for param %s to be a map, recieved %T", k, v)
			}
			bodyVals, ok := inBody["properties"]
			if !ok {
				debug.Log("Missing properties in body param for '%s'. Assuming body already folded in.", k)
				bodyVals = inBody
			} else if len(inBody) > 1 {
				debug.Log("items in body outside of 'properties' field, will merge them into properties")
				for k, v := range inBody {
					if k != "properties" {
						bodyVals.(map[string]interface{})[k] = v
					}
				}
			}
			bodyValsMap := bodyVals.(map[string]interface{})
			bodyProps := map[string]provider.AzureApiProperty{}
			for k, v := range paramMetadata.Body.Properties {
				props := v
				props.Container = ""
				bodyProps[k] = props
			}
			flattened := transformProperties(bodyProps, types, bodyValsMap)
			mergeMap(out, flattened)
			continue
		}

		// replace param name with value in SdkName if provided.
		if paramMetadata.Value != nil && paramMetadata.Value.SdkName != "" {
			k = paramMetadata.Value.SdkName
		}

		out[k] = v
	}
	return out, nil
}

func transformProperty(prop *provider.AzureApiProperty, types map[string]provider.AzureApiType, val interface{}) interface{} {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Map:
		if prop.Ref == "" {
			typeName := strings.TrimPrefix(prop.Type, "#/types/")
			if _, typed := types[typeName]; typed {
				return transformProperties(types[prop.Type].Properties, types, val.(map[string]interface{}))
			}
			// Return untyped dictionaries as-is if no container unwrapping required.
			return val
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := types[typeName]
		return transformProperties(typ.Properties, types, val.(map[string]interface{}))
	case reflect.Slice, reflect.Array:
		var result []interface{}
		s := reflect.ValueOf(val)

		for i := 0; i < s.Len(); i++ {
			result = append(result, transformProperty(prop.Items, types, s.Index(i).Interface()))
		}
		return result
	}
	return val
}

func transformProperties(props map[string]provider.AzureApiProperty, types map[string]provider.AzureApiType, values map[string]interface{}) map[string]interface{} {
	containers := codegen.NewStringSet()
	for _, v := range props {
		if v.Container != "" {
			containers.Add(v.Container)
		}
	}

	result := map[string]interface{}{}
	for k, v := range values {
		prop, ok := props[k]
		if !ok {
			debug.Log("missing '%s' in props: %#v", k, props)
			continue
		}
		sdkName := k

		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		if containers.Has(k) {
			if v != nil {
				container := transformProperty(&prop, types, v)
				// Expect container types to wrap maps.
				mergeMap(result, container.(map[string]interface{}))
			}
			continue
		}
		if v != nil {
			result[sdkName] = transformProperty(&prop, types, v)
		}
	}
	return result
}

func mergeMap(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
