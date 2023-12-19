// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
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
	actual := calculateDetailedDiff(&res, emptyTypes, &diff)
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
	actual := calculateDetailedDiff(&res, emptyTypes, &diff)
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
	lookupType := func(t string) (*resources.AzureAPIType, bool, error) {
		if strings.HasSuffix(t, fooTypeName) {
			return &resources.AzureAPIType{
				Properties: map[string]resources.AzureAPIProperty{
					"ps1": {},
					"ps2": {ForceNew: true},
				},
			}, true, nil
		}
		return nil, false, nil
	}

	_, ok, err := lookupType(fooTypeName)
	require.NoError(t, err)
	require.True(t, ok)

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
	actual := calculateDetailedDiff(&res, lookupType, &diff)
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

func TestResourceGroupNameDiffingIsCaseInsensitive(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "path",
				Name:     "resourceGroupName",
			},
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"something": {Type: "string"},
					},
				},
			},
		},
	}
	variantsSame := []string{"MyRG", "myrg", "MYRG", "MyRg"}
	for _, oldValue := range variantsSame {
		for _, newValue := range variantsSame {
			diff := resource.ObjectDiff{
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"resourceGroupName": {
						Old: resource.PropertyValue{V: oldValue},
						New: resource.PropertyValue{V: newValue},
					},
				},
			}
			actual := calculateDetailedDiff(&res, emptyTypes, &diff)
			expected := map[string]*rpc.PropertyDiff{}
			assert.Equal(t, expected, actual)
		}
	}
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
			actual := calculateDetailedDiff(&res, emptyTypes, &diff)
			expected := map[string]*rpc.PropertyDiff{}
			assert.Equal(t, expected, actual)
		}
	}
}

func TestSkuDiffingIsInsensitiveToAksPermutations(t *testing.T) {
	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"sku": {Type: "object"},
					},
				},
			},
		},
	}
	testCases := [][]string{
		[]string{"Basic", "Paid", "Base", "Standard", "equal"},
		[]string{"Base", "Standard", "Basic", "Paid", "equal"},
		[]string{"Base", "Free", "Basic", "Free", "equal"},
		[]string{"Basic", "Paid", "Basic", "Standard", "equal"},
		[]string{"Basic", "Paid", "Basic", "Paid", "equal"},
		[]string{"Base", "Standard", "Basic", "Free", "not equal"},
		[]string{"Premium", "Standard", "Basic", "Free", "not equal"},
	}
	for _, testCase := range testCases {
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"sku": {
					Old: resource.PropertyValue{V: resource.PropertyMap{
						"name": {V: testCase[0]},
						"tier": {V: testCase[1]}},
					},
					New: resource.PropertyValue{V: resource.PropertyMap{
						"name": {V: testCase[2]},
						"tier": {V: testCase[3]}},
					},
				},
			},
		}
		actual := calculateDetailedDiff(&res, emptyTypes, &diff)
		if testCase[4] == "equal" {
			expected := map[string]*rpc.PropertyDiff{}
			assert.Equal(t, expected, actual)
		} else {
			assert.Equal(t, 1, len(actual))
		}
	}
}

var emptyTypes resources.TypeLookupFunc = func(t string) (*resources.AzureAPIType, bool, error) {
	return nil, false, nil
}

func TestChangesAndReplacements_AddedPropertyCausesDiff(t *testing.T) {
	changes, replacements := calculateChangesAndReplacementsForOneAddedProperty(t, "newvalue", nil)
	assert.Len(t, changes, 1)
	assert.Len(t, replacements, 0)
}

func TestChangesAndReplacements_AddedPropertyWithDefaultCausesNoDiff(t *testing.T) {
	changes, replacements := calculateChangesAndReplacementsForOneAddedProperty(t, "defaultvalue", "defaultvalue")
	assert.Len(t, changes, 0)
	assert.Len(t, replacements, 0)
}

func TestChangesAndReplacements_AddedPropertyWithDefaultButDifferentValueCausesDiff(t *testing.T) {
	changes, replacements := calculateChangesAndReplacementsForOneAddedProperty(t, "newvalue", "defaultvalue")
	assert.Len(t, changes, 1)
	assert.Len(t, replacements, 0)
}

func calculateChangesAndReplacementsForOneAddedProperty(t *testing.T, value string, defaultValue interface{}) ([]string, []string) {
	propertyName := "p1"

	var detailedDiff map[string]*rpc.PropertyDiff = map[string]*rpc.PropertyDiff{
		propertyName: &rpc.PropertyDiff{
			Kind: rpc.PropertyDiff_ADD,
		},
	}

	oldInputs := resource.PropertyMap{}
	oldState := resource.PropertyMap{}

	newInputs := resource.PropertyMap{
		resource.PropertyKey(propertyName): {V: value},
	}

	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						propertyName: {
							Type:    "string",
							Default: defaultValue,
						},
					},
				},
			},
		},
	}

	return calculateChangesAndReplacements(detailedDiff, oldInputs, newInputs, oldState, res)
}

func TestDiffKeyedArrays(t *testing.T) {
	old := []resource.PropertyValue{
		{V: resource.PropertyMap{
			"p1": {V: "unchanged"},
			"p2": {V: "v2"},
			"p3": {V: "v3"},
		}},
		{V: resource.PropertyMap{
			"p1": {V: "updated"},
			"p2": {V: "v2"},
			"p3": {V: "v3"},
		}},
		{V: resource.PropertyMap{
			"p1": {V: "will be deleted"},
			"p2": {V: "v2"},
			"p3": {V: "v3"},
		}},
	}
	new := []resource.PropertyValue{
		{V: resource.PropertyMap{
			"p1": {V: "updated"},
			"p2": {V: "v2222"},
			"p3": {V: "v3333"},
		}},
		{V: resource.PropertyMap{
			"p1": {V: "new"},
			"p2": {V: "v222"},
		}},
		{V: resource.PropertyMap{
			"p1": {V: "unchanged"},
			"p2": {V: "v2"},
			"p3": {V: "v3"},
		}},
	}
	keys := []string{"p1"}

	properties := map[string]resources.AzureAPIProperty{}
	diff := diffKeyedArrays(properties, keys, old, new, "")

	assert.NotNil(t, diff)
	assert.NotNil(t, diff.Old)
	assert.NotNil(t, diff.New)

	assert.Equal(t, 1, len(diff.Array.Adds))
	assert.Equal(t, new[1], diff.Array.Adds[1])

	assert.Equal(t, 1, len(diff.Array.Updates))
	assert.Contains(t, diff.Array.Updates, 0)
	update := diff.Array.Updates[0]
	assert.NotNil(t, update.Object)
	assert.Contains(t, update.Object.Updates, resource.PropertyKey("p2"))
	assert.Contains(t, update.Object.Updates, resource.PropertyKey("p3"))
	assert.Contains(t, update.Object.Sames, resource.PropertyKey("p1"))

	assert.Equal(t, 1, len(diff.Array.Sames))
	assert.Equal(t, old[0], diff.Array.Sames[2])

	assert.Equal(t, 1, len(diff.Array.Deletes))
	assert.Equal(t, old[2], diff.Array.Deletes[2])
}
