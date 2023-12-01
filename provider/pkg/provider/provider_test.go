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

func TestSpecialInputs(t *testing.T) {
	inputs := resource.PropertyMap{
		"unrelated": resource.NewStringProperty("foo"),
	}
	olds := resource.PropertyMap{
		"unrelated":      resource.NewStringProperty("foo"),
		"networkRuleSet": resource.NewObjectProperty(resource.PropertyMap{}),
	}

	err := adjustInputsForSpecialCases(inputs,
		"azure-native:storage:StorageAccount",
		func() (resource.PropertyMap, error) {
			return olds, nil
		})
	assert.NoError(t, err)

	// Was not in inputs but was added to reset it back to default.
	assert.Contains(t, inputs, resource.PropertyKey("networkRuleSet"))
}
