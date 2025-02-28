// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ensure our VersionMetadata type implements the gen.Versioning interface
// The compiler will raise an error here if the interface isn't implemented
var _ Versioning = (*versioningStub)(nil)

type versioningStub struct {
	shouldInclude      func(moduleName openapi.ModuleName, version *openapi.ApiVersion, typeName, token string) bool
	getDeprecations    func(token string) (ResourceDeprecation, bool)
	getAllVersions     func(moduleName openapi.ModuleName, resource string) []openapi.ApiVersion
	previousTokenPaths map[string]string
}

func (v versioningStub) ShouldInclude(moduleName openapi.ModuleName, version *openapi.ApiVersion, typeName, token string) bool {
	if v.shouldInclude != nil {
		return v.shouldInclude(moduleName, version, typeName, token)
	}
	return true
}

func (v versioningStub) GetDeprecation(token string) (ResourceDeprecation, bool) {
	if v.getDeprecations != nil {
		return v.getDeprecations(token)
	}
	return ResourceDeprecation{}, false
}

func (v versioningStub) GetAllVersions(moduleName openapi.ModuleName, resource string) []openapi.ApiVersion {
	if v.getAllVersions != nil {
		return v.getAllVersions(moduleName, resource)
	}
	return []openapi.ApiVersion{}
}

func (v versioningStub) GetPreviousCompatibleTokensLookup() (*CompatibleTokensLookup, error) {
	return NewCompatibleTokensLookup(v.previousTokenPaths)
}

