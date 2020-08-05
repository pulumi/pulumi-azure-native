package provider

import "strings"

func (k *azurermProvider) sdkPropertyToRequest(prop *AzureApiProperty, value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
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

func (k *azurermProvider) sdkPropertiesToRequest(props map[string]AzureApiProperty, values map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, has := values[sdkName]; has {
			if prop.Container == "" {
				result[name] = k.sdkPropertyToRequest(&prop, value)
			} else {
				// "Unflatten" the property into a container property.
				container := map[string]interface{}{}
				if v, ok := result[prop.Container]; ok {
					if v, ok := v.(map[string]interface{}); ok {
						container = v
					}
				}
				container[name] = k.sdkPropertyToRequest(&prop, value)
				result[prop.Container] = container
			}
		}
	}
	return result
}

func (k *azurermProvider) responsePropertyToSdk(prop *AzureApiProperty, value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
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

func (k *azurermProvider) responseToSdkOutputs(props map[string]AzureApiProperty, response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
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
			result[sdkName] = k.responsePropertyToSdk(&prop, value)
		}
	}
	return result
}
