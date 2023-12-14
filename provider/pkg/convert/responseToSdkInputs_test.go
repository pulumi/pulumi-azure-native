// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestPathParams(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			pathParameters: []resources.AzureAPIParameter{
				pathParam("resourceGroupName"),
			},
			pathValues: map[string]string{
				"resourceGroupName": "rg-name",
			},
		})

		// SubscriptionId is not included in the inputs
		var expected = map[string]interface{}{
			"resourceGroupName": "rg-name",
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("subscription is removed", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			pathParameters: []resources.AzureAPIParameter{
				pathParam("subscriptionId"),
			},
			pathValues: map[string]string{
				"subscriptionID": "0282681f-7a9e-123b-40b2-96babd57a8a1",
			},
		})

		var expected = map[string]interface{}{}

		assert.Equal(t, expected, actual)
	})

	t.Run("renamed", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			pathParameters: []resources.AzureAPIParameter{
				{
					Name:     "x-threshold",
					Location: "path",
					Value: &resources.AzureAPIProperty{
						SdkName: "threshold",
					},
				},
			},
			pathValues: map[string]string{
				"x-threshold": "123",
			},
		})

		var expected = map[string]interface{}{
			"threshold": "123",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestBodyProps(t *testing.T) {
	t.Run("untyped non-empty values remain unchanged", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"untyped": {},
			},
			body: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"untyped": value,
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("any type values", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"untyped": {
					Ref: TypeAny,
				},
			},
			body: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"untyped": value,
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("id-ignored", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"id": {},
			},
			body: map[string]interface{}{
				"id": "123",
			},
		})

		// Top-level SubscriptionId is not included in the inputs
		var expected = map[string]interface{}{}

		assert.Equal(t, expected, actual)
	})

	t.Run("renamed", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"x-threshold": {
					SdkName: "threshold",
				},
			},
			body: map[string]interface{}{
				"x-threshold": 123,
			},
		})

		var expected = map[string]interface{}{
			"threshold": 123,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("containers", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"container"},
				},
			},
			body: map[string]interface{}{
				"container": map[string]interface{}{
					"prop": "value",
				},
			},
		})

		var expected = map[string]interface{}{
			"prop": "value",
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("nested containers", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"a", "b", "c"},
				},
			},
			body: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": map[string]interface{}{
							"prop": "value",
						},
					},
				},
			},
		})

		var expected = map[string]interface{}{
			"prop": "value",
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatched const ignored", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"const": {
					Const: "value",
				},
			},
			body: map[string]interface{}{
				"const": "other",
			},
		})

		var expected = map[string]interface{}{}

		assert.Equal(t, expected, actual)
	})

	t.Run("array of empties replaced with nil", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"emptyArray": {
					Type: "array",
				},
			},
			body: map[string]interface{}{
				"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyArray": nil,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("map of empties replaced with nil", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"emptyDict": {
					Type: "object",
				},
			},
			body: map[string]interface{}{
				"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyDict": nil,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed array doesn't change items", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"typedArray": {
					Type: "array",
					Items: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			body: map[string]interface{}{
				"typedArray": []interface{}{"a", "b", 3},
			},
		})

		var expected = map[string]interface{}{
			"typedArray": []interface{}{"a", "b", 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed map doesn't change items", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"typedMap": {
					Type: "object",
					AdditionalProperties: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			body: map[string]interface{}{
				"typedMap": map[string]interface{}{"a": "b", "c": 3},
			},
		})

		var expected = map[string]interface{}{
			"typedMap": map[string]interface{}{"a": "b", "c": 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("string set", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					Type:        "object",
					IsStringSet: true,
				},
			},
			body: map[string]interface{}{
				"userAssignedIdentities": map[string]interface{}{
					"a": "b",
					"c": map[string]interface{}{
						"d": "e",
					},
				},
			},
		})

		var expected = map[string]interface{}{
			"userAssignedIdentities": []interface{}{"a", "c"},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestNestedTypes(t *testing.T) {
	bodyParams := map[string]resources.AzureAPIProperty{
		"nested": {
			Ref: "#/types/azure-native:testing:SubType",
		},
	}
	t.Run("untyped simple value", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"value": {},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{
					"value": value,
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"value": value,
			},
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("empty object replaced with nil", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"name": {},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{},
			},
		})

		var expected = map[string]interface{}{
			"nested": nil,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("sub-id not ignored", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"id": {},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{
					"id": "id-value",
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"id": "id-value",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("renamed", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"x-renamed": {
						SdkName: "renamed",
					},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{
					"x-renamed": "value",
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"renamed": "value",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("containered", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"containered": {
						Containers: []string{"props"},
					},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{
					"props": map[string]interface{}{
						"containered": true,
					},
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"containered": true,
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatched const ignored", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"const": {
						Const: "value",
					},
				},
			},
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"nested": map[string]interface{}{
					"const": "other",
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUnionTypes(t *testing.T) {
	types := map[string]map[string]resources.AzureAPIProperty{
		"azure-native:testing:OptionA": {
			"type": {
				Const: "AAA",
			},
			"a": {
				Containers: []string{"aa"},
			},
		},
		"azure-native:testing:OptionB": {
			"type": {
				Const: "BBB",
			},
			"b": {
				Containers: []string{"bb"},
			},
		},
	}
	bodyParams := map[string]resources.AzureAPIProperty{
		"union": {
			OneOf: []string{"#/types/azure-native:testing:OptionA", "#/types/azure-native:testing:OptionB"},
		},
	}
	t.Run("neither", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types:          types,
			bodyParameters: bodyParams,
			body:           map[string]interface{}{},
		})
		expected := map[string]interface{}{}
		assert.Equal(t, expected, actual)
	})
	t.Run("option a", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types:          types,
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"union": map[string]interface{}{
					"type": "AAA",
					"aa": map[string]interface{}{
						"a": "value",
					},
				},
			},
		})
		expected := map[string]interface{}{
			"union": map[string]interface{}{
				"type": "AAA",
				"a":    "value",
			},
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("option b", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			types:          types,
			bodyParameters: bodyParams,
			body: map[string]interface{}{
				"union": map[string]interface{}{
					"type": "BBB",
					"bb": map[string]interface{}{
						"b": "value",
					},
				},
			},
		})
		expected := map[string]interface{}{
			"union": map[string]interface{}{
				"type": "BBB",
				"b":    "value",
			},
		}
		assert.Equal(t, expected, actual)
	})
}

