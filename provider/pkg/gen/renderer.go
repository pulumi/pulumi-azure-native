package gen

import (
	"fmt"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/sourcegraph/jsonx"
)

func Render(reader io.Reader, metadata *provider.AzureAPIMetadata, pkgSpec *schema.PackageSpec) (*model.Body, map[string][]Diagnostic, error) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, nil, err
	}
	root, err := parseJsonxTree(buf.String())
	if err != nil {
		return nil, nil, err
	}
	return RenderTemplate(map[string]*jsonx.Node{"azure-deploy.json": root}, metadata, pkgSpec)
}

func RenderFile(path string, metadata *provider.AzureAPIMetadata, pkgSpec *schema.PackageSpec) (*model.Body, map[string][]Diagnostic, error) {
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

	return RenderTemplate(parseTrees, metadata, pkgSpec)
}

// RenderTemplate renders a parsed CloudFormation template to a PCL program body. If there are errors in the template,
// the function returns an error.
func RenderTemplate(templates map[string]*jsonx.Node,
	metadata *provider.AzureAPIMetadata,
	pkgSpec *schema.PackageSpec) (*model.Body, map[string][]Diagnostic, error) {
	switch len(templates) {
	case 0:
		return &model.Body{}, nil, nil
	}

	templ := NewTemplateElements()

	for _, root := range templates {
		rootValue := jsonx.NodeValue(*root).(map[string]interface{})
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

				if err := templ.AddParameter(param, fMap, true); err != nil {
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
				if _, err := templ.AddVariable(varName, f, true); err != nil {
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

				if err := templ.AddResource(resMap); err != nil {
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
				if err := templ.AddOutput(outputName, args.(map[string]interface{}), true); err != nil {
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
