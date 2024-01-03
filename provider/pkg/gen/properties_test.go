package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestPropChangeForcesRecreate(t *testing.T) {
	t.Run("no extensions", func(t *testing.T) {
		schema := &openapi.Schema{
			Schema: &spec.Schema{},
		}
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.False(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})

	makeSchema := func(mutability []string) *openapi.Schema {
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
		}
	}

	t.Run("create only", func(t *testing.T) {
		schema := makeSchema([]string{extensionMutabilityCreate})
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.True(t, forcesRecreate)
	})
	t.Run("create and update", func(t *testing.T) {
		schema := makeSchema([]string{extensionMutabilityCreate, extensionMutabilityUpdate})
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
	t.Run("create and read", func(t *testing.T) {
		schema := makeSchema([]string{extensionMutabilityCreate, extensionMutabilityRead})
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.True(t, forcesRecreate)
	})
	t.Run("read only", func(t *testing.T) {
		schema := makeSchema([]string{extensionMutabilityRead})
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
	t.Run("all", func(t *testing.T) {
		schema := makeSchema([]string{extensionMutabilityCreate, extensionMutabilityUpdate, extensionMutabilityRead})
		hasMutabilityInfo, forcesRecreate := propChangeForcesRecreate(schema)
		assert.True(t, hasMutabilityInfo)
		assert.False(t, forcesRecreate)
	})
}
