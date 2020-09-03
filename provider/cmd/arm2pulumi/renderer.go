package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/goccy/go-yaml/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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

	for _, root := range templates {
		if root.Type != jsonx.Object {
			continue
		}
		ctx := &renderContext{
			parameters: map[string]*model.Variable{},
			variables:  map[string]*model.Variable{},
			functions:  map[string]*model.Variable{},
			resources:  map[string]*model.Variable{},
			outputs:    map[string]*model.Variable{},
		}

		// Declare parameters, mappings, conditions, resources, and outputs.
		rootValue := jsonx.NodeValue(*root)
		for key, f := range rootValue.(map[string]interface{}) {
			var table map[string]*model.Variable
			switch key {
			case "parameters":
				ctx.detectPseudoParameters(f.Value)
				if values, ok := mapValues(f.Value); ok {
					for _, f := range values {
						paramName := keyString(f)
						if isSSMParameter(f.Value) {
							ctx.ssmParameters[paramName] = nil
						}
						ctx.parameters[paramName] = nil
					}
				}
				continue
			case "Mappings":
				table = ctx.mappings
			case "Conditions":
				table = ctx.conditions
			case "Resources":
				table = ctx.resources
			case "Outputs":
				table = ctx.outputs
			default:
				continue
			}
			ctx.detectPseudoParameters(f.Value)
			if values, ok := mapValues(f.Value); ok {
				for _, f := range values {
					table[keyString(f)] = nil
				}
			}
		}
		ctx.assignNames()

		// Swap in SSM parameter value names for their config names.
		for k, value := range ctx.ssmParameters {
			config, ok := ctx.parameters[k]
			contract.Assert(ok)
			ctx.parameters[k], ctx.ssmParameters[k] = value, config
		}

		var items []model.BodyItem
		for _, f := range root.Values {
			switch keyString(f) {
			case "AWSTemplateFormatVersion":
				// Ignore this
			case "Description":
				if f.Value.Type() != ast.StringType {
					return nil, fmt.Errorf("Description must be a string")
				}
				//comment = f.Value.(jsonast.String).String()
			case "Metadata":
				// TODO: render metadata
			case "Parameters":
				ctx.renderPseudoParameters(&items)

				values, ok := mapValues(f.Value)
				if !ok {
					return nil, fmt.Errorf("Parameters must be a map")
				}

				for _, f := range values {
					parameter, err := ctx.renderParameter(f)
					if err != nil {
						return nil, err
					}
					items = append(items, parameter...)
				}
			case "Mappings":
				mappings, err := ctx.renderObjects(f, ctx.renderMapping)
				if err != nil {
					return nil, err
				}
				items = append(items, mappings...)
			case "Conditions":
				ctx.renderPseudoParameters(&items)

				conditions, err := ctx.renderObjects(f, ctx.renderCondition)
				if err != nil {
					return nil, err
				}
				items = append(items, conditions...)
			case "Transform":
				return nil, fmt.Errorf("template transforms are not supported")
			case "Resources":
				ctx.renderPseudoParameters(&items)

				resources, err := ctx.renderObjects(f, ctx.renderResource)
				if err != nil {
					return nil, err
				}
				items = append(items, resources...)
			case "Outputs":
				ctx.renderPseudoParameters(&items)

				outputs, err := ctx.renderObjects(f, ctx.renderOutput)
				if err != nil {
					return nil, err
				}
				items = append(items, outputs...)
			default:
				return nil, fmt.Errorf("unexpected template property %v", f.Key)
			}
		}
	}

	body := &model.Body{Items: items}
	formatBody(body)
	return body, nil
}

