// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const body = "body"

const TypeAny = "pulumi.json#/Any"

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types        map[string]AzureAPIType
	PartialTypes PartialMap[AzureAPIType]
}

func NewSdkShapeConverterPartial(ptypes PartialMap[AzureAPIType]) SdkShapeConverter {
	return SdkShapeConverter{
		Types:        nil,
		PartialTypes: ptypes,
	}
}

func NewSdkShapeConverterFull(types map[string]AzureAPIType) SdkShapeConverter {
	return SdkShapeConverter{
		Types:        types,
		PartialTypes: NewPartialMap[AzureAPIType](),
	}
}

func (k *SdkShapeConverter) GetType(name string) (AzureAPIType, bool, error) {
	if k.Types != nil {
		typ, ok := k.Types[name]
		if ok {
			return typ, true, nil
		}
		return AzureAPIType{}, false, nil
	}

	return k.PartialTypes.Get(name)
}

type convertPropValues func(props map[string]AzureAPIProperty, values map[string]interface{}) map[string]interface{}

func (k *SdkShapeConverter) convertPropValue(prop *AzureAPIProperty, value interface{}, convertMap convertPropValues) interface{} {
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

			request := convertMap(typ.Properties, value.(map[string]interface{}))
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
			return convertMap(typ.Properties, valueMap)
		}

		if prop.AdditionalProperties != nil {
			result := map[string]interface{}{}
			for key, item := range valueMap {
				result[key] = k.convertPropValue(prop.AdditionalProperties, item, convertMap)
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
			result = append(result, k.convertPropValue(prop.Items, s.Index(i).Interface(), convertMap))
		}
		return result
	}
	return value
}

// SdkPropertiesToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
func (k *SdkShapeConverter) SdkPropertiesToRequestBody(props map[string]AzureAPIProperty,
	values map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, has := values[sdkName]; has {
			if prop.Const != nil && value != prop.Const {
				return nil
			}

			container := k.buildContainer(result, prop.Containers)
			container[name] = k.convertPropValue(&p, value, k.SdkPropertiesToRequestBody)
		}
	}
	return result
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

// BodyPropertiesToSDK converts a JSON request- or response body to the SDK shape.
func (k *SdkShapeConverter) BodyPropertiesToSDK(props map[string]AzureAPIProperty,
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
				return nil
			}

			if value != nil {
				result[sdkName] = k.convertPropValue(&p, value, k.BodyPropertiesToSDK)
			}
		}
	}
	return result
}

// ResponseToSdkInputs calculates a map of input values that would produce the given resource path and
// response. This is useful when we need to import an existing resource based on its current properties.
func (k *SdkShapeConverter) ResponseToSdkInputs(parameters []AzureAPIParameter,
	pathValues map[string]string, response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, param := range parameters {
		switch {
		case strings.EqualFold(param.Name, "subscriptionId"):
			// Ignore
		case param.Location == "path":
			name := param.Name
			sdkName := name
			if param.Value != nil && param.Value.SdkName != "" {
				sdkName = param.Value.SdkName
			}
			result[sdkName] = pathValues[name]
		case param.Location == body:
			bodyProps := k.BodyPropertiesToSDK(param.Body.Properties, response)
			for k, v := range bodyProps {
				switch {
				case k == "id":
					// Some resources have a top-level `id` property which is (probably incorrectly) marked as
					// non-readonly. `id` is a special property to Pulumi and will always cause diffs if set in
					// the result of a Read operation and block import. So, don't copy it to inputs.
					continue
				default:
					// Attempt to exclude insignificant properties from the inputs. A resource response would
					// contain a lot of default values, e.g. empty arrays when no values were specified, empty
					// strings, or false booleans. The decision to remove them is somewhat arbitrary but it
					// seems to make the practical import experience smoother.
					result[k] = removeEmptyCollections(v)
				}
			}
		}
	}
	return result
}

func (k *SdkShapeConverter) sdkOutputsToSDKInputs(props map[string]AzureAPIProperty, outputs map[string]interface{}) map[string]interface{} {
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

			p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
			result[sdkName] = k.convertPropValue(&p, value, k.sdkOutputsToSDKInputs)
		}
	}
	return result
}

