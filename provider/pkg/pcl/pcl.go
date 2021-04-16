// Copyright 2021, Pulumi Corporation.  All rights reserved.

package pcl

import (
	"reflect"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

// null represents PCL's builtin `null` variable
var null = &model.Variable{
	Name:         "null",
	VariableType: model.NoneType,
}

// RenderValue renders an object that represents a json value into PCL. Most nodes are rendered as one
// would expect (e.g. sequences -> tuple construction, maps -> object construction, etc.).
func RenderValue(node interface{}) (model.Expression, error) {
	if node == nil {
		return model.VariableReference(null), nil
	}

	typ := reflect.TypeOf(node)
	val := reflect.ValueOf(node)
	kind := typ.Kind()
	switch kind {
	case reflect.Slice:
		var expressions []model.Expression
		for i := 0; i < val.Len(); i++ {
			e, err := RenderValue(val.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, e)
		}
		return &model.TupleConsExpression{
			Expressions: expressions,
		}, nil
	case reflect.Bool:
		b := reflect.ValueOf(node).Bool()
		return &model.LiteralValueExpression{
			Value: cty.BoolVal(b),
		}, nil
	case reflect.Float32, reflect.Float64:
		return &model.LiteralValueExpression{
			Value: cty.NumberFloatVal(val.Float()),
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &model.LiteralValueExpression{Value: cty.NumberIntVal(val.Int())}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &model.LiteralValueExpression{Value: cty.NumberUIntVal(val.Uint())}, nil
	case reflect.String:
		return QuotedLit(val.String()), nil
	case reflect.Map:
		consItems := map[string]model.Expression{}
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			rendered, err := RenderValue(v.Interface())
			if err != nil {
				return nil, err
			}
			consItems[k.String()] = rendered
		}
		var items []model.ObjectConsItem
		for _, k := range codegen.SortedKeys(consItems) {
			items = append(items, ObjectConsItem(k, consItems[k]))
		}
		return &model.ObjectConsExpression{
			Items: items,
		}, nil
	case reflect.Ptr:
		nodeExpr, ok := node.(model.Expression)
		if !ok {
			// only expect model.Expression as the embedded interface
			panic(val)
		}
		return nodeExpr, nil
	default:
		contract.Failf("unexpected type %T", node)
		return nil, nil
	}
}
