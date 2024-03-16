package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const OriginalStateKey = "__orig_state"

type Rule struct {
	Id string
}

type PolicyWithRules struct {
	Rules []Rule
}

func pimRoleManagementPolicy(lookupResource resources.ResourceLookupFunc, crudClientFactory crud.ResourceCrudClientFactory) (*CustomResource, error) {
	// A bit of a hack to initialize some resource. This func's parameters are all nil when the
	// function is called for the first time, for customresources.featureLookup, which is ok but
	// would break our initialization here.
	var client crud.ResourceCrudClient
	var res resources.AzureAPIResource
	if lookupResource != nil && crudClientFactory != nil {
		var found bool
		var err error
		res, found, err = lookupResource("azure-native:authorization:RoleManagementPolicy")
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", "azure-native:authorization:RoleManagementPolicy")
		}

		client = crudClientFactory(&res)
	}

	return &CustomResource{
		path: "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}",

		// PIM Role Management Policies cannot be created. Instead, each resource has a default policy.
		// But we need to allow the user to create the Pulumi resource, so we return true here and then
		// let Create just return the existing policy.
		CanCreate: func(ctx context.Context, id string) error {
			return nil
		},

		// PIM Role Management Policies cannot be created. Instead, each resource has a default policy.
		// Simply look up the default policy and return it.
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			// First read the existing policy to capture the default for resetting it later on delete.
			originalState, err := client.Read(ctx, id)
			if err != nil {
				return nil, err
			}

			bodyParams, err := client.PrepareAzureRESTBody(id, inputs)
			if err != nil {
				return nil, err
			}
			queryParams := map[string]any{"api-version": client.ApiVersion()}

			// TODO we could skip this if bodyParams = originalState, i.e., the user adds a policy
			// in its default configuration to their program
			resp, _, err := client.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				return nil, err
			}

			// TODO sort rules?

			outputs := client.ResponseBodyToSdkOutputs(resp)
			outputs[OriginalStateKey] = originalState
			return outputs, nil
		},

		// Tricky because rules can be removed from the list of rules of the policy, but simply removing them from the
		// request will leave them in their current state, not the default state.
		Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
			if !olds.HasValue(OriginalStateKey) {
				logging.V(3).Infof("Warning: no original state found for %s, cannot reset deleted rules", id)
			} else {
				restoreDefaultsForDeletedRules(olds, news)
			}

			bodyParams, err := client.PrepareAzureRESTBody(id, news)
			if err != nil {
				return nil, err
			}
			queryParams := map[string]any{"api-version": client.ApiVersion()}

			resp, _, err := client.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				return nil, err
			}

			outputs := client.ResponseBodyToSdkOutputs(resp)
			outputs[OriginalStateKey] = olds[OriginalStateKey]
			return outputs, nil
		},

		// PIM Role Management Policies cannot be deleted. Instead, we reset the policy to its default.
		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			queryParams := map[string]any{"api-version": client.ApiVersion()}
			if !state.HasValue(OriginalStateKey) {
				logging.V(3).Infof("Warning: no original state found for %s, cannot reset", id)
				return nil
			}

			origState := state[OriginalStateKey].Mappable().(map[string]any)
			pathItems, err := resources.ParseResourceID(id, res.Path)
			if err != nil {
				return err
			}
			origSdkInputs := client.ResponseToSdkInputs(pathItems, origState)
			origRequest, err := client.SdkInputsToRequestBody(origSdkInputs, id)
			if err != nil {
				return err
			}

			_, _, err = client.CreateOrUpdate(ctx, id, origRequest, queryParams)
			return err
		},
	}, nil
}

// restoreDefaultsForDeletedRules restores the original values for rules that were deleted from the policy.
// For each rule in olds that's not in news, it looks up the original rule in olds and adds it to news.
func restoreDefaultsForDeletedRules(olds, news resource.PropertyMap) {
	oldRules := mapRulesById(olds)
	if len(oldRules) == 0 {
		return
	}

	newRules := mapRulesById(news)

	origState := olds[OriginalStateKey].ObjectValue()
	if !origState.HasValue("properties") {
		logging.V(3).Infof("Warning: restoreDefaultsForDeletedRules: no 'properties' in original state")
		return
	}
	origRules := mapRulesById(origState["properties"].ObjectValue())

	newRulesList := []resource.PropertyValue{}
	if len(newRules) > 0 {
		newRulesList = news["rules"].ArrayValue()
	}

	for id := range oldRules {
		if _, ok := newRules[id]; !ok {
			if origRule, ok := origRules[id]; ok {
				newRulesList = append(newRulesList, origRule)
			} else {
				logging.V(3).Infof("Warning: restoreDefaultsForDeletedRules: rule %s not found in original state", id)
			}
		}
	}

	// TODO sort?
	news["rules"] = resource.NewArrayProperty(newRulesList)
}

func mapRulesById(managementPolicy resource.PropertyMap) map[string]resource.PropertyValue {
	if !managementPolicy.HasValue("rules") {
		return nil
	}
	rules := managementPolicy["rules"].ArrayValue()

	rulesById := map[string]resource.PropertyValue{}
	for _, rule := range rules {
		if !rule.IsObject() || !rule.ObjectValue().HasValue("id") {
			logging.V(3).Infof("Warning: mapRulesById: rule has no id: %v", rule)
			return nil
		}
		rulesById[rule.ObjectValue()["id"].StringValue()] = rule
	}
	return rulesById
}