// A renderContext contains the state necessary to render an Arm template to a Pulumi program.
type renderContext struct {
	// The maps below map from names in the ARM template to output variables. These maps are used during
	// rendering to emit appropriate references to the entities.

	parameters map[string]*model.Variable // Template parameters
	variables  map[string]*model.Variable // Template variables
	functions  map[string]*model.Variable // Template user-defined functions
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

// pseudoParameters is the set of CFN pseudo-parameters.
var pseudoParameters = codegen.NewStringSet("AWS::NoValue",
	"AWS::AccountId",
	"AWS::Partition",
	"AWS::Region",
	"AWS::StackName",
	"AWS::StackId",
	"AWS::URLSuffix")

// functions is the set of CFN functions.
var functions = codegen.NewStringSet("Condition",
	"Fn::And",
	"Fn::Base64",
	"Fn::Cidr",
	"Fn::Equals",
	"Fn::FindInMap",
	"Fn::GetAtt",
	"Fn::GetAZs",
	"Fn::If",
	"Fn::ImportValue",
	"Fn::Join",
	"Fn::Not",
	"Fn::Or",
	"Fn::Select",
	"Fn::Split",
	"Fn::Sub",
	"Fn::Transform",
	"Ref")

// functionTags maps from tags that represent the short forms of functions to their corresponding function names.
var functionTags = map[string]string{
	"!Condition":   "Condition",
	"!And":         "Fn::And",
	"!Base64":      "Fn::Base64",
	"!Cidr":        "Fn::Cidr",
	"!Equals":      "Fn::Equals",
	"!FindInMap":   "Fn::FindInMap",
	"!GetAtt":      "Fn::GetAtt",
	"!GetAZs":      "Fn::GetAZs",
	"!If":          "Fn::If",
	"!ImportValue": "Fn::ImportValue",
	"!Join":        "Fn::Join",
	"!Not":         "Fn::Not",
	"!Or":          "Fn::Or",
	"!Select":      "Fn::Select",
	"!Split":       "Fn::Split",
	"!Sub":         "Fn::Sub",
	"!Transform":   "Fn::Transform",
	"!Ref":         "Ref",
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

// renderRef renders a reference to a pseudo-parameter, parameter, or resource. These entities all correspond to
// top-level variables in the Pulumi program, so each reference is rendered as a scope traversal expression.
func (ctx *renderContext) renderRef(name string) (model.Expression, error) {
	if pseudoParameters.Has(name) {
		pseudoVar, ok := ctx.pseudoParameters[name]
		contract.Assert(ok)
		return model.VariableReference(pseudoVar), nil
	}

	v, ok := ctx.parameters[name]
	if ok {
		return model.VariableReference(v), nil
	}

	v, ok = ctx.resources[name]
	if !ok {
		return nil, fmt.Errorf("unknown parameter or resource '%v'", name)
	}
	return &model.ScopeTraversalExpression{
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: v.Name}, hcl.TraverseAttr{Name: "id"}},
		Parts:     []model.Traversable{v, model.DynamicType},
	}, nil
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
			v, err := ctx.renderValue(f.Value)
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
		parts = append(parts, plainLit(l))
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

// renderJoin renders a call to Fn::Join. If the arguments to the join are a string literal and a sequence, the call
// is rendered as a template expression. Otherwise, it is rendered as a call to the `join` function.
func (ctx *renderContext) renderJoin(name string, arg ast.Node) (model.Expression, error) {
	args, err := ctx.renderArgsArray(name, arg, 2, 2)
	if err != nil {
		return nil, err
	}

	call := &model.FunctionCallExpression{
		Name: "join",
		Args: args,
	}

	// Turn `join("lit", [ ... ])` into a template expression.
	sep, ok := getStringValue(args[0])
	if !ok {
		return call, nil
	}

	tuple, ok := args[1].(*model.TupleConsExpression)
	if !ok {
		return call, nil
	}

	afterLiteral := false
	var parts []model.Expression
	for _, x := range tuple.Expressions {
		lit, ok := getStringValue(x)
		if ok {
			if afterLiteral {
				last := parts[len(parts)-1].(*model.LiteralValueExpression)
				last.Value = cty.StringVal(last.Value.AsString() + sep + lit)
			} else {
				if len(parts) > 0 {
					lit = sep + lit
				}
				parts = append(parts, plainLit(lit))
				afterLiteral = true
			}
		} else {
			if afterLiteral {
				last := parts[len(parts)-1].(*model.LiteralValueExpression)
				last.Value = cty.StringVal(last.Value.AsString() + sep + lit)
			} else {
				parts = append(parts, plainLit(sep))
			}
			parts = append(parts, x)
			afterLiteral = false
		}
	}
	if !afterLiteral {
		parts = append(parts, plainLit(""))
	}
	return &model.TemplateExpression{Parts: parts}, nil
}

// renderFunctionCall renders a call to an AWS intrinsic function. The way the function is rendered depends on the
// function:
//
// - `Condition` is rendered as a reference to the equivalent condition variable
// - `Fn::And` is rendered as the logical and operator
// - `Fn::Base64` is rendered as a call to the `toBase64` function
// - `Fn::Cidr` is rendered as an invocation of `cloudformation::cidr`
// - `Fn::Equals` is rendered as the equality operator
// - `Fn::FindInMap` is rendered as a two index expressions
// - `Fn::GetAtt` is rendered as a scope traversal expression on the referenced resource to fetch the referenced
//   attribute
// - `Fn::GetAZs` is rendered as an invocation of `cloudformation::getAzs`
// - `Fn::If` is rendered as the conditional operator
// - `Fn::ImportValue` is rendered as an invocation of `cloudformation::importValue`
// - `Fn::Join` is rendered as either a template expression or a call to `join`
// - `Fn::Not` is rendered as the logical not operator
// - `Fn::Or` is rendered as the logical or operator
// - `Fn::Select` is rendered as an index expressions
// - `Fn::Split` is rendered as a call to `split`
// - `Fn::Sub` is rendered as a template expression
// - `Fn::Transform` is not supported
// - `Ref` is rendered as a variable reference
//
func (ctx *renderContext) renderFunctionCall(name string, arg ast.Node) (model.Expression, error) {
	switch name {
	case "Condition":
		arg, err := ctx.renderValue(arg)
		if err != nil {
			return nil, err
		}
		condition, ok := getStringValue(arg)
		if !ok {
			return nil, fmt.Errorf("the argument to Condition must be a string")
		}
		condVar, ok := ctx.conditions[condition]
		if !ok {
			return nil, fmt.Errorf("unknown condition '%v'", condition)
		}
		return model.VariableReference(condVar), nil
	case "Fn::And":
		args, err := ctx.renderArgsArray(name, arg, 2, -1)
		if err != nil {
			return nil, err
		}

		leftOperand := args[0]
		for _, arg := range args[1:] {
			leftOperand = &model.BinaryOpExpression{
				LeftOperand:  leftOperand,
				Operation:    hclsyntax.OpLogicalAnd,
				RightOperand: arg,
			}
		}
		return leftOperand, nil
	case "Fn::Base64":
		arg, err := ctx.renderValue(arg)
		if err != nil {
			return nil, err
		}
		return &model.FunctionCallExpression{
			Name: "toBase64",
			Args: []model.Expression{
				arg,
			},
		}, nil
	case "Fn::Cidr":
		args, err := ctx.renderArgsArray(name, arg, 3, 3)
		if err != nil {
			return nil, err
		}

		inputs := []model.ObjectConsItem{
			objectConsItem("ipBlock", args[0]),
			objectConsItem("count", args[1]),
			objectConsItem("cidrBits", args[2]),
		}
		return relativeTraversal(invoke("cloudformation::cidr", inputs...), "subnets"), nil
	case "Fn::Equals":
		args, err := ctx.renderArgsArray(name, arg, 2, 2)
		if err != nil {
			return nil, err
		}
		return &model.BinaryOpExpression{
			LeftOperand:  args[0],
			Operation:    hclsyntax.OpEqual,
			RightOperand: args[1],
		}, nil
	case "Fn::FindInMap":
		args, err := ctx.renderArgsArray(name, arg, 3, 3)
		if err != nil {
			return nil, err
		}

		mapName, ok := getStringValue(args[0])
		if !ok {
			return nil, fmt.Errorf("the first argument to 'Fn::FindInMap' must be a string")
		}
		mapVar, ok := ctx.mappings[mapName]
		if !ok {
			return nil, fmt.Errorf("unknown mapping '%v'", mapName)
		}
		return &model.IndexExpression{
			Collection: &model.IndexExpression{
				Collection: model.VariableReference(mapVar),
				Key:        args[1],
			},
			Key: args[2],
		}, nil
	case "Fn::GetAtt":
		args, err := ctx.renderArgsArray(name, arg, 1, 2)
		if err != nil {
			return nil, err
		}

		var resourceName string
		var attrArg model.Expression
		if len(args) == 1 {
			attrPath, ok := getStringValue(args[0])
			if !ok {
				return nil, fmt.Errorf("the argument to 'Fn::GetAtt' must be a string")
			}
			dotIndex := strings.Index(attrPath, ".")
			if dotIndex == -1 {
				return nil, fmt.Errorf("attribute paths must be of the form \"resourceName.attrName\"")
			}
			resourceName, attrArg = attrPath[:dotIndex], plainLit(attrPath[dotIndex+1:])
		} else {
			rn, ok := getStringValue(args[0])
			if !ok {
				return nil, fmt.Errorf("the first argument to 'Fn::GetAtt' must be a string")
			}
			resourceName, attrArg = rn, args[1]
		}

		resourceVar, ok := ctx.resources[resourceName]
		if !ok {
			return nil, fmt.Errorf("unknown resource '%v'", resourceName)
		}

		attrName, ok := getStringValue(attrArg)
		if !ok {
			return &model.IndexExpression{
				Collection: &model.ScopeTraversalExpression{
					Traversal: hcl.Traversal{
						hcl.TraverseRoot{Name: resourceVar.Name},
						hcl.TraverseAttr{Name: "attributes"},
					},
					Parts: []model.Traversable{
						resourceVar,
						model.DynamicType,
					},
				},
				Key: args[1],
			}, nil
		}

		return &model.ScopeTraversalExpression{
			Traversal: hcl.Traversal{
				hcl.TraverseRoot{Name: resourceVar.Name},
				hcl.TraverseAttr{Name: "attributes"},
				hcl.TraverseAttr{Name: strings.Replace(attrName, ".", "", -1)},
			},
			Parts: []model.Traversable{
				resourceVar,
				model.DynamicType,
				model.DynamicType,
			},
		}, nil
	case "Fn::GetAZs":
		region, err := ctx.renderValue(arg)
		if err != nil {
			return nil, err
		}
		return relativeTraversal(invoke("cloudformation::getAzs", objectConsItem("region", region)), "azs"), nil
	case "Fn::If":
		args, err := ctx.renderArgsArray(name, arg, 3, 3)
		if err != nil {
			return nil, err
		}

		condition, ok := getStringValue(args[0])
		if !ok {
			return nil, fmt.Errorf("the first argument to 'Fn::If' must be a string")
		}
		condVar, ok := ctx.conditions[condition]
		if !ok {
			return nil, fmt.Errorf("unknown condition '%v'", condition)
		}
		return &model.ConditionalExpression{
			Condition:   model.VariableReference(condVar),
			TrueResult:  args[1],
			FalseResult: args[2],
		}, nil
	case "Fn::ImportValue":
		arg, err := ctx.renderValue(arg)
		if err != nil {
			return nil, err
		}
		return relativeTraversal(invoke("cloudformation::importValue", objectConsItem("name", arg)), "value"), nil
	case "Fn::Join":
		return ctx.renderJoin(name, arg)
	case "Fn::Not":
		args, err := ctx.renderArgsArray(name, arg, 1, 1)
		if err != nil {
			return nil, err
		}
		return &model.UnaryOpExpression{
			Operation: hclsyntax.OpLogicalNot,
			Operand:   args[0],
		}, nil
	case "Fn::Or":
		args, err := ctx.renderArgsArray(name, arg, 2, -1)
		if err != nil {
			return nil, err
		}

		leftOperand := args[0]
		for _, arg := range args[1:] {
			leftOperand = &model.BinaryOpExpression{
				LeftOperand:  leftOperand,
				Operation:    hclsyntax.OpLogicalOr,
				RightOperand: arg,
			}
		}
		return leftOperand, nil
	case "Fn::Select":
		args, err := ctx.renderArgsArray(name, arg, 2, 2)
		if err != nil {
			return nil, err
		}

		indexStr, ok := getStringValue(args[0])
		if ok {
			indexInt, err := convert.Convert(cty.StringVal(indexStr), cty.Number)
			if err == nil {
				args[0] = &model.LiteralValueExpression{
					Value: indexInt,
				}
			} else {
				args[0] = quotedLit(indexStr)
			}
		}

		return &model.IndexExpression{
			Collection: args[1],
			Key:        args[0],
		}, nil
	case "Fn::Split":
		args, err := ctx.renderArgsArray(name, arg, 2, 2)
		if err != nil {
			return nil, err
		}

		return &model.FunctionCallExpression{
			Name: "split",
			Args: args,
		}, nil
	case "Fn::Sub":
		return ctx.renderSub(name, arg)
	case "Fn::Transform":
		return nil, fmt.Errorf("NYI: Fn::Transform")
	case "Ref":
		arg, err := ctx.renderValue(arg)
		if err != nil {
			return nil, err
		}

		entityName, ok := getStringValue(arg)
		if !ok {
			return nil, fmt.Errorf("the argument to 'Ref' must be a string")
		}

		return ctx.renderRef(entityName)
	default:
		contract.Failf("unexpected function %v", name)
		return nil, nil
	}
}

// renderValue renders an AST node that represents a YAML value as its equivalent PCL. Most nodes are rendered as one
// would expect (e.g. sequences -> tuple construction, maps -> object construction, etc.). Function calls are the lone
// exception; see renderFunction for more details.
func (ctx *renderContext) renderValue(node ast.Node) (model.Expression, error) {
	switch node := node.(type) {
	case *ast.SequenceNode:
		var expressions []model.Expression
		for _, v := range node.Values {
			e, err := ctx.renderValue(v)
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, e)
		}
		return &model.TupleConsExpression{
			Expressions: expressions,
		}, nil
	case *ast.BoolNode:
		return &model.LiteralValueExpression{
			Value: cty.BoolVal(node.Value),
		}, nil
	case *ast.NullNode:
		return model.VariableReference(null), nil
	case *ast.FloatNode:
		return &model.LiteralValueExpression{
			Value: cty.NumberFloatVal(node.Value),
		}, nil
	case *ast.IntegerNode:
		var value cty.Value
		switch v := node.Value.(type) {
		case int64:
			value = cty.NumberIntVal(v)
		case uint64:
			value = cty.NumberUIntVal(v)
		default:
			contract.Failf("unexpected value of type %T in integer node", v)
		}
		return &model.LiteralValueExpression{Value: value}, nil
	case *ast.LiteralNode:
		return ctx.renderValue(node.Value)
	case *ast.MappingNode, *ast.MappingValueNode:
		values, ok := mapValues(node)
		contract.Assert(ok)

		if len(values) == 1 && functions.Has(keyString(values[0])) {
			return ctx.renderFunctionCall(keyString(values[0]), values[0].Value)
		}

		var items []model.ObjectConsItem
		for _, f := range values {
			v, err := ctx.renderValue(f.Value)
			if err != nil {
				return nil, err
			}
			items = append(items, objectConsItem(keyString(f), v))
		}
		return &model.ObjectConsExpression{
			Items: items,
		}, nil
	case *ast.StringNode:
		return quotedLit(node.Value), nil
	case *ast.TagNode:
		fnName, ok := functionTags[node.Start.Value]
		if !ok {
			return nil, fmt.Errorf("unknown tag %v", node.Start.Value)
		}
		return ctx.renderFunctionCall(fnName, node.Value)
	default:
		contract.Failf("unexpected YAML node of type %T", node)
		return nil, nil
	}
}

