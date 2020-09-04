package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/goccy/go-yaml/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/pcl"
	"github.com/pulumi/pulumi/pkg/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
	"github.com/sourcegraph/jsonx"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

func RenderFile(path string) (*model.Body, error) {
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

	return RenderTemplate(parseTrees)
}

// RenderTemplate renders a parsed CloudFormation template to a PCL program body. If there are errors in the template,
// the function returns an error.
func RenderTemplate(templates map[string]*jsonx.Node) (*model.Body, error) {
	switch len(templates) {
	case 0:
		return &model.Body{}, nil
	}

	ctx := &renderContext{
		parameters: map[string]*model.Variable{},
		variables:  map[string]*model.Variable{},
		functions:  map[string]*model.Function{},
		resources:  map[string]*model.Variable{},
		outputs:    map[string]*model.Variable{},
	}

	for _, root := range templates {
		if root.Type != jsonx.Object {
			continue
		}

		// Declare parameters, mappings, conditions, resources, and outputs.
		rootValue := jsonx.NodeValue(*root)
		for key, f := range rootValue.(map[string]interface{}) {
			var table map[string]*model.Variable
			containerType := reflect.Map
			switch key {
			case "parameters":
				table = ctx.parameters
			case "variables":
				table = ctx.variables
			case "resources":
				containerType = reflect.Slice
				table = ctx.resources
			case "outputs":
				table = ctx.outputs
			case "apiProfile":
				ctx.apiProfile = f.(string)
			default:
				continue
			}

			items := reflect.ValueOf(f)

			if items.Kind() != containerType {
				return nil, fmt.Errorf("unexpected type of block %s", key)
			}
			switch items.Kind() {
			case reflect.Slice:
				for i := 0; i < items.Len(); i++ {
					item := items.Index(i)
					if item.Kind() != reflect.Map {
						return nil, fmt.Errorf("expect block %s to consist of maps, received %T", item.Interface())
					}
					name, err := extractNameFromMap(item.Interface())
					if err != nil {
						return nil, err
					}
					table[name] = nil
				}
			case reflect.Map:
				for itemIter := items.MapRange(); itemIter.Next(); {
					table[itemIter.Key().String()] = nil
				}
			}
		}
	}
	ctx.assignNames()

	var items []model.BodyItem

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
					parameter, err := ctx.renderParameter(param, fMap)
					if err != nil {
						return nil, err
					}
					items = append(items, parameter...)
				}
			case "variables":
				fObj, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}
				variables, err := ctx.renderObjects(fObj, ctx.renderVariable)
				if err != nil {
					return nil, err
				}
				items = append(items, variables...)
			case "Resources":
				fObj, ok := f.([]map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}
				name, err := extractNameFromMap(fObj)
				if err != nil {
					return nil, fmt.Errorf("failed processing %s block", key)
				}
				res := map[string]interface{}{
					name: fObj,
				}
				resources, err := ctx.renderObjects(res, ctx.renderResource)
				if err != nil {
					return nil, err
				}
				items = append(items, resources...)
			case "Outputs":
				fObj, ok := f.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expect %s block to be a map, got: %T", key, f)
				}
				mappings, err := ctx.renderObjects(fObj, ctx.renderOutput)
				if err != nil {
					return nil, err
				}
				items = append(items, mappings...)
			default:
				return nil, fmt.Errorf("unexpected template property %v", key)
			}
		}
	}

	body := &model.Body{Items: items}
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

// A renderContext contains the state necessary to render an Arm template to a Pulumi program.
type renderContext struct {
	// The maps below map from names in the ARM template to output variables. These maps are used during
	// rendering to emit appropriate references to the entities.

	parameters map[string]*model.Variable // Template parameters
	variables  map[string]*model.Variable // Template variables
	functions  map[string]*model.Function // Template user-defined functions
	resources  map[string]*model.Variable // Template resources
	outputs    map[string]*model.Variable // Template outputs

	schema         string
	contentVersion string
	apiProfile     string
}

// null represents PCL's builtin `null` variable
var null = &model.Variable{
	Name:         "null",
	VariableType: model.NoneType,
}

