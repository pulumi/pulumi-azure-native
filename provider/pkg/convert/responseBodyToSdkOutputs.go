package convert

import (
	"reflect"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// ResponseBodyToSdkOutputs converts a JSON request- or response body to the SDK shape.
func (k *SdkShapeConverter) ResponseBodyToSdkOutputs(props map[string]resources.AzureAPIProperty,
	response map[string]interface{},
) map[string]interface{} {
	result := map[string]interface{}{}

	for name, prop := range props {
		p := prop // https://go.dev/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
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
	return k.convertTypedResponseBodyObjectsToSdkOutputs(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		// Reverse of outbound AKS autoscaler profile workaround: convert kebab-case wire keys
		// back to camelCase expected by generated SDK property names.
		if strings.HasSuffix(typeName, "ManagedClusterPropertiesAutoScalerProfile") || strings.HasSuffix(typeName, "ManagedClusterPropertiesResponseAutoScalerProfile") {
			values = restoreAutoScalerProfileKeys(values)
		}
		return k.ResponseBodyToSdkOutputs(props, values)
	})
}

// restoreAutoScalerProfileKeys performs the inverse mapping of rewriteAutoScalerProfileKeys.
func restoreAutoScalerProfileKeys(in map[string]interface{}) map[string]interface{} {
	if in == nil {
		return in
	}
	mappings := map[string]string{
		"scan-interval":                         "scanInterval",
		"scale-down-delay-after-add":            "scaleDownDelayAfterAdd",
		"scale-down-delay-after-delete":         "scaleDownDelayAfterDelete",
		"scale-down-delay-after-failure":        "scaleDownDelayAfterFailure",
		"scale-down-unneeded-time":              "scaleDownUnneededTime",
		"scale-down-unready-time":               "scaleDownUnreadyTime",
		"ignore-daemonsets-utilization":         "ignoreDaemonsetsUtilization",
		"daemonset-eviction-for-empty-nodes":    "daemonsetEvictionForEmptyNodes",
		"daemonset-eviction-for-occupied-nodes": "daemonsetEvictionForOccupiedNodes",
		"scale-down-utilization-threshold":      "scaleDownUtilizationThreshold",
		"max-graceful-termination-sec":          "maxGracefulTerminationSec",
		"balance-similar-node-groups":           "balanceSimilarNodeGroups",
		"expander":                              "expander",
		"skip-nodes-with-local-storage":         "skipNodesWithLocalStorage",
		"skip-nodes-with-system-pods":           "skipNodesWithSystemPods",
		"max-empty-bulk-delete":                 "maxEmptyBulkDelete",
		"new-pod-scale-up-delay":                "newPodScaleUpDelay",
		"max-total-unready-percentage":          "maxTotalUnreadyPercentage",
		"max-node-provision-time":               "maxNodeProvisionTime",
		"ok-total-unready-count":                "okTotalUnreadyCount",
	}
	out := map[string]interface{}{}
	for k, v := range in {
		if sdk, ok := mappings[k]; ok {
			out[sdk] = v
		} else {
			out[k] = v
		}
	}
	return out
}

// convertTypedResponseBodyObjectsToSdkOutputs recursively finds map types with a known type and calls convertMap on them.
func (k *SdkShapeConverter) convertTypedResponseBodyObjectsToSdkOutputs(prop *resources.AzureAPIProperty, value interface{}, convertObject convertTypedObject) interface{} {
	// This line is unreachable in the current implementation, but good to have for safety to prevent Kind() of a nil type from causing a nil pointer dereference.
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
				result[key] = k.convertTypedResponseBodyObjectsToSdkOutputs(prop.AdditionalProperties, item, convertObject)
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
			result = append(result, k.convertTypedResponseBodyObjectsToSdkOutputs(prop.Items, s.Index(i).Interface(), convertObject))
		}
		return result
	}
	return value
}
