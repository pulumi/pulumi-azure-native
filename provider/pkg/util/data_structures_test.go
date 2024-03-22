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
