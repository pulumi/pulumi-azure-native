package gen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/pcl"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

var (
	exprRegex     = regexp.MustCompile(`^\[(.*)\]$`)
	funcExprRegex = regexp.MustCompile(`(?P<funcName>\w+)\((?P<funcParams>.*)\)(?P<funcProps>[?:\.|\[].*)`)
)

func isTemplateExpression(s string) bool {
	return exprRegex.MatchString(s)
}

// EvalGlobal takes the potential expression string as input and evaluates
// it to the most basic level possible given current knowledge. So,
// the first returned result may be a variable reference model expression if
// it refers to an existing (or new variable). If not an expression, the input
// is returned unmodified.
func (t *TemplateElements) EvalGlobal(expr string) (interface{}, error) {
	matches := exprRegex.FindStringSubmatch(expr)
	contract.Assert(len(matches) == 2)
	expr = matches[1]
	expr = strings.TrimSpace(expr)

	if funcExprRegex.MatchString(expr) {
		var funcName, funcParams, funcProps string
		subMatches := funcExprRegex.FindStringSubmatch(expr)

		for i, name := range funcExprRegex.SubexpNames() {
			if i != 0 {
				switch name {
				case "funcName":
					funcName = subMatches[i]
				case "funcParams":
					funcParams = subMatches[i]
				case "funcProps":
					funcProps = subMatches[i]
				}
			}
		}

		switch funcName {
		// Catch some special cases, with referenced properties, e.g. resourceGroup().location
		case "resourceGroup":
			v := &model.Variable{
				Name:         "resourceGroup",
				VariableType: model.DynamicType,
			}
			/* TODO invoke get on resource group and emit a reference to the containing variable with traversal for prop */
			if _, ok := t.elements["resourceGroup"]; !ok {
				if err := t.AddParameter("resourceGroup", map[string]interface{}{
					"type": "string",
				}); err != nil {
					return nil, err
				}
			}
			
			return model.VariableReference(v), nil
		case "subscription":
			if funcProps == ".subscriptionid" {
				/* TODO invoke for subscription? */
				return nil, fmt.Errorf("NYI: deployment functions not supported")
			}

		case "deployment":
			return nil, fmt.Errorf("NYI: deployment functions not supported")

		case "variables":
			params, err := t.EvalGlobal(funcParams)
			if err != nil {
				return nil, err
			}
			paramsStr, ok := params.(string)
			if !ok {
				return nil, fmt.Errorf("expect variables function params to be a string, received: %T", params)
			}
			return t.evalVariable(paramsStr, funcProps)

		case "parameters":
			params, err := t.EvalGlobal(funcParams)
			if err != nil {
				return nil, err
			}

			paramsStr, ok := params.(string)
			if !ok {
				return nil, fmt.Errorf("expect parameters function params to be a string, received: %T", params)
			}
			return t.evalParameter(paramsStr, funcProps)
		case "uniquestring":
			return nil, fmt.Errorf("NYI: deployment functions not supported")
		case "concat":
			return t.evalConcat(funcParams)
		}
	}
	return nil, nil
}

func parseParams(paramString string) []string {
	depth := 0
	var parts []string
	lastSplit := 0
	for i := 0; i < len(paramString); i++ {
		c := paramString[i] // paramString[i];
		if c == '(' {
			depth++
		}
		if c == ')' {
			depth--
		}

		endOfString := i == len(paramString)-1
		if (c == ',' && depth == 0) || endOfString {
			endPoint := i
			if endOfString {
				endPoint = len(paramString)
			}
			parts = append(parts, strings.TrimSpace(paramString[lastSplit:endPoint]))
			lastSplit = i + 1
		}
	}
	return parts
}

func (t *TemplateElements) evalConcat(params string) (interface{}, error) {
	paramList := parseParams(params)

	var tuple []model.Expression
	for _, p := range paramList {
		res, err := t.evalVariable(p, "")
		if err != nil {
			return nil, err
		}
		x, err := pcl.RenderValue(res)
		if err != nil {
			return nil, err
		}
		tuple = append(tuple, x)
	}
	tupleCons := &model.TupleConsExpression{
		Expressions: tuple,
	}
	return &model.TemplateJoinExpression{
		Tuple: tupleCons,
	}, nil
}

func (t *TemplateElements) evalVariable(varName string, propAccessor string) (interface{}, error) {
	propAccessor = strings.TrimSpace(propAccessor)

	if _, found := t.elements[varName]; !found {
		return nil, fmt.Errorf("missing variable: %s", varName)
	}

	return t.resolveReference(varName, propAccessor)
}

func (t *TemplateElements) evalParameter(paramName string, propAccessor string) (interface{}, error) {
	propAccessor = strings.TrimSpace(propAccessor)

	if _, found := t.elements[paramName]; !found {
		return nil, fmt.Errorf("missing variable: %s", paramName)
	}

	return t.resolveReference(paramName, propAccessor)
}

func (t *TemplateElements) resolveReference(varName string, propAccessor string) (interface{}, error) {
	v := &model.Variable{
		Name:         varName,
		VariableType: model.DynamicType,
	}

	if len(propAccessor) > 0 && propAccessor[0] == '[' && propAccessor[len(propAccessor)-1] == ']' {
		accessor := propAccessor[1 : len(propAccessor)-1]
		resolved, err := t.EvalGlobal(accessor)
		if err != nil {
			return nil, err
		}

		resolvedExpr, ok := resolved.(model.Expression)
		if ok {
			return &model.IndexExpression{
				Collection: model.VariableReference(v),
				Key:        resolvedExpr,
			}, nil
		}

		resolvedStr, ok := resolved.(string)
		if ok {
			num, err := strconv.Atoi(resolvedStr)
			if err != nil {
				// Regular string?
				key, err := pcl.RenderValue(resolvedStr)
				if err != nil {
					return nil, err
				}
				return &model.IndexExpression{
					Collection: model.VariableReference(v),
					Key:        key,
				}, nil
			}
			numRendered, err := pcl.RenderValue(num)
			if err != nil {
				return nil, err
			}
			return &model.IndexExpression{
				Collection: model.VariableReference(v),
				Key:        numRendered,
			}, nil

		}
	} else if len(propAccessor) > 0 {
		return &model.RelativeTraversalExpression{
			Source:    model.VariableReference(v),
			Traversal: hcl.Traversal{hcl.TraverseAttr{Name: propAccessor}},
			Parts:     []model.Traversable{model.DynamicType, model.DynamicType},
		}, nil
	}

	return v, nil
}
