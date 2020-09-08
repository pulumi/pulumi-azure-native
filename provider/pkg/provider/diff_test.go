// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDiffBodyProperties(t *testing.T) {
	res := AzureAPIResource{
		PutParameters: []AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &AzureAPIType{
					Properties: map[string]AzureAPIProperty{
						"p1": {Type: "string"},
						"p2": {Type: "number"},
						"p3": {Type: "boolean"},
						"p4": {Ref: "#/types/azurerm:foobar/v20200101:SomeType"},
					},
				},
			},
		},
	}
	diff := resource.ObjectDiff{
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
			"p4": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{
						"p41": {
							Old: resource.PropertyValue{V: "oldvalue"},
							New: resource.PropertyValue{V: "newvalue"},
						},
						"p44": {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"p441": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
								},
							},
						},
					},
					Adds: map[resource.PropertyKey]resource.PropertyValue{
						"p42": {V: 123},
					},
					Deletes: map[resource.PropertyKey]resource.PropertyValue{
						"p43": {V: true},
					},
				},
			},
		},
		Adds: map[resource.PropertyKey]resource.PropertyValue{
			"p2": {V: 123},
		},
		Deletes: map[resource.PropertyKey]resource.PropertyValue{
			"p3": {V: true},
		},
	}
	actual, err := calculateDetailedDiff(&res, &diff)
	assert.NoError(t, err)
	expected := map[string]*rpc.PropertyDiff{
		"p1":          {Kind: rpc.PropertyDiff_UPDATE},
		"p2":          {Kind: rpc.PropertyDiff_ADD},
		"p3":          {Kind: rpc.PropertyDiff_DELETE},
		"p4.p41":      {Kind: rpc.PropertyDiff_UPDATE},
		"p4.p42":      {Kind: rpc.PropertyDiff_ADD},
		"p4.p43":      {Kind: rpc.PropertyDiff_DELETE},
		"p4.p44.p441": {Kind: rpc.PropertyDiff_UPDATE},
	}
	assert.Equal(t, expected, actual)
}

func TestCalculateDiffReplacesPathParameters(t *testing.T) {
	res := AzureAPIResource{
		PutParameters: []AzureAPIParameter{
			{
				Location: "path",
				Name:     "p1",
			},
			{
				Location: "path",
				Name:     "p2",
				Value: &AzureAPIProperty{
					SdkName: "Prop2",
				},
			},
			{
				Location: "query",
				Name:     "q1",
			},
			{
				Location: "body",
				Name:     "b1",
			},
		},
	}
	diff := resource.ObjectDiff{
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldpath"},
				New: resource.PropertyValue{V: "newpath"},
			},
			"Prop2": {
				Old: resource.PropertyValue{V: "oldpath"},
				New: resource.PropertyValue{V: "newpath"},
			},
			"q1": {
				Old: resource.PropertyValue{V: "oldquery"},
				New: resource.PropertyValue{V: "newquery"},
			},
			"b1": {
				Old: resource.PropertyValue{V: "oldbody"},
				New: resource.PropertyValue{V: "newbody"},
			},
		},
	}
	actual, err := calculateDetailedDiff(&res, &diff)
	assert.NoError(t, err)
	expected := map[string]*rpc.PropertyDiff{
		"p1":    {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"Prop2": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"q1":    {Kind: rpc.PropertyDiff_UPDATE},
		"b1":    {Kind: rpc.PropertyDiff_UPDATE},
	}
	assert.Equal(t, expected, actual)
}