// SDKOutputsToSDKInputs converts resource outputs (not a response body, but already valid outputs) to corresponding
// resource inputs, excluding the read-only properties from the map.
func (k *SdkShapeConverter) SDKOutputsToSDKInputs(parameters []AzureAPIParameter, outputs map[string]interface{}) map[string]interface{} {
	for _, param := range parameters {
		if param.Location == body {
			return k.sdkOutputsToSDKInputs(param.Body.Properties, outputs)
		}
	}
	return nil
}

// IsDefaultResponse returns true if the shape of the HTTP response matches the expected shape.
// The following comparison rules apply:
// - response is converted to the SDK shape of inputs (so, the structure is flattened and read-only props are removed)
// - A boolean 'false' in the response is equivalent to a no-value in the expected map
// - Any non-empty map or slice leads to the 'false' result (may need to revise if any API endpoints have default
//   non-empty collections, but none are found yet)
func (k *SdkShapeConverter) IsDefaultResponse(putParameters []AzureAPIParameter, response map[string]interface{},
	defaultBody map[string]interface{}) bool {
	for _, param := range putParameters {
		if param.Location == body {
			for key, value := range k.BodyPropertiesToSDK(param.Body.Properties, response) {
				switch reflect.TypeOf(value).Kind() {
				case reflect.Slice, reflect.Array:
					collection := reflect.ValueOf(value)
					if collection.Len() > 0 {
						return false
					}
				case reflect.Map:
					iter := reflect.ValueOf(value).MapRange()
					for iter.Next() {
						mk := iter.Key().String()
						mv := iter.Value().Interface()
						defaultMap, ok := defaultBody[key].(map[string]interface{})
						if !ok {
							return false
						}
						defaultValue, has := defaultMap[mk]
						if !has || defaultValue != mv {
							return false
						}
					}
				case reflect.Bool:
					b := value.(bool)
					if b && defaultBody[key] != value {
						return false
					}
				default:
					// `*` default body means that we want to accept any value there.
					// It's used for values that are determined dynamically by Azure API.
					if defaultBody[key] != value && defaultBody[key] != "*" {
						return false
					}
				}
			}
		}
	}
	return true
}

// PreviewOutputs calculates a map of outputs at the time of initial resource creation. It takes the provided resource
// inputs and maps them to the outputs shape, adding unknowns for all properties that are not defined in inputs.
func (k *SdkShapeConverter) PreviewOutputs(inputs resource.PropertyMap,
	props map[string]AzureAPIProperty) resource.PropertyMap {
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
			result[key] = resource.MakeComputed(resource.NewStringProperty(""))
		}
	}
	return result
}

func (k *SdkShapeConverter) previewOutputValue(inputValue resource.PropertyValue,
	prop *AzureAPIProperty) resource.PropertyValue {
	switch {
	case prop.Items != nil && inputValue.IsArray():
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
	default:
		return inputValue
	}
}

// ParseResourceID extracts templated values from the given resource ID based on the names of those templated
// values in an HTTP path. The structure of id and path must match: we validate it by building a regular
// expression based on the path parameters and matching the id.
func ParseResourceID(id, path string) (map[string]string, error) {
	pathParts := strings.Split(path, "/")
	regexParts := make([]string, len(pathParts))
	for i, s := range pathParts {
		if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
			name := s[1 : len(s)-1]
			regexParts[i] = fmt.Sprintf("(?P<%s>.+)", name)
		} else {
			regexParts[i] = pathParts[i]
		}
	}

	expr := fmt.Sprintf("(?i)^%s$", strings.Join(regexParts, "/"))
	pattern, err := regexp.Compile(expr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compile expression '%s' for path '%s'", expr, path)
	}

	match := pattern.FindStringSubmatch(id)
	if len(match) < len(pattern.SubexpNames()) {
		return nil, errors.Errorf("failed to parse '%s' against the path '%s'", id, path)
	}

	result := map[string]string{}
	for i, name := range pattern.SubexpNames() {
		if i > 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result, nil
}

// removeEmptyCollections returns nil if the given value is a default map or array with no values.
func removeEmptyCollections(value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		result := map[string]interface{}{}
		for k, v := range value {
			resultValue := removeEmptyCollections(v)
			if resultValue != nil {
				result[k] = resultValue
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, v := range value {
			resultValue := removeEmptyCollections(v)
			if resultValue != nil {
				result = append(result, resultValue)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	}
	return value
}