// renderPseudoParameter renders a reference to a pseudo parameter as an invocation of the corresponding getter in the
// cloudformation provider (or a reference to `null` for `AWS::NoValue`).
func renderPseudoParameter(name string) model.Expression {
	getter := func(name string) model.Expression {
		return relativeTraversal(invoke("cloudformation::get"+name), camel(name))
	}

	switch name {
	case "AWS::NoValue":
		return model.VariableReference(null)
	case "AWS::AccountId", "AWS::Partition", "AWS::Region", "AWS::StackName", "AWS::StackId":
		return getter(name[strings.Index(name, "::")+2:])
	case "AWS::URLSuffix":
		return getter("UrlSuffix")
	default:
		contract.Failf("unexpected pseudo parameter '%v'", name)
		return nil
	}
}

// renderPseudoParameters renders all referenced pseudo parameters into PCL local variables whose values are
// invocations of the corresponding provider functions.
func (ctx *renderContext) renderPseudoParameters(items *[]model.BodyItem) {
	if !ctx.renderedPseudoParameters {
		for _, v := range pseudoParameters.SortedValues() {
			if pseudoVar, ok := ctx.pseudoParameters[v]; ok {
				*items = append(*items, &model.Attribute{
					Name:  pseudoVar.Name,
					Value: renderPseudoParameter(v),
				})
			}
		}
		ctx.renderedPseudoParameters = true
	}
}

