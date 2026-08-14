package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeSchema(mutability ...string) *openapi.Schema {
	mutabilityInterface := make([]interface{}, len(mutability))
	for i, m := range mutability {
		mutabilityInterface[i] = m
	}
	return &openapi.Schema{
		Schema: &spec.Schema{
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					extensionMutability: mutabilityInterface,
				},
			},
		},
		ReferenceContext: &openapi.ReferenceContext{
			ReferenceName: "foo",
		},
	}
}

func TestPropChangeForcesRecreate(t *testing.T) {
	t.Run("no extensions", func(t *testing.T) {
		schema := &openapi.Schema{
			Schema: &spec.Schema{},
		}
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.False(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})

	t.Run("create only", func(t *testing.T) {
		schema := makeSchema(extensionMutabilityCreate)
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.True(t, forcesRecreate)
	})
	t.Run("create and update", func(t *testing.T) {
		schema := makeSchema(extensionMutabilityCreate, extensionMutabilityUpdate)
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
	t.Run("create and read", func(t *testing.T) {
		schema := makeSchema(extensionMutabilityCreate, extensionMutabilityRead)
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.True(t, forcesRecreate)
	})
	t.Run("read only", func(t *testing.T) {
		schema := makeSchema(extensionMutabilityRead)
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
	t.Run("all", func(t *testing.T) {
		schema := makeSchema(extensionMutabilityCreate, extensionMutabilityUpdate, extensionMutabilityRead)
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
}

func TestForceNew(t *testing.T) {
	m := moduleGenerator{
		moduleName: "foo",
	}

	t.Run("forceNew", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(extensionMutabilityCreate), "prop", false)
		assert.Equal(t, forceNew, forceNewMetadata)
	})
	t.Run("noForceNew, mutability spec", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(extensionMutabilityCreate, extensionMutabilityUpdate), "prop", false)
		assert.Equal(t, noForceNew, forceNewMetadata)
	})
	t.Run("noForceNew, no mutability spec", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(), "prop", false)
		assert.Equal(t, noForceNew, forceNewMetadata)
	})

	t.Run("forceNew, is type", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(extensionMutabilityCreate), "prop", true)
		assert.Equal(t, forceNewSetOnReferencedType, forceNewMetadata)
	})
	t.Run("noForceNew, mutability spec, is type", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(extensionMutabilityCreate, extensionMutabilityUpdate), "prop", true)
		assert.Equal(t, noForceNew, forceNewMetadata)
	})
	t.Run("noForceNew, no mutability spec, is type", func(t *testing.T) {
		forceNewMetadata := m.forceNew(makeSchema(), "prop", true)
		assert.Equal(t, noForceNew, forceNewMetadata)
	})

	elasticSan := moduleGenerator{moduleName: "ElasticSan", resourceName: "ElasticSan"}
	t.Run("noForceNewMap overrides create-only mutability", func(t *testing.T) {
		forceNewMetadata := elasticSan.forceNew(makeSchema(extensionMutabilityCreate, extensionMutabilityRead), "baseSizeTiB", false)
		assert.Equal(t, noForceNew, forceNewMetadata)
		forceNewMetadata = elasticSan.forceNew(makeSchema(extensionMutabilityCreate, extensionMutabilityRead), "extendedCapacitySizeTiB", false)
		assert.Equal(t, noForceNew, forceNewMetadata)
	})
	t.Run("noForceNewMap does not affect other properties", func(t *testing.T) {
		forceNewMetadata := elasticSan.forceNew(makeSchema(extensionMutabilityCreate, extensionMutabilityRead), "sku", false)
		assert.Equal(t, forceNew, forceNewMetadata)
	})
}

