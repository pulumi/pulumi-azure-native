// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestSdkInputsToRequestBodySubResource(t *testing.T) {
	c := SdkShapeConverter{
		Types: map[string]resources.AzureAPIType{
			"azure-native:testing:Structure": {
				Properties: map[string]resources.AzureAPIProperty{
					"v5": {
						Ref: "#/types/azure-native:testing:SubResource",
					},
				},
			},
			"azure-native:testing:SubResource": {
				Properties: map[string]resources.AzureAPIProperty{
					"id": {
						Type: "string",
					},
				},
			},
		},
	}
	var sdkData = map[string]interface{}{
		"structure": map[string]interface{}{
			"v5": map[string]interface{}{
				"id": "$self/relative/id/123",
			},
		},
	}
	var expectedBody = map[string]interface{}{
		"structure": map[string]interface{}{
			"v5": map[string]interface{}{
				"id": "/sub/456/rg/my/network/abc/relative/id/123",
			},
		},
	}
	bodyProperties := map[string]resources.AzureAPIProperty{
		"structure": {
			Ref: "#/types/azure-native:testing:Structure",
		},
	}
	actualBody, err := c.SdkInputsToRequestBody(bodyProperties, sdkData, "/sub/456/rg/my/network/abc")
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, actualBody)
}

func captureStderr(f func()) string {
	// Create a pipe
	r, w, _ := os.Pipe()
	// Save the original stderr
	originalStderr := os.Stderr
	// Redirect stderr to the write end of the pipe
	os.Stderr = w

	// Run the target function
	f()

	// Close the write end of the pipe and restore stderr
	w.Close()
	os.Stderr = originalStderr

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	reader := bufio.NewReader(r)
	buf.ReadFrom(reader)

	return buf.String()
}

