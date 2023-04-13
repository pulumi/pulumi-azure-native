// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"encoding/json"
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateDiffBodyProperties(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"p1": {Type: "string"},
						"p2": {Type: "number"},
						"p3": {Type: "boolean"},
						"p4": {Ref: "#/types/azure-native:foobar/v20200101:SomeType"},
						"p5": {Type: "array", Items: &resources.AzureAPIProperty{
							Type: "string",
						}},
						"p6": {Type: "array", Items: &resources.AzureAPIProperty{
							Ref: "#/types/azure-native:foobar/v20200101:SomeType",
						}},
						"location": {Type: "string"},
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
			"p5": {
				Array: &resource.ArrayDiff{
					Updates: map[int]resource.ValueDiff{
						0: {
							Old: resource.PropertyValue{V: "old0value"},
							New: resource.PropertyValue{V: "new0value"},
						},
					},
					Adds: map[int]resource.PropertyValue{
						1: {V: "new1value"},
					},
					Deletes: map[int]resource.PropertyValue{
						2: {V: "old2value"},
					},
				},
			},
			"p6": {
				Array: &resource.ArrayDiff{
					Updates: map[int]resource.ValueDiff{
						0: {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"p61": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
								},
								Adds: map[resource.PropertyKey]resource.PropertyValue{
									"p62": {V: 123},
								},
							},
						},
					},
				},
			},
			"location": {
				Old: resource.PropertyValue{V: "East US"},
				New: resource.PropertyValue{V: "East US 2"},
			},
		},
		Adds: map[resource.PropertyKey]resource.PropertyValue{
			"p2": {V: 123},
		},
		Deletes: map[resource.PropertyKey]resource.PropertyValue{
			"p3": {V: true},
		},
	}
	emptyTypes := resources.NewPartialMap[resources.AzureAPIType]()
	actual := calculateDetailedDiff(&res, &emptyTypes, &diff)
	expected := map[string]*rpc.PropertyDiff{
		"p1":          {Kind: rpc.PropertyDiff_UPDATE},
		"p2":          {Kind: rpc.PropertyDiff_ADD},
		"p3":          {Kind: rpc.PropertyDiff_DELETE},
		"p4.p41":      {Kind: rpc.PropertyDiff_UPDATE},
		"p4.p42":      {Kind: rpc.PropertyDiff_ADD},
		"p4.p43":      {Kind: rpc.PropertyDiff_DELETE},
		"p4.p44.p441": {Kind: rpc.PropertyDiff_UPDATE},
		"p5[0]":       {Kind: rpc.PropertyDiff_UPDATE},
		"p5[1]":       {Kind: rpc.PropertyDiff_ADD},
		"p5[2]":       {Kind: rpc.PropertyDiff_DELETE},
		"p6[0].p61":   {Kind: rpc.PropertyDiff_UPDATE},
		"p6[0].p62":   {Kind: rpc.PropertyDiff_ADD},
		"location":    {Kind: rpc.PropertyDiff_UPDATE},
	}
	assert.Equal(t, expected, actual)
}

