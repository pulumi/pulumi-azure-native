package arm2pulumi

import (
	"errors"
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/tle"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	exprRegex = regexp.MustCompile(`^\[(.*)\]$`)
)

func isTemplateExpression(s string) bool {
	return exprRegex.MatchString(unquote(s))
}

// Eval takes the potential expression string as input and evaluates
// it to the most basic level possible given current knowledge. So,
// the first returned result may be a variable reference if
// it refers to an existing (or new variable). If not an expression, the input
// is returned unmodified. The third argument - exclusions is an optional
// set of exclusion for template functions. If a template function matches
// something in the exclusion set, the resulting
func (t *TemplateElements) eval(
	source *dependencyTracking,
	expression string,
	exclusions codegen.StringSet,
) (interface{}, error) {
	result, err := tle.Parse(expression)
	if err != nil {
		return nil, err
	}

	val, err := t.evalValue(source, result.Expression, exclusions)
	if err != nil {
		var nyi *notYetImplementedError
		var excluded *excludedError
		if errors.As(err, &nyi) {
			t.addDiagnostic(Diagnostic{
				SourceElement: source.Name(),
				SourceToken:   expression,
				Severity:      SevMed,
				Description:   err.Error(),
			})
			return expression, nil
		}
		if errors.As(err, &excluded) {
			return expression, nil
		}
		return nil, err
	}
	return val, nil
}

func (t *TemplateElements) evalValue(
	source *dependencyTracking,
	value tle.Value,
	exclusions codegen.StringSet,
) (interface{}, error) {
	switch value.(type) {
	case *tle.StringValue:
		var str string
		if err := value.Accept(tle.StringValueVisitor(func(stringValue *tle.StringValue) error {
			str = unquote(stringValue.Token().StringValue())
			return nil
		})); err != nil {
			return nil, err
		}
		return str, nil
	case *tle.NumberValue:
		var num int
		if err := value.Accept(tle.NumberValueVisitor(func(numberValue *tle.NumberValue) error {
			var err error
			str := unquote(numberValue.Token().StringValue())
			num, err = strconv.Atoi(str)
			return err
		})); err != nil {
			return nil, err
		}
		return num, nil
	case *tle.ArrayAccessValue:
		var target interface{}
		if err := value.Accept(tle.ArrayAccessValueVisitor(t.evalArrayAccess(source, &target, exclusions))); err != nil {
			return nil, err
		}
		return target, nil
	case *tle.FunctionCallValue:
		var target interface{}
		if err := value.Accept(
			tle.FunctionCallValueVisitor(t.evalFunctionCall(source, &target, exclusions))); err != nil {
			return nil, err
		}
		return target, nil
	case *tle.PropertyAccessValue:
		var target interface{}
		if err := value.Accept(
			tle.PropertyAccessValueVisitor(t.evalPropertyAccess(source, &target, exclusions))); err != nil {
			return nil, err
		}
		return target, nil
	default:
		return nil, fmt.Errorf("unexpected value type: %T for value: %v", value, value)
	}
}

// nyiError Creates a new notYetImplementedError
func nyiError(functionName string) *notYetImplementedError {
	return &notYetImplementedError{function: functionName}
}

type notYetImplementedError struct {
	function string
}

func (e notYetImplementedError) Error() string {
	return fmt.Sprintf("NYI: '%s' not supported as yet", e.function)
}

type excludedError struct {
	function string
}

func (e excludedError) Error() string {
	return fmt.Sprintf("%s is currently excluded", e.function)
}

