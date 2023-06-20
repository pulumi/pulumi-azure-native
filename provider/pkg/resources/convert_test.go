// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

var resourceMap = &AzureAPIMetadata{
	Types: map[string]AzureAPIType{
		"azure-native:testing:Structure": {
			Properties: map[string]AzureAPIProperty{
				"v1": {},
				"v2": {},
				"v3-odd": {
					SdkName: "v3",
				},
				"v4-nested": {
					SdkName:    "v4",
					Containers: []string{"props"},
				},
			},
		},
		"azure-native:testing:StructureResponse": {
			Properties: map[string]AzureAPIProperty{
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
		},
		"azure-native:testing:More": {
			Properties: map[string]AzureAPIProperty{
				"items": {
					Items: &AzureAPIProperty{
						Ref: "#/types/azure-native:testing:MoreItem",
					},
				},
				"itemsMap": {
					Type: "object",
					AdditionalProperties: &AzureAPIProperty{
						Ref: "#/types/azure-native:testing:MoreItem",
					},
				},
			},
		},
		"azure-native:testing:MoreItem": {
			Properties: map[string]AzureAPIProperty{
				"aaa": {
					SdkName: "Aaa",
				},
				"bbb": {
					Containers: []string{"ccc"},
				},
			},
		},
		"azure-native:testing:OptionA": {
			Properties: map[string]AzureAPIProperty{
				"type": {
					Const: "AAA",
				},
				"a": {
					Containers: []string{"aa"},
				},
			},
		},
		"azure-native:testing:OptionB": {
			Properties: map[string]AzureAPIProperty{
				"type": {
					Const: "BBB",
				},
				"b": {
					Containers: []string{"bb"},
				},
			},
		},
	},
	Resources: map[string]AzureAPIResource{
		"r1": {
			PutParameters: []AzureAPIParameter{
				{
					Location: "body",
					Body: &AzureAPIType{
						Properties: map[string]AzureAPIProperty{
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
					},
				},
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
					Value: &AzureAPIProperty{
						SdkName: "networkInterfaceName",
					},
				},
			},
			Response: map[string]AzureAPIProperty{
				"name": {},
				"x-threshold": {
					SdkName: "threshold",
				},
				"structure": {
					Ref: "#/types/azure-native:testing:StructureResponse",
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
				"readOnly": {},
			},
		},
	},
}

var c = SdkShapeConverter{Types: resourceMap.Types}

var sampleAPIPackage = map[string]interface{}{
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
		"document": map[string]interface{}{
			"body": map[string]interface{}{
				"p3": "prop3",
			},
		},
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
	"untypedArray": []interface{}{
		map[string]interface{}{"key1": "value1"},
		map[string]interface{}{"key1": "value2"},
	},
	"untypedDict": map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	},
}
var sampleSdkProps = map[string]interface{}{
	"name":      "MyResource",
	"threshold": 123,
	"structure": map[string]interface{}{
		"v1": "value1",
		"v2": 2,
		"v3": "odd-value",
		"v4": true,
	},
	"p1": "prop1",
	"p2": "prop2",
	"p3": "prop3",
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
	"untypedArray": []interface{}{
		map[string]interface{}{"key1": "value1"},
		map[string]interface{}{"key1": "value2"},
	},
	"untypedDict": map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	},
}

func TestResponseToSdkOutputs(t *testing.T) {
	outputs := c.BodyPropertiesToSDK(resourceMap.Resources["r1"].Response, sampleAPIPackage)
	assert.Equal(t, sampleSdkProps, outputs)
}

func TestSdkPropertiesToRequestBody(t *testing.T) {
	bodyProperties := resourceMap.Resources["r1"].PutParameters[0].Body.Properties
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps)
	assert.Equal(t, sampleAPIPackage, data)
}

func TestSdkPropertiesToRequestBodyEmptyCollections(t *testing.T) {
	var emptyCollectionData = map[string]interface{}{
		"more": map[string]interface{}{
			"items":    make([]interface{}, 0),
			"itemsMap": make(map[string]interface{}),
		},
	}
	var expectedBody = map[string]interface{}{
		"properties": map[string]interface{}{
			"more": map[string]interface{}{
				"items":    make([]interface{}, 0),
				"itemsMap": make(map[string]interface{}),
			},
		},
	}
	bodyProperties := resourceMap.Resources["r1"].PutParameters[0].Body.Properties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, emptyCollectionData)
	assert.Equal(t, expectedBody, actualBody)
}

func TestSdkPropertiesToRequestBodyNilCollections(t *testing.T) {
	var nilCollectionData = map[string]interface{}{
		"more": map[string]interface{}{
			"items":    nil,
			"itemsMap": nil,
		},
	}
	var expectedBody = map[string]interface{}{
		"properties": map[string]interface{}{
			"more": map[string]interface{}{
				"items":    nil,
				"itemsMap": nil,
			},
		},
	}
	bodyProperties := resourceMap.Resources["r1"].PutParameters[0].Body.Properties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, nilCollectionData)
	assert.Equal(t, expectedBody, actualBody)
}

type resourceIDCase struct {
	id   string
	path string
}

