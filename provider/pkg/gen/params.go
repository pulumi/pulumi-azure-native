package gen

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/debug"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
)

// FlattenParams takes the parameters specified in Azure API specs/ARM templates and
// flattens them to match the desired format for the Pulumi schema.
// resourceParams is a mapping of API parameter names to provider.AzureApiParameter
// and types is a mapping for the API type names to provider.AzureApiType.
// The latter two can be derived from the metadata generated during schema generation.
func FlattenParams(
	input map[string]interface{},
	resourceParams map[string]provider.AzureAPIParameter,
	types map[string]provider.AzureAPIType,
) (map[string]interface{}, error) {
	containers := map[string]codegen.StringSet{}
	for k, v := range resourceParams {
		if v.Value != nil && v.Value.Container != "" {
			if _, ok := containers[v.Value.Container]; !ok {
				containers[v.Value.Container] = codegen.NewStringSet()
			}
			containers[v.Value.Container].Add(k)
		}
	}

	out := map[string]interface{}{}
	for k, v := range input {
		switch k {
		case "If-Match": // TODO: Not handled in schema
			continue
		case "api-version", "apiVersion", "subscriptionId":
			continue // No need to emit these since we auto inject them
		}

		if _, ok := containers[k]; ok {
			contained, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("property %s is expected to be a map, recieved: %T", k, v)
			}
			flattened, err := FlattenParams(contained, resourceParams, types)
			if err != nil {
				return nil, err
			}

			mergeMap(out, flattened)
			continue
		}

		paramMetadata, ok := resourceParams[k]
		if !ok {
			debug.Log("missing item '%s' from resource metadata for resource, skipping", k)
			continue
		}
		// body parameters are folded in
		if paramMetadata.Body != nil {
			inBody, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expect body for param %s to be a map, recieved %T", k, v)
			}
			bodyVals, ok := inBody["properties"]
			if !ok {
				debug.Log("Missing properties in body param for '%s'. Assuming body already folded in.", k)
				bodyVals = inBody
			} else if len(inBody) > 1 {
				debug.Log("items in body outside of 'properties' field, will merge them into properties")
				for k, v := range inBody {
					if k != "properties" {
						bodyVals.(map[string]interface{})[k] = v
					}
				}
			}
			bodyValsMap := bodyVals.(map[string]interface{})
			bodyProps := map[string]provider.AzureAPIProperty{}
			for k, v := range paramMetadata.Body.Properties {
				props := v
				props.Container = ""
				bodyProps[k] = props
			}
			flattened := transformProperties(bodyProps, types, bodyValsMap)
			mergeMap(out, flattened)
			continue
		}

		// replace param name with value in SdkName if provided.
		if paramMetadata.Value != nil && paramMetadata.Value.SdkName != "" {
			k = paramMetadata.Value.SdkName
		}

		out[k] = v
	}
	return out, nil
}

func transformProperty(prop *provider.AzureAPIProperty, types map[string]provider.AzureAPIType, val interface{}) interface{} {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Map:
		if prop.Ref == "" {
			typeName := strings.TrimPrefix(prop.Type, "#/types/")
			if _, typed := types[typeName]; typed {
				return transformProperties(types[prop.Type].Properties, types, val.(map[string]interface{}))
			}
			// Return untyped dictionaries as-is if no container unwrapping required.
			return val
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := types[typeName]
		return transformProperties(typ.Properties, types, val.(map[string]interface{}))
	case reflect.Slice, reflect.Array:
		var result []interface{}
		s := reflect.ValueOf(val)
		for i := 0; i < s.Len(); i++ {
			if prop.Items != nil {
				result = append(result, transformProperty(prop.Items, types, s.Index(i).Interface()))
			} else {
				result = append(result, s.Index(i).Interface())
			}
		}
		return result
	}
	return val
}

func transformProperties(
	props map[string]provider.AzureAPIProperty,
	types map[string]provider.AzureAPIType,
	values map[string]interface{},
) map[string]interface{} {
	containers := codegen.NewStringSet()
	for _, v := range props {
		if v.Container != "" {
			containers.Add(v.Container)
		}
	}

	result := map[string]interface{}{}
	for k, v := range values {
		prop, ok := props[k]
		if !ok && !containers.Has(k) {
			debug.Log("missing '%s' in props: %#v", k, props)
			continue
		}
		if containers.Has(k) {
			if v != nil {
				container := transformProperty(&prop, types, v)
				// Expect container types to wrap maps.
				mergeMap(result, container.(map[string]interface{}))
			}
			continue
		}

		sdkName := k

		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		if v != nil {
			result[sdkName] = transformProperty(&prop, types, v)
		}
	}
	return result
}

func mergeMap(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
