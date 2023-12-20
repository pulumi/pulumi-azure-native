package provider

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWritePropertiesToBody(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		missingProperties := []propertyPath{}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{}

		writePropertiesToBody(missingProperties, bodyParams, response)
		assert.Equal(t, map[string]interface{}{}, bodyParams)
	})

	t.Run("top-level", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{},
		}}
		bodyParams := map[string]interface{}{
			"existing": "value",
		}
		response := map[string]interface{}{
			"remote": "foo",
		}

		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"remote":   "foo",
			"existing": "value",
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("properties container", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			}}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"remote": "foo",
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"remote":   "foo",
				"existing": "value",
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("existing properties are maintained", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "bar",
				"remote":   "foo",
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
				"remote":   "foo",
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("properties missed from remote", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "new-value",
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("properties container missing from remote", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
		response := map[string]interface{}{}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("properties container missing in body", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"remote": "value",
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"remote": "value",
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("empty with container", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "remote",
			path:         []string{"properties"},
		}}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			// Container is auto-created
			"properties": map[string]interface{}{},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("Nested array missing from body", func(t *testing.T) {
		missingProperties := []propertyPath{{
			propertyName: "accessPolicies",
			path:         []string{"properties", "accessPolicies"},
		}}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{},
		}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"accessPolicies": []interface{}{},
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			// Container is auto-created
			"properties": map[string]interface{}{
				"accessPolicies": []interface{}{},
			},
		}
		assert.Equal(t, expected, bodyParams)
	})
}

func TestFindUnsetSubResourceProperties(t *testing.T) {
	resWithSubResource := &resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"subResource": {
							Type:                       "string",
							MaintainSubResourceIfUnset: true,
						},
					},
				},
			},
		},
	}

	provider := azureNativeProvider{}

	t.Run("empty", func(t *testing.T) {
		res := &resources.AzureAPIResource{}
		oldInputs := map[string]any{}
		actual := provider.findUnsetPropertiesToMaintain(res, oldInputs)
		expected := []propertyPath{}
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource not set", func(t *testing.T) {
		oldInputs := map[string]any{
			"existing": "value",
		}
		actual := provider.findUnsetPropertiesToMaintain(resWithSubResource, oldInputs)
		expected := []propertyPath{{
			propertyName: "subResource",
			path:         []string{"subResource"},
		}}
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource set", func(t *testing.T) {
		oldInputs := map[string]any{
			"existing":    "value",
			"subResource": "value",
		}
		actual := provider.findUnsetPropertiesToMaintain(resWithSubResource, oldInputs)
		expected := []propertyPath{}
		assert.Equal(t, expected, actual)
	})
}

func TestFindUnsetSubResourcePropertiesFollowingTypeRefs(t *testing.T) {
	res, provider := setUpResourceWithRefAndProviderWithTypeLookup()

	t.Run("KV accessPolicies is not set", func(t *testing.T) {
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{},
		}
		unset := provider.findUnsetPropertiesToMaintain(res, bodyParams)
		require.Equal(t, 1, len(unset))
		assert.Equal(t, "accessPolicies", unset[0].propertyName)
		assert.Equal(t, []string{"properties", "accessPolicies"}, unset[0].path)
	})

	t.Run("KV accessPolicies is set", func(t *testing.T) {
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"accessPolicies": []interface{}{},
			},
		}
		unset := provider.findUnsetPropertiesToMaintain(res, bodyParams)
		assert.Empty(t, unset)
	})
}

func TestRestoreDefaultInputs(t *testing.T) {
	inputs := resource.PropertyMap{
		"unrelated": resource.NewStringProperty("foo"),
	}
	olds := resource.PropertyMap{
		"unrelated":      resource.NewStringProperty("foo"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)

	// Was not in inputs but was added to reset it back to default.
	assert.Contains(t, inputs, resource.PropertyKey("networkRuleSet"))
}

func TestDoNotRestoreDefaultInputsIfInputPresent(t *testing.T) {
	inputs := resource.PropertyMap{
		"unrelated": resource.NewStringProperty("bar"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
			"defaultAction": resource.NewStringProperty("Deny"),
		}),
	}
	olds := resource.PropertyMap{
		"unrelated":      resource.NewStringProperty("foo"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)

	assert.Contains(t, inputs, resource.PropertyKey("networkRuleSet"))
	// Input "deny" was not overwritten with default "allow"
	assert.Equal(t, "Deny", inputs["networkRuleSet"].ObjectValue()["defaultAction"].StringValue())
}

func TestRestoreDefaultInputsIsNoopWithoutDefaultProperties(t *testing.T) {
	inputs := resource.PropertyMap{}

	olds := resource.PropertyMap{
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	res := resources.AzureAPIResource{} // no defaults

	err := restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)
	assert.Empty(t, inputs)

	// same with empty defaults
	res.DefaultProperties = map[string]interface{}{}
	err = restoreDefaultInputsForRemovedProperties(inputs, res, olds)
	assert.NoError(t, err)
	assert.Empty(t, inputs)
}

func TestMappableOldStateIsNoopWithoutDefaults(t *testing.T) {
	res := resources.AzureAPIResource{} // no defaults
	m := mappableOldState(res, resource.PropertyMap{
		"foo": resource.NewStringProperty("bar"),
	})
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, m)
}

