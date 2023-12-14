// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestSdkOutputsToSdkInputs(t *testing.T) {
	t.Run("untyped non-empty values remain unchanged", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"untyped": {},
			},
			outputs: map[string]interface{}{
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
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"untyped": {
					Ref: TypeAny,
				},
			},
			outputs: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"untyped": value,
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("renamed", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"x-threshold": {
					SdkName: "threshold",
				},
			},
			outputs: map[string]interface{}{
				"threshold": 123,
			},
		})

		var expected = map[string]interface{}{
			"threshold": 123,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatched const returns nil", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"const": {
					Const: "value",
				},
			},
			outputs: map[string]interface{}{
				"const": "other",
			},
		})

		var expected map[string]interface{} = nil

		assert.Equal(t, expected, actual)
	})

	t.Run("array of empties not changed", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"emptyArray": {
					Type: "array",
				},
			},
			outputs: map[string]interface{}{
				"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("map of empties unchanged", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"emptyDict": {
					Type: "object",
				},
			},
			outputs: map[string]interface{}{
				"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed array doesn't change items", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"typedArray": {
					Type: "array",
					Items: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			outputs: map[string]interface{}{
				"typedArray": []interface{}{"a", "b", 3},
			},
		})

		var expected = map[string]interface{}{
			"typedArray": []interface{}{"a", "b", 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed map doesn't change items", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"typedMap": {
					Type: "object",
					AdditionalProperties: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			outputs: map[string]interface{}{
				"typedMap": map[string]interface{}{"a": "b", "c": 3},
			},
		})

		var expected = map[string]interface{}{
			"typedMap": map[string]interface{}{"a": "b", "c": 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("string set mapped back to string list", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					Type:        "object",
					IsStringSet: true,
				},
			},
			outputs: map[string]interface{}{
				"userAssignedIdentities": map[string]interface{}{
					"a": "b",
					"c": map[string]interface{}{
						"d": "e",
					},
				},
			},
		})

		var expected = map[string]interface{}{
			"userAssignedIdentities": []string{"a", "c"},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestSdkOutputsToSdkInputsNestedTypes(t *testing.T) {
	bodyParams := map[string]resources.AzureAPIProperty{
		"nested": {
			Ref: "#/types/azure-native:testing:SubType",
		},
	}
	t.Run("untyped simple value", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"value": {},
				},
			},
			bodyParameters: bodyParams,
			outputs: map[string]interface{}{
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

	t.Run("empty object unchanged", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"name": {},
				},
			},
			bodyParameters: bodyParams,
			outputs: map[string]interface{}{
				"nested": map[string]interface{}{},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("sub-id not ignored", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"id": {},
				},
			},
			bodyParameters: bodyParams,
			outputs: map[string]interface{}{
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
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"x-renamed": {
						SdkName: "renamed",
					},
				},
			},
			bodyParameters: bodyParams,
			outputs: map[string]interface{}{
				"nested": map[string]interface{}{
					"renamed": "value",
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

	t.Run("mismatched const ignored", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:SubType": {
					"const": {
						Const: "value",
					},
				},
			},
			bodyParameters: bodyParams,
			outputs: map[string]interface{}{
				"nested": map[string]interface{}{
					"const": "other",
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}(nil),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("string set mapped back to string list", func(t *testing.T) {
		actual := testSdkOutputsToSDKInputs(sdkOutputsToSDKInputsTestCase{
			bodyParameters: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					Type:        "object",
					IsStringSet: true,
				},
			},
			outputs: map[string]interface{}{
				"nested": map[string]interface{}{
					"userAssignedIdentities": map[string]interface{}{
						"a": "b",
						"c": map[string]interface{}{
							"d": "e",
						},
					},
				},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"userAssignedIdentities": []string{"a", "c"},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

type sdkOutputsToSDKInputsTestCase struct {
	outputs         map[string]interface{}
	bodyParameters  map[string]resources.AzureAPIProperty
	extraParameters []resources.AzureAPIParameter
	types           map[string]map[string]resources.AzureAPIProperty
}

func testSdkOutputsToSDKInputs(testCase sdkOutputsToSDKInputsTestCase) map[string]interface{} {
	types := map[string]resources.AzureAPIType{}
	if testCase.types != nil {
		for typeName, typeProperties := range testCase.types {
			types[typeName] = resources.AzureAPIType{
				Properties: typeProperties,
			}
		}
	}
	c := SdkShapeConverter{Types: types}
	parameters := testCase.extraParameters
	if testCase.bodyParameters != nil {
		parameters = append(parameters, resources.AzureAPIParameter{
			Location: body,
			Body: &resources.AzureAPIType{
				Properties: testCase.bodyParameters,
			},
		})
	}
	return c.SDKOutputsToSDKInputs(parameters, testCase.outputs)
}
