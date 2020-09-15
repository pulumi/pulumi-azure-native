// Copyright 2016-2020, Pulumi Corporation.

package provider

import "strings"

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
