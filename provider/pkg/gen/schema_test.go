// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestTypeAliasFormatting(t *testing.T) {
	generator := packageGenerator{
		pkg:        &pschema.PackageSpec{Name: "azure-native"},
		apiVersion: "v20220222",
		provider:   "Compute",
	}

	actual := generator.makeTypeAlias("VirtualMachine", "v18851225")
	assert.NotNil(t, actual)
	assert.Equal(t, "azure-native:compute/v18851225:VirtualMachine", *actual.Type)

	actual = generator.makeTypeAlias("VirtualMachine", "")
	assert.NotNil(t, actual)
	assert.Equal(t, "azure-native:compute:VirtualMachine", *actual.Type)
}

// Ensure our VersionMetadata type implements the gen.Versioning interface
// The compiler will raise an error here if the interface isn't implemented
var _ Versioning = (*versioningStub)(nil)

type versioningStub struct {
}

func (v versioningStub) ShouldInclude(provider string, version string, typeName, token string) bool {
	return true
}

func (v versioningStub) GetDeprecation(token string) (ResourceDeprecation, bool) {
	return ResourceDeprecation{}, false
}

func (v versioningStub) GetAllVersions(provider, resource string) []string {
	return []string{}
}

func TestAliases(t *testing.T) {
	generator := packageGenerator{
		pkg:        &pschema.PackageSpec{Name: "azure-native"},
		apiVersion: "v20220222",
		versioning: versioningStub{},
		provider:   "Insights",
	}

	resource := &resourceVariant{
		ResourceSpec: &openapi.ResourceSpec{
			CompatibleVersions: []string{"v20210111"},
		},
		typeName: "PrivateLinkForAzureAd",
	}

	aliases := generator.generateAliases(resource, "privateLinkForAzureAd")
	actual := codegen.NewStringSet()
	for _, alias := range aliases {
		actual.Add(*alias.Type)
	}
	expected := codegen.NewStringSet(
		"azure-native:insights/v20210111:privateLinkForAzureAd",
		"azure-native:insights/v20210111:PrivateLinkForAzureAd",
		"azure-native:insights/v20220222:privateLinkForAzureAd",
	)
	assert.Equal(t, expected, actual)
}

func TestFindNestedResources(t *testing.T) {
	parent := &openapi.ResourceSpec{
		Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}",
	}

	otherResources := map[string]*openapi.ResourceSpec{
		"VirtualNetwork": {
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}",
		},
		"VirtualNetworks": {
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks",
		},
		"Subnet": {
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}",
		},
		"SubnetPrivateEndpoint": {
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}",
		},
		"VirtualNetworkPeering": {
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}",
		},
	}

	actual := findNestedResources(parent, otherResources)

	expected := []*openapi.ResourceSpec{
		{
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}",
		},
		{
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}",
		},
		{
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestFindResourcesBodyRefs(t *testing.T) {
	nestedResources := []*openapi.ResourceSpec{
		{
			Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}",
			PathItem: &spec.PathItem{
				PathItemProps: spec.PathItemProps{
					Put: &spec.Operation{
						OperationProps: spec.OperationProps{
							Parameters: []spec.Parameter{
								{
									ParamProps: spec.ParamProps{
										In: "body",
										Schema: &spec.Schema{
											SchemaProps: spec.SchemaProps{
												Ref: spec.MustCreateRef("#/definitions/Subnet"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	actual := findResourcesBodyRefs(nestedResources)
	expected := []string{"#/definitions/Subnet"}
	assert.Equal(t, expected, actual)
}

func TestMergeSchemaOverlays(t *testing.T) {
	strProp := pschema.PropertySpec{
		TypeSpec: pschema.TypeSpec{
			Type: "string",
		},
	}
	intProp := pschema.PropertySpec{
		TypeSpec: pschema.TypeSpec{
			Type: "integer",
		},
	}

	t.Run("addition", func(t *testing.T) {
		res := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
			},
		}

		overlay := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"bar": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"barInput": strProp,
			},
		}

		expected := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": strProp,
					"bar": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
				"barInput": strProp,
			},
		}

		assert.Equal(t, expected, mergeSchemaOverlay(res, overlay))
	})

	t.Run("overwrite", func(t *testing.T) {
		res := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
			},
		}

		overlay := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": intProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": intProp,
			},
		}

		expected := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": intProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": intProp,
			},
		}

		assert.Equal(t, expected, mergeSchemaOverlay(res, overlay))
	})

	t.Run("combined", func(t *testing.T) {
		res := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
			},
		}

		overlay := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": intProp,
					"bar": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"barInput":     strProp,
				"anotherInput": strProp,
			},
		}

		expected := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": intProp,
					"bar": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput":     strProp,
				"barInput":     strProp,
				"anotherInput": strProp,
			},
		}

		assert.Equal(t, expected, mergeSchemaOverlay(res, overlay))
	})

	t.Run("no-op", func(t *testing.T) {
		res := pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"foo": strProp,
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
			},
		}

		overlay := pschema.ResourceSpec{
			InputProperties: map[string]pschema.PropertySpec{
				"fooInput": strProp,
			},
		}

		assert.Equal(t, res, mergeSchemaOverlay(res, overlay))
	})
}

func TestMergeMetaOverlays(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		meta := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "path",
					Name:     "foo",
				},
			},
		}

		overlay := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "path",
					Name:     "bar",
				},
			},
		}

		merged := mergeMetaOverlay(meta, overlay)
		assert.Len(t, merged.PutParameters, 2)
		assert.Contains(t, merged.PutParameters, meta.PutParameters[0])
		assert.Contains(t, merged.PutParameters, overlay.PutParameters[0])
	})

	t.Run("addition in body", func(t *testing.T) {
		meta := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Name:     "properties",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"foo": {
								Type: "string",
							},
						},
					},
				},
			},
		}

		overlay := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Name:     "properties",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"bar": {
								Type: "string",
							},
						},
					},
				},
			},
		}

		merged := mergeMetaOverlay(meta, overlay)
		assert.Len(t, merged.PutParameters, 1)
		assert.Len(t, merged.PutParameters[0].Body.Properties, 2)
		assert.Contains(t, merged.PutParameters[0].Body.Properties, "foo")
		assert.Contains(t, merged.PutParameters[0].Body.Properties, "bar")
	})

	t.Run("no-op", func(t *testing.T) {
		meta := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{},
		}

		overlay := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{},
		}

		expected := resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{},
		}

		assert.Equal(t, expected, mergeMetaOverlay(meta, overlay))
	})
}
