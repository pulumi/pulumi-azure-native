package convert

import (
	"reflect"
	"sort"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// SdkOutputsToSdkInputs converts resource outputs (not a response body, but already valid outputs) to corresponding
// resource inputs, excluding the read-only properties from the map.
func (k *SdkShapeConverter) SdkOutputsToSdkInputs(parameters []resources.AzureAPIParameter, outputs map[string]interface{}) map[string]interface{} {
	for _, param := range parameters {
		if param.Location == body {
			return k.sdkOutputsToSDKInputs(param.Body.Properties, outputs)
		}
	}
	return nil
}

func (k *SdkShapeConverter) sdkOutputsToSDKInputs(props map[string]resources.AzureAPIProperty, outputs map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, ok := outputs[sdkName]; ok {
			if prop.Const != nil && value != prop.Const {
				return nil
			}

			p := prop // https://go.dev/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
			result[sdkName] = k.convertOutputToInputPropValue(&p, value)
		}
	}
	return result
}

// convertOutputToInputPropValue converts an output value back to an input value.
func (k *SdkShapeConverter) convertOutputToInputPropValue(prop *resources.AzureAPIProperty, value interface{}) interface{} {
	return k.convertTypedSdkOutputObjectsToSdkInput(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		return k.sdkOutputsToSDKInputs(props, values)
	})
}

// convertTypedSdkOutputObjectsToSdkInput recursively finds map types with a known type and calls convertMap on them.
func (k *SdkShapeConverter) convertTypedSdkOutputObjectsToSdkInput(prop *resources.AzureAPIProperty, value interface{}, convertObject convertTypedObject) interface{} {
	// This line is unreachable in the current implementation, but good to have for safety to prevent Kind() of a nil type from causing a nil pointer dereference.
	if value == nil {
		return nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		// For union types, iterate through types and find the first one that matches the shape.
		for _, t := range prop.OneOf {
			typeName := strings.TrimPrefix(t, "#/types/")
			typ, ok, err := k.GetType(typeName)
			if !ok || err != nil {
				continue
			}

			request := convertObject(typeName, typ.Properties, value.(map[string]interface{}))
			if request != nil {
				return request
			}
		}

		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}

		if strings.HasPrefix(prop.Ref, "#/types/") {
			typeName := strings.TrimPrefix(prop.Ref, "#/types/")
			typ, ok, err := k.GetType(typeName)
			if !ok || err != nil {
				return value
			}
			return convertObject(typeName, typ.Properties, valueMap)
		}

		if prop.AdditionalProperties != nil {
			result := map[string]interface{}{}
			for key, item := range valueMap {
				result[key] = k.convertTypedSdkOutputObjectsToSdkInput(prop.AdditionalProperties, item, convertObject)
			}
			return result
		}
		if prop.IsStringSet {
			result := make([]interface{}, 0)
			for key := range valueMap {
				result = append(result, key)
			}
			sort.SliceStable(result, func(i, j int) bool {
				return result[i].(string) < result[j].(string)
			})
			return result
		}
		return value
	case reflect.Slice, reflect.Array:
		if prop.Items == nil {
			return value
		}

		result := make([]interface{}, 0)
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			result = append(result, k.convertTypedSdkOutputObjectsToSdkInput(prop.Items, s.Index(i).Interface(), convertObject))
		}
		return result
	}
	return value
}
