// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"encoding/json"
	"fmt"
	"testing"
	"unicode"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
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
						"has.period": {Type: "string"},
						"location":   {Type: "string"},
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
						"p42":        {V: 123},
						"p45.period": {V: "yes"},
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
			"has.period": {},
			"location": {
				Old: resource.PropertyValue{V: "East US"},
				New: resource.PropertyValue{V: "East US 2"},
			},
		},
		Adds: map[resource.PropertyKey]resource.PropertyValue{
			"p2":         {V: 123},
			"has.period": {V: "yes"},
		},
		Deletes: map[resource.PropertyKey]resource.PropertyValue{
			"p3": {V: true},
		},
	}
	actual := calculateDetailedDiff(&res, emptyTypes, &diff)
	expected := map[string]*rpc.PropertyDiff{
		"p1":                  {Kind: rpc.PropertyDiff_UPDATE},
		"p2":                  {Kind: rpc.PropertyDiff_ADD},
		"p3":                  {Kind: rpc.PropertyDiff_DELETE},
		"p4.p41":              {Kind: rpc.PropertyDiff_UPDATE},
		"p4.p42":              {Kind: rpc.PropertyDiff_ADD},
		"p4.p43":              {Kind: rpc.PropertyDiff_DELETE},
		"p4.[\"p45.period\"]": {Kind: rpc.PropertyDiff_ADD},
		"p4.p44.p441":         {Kind: rpc.PropertyDiff_UPDATE},
		"p5[0]":               {Kind: rpc.PropertyDiff_UPDATE},
		"p5[1]":               {Kind: rpc.PropertyDiff_ADD},
		"p5[2]":               {Kind: rpc.PropertyDiff_DELETE},
		"p6[0].p61":           {Kind: rpc.PropertyDiff_UPDATE},
		"p6[0].p62":           {Kind: rpc.PropertyDiff_ADD},
		"[\"has.period\"]":    {Kind: rpc.PropertyDiff_ADD},
		"location":            {Kind: rpc.PropertyDiff_UPDATE},
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
	const topLevelProperty = "p1"

	oldToNew := resource.ValueDiff{
		Old: resource.PropertyValue{V: "oldvalue"},
		New: resource.PropertyValue{V: "newvalue"},
	}

	test := func(t *testing.T, property resources.AzureAPIProperty, refProperties map[string]resources.AzureAPIProperty, diff resource.ObjectDiff, expected map[string]*rpc.PropertyDiff) {
		res := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Name:     "bodyProperties",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							topLevelProperty: property,
						},
					},
				},
			},
		}

		lookupType := func(t string) (*resources.AzureAPIType, bool, error) {
			return &resources.AzureAPIType{
				Properties: refProperties,
			}, true, nil
		}

		actual := calculateDetailedDiff(&res, lookupType, &diff)
		assert.Equal(t, expected, actual)
	}

	t.Run("primitive property", func(t *testing.T) {
		property := resources.AzureAPIProperty{Type: "string"}
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				topLevelProperty: oldToNew,
			},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty: {Kind: rpc.PropertyDiff_UPDATE},
		}
		test(t, property, nil, diff, expected)
	})

	t.Run("primitive property with forceNew", func(t *testing.T) {
		property := resources.AzureAPIProperty{Type: "string", ForceNew: true}
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				topLevelProperty: oldToNew,
			},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty: {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		}
		test(t, property, nil, diff, expected)
	})

	t.Run("object with updates", func(t *testing.T) {
		property := resources.AzureAPIProperty{Ref: "#/types/foo"}
		refProperties := map[string]resources.AzureAPIProperty{
			"inner1": {},
			"inner2": {ForceNew: true},
		}
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				topLevelProperty: {
					Object: &resource.ObjectDiff{
						Updates: map[resource.PropertyKey]resource.ValueDiff{
							"inner1": oldToNew,
							"inner2": oldToNew,
						},
					},
				},
			},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty + ".inner1": {Kind: rpc.PropertyDiff_UPDATE},
			topLevelProperty + ".inner2": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		}
		test(t, property, refProperties, diff, expected)
	})

	t.Run("object with updates and forceNew", func(t *testing.T) {
		property := resources.AzureAPIProperty{Ref: "#/types/foo", ForceNew: true}
		refProperties := map[string]resources.AzureAPIProperty{
			"inner1": {},
		}
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				topLevelProperty: {
					Object: &resource.ObjectDiff{
						Updates: map[resource.PropertyKey]resource.ValueDiff{
							"inner1": oldToNew,
						},
					},
				},
			},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty + ".inner1": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		}
		test(t, property, refProperties, diff, expected)
	})

	diffObjectHasPropertyAddedAndDeleted := resource.ObjectDiff{
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			topLevelProperty: {
				Object: &resource.ObjectDiff{
					Deletes: map[resource.PropertyKey]resource.PropertyValue{
						"inner1": {V: "removed"},
					},
					Adds: map[resource.PropertyKey]resource.PropertyValue{
						"inner3": {V: "added"},
					},
				},
			},
		},
	}

	t.Run("object with additions and deletions only", func(t *testing.T) {
		property := resources.AzureAPIProperty{Ref: "#/types/foo"}
		refProperties := map[string]resources.AzureAPIProperty{
			"inner1": {},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty + ".inner1": {Kind: rpc.PropertyDiff_DELETE},
			topLevelProperty + ".inner3": {Kind: rpc.PropertyDiff_ADD},
		}
		test(t, property, refProperties, diffObjectHasPropertyAddedAndDeleted, expected)
	})

	t.Run("object with additions and deletions only and forceNew", func(t *testing.T) {
		property := resources.AzureAPIProperty{Ref: "#/types/foo", ForceNew: true}
		refProperties := map[string]resources.AzureAPIProperty{
			"inner1": {},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty:             {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
			topLevelProperty + ".inner1": {Kind: rpc.PropertyDiff_DELETE},
			topLevelProperty + ".inner3": {Kind: rpc.PropertyDiff_ADD},
		}
		test(t, property, refProperties, diffObjectHasPropertyAddedAndDeleted, expected)
	})

	t.Run("array of object with updates", func(t *testing.T) {
		property := resources.AzureAPIProperty{
			Items: &resources.AzureAPIProperty{Ref: "#/types/foo"},
		}
		refProperties := map[string]resources.AzureAPIProperty{
			"inner1": {},
			"inner2": {ForceNew: true},
		}
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				topLevelProperty: {
					Array: &resource.ArrayDiff{
						Updates: map[int]resource.ValueDiff{
							0: {
								Object: &resource.ObjectDiff{
									Updates: map[resource.PropertyKey]resource.ValueDiff{
										"inner1": oldToNew,
										"inner2": oldToNew,
									},
								},
							},
						},
					},
				},
			},
		}
		expected := map[string]*rpc.PropertyDiff{
			topLevelProperty + "[0].inner1": {Kind: rpc.PropertyDiff_UPDATE},
			topLevelProperty + "[0].inner2": {Kind: rpc.PropertyDiff_UPDATE_REPLACE},
		}
		test(t, property, refProperties, diff, expected)
	})
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

