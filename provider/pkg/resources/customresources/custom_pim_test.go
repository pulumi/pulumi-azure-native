package customresources

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanCreateRoleManagementPolicy(t *testing.T) {
	t.Parallel()

	t.Run("returns nil", func(t *testing.T) {
		t.Parallel()
		r, err := pimRoleManagementPolicy(nil, nil)
		require.NoError(t, err)
		assert.Nil(t, r.CanCreate(context.Background(), ""))
	})
}

func TestCreateRoleManagementPolicy(t *testing.T) {
	t.Parallel()

	roleManagementPolicyResource := resources.AzureAPIResource{}

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		azureClient := azure.MockAzureClient{
			GetResponse: map[string]any{"original": "policy"},
		}
		mockClient := crud.NewResourceCrudClient(&azureClient, nil, nil, "123", &roleManagementPolicyResource)
		roleManagementPolicyClient := &roleManagementPolicyClient{
			client: mockClient,
		}

		out, err := roleManagementPolicyClient.create(context.Background(), "/id", resource.PropertyMap{})
		require.NoError(t, err)
		// we read the original policy
		assert.Equal(t, []string{"/id"}, azureClient.GetIds)
		// we preserved the original policy in state
		assert.Equal(t, out[OriginalStateKey], map[string]any{"original": "policy"})
	})

	t.Run("error on read", func(t *testing.T) {
		t.Parallel()

		azureClient := azure.MockAzureClient{
			GetResponseErr: errors.New("error"),
		}
		mockClient := crud.NewResourceCrudClient(&azureClient, nil, nil, "123", &roleManagementPolicyResource)
		roleManagementPolicyClient := &roleManagementPolicyClient{
			client: mockClient,
		}
		out, err := roleManagementPolicyClient.create(context.Background(), "/id", resource.PropertyMap{})
		require.Error(t, err)
		assert.Nil(t, out)
	})

	t.Run("error on put", func(t *testing.T) {
		t.Parallel()

		azureClient := azure.MockAzureClient{
			PutResponseErr: errors.New("error"),
		}
		mockClient := crud.NewResourceCrudClient(&azureClient, nil, nil, "123", &roleManagementPolicyResource)
		roleManagementPolicyClient := &roleManagementPolicyClient{
			client: mockClient,
		}
		out, err := roleManagementPolicyClient.create(context.Background(), "/id", resource.PropertyMap{})
		require.Error(t, err)
		assert.Nil(t, out)
	})
}