// checkArgsArray validates the argument to a CloudFormation function. If the argument is a sequence, its length is
// checked against the provided minimum and maximum arg count. If the sequence is valid, its values are returned.
//
// If the argument is a single value and the minimum count is greater than one, the argument is wrapped in an array
// and returned.
func (ctx *renderContext) checkArgsArray(name string, v ast.Node, min, max int) ([]ast.Node, error) {
	if v.Type() != ast.SequenceType {
		if min > 1 {
			return nil, fmt.Errorf("the argument to '%s' must be an array", name)
		}
		return []ast.Node{v}, nil
	}

	arr := v.(*ast.SequenceNode).Values
	if len(arr) < min || (max >= 0 && len(arr) > max) {
		if min == max {
			return nil, fmt.Errorf("the argument to '%s' must have exactly %v elements", name, min)
		}
		if max >= 0 {
			return nil, fmt.Errorf("the argument to '%s' must have between %v and %v elements", name, min, max)
		}
		return nil, fmt.Errorf("the argument to '%s' must have at least %v elements", name, min)
	}
	return arr, nil
}

var holePattern = regexp.MustCompile(`\${([^}]*)}`)

// renderSub renders a call to Fn::Sub. The call is converted to a template expression. If an envrionment map is
// provided, references to map elements are replaced with the corresponding elements.
func (ctx *renderContext) renderSub(name string, value ast.Node) (model.Expression, error) {
	arr, err := ctx.checkArgsArray(name, value, 1, 2)
	if err != nil {
		return nil, err
	}

	if arr[0].Type() != ast.StringType {
		return nil, fmt.Errorf("the first argument to 'Fn::Sub' must be a string")
	}

	var environment map[string]model.Expression
	if len(arr) == 2 {
		values, ok := mapValues(arr[1])
		if !ok {
			return nil, fmt.Errorf("the second argument to 'Fn::Sub' must be a mapping")
		}

		for _, f := range values {
			v, err := pcl.RenderValue(f.Value)
			if err != nil {
				return nil, err
			}
			environment[keyString(f)] = v
		}
	}

	text := arr[0].(*ast.StringNode).Value
	holes := holePattern.FindAllStringSubmatchIndex(text, -1)

	var literals []string
	var refs []model.Expression
	start, end := 0, 0
	for _, hole := range holes {
		end = hole[0]

		literals = append(literals, text[start:end])

		ref := text[hole[2]:hole[3]]
		switch {
		case len(ref) == 0:
			return nil, fmt.Errorf("empty variable in 'Fn::Sub' input string")
		case ref[0] == '!':
			literals[len(literals)-1] = literals[len(literals)-1] + ref
		default:
			x, err := ctx.renderRef(ref)
			if err != nil {
				v, ok := environment[ref]
				if !ok {
					return nil, err
				}
				x = v
			}
			refs = append(refs, x)
		}

		start = hole[1]
	}
	literals = append(literals, text[start:])

	contract.Assert(len(literals) == len(refs)+1)

	var parts []model.Expression
	for i, l := range literals {
		parts = append(parts, pcl.PlainLit(l))
		if i < len(refs) {
			parts = append(parts, refs[i])
		}
	}
	return &model.TemplateExpression{Parts: parts}, nil
}

