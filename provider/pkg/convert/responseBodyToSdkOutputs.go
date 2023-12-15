package convert

import (
	"reflect"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// ResponseBodyToSdkOutputs converts a JSON request- or response body to the SDK shape.
func (k *SdkShapeConverter) ResponseBodyToSdkOutputs(props map[string]resources.AzureAPIProperty,
	response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		values := response
		for _, containerName := range prop.Containers {
			if v, has := values[containerName]; has {
				if v, ok := v.(map[string]interface{}); ok {
					values = v
				}
			} else {
				break
			}
		}

		if value, has := values[name]; has {
			if prop.Const != nil && value != prop.Const {
				// Returning nil indicates that the given value is not of the expected type.
				return nil
			}

			if value != nil {
				result[sdkName] = k.convertBodyPropToSdkPropValue(&p, value)
			}
		}
	}
	return result
}

// convertBodyPropToSdkPropValue converts a value from a request body to an SDK property value.
func (k *SdkShapeConverter) convertBodyPropToSdkPropValue(prop *resources.AzureAPIProperty, value interface{}) interface{} {
	return k.convertTypedResponseBodyObjectsToSdkOutputs(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		return k.ResponseBodyToSdkOutputs(props, values)
	})
}

// convertTypedResponseBodyObjectsToSdkOutputs recursively finds map types with a known type and calls convertMap on them.
func (k *SdkShapeConverter) convertTypedResponseBodyObjectsToSdkOutputs(prop *resources.AzureAPIProperty, value interface{}, convertObject convertTypedObject) interface{} {
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
				result[key] = k.convertTypedResponseBodyObjectsToSdkOutputs(prop.AdditionalProperties, item, convertObject)
			}
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
			result = append(result, k.convertTypedResponseBodyObjectsToSdkOutputs(prop.Items, s.Index(i).Interface(), convertObject))
		}
		return result
	}
	return value
}