// renderParameterType converts a CloudFormation parameter type to its equivalent PCL type.
func renderParameterType(s string) (string, bool) {
	switch s {
	case "String":
		return "string", true
	case "Number":
		return "number", true
	case "List<Number>":
		return "list(number)", true
	case "CommaDelimitedList", "List<String>":
		return "list(string)", true
	case "AWS::EC2::AvailabilityZone::Name",
		"AWS::EC2::Image::Id",
		"AWS::EC2::Instance::Id",
		"AWS::EC2::KeyPair::KeyName",
		"AWS::EC2::SecurityGroup::GroupName",
		"AWS::EC2::SecurityGroup::Id",
		"AWS::EC2::Subnet::Id",
		"AWS::EC2::Volume::Id",
		"AWS::EC2::VPC::Id",
		"AWS::Route53::HostedZone::Id",
		"AWS::SSM::Parameter::Name":
		return "string", true
	case "List<AWS::EC2::AvailabilityZone::Name>",
		"List<AWS::EC2::Image::Id>",
		"List<AWS::EC2::Instance::Id>",
		"List<AWS::EC2::SecurityGroup::GroupName>",
		"List<AWS::EC2::SecurityGroup::Id>",
		"List<AWS::EC2::Subnet::Id>",
		"List<AWS::EC2::Volume::Id>",
		"List<AWS::EC2::VPC::Id>",
		"List<AWS::Route53::HostedZone::Id>":
		return "list(string)", true
	default:
		return "", false
	}
}