// renderArgsArray validates and renders the argument to a function as an array of expressions.
func (ctx *renderContext) renderArgsArray(name string, arg ast.Node, min, max int) ([]model.Expression, error) {
	arr, err := ctx.checkArgsArray(name, arg, min, max)
	if err != nil {
		return nil, err
	}

	var args []model.Expression
	for _, v := range arr {
		arg, err := ctx.renderValue(v)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	return args, nil
}

type armDataType string

const (
	intType          armDataType = "int"
	boolType         armDataType = "bool"
	stringType       armDataType = "string"
	secureStringType armDataType = "securestring"
	objectType       armDataType = "object"
	secureObjectType armDataType = "secureObject"
	arrayType        armDataType = "array"
)

func extractType(info map[string]interface{}) (armDataType, error) {
	typ, ok := info["type"]
	if !ok {
		return "", errors.New("missing required field 'type'")
	}

	typeStr, ok := typ.(string)
	if !ok {
		return "", errors.New("expect type to be a string")
	}

	armDType := armDataType(typeStr)
	switch armDType {
	case intType, boolType, stringType, secureStringType, objectType, secureObjectType, arrayType:
	default:
		return "", fmt.Errorf("unexpected type: %s", typeStr)
	}

	return armDType, nil
}

// renderParameter renders a template parameter.
// The parameter is rendered as either a sinmple config variable
// definition or a config variable definition and an invocation of one of the two SSM parameter getters (in the case of
// SSM parameters). Uses of SSM parameters reference the result of the invocation.
func (ctx *renderContext) renderParameter(name string, paramInfo map[string]interface{}) ([]model.BodyItem, error) {
	typ, err := extractType(paramInfo)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter %s: %w", name, err)
	}

	typeExpr := string(typ)

	paramRefVar, ok := ctx.parameters[name]
	contract.Assert(ok)

	paramDefVar := paramRefVar

	var defaultValue model.Expression
	if def, ok := paramInfo["defaultValue"]; ok {
		v, err := pcl.RenderValue(def)
		if err != nil {
			return nil, err
		}
		defaultValue = v
	}

	configDef := &model.Block{
		Type:   "config",
		Labels: []string{paramDefVar.Name, typeExpr},
		Body:   &model.Body{},
	}

	if defaultValue != nil {
		configDef.Body.Items = append(configDef.Body.Items, &model.Attribute{
			Name:  "default",
			Value: defaultValue,
		})
	}
	return []model.BodyItem{configDef}, nil

}

// renderVariable renders an ARM template variable as a PCL local variable.
func (ctx *renderContext) renderVariable(name string, value interface{}) (model.BodyItem, error) {
	v, ok := ctx.variables[name]
	contract.Assert(ok)

	m, err := pcl.RenderValue(value)
	if err != nil {
		return nil, err
	}

	return &model.Attribute{
		Name:  v.Name,
		Value: m,
	}, nil
}

func (ctx *renderContext) renderResources(resources []map[string]interface{}) ([]model.BodyItem, error) {
	// Need to load all resources, build a DAG and then topological sort to render
	
}

// renderResource renders a CloudFormation resource as a PCL resource.
func (ctx *renderContext) renderResource(name string, value interface{}) (model.BodyItem, error) {
	values, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("resource '%v' must be a mapping", name)
	}

	resourceVar, ok := ctx.resources[name]
	contract.Assert(ok)

	var token string
	var items []model.BodyItem
	for k, v := range values {
		switch k {
		case "CreationPolicy", "DeletionPolicy", "Metadata", "Properties", "UpdatePolicy", "UpdateReplacePolicy":
			v, err := ctx.renderValue(f.Value)
			if err != nil {
				return nil, err
			}
			items = append(items, &model.Attribute{
				Name:  camel(keyString(f)),
				Value: v,
			})
		case "DependsOn":
			var arr []ast.Node
			switch f.Value.Type() {
			case ast.StringType:
				arr = []ast.Node{f.Value}
			case ast.SequenceType:
				arr = f.Value.(*ast.SequenceNode).Values
			default:
				return nil, fmt.Errorf("the \"DependsOn\" attribute for resource '%v' must be a string or list of strings", name)
			}

			var refs []model.Expression
			for _, v := range arr {
				if v.Type() != ast.StringType {
					return nil, fmt.Errorf("the \"DependsOn\" attribute for resource '%v' must be a string or list of strings", name)
				}
				resourceName := v.(*ast.StringNode).Value
				resourceVar, ok := ctx.resources[resourceName]
				if !ok {
					return nil, fmt.Errorf("unknown resource '%v'", resourceName)
				}
				refs = append(refs, model.VariableReference(resourceVar))
			}
			items = append(items, &model.Block{
				Type: "options",
				Body: &model.Body{
					Items: []model.BodyItem{
						&model.Attribute{
							Name: "dependsOn",
							Value: &model.TupleConsExpression{
								Expressions: refs,
							},
						},
					},
				},
			})
		case "Type":
			if f.Value.Type() != ast.StringType {
				return nil, fmt.Errorf("the \"Type\" of reosurce '%v' must be a string", name)
			}
			token = resourceToken(f.Value.(*ast.StringNode).Value)
		default:
			return nil, fmt.Errorf("unsupported property '%v' in resource '%v'", f.Key, name)
		}
	}

	if token == "" {
		return nil, fmt.Errorf("resource '%v' has no \"Type\" attribute", name)
	}
	return &model.Block{
		Type:   "resource",
		Labels: []string{resourceVar.Name, token},
		Body:   &model.Body{Items: items},
	}, nil
}