func TestChangesAndReplacements_ReduceNestedPropertiesToTopLevel(t *testing.T) {
	detailedDiff := map[string]*rpc.PropertyDiff{
		"agentPoolProfiles[0].count": &rpc.PropertyDiff{
			Kind: rpc.PropertyDiff_ADD,
		},
	}
	oldState := resource.PropertyMap{}
	oldInputs := resource.PropertyMap{
		resource.PropertyKey("agentPoolProfiles[0].count"): {V: 3},
	}
	newInputs := resource.PropertyMap{
		resource.PropertyKey("agentPoolProfiles[0].count"): {V: 4},
	}

	res := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"agentPoolProfiles": {Type: "array", Items: &resources.AzureAPIProperty{
							Type: "object",
						}},
					},
				},
			},
		},
	}

	changes, replacements := calculateChangesAndReplacements(detailedDiff, oldInputs, newInputs, oldState, res)
	assert.Len(t, changes, 1)
	assert.Equal(t, "agentPoolProfiles", changes[0])
	assert.Empty(t, replacements)
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

func TestNormalizeAzureId(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "",
			output: "",
		},
		{
			input:  "foo/bar",
			output: "foo/bar",
		},
		{
			input:  "/subscriptions/foo/bar",
			output: "/subscriptions/foo/bar",
		},
		{
			input:  "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			output: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
		},
		{
			input:  "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/something",
			output: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
		},
		{
			input:  "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/something/with/a/path",
			output: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something/with/a/path",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, normalizeAzureId(testCase.input))
	}
}

