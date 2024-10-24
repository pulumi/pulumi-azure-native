// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// SdkInputsToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
// Returns a map of request body properties and a map of unmapped values.
func (k *SdkShapeConverter) SdkInputsToRequestBody(props map[string]resources.AzureAPIProperty,
	values map[string]interface{}, id string) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	unusedValues := map[string]interface{}{}
	for name, value := range values {
		if _, ok := props[name]; !ok {
			unusedValues[name] = value
		}
	}

	for name, prop := range props {
		p := prop // https://go.dev/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, has := values[sdkName]; has {
			delete(unusedValues, sdkName)

			if prop.Const != nil && value != prop.Const {
				return nil, fmt.Errorf("property %s is a constant of value %q and cannot be modified to be %q", name, prop.Const, value)
			}

			container := k.buildContainer(result, prop.Containers)
			container[name] = k.convertSdkPropToRequestBodyPropValue(id, &p, value)
		}
	}

	if len(unusedValues) > 0 {
		unusedKeys := make([]string, 0, len(unusedValues))
		for k := range unusedValues {
			unusedKeys = append(unusedKeys, k)
		}
		sort.Strings(unusedKeys)
		logging.V(9).Infof("Unrecognized properties in SdkInputsToRequestBody: %v", unusedKeys)
	}
	return result, nil
}

// convertSdkPropToRequestBodyPropValue converts an SDK property to a value to be used in a request body.
func (k *SdkShapeConverter) convertSdkPropToRequestBodyPropValue(id string, prop *resources.AzureAPIProperty, value interface{}) interface{} {
	return k.convertTypedSdkInputObjectsToRequestBody(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		// Detect if we are dealing with a special case of a SubResource type with an ID property.
		// These properties reference a sub-ID of the currently modified resource (e.g.
		// an ID of a backend pool in a load balancer while creating the load balancer).
		// In that case, we allow users to specify a relative ID ($self/backendPool/abc) instead of
		// specifying the full Azure resource ID explicitly.
		// The block below takes care of resolving those relative IDs to absolute IDs.
		if _, _, resourceName, err := resources.ParseToken(typeName); err == nil && resourceName == "SubResource" {
			if relId, ok := values["id"].(string); ok && strings.HasPrefix(relId, "$self/") {
				values["id"] = strings.Replace(relId, "$self", id, 1)
			}
		}

		// Otherwise, delegate to the normal map conversion flow.
		converted, _ := k.SdkInputsToRequestBody(props, values, id)
		// We ignore the error here as we haven't previously handled these errors recursively through convertTypedSdkInputObjectsToRequestBody.
		// A difficulty is that the error is treated as a warning if we got a valid result, so should only be thrown if `converted` is nil.
		return converted
	})
}

// buildContainer creates a nested container for each item in 'path' and returns that inner-most container.
// For instance, a 'path' of ["top", "bottom"] would return a map, which is assigned to a key "bottom" in another
// map, which is assigned to a key "top" in the 'parent' map.
func (k *SdkShapeConverter) buildContainer(parent map[string]interface{}, path []string) map[string]interface{} {
	for _, containerName := range path {
		container := map[string]interface{}{}
		if v, ok := parent[containerName]; ok {
			if v, ok := v.(map[string]interface{}); ok {
				container = v
			}
		}
		parent[containerName] = container
		parent = container
	}
	return parent
}

// convertTypedSdkInputObjectsToRequestBody recursively finds map types with a known type and calls convertMap on them.
func (k *SdkShapeConverter) convertTypedSdkInputObjectsToRequestBody(prop *resources.AzureAPIProperty, value interface{}, convertObject convertTypedObject) interface{} {
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
				result[key] = k.convertTypedSdkInputObjectsToRequestBody(prop.AdditionalProperties, item, convertObject)
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
			result = append(result, k.convertTypedSdkInputObjectsToRequestBody(prop.Items, s.Index(i).Interface(), convertObject))
		}
		return result
	}
	return value
}