func TestMappableOldStatePreservesNonDefaults(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := mappableOldState(res, resource.PropertyMap{
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
			"defaultAction": resource.NewStringProperty("Deny"),
		}),
	})
	assert.Equal(t, "Deny", m["networkRuleSet"].(map[string]interface{})["defaultAction"])
}

func TestMappableOldStateRemovesDefaultsThatWereInputs(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := mappableOldState(res, resource.PropertyMap{
		"__inputs": resource.NewObjectProperty(resource.PropertyMap{
			"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
				"defaultAction": resource.NewStringProperty("Allow"),
			}),
		}),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
			"defaultAction": resource.NewStringProperty("Allow"),
		}),
	})
	assert.Contains(t, m, "__inputs")
	assert.NotContains(t, m, "networkRuleSet")
}

func TestMappableOldStatePreservesDefaultsThatWereNotInputs(t *testing.T) {
	res := resources.AzureAPIResource{
		DefaultProperties: map[string]interface{}{
			"networkRuleSet": map[string]interface{}{
				"defaultAction": "Allow",
			},
		},
	}
	m := mappableOldState(res, resource.PropertyMap{
		"__inputs": resource.NewObjectProperty(resource.PropertyMap{}),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{
			"defaultAction": resource.NewStringProperty("Allow"),
		}),
	})
	assert.Contains(t, m, "__inputs")
	assert.Contains(t, m, "networkRuleSet")
}

func TestDeleteFromMap(t *testing.T) {
	m := map[string]any{
		"a": "scalar",
		"b": map[string]any{
			"b.a": "scalar",
			"b.b": map[string]any{
				"b.b.a": map[string]any{},
			},
		},
	}

	deleted := deleteFromMap(m, []string{"a"})
	assert.True(t, deleted)
	assert.NotContains(t, m, "a")
	assert.Contains(t, m, "b")

	deleted = deleteFromMap(m, []string{"b", "b.b", "b.b.a"})
	assert.True(t, deleted)
	assert.Contains(t, m, "b")
	b := m["b"].(map[string]any)
	assert.Contains(t, b, "b.a")
	assert.Contains(t, b, "b.b")
	bb := b["b.b"].(map[string]any)
	assert.NotContains(t, bb, "b.b.a")

	assert.False(t, deleteFromMap(m, []string{"b", "notfound"}))
}

func TestRemoveUnsetSubResourceProperties(t *testing.T) {
	ctx := context.Background()

	res, provider := setUpResourceWithRefAndProviderWithTypeLookup()

	t.Run("empty", func(t *testing.T) {
		empty := &resources.AzureAPIResource{}
		oldInputs := resource.PropertyMap{}
		sdkResponse := map[string]any{}
		actual := provider.removeUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, empty)
		expected := map[string]any{}
		assert.Equal(t, expected, actual)
	})

	t.Run("remove", func(t *testing.T) {
		oldInputs := resource.PropertyMap{}
		sdkResponse := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}
		actual := provider.removeUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, res)
		expected := map[string]any{"properties": map[string]any{}}
		assert.Equal(t, expected, actual)
	})

	t.Run("preserve", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			resource.PropertyKey("properties"): resource.NewObjectProperty(resource.PropertyMap{
				resource.PropertyKey("accessPolicies"): resource.NewArrayProperty([]resource.PropertyValue{}),
			}),
		}
		sdkResponse := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}
		actual := provider.removeUnsetSubResourceProperties(ctx, "urn", sdkResponse, oldInputs, res)
		expected := sdkResponse
		assert.Equal(t, expected, actual)
	})
}

// Helper to avoid repeating the same setup code in multiple tests. Returns a resource with a
// "properties" property of type azure-native:keyvault:VaultProperties, which the returned provider
// will return when asked to look up that type.
func setUpResourceWithRefAndProviderWithTypeLookup() (*resources.AzureAPIResource, *azureNativeProvider) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"properties": {
							Type: "object",
							Ref:  "#/types/azure-native:keyvault:VaultProperties",
						},
					},
				},
			},
		},
	}

	provider := azureNativeProvider{
		// Mock the type lookup to only return the type referenced in the resource above
		lookupType: func(ref string) (*resources.AzureAPIType, bool, error) {
			if ref == "#/types/azure-native:keyvault:VaultProperties" {
				return &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"accessPolicies": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								Type: "object",
								Ref:  "#/types/azure-native:keyvault:AccessPolicyEntry",
							},
							MaintainSubResourceIfUnset: true,
						},
					},
				}, true, nil
			}
			if ref == "#/types/azure-native:keyvault:AccessPolicyEntry" {
				return &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"permissions": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								Type: "string",
							},
							Containers: []string{"container2", "container3"},
						}},
				}, true, nil
			}
			return nil, false, nil
		},
	}

	return &res, &provider
}

func TestSetUnsetSubresourcePropertiesToDefaults(t *testing.T) {
	res, provider := setUpResourceWithRefAndProviderWithTypeLookup()

	t.Run("unchanged", func(t *testing.T) {
		body := map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body)
		assert.Equal(t, map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}, body)
	})

	t.Run("simple missing", func(t *testing.T) {
		body := map[string]any{
			"properties": map[string]any{},
		}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body)
		assert.Equal(t, map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}, body)
	})

	t.Run("nested missing", func(t *testing.T) {
		body := map[string]any{}
		provider.setUnsetSubresourcePropertiesToDefaults(*res, body)
		assert.Equal(t, map[string]any{
			"properties": map[string]any{
				"accessPolicies": []any{},
			},
		}, body)
	})
}
