// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

// TODO: Remove this test once we've created new tests for all kinds of data covered here.
func TestLegacy(t *testing.T) {
	var body = map[string]interface{}{
		"name":        "MyResource",
		"x-threshold": 123,
		"structure": map[string]interface{}{
			"v1":     "value1",
			"v2":     2,
			"v3-odd": "odd-value",
			"props": map[string]interface{}{
				"v4-nested": true,
			},
		},
		"properties": map[string]interface{}{
			"p1": "prop1",
			"p2": "prop2",
			"more": map[string]interface{}{
				"items": []interface{}{
					map[string]interface{}{"aaa": "111", "ccc": map[string]interface{}{"bbb": "333"}},
					map[string]interface{}{"aaa": "222"},
				},
				"itemsMap": map[string]interface{}{
					"key1": map[string]interface{}{"aaa": "444", "ccc": map[string]interface{}{"bbb": "555"}},
					"key2": map[string]interface{}{"aaa": "666"},
				},
			},
		},
		"union": map[string]interface{}{
			"type": "BBB",
			"bb": map[string]interface{}{
				"b": "valueOfB",
			},
		},
		"tags": map[string]interface{}{
			"createdBy":   "admin",
			"application": "dashboard",
		},
		"trivia": map[string]interface{}{
			"emptyBool":   false,
			"emptyNumber": 0,
			"emptyString": "",
			"emptyArray":  []string{},
		},
	}
	pathValues := map[string]string{
		"subscriptionID":       "0282681f-7a9e-123b-40b2-96babd57a8a1",
		"resourceGroupName":    "rg-name",
		"NetworkInterfaceName": "nic-name",
	}

	types := map[string]map[string]resources.AzureAPIProperty{
		"azure-native:testing:Structure": {
			"v1": {},
			"v2": {},
			"v3-odd": {
				SdkName: "v3",
			},
			"v4-nested": {
				SdkName:    "v4",
				Containers: []string{"props"},
			},
			"v5": {
				Ref: "#/types/azure-native:testing:SubResource",
			},
		},
		"azure-native:testing:StructureResponse": {
			"v1": {},
			"v2": {},
			"v3-odd": {
				SdkName: "v3",
			},
			"v4-nested": {
				SdkName:    "v4",
				Containers: []string{"props"},
			},
			"v5ReadOnly": {},
		},
		"azure-native:testing:More": {
			"items": {
				Items: &resources.AzureAPIProperty{
					Ref: "#/types/azure-native:testing:MoreItem",
				},
			},
			"itemsMap": {
				Type: "object",
				AdditionalProperties: &resources.AzureAPIProperty{
					Ref: "#/types/azure-native:testing:MoreItem",
				},
			},
		},
		"azure-native:testing:MoreItem": {
			"aaa": {
				SdkName: "Aaa",
			},
			"bbb": {
				Containers: []string{"ccc"},
			},
		},
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
		"azure-native:testing:SubResource": {
			"id": {
				Type: "string",
			},
		},
	}

	actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
		pathParameters: []resources.AzureAPIParameter{
			{
				Location: "path",
				Name:     "subscriptionId",
			},
			{
				Location: "path",
				Name:     "resourceGroupName",
			},
			{
				Location: "path",
				Name:     "NetworkInterfaceName",
				Value: &resources.AzureAPIProperty{
					SdkName: "networkInterfaceName",
				},
			},
		},
		bodyParameters: map[string]resources.AzureAPIProperty{
			"name": {},
			"x-threshold": {
				SdkName: "threshold",
			},
			"structure": {
				Ref: "#/types/azure-native:testing:Structure",
			},
			"p1": {
				Containers: []string{"properties"},
			},
			"p2": {
				Containers: []string{"properties"},
			},
			"p3": {
				Containers: []string{"properties", "document", "body"},
			},
			"more": {
				Containers: []string{"properties"},
				Ref:        "#/types/azure-native:testing:More",
			},
			"union": {
				OneOf: []string{"#/types/azure-native:testing:OptionA", "#/types/azure-native:testing:OptionB"},
			},
			"tags":         {},
			"untypedArray": {},
			"untypedDict": {
				Ref: TypeAny,
			},
		},
		pathValues: pathValues,
		types:      types,
		body:       body,
	})

	var expected = map[string]interface{}{
		"resourceGroupName":    "rg-name",
		"networkInterfaceName": "nic-name",
		"name":                 "MyResource",
		"threshold":            123,
		"structure": map[string]interface{}{
			"v1": "value1",
			"v2": 2,
			"v3": "odd-value",
			"v4": true,
		},
		"p1": "prop1",
		"p2": "prop2",
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"Aaa": "111", "bbb": "333"},
				map[string]interface{}{"Aaa": "222"},
			},
			"itemsMap": map[string]interface{}{
				"key1": map[string]interface{}{"Aaa": "444", "bbb": "555"},
				"key2": map[string]interface{}{"Aaa": "666"},
			},
		},
		"union": map[string]interface{}{
			"type": "BBB",
			"b":    "valueOfB",
		},
		"tags": map[string]interface{}{
			"createdBy":   "admin",
			"application": "dashboard",
		},
	}

	assert.Equal(t, expected, actual)
}

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

	t.Run("subscription", func(t *testing.T) {
		actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
			pathParameters: []resources.AzureAPIParameter{
				pathParam("subscriptionId"),
			},
			pathValues: map[string]string{
				"subscriptionID": "0282681f-7a9e-123b-40b2-96babd57a8a1",
			},
		})

		// SubscriptionId is not included in the inputs
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
}

func TestBodyProps(t *testing.T) {
	t.Run("untyped simple value", rapid.MakeCheck(func(t *rapid.T) {
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

	t.Run("containered", func(t *testing.T) {
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

func TestReferencedTypes(t *testing.T) {
	actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
		types: map[string]map[string]resources.AzureAPIProperty{
			"azure-native:testing:Structure": {
				"x-renamed": {
					SdkName: "renamed",
				},
				"containered": {
					Containers: []string{"props"},
				},
				"subResource": {
					Ref: "#/types/azure-native:testing:SubResource",
				},
			},
			"azure-native:testing:SubResource": {
				"subField": {
					Type: "string",
				},
			},
		},
		bodyParameters: map[string]resources.AzureAPIProperty{
			"structure": {
				Ref: "#/types/azure-native:testing:Structure",
			},
		},
		body: map[string]interface{}{
			"structure": map[string]interface{}{
				"x-renamed": "renamed-value",
				"props": map[string]interface{}{
					"containered": true,
				},
				"subResource": map[string]interface{}{
					"subField": "sub-value",
				},
			},
		},
	})

	var expected = map[string]interface{}{
		"structure": map[string]interface{}{
			"renamed":     "renamed-value",
			"containered": true,
			"subResource": map[string]interface{}{
				"subField": "sub-value",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestSubTypeRename(t *testing.T) {
	actual := testResponseToSdkInputs(responseToSdkInputsTestCase{
		types: map[string]map[string]resources.AzureAPIProperty{
			"azure-native:testing:Structure": {
				"x-renamed": {
					SdkName: "renamed",
				},
			},
		},
		bodyParameters: map[string]resources.AzureAPIProperty{
			"structure": {
				Ref: "#/types/azure-native:testing:Structure",
			},
		},
		body: map[string]interface{}{
			"structure": map[string]interface{}{
				"x-renamed": "renamed-value",
			},
		},
	})

	var expected = map[string]interface{}{
		"structure": map[string]interface{}{
			"renamed": "renamed-value",
		},
	}

	assert.Equal(t, expected, actual)
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
