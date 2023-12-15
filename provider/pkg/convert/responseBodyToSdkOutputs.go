package convert

import (
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
	return k.convertTypedObjects(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		return k.ResponseBodyToSdkOutputs(props, values)
	})
}
