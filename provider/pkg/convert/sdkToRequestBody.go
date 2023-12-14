// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// convertSdkPropToRequestBodyPropValue converts an SDK property to a value to be used in a request body.
func (k *SdkShapeConverter) convertSdkPropToRequestBodyPropValue(id string, prop *resources.AzureAPIProperty, value interface{}) interface{} {
	return k.convertTypedObjects(prop, value, func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{} {
		// Detect if we are dealing with a special case of a SubResource type with an ID property.
		// These properties reference a sub-ID of the currently modified resource (e.g.
		// an ID of a backend pool in a load balancer while creating the load balancer).
		// In that case, we allow users to specify a relative ID ($self/backendPool/abc) instead of
		// specifying the full Azure resource ID explicitly.
		// The block below takes care of resolving those relative IDs to absolute IDs.
		typeNameParts := strings.Split(typeName, ":")
		if len(typeNameParts) == 3 && typeNameParts[2] == "SubResource" {
			if relId, ok := values["id"].(string); ok && strings.HasPrefix(relId, "$self/") {
				values["id"] = strings.Replace(relId, "$self", id, 1)
			}
		}

		// Otherwise, delegate to the normal map convertion flow.
		return k.SdkPropertiesToRequestBody(props, values, id)
	})
}

// SdkPropertiesToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
func (k *SdkShapeConverter) SdkPropertiesToRequestBody(props map[string]resources.AzureAPIProperty,
	values map[string]interface{}, id string) map[string]interface{} {
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
			container[name] = k.convertSdkPropToRequestBodyPropValue(id, &p, value)
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
