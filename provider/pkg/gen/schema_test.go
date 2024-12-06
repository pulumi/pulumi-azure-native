// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeAliasFormatting(t *testing.T) {
	generator := packageGenerator{
		pkg:        &pschema.PackageSpec{Name: "azure-native"},
		sdkVersion: "v20220222",
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
	shouldInclude   func(provider string, version *openapi.ApiVersion, typeName, token string) bool
	getDeprecations func(token string) (ResourceDeprecation, bool)
	getAllVersions  func(provider, resource string) []openapi.ApiVersion
}

func (v versioningStub) ShouldInclude(provider string, version *openapi.ApiVersion, typeName, token string) bool {
	if v.shouldInclude != nil {
		return v.shouldInclude(provider, version, typeName, token)
	}
	return true
}

func (v versioningStub) GetDeprecation(token string) (ResourceDeprecation, bool) {
	if v.getDeprecations != nil {
		return v.getDeprecations(token)
	}
	return ResourceDeprecation{}, false
}

func (v versioningStub) GetAllVersions(provider, resource string) []openapi.ApiVersion {
	if v.getAllVersions != nil {
		return v.getAllVersions(provider, resource)
	}
	return []openapi.ApiVersion{}
}

func TestAliases(t *testing.T) {
	generator := packageGenerator{
		pkg:          &pschema.PackageSpec{Name: "azure-native"},
		sdkVersion:   "v20220222",
		versioning:   versioningStub{},
		provider:     "Insights",
		majorVersion: 2,
	}

	resource := &resourceVariant{
		ResourceSpec: &openapi.ResourceSpec{
			CompatibleVersions: []openapi.SdkVersion{"v20210111"},
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

func TestNormalizePattern(t *testing.T) {
	openapiParam := func(pattern string, maxLength int64) *openapi.Parameter {
		return &openapi.Parameter{
			Parameter: &spec.Parameter{
				CommonValidations: spec.CommonValidations{
					Pattern:   pattern,
					MaxLength: &maxLength,
				},
				ParamProps: spec.ParamProps{
					Name: "dashboardName",
				},
			},
		}
	}

	t.Run("not a dashboard name", func(t *testing.T) {
		p := openapiParam("^[a-zA-Z0-9-]{3,24}$", 90)
		p.Name = "foo"
		assert.Equal(t, "^[a-zA-Z0-9-]{3,24}$", normalizeParamPattern(p))
	})

	t.Run("empty", func(t *testing.T) {
		p := openapiParam("", 0)
		assert.Equal(t, "", normalizeParamPattern(p))
	})

	t.Run("empty with length", func(t *testing.T) {
		p := openapiParam("", 90)
		assert.Equal(t, "", normalizeParamPattern(p))
	})

	t.Run("pattern without length", func(t *testing.T) {
		p := openapiParam("^[a-zA-Z0-9-]{3,24}$", 0)
		assert.Equal(t, "^[a-zA-Z0-9-]{3,24}$", normalizeParamPattern(p))
	})

	t.Run("pattern with correct length", func(t *testing.T) {
		p := openapiParam("^[a-zA-Z0-9-]{3,24}$", 24)
		assert.Equal(t, "^[a-zA-Z0-9-]{3,24}$", normalizeParamPattern(p))
	})

	t.Run("pattern with shorter length", func(t *testing.T) {
		p := openapiParam("^[a-zA-Z0-9-]{3,24}$", 23)
		assert.Equal(t, "^[a-zA-Z0-9-]{3,24}$", normalizeParamPattern(p))
	})

	t.Run("pattern with longer length", func(t *testing.T) {
		p := openapiParam("^[a-zA-Z0-9-]{3,24}$", 60)
		assert.Equal(t, "^[a-zA-Z0-9-]{3,60}$", normalizeParamPattern(p))
	})
}

func TestPropertyDescriptions(t *testing.T) {
	t.Run("no description", func(t *testing.T) {
		s := &spec.Schema{}
		desc, err := getPropertyDescription(s, nil, false)
		require.NoError(t, err)
		assert.Equal(t, "", desc)
	})

	t.Run("simple description", func(t *testing.T) {
		s := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "A simple description",
			},
		}
		desc, err := getPropertyDescription(s, nil, false)
		require.NoError(t, err)
		assert.Equal(t, "A simple description", desc)
	})

	t.Run("line breaks in description are preserved", func(t *testing.T) {
		s := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "A simple \ndescription",
			},
		}
		desc, err := getPropertyDescription(s, nil, false)
		require.NoError(t, err)
		assert.Equal(t, "A simple \ndescription", desc)
	})

	t.Run("description with weird whitespace", func(t *testing.T) {
		s := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This \u2028 is not \u2029 allowed",
			},
		}
		desc, err := getPropertyDescription(s, nil, false)
		require.NoError(t, err)
		assert.Equal(t, "This  is not  allowed", desc)
	})
}
