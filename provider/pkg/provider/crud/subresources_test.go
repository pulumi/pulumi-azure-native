package crud

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
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
		missingProperties := []propertyPath{
			{"remote"},
		}
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

	t.Run("extra property", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"existing": "value",
			},
		}
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
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
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
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
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
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
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

	// Regression test for #3036 - do not add empty containers to the body that will not be filled
	t.Run("issue-3036", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "privateNetworks"},
		}
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

	// Regression test for #4094 - do not add empty container in update case
	t.Run("issue-4094", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "securityRules"},
		}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{
			"properties": map[string]interface{}{
				"securityRules": []interface{}{},
			},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			"properties": map[string]interface{}{
				"securityRules": []interface{}{},
			},
		}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("properties container missing in body", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
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
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("intermediate containers", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "remote"},
		}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{
			"properties": map[string]interface{}{},
		}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{}
		assert.Equal(t, expected, bodyParams)
	})

	t.Run("Nested array missing from body", func(t *testing.T) {
		missingProperties := []propertyPath{
			{"properties", "accessPolicies"},
		}
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

	var dummyLookupType resources.TypeLookupFunc = func(ref string) (*resources.AzureAPIType, bool, error) {
		return nil, false, nil
	}

	t.Run("empty", func(t *testing.T) {
		res := &resources.AzureAPIResource{}
		oldInputs := map[string]any{}
		actual := findUnsetPropertiesToMaintain(res, oldInputs, true, dummyLookupType)
		expected := []propertyPath{}
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource not set", func(t *testing.T) {
		oldInputs := map[string]any{
			"existing": "value",
		}
		actual := findUnsetPropertiesToMaintain(resWithSubResource, oldInputs, true, dummyLookupType)
		expected := []propertyPath{
			{"subResource"},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource set", func(t *testing.T) {
		oldInputs := map[string]any{
			"existing":    "value",
			"subResource": "value",
		}
		actual := findUnsetPropertiesToMaintain(resWithSubResource, oldInputs, true, dummyLookupType)
		expected := []propertyPath{}
		assert.Equal(t, expected, actual)
	})
}

func TestFindUnsetSubResourcePropertiesFollowingTypeRefs(t *testing.T) {
	res, lookupType := setUpResourceWithRefAndProviderWithTypeLookup()

	t.Run("properties is not set", func(t *testing.T) {
		bodyParams := map[string]interface{}{}
		unset := findUnsetPropertiesToMaintain(res, bodyParams, false, lookupType)
		require.Equal(t, 1, len(unset))
		assert.Equal(t, []string{"properties", "accessPolicies"}, unset[0])
	})

	t.Run("KV accessPolicies is not set", func(t *testing.T) {
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{},
		}
		unset := findUnsetPropertiesToMaintain(res, bodyParams, false, lookupType)
		require.Equal(t, 1, len(unset))
		assert.Equal(t, []string{"properties", "accessPolicies"}, unset[0])
	})

	t.Run("KV accessPolicies is set", func(t *testing.T) {
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"accessPolicies": []interface{}{},
			},
		}
		unset := findUnsetPropertiesToMaintain(res, bodyParams, false, lookupType)
		assert.Empty(t, unset)
	})
}

// Helper to avoid repeating the same setup code in multiple tests. Returns a resource with a
// "properties" property of type azure-native:keyvault:VaultProperties, which the returned provider
// will return when asked to look up that type.
func setUpResourceWithRefAndProviderWithTypeLookup() (*resources.AzureAPIResource, resources.TypeLookupFunc) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"properties": {
							Type:       "object",
							Ref:        "#/types/azure-native:keyvault:VaultProperties",
							Containers: []string{"container"},
						},
					},
				},
			},
		},
	}

	// Mock the type lookup to only return the type referenced in the resource above
	lookupType := func(ref string) (*resources.AzureAPIType, bool, error) {
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
	}

	return &res, lookupType
}
