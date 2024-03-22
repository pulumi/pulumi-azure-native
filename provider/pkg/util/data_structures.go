package util

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