func TestCaseInsensitiveDiff(t *testing.T) {
	managedCluster := moduleGenerator{moduleName: "ContainerService", resourceName: "ManagedCluster"}
	otherResource := moduleGenerator{moduleName: "ContainerService", resourceName: "AgentPool"}
	otherModule := moduleGenerator{moduleName: "Storage", resourceName: "StorageAccount"}

	assert.True(t, managedCluster.caseInsensitiveDiff("backendPoolType"))
	assert.False(t, managedCluster.caseInsensitiveDiff("someOtherProperty"))
	assert.False(t, otherResource.caseInsensitiveDiff("backendPoolType"))
	assert.False(t, otherModule.caseInsensitiveDiff("backendPoolType"))

	searchService := moduleGenerator{moduleName: "Search", resourceName: "Service"}
	assert.True(t, searchService.caseInsensitiveDiff("hostingMode"))
	assert.False(t, searchService.caseInsensitiveDiff("someOtherProperty"))

	vault := moduleGenerator{moduleName: "KeyVault", resourceName: "Vault"}
	otherKeyVaultResource := moduleGenerator{moduleName: "KeyVault", resourceName: "ManagedHsm"}
	assert.True(t, vault.caseInsensitiveDiff("name"))
	assert.False(t, vault.caseInsensitiveDiff("someOtherProperty"))
	assert.False(t, otherKeyVaultResource.caseInsensitiveDiff("name"))
}

func TestNoDefault(t *testing.T) {
	webApp := moduleGenerator{moduleName: "Web", resourceName: "WebApp"}
	webAppSlot := moduleGenerator{moduleName: "Web", resourceName: "WebAppSlot"}
	otherResource := moduleGenerator{moduleName: "Web", resourceName: "AppServicePlan"}
	otherModule := moduleGenerator{moduleName: "Storage", resourceName: "StorageAccount"}

	assert.True(t, webApp.noDefault("http20ProxyFlag"))
	assert.True(t, webAppSlot.noDefault("http20ProxyFlag"))
	assert.True(t, webApp.noDefault("reserved"))
	assert.True(t, webAppSlot.noDefault("reserved"))
	// Sibling properties of the same type keep their spec defaults.
	assert.False(t, webApp.noDefault("http20Enabled"))
	assert.False(t, otherResource.noDefault("http20ProxyFlag"))
	assert.False(t, otherModule.noDefault("http20ProxyFlag"))
	assert.False(t, otherResource.noDefault("reserved"))
	assert.False(t, otherModule.noDefault("reserved"))
}

func TestIsReadableOutput(t *testing.T) {
	webApp := moduleGenerator{moduleName: "Web", resourceName: "WebApp"}
	webAppSlot := moduleGenerator{moduleName: "Web", resourceName: "WebAppSlot"}
	plan := moduleGenerator{moduleName: "Web", resourceName: "AppServicePlan"}
	otherModule := moduleGenerator{moduleName: "Storage", resourceName: "StorageAccount"}

	assert.True(t, webApp.isReadableOutput("siteConfig"))
	assert.True(t, webAppSlot.isReadableOutput("siteConfig"))
	assert.False(t, webApp.isReadableOutput("someOtherProperty"))
	assert.False(t, plan.isReadableOutput("siteConfig"))
	assert.False(t, otherModule.isReadableOutput("siteConfig"))
}

func TestWriteOnlyOutputRetained(t *testing.T) {
	// siteConfig as Azure models it on Microsoft.Web/sites since api-version 2024-11-01:
	// settable on create/update, but not returned by a GET of the parent resource.
	writeOnlySiteConfig := func() *openapi.Schema {
		return &openapi.Schema{
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type: []string{"object"},
					Properties: map[string]spec.Schema{
						"siteConfig": {
							SchemaProps: spec.SchemaProps{Type: []string{"string"}},
							VendorExtensible: spec.VendorExtensible{
								Extensions: spec.Extensions{
									extensionMutability: []interface{}{
										extensionMutabilityCreate,
										extensionMutabilityUpdate,
									},
								},
							},
						},
					},
				},
			},
		}
	}
	variant := genPropertiesVariant{isOutput: true}

	t.Run("stripped as write-only by default", func(t *testing.T) {
		m := moduleGenerator{moduleName: "Web", resourceName: "AppServicePlan"}
		props, err := m.genProperties(writeOnlySiteConfig(), variant)
		require.NoError(t, err)
		assert.NotContains(t, props.specs, "siteConfig")
		assert.NotContains(t, props.properties, "siteConfig")
	})

	t.Run("retained for Web/WebApp via readableOutputs override", func(t *testing.T) {
		m := moduleGenerator{moduleName: "Web", resourceName: "WebApp"}
		props, err := m.genProperties(writeOnlySiteConfig(), variant)
		require.NoError(t, err)
		assert.Contains(t, props.specs, "siteConfig")
		assert.Contains(t, props.properties, "siteConfig")
	})
}

