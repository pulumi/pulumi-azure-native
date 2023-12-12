package convert

import "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

// SDKOutputsToSDKInputs converts resource outputs (not a response body, but already valid outputs) to corresponding
// resource inputs, excluding the read-only properties from the map.
func (k *SdkShapeConverter) SDKOutputsToSDKInputs(parameters []resources.AzureAPIParameter, outputs map[string]interface{}) map[string]interface{} {
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

			p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
			result[sdkName] = k.convertOutputToInputPropValue(&p, value)
		}
	}
	return result
}

// convertOutputToInputPropValue converts an output value back to an input value.
func (k *SdkShapeConverter) convertOutputToInputPropValue(prop *resources.AzureAPIProperty, value interface{}) interface{} {
	return k.convertPropValue(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		return k.sdkOutputsToSDKInputs(props, values)
	})
}
