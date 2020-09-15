package gen

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/pcl"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	exprRegex     = regexp.MustCompile(`^\[(.*)\]$`)
	funcExprRegex = regexp.MustCompile(`(?P<funcName>\w+)\((?P<funcParams>.*)\)(?P<funcProps>[\.|\[]?.*)`)
)

func isTemplateExpression(s string) bool {
	return exprRegex.MatchString(s)
}

func isFunctionCall(s string) (matched bool, funcName, funcParams, funcProps string) {
	matched = funcExprRegex.MatchString(s)
	if !matched {
		return
	}

	subMatches := funcExprRegex.FindStringSubmatch(s)

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
	return
}

// EvalGlobal takes the potential expression string as input and evaluates
// it to the most basic level possible given current knowledge. So,
// the first returned result may be a variable reference model expression if
// it refers to an existing (or new variable). If not an expression, the input
// is returned unmodified.
func (t *TemplateElements) EvalGlobal(expr string, exclusions codegen.StringSet) (interface{}, codegen.StringSet, error) {
	matches := exprRegex.FindStringSubmatch(expr)
	if len(matches) == 2 {
		expr = matches[1]
		expr = strings.TrimSpace(expr)
	}

	var funcName, funcParams, funcProps string
	funcMatch, funcName, funcParams, funcProps := isFunctionCall(expr)

	if exclusions.Has(funcName) {
		return expr, codegen.NewStringSet(), nil
	}

	if funcMatch {
		switch funcName {
		// Catch some special cases, with referenced properties, e.g. resourceGroup().location
		case "resourceGroup":
			if _, ok := t.elements["resourceGroupNameParam"]; !ok {
				if err := t.AddParameter("resourceGroupNameParam", map[string]interface{}{
					"type": "string",
				}); err != nil {
					return nil, nil, err
				}
			}

			v := &model.Variable{
				Name:         "resourceGroupNameParam",
				VariableType: model.DynamicType,
			}

			invoke := pcl.Invoke("azure-nextgen:resources/latest:getResourceGroup", pcl.ObjectConsItem("resourceGroupName", model.VariableReference(v)))
			referenced := codegen.StringSet{"resourceGroupNameParam": struct{}{}}
			funcProps := strings.TrimPrefix(funcProps, ".")
			if len(funcProps) > 0 {
				return pcl.RelativeTraversal(invoke, funcProps), referenced, nil
			}
			return invoke, referenced, nil
		case "subscription":
			if funcProps == ".subscriptionid" {
				/* TODO invoke for subscription? */
			}
			return nil, nil, fmt.Errorf("NYI: deployment functions not supported")

		case "deployment":
			return nil, nil, fmt.Errorf("NYI: deployment functions not supported")

		case "variables":
			params, referenced, err := t.EvalGlobal(funcParams, exclusions)
			if err != nil {
				return nil, nil, err
			}
			paramsStr, ok := params.(string)
			if !ok {
				return nil, nil, fmt.Errorf("expect variables function params to be a string, received: %T", params)
			}
			paramsStr = strings.Trim(paramsStr, "'")
			props, ref, err := t.evalVariable(paramsStr, funcProps, exclusions)
			if err != nil {
				return nil, nil, err
			}
			for k := range ref {
				referenced.Add(k)
			}
			return props, referenced, nil

		case "parameters":
			params, referenced, err := t.EvalGlobal(funcParams, exclusions)
			if err != nil {
				return nil, nil, err
			}

			paramsStr, ok := params.(string)
			if !ok {
				return nil, nil, fmt.Errorf("expect parameters function params to be a string, received: %T", params)
			}
			paramsStr = strings.Trim(paramsStr, "'")
			props, ref, err := t.evalParameter(paramsStr, funcProps, exclusions)
			if err != nil {
				return nil, nil, err
			}
			for k := range ref {
				referenced.Add(k)
			}
			return props, referenced, nil
		case "resourceId":
			// TODO: Add subscriptionid and resource group id to the concat list
			return t.evalConcat(funcParams, exclusions)
		case "reference":
			referenced := codegen.NewStringSet()
			params := parseParams(funcParams)
			var referencedResource interface{}
			var version string
			for _, p := range params {
				refParam, refs, err := t.EvalGlobal(p, exclusions)
				if err != nil {
					return nil, nil, err
				}
				for r := range refs {
					referenced.Add(r)
				}
				if referencedResource == nil {
					referencedResource = refParam
				} else if _, ok := refParam.(string); ok && version == "" {
					version = refParam.(string)
				} else {
					// ignore [full] since we always get the full resource
				}
			}

			referencedResStr, ok := referencedResource.(string)
			if ok {
				if isResourceId(referencedResStr) {
					// TODO: Provider needs support for
					// https://github.com/Azure/azure-rest-api-specs/blob/master/specification/resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json#L3445
					return nil, nil, fmt.Errorf("NYI: resource by id function not supported")
				}
			}
			// otherwise this is a resource name literal or a reference to a variable
			// (parameter) which indirectly refers to the resource name.
			//
			// We need to do a slow search through all known resources to find a match.
			// This also means that this search will only work in the final evaluation pass -
			// i.e. all resources have already been fully evaluated.
			found := false
			for k, v := range t.elements {
				switch v.TemplateElement.(type) {
				case *resource:
					args := v.Args()
					argsMap, ok := args.(map[string]interface{})
					if !ok {
						return nil, nil, fmt.Errorf("invalid payload for resource %s: %T", k, args)
					}
					name, ok := argsMap["name"]
					if !ok {
						continue
					}
					if reflect.DeepEqual(name, referencedResource) {
						referencedResource = model.VariableReference(&model.Variable{
							Name:         k,
							VariableType: model.DynamicType,
						})
						referenced.Add(k)
						found = true
						break
					}

				}
			}
			if !found {
				// TODO: Need the ability to get by resourceId
				// https://github.com/Azure/azure-rest-api-specs/blob/master/specification/resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json#L3445
				return nil, nil, fmt.Errorf("NYI: resource by id function not supported")
			}

			expr, ok := referencedResource.(model.Expression)
			if !ok {
				return nil, nil, fmt.Errorf("expected expression, received: %T", referencedResource)
			}
			funcProps := strings.TrimPrefix(funcProps, ".")
			if len(funcProps) > 0 {
				return pcl.RelativeTraversal(expr, funcProps), referenced, nil
			}
			return referencedResource, referenced, nil
		case "uniquestring":
			return nil, nil, fmt.Errorf("NYI: deployment functions not supported")
		case "concat":
			return t.evalConcat(funcParams, exclusions)
		}
	}
	return expr, codegen.NewStringSet(), nil
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

func (t *TemplateElements) evalConcat(params string, exclusions codegen.StringSet) (interface{}, codegen.StringSet, error) {
	paramList := parseParams(params)
	referenced := codegen.NewStringSet()
	var tuple []model.Expression
	for _, p := range paramList {
		res, ref, err := t.EvalGlobal(p, exclusions)
		if err != nil {
			return nil, nil, err
		}
		for k := range ref {
			referenced.Add(k)
		}
		x, err := pcl.RenderValue(res)
		if err != nil {
			return nil, nil, err
		}
		tuple = append(tuple, x)
	}
	tupleCons := &model.TupleConsExpression{
		Expressions: tuple,
	}
	return &model.TemplateJoinExpression{
		Tuple: tupleCons,
	}, referenced, nil
}

func (t *TemplateElements) evalVariable(varName string, propAccessor string, exclusions codegen.StringSet) (interface{}, codegen.StringSet, error) {
	propAccessor = strings.TrimSpace(propAccessor)

	if !strings.HasSuffix(varName, "Var") {
		varName = varName + "Var"
	}
	if _, found := t.elements[varName]; !found {
		return nil, nil, fmt.Errorf("missing variable: %s", varName)
	}

	return t.resolveReference(varName, propAccessor, exclusions)
}

func (t *TemplateElements) evalParameter(paramName string, propAccessor string, exclusions codegen.StringSet) (interface{}, codegen.StringSet, error) {
	propAccessor = strings.TrimSpace(propAccessor)

	if !strings.HasSuffix(paramName, "Param") {
		paramName = paramName + "Param"
	}
	if _, found := t.elements[paramName]; !found {
		return nil, nil, fmt.Errorf("missing variable: %s", paramName)
	}

	return t.resolveReference(paramName, propAccessor, exclusions)
}

func (t *TemplateElements) resolveReference(varName string, propAccessor string, exclusions codegen.StringSet) (interface{}, codegen.StringSet, error) {
	v := &model.Variable{
		Name:         varName,
		VariableType: model.DynamicType,
	}

	referenced := codegen.NewStringSet()
	referenced.Add(varName)
	if len(propAccessor) > 0 && propAccessor[0] == '[' && propAccessor[len(propAccessor)-1] == ']' {
		accessor := propAccessor[1 : len(propAccessor)-1]
		resolved, ref, err := t.EvalGlobal(accessor, exclusions)
		if err != nil {
			return nil, referenced, err
		}

		for k := range ref {
			referenced.Add(k)
		}

		resolvedExpr, ok := resolved.(model.Expression)
		if ok {
			return &model.IndexExpression{
				Collection: model.VariableReference(v),
				Key:        resolvedExpr,
			}, referenced, nil
		}

		resolvedStr, ok := resolved.(string)
		if ok {
			num, err := strconv.Atoi(resolvedStr)
			if err != nil {
				// Regular string?
				key, err := pcl.RenderValue(resolvedStr)
				if err != nil {
					return nil, nil, err
				}
				return &model.IndexExpression{
					Collection: model.VariableReference(v),
					Key:        key,
				}, referenced, nil
			}
			numRendered, err := pcl.RenderValue(num)
			if err != nil {
				return nil, nil, err
			}
			return &model.IndexExpression{
				Collection: model.VariableReference(v),
				Key:        numRendered,
			}, referenced, nil

		}
	} else if len(propAccessor) > 0 {
		propAccessor = propAccessor[1:] // Trim the leading '.'
		return &model.RelativeTraversalExpression{
			Source:    model.VariableReference(v),
			Traversal: hcl.Traversal{hcl.TraverseAttr{Name: propAccessor}},
			Parts:     []model.Traversable{model.DynamicType, model.DynamicType},
		}, referenced, nil
	}

	return model.VariableReference(v), referenced, nil
}

var isResourceId = regexp.MustCompile(`(?i)/subscriptions/(.+?)/resourcegroups/(.+?)/providers/Microsoft.(.*)/(.*)/(.*)`).MatchString
