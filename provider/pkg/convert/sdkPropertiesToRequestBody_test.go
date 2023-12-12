// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
)

func getSdkPropertiesToRequestBodyConverter() SdkShapeConverter {
	return SdkShapeConverter{
		Types: map[string]resources.AzureAPIType{
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
		},
	}
}

func getSampleResource() resources.AzureAPIResource {
	return resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
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
				Value: &resources.AzureAPIProperty{
					SdkName: "networkInterfaceName",
				},
			},
		},
		Response: map[string]resources.AzureAPIProperty{
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
	}
}

func TestSdkPropertiesToRequestBody(t *testing.T) {
	c := getSdkPropertiesToRequestBodyConverter()
	bodyProperties := getSampleResource().PutParameters[0].Body.Properties
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
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps, "/sub/123/rg/my/virtualNetwork/abc")
	sampleAPIPackage := map[string]interface{}{
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
	assert.Equal(t, sampleAPIPackage, data)
}

func TestSdkPropertiesToRequestBodyEmptyCollections(t *testing.T) {
	c := getSdkPropertiesToRequestBodyConverter()
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
	bodyProperties := getSampleResource().PutParameters[0].Body.Properties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, emptyCollectionData, "")
	assert.Equal(t, expectedBody, actualBody)
}

func TestSdkPropertiesToRequestBodyNilCollections(t *testing.T) {
	c := getSdkPropertiesToRequestBodyConverter()
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
	bodyProperties := getSampleResource().PutParameters[0].Body.Properties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, nilCollectionData, "")
	assert.Equal(t, expectedBody, actualBody)
}

func TestSdkPropertiesToRequestBodySubResource(t *testing.T) {
	c := getSdkPropertiesToRequestBodyConverter()
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
	bodyProperties := getSampleResource().PutParameters[0].Body.Properties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, sdkData, "/sub/456/rg/my/network/abc")
	assert.Equal(t, expectedBody, actualBody)
}
