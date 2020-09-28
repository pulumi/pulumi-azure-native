// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = &azureNextGenProvider{
	resourceMap: &AzureAPIMetadata{
		Types: map[string]AzureAPIType{
			"azure-nextgen:testing:Structure": {
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
			"azure-nextgen:testing:More": {
				Properties: map[string]AzureAPIProperty{
					"items": {
						Items: &AzureAPIProperty{
							Ref: "#/types/azure-nextgen:testing:MoreItem",
						},
					},
				},
			},
			"azure-nextgen:testing:MoreItem": {
				Properties: map[string]AzureAPIProperty{
					"aaa": {
						SdkName: "Aaa",
					},
					"bbb": {
						Container: "ccc",
					},
				},
			},
			"azure-nextgen:testing:OptionA": {
				Properties: map[string]AzureAPIProperty{
					"type": {
						Const: "AAA",
					},
					"a": {
						Container: "aa",
					},
				},
			},
			"azure-nextgen:testing:OptionB": {
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
						Location: "body",
						Body: &AzureAPIType{
							Properties: map[string]AzureAPIProperty{
								"name": {},
								"x-threshold": {
									SdkName: "threshold",
								},
								"structure": {
									Ref: "#/types/azure-nextgen:testing:Structure",
								},
								"p1": {
									Container: "properties",
								},
								"p2": {
									Container: "properties",
								},
								"more": {
									Container: "properties",
									Ref:       "#/types/azure-nextgen:testing:More",
								},
								"union": {
									OneOf: []string{"#/types/azure-nextgen:testing:OptionA", "#/types/azure-nextgen:testing:OptionB"},
								},
								"tags": {},
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
						Name:     "networkInterfaceName",
					},
				},
				Response: map[string]AzureAPIProperty{
					"name": {},
					"x-threshold": {
						SdkName: "threshold",
					},
					"structure": {
						Ref: "#/types/azure-nextgen:testing:Structure",
					},
					"p1": {
						Container: "properties",
					},
					"p2": {
						Container: "properties",
					},
					"more": {
						Container: "properties",
						Ref:       "#/types/azure-nextgen:testing:More",
					},
					"union": {
						OneOf: []string{"#/types/azure-nextgen:testing:OptionA", "#/types/azure-nextgen:testing:OptionB"},
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
		_, err := p.parseResourceID(testCase.id, testCase.path)
		assert.Error(t, err)
	}
}

func TestParseValidResourceID(t *testing.T) {
	id := "/subscriptions/0282681f-7a9e-123b-40b2-96babd57a8a1/resourcegroups/pulumi-name/providers/Microsoft.Network/networkInterfaces/pulumi-nic/ipConfigurations/ipconfig1"
	path := "/subscriptions/{subscriptionID}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	actual, err := p.parseResourceID(id, path)
	assert.NoError(t, err)
	expected := map[string]string{
		"subscriptionID":       "0282681f-7a9e-123b-40b2-96babd57a8a1",
		"resourceGroupName":    "pulumi-name",
		"networkInterfaceName": "pulumi-nic",
		"ipConfigurationName":  "ipconfig1",
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
		"networkInterfaceName": "nic-name",
	}
	inputs := p.responseToSdkInputs(p.resourceMap.Resources["r1"].PutParameters, pathValues, responseForInputCalculation)
	assert.Equal(t, calculatedInputs, inputs)
}
