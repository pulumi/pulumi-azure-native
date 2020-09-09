package gen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi-azurerm/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/sourcegraph/jsonx"
)

func RenderFile(path string, metadata *provider.AzureApiMetadata) (*model.Body, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	var files []string
	if fileInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		root, err := parseJsonxTree(string(b))
		if err != nil {
			return nil, err
		}
		parseTrees[file] = root

	}

	return RenderTemplate(parseTrees, metadata)
}

// RenderTemplate renders a parsed CloudFormation template to a PCL program body. If there are errors in the template,
// the function returns an error.
func RenderTemplate(templates map[string]*jsonx.Node, metadata *provider.AzureApiMetadata) (*model.Body, error) {
	switch len(templates) {
	case 0:
		return &model.Body{}, nil
	}

	templ := TemplateElements{}

	for _, root := range templates {
		rootValue := jsonx.NodeValue(*root)
		for key, f := range rootValue.(map[string]interface{}) {
			switch key {
			case "$schema", "contentVersion", "apiProfile":
				// Ignore this
			case "parameters":
				fObj, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}

				for param, f := range fObj {
					fMap, ok := f.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("expect param %s to be defined with a map, got %T", param, f)
					}

					if err := templ.AddParameter(param, fMap); err != nil {
						return nil, err
					}
				}
			case "variables":
				fObj, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}

				for varName, f := range fObj {
					if err := templ.AddVariable(varName, f); err != nil {
						return nil, err
					}
				}
			case "resources":
				fObj, ok := f.([]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a list, got: %T", key, f)
				}

				for _, res := range fObj {
					resMap, ok := res.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("expect resource body to be a map, got: %T", res)
					}

					if err := templ.AddResource(resMap); err != nil {
						return nil, err
					}
				}
			case "outputs":
				fObj, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}

				for outputName, args := range fObj {
					if err := templ.AddOutput(outputName, args.(map[string]interface{})); err != nil {
						return nil, err
					}
				}
			}
		}
	}

	body, err := templ.RenderPCL(metadata)
	if err != nil {
		return nil, err
	}
	pcl.FormatBody(body)
	return body, nil
}

func extractNameFromMap(m interface{}) (string, error) {
	mMap, ok := m.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected map, got %T", m)
	}
	for k, v := range mMap {
		if k == "name" {
			return v.(string), nil
		}
	}
	return "", fmt.Errorf("no 'name' field found")
}

func parseJsonxTree(text string) (*jsonx.Node, error) {
	root, errs := jsonx.ParseTree(text, jsonx.ParseOptions{Comments: true, TrailingCommas: true})
	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to parse JSON: %v", errs)
	}
	return root, nil
}