func TestUpdateRoleManagementPolicy(t *testing.T) {
	t.Parallel()

	roleManagementPolicyResource := resources.AzureAPIResource{
		PutParameters: []resources.AzureAPIParameter{
			{
				Name:     "properties",
				Location: "body",
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"rules": {
							Type: "array",
							Items: &resources.AzureAPIProperty{
								Type:  "object",
								OneOf: []string{"azure-native:authorization:RoleManagementPolicyExpirationRule"},
							},
						},
					},
				},
			},
		},
	}

	expirationRule := resources.AzureAPIType{
		Properties: map[string]resources.AzureAPIProperty{
			"id":                   {Type: "string"},
			"isExpirationRequired": {Type: "boolean"},
		},
	}

	lookupType := func(ref string) (*resources.AzureAPIType, bool, error) {
		if ref == "#/types/azure-native:authorization:RoleManagementPolicyExpirationRule" {
			return &expirationRule, true, nil
		}
		return nil, false, nil
	}

	converter := convert.NewSdkShapeConverterFull(map[string]resources.AzureAPIType{
		"azure-native:authorization:RoleManagementPolicyExpirationRule": expirationRule,
	})

	// "isExpirationRequired" changes from false to true. The original "false" value is preserved.
	t.Run("restores deleted rules", func(t *testing.T) {
		t.Parallel()

		azureClient := azure.MockAzureClient{}
		mockClient := crud.NewResourceCrudClient(&azureClient, lookupType, &converter, "123", &roleManagementPolicyResource)
		roleManagementPolicyClient := &roleManagementPolicyClient{
			client: mockClient,
		}

		oldsPlain := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": false},
			},
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": []map[string]any{
						{"id": "rule1", "isExpirationRequired": false},
					},
				},
			},
		}
		olds := resource.NewPropertyMapFromMap(oldsPlain)

		newsPlain := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "isExpirationRequired": true},
			},
		}
		news := resource.NewPropertyMapFromMap(newsPlain)

		// The return value is mostly empty because `roleManagementPolicyResource` is a dummy without the response properties.
		out, err := roleManagementPolicyClient.update(context.Background(), "/id", news, olds)
		require.NoError(t, err)
		assert.Equal(t, []string{"/id"}, azureClient.PutIds)
		assert.Len(t, azureClient.PutBodies, 1)
		// We cannot use `assert.Equals` because in one object "rules" is a `[]map[string]map[string]any`, in the other a `[]map[string]any`.
		expectedBody, err := json.Marshal(newsPlain)
		require.NoError(t, err)
		actualBody, err := json.Marshal(azureClient.PutBodies[0])
		require.NoError(t, err)
		assert.JSONEq(t, string(expectedBody), string(actualBody))

		// Check the original state is preserved.
		// We cannot use `assert.Equals` because in one object "rules" is a `[]map[string]map[string]any`, in the other a `[]map[string]any`.
		assert.Contains(t, out, OriginalStateKey)
		expectedOriginalState, err := json.Marshal(oldsPlain[OriginalStateKey])
		require.NoError(t, err)
		actualOriginalState, err := json.Marshal(out[OriginalStateKey])
		require.NoError(t, err)
		assert.JSONEq(t, string(expectedOriginalState), string(actualOriginalState))
	})

	t.Run("error on PUT", func(t *testing.T) {
		t.Parallel()

		azureClient := azure.MockAzureClient{
			PutResponseErr: errors.New("error"),
		}
		mockClient := crud.NewResourceCrudClient(&azureClient, lookupType, &converter, "123", &roleManagementPolicyResource)
		roleManagementPolicyClient := &roleManagementPolicyClient{
			client: mockClient,
		}
		out, err := roleManagementPolicyClient.update(context.Background(), "/id", resource.PropertyMap{}, resource.PropertyMap{})
		require.Error(t, err)
		assert.Nil(t, out)
	})
}

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

	t.Run("restores rules never specified by user", func(t *testing.T) {
		// Original state from Azure has 3 rules, but user only ever specified 1 in their Pulumi program
		origRules := []map[string]any{
			{"id": "rule1", "maximumDuration": "P365D"},
			{"id": "rule2", "isExpirationRequired": false},
			{"id": "rule3", "notificationLevel": "All"},
		}

		// User only specified rule1 in their Pulumi program
		olds := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "maximumDuration": "P365D"},
			},
			OriginalStateKey: map[string]any{
				"properties": map[string]any{
					"rules": origRules,
				},
			},
		}

		// User modifies rule1
		news := map[string]any{
			"rules": []map[string]any{
				{"id": "rule1", "maximumDuration": "P90D"},
			},
		}

		newsPropertyMap := resource.NewPropertyMapFromMap(news)
		restoreDefaultsForDeletedRules(resource.NewPropertyMapFromMap(olds), newsPropertyMap)

		// Should include the modified rule1 AND the unspecified rule2 and rule3 from original state
		rules := newsPropertyMap["rules"].ArrayValue()
		require.Len(t, rules, 3)
		assert.Equal(t,
			[]resource.PropertyValue{
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{"id": "rule1", "maximumDuration": "P90D"})),
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{"id": "rule2", "isExpirationRequired": false})),
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]any{"id": "rule3", "notificationLevel": "All"})),
			},
			newsPropertyMap["rules"].ArrayValue(),
		)
	})
}
