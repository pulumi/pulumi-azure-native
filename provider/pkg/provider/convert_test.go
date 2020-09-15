// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = &azurermProvider{
	resourceMap: &AzureAPIMetadata{
		Types: map[string]AzureAPIType{
			"azurerm:testing:Structure": {
				Properties: map[string]AzureAPIProperty{
					"v1": {},
					"v2": {},
					"v3-odd": {
						SdkName: "v3",
					},
					"v4-nested": {
						SdkName:   "v4",
						Container: "props",
					},
				},
			},
			"azurerm:testing:More": {
				Properties: map[string]AzureAPIProperty{
					"items": {
						Items: &AzureAPIProperty{
							Ref: "#/types/azurerm:testing:MoreItem",
						},
					},
				},
			},
			"azurerm:testing:MoreItem": {
				Properties: map[string]AzureAPIProperty{
					"aaa": {
						SdkName: "Aaa",
					},
					"bbb": {
						Container: "ccc",
					},
				},
			},
			"azurerm:testing:OptionA": {
				Properties: map[string]AzureAPIProperty{
					"type": {
						Const: "AAA",
					},
					"a": {
						Container: "aa",
					},
				},
			},
			"azurerm:testing:OptionB": {
				Properties: map[string]AzureAPIProperty{
					"type": {
						Const: "BBB",
					},
					"b": {
						Container: "bb",
					},
				},
			},
		},
		Resources: map[string]AzureAPIResource{
			"r1": {
				PutParameters: []AzureAPIParameter{
					{
						Body: &AzureAPIType{
							Properties: map[string]AzureAPIProperty{
								"name": {},
								"x-threshold": {
									SdkName: "threshold",
								},
								"structure": {
									Ref: "#/types/azurerm:testing:Structure",
								},
								"p1": {
									Container: "properties",
								},
								"p2": {
									Container: "properties",
								},
								"more": {
									Container: "properties",
									Ref:       "#/types/azurerm:testing:More",
								},
								"union": {
									OneOf: []string{"#/types/azurerm:testing:OptionA", "#/types/azurerm:testing:OptionB"},
								},
								"tags": {},
							},
						},
					},
				},
				Response: map[string]AzureAPIProperty{
					"name": {},
					"x-threshold": {
						SdkName: "threshold",
					},
					"structure": {
						Ref: "#/types/azurerm:testing:Structure",
					},
					"p1": {
						Container: "properties",
					},
					"p2": {
						Container: "properties",
					},
					"more": {
						Container: "properties",
						Ref:       "#/types/azurerm:testing:More",
					},
					"union": {
						OneOf: []string{"#/types/azurerm:testing:OptionA", "#/types/azurerm:testing:OptionB"},
					},
					"tags": {},
				},
			},
		},
	},
}

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
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"aaa": "111", "ccc": map[string]interface{}{"bbb": "333"}},
				map[string]interface{}{"aaa": "222"},
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
	"more": map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{"Aaa": "111", "bbb": "333"},
			map[string]interface{}{"Aaa": "222"},
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

func TestResponseToSdkOutputs(t *testing.T) {
	outputs := p.responseToSdkOutputs(p.resourceMap.Resources["r1"].Response, sampleAPIPackage)
	assert.Equal(t, sampleSdkProps, outputs)
}

func TestSdkPropertiesToRequest(t *testing.T) {
	bodyProperties := p.resourceMap.Resources["r1"].PutParameters[0].Body.Properties
	data := p.sdkPropertiesToRequest(bodyProperties, sampleSdkProps)
	assert.Equal(t, sampleAPIPackage, data)
}