func TestParseInvalidResourceID(t *testing.T) {
	cases := []resourceIDCase{
		// ID shorter than Path
		{"/resourceGroup/myrg", "/resourceGroup/{resourceGroup}/subResource"},
		// ID longer than Path
		{"/resourceGroup/myrg/cdn/mycdn", "/resourceGroup/{resourceGroup}/cdn"},
		// Segment names don't match
		{"/resourceGroup/myrg/foo/mycdn", "/resourceGroup/{resourceGroup}/bar/{cdn}"},
	}
	for _, testCase := range cases {
		_, err := ParseResourceID(testCase.id, testCase.path)
		assert.Error(t, err)
	}
}

func TestParseFullResourceID(t *testing.T) {
	id := "/subscriptions/0282681f-7a9e-123b-40b2-96babd57a8a1/resourcegroups/pulumi-name/providers/Microsoft.Network/networkInterfaces/pulumi-nic/ipConfigurations/ipconfig1"
	path := "/subscriptions/{subscriptionID}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	actual, err := ParseResourceID(id, path)
	assert.NoError(t, err)
	expected := map[string]string{
		"subscriptionID":       "0282681f-7a9e-123b-40b2-96babd57a8a1",
		"resourceGroupName":    "pulumi-name",
		"networkInterfaceName": "pulumi-nic",
		"ipConfigurationName":  "ipconfig1",
	}
	assert.Equal(t, expected, actual)
}

func TestParseScopedResourceID(t *testing.T) {
	id := "/subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev/providers/Microsoft.Authorization/roleAssignments/2a88abc7-f599-0eba-a21f-a1817e597115"
	path := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	actual, err := ParseResourceID(id, path)
	assert.NoError(t, err)
	expected := map[string]string{
		"scope":              "subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev",
		"roleAssignmentName": "2a88abc7-f599-0eba-a21f-a1817e597115",
	}
	assert.Equal(t, expected, actual)
}

