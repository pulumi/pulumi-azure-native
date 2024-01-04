package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
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