func (t *TemplateElements) evalArrayAccess(
	d *dependencyTracking,
	target *interface{},
	exclusions codegen.StringSet,
) func(*tle.ArrayAccessValue) error {
	return func(a *tle.ArrayAccessValue) error {
		srcEvaled, err := t.evalValue(d, a.Source(), exclusions)
		if err != nil {
			return err
		}
		srcExpr, ok := srcEvaled.(model.Expression)
		if !ok {
			return fmt.Errorf("expected array access source to be model.Expression, got %T", srcEvaled)
		}
		indexEvaled, err := t.evalValue(d, a.Index(), exclusions)
		if err != nil {
			return err
		}

		keyExpr, ok := indexEvaled.(model.Expression)
		if !ok {
			key, err := pcl.RenderValue(indexEvaled)
			if err != nil {
				return err
			}
			*target = &model.IndexExpression{
				Collection: srcExpr,
				Key:        key,
			}
			return nil
		}

		*target = &model.IndexExpression{
			Collection: srcExpr,
			Key:        keyExpr,
		}
		return nil
	}
}

func (t *TemplateElements) evalPropertyAccess(
	d *dependencyTracking,
	target *interface{},
	exclusions codegen.StringSet,
) func(*tle.PropertyAccessValue) error {
	return func(p *tle.PropertyAccessValue) error {
		srcEvaled, err := t.evalValue(d, p.Source(), exclusions)
		if err != nil {
			return err
		}
		propString := p.PropertyName()
		expr, ok := srcEvaled.(model.Expression)
		if !ok {
			return fmt.Errorf("expected property access source to be model.Expression, got %T", srcEvaled)
		}
		*target = &model.RelativeTraversalExpression{
			Source:    expr,
			Traversal: hcl.Traversal{hcl.TraverseAttr{Name: propString}},
			Parts:     []model.Traversable{model.DynamicType, model.DynamicType},
		}
		return nil
	}
}

