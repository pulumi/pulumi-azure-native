// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestPreviewOutputs(t *testing.T) {
	type previewTestCase struct {
		inputs   map[string]interface{}
		metadata map[string]resources.AzureAPIProperty
		types    map[string]resources.AzureAPIType
	}
	preview := func(testCase previewTestCase) resource.PropertyMap {
		inputMap := resource.NewPropertyMapFromMap(testCase.inputs)
		converter := NewSdkShapeConverterFull(testCase.types)
		return converter.PreviewOutputs(inputMap, testCase.metadata)
	}

	t.Run("sdk renamed property", func(t *testing.T) {
		actual := preview(previewTestCase{
			inputs: map[string]interface{}{
				"prop1": "value",
			},
			metadata: map[string]resources.AzureAPIProperty{
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
		metadata := map[string]resources.AzureAPIProperty{
			"type": {
				Const: "AAA",
			},
		}
		expected := resource.PropertyMap{
			"type": resource.NewStringProperty("AAA"),
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"type": "AAA",
				},
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"type": 123,
				},
				metadata: metadata,
			})
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("string property", func(t *testing.T) {
		metadata := map[string]resources.AzureAPIProperty{
			"name": {
				Type: "string",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"name": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
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
		metadata := map[string]resources.AzureAPIProperty{
			"threshold": {
				Type: "number",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"threshold": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"threshold": "123",
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"threshold": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("bool property", func(t *testing.T) {
		metadata := map[string]resources.AzureAPIProperty{
			"enabled": {
				Type: "boolean",
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"enabled": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"enabled": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"enabled": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("union property", func(t *testing.T) {
		metadata := map[string]resources.AzureAPIProperty{
			"union": {
				OneOf: []string{"#/types/azure-native:testing:OptionA", "#/types/azure-native:testing:OptionB"},
			},
		}
		types := map[string]resources.AzureAPIType{
			"azure-native:testing:OptionA": {
				Properties: map[string]resources.AzureAPIProperty{
					"type": {
						Const: "AAA",
					},
					"a": {
						Containers: []string{"aa"},
					},
				},
			},
			"azure-native:testing:OptionB": {
				Properties: map[string]resources.AzureAPIProperty{
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
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"union": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
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
		metadata := map[string]resources.AzureAPIProperty{
			"array": {
				Type: "array",
				Items: &resources.AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"array": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"array": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(previewTestCase{
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
		metadata := map[string]resources.AzureAPIProperty{
			"object": {
				Type: "object",
				AdditionalProperties: &resources.AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"object": 123,
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				"object": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(previewTestCase{
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
		metadata := map[string]resources.AzureAPIProperty{
			"complex": {
				Ref: "#/types/azure-native:testing:StructureResponse",
			},
		}
		types := map[string]resources.AzureAPIType{
			"azure-native:testing:StructureResponse": {
				Properties: map[string]resources.AzureAPIProperty{
					"v1": {
						Type: "string",
					},
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
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
			actual := preview(previewTestCase{
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"complex": 123,
				},
				metadata: metadata,
				types:    types,
			})
			expected := resource.PropertyMap{
				"complex": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(previewTestCase{
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
		metadata := map[string]resources.AzureAPIProperty{
			"userAssignedIdentities": {
				Type: "object",
				AdditionalProperties: &resources.AzureAPIProperty{
					Type: "string",
				},
			},
		}
		t.Run("valid", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"userAssignedIdentities": []interface{}{
						"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name",
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("unknown", func(t *testing.T) {
			actual := preview(previewTestCase{
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch", func(t *testing.T) {
			actual := preview(previewTestCase{
				inputs: map[string]interface{}{
					"userAssignedIdentities": []interface{}{
						123,
					},
				},
				metadata: metadata,
			})
			expected := resource.PropertyMap{
				// Whole object just marked as computed
				"userAssignedIdentities": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
		t.Run("mismatch element", func(t *testing.T) {
			actual := preview(previewTestCase{
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
				"userAssignedIdentities": resource.MakeComputed(resource.NewStringProperty("")),
			}
			assert.Equal(t, expected, actual)
		})
	})
}
