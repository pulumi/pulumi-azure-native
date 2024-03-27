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
		prov: "foo",
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
}

func TestNonObjectInvokeResponses(t *testing.T) {
	m := moduleGenerator{
		prov: "foo",
	}

	resolvedSchema := &openapi.Schema{
		Schema: &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"string"},
			},
		},
	}

	t.Run("string type, response properties requested", func(t *testing.T) {
		props, err := m.genProperties(resolvedSchema, true, false, true)
		require.NoError(t, err)

		require.Len(t, props.specs, 1)
		assert.Contains(t, props.specs, resources.SingleValueProperty)

		require.Len(t, props.properties, 1)
		assert.Contains(t, props.properties, resources.SingleValueProperty)
	})

	t.Run("string type, response properties not requested", func(t *testing.T) {
		props, err := m.genProperties(resolvedSchema, true, false, false)
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

		props, err := m.genProperties(schema, true, false, true)
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

		props, err := m.genProperties(schema, true, false, true)
		require.NoError(t, err)

		require.Len(t, props.specs, 1)
		assert.NotContains(t, props.specs, resources.SingleValueProperty)
		require.Len(t, props.properties, 1)
		assert.NotContains(t, props.properties, resources.SingleValueProperty)
	})
}
