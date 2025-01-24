package customresources

import (
	"encoding/json"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
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

func TestPrepareDeleteRequest(t *testing.T) {
	// A minimal copy of the actual resource metadata, reduced to what's needed for the test.
	roleManagementPolicyResourceMetadata := &resources.AzureAPIResource{
		APIVersion:   "2024-09-01-preview",
		Path:         "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}",
		UpdateMethod: "PATCH",
		PutParameters: []resources.AzureAPIParameter{
			{
				Name:     "scope",
				Location: "path",
			},
			{
				Name:     "roleManagementPolicyName",
				Location: "path",
			},
			{
				Name:     "parameters",
				Location: "body",
				// IsRequired: true,
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"description": {
							Type:       "string",
							Containers: []string{"properties"},
						},
						"displayName": {
							Type:       "string",
							Containers: []string{"properties"},
						},
						"isOrganizationDefault": {
							Type:       "boolean",
							Containers: []string{"properties"},
						},
						"rules": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								OneOf: []string{"#/types/azure-native:authorization:RoleManagementPolicyExpirationRule"},
							},
							Containers: []string{"properties"},
						},
						"scope": {
							Type:       "string",
							Containers: []string{"properties"},
						},
					},
				},
			},
		},
	}

	sdkShapeConverter := convert.NewSdkShapeConverterFull(nil)
	client := crud.NewResourceCrudClient(&azure.MockAzureClient{}, nil, &sdkShapeConverter, "123", roleManagementPolicyResourceMetadata)

	id := "/subscriptions/123/providers/Microsoft.Authorization/roleManagementPolicies/policy1"
	resourcePath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"

	t.Run("happy path", func(t *testing.T) {
		state := resource.NewPropertyMapFromMap(map[string]any{
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": []map[string]any{
						{"id": "rule1", "isExpirationRequired": false},
						{"id": "rule2"},
					},
				},
			},
		})

		req, err := prepareUpdateRequestForDelete(client, id, resourcePath, state)
		require.NoError(t, err)

		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		expectedJson, err := json.Marshal(map[string]any{
			"properties": map[string]any{
				"rules": []map[string]any{
					{"id": "rule1", "isExpirationRequired": false},
					{"id": "rule2"},
				},
				"scope": "subscriptions/123",
			},
		})
		require.NoError(t, err)

		assert.JSONEq(t, string(expectedJson), string(jsonReq))
	})

	t.Run("no saved state", func(t *testing.T) {
		state := resource.PropertyMap{}
		_, err := prepareUpdateRequestForDelete(client, id, resourcePath, state)
		require.Error(t, err)
	})
}