func TestSdkInputsToRequestBody(t *testing.T) {
	type testCaseArgs struct {
		id     string
		types  map[string]map[string]resources.AzureAPIProperty
		props  map[string]resources.AzureAPIProperty
		inputs map[string]interface{}
	}

	convertWithError := func(args testCaseArgs) (map[string]interface{}, error) {
		types := map[string]resources.AzureAPIType{}
		if args.types != nil {
			for typeName, typeProperties := range args.types {
				types[typeName] = resources.AzureAPIType{
					Properties: typeProperties,
				}
			}
		}
		c := SdkShapeConverter{
			Types: types,
		}
		return c.SdkInputsToRequestBody(args.props, args.inputs, args.id)
	}

	convert := func(args testCaseArgs) map[string]interface{} {
		body, err := convertWithError(args)
		assert.Nil(t, err)
		return body
	}

	t.Run("nil inputs", func(t *testing.T) {
		actual := convert(testCaseArgs{
			inputs: nil,
		})

		expected := map[string]interface{}{}

		assert.Equal(t, expected, actual)
	})

	t.Run("unmatched inputs are reported", func(t *testing.T) {
		prevLogToStderr := logging.LogToStderr
		prevVerbose := logging.Verbose
		prevLogFlow := logging.LogFlow
		logging.InitLogging(true, 9, true)
		defer func() {
			logging.InitLogging(prevLogToStderr, prevVerbose, prevLogFlow)
		}()

		var actual map[string]any
		stderr := captureStderr(func() {
			actual = convert(testCaseArgs{
				props: map[string]resources.AzureAPIProperty{
					"propA": {
						Type: "string",
					},
				},
				inputs: map[string]interface{}{
					"propA": "a",
					"propB": "b",
				},
			})
		})

		assert.Equal(t, map[string]interface{}{"propA": "a"}, actual)
		assert.Contains(t, stderr, "Unrecognized properties in SdkInputsToRequestBody: [propB]")
	})

	t.Run("untyped non-empty values remain unchanged", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"untyped": {},
			},
			inputs: map[string]interface{}{
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
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"untyped": {
					Ref: TypeAny,
				},
			},
			inputs: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"untyped": value,
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("renamed", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"x-threshold": {
					SdkName: "threshold",
				},
			},
			inputs: map[string]interface{}{
				"threshold": 123,
			},
		})

		var expected = map[string]interface{}{
			"x-threshold": 123,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("containers", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"container"},
				},
			},
			inputs: map[string]interface{}{
				"prop": "value",
			},
		})

		var expected = map[string]interface{}{
			"container": map[string]interface{}{
				"prop": "value",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("nested containers", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"a", "b", "c"},
				},
			},
			inputs: map[string]interface{}{
				"prop": "value",
			},
		})

		var expected = map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": map[string]interface{}{
						"prop": "value",
					},
				},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatched const returns nil", func(t *testing.T) {
		actual, err := convertWithError(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"const": {
					Const: "value",
				},
			},
			inputs: map[string]interface{}{
				"const": "other",
			},
		})

		assert.Nil(t, actual)
		assert.Equal(t, err, fmt.Errorf("property const is a constant of value \"value\" and cannot be modified to be \"other\""))
	})

	t.Run("array of empties not changed", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"emptyArray": {
					Type: "array",
				},
			},
			inputs: map[string]interface{}{
				"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("map of empties unchanged", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"emptyDict": {
					Type: "object",
				},
			},
			inputs: map[string]interface{}{
				"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed array doesn't change items", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"typedArray": {
					Type: "array",
					Items: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			inputs: map[string]interface{}{
				"typedArray": []interface{}{"a", "b", 3},
			},
		})

		var expected = map[string]interface{}{
			"typedArray": []interface{}{"a", "b", 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed map doesn't change items", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"typedMap": {
					Type: "object",
					AdditionalProperties: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			inputs: map[string]interface{}{
				"typedMap": map[string]interface{}{"a": "b", "c": 3},
			},
		})

		var expected = map[string]interface{}{
			"typedMap": map[string]interface{}{"a": "b", "c": 3},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("string set", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					IsStringSet: true,
					Type:        "object",
				},
			},
			inputs: map[string]interface{}{
				"userAssignedIdentities": []interface{}{
					"a", "b",
				},
			},
		})
		expected := map[string]interface{}{
			"userAssignedIdentities": map[string]interface{}{
				"a": struct{}{},
				"b": struct{}{},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("string set with non-strings", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					IsStringSet: true,
					Type:        "object",
				},
			},
			inputs: map[string]interface{}{
				"userAssignedIdentities": []interface{}{
					"a", "b", 3,
				},
			},
		})
		expected := map[string]interface{}{
			"userAssignedIdentities": map[string]interface{}{
				"a": struct{}{},
				"b": struct{}{},
				// 3 gets ignored
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("missing ref type continues with no change", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"p": {
					Ref: "#/types/azure-native:testing:Type1",
				},
			},
			inputs: map[string]interface{}{
				"p": map[string]interface{}{
					"k": "v",
				},
			},
		})

		expected := map[string]interface{}{
			"p": map[string]interface{}{
				"k": "v",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("missing oneOf type continues with no change", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"oneOf": {
					OneOf: []string{"#types/azure-native:testing:Type1", "#types/azure-native:testing:Type2"},
				},
			},
			inputs: map[string]interface{}{
				"oneOf": map[string]interface{}{
					"prop1": "value",
				},
			},
		})

		expected := map[string]interface{}{
			"oneOf": map[string]interface{}{
				"prop1": "value",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("oneOf type", func(t *testing.T) {
		actual := convert(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"oneOf": {
					OneOf: []string{"#types/azure-native:testing:Type1", "#types/azure-native:testing:Type2"},
				},
			},
			inputs: map[string]interface{}{
				"oneOf": map[string]interface{}{
					"prop1": "type1",
				},
			},
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:Type1": {
					"prop1": {
						Const: "type1",
					},
				},
			},
		})

		expected := map[string]interface{}{
			"oneOf": map[string]interface{}{
				"prop1": "type1",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatching oneOf type", func(t *testing.T) {
		actual := convert(testCaseArgs{
			types: map[string]map[string]resources.AzureAPIProperty{
				"azure-native:testing:Type1": {
					"prop1": {
						Const: "type1",
					},
				},
				"azure-native:testing:Type2": {
					"prop1": {
						Const: "type2",
					},
				},
			},
			props: map[string]resources.AzureAPIProperty{
				"oneOf": {
					OneOf: []string{"#types/azure-native:testing:Type1", "#types/azure-native:testing:Type2"},
				},
			},
			inputs: map[string]interface{}{
				"oneOf": map[string]interface{}{
					"prop1": "foo",
				},
			},
		})

		expected := map[string]interface{}{
			"oneOf": map[string]interface{}{
				"prop1": "foo",
			},
		}

		assert.Equal(t, expected, actual)
	})

	convertNested := func(args testCaseArgs) map[string]interface{} {
		props := map[string]resources.AzureAPIProperty{
			"nested": {
				Ref: "#/types/azure-native:testing:SubType",
			},
		}
		types := map[string]resources.AzureAPIType{
			"azure-native:testing:SubType": {
				Properties: args.props,
			},
		}
		if args.types != nil {
			for typeName, typeProperties := range args.types {
				types[typeName] = resources.AzureAPIType{
					Properties: typeProperties,
				}
			}
		}
		inputs := map[string]interface{}{
			"nested": args.inputs,
		}
		c := SdkShapeConverter{
			Types: types,
		}
		body, err := c.SdkInputsToRequestBody(props, inputs, args.id)
		assert.Nil(t, err)
		return body
	}

	t.Run("nil inputs", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			inputs: nil,
		})

		expected := map[string]interface{}{
			"nested": map[string]interface{}{},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("untyped non-empty values remain unchanged", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"untyped": {},
			},
			inputs: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"untyped": value,
			},
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("any type values", rapid.MakeCheck(func(t *rapid.T) {
		value := propNestedComplex().Draw(t, "value")
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"untyped": {
					Ref: TypeAny,
				},
			},
			inputs: map[string]interface{}{
				"untyped": value,
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"untyped": value,
			},
		}

		assert.Equal(t, expected, actual)
	}))

	t.Run("renamed", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"x-threshold": {
					SdkName: "threshold",
				},
			},
			inputs: map[string]interface{}{
				"threshold": 123,
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"x-threshold": 123,
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("containers", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"container"},
				},
			},
			inputs: map[string]interface{}{
				"prop": "value",
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"container": map[string]interface{}{
					"prop": "value",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("nested containers", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"prop": {
					Containers: []string{"a", "b", "c"},
				},
			},
			inputs: map[string]interface{}{
				"prop": "value",
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": map[string]interface{}{
							"prop": "value",
						},
					},
				},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("mismatched const returns nil", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"const": {
					Const: "value",
				},
			},
			inputs: map[string]interface{}{
				"const": "other",
			},
		})

		expected := map[string]interface{}{
			"nested": map[string]interface{}(nil),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("array of empties not changed", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"emptyArray": {
					Type: "array",
				},
			},
			inputs: map[string]interface{}{
				"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"emptyArray": []interface{}{nil, []interface{}{}, map[string]interface{}{}},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("map of empties unchanged", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"emptyDict": {
					Type: "object",
				},
			},
			inputs: map[string]interface{}{
				"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"emptyDict": map[string]interface{}{"a": nil, "b": map[string]interface{}{}, "c": []interface{}{}},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed array doesn't change items", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"typedArray": {
					Type: "array",
					Items: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			inputs: map[string]interface{}{
				"typedArray": []interface{}{"a", "b", 3},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"typedArray": []interface{}{"a", "b", 3},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("typed map doesn't change items", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"typedMap": {
					Type: "object",
					AdditionalProperties: &resources.AzureAPIProperty{
						Type: "string",
					},
				},
			},
			inputs: map[string]interface{}{
				"typedMap": map[string]interface{}{"a": "b", "c": 3},
			},
		})

		var expected = map[string]interface{}{
			"nested": map[string]interface{}{
				"typedMap": map[string]interface{}{"a": "b", "c": 3},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("string set", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"userAssignedIdentities": {
					IsStringSet: true,
					Type:        "object",
				},
			},
			inputs: map[string]interface{}{
				"userAssignedIdentities": []interface{}{
					"a", "b",
				},
			},
		})
		expected := map[string]interface{}{
			"nested": map[string]interface{}{
				"userAssignedIdentities": map[string]interface{}{
					"a": struct{}{},
					"b": struct{}{},
				},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("missing ref type continues with no change", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"p": {
					Ref: "#/types/azure-native:testing:Type1",
				},
			},
			inputs: map[string]interface{}{
				"p": map[string]interface{}{
					"k": "v",
				},
			},
		})

		expected := map[string]interface{}{
			"nested": map[string]interface{}{
				"p": map[string]interface{}{
					"k": "v",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("missing oneOf type continues with no change", func(t *testing.T) {
		actual := convertNested(testCaseArgs{
			props: map[string]resources.AzureAPIProperty{
				"oneOf": {
					OneOf: []string{"#types/azure-native:testing:Type1", "#types/azure-native:testing:Type2"},
				},
			},
			inputs: map[string]interface{}{
				"oneOf": map[string]interface{}{
					"prop1": "value",
				},
			},
		})

		expected := map[string]interface{}{
			"nested": map[string]interface{}{
				"oneOf": map[string]interface{}{
					"prop1": "value",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}