var ssmPattern = regexp.MustCompile("AWS::SSM::Parameter::Value<(.*)>")

// renderParameter renders a template parameter. The parameter is rendered as either a sinmple config variable
// definition or a config variable definition and an invocation of one of the two SSM parameter getters (in the case of
// SSM parameters). Uses of SSM parameters reference the result of the invocation.
func (ctx *renderContext) renderParameter(attr *ast.MappingValueNode) ([]model.BodyItem, error) {
	name := keyString(attr)
	values, ok := mapValues(attr.Value)
	if !ok {
		return nil, fmt.Errorf("parameter '%s' must be a mapping", name)
	}

	typ, ok := valueAt(values, "Type")
	if !ok {
		return nil, fmt.Errorf("parameter '%s' is missing required field 'Type'", name)
	}
	if typ.Type() != ast.StringType {
		return nil, fmt.Errorf("'Type' of parameter '%s' must be a string", name)
	}

	typeValue, isSSMParameter := typ.(*ast.StringNode).Value, false
	if matches := ssmPattern.FindStringSubmatch(typeValue); len(matches) != 0 {
		typeValue, isSSMParameter = matches[1], true
	}

	typeExpr, ok := renderParameterType(typeValue)
	if !ok {
		return nil, fmt.Errorf("Unrecognized type '%v' for parameter '%s'", typeValue, name)
	}

	paramRefVar, ok := ctx.parameters[keyString(attr)]
	contract.Assert(ok)

	paramDefVar := paramRefVar
	if isSSMParameter {
		paramDefVar, ok = ctx.ssmParameters[keyString(attr)]
		contract.Assert(ok)

		typeExpr = "string"
	}

	var defaultValue model.Expression
	if def, ok := valueAt(values, "Default"); ok {
		v, err := ctx.renderValue(def)
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
	if !isSSMParameter {
		if defaultValue != nil {
			configDef.Body.Items = append(configDef.Body.Items, &model.Attribute{
				Name:  "default",
				Value: defaultValue,
			})
		}
		return []model.BodyItem{configDef}, nil
	}

	ssmGetter := "getSsmParameterString"
	switch typeExpr {
	case "number", "list(number)":
		return nil, fmt.Errorf("type '%v' is not supported for SSM parameters", typeValue)
	case "string":
		// OK
	default:
		ssmGetter = "getSsmParameterList"
	}

	var getParamValue model.Expression = relativeTraversal(
		invoke("cloudformation::"+ssmGetter,
			objectConsItem("name", model.VariableReference(paramDefVar))),
		"value")
	if defaultValue != nil {
		configDef.Body.Items = append(configDef.Body.Items, &model.Attribute{
			Name:  "default",
			Value: quotedLit(""),
		})

		getParamValue = &model.ConditionalExpression{
			Condition: &model.BinaryOpExpression{
				LeftOperand:  model.VariableReference(paramDefVar),
				Operation:    hclsyntax.OpEqual,
				RightOperand: quotedLit(""),
			},
			TrueResult:  defaultValue,
			FalseResult: getParamValue,
		}
	}
	return []model.BodyItem{
		configDef,
		&model.Attribute{
			Name:  paramRefVar.Name,
			Value: getParamValue,
		},
	}, nil
}

// renderMapping renders a CloudFormation mapping as a PCL local variable whose value is the value of the mapping.
func (ctx *renderContext) renderMapping(attr *ast.MappingValueNode) (model.BodyItem, error) {
	v, ok := ctx.mappings[keyString(attr)]
	contract.Assert(ok)

	m, err := ctx.renderValue(attr.Value)
	if err != nil {
		return nil, err
	}

	return &model.Attribute{
		Name:  v.Name,
		Value: m,
	}, nil
}

// renderCondition renders a CloudFormation condition as a PCL local variable whose value is the value of the
// condition.
func (ctx *renderContext) renderCondition(attr *ast.MappingValueNode) (model.BodyItem, error) {
	v, ok := ctx.conditions[keyString(attr)]
	contract.Assert(ok)

	m, err := ctx.renderValue(attr.Value)
	if err != nil {
		return nil, err
	}

	return &model.Attribute{
		Name:  v.Name,
		Value: m,
	}, nil
}

// renderResource renders a CloudFormation resource as a PCL resource.
func (ctx *renderContext) renderResource(attr *ast.MappingValueNode) (model.BodyItem, error) {
	name := keyString(attr)
	values, ok := mapValues(attr.Value)
	if !ok {
		return nil, fmt.Errorf("resource '%v' must be a mapping", name)
	}

	resourceVar, ok := ctx.resources[name]
	contract.Assert(ok)

	var token string
	var items []model.BodyItem
	for _, f := range values {
		switch keyString(f) {
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

// renderOutput renders a CloudFormation output as a PCL output.
func (ctx *renderContext) renderOutput(attr *ast.MappingValueNode) (model.BodyItem, error) {
	name := keyString(attr)
	values, ok := mapValues(attr.Value)
	if !ok {
		return nil, fmt.Errorf("output '%v' must be a mapping", name)
	}

	outputVar, ok := ctx.outputs[name]
	contract.Assert(ok)

	// TODO: description, export

	value, ok := valueAt(values, "Value")
	if !ok {
		return nil, fmt.Errorf("output '%v' has no \"Value\" attribute", name)
	}

	x, err := ctx.renderValue(value)
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

// An objectRenderer renders a mapping value node into a body item.
type objectRenderer func(*ast.MappingValueNode) (model.BodyItem, error)

// renderObjects is a helper that renders a set of named objects in a mapping. Each named object is passed to the
// provided renderer.
func (ctx *renderContext) renderObjects(attr *ast.MappingValueNode, render objectRenderer) ([]model.BodyItem, error) {
	values, ok := mapValues(attr.Value)
	if !ok {
		return nil, fmt.Errorf("%s must be a mapping", keyString(attr))
	}

	var items []model.BodyItem
	for _, f := range values {
		i, err := render(f)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

// assignNames assigns names to the variables used to represent template parameters outputs, SSM values, resources,
// mappings, conditions, and pseudo-parameters. Care is taken to keep parameter and output names as close to their
// original names as possible.
func (ctx *renderContext) assignNames() {
	assigned := codegen.StringSet{}

	assign := func(name, suffix string) *model.Variable {
		assignName := func(name, suffix string) string {
			name = camel(makeLegalIdentifier(name))
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
	assignNames(ctx.outputs, "")
	assignNames(ctx.ssmParameters, "Value")
	assignNames(ctx.resources, "Resource")
	assignNames(ctx.mappings, "Mapping")
	assignNames(ctx.conditions, "Condition")
	assignNames(ctx.pseudoParameters, "")
}

func isSSMParameter(node ast.Node) bool {
	values, ok := mapValues(node)
	if !ok {
		return false
	}
	typ, ok := valueAt(values, "Type")
	if !ok || typ.Type() != ast.StringType {
		return false
	}
	return ssmPattern.MatchString(typ.(*ast.StringNode).Value)
}

// detectFunctionPseudoParameters detects uses of pseudo-parameters in a function. This involves inspecting the holes
// in Fn::Sub template strings and the entities named in calls to Ref.
func (ctx *renderContext) detectFunctionPseudoParameters(name string, arg ast.Node) {
	switch name {
	case "Fn::Sub":
		if seq, ok := arg.(*ast.SequenceNode); ok {
			if len(seq.Values) == 0 {
				return
			}
			arg = seq.Values[0]

			for _, v := range seq.Values[1:] {
				ctx.detectPseudoParameters(v)
			}
		}
		str, ok := arg.(*ast.StringNode)
		if !ok {
			return
		}

		text := str.Value
		holes := holePattern.FindAllStringSubmatchIndex(text, -1)

		for _, hole := range holes {
			ref := text[hole[2]:hole[3]]
			if len(ref) > 0 && ref[0] != '!' && pseudoParameters.Has(ref) {
				ctx.pseudoParameters[ref] = nil
			}
		}
	case "Ref":
		if s, ok := arg.(*ast.StringNode); ok && pseudoParameters.Has(s.Value) {
			ctx.pseudoParameters[s.Value] = nil
		} else {
			ctx.detectPseudoParameters(arg)
		}
	default:
		ctx.detectPseudoParameters(arg)
	}
}

func parseJsonxTree(text string) (*jsonx.Node, error) {
	root, errs := jsonx.ParseTree(text, jsonx.ParseOptions{Comments: true, TrailingCommas: true})
	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to parse JSON: %v", errs)
	}
	return root, nil
}