func TestStringsEqualCaseInsensitiveAzureIds(t *testing.T) {
	for _, equalCase := range []struct {
		id1 string
		id2 string
	}{
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourcegroups/rg/providers/Microsoft.compute/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourcegroups/rg/providers/microsoft.COMPUTE/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourceGroups/rg/Providers/microsoft.Compute/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg",
			id2: "/subscriptions/123/resourcegroups/rg",
		},
	} {
		assert.True(t, stringsEqualCaseInsensitiveAzureIds(equalCase.id1, equalCase.id2))
	}

	for _, notEqualCase := range []struct {
		id1 string
		id2 string
	}{
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/456/resourcegroups/rg/providers/microsoft.compute/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourcegroups/rg2/providers/microsoft.COMPUTE/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourcegroups/rg/providers/Microsoft.compute/somethingElse",
		},

		{
			id1: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something",
			id2: "/subscriptions/123/resourceGroups/rg/providers/microsoft.ComputeBetter/something",
		},
		{
			id1: "/subscriptions/123/resourcegroups/rg",
			id2: "/subscriptions/123/resourceGroups/rg",
		},
	} {
		assert.False(t, stringsEqualCaseInsensitiveAzureIds(notEqualCase.id1, notEqualCase.id2))
	}
}

