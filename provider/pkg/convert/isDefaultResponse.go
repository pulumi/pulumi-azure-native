package convert

import (
	"reflect"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// IsDefaultResponse returns true if the shape of the HTTP response matches the expected shape.
// The following comparison rules apply:
//   - response is converted to the SDK shape of inputs (so, the structure is flattened and read-only props are removed)
//   - A boolean 'false' in the response is equivalent to a no-value in the expected map
//   - Any non-empty map or slice leads to the 'false' result (may need to revise if any API endpoints have default
//     non-empty collections, but none are found yet)
//   - The response need not contain all properties that are present in the default body.
func (k *SdkShapeConverter) IsDefaultResponse(putParameters []resources.AzureAPIParameter, response map[string]interface{},
	defaultBody map[string]interface{}) bool {
	for _, param := range putParameters {
		if param.Location == body {
			for key, value := range k.ResponseBodyToSdkOutputs(param.Body.Properties, response) {
				defaultValue, _ := defaultBody[key]
				if !isDefaultValue(value, defaultValue) {
					return false
				}
			}
		}
	}
	return true
}

func isDefaultValue(value any, defaultValue any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		collection := reflect.ValueOf(value)
		// TODO: Add support for non-empty collections in the case that provider/pkg/openapi/defaults/defaultResourcesState has non-empty arrays added.
		if collection.Len() > 0 {
			return false
		}
	case reflect.Map:
		iter := reflect.ValueOf(value).MapRange()
		for iter.Next() {
			mk := iter.Key().String()
			mv := iter.Value().Interface()
			defaultMap, ok := defaultValue.(map[string]interface{})
			if !ok {
				return false
			}
			defaultValue, has := defaultMap[mk]
			if !has || !isDefaultValue(mv, defaultValue) {
				return false
			}
		}
	case reflect.Bool:
		b := value.(bool)
		if b && defaultValue != value {
			return false
		}
	default:
		// `*` default body means that we want to accept any value there.
		// It's used for values that are determined dynamically by Azure API.
		if defaultValue != value && defaultValue != "*" {
			return false
		}
	}
	return true
}
