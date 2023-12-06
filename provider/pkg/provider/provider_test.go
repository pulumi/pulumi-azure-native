package provider

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestWritePropertiesToBody(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		missingProperties := map[string]resources.AzureAPIProperty{}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{}

		writePropertiesToBody(missingProperties, bodyParams, response)
		assert.Equal(t, map[string]interface{}{}, bodyParams)
	})

	t.Run("top-level", func(t *testing.T) {
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type: "string",
			},
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

	t.Run("properties container", func(t *testing.T) {
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
		}
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
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
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
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
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
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
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

	t.Run("properties container missing in body", func(t *testing.T) {
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
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
		missingProperties := map[string]resources.AzureAPIProperty{
			"remote": {
				Type:       "string",
				Containers: []string{"properties"},
			},
		}
		bodyParams := map[string]interface{}{}
		response := map[string]interface{}{}
		writePropertiesToBody(missingProperties, bodyParams, response)
		expected := map[string]interface{}{
			// Container is auto-created
			"properties": map[string]interface{}{},
		}
		assert.Equal(t, expected, bodyParams)
	})
}

func TestFindSubResourcePropertiesToMaintain(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		res := &resources.AzureAPIResource{}
		bodyParams := map[string]interface{}{}
		actual := findSubResourcePropertiesToMaintain(res, bodyParams)
		expected := map[string]resources.AzureAPIProperty{}
		assert.Equal(t, expected, actual)
	})

	t.Run("single prop", func(t *testing.T) {
		prop := resources.AzureAPIProperty{
			Type:                       "string",
			MaintainSubResourceIfUnset: true,
		}
		res := &resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"prop": prop,
						},
					},
				},
			},
		}
		bodyParams := map[string]interface{}{}
		actual := findSubResourcePropertiesToMaintain(res, bodyParams)
		expected := map[string]resources.AzureAPIProperty{
			"prop": prop,
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("containered missing prop", func(t *testing.T) {
		prop := resources.AzureAPIProperty{
			Type:                       "string",
			MaintainSubResourceIfUnset: true,
			Containers:                 []string{"properties"},
		}
		res := &resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"prop": prop,
						},
					},
				},
			},
		}
		bodyParams := map[string]interface{}{}
		actual := findSubResourcePropertiesToMaintain(res, bodyParams)
		expected := map[string]resources.AzureAPIProperty{
			"prop": prop,
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("containered existing prop", func(t *testing.T) {
		prop := resources.AzureAPIProperty{
			Type:                       "string",
			MaintainSubResourceIfUnset: true,
			Containers:                 []string{"properties"},
		}
		res := &resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"prop": prop,
						},
					},
				},
			},
		}
		bodyParams := map[string]interface{}{
			"properties": map[string]interface{}{
				"prop": "value",
			},
		}
		actual := findSubResourcePropertiesToMaintain(res, bodyParams)
		expected := map[string]resources.AzureAPIProperty{}
		assert.Equal(t, expected, actual)
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
	t.Run("empty", func(t *testing.T) {
		res := &resources.AzureAPIResource{}
		oldInputs := resource.PropertyMap{}
		actual := findUnsetSubResourceProperties(res, oldInputs)
		var expected []string
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource not set", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"existing": resource.NewStringProperty("value"),
		}
		actual := findUnsetSubResourceProperties(resWithSubResource, oldInputs)
		expected := []string{"subResource"}
		assert.Equal(t, expected, actual)
	})

	t.Run("sub-resource set", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"existing":    resource.NewStringProperty("value"),
			"subResource": resource.NewStringProperty("value"),
		}
		actual := findUnsetSubResourceProperties(resWithSubResource, oldInputs)
		var expected []string
		assert.Equal(t, expected, actual)
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
