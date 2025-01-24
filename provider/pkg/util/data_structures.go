package util

import (
	"cmp"
	"iter"
	"slices"

	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
)

// GetInnerMap returns a map nested inside another. It traverses the list of keys, such that in
// {a: {b: {c: 1}}}, GetInnerMap(m, "a", "b") returns {c: 1}.
// Its purpose is to let callers avoid the repeated `if ..., ok` double check of does the key
// exist, and is the value a map, at each level.
func GetInnerMap(m map[string]any, keys ...string) (map[string]any, bool) {
	cur := m
	for i, key := range keys {
		val, ok := cur[key]
		if !ok {
			return nil, false
		}
		if i == len(keys)-1 {
			if valMap, ok := val.(map[string]any); ok {
				return valMap, true
			}
		}
		cur, ok = val.(map[string]any)
		if !ok {
			return nil, false
		}
	}
	return nil, false
}

// UnsortedKeys returns the keys of a map in an arbitrary order.
func UnsortedKeys[K comparable, V any](m map[K]V) []K {
	keys := slice.Prealloc[K](len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns the keys of a map in sorted order.
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := UnsortedKeys(m)
	slices.Sort(keys)
	return keys
}

// MapOrdered returns a sequence of key-value pairs from a map, ordered by key.
// Example usage:
//
//	for key, value := range util.MapOrdered(m) {
//	    ...
//	}
func MapOrdered[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	keys := SortedKeys(m)
	return func(yield func(K, V) bool) {
		for _, key := range keys {
			if !yield(key, m[key]) {
				return
			}
		}
	}
}