func TestNonObjectInvokeResponses(t *testing.T) {
	m := moduleGenerator{
		moduleName: "foo",
	}

	resolvedSchema := &openapi.Schema{
		Schema: &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"string"},
			},
		},
	}

	t.Run("string type, response properties requested", func(t *testing.T) {
		variant := genPropertiesVariant{
			isOutput:   true,
			isType:     false,
			isResponse: true,
		}
		props, err := m.genProperties(resolvedSchema, variant)
		require.NoError(t, err)

		require.Len(t, props.specs, 1)
		assert.Contains(t, props.specs, resources.SingleValueProperty)

		require.Len(t, props.properties, 1)
		assert.Contains(t, props.properties, resources.SingleValueProperty)
	})

	t.Run("string type, response properties not requested", func(t *testing.T) {
		variant := genPropertiesVariant{
			isOutput:   true,
			isType:     false,
			isResponse: false,
		}
		props, err := m.genProperties(resolvedSchema, variant)
		require.NoError(t, err)

		require.Len(t, props.specs, 0)
		require.Len(t, props.properties, 0)
	})

	t.Run("object type, response properties requested", func(t *testing.T) {
		schema := &openapi.Schema{
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type: []string{"object"},
				},
			},
		}

		variant := genPropertiesVariant{
			isOutput:   true,
			isType:     false,
			isResponse: true,
		}
		props, err := m.genProperties(schema, variant)
		require.NoError(t, err)

		require.Len(t, props.specs, 0)
		require.Len(t, props.properties, 0)
	})

	t.Run("string type, response properties requested but there are other properties", func(t *testing.T) {
		schema := &openapi.Schema{
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type: []string{"string"},
					Properties: map[string]spec.Schema{
						"foo": {
							SchemaProps: spec.SchemaProps{},
						},
					},
				},
			},
		}

		variant := genPropertiesVariant{
			isOutput:   true,
			isType:     false,
			isResponse: true,
		}
		props, err := m.genProperties(schema, variant)
		require.NoError(t, err)

		require.Len(t, props.specs, 1)
		assert.NotContains(t, props.specs, resources.SingleValueProperty)
		require.Len(t, props.properties, 1)
		assert.NotContains(t, props.properties, resources.SingleValueProperty)
	})
}

func TestPropertyIntersection(t *testing.T) {
	t.Run("no conflict", func(t *testing.T) {
		outer := propertyBag{properties: map[string]resources.AzureAPIProperty{
			"foo": {},
		}}
		inner := propertyBag{properties: map[string]resources.AzureAPIProperty{
			"bar":  {},
			"foo2": {},
		}}
		assert.Empty(t, outer.propertyIntersection(&inner))
	})

	t.Run("conflict", func(t *testing.T) {
		outer := propertyBag{properties: map[string]resources.AzureAPIProperty{
			"foo":  {},
			"foo2": {},
			"bla":  {},
		}}
		inner := propertyBag{properties: map[string]resources.AzureAPIProperty{
			"foo": {},
			"bar": {},
		}}
		assert.Equal(t, []string{"foo"}, outer.propertyIntersection(&inner))
	})

}
