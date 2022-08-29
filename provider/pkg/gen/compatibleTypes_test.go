// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplexTypeMerge(t *testing.T) {
	t.Run("Empty output", func(t *testing.T) {
		t1 := schema.ComplexTypeSpec{}
		t2 := schema.ComplexTypeSpec{}
		actual, err := mergeTypes(t1, t2, false)
		expected := schema.ComplexTypeSpec{}
		assert.NoError(t, err)
		assert.Equal(t, expected, *actual)
	})
	t.Run("Single Property", func(t *testing.T) {
		t1 := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"a": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		}
		t2 := t1
		actual, err := mergeTypes(t1, t2, false)
		expected := t1
		assert.NoError(t, err)
		assert.Equal(t, expected, *actual)
	})
	t.Run("Diff optional property", func(t *testing.T) {
		t1 := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"a": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		}
		t2 := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"b": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		}
		actual, err := mergeTypes(t1, t2, false)
		expected := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"a": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
					"b": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, *actual)
	})
	t.Run("Diff required property", func(t *testing.T) {
		t1 := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"a": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
				Required: []string{"a"},
			},
		}
		t2 := schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"b": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		}
		_, err := mergeTypes(t1, t2, false)
		assert.Error(t, err)
	})
}
