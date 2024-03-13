package customresources

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestoreDefaultsForDeletedRules(t *testing.T) {
	t.Run("nothing to restore", func(t *testing.T) {
		// There's one rule, it's in the original state with a different value for isExpirationRequired.
		olds := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": true},
			},
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": []map[string]any{
						{"id": "rule1", "isExpirationRequired": false},
					},
				},
			},
		}

		// One rule is modified and another added, but none is removed.
		news := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": false},
				{"id": "rule2"},
			},
		}

		newsPropertyMap := resource.NewPropertyMapFromMap(news)
		newsPropertyMapCopy := deepcopy.Copy(newsPropertyMap)
		restoreDefaultsForDeletedRules(resource.NewPropertyMapFromMap(olds), newsPropertyMap)
		assert.Equal(t, newsPropertyMap, newsPropertyMapCopy)
	})

	t.Run("one deleted rule to restore", func(t *testing.T) {
		// There's one rule, it's in the original state with a different value for isExpirationRequired.
		olds := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": true},
			},
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": []map[string]any{
						{"id": "rule1", "isExpirationRequired": false},
					},
				},
			},
		}

		// The old rule is removed and another added.
		news := map[string]any{
			"rules": []map[string]any{
				{"id": "rule2"},
			},
		}

		newsPropertyMap := resource.NewPropertyMapFromMap(news)
		newsPropertyMapCopy := deepcopy.Copy(newsPropertyMap).(resource.PropertyMap)
		restoreDefaultsForDeletedRules(resource.NewPropertyMapFromMap(olds), newsPropertyMap)

		assert.NotEqual(t, newsPropertyMap, newsPropertyMapCopy)
		assert.Equal(t,
			[]resource.PropertyValue{
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{"id": "rule1", "isExpirationRequired": false})),
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{"id": "rule2"})),
			},
			newsPropertyMap["rules"].ArrayValue(),
		)
	})

	t.Run("all rules deleted", func(t *testing.T) {
		// There are two rules in the original state, one with a different value for isExpirationRequired.
		origRules := []map[string]any{
			{"id": "rule1", "isExpirationRequired": false},
			{"id": "rule2"},
		}

		olds := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": true},
				{"id": "rule2"},
			},
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": origRules,
				},
			},
		}

		// No rules left. OriginalState is optional, it's read from `olds`.
		news := map[string]any{}

		newsPropertyMap := resource.NewPropertyMapFromMap(news)
		newsPropertyMapCopy := deepcopy.Copy(newsPropertyMap).(resource.PropertyMap)
		restoreDefaultsForDeletedRules(resource.NewPropertyMapFromMap(olds), newsPropertyMap)

		assert.NotEqual(t, newsPropertyMap, newsPropertyMapCopy)
		rules := newsPropertyMap["rules"].ArrayValue()
		require.Len(t, rules, len(origRules))
		assert.Equal(t,
			[]resource.PropertyValue{
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(origRules[0])),
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(origRules[1])),
			},
			newsPropertyMap["rules"].ArrayValue(),
		)
	})
}
