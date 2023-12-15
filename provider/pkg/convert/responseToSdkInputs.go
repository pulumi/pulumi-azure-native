package convert

import (
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// ResponseToSdkInputs calculates a map of input values that would produce the given resource path and
// response. This is useful when we need to import an existing resource based on its current properties.
func (k *SdkShapeConverter) ResponseToSdkInputs(parameters []resources.AzureAPIParameter,
	pathValues map[string]string, responseBody map[string]interface{}) map[string]interface{} {
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
			outputs := k.ResponseBodyToSdkOutputs(param.Body.Properties, responseBody)
			inputs := k.SdkOutputsToSdkInputs([]resources.AzureAPIParameter{param}, outputs)
			for k, v := range inputs {
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
