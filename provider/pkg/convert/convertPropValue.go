// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"reflect"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

type convertPropValues func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{}

func (k *SdkShapeConverter) convertPropValue(prop *resources.AzureAPIProperty, value interface{}, convertMap convertPropValues) interface{} {
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

			request := convertMap(typeName, typ.Properties, value.(map[string]interface{}))
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
			return convertMap(typeName, typ.Properties, valueMap)
		}

		if prop.AdditionalProperties != nil {
			result := map[string]interface{}{}
			for key, item := range valueMap {
				result[key] = k.convertPropValue(prop.AdditionalProperties, item, convertMap)
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
			result = append(result, k.convertPropValue(prop.Items, s.Index(i).Interface(), convertMap))
		}
		return result
	}
	return value
}
