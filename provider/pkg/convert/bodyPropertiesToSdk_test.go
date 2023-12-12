// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
)

func TestResponseToSdkOutputs(t *testing.T) {
	props := map[string]resources.AzureAPIProperty{
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
	}

	var c = SdkShapeConverter{Types: map[string]resources.AzureAPIType{
		"azure-native:testing:Structure": {
			Properties: map[string]resources.AzureAPIProperty{
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
		},
		"azure-native:testing:StructureResponse": {
			Properties: map[string]resources.AzureAPIProperty{
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
			Properties: map[string]resources.AzureAPIProperty{
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
		},
		"azure-native:testing:MoreItem": {
			Properties: map[string]resources.AzureAPIProperty{
				"aaa": {
					SdkName: "Aaa",
				},
				"bbb": {
					Containers: []string{"ccc"},
				},
			},
		},
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
		"azure-native:testing:SubResource": {
			Properties: map[string]resources.AzureAPIProperty{
				"id": {
					Type: "string",
				},
			},
		},
	}}

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
	outputs := c.BodyPropertiesToSDK(props, sampleAPIPackage)
	assert.Equal(t, sampleSdkProps, outputs)
}