var responseForInputCalculation = map[string]interface{}{
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
var calculatedInputs = map[string]interface{}{
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

func TestResponseToSdkInputs(t *testing.T) {
	pathValues := map[string]string{
		"subscriptionID":       "0282681f-7a9e-123b-40b2-96babd57a8a1",
		"resourceGroupName":    "rg-name",
		"NetworkInterfaceName": "nic-name",
	}
	inputs := c.ResponseToSdkInputs(resourceMap.Resources["r1"].PutParameters, pathValues, responseForInputCalculation)
	assert.Equal(t, calculatedInputs, inputs)
}

func TestSDKOutputsToSDKInputs(t *testing.T) {
	// Produced by copying sampleSdkProps plus some extra properties.
	outputs := map[string]interface{}{
		"name":      "MyResource",
		"threshold": 123,
		"structure": map[string]interface{}{
			"v1":         "value1",
			"v2":         2,
			"v3":         "odd-value",
			"v4":         true,
			"v5ReadOnly": "calculated",
		},
		"p1": "prop1",
		"p2": "prop2",
		"p3": "prop3",
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
		"untypedArray": []interface{}{
			map[string]interface{}{"key1": "value1"},
			map[string]interface{}{"key1": "value2"},
		},
		"untypedDict": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
		"readOnly": 12345,
	}
	inputs := c.SDKOutputsToSDKInputs(resourceMap.Resources["r1"].PutParameters, outputs)
	assert.Equal(t, sampleSdkProps, inputs)
}

func TestPreviewOutputs(t *testing.T) {
	type testCase struct {
		inputs   map[string]interface{}
		metadata map[string]AzureAPIProperty
		types    map[string]AzureAPIType
	}
	preview := func(testCase testCase) resource.PropertyMap {
		inputMap := resource.NewPropertyMapFromMap(testCase.inputs)
		converter := SdkShapeConverter{Types: testCase.types}
		return converter.PreviewOutputs(inputMap, testCase.metadata)
	}

	t.Run("sdk renamed property", func(t *testing.T) {
		actual := preview(testCase{
			inputs: map[string]interface{}{
				"prop1": "value",
			},
			metadata: map[string]AzureAPIProperty{
				"x-renamed": {
					SdkName: "prop1",
				},
			},
		})
		expected := resource.PropertyMap{
			"prop1": resource.NewStringProperty("value"),
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("constant property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"type": {
				Const: "AAA",
			},
		}
		expected := resource.PropertyMap{
			"type": resource.NewStringProperty("AAA"),
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"type": "AAA",
				},
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"type": 123,
				},
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("string property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"name": {
				Type: "string",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"name": "MyResource",
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"name": resource.NewStringProperty("MyResource"),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"name": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"name": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"name": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("number property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"threshold": {
				Type: "number",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"threshold": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"threshold": resource.NewNumberProperty(123),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"threshold": resource.MakeComputed(resource.NewNumberProperty(0)),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"threshold": "123",
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"threshold": resource.MakeComputed(resource.NewNumberProperty(0)),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("bool property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"enabled": {
				Type: "boolean",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"enabled": true,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"enabled": resource.NewBoolProperty(true),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"enabled": resource.MakeComputed(resource.NewBoolProperty(false)),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"enabled": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"enabled": resource.MakeComputed(resource.NewBoolProperty(false)),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("union property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"union": {
				OneOf: []string{"#/types/azure-native:testing:OptionA", "#/types/azure-native:testing:OptionB"},
			},
		}
		types := map[string]AzureAPIType{
			"azure-native:testing:OptionA": {
				Properties: map[string]AzureAPIProperty{
					"type": {
						Const: "AAA",
					},
					"a": {
						Containers: []string{"aa"},
					},
				},
			},
			"azure-native:testing:OptionB": {
				Properties: map[string]AzureAPIProperty{
					"type": {
						Const: "BBB",
					},
					"b": {
						Containers: []string{"bb"},
					},
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"union": map[string]interface{}{
						"type": "AAA",
						"a":    "a",
					},
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"union": resource.NewObjectProperty(resource.PropertyMap{
					"type": resource.NewStringProperty("AAA"),
					"a":    resource.NewStringProperty("a"),
				}),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"union": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"union": 123,
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"union": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("string array property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"array": {
				Type: "array",
				Items: &AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"array": []interface{}{
						"hello",
						"world",
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("hello"),
					resource.NewStringProperty("world"),
				}),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.MakeComputed(resource.NewArrayProperty([]resource.PropertyValue{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"array": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.MakeComputed(resource.NewArrayProperty([]resource.PropertyValue{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"array": []interface{}{
						"hello",
						123,
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("hello"),
					resource.MakeComputed(resource.NewStringProperty("")),
				})}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("string map property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"object": {
				Type: "object",
				AdditionalProperties: &AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"object": map[string]interface{}{
						"key1": "value1",
						"key2": "value2",
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.NewObjectProperty(resource.PropertyMap{
					"key1": resource.NewStringProperty("value1"),
					"key2": resource.NewStringProperty("value2"),
				}),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"object": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"object": map[string]interface{}{
						"key1": "value1",
						"key2": 123,
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.NewObjectProperty(resource.PropertyMap{
					"key1": resource.NewStringProperty("value1"),
					"key2": resource.MakeComputed(resource.NewStringProperty("")),
				}),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("complex property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"complex": {
				Ref: "#/types/azure-native:testing:StructureResponse",
			},
		}
		types := map[string]AzureAPIType{
			"azure-native:testing:StructureResponse": {
				Properties: map[string]AzureAPIProperty{
					"v1": {
						Type: "string",
					},
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"complex": map[string]interface{}{
						"v1": "value1",
					},
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.NewObjectProperty(resource.PropertyMap{
					"v1": resource.NewStringProperty("value1"),
				}),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"complex": 123,
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"complex": map[string]interface{}{
						"v1": 123,
					},
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.NewObjectProperty(resource.PropertyMap{
					"v1": resource.MakeComputed(resource.NewStringProperty("")),
				}),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("string set property", func(t *testing.T) {
		metadata := map[string]AzureAPIProperty{
			"userAssignedIdentities": {
				Type: "object",
				AdditionalProperties: &AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"userAssignedIdentities": []interface{}{
						"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name",
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(testCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"userAssignedIdentities": []interface{}{
						123,
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(testCase{
				inputs: map[string]interface{}{
					"userAssignedIdentities": []interface{}{
						"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name",
						123,
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewObjectProperty(resource.PropertyMap{})),
			}
			assert.Equal(t, expected, actual)
		})
	})
}

func TestIsDefaultResponseTrue(t *testing.T) {
	response := map[string]interface{}{
		"p1":             true,
		"name":           "foo",
		"irrelevantBool": false,
		"untypedDict": map[string]interface{}{
			"flag":   true,
			"string": "bar",
		},
	}
	defaultBody := map[string]interface{}{
		"p1":   true,
		"name": "foo",
		"untypedDict": map[string]interface{}{
			"flag":   true,
			"string": "bar",
		},
	}
	actual := c.IsDefaultResponse(resourceMap.Resources["r1"].PutParameters, response, defaultBody)
	assert.True(t, actual)
}

func TestIsDefaultResponseFalse(t *testing.T) {
	defaultBody := map[string]interface{}{}
	response1 := map[string]interface{}{
		"p1": true,
	}
	actual1 := c.IsDefaultResponse(resourceMap.Resources["r1"].PutParameters, response1, defaultBody)
	assert.False(t, actual1)

	response2 := map[string]interface{}{
		"untypedArray": []string{"1", "2"},
	}
	actual2 := c.IsDefaultResponse(resourceMap.Resources["r1"].PutParameters, response2, defaultBody)
	assert.False(t, actual2)

	response3 := map[string]interface{}{
		"untypedDict": map[string]interface{}{
			"flag":   true,
			"string": "buzz",
		},
	}
	defaultBody = map[string]interface{}{
		"untypedDict": map[string]interface{}{
			"flag":   true,
			"string": "bar",
		},
	}
	actual3 := c.IsDefaultResponse(resourceMap.Resources["r1"].PutParameters, response3, defaultBody)
	assert.False(t, actual3)
}
