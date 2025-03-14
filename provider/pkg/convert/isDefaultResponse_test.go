package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
)

func TestIsDefaultResponseTrue(t *testing.T) {
	c := getMockConverter()
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
	actual := c.IsDefaultResponse(resourcePutParameters(), response, defaultBody)
	assert.True(t, actual)
}

func TestIsDefaultResponseFalse(t *testing.T) {
	c := getMockConverter()
	defaultBody := map[string]interface{}{}
	response1 := map[string]interface{}{
		"p1": true,
	}
	actual1 := c.IsDefaultResponse(resourcePutParameters(), response1, defaultBody)
	assert.False(t, actual1)

	response2 := map[string]interface{}{
		"untypedArray": []string{"1", "2"},
	}
	actual2 := c.IsDefaultResponse(resourcePutParameters(), response2, defaultBody)
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
	actual3 := c.IsDefaultResponse(resourcePutParameters(), response3, defaultBody)
	assert.False(t, actual3)
}

func resourcePutParameters() []resources.AzureAPIParameter {
	return []resources.AzureAPIParameter{
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
	}
}

func getMockConverter() SdkShapeConverter {
	return NewSdkShapeConverterFull(map[string]resources.AzureAPIType{
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
	})
}
