// Copyright 2024, Pulumi Corporation.  All rights reserved.

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInnerMap(t *testing.T) {
	t.Run("no keys", func(t *testing.T) {
		m := map[string]any{
			"a": "b",
		}
		_, ok := GetInnerMap(m)
		assert.False(t, ok)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]any{}
		_, ok := GetInnerMap(m, "a", "b")
		assert.False(t, ok)
	})

	t.Run("key has no map value", func(t *testing.T) {
		m := map[string]any{
			"a": 1,
		}
		_, ok := GetInnerMap(m, "a")
		assert.False(t, ok)
	})

	t.Run("single key", func(t *testing.T) {
		m := map[string]any{
			"a": map[string]any{"b": 1},
		}
		inner, ok := GetInnerMap(m, "a")
		assert.True(t, ok)
		assert.Equal(t, map[string]any{"b": 1}, inner)
	})

	t.Run("multiple keys", func(t *testing.T) {
		m := map[string]any{
			"a": map[string]any{"b": map[string]any{"c": 1}},
		}
		inner, ok := GetInnerMap(m, "a", "b")
		assert.True(t, ok)
		assert.Equal(t, map[string]any{"c": 1}, inner)
	})

	t.Run("multiple keys, one has no map value", func(t *testing.T) {
		m := map[string]any{
			"a": map[string]any{"b": 1},
		}
		_, ok := GetInnerMap(m, "a", "b", "c")
		assert.False(t, ok)
	})
}

func TestRemoveFromSlice(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		s := []int{}
		s = RemoveFromSlice(s, 1)
		assert.Equal(t, []int{}, s)
	})

	t.Run("value not found", func(t *testing.T) {
		s := []int{1, 2, 3}
		s = RemoveFromSlice(s, 4)
		assert.Equal(t, []int{1, 2, 3}, s)
	})

	t.Run("value found", func(t *testing.T) {
		s := []int{1, 2, 3}
		s = RemoveFromSlice(s, 2)
		assert.Equal(t, []int{1, 3}, s)
	})
}
