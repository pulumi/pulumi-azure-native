// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"reflect"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// Returning nil indicates that the given value is not of the expected type.
type convertTypedObject func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{}

// convertTypedObjects recursively finds map types with a known type and calls convertMap on them.
func (k *SdkShapeConverter) convertTypedObjects(prop *resources.AzureAPIProperty, value interface{}, convertObject convertTypedObject) interface{} {
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
				result[key] = k.convertTypedObjects(prop.AdditionalProperties, item, convertObject)
			}
			return result
		}
		// Convert can be called for either turning user inputs into a request body or for turning a response body
		// into SDK outputs. In the latter case, we need to turn string sets back into arrays.
		if prop.IsStringSet {
			result := make([]interface{}, 0)
			for key := range valueMap {
				result = append(result, key)
			}
			return result
		}
		return value
	case reflect.Slice, reflect.Array:
		if prop.IsStringSet {
			emptyValue := struct{}{}
			setResult := map[string]interface{}{}
			for _, setItem := range value.([]interface{}) {
				if reflect.TypeOf(setItem).Kind() != reflect.String {
					// This should have been handled by validation
					continue
				}
				setResult[setItem.(string)] = emptyValue
			}
			return setResult
		}
		if prop.Items == nil {
			return value
		}

		result := make([]interface{}, 0)
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			result = append(result, k.convertTypedObjects(prop.Items, s.Index(i).Interface(), convertObject))
		}
		return result
	}
	return value
}