func TestAliases(t *testing.T) {
	// Wrap the generation of type aliases in a function to make it easier to test
	generateTypeAliases := func(moduleName, typeName string, sdkVersion openapi.SdkVersion, oldTokens []string, typeNameAliases []string, versions []openapi.ApiVersion) []string {
		path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Foo/resources/{name}"
		previousTokenPaths := map[string]string{}
		for _, token := range oldTokens {
			previousTokenPaths[token] = path
		}
		previousCompatibleTokensLookup, err := NewCompatibleTokensLookup(previousTokenPaths)
		require.NoError(t, err, "failed to create previous compatible tokens lookup")
		generator := packageGenerator{
			pkg:                            &pschema.PackageSpec{Name: "azure-native"},
			sdkVersion:                     sdkVersion,
			versioning:                     versioningStub{},
			moduleName:                     openapi.ModuleName(moduleName),
			majorVersion:                   2,
			previousCompatibleTokensLookup: previousCompatibleTokensLookup,
		}

		resource := &resourceVariant{
			ResourceSpec: &openapi.ResourceSpec{
				Path:               path,
				CompatibleVersions: versions,
			},
			typeName: typeName,
		}

		resourceTok := generateTok(openapi.ModuleName(moduleName), typeName, sdkVersion)

		aliasSpecs := generator.generateAliases(resourceTok, resource, typeNameAliases...)
		typeAliases := []string{}
		for _, alias := range aliasSpecs {
			typeAliases = append(typeAliases, *alias.Type)
		}
		return typeAliases
	}

	t.Run("compatible version", func(t *testing.T) {
		actual := generateTypeAliases("Insights", "PrivateLinkForAzureAd", "v20220222", nil, nil, []openapi.ApiVersion{"2020-01-10", "2021-01-11"})
		expected := []string{
			"azure-native:insights/v20200110:PrivateLinkForAzureAd",
			"azure-native:insights/v20210111:PrivateLinkForAzureAd",
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("type alias", func(t *testing.T) {
		actual := generateTypeAliases("Insights", "PrivateLinkForAzureAd", "v20220222", nil, []string{"privateLinkForAzureAd"}, []openapi.ApiVersion{})
		expected := []string{
			"azure-native:insights/v20220222:privateLinkForAzureAd",
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("compatible version & type alias", func(t *testing.T) {
		actual := generateTypeAliases("Insights", "PrivateLinkForAzureAd", "v20220222", nil, []string{"privateLinkForAzureAd"}, []openapi.ApiVersion{"2020-01-10", "2021-01-11"})
		expected := []string{
			"azure-native:insights/v20200110:privateLinkForAzureAd",
			"azure-native:insights/v20200110:PrivateLinkForAzureAd",
			"azure-native:insights/v20210111:privateLinkForAzureAd",
			"azure-native:insights/v20210111:PrivateLinkForAzureAd",
			"azure-native:insights/v20220222:privateLinkForAzureAd",
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("previous module", func(t *testing.T) {
		actual := generateTypeAliases("Monitor", "PrivateLinkForAzureAd", "v20220222", []string{"azure-native:insights/v20220222:PrivateLinkForAzureAd"}, nil, []openapi.ApiVersion{})
		expected := []string{
			"azure-native:insights/v20220222:PrivateLinkForAzureAd",
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("previous module & compatible version", func(t *testing.T) {
		actual := generateTypeAliases("Monitor", "PrivateLinkForAzureAd", "v20220222", []string{"azure-native:insights/v20220222:PrivateLinkForAzureAd"}, nil, []openapi.ApiVersion{"2020-01-10", "2021-01-11"})
		expected := []string{
			"azure-native:monitor/v20200110:PrivateLinkForAzureAd",  // change version
			"azure-native:insights/v20220222:PrivateLinkForAzureAd", // previous major version
			"azure-native:monitor/v20210111:PrivateLinkForAzureAd",  // change version
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("previous module, compatible version & type alias", func(t *testing.T) {
		actual := generateTypeAliases("Monitor", "PrivateLinkForAzureAd", "v20220222", []string{"azure-native:insights/v20220222:PrivateLinkForAzureAd"}, []string{"privateLinkForAzureAd"}, []openapi.ApiVersion{"2020-01-10", "2021-01-11"})
		expected := []string{
			"azure-native:monitor/v20200110:PrivateLinkForAzureAd",  // change version
			"azure-native:monitor/v20200110:privateLinkForAzureAd",  // change version & case
			"azure-native:monitor/v20210111:PrivateLinkForAzureAd",  // change version
			"azure-native:monitor/v20210111:privateLinkForAzureAd",  // change version & case
			"azure-native:monitor/v20220222:privateLinkForAzureAd",  // change case
			"azure-native:insights/v20220222:PrivateLinkForAzureAd", // previous major version
		}
		assert.ElementsMatch(t, expected, actual)
	})
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

func TestResourceIsSingleton(t *testing.T) {
	t.Run("singleton", func(t *testing.T) {
		res := &resourceVariant{
			ResourceSpec: &openapi.ResourceSpec{
				Path:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}",
				PathItem: &spec.PathItem{},
			},
		}
		assert.True(t, isSingleton(res))
	})

	t.Run("implicit singleton", func(t *testing.T) {
		res := &resourceVariant{
			ResourceSpec: &openapi.ResourceSpec{
				Path:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Foo/bar",
				PathItem: &spec.PathItem{},
			},
		}
		assert.True(t, isSingleton(res))
	})

	t.Run("not a singleton", func(t *testing.T) {
		res := &resourceVariant{
			ResourceSpec: &openapi.ResourceSpec{
				Path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Foo/bar",
				PathItem: &spec.PathItem{
					PathItemProps: spec.PathItemProps{
						Delete: &spec.Operation{},
					},
				}},
		}
		assert.False(t, isSingleton(res))
	})
}

func TestGoModuleName(t *testing.T) {
	t.Run("explicit version", func(t *testing.T) {
		assert.Equal(t, "github.com/pulumi/pulumi-azure-native-sdk/network/v2/v20220222", goModuleName("Network", "v20220222"))
	})

	t.Run("default version", func(t *testing.T) {
		assert.Equal(t, "github.com/pulumi/pulumi-azure-native-sdk/network/v2", goModuleName("Network", ""))
	})
}

func TestResourcePathTracker(t *testing.T) {
	t.Run("no conflicts, one module", func(t *testing.T) {
		tracker := newResourcesPathConflictsTracker()
		tracker.addPathConflictsForModule("compute", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"VirtualMachine": {"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachines/{}": {openapi.ApiVersion("2022-02-22")}},
		})
		assert.False(t, tracker.hasConflicts())
	})

	t.Run("conflicts, one module", func(t *testing.T) {
		tracker := newResourcesPathConflictsTracker()
		tracker.addPathConflictsForModule("compute", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"VirtualMachine": {
				"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachines/{}":    {openapi.ApiVersion("2022-02-22")},
				"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachinesFoo/{}": {openapi.ApiVersion("2024-04-22")},
			},
		})
		assert.True(t, tracker.hasConflicts())
	})

	t.Run("no conflicts, multiple modules", func(t *testing.T) {
		tracker := newResourcesPathConflictsTracker()
		tracker.addPathConflictsForModule("compute", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"VirtualMachine": {"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachines/{}": {openapi.ApiVersion("2022-02-22")}},
		})
		tracker.addPathConflictsForModule("storage", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"StorageAccount": {"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Storage/storageAccounts/{}": {openapi.ApiVersion("2022-02-22")}},
		})
		assert.False(t, tracker.hasConflicts())
	})

	t.Run("conflicts, multiple modules", func(t *testing.T) {
		tracker := newResourcesPathConflictsTracker()
		tracker.addPathConflictsForModule("storage", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"StorageAccount": {"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Storage/storageAccounts/{}": {openapi.ApiVersion("2022-02-22")}},
		})
		tracker.addPathConflictsForModule("compute", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"VirtualMachine": {
				"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachines/{}":    {openapi.ApiVersion("2022-02-22")},
				"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Compute/virtualMachinesFoo/{}": {openapi.ApiVersion("2024-04-22")},
			},
		})
		tracker.addPathConflictsForModule("migrate", map[openapi.ResourceName]map[string][]openapi.ApiVersion{
			"AssessmentProject": {"/subscriptions/{}/resourceGroups/{}/providers/Microsoft.Migrate/assessmentProjects/{}": {openapi.ApiVersion("2022-02-22")}},
		})
		assert.True(t, tracker.hasConflicts())
	})
}

func TestParamIsSetByProvider(t *testing.T) {
	createParam := func(name, in string) *openapi.Parameter {
		return &openapi.Parameter{
			Parameter: &spec.Parameter{
				ParamProps: spec.ParamProps{
					Name: name,
					In:   in,
				},
			},
		}
	}

	t.Run("not api-version param", func(t *testing.T) {
		m := &moduleGenerator{
			moduleName:   "SomeModule",
			resourceName: "SomeResource",
		}
		param := createParam("some-other-param", "query")
		assert.False(t, m.paramIsSetByProvider(param, nil))
	})

	t.Run("not query param", func(t *testing.T) {
		m := &moduleGenerator{
			moduleName:   "SomeModule",
			resourceName: "SomeResource",
		}
		param := createParam("some-other-param", "path")
		assert.False(t, m.paramIsSetByProvider(param, nil))
	})

	t.Run("Resources.Resource special case", func(t *testing.T) {
		m := &moduleGenerator{
			moduleName:   "Resources",
			resourceName: "Resource",
		}
		param := createParam("api-version", "query")
		assert.False(t, m.paramIsSetByProvider(param, nil))
	})

	t.Run("Authorization.ManagementLockAtResourceLevel special case", func(t *testing.T) {
		m := &moduleGenerator{
			moduleName:   "Authorization",
			resourceName: "ManagementLockAtResourceLevel",
		}
		param := createParam("api-version", "query")
		assert.False(t, m.paramIsSetByProvider(param, nil))
	})

	t.Run("standard api-version query param", func(t *testing.T) {
		m := &moduleGenerator{
			moduleName:   "SomeModule",
			resourceName: "SomeResource",
		}
		param := createParam("api-version", "query")
		assert.True(t, m.paramIsSetByProvider(param, nil))
	})
}

func TestAddApiVersionOverride(t *testing.T) {
	t.Run("major version < 3", func(t *testing.T) {
		g := &packageGenerator{
			majorVersion: 2,
		}
		resourceRequest := &parameterBag{specs: map[string]pschema.PropertySpec{}}
		g.addApiVersionOverride(resourceRequest, "2022-02-22")
		assert.Equal(t, 0, len(resourceRequest.parameters))
	})

	t.Run("major version >= 3", func(t *testing.T) {
		g := &packageGenerator{majorVersion: 3}
		resourceRequest := &parameterBag{
			specs:      map[string]pschema.PropertySpec{},
			parameters: []resources.AzureAPIParameter{},
		}
		g.addApiVersionOverride(resourceRequest, "2022-02-22")
		assert.Equal(t, 1, len(resourceRequest.specs))
		assert.Contains(t, resourceRequest.specs, apiVersionOverride)
		assert.Len(t, resourceRequest.parameters, 1)
		assert.Equal(t, apiVersionOverride, resourceRequest.parameters[0].Name)
	})

	t.Run("major version >= 3, api version already set", func(t *testing.T) {
		g := &packageGenerator{
			majorVersion: 3,
		}
		resourceRequest := &parameterBag{
			specs: map[string]pschema.PropertySpec{
				"apiVersion": {
					TypeSpec: pschema.TypeSpec{Type: "string"},
				},
			},
			parameters: []resources.AzureAPIParameter{},
		}
		g.addApiVersionOverride(resourceRequest, "2022-02-22")
		assert.Equal(t, 1, len(resourceRequest.specs))
	})
}
