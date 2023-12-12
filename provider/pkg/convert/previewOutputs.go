package convert

import (
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const TypeAny = "pulumi.json#/Any"

// PreviewOutputs calculates a map of outputs at the time of initial resource creation. It takes the provided resource
// inputs and maps them to the outputs shape, adding unknowns for all properties that are not defined in inputs.
func (k *SdkShapeConverter) PreviewOutputs(inputs resource.PropertyMap,
	props map[string]resources.AzureAPIProperty) resource.PropertyMap {
	result := resource.PropertyMap{}
	for name, prop := range props {
		p := prop
		if prop.SdkName != "" {
			name = prop.SdkName
		}
		key := resource.PropertyKey(name)
		if inputValue, ok := inputs[key]; ok {
			result[key] = k.previewOutputValue(inputValue, &p)
		} else {
			placeholderValue := resource.MakeComputed(resource.NewStringProperty(""))
			result[key] = k.previewOutputValue(placeholderValue, &p)
		}
	}
	return result
}

func (k *SdkShapeConverter) previewOutputValue(inputValue resource.PropertyValue,
	prop *resources.AzureAPIProperty) resource.PropertyValue {
	if prop == nil {
		return resource.MakeComputed(resource.NewStringProperty(""))
	}
	if prop.Const != nil {
		if asString, ok := prop.Const.(string); ok {
			return resource.NewStringProperty(asString)
		}
	}
	switch prop.Type {
	case "boolean":
		if inputValue.IsBool() {
			return inputValue
		}
	case "integer", "number":
		if inputValue.IsNumber() {
			return inputValue
		}
	case "string":
		if inputValue.IsString() {
			return inputValue
		}
	}
	switch {
	case prop.Type == "boolean" && inputValue.IsBool():
		return inputValue
	case prop.Type == "integer" && inputValue.IsNumber():
		return inputValue
	case prop.Type == "number" && inputValue.IsNumber():
		return inputValue
	case prop.Type == "string" && inputValue.IsString():
		return inputValue
	case (prop.Type == "array" || prop.Items != nil) && inputValue.IsArray():
		var items []resource.PropertyValue
		for _, item := range inputValue.ArrayValue() {
			items = append(items, k.previewOutputValue(item, prop.Items))
		}
		return resource.NewArrayProperty(items)

	case strings.HasPrefix(prop.Ref, "#/types/") && inputValue.IsObject():
		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ, _, _ := k.GetType(typeName)
		v := k.PreviewOutputs(inputValue.ObjectValue(), typ.Properties)
		return resource.NewObjectProperty(v)

	case prop.AdditionalProperties != nil && inputValue.IsObject():
		inputObject := inputValue.ObjectValue()
		result := resource.PropertyMap{}
		for name, value := range inputObject {
			p := value
			result[name] = k.previewOutputValue(p, prop.AdditionalProperties)
		}
		return resource.NewObjectProperty(result)

	case prop.OneOf != nil && inputValue.IsObject():
		// TODO: It would be nice to do something smart here and pick the right oneOf branch based on the input.
		// The challenge is differentiating between a mis-match and valid unknowns in the input.
		// inputObject := inputValue.ObjectValue()
		// for _, oneOf := range prop.OneOf {
		// 	typeName := strings.TrimPrefix(oneOf, "#/types/")
		// 	typ, _, _ := k.GetType(typeName)
		// 	v := k.PreviewOutputs(inputObject, typ.Properties)
		// 	v.Mappable()
		// 	if !v.ContainsUnknowns() {
		// 		return resource.NewObjectProperty(v)
		// 	}
		// }

		// Fallback to legacy behaviour - assuming the input is same as output
		return inputValue

	case prop.Ref == TypeAny:
		return inputValue

	case prop.Type == "" && prop.Items == nil && prop.AdditionalProperties == nil && prop.OneOf == nil && prop.Ref == "":
		// Untyped property
		if inputValue.IsString() || inputValue.IsNumber() || inputValue.IsBool() {
			// Assume simple types with no type information remain unchanged
			return inputValue
		}
	}
	return resource.MakeComputed(k.makeComputedValue(prop))
}

func (k *SdkShapeConverter) makeComputedValue(prop *resources.AzureAPIProperty) resource.PropertyValue {
	if prop != nil && prop.Const != nil {
		return resource.NewStringProperty(prop.Const.(string))
	}
	// To mark something as computed, we always use a string property with an
	// empty string, regardless of the type.
	return resource.NewStringProperty("")
}