type responseToSdkInputsTestCase struct {
	body           map[string]interface{}
	bodyParameters map[string]resources.AzureAPIProperty
	pathParameters []resources.AzureAPIParameter
	pathValues     map[string]string
	types          map[string]map[string]resources.AzureAPIProperty
}

func testResponseToSdkInputs(testCase responseToSdkInputsTestCase) map[string]interface{} {
	types := map[string]resources.AzureAPIType{}
	if testCase.types != nil {
		for typeName, typeProperties := range testCase.types {
			types[typeName] = resources.AzureAPIType{
				Properties: typeProperties,
			}
		}
	}
	if testCase.pathValues == nil {
		testCase.pathValues = map[string]string{}
	}
	c := SdkShapeConverter{Types: types}
	if testCase.bodyParameters == nil {
		testCase.bodyParameters = map[string]resources.AzureAPIProperty{}
	}
	parameters := []resources.AzureAPIParameter{
		bodyParam(testCase.bodyParameters),
	}
	parameters = append(parameters, testCase.pathParameters...)
	return c.ResponseToSdkInputs(parameters, testCase.pathValues, testCase.body)
}

func pathParam(name string) resources.AzureAPIParameter {
	return resources.AzureAPIParameter{
		Location: "path",
		Name:     name,
	}
}

func bodyParam(properties map[string]resources.AzureAPIProperty) resources.AzureAPIParameter {
	return resources.AzureAPIParameter{
		Location: "body",
		Body: &resources.AzureAPIType{
			Properties: properties,
		},
	}
}