func (t *TemplateElements) evalFunctionCall(
	d *dependencyTracking,
	target *interface{},
	exclusions codegen.StringSet,
) func(f *tle.FunctionCallValue) error {
	return func(f *tle.FunctionCallValue) error {
		if f.NamespaceToken != nil && f.NamespaceToken.StringValue() != "" {
			return fmt.Errorf("%w",
				nyiError(fmt.Sprintf(f.NamespaceToken.StringValue(), f.NameToken.StringValue())))
		}

		if exclusions.Has(f.NameToken.StringValue()) {
			return &excludedError{function: f.NameToken.StringValue()}
		}

		switch f.NameToken.StringValue() {
		case "resourceGroup":
			// ARM templates are often executed in the context of a resource group and the resourcegroup()
			// function gives access to it. Here we try to inject similar semantics by adding a parameter
			// specifying the resource group name. Then making all template references refer to a variable
			// that holds the resolved resource group (determined through invoke)
			rgNameParam, rg, err := t.ensureGlobals()
			if err != nil {
				return err
			}
			d.RefersTo(rgNameParam)
			d.RefersTo(rg)
			*target = model.VariableReference(&model.Variable{
				Name:         rg.Name(),
				VariableType: model.DynamicType,
			})
			return nil
		case "subscription":
			// Invoke the TF azure provider's "azure:core/getSubscription:getSubscription" to get subscription?
			/* TODO invoke for subscription? */

			return nyiError("subscription")

		case "deployment":
			return nyiError("deployment")

		case "variables":
			if len(f.ArgumentExpressions) != 1 {
				return fmt.Errorf("expect exactly one argument to 'variables' function, got %d", len(f.ArgumentExpressions))
			}
			value := f.ArgumentExpressions[0]
			evaled, err := t.evalValue(d, value, exclusions)
			if err != nil {
				return err
			}
			evaledVariableName, ok := evaled.(string)
			if !ok {
				return fmt.Errorf("expect variable to be a string, received: %T", evaled)
			}
			var varName string
			if !strings.HasSuffix(evaledVariableName, "Var") {
				varName = evaledVariableName + "Var"
			}
			varElement := t.lookup(varName)
			if varElement == nil {
				return fmt.Errorf("missing variable: %s", varName)
			}
			d.RefersTo(varElement)
			*target = model.VariableReference(&model.Variable{
				Name:         varElement.Name(),
				VariableType: model.DynamicType,
			})
			return nil
		case "parameters":
			if len(f.ArgumentExpressions) != 1 {
				return fmt.Errorf("expect exactly one argument to 'parameters' function, got %d", len(f.ArgumentExpressions))
			}
			value := f.ArgumentExpressions[0]
			evaled, err := t.evalValue(d, value, exclusions)
			if err != nil {
				return err
			}
			evaledParamName, ok := evaled.(string)
			if !ok {
				return fmt.Errorf("expect parameter to be a string, received: %T", evaledParamName)
			}
			if !strings.HasSuffix(evaledParamName, "Param") {
				evaledParamName = evaledParamName + "Param"
			}
			paramElement := t.lookup(evaledParamName)
			if paramElement == nil {
				return fmt.Errorf("missing parameter: %s", evaledParamName)
			}

			d.RefersTo(paramElement)
			*target = model.VariableReference(&model.Variable{
				Name:         paramElement.Name(),
				VariableType: model.DynamicType,
			})
			return nil
		case "resourceId":
			expression, err := tle.ToTemplateExpressionString(f)
			if err != nil {
				return err
			}
			expression = unquote(expression)

			if len(f.ArgumentExpressions) == 0 {
				return nyiError("resourceId") // empty resourceId?
			}
			arg, err := t.evalValue(d, f.ArgumentExpressions[0], exclusions)
			if err != nil {
				return err
			}
			if refString, ok := arg.(string); ok {
				// If the first arg is a provider name then we are dealing with
				// resources initialized in this template. Otherwise, we might be
				// dealing with global resources which we don't support
				// TODO: How to detect providers? There may be third-party providers which
				// don't use the "Microsoft.XXXX" format.
				if !strings.HasPrefix(strings.ToLower(refString), "microsoft.") {
					return nyiError("resourceId")
				}
			} else {
				// need resource type to resolve to a string
				return nyiError("resourceId")
			}

			var evaledResourceName interface{}
			if len(f.ArgumentExpressions) == 2 {
				evaledResourceName, err = t.evalValue(d, f.ArgumentExpressions[1], exclusions)
				if err != nil {
					return err
				}
			} else {
				// Multiple args likely mean the name had a concatenation
				resourceNameConcatFunc := tle.FunctionCallValue{
					NameToken:           tle.CreateLiteral(0, "concat"),
					ArgumentExpressions: []tle.Value{},
				}
				for _, a := range f.ArgumentExpressions[1:] {
					if len(resourceNameConcatFunc.ArgumentExpressions) > 0 {
						resourceNameConcatFunc.ArgumentExpressions = append(resourceNameConcatFunc.ArgumentExpressions, tle.NewStringValue(tle.CreateQuotedString(0, "/")))
					}
					resourceNameConcatFunc.ArgumentExpressions = append(resourceNameConcatFunc.ArgumentExpressions, a)
					evaledResourceName, err = t.evalValue(d, &resourceNameConcatFunc, exclusions)
					if err != nil {
						return err
					}
				}
			}

			res, err := t.lookupResourceByEvaledName(evaledResourceName)
			if err != nil {
				return err
			}

			if res == nil {
				*target = expression
				return nil
			}

			t.resourceIdMap[expression] = resourceIdTargetEntry{
				variableExpression: model.VariableReference(&model.Variable{
					Name:         res.Name(),
					VariableType: model.DynamicType,
				}),
				targetElementName: res.Name(),
				RawName:           res.rawName(),
			}

			// The intent of using resourceId is different based on context
			// - e.g. in dependsOn resourceId should translate to reference to resource
			// while in other contexts resourceId should actually be the resource's Id.
			// We can't distinguish between the use cases here so store the canonical
			// expression instead. At a later step where we have the appropriate context
			// we will use the expression to lookup in resourceIdMap and perform the
			// necessary translation of the value stored there.
			*target = expression
			return nil
			// TODO: Add subscriptionid and resource group id to the concat list
		case "reference":
			var referencedResource interface{}
			var version string

			for _, a := range f.ArgumentExpressions {
				refParam, err := t.evalValue(d, a, exclusions)
				if err != nil {
					return err
				}
				if referencedResource == nil {
					referencedResource = refParam
				} else if _, ok := refParam.(string); ok && version == "" {
					version = refParam.(string)
				}
				// else {
				// ignore [full] since we always get the full resource
				//}
			}

			referencedResStr, ok := referencedResource.(string)
			if ok {
				if isResourceID(referencedResStr) {
					// TODO: Provider needs support for
					// https://tinyurl.com/yxqadjr9
					return nyiError("resourceId")
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
						return fmt.Errorf("invalid payload for resource %s: %T", k, args)
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
						d.RefersTo(v)
						found = true
						break
					}

				}
			}
			if !found {
				// TODO: Need the ability to get by resourceId
				// https://tinyurl.com/yxqadjr9
				return nyiError("resourceId")
			}
			*target = referencedResource
			return nil
		case "uniquestring":
			return nyiError("uniquestring")
		case "concat":
			var tuple []model.Expression
			for _, a := range f.ArgumentExpressions {
				refArg, err := t.evalValue(d, a, exclusions)
				if err != nil {
					return err
				}
				x, err := pcl.RenderValue(refArg)
				if err != nil {
					return err
				}
				tuple = append(tuple, x)
			}

			tupleCons := &model.TupleConsExpression{
				Expressions: tuple,
			}
			afterLiteral := false
			var parts []model.Expression
			for _, x := range tupleCons.Expressions {
				lit, ok := getStringValue(x)
				if ok {
					if afterLiteral {
						last := parts[len(parts)-1].(*model.LiteralValueExpression)
						last.Value = cty.StringVal(last.Value.AsString() + lit)
					} else {
						parts = append(parts, plainLit(lit))
						afterLiteral = true
					}
				} else {
					if afterLiteral {
						last := parts[len(parts)-1].(*model.LiteralValueExpression)
						last.Value = cty.StringVal(last.Value.AsString() + lit)
					} else {
						parts = append(parts, plainLit(""))
					}

					parts = append(parts, x)
					afterLiteral = false
				}
			}
			if !afterLiteral {
				parts = append(parts, plainLit(""))
			}
			*target = &model.TemplateExpression{Parts: parts}
		default:
			return nyiError(f.NameToken.StringValue())
		}

		return nil
	}
}