func TestCalculateDiffReplacesPathParameters(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "path",
				Name:     "p1",
			},
			{
				Location: "path",
				Name:     "p2",
				Value: &resources.AzureAPIProperty{
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
	emptyTypes := resources.NewPartialMap[resources.AzureAPIType]()
	actual := calculateDetailedDiff(&res, &emptyTypes, &diff)
	expected := map[string]*rpc.PropertyDiff{
		"p1":    {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"Prop2": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"q1":    {Kind: rpc.PropertyDiff_UPDATE},
		"b1":    {Kind: rpc.PropertyDiff_UPDATE},
	}
	assert.Equal(t, expected, actual)
}

func TestCalculateDiffReplacesBodyProperties(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"p1": {Type: "string"},
						"p2": {Type: "number", ForceNew: true},
						"p3": {Ref: "#/types/azure-native:foobar/v20200101:FooType"},
						"p4": {Ref: "#/types/azure-native:foobar/v20200101:FooType", ForceNew: true},
						"p5": {
							Items: &resources.AzureAPIProperty{
								Ref: "#/types/azure-native:foobar/v20200101:FooType",
							},
						},
					},
				},
			},
		},
	}
	fooTypeName := "azure-native:foobar/v20200101:FooType"
	fullTypes := map[string]resources.AzureAPIType{
		"azure-native:foobar/v20200101:FooType": {
			Properties: map[string]resources.AzureAPIProperty{
				"ps1": {},
				"ps2": {ForceNew: true},
			},
		},
	}
	typeData, err := json.Marshal(fullTypes)
	require.NoError(t, err)
	var types resources.PartialMap[resources.AzureAPIType]
	err = json.Unmarshal(typeData, &types)
	require.NoError(t, err)

	testType, ok, err := types.Get(fooTypeName)
	require.NoError(t, err)
	require.True(t, ok)
	assert.Equal(t, testType, fullTypes[fooTypeName])

	diff := resource.ObjectDiff{
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
			"p2": {
				Old: resource.PropertyValue{V: 1},
				New: resource.PropertyValue{V: 2},
			},
			"p3": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{
						"ps1": {
							Old: resource.PropertyValue{V: "oldvalue"},
							New: resource.PropertyValue{V: "newvalue"},
						},
						"ps2": {
							Old: resource.PropertyValue{V: "oldvalue"},
							New: resource.PropertyValue{V: "newvalue"},
						},
					},
				},
			},
			"p4": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{
						"ps1": {
							Old: resource.PropertyValue{V: "oldvalue"},
							New: resource.PropertyValue{V: "newvalue"},
						},
					},
				},
			},
			"p5": {
				Array: &resource.ArrayDiff{
					Updates: map[int]resource.ValueDiff{
						0: {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"ps1": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
									"ps2": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	actual := calculateDetailedDiff(&res, resources.GoMap[resources.AzureAPIType](fullTypes), &diff)
	expected := map[string]*rpc.PropertyDiff{
		"p1":        {Kind: rpc.PropertyDiff_UPDATE},
		"p2":        {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"p3.ps1":    {Kind: rpc.PropertyDiff_UPDATE},
		"p3.ps2":    {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"p4.ps1":    {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		"p5[0].ps1": {Kind: rpc.PropertyDiff_UPDATE},
		"p5[0].ps2": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
	}
	assert.Equal(t, expected, actual)
}

func TestApplyDiff(t *testing.T) {
	// Inputs are the base to apply to.
	inputs := resource.PropertyMap{
		"p1": {V: "oldvalue"},
		"p2": {V: "iamdeleted"},
		"p4": {
			V: resource.PropertyMap{
				"p4.1": {V: "oldervalue"},
				"p4.3": {V: "willbedeleted"},
			},
		},
		"p5": {
			V: []resource.PropertyValue{
				{
					V: resource.PropertyMap{
						"p5.1": {V: "nochange1"},
						"p5.2": {
							V: resource.PropertyMap{
								"p5.2.1": {V: "oldestvalue"},
								"p5.2.2": {V: "nochange2"},
							},
						},
					},
				},
			},
		},
	}
	// Diff represents a difference between old outputs and new outputs.
	diff := resource.ObjectDiff{
		Adds: resource.PropertyMap{
			"p3": {V: "newkey"},
		},
		Deletes: resource.PropertyMap{
			"p2": {V: "iamdeleted"},
		},
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
			"p4": {
				Old: resource.PropertyValue{
					V: resource.PropertyMap{
						"p4.1": {V: "oldervalue"},
						"p4.2": {V: "same"},
						"p4.3": {V: "willbedeleted"},
					},
				},
				New: resource.PropertyValue{
					V: resource.PropertyMap{
						"p4.1": {V: "newervalue"},
						"p4.2": {V: "same"},
						"p4.4": {V: "thiswasadded"},
					},
				},
			},
			"p5": {
				Old: resource.PropertyValue{
					V: []resource.PropertyValue{
						{
							V: resource.PropertyMap{
								"p5.1": {V: "nochange1"},
								"p5.2": {
									V: resource.PropertyMap{
										"p5.2.1": {V: "oldestvalue"},
										"p5.2.2": {V: "nochange2"},
										"p5.2.3": {V: "nochange3"},
									},
								},
								"p5.3": {V: "nochange4"},
							},
						},
					},
				},
				New: resource.PropertyValue{
					V: []resource.PropertyValue{
						{
							V: resource.PropertyMap{
								"p5.1": {V: "nochange1"},
								"p5.2": {
									V: resource.PropertyMap{
										"p5.2.1": {V: "newestvalue"},
										"p5.2.2": {V: "nochange2"},
										"p5.2.3": {V: "nochange3"},
									},
								},
								"p5.3": {V: "nochange4"},
							},
						},
						{
							V: resource.PropertyMap{
								"p6.1": {V: "new1"},
								"p6.2": {
									V: resource.PropertyMap{
										"p6.2.1": {V: "new2"},
									},
								},
								"p6.3": {V: "new3"},
							},
						},
					},
				},
			},
		},
	}
	actual := applyDiff(inputs, &diff)
	expected := resource.PropertyMap{
		"p1": {V: "newvalue"},
		"p3": {V: "newkey"},
		"p4": {V: resource.PropertyMap{
			"p4.1": {V: "newervalue"},
			"p4.4": {V: "thiswasadded"},
		}},
		"p5": {
			V: []resource.PropertyValue{
				{
					V: resource.PropertyMap{
						"p5.1": {V: "nochange1"},
						"p5.2": {
							V: resource.PropertyMap{
								"p5.2.1": {V: "newestvalue"},
								"p5.2.2": {V: "nochange2"},
							},
						},
					},
				},
				{
					V: resource.PropertyMap{
						"p6.1": {V: "new1"},
						"p6.2": {
							V: resource.PropertyMap{
								"p6.2.1": {V: "new2"},
							},
						},
						"p6.3": {V: "new3"},
					},
				},
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestLocationDiffingIsInsensitiveToSpacesAndCasing(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"location": {Type: "string"},
					},
				},
			},
		},
	}
	variantsSame := []string{"West US 2", "WestUS2", "west us 2", "westus2"}
	for _, oldValue := range variantsSame {
		for _, newValue := range variantsSame {
			diff := resource.ObjectDiff{
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"location": {
						Old: resource.PropertyValue{V: oldValue},
						New: resource.PropertyValue{V: newValue},
					},
				},
			}
			emptyTypes := resources.NewPartialMap[resources.AzureAPIType]()
			actual := calculateDetailedDiff(&res, &emptyTypes, &diff)
			expected := map[string]*rpc.PropertyDiff{}
			assert.Equal(t, expected, actual)
		}
	}
}