func TestDiffKeyedArrays(t *testing.T) {
	t.Run("basic example", func(t *testing.T) {
		// The unique identifier for each object is the "p1" property.
		keys := []string{"p1"}

		// The first object will be unchanged, the second updated, the third removed and replaced
		// with a different one. The unchanged one has case differences that shouldn't matter.
		old := []resource.PropertyValue{
			{V: resource.PropertyMap{
				"p1": {V: "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Compute/something"},
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
		// Change the order compared to `old` to prove that objects are identified by their key properties.
		new := []resource.PropertyValue{
			{V: resource.PropertyMap{
				"p1": {V: "updated"},
				"p2": {V: "v2222"},
				// "p3" is deleted
				"p4": {V: "v4444"}, // added property
			}},
			{V: resource.PropertyMap{
				"p1": {V: "new"},
				"p2": {V: "v222"},
			}},
			{V: resource.PropertyMap{
				"p1": {V: "/subscriptions/123/resourcegroups/rg/providers/microsoft.compute/something"},
				"p2": {V: "v2"},
				"p3": {V: "v3"},
			}},
		}

		properties := map[string]resources.AzureAPIProperty{}
		diff, validKeyedArray := diffKeyedArrays(properties, keys, old, new, "")
		assert.True(t, validKeyedArray)

		assert.NotNil(t, diff)
		assert.NotNil(t, diff.Old)
		assert.NotNil(t, diff.New)

		assert.Equal(t, 1, len(diff.Array.Adds))
		assert.Equal(t, new[1], diff.Array.Adds[1])

		require.Equal(t, 1, len(diff.Array.Sames))
		require.Contains(t, diff.Array.Sames, 2)
		sameOld := old[0].ObjectValue()
		sameNew := diff.Array.Sames[2].ObjectValue()
		assert.Equal(t, normalizeAzureId(sameOld["p1"].StringValue()), normalizeAzureId(sameNew["p1"].StringValue()))
		assert.Equal(t, sameOld["p2"], sameNew["p2"])

		assert.Equal(t, 1, len(diff.Array.Deletes))
		assert.Equal(t, old[2], diff.Array.Deletes[2])

		assert.Equal(t, 1, len(diff.Array.Updates))
		assert.Contains(t, diff.Array.Updates, 0)
		update := diff.Array.Updates[0]
		assert.NotNil(t, update.Object)
		assert.Contains(t, update.Object.Sames, resource.PropertyKey("p1"))
		assert.Contains(t, update.Object.Updates, resource.PropertyKey("p2"))
		assert.Contains(t, update.Object.Deletes, resource.PropertyKey("p3"))
		assert.Contains(t, update.Object.Adds, resource.PropertyKey("p4"))
	})

	t.Run("recognizes invalid keyed array", func(t *testing.T) {
		olds := map[string][]resource.PropertyValue{
			"not an object": {
				{V: resource.PropertyMap{"p1": {V: "oldvalue"}}},
				{V: "oldvalue"},
			},
			"missing keys": {
				{V: resource.PropertyMap{"p1": {V: "oldvalue"}}},
				{V: resource.PropertyMap{"foo": {V: "bar"}}},
			},
			"duplicate key value": {
				{V: resource.PropertyMap{"p1": {V: "oldvalue"}}},
				{V: resource.PropertyMap{"p1": {V: "oldvalue"}}},
			},
		}

		for name, old := range olds {
			diff, ok := diffKeyedArrays(map[string]resources.AzureAPIProperty{}, []string{"p1"}, old, old, "")
			require.False(t, ok, name)
			require.Nil(t, diff, name)
		}
	})
}

// A series of property-based tests that use rapid to generate random inputs and random mutations
// of these inputs, then verify that the diffing algorithm produces the expected results.
func TestDiffKeyedArraysRapid(tt *testing.T) {
	runRapidTest := func(t *rapid.T,
		keyProperties []string,
		objectGenerator *rapid.Generator[resource.PropertyValue],
		mutator mutator) {

		array := makeArrayGenerator(objectGenerator)
		old := array.Draw(t, "old").ArrayValue()

		changes, new := mutator(t, old, keyProperties)

		diff, ok := diffKeyedArrays(map[string]resources.AzureAPIProperty{}, keyProperties, old, new, "")
		assert.True(t, ok)
		compareDiffWithRecordedChanges(t, diff, changes, old, new)
	}

	tt.Run("no change, single key", rapid.MakeCheck(func(t *rapid.T) {
		keyProperty := []string{"p1"}
		runRapidTest(t, keyProperty, makeNestedObjectGenerator(keyProperty...), noopMutator)
	}))

	tt.Run("no change, multiple keys", rapid.MakeCheck(func(t *rapid.T) {
		keyProperties := []string{"p1", "p2", "p3"}
		runRapidTest(t, keyProperties, makeNestedObjectGenerator(keyProperties...), noopMutator)
	}))

	tt.Run("single key, single level", rapid.MakeCheck(func(t *rapid.T) {
		keyProperty := []string{"p1"}
		runRapidTest(t, keyProperty, makeObjectGenerator(allStringValues, keyProperty...), keyedArrayMutator)
	}))

	tt.Run("single key, nested", rapid.MakeCheck(func(t *rapid.T) {
		keyProperty := []string{"p1"}
		runRapidTest(t, keyProperty, makeNestedObjectGenerator(keyProperty...), keyedArrayMutator)
	}))

	tt.Run("multiple keys, nested", rapid.MakeCheck(func(t *rapid.T) {
		keyProperties := []string{"p1", "p2", "p3"}
		runRapidTest(t, keyProperties, makeNestedObjectGenerator(keyProperties...), keyedArrayMutator)
	}))
}

//// Rapid book-keeping

type changeKind string

const (
	add    changeKind = "add"
	update changeKind = "update"
	same   changeKind = "same"
	del    changeKind = "delete"
)

type objectChanges map[resource.PropertyKey]changeKind

// allSames returns true if all changes are "same", i.e., the object as a whole is unchanged.
func (o objectChanges) allSames() bool {
	for _, operation := range o {
		if operation != same {
			return false
		}
	}
	return true
}

type arrayElementChanges struct {
	update      objectChanges
	add, delete bool
	index       int
}

//// Rapid generators

var allStringValues = rapid.Custom[resource.PropertyValue](func(t *rapid.T) resource.PropertyValue {
	return resource.PropertyValue{V: rapid.String().Draw(t, "any string")}
})

// Since the key properties of keyed arrays have unique id semantics, they must be unique.
// While random strings are unlikely to collide, when there's a test failure rapid minimizes the
// input which produces short repeating strings.
var seen = map[string]struct{}{}
var uniqueStrings = rapid.Custom[string](func(t *rapid.T) string {
	hexValues := rapid.StringOfN(rapid.RuneFrom(nil, unicode.ASCII_Hex_Digit), 8, 8, -1)
	uniqueStrings := hexValues.Filter(func(v string) bool {
		_, ok := seen[v]
		if ok {
			return false
		}
		seen[v] = struct{}{}
		return true
	})
	return uniqueStrings.Draw(t, "S")
})

var uniqueStringValues = rapid.Custom[resource.PropertyValue](func(t *rapid.T) resource.PropertyValue {
	return resource.PropertyValue{V: uniqueStrings.Draw(t, "V")}
})

// makeArrayGenerator generates an array whose values are generated via the given `valueGenerator`.
// Note: if the result is used as a keyed array, valueGenerator must generate objects.
func makeArrayGenerator(valueGenerator *rapid.Generator[resource.PropertyValue]) *rapid.Generator[resource.PropertyValue] {
	return rapid.Custom[resource.PropertyValue](func(t *rapid.T) resource.PropertyValue {
		result := []resource.PropertyValue{}

		len := rapid.IntRange(0, 8).Draw(t, "len")
		for i := 0; i < len; i++ {
			result = append(result, valueGenerator.Draw(t, "array value"))
		}
		return resource.NewArrayProperty(result)
	})
}

func makeObjectGenerator(valueGenerator *rapid.Generator[resource.PropertyValue], keys ...string) *rapid.Generator[resource.PropertyValue] {
	return rapid.Custom[resource.PropertyValue](func(t *rapid.T) resource.PropertyValue {
		result := resource.PropertyMap{}

		numOtherProperties := rapid.IntRange(0, 4).Draw(t, "numProperties")
		otherProperties := make([]string, numOtherProperties)
		for i := 0; i < numOtherProperties; i++ {
			otherProperties[i] = uniqueStrings.Draw(t, "other property")
		}

		for _, key := range keys {
			result[resource.PropertyKey(key)] = uniqueStringValues.Draw(t, "val")
		}
		for _, prop := range otherProperties {
			result[resource.PropertyKey(prop)] = valueGenerator.Draw(t, "val")
		}
		// add some divergent properties as well, objects don't need to be totally homogenous
		for i := 0; i < rapid.IntRange(0, 4).Draw(t, "numProperties"); i++ {
			prop := uniqueStrings.Draw(t, "key")
			result[resource.PropertyKey(prop)] = valueGenerator.Draw(t, "val")
		}

		return resource.PropertyValue{V: result}
	})
}

func makeNestedObjectGenerator(keys ...string) *rapid.Generator[resource.PropertyValue] {
	objects := makeObjectGenerator(allStringValues, keys...)
	return makeObjectGenerator(objects, keys...)
}

//// Rapid mutators of PropertyValues to create changes between old and new

// mutateObject makes changes to `new` and records them in the return value.
// For each property of the object, it draws from the four operations
// [update, add, delete, same] and applies the operation.
// Key (identifier) properties in `keysToPreserve` are not mutated.
func mutateObject(t *rapid.T, old, new resource.PropertyMap, keysToPreserve []string) objectChanges {
	record := objectChanges{}
	operations := rapid.SampledFrom([]changeKind{add, update, same, del})

	isKey := func(s string) bool {
		for _, key := range keysToPreserve {
			if key == s {
				return true
			}
		}
		return false
	}

	for k := range old {
		// keys cannot be changed
		if isKey(string(k)) {
			continue
		}

		operation := operations.Draw(t, "object operation")
		switch operation {
		case add:
			newKey := resource.PropertyKey(uniqueStringValues.Draw(t, "key").StringValue())
			new[newKey] = resource.PropertyValue{V: "added"}
			record[newKey] = operation
		case update:
			new[k] = resource.PropertyValue{V: "updated"}
			record[k] = operation
		case del:
			delete(new, k)
			record[k] = operation
		}
	}

	return record
}

// mutator makes random changes to `old` and records them in the return value.
// Implementations should not mutate `old` but make copies.
type mutator func(t *rapid.T, old []resource.PropertyValue, keyProperties []string) ([]arrayElementChanges, []resource.PropertyValue)

var noopMutator mutator = func(t *rapid.T, old []resource.PropertyValue, keyProperties []string) ([]arrayElementChanges, []resource.PropertyValue) {
	return []arrayElementChanges{}, old
}

// keyedArrayMutator makes random changes to `new` and records them in the return value. Changes can
// be addition or removal of array elements, or random changes to existing elements.
var keyedArrayMutator mutator = func(t *rapid.T, old []resource.PropertyValue, keyProperties []string) ([]arrayElementChanges, []resource.PropertyValue) {
	new := make([]resource.PropertyValue, 0, len(old))
	changes := []arrayElementChanges{}
	objects := makeObjectGenerator(allStringValues, keyProperties...)

	// Note: in the loop below, for "add" operations a new object is added and the existing one at
	// the current index is appended to `new` verbatim. Therefore, "add" also implies "same" and we
	// don't need to explicitly draw "same".
	operationsBiasedToUpdate := rapid.SampledFrom([]changeKind{add, update, update, del})

	for i, m := range old {
		kind := operationsBiasedToUpdate.Draw(t, "array operation")
		switch kind {
		case add:
			new = append(new, objects.Draw(t, "new element"))
			new = append(new, m)
			changes = append(changes, arrayElementChanges{add: true, index: len(new) - 2})
		case update:
			copy := deepcopy.Copy(m).(resource.PropertyValue)
			objectChanges := mutateObject(t, m.ObjectValue(), copy.ObjectValue(), keyProperties)
			new = append(new, copy)
			if len(objectChanges) > 0 {
				changes = append(changes, arrayElementChanges{update: objectChanges, index: len(new) - 1})
			}
		case del:
			changes = append(changes, arrayElementChanges{delete: true, index: i})
		case same:
			new = append(new, m)
		}
	}
	return changes, new
}

// compareDiffWithRecordedChanges asserts that a diff matches the previously recorded changes to `old`.
func compareDiffWithRecordedChanges(
	t *rapid.T,
	diff *resource.ValueDiff,
	changes []arrayElementChanges,
	old, new []resource.PropertyValue) {

	// It's possible that rapid rolls a no-op change.
	if diff == nil {
		assert.Equal(t, old, new, "no diff for different objects")
		return
	}
	// The flip side of the contract is that if there were no changes, the diff must be nil.
	if len(changes) == 0 {
		assert.Nil(t, diff)
		return
	}

	assert.Nil(t, diff.Object)
	require.NotNil(t, diff.Array)

	for _, elementChanges := range changes {
		if elementChanges.add {
			assert.Equal(t, diff.Array.Adds[elementChanges.index], new[elementChanges.index])
		} else if elementChanges.delete {
			assert.Equal(t, diff.Array.Deletes[elementChanges.index], old[elementChanges.index])
		} else if elementChanges.update != nil {
			assertObjectChange(t, diff, elementChanges.update, elementChanges.index)
		}
	}
}

func assertObjectChange(t *rapid.T, diff *resource.ValueDiff, objectChanges objectChanges, i int) {
	for key, operation := range objectChanges {
		// If none of the properties changed, the object is in diff->sames.
		if operation == same && objectChanges.allSames() {
			assert.Contains(t, diff.Array.Sames, i, "same of key %s, entry %d", key, i)
			continue
		}

		require.Contains(t, diff.Array.Updates, i, "updates for key %s, entry %d", key, i)
		diffObj := diff.Array.Updates[i].Object

		failureMsgAndArgs := []any{"%s of key %s, entry %d, updates %+v", operation, key, i, diffObj}
		require.NotNil(t, diffObj, failureMsgAndArgs...)
		switch operation {
		case "update":
			assert.Contains(t, diffObj.Updates, key, failureMsgAndArgs...)
		case "delete":
			assert.Contains(t, diffObj.Deletes, key, failureMsgAndArgs...)
		case "add":
			assert.Contains(t, diffObj.Adds, key, failureMsgAndArgs...)
		case "same":
			assert.Contains(t, diffObj.Sames, key, failureMsgAndArgs...)
		}
	}
}

// printDebugJson pretty-prints the arguments in JSON. Can be useful for investigating a test failure.
func printDebugJson(diff *resource.ValueDiff, changes []arrayElementChanges, old, new []resource.PropertyValue) {
	jsonOld, _ := json.MarshalIndent(old, "", "    ")
	fmt.Println("OLD", string(jsonOld))
	fmt.Println("UPDATES", changes)
	jsonNew, _ := json.MarshalIndent(new, "", "    ")
	fmt.Println("NEW", string(jsonNew))
	jsonDiff, _ := json.MarshalIndent(diff, "", "    ")
	fmt.Println("DIFF", string(jsonDiff))
}
