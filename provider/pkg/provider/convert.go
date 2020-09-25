// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"strings"

	"github.com/pkg/errors"
)

func (k *azureNextGenProvider) sdkPropertyToRequest(prop *AzureAPIProperty, value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		// For union types, iterate through types and find the first one that matches the shape.
		for _, t := range prop.OneOf {
			typeName := strings.TrimPrefix(t, "#/types/")
			typ := k.resourceMap.Types[typeName]
			request := k.sdkPropertiesToRequest(typ.Properties, value)
			if request != nil {
				return request
			}
		}

		if prop.Ref == "" {
			// Return untyped dictionaries as-is.
			return value
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := k.resourceMap.Types[typeName]
		return k.sdkPropertiesToRequest(typ.Properties, value)
	case []interface{}:
		var result []interface{}
		for _, item := range value {
			result = append(result, k.sdkPropertyToRequest(prop.Items, item))
		}
		return result
	}
	return value
}

func (k *azureNextGenProvider) sdkPropertiesToRequest(props map[string]AzureAPIProperty,
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

			if prop.Container == "" {
				result[name] = k.sdkPropertyToRequest(&p, value)
			} else {
				// "Unflatten" the property into a container property.
				container := map[string]interface{}{}
				if v, ok := result[prop.Container]; ok {
					if v, ok := v.(map[string]interface{}); ok {
						container = v
					}
				}
				container[name] = k.sdkPropertyToRequest(&p, value)
				result[prop.Container] = container
			}
		}
	}
	return result
}

func (k *azureNextGenProvider) responsePropertyToSdk(prop *AzureAPIProperty, value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		// For union types, iterate through types and find the first one that matches the shape.
		for _, t := range prop.OneOf {
			typeName := strings.TrimPrefix(t, "#/types/")
			typ := k.resourceMap.Types[typeName]
			response := k.responseToSdkOutputs(typ.Properties, value)
			if response != nil {
				return response
			}
		}

		if prop.Ref == "" {
			// Return untyped dictionaries as-is.
			return value
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := k.resourceMap.Types[typeName]
		return k.responseToSdkOutputs(typ.Properties, value)
	case []interface{}:
		var result []interface{}
		for _, item := range value {
			result = append(result, k.responsePropertyToSdk(prop.Items, item))
		}
		return result
	}
	return value
}

func (k *azureNextGenProvider) responseToSdkOutputs(props map[string]AzureAPIProperty,
	response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		values := response
		if prop.Container != "" {
			if v, has := values[prop.Container]; has {
				if v, ok := v.(map[string]interface{}); ok {
					values = v
				}
			}
		}

		if value, has := values[name]; has {
			if prop.Const != nil && value != prop.Const {
				return nil
			}

			result[sdkName] = k.responsePropertyToSdk(&p, value)
		}
	}
	return result
}

func (k *azureNextGenProvider) parseResourceID(id, path string) (map[string]string, error) {
	pathParts := strings.Split(path, "/")
	idParts := strings.Split(id, "/")
	if len(pathParts) != len(idParts) {
		return nil, errors.Errorf("length of id doesn't match the path: '%s' vs. '%s'", id, path)
	}

	result := map[string]string{}
	for i, s := range pathParts {
		value := idParts[i]
		switch {
		case strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}"):
			name := s[1 : len(s)-1]
			result[name] = value
		case !strings.EqualFold(s, value):
			return nil, errors.Errorf("failed parsing id: %s <> %s", s, value)
		}
	}
	return result, nil
}

func (k *azureNextGenProvider) responseToSdkInputs(parameters []AzureAPIParameter,
	pathValues map[string]string, response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, param := range parameters {
		switch {
		case strings.EqualFold(param.Name, "subscriptionId"):
			// Ignore
		case param.Location == "path":
			name := param.Name
			result[name] = pathValues[name]
		case param.Location == body:
			bodyProps := k.responseToSdkOutputs(param.Body.Properties, response)
			for k, v := range bodyProps {
				switch {
				case k == "id":
					// Some resources have a top-level `id` property which is (probably incorrectly) marked as
					// non-readonly. `id` is a special property to Pulumi and will always cause diffs if set in
					// the result of a Read operation and block import. So, don't copy it to inputs.
					continue
				default:
					result[k] = removeTrivialValues(v)
				}
			}
		}
	}
	return result
}

func removeTrivialValues(value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		result := map[string]interface{}{}
		for k, v := range value {
			resultValue := removeTrivialValues(v)
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
			resultValue := removeTrivialValues(v)
			if resultValue != nil {
				result = append(result, resultValue)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case bool:
		if !value {
			return nil
		}
	case string:
		if len(value) == 0 {
			return nil
		}
	}
	return value
}