// renderOutput renders an ARM output as a PCL output.
func (ctx *renderContext) renderOutput(name string, values interface{}) (model.BodyItem, error) {
	outputVar, ok := ctx.outputs[name]
	contract.Assert(ok)

	value, ok := values.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expect output %s to be map, got %T", name, values)
	}

	if cond, ok := value["condition"]; ok {
		if keep, ok := cond.(bool); !ok {
			return nil, fmt.Errorf("expect condition field in output %s to be boolean, received: %T", name, cond)
		} else if !keep {
			return nil, nil
		}
	}

	// TODO: Need to add support for copy
	val, ok := value["value"]
	if !ok {
		return nil, fmt.Errorf("output '%v' has no \"value\" attribute", name)
	}

	x, err := pcl.RenderValue(val)
	if err != nil {
		return nil, err
	}
	return &model.Block{
		Type:   "output",
		Labels: []string{outputVar.Name},
		Body: &model.Body{
			Items: []model.BodyItem{
				&model.Attribute{
					Name:  "value",
					Value: x,
				},
			},
		},
	}, nil
}

// An objectRenderer renders a map key/value into a body item.
type objectRenderer func(string, interface{}) (model.BodyItem, error)

// renderObjects is a helper that renders a set of named objects in a mapping. Each named object is passed to the
// provided renderer.
func (ctx *renderContext) renderObjects(obj map[string]interface{}, render objectRenderer) ([]model.BodyItem, error) {
	var body []model.BodyItem
	for k, v := range obj {
		i, err := render(k, v)
		if err != nil {
			return nil, err
		}
		if i == nil {
			continue
		}
		body = append(body, i)
	}

	return body, nil
}

// assignNames assigns names to the variables used to represent template parameters outputs, SSM values, resources,
// mappings, conditions, and pseudo-parameters. Care is taken to keep parameter and output names as close to their
// original names as possible.
func (ctx *renderContext) assignNames() {
	assigned := codegen.StringSet{}

	assign := func(name, suffix string) *model.Variable {
		assignName := func(name, suffix string) string {
			name = pcl.Camel(pcl.MakeLegalIdentifier(name))
			if !assigned.Has(name) {
				assigned.Add(name)
				return name
			}

			base := name + suffix
			name = base
			for i := 0; assigned.Has(name); i++ {
				name = fmt.Sprintf("%s%d", base, i)
			}
			assigned.Add(name)
			return name
		}

		return &model.Variable{
			Name:         assignName(name, suffix),
			VariableType: model.DynamicType,
		}
	}

	// TODO: do this in source order
	assignNames := func(m map[string]*model.Variable, suffix string) {
		names := make([]string, 0, len(m))
		for n := range m {
			names = append(names, n)
		}
		sort.Strings(names)

		for _, n := range names {
			m[n] = assign(n, suffix)
		}
	}

	assignNames(ctx.parameters, "")
	assignNames(ctx.variables, "")
	assignNames(ctx.outputs, "")
	assignNames(ctx.resources, "Resource")
}

func parseJsonxTree(text string) (*jsonx.Node, error) {
	root, errs := jsonx.ParseTree(text, jsonx.ParseOptions{Comments: true, TrailingCommas: true})
	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to parse JSON: %v", errs)
	}
	return root, nil
}