func (t *TemplateElements) lookupResourceByEvaledName(evaledResourceName interface{}) (*resource, error) {
	for k, v := range t.elements {
		switch typ := v.TemplateElement.(type) {
		case *resource:
			args := v.Args()
			argsMap, ok := args.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid payload for resource %s: %T", k, args)
			}
			name, ok := argsMap["name"]
			if !ok {
				continue
			}
			_, ok = argsMap["type"]
			if !ok {
				continue
			}
			if reflect.DeepEqual(name, evaledResourceName) {
				return typ, nil
			}
		}
	}
	return nil, nil
}

// getStringValue returns the literal string value of the given expression, if any.
func getStringValue(v model.Expression) (string, bool) {
	switch v := v.(type) {
	case *model.LiteralValueExpression:
		if v.Value.Type() == cty.String {
			return v.Value.AsString(), true
		}
	case *model.TemplateExpression:
		if len(v.Parts) == 1 {
			if lit, ok := v.Parts[0].(*model.LiteralValueExpression); ok {
				if lit.Value.Type() == cty.String {
					return lit.Value.AsString(), true
				}
			}
		}
	}
	return "", false
}

// plainLit returns an unquoted string literal expression.
func plainLit(v string) *model.LiteralValueExpression {
	return &model.LiteralValueExpression{Value: cty.StringVal(v)}
}

var isResourceID = regexp.MustCompile(
	`(?i)/subscriptions/(.+?)/resourcegroups/(.+?)/providers/Microsoft..*`).MatchString
