package collections

import (
	"cmp"
	"slices"
)

func OrderedKeys[T cmp.Ordered](m map[T]any) []T {
	keys := make([]T, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
