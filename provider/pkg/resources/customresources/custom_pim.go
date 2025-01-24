package customresources

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// OriginalStateKey is an internal key in the state of a PIM Role Management Policy that contains the original state of the policy.
const OriginalStateKey = "__orig_state"

// pimRoleManagementPolicy returns a custom resource for
// [PIM Role Management Policies](https://learn.microsoft.com/en-us/entra/id-governance/privileged-identity-management/pim-how-to-add-role-to-user#next-steps).
// The custom resource uses the existing schema and metadata for the resource
// azure-native:authorization:RoleManagementPolicy from the Azure spec. The CRUD methods are custom because Role
// Management Policies cannot be created or deleted, just modified resp. restored to their default state. Azure doesn't
// provide a way to get the default state of a policy, so we need to store it in the state of the resource.
//
// Note also the [license requirements](https://learn.microsoft.com/en-us/entra/id-governance/licensing-fundamentals#valid-licenses-for-pim).
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
		isSingleton: true,
		path:        "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}",

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

			// We could skip this if bodyParams = originalState, i.e., the user adds a policy
			// in its default configuration to their program, but we don't have a diff function.
			resp, _, err := client.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				return nil, err
			}

			outputs := client.ResponseBodyToSdkOutputs(resp)
			outputs[OriginalStateKey] = originalState
			return outputs, nil
		},

		// When rules are removed from the PIM policy, simply removing them from the rules in the update request will
		// leave them in their current state. Therefore, we need to restore deleted rules to their original values.
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
			if olds.HasValue(OriginalStateKey) {
				outputs[OriginalStateKey] = olds[OriginalStateKey].Mappable()
			}
			return outputs, nil
		},

		// PIM Role Management Policies cannot be deleted. Instead, we reset the policy to its default.
		Delete: func(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
			req, err := prepareUpdateRequestForDelete(client, id, res.Path, state)
			if err != nil {
				return err
			}

			queryParams := map[string]any{"api-version": client.ApiVersion()}

			_, _, err = client.CreateOrUpdate(ctx, id, req, queryParams)
			return err
		},
	}, nil
}

func prepareUpdateRequestForDelete(client crud.ResourceCrudClient, id, resourcePath string, state resource.PropertyMap) (map[string]any, error) {
	origState, err := getOriginalState(state)
	if err != nil {
		return nil, err
	}
	pathItems, err := resources.ParseResourceID(id, resourcePath)
	if err != nil {
		return nil, err
	}
	sdkInputs := client.ResponseToSdkInputs(pathItems, origState.Mappable())
	req, err := client.SdkInputsToRequestBody(sdkInputs, id)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// restoreDefaultsForDeletedRules restores the original values for rules that were deleted from the policy.
// For each rule in olds that's not in news, it looks up the original rule in olds and adds it to news.
func restoreDefaultsForDeletedRules(olds, news resource.PropertyMap) error {
	oldRules := mapRulesById(olds)
	if len(oldRules) == 0 {
		return nil
	}

	newRules := mapRulesById(news)

	origState, err := getOriginalState(olds)
	if err != nil {
		return err
	}
	if !origState.HasValue("properties") {
		return fmt.Errorf("[PIM] restoreDefaultsForDeletedRules: original state to restore has no 'properties'")
	}
	origProperties := mapRulesById(origState["properties"].ObjectValue())

	newRulesList := []resource.PropertyValue{}
	if len(newRules) > 0 {
		newRulesList = news["rules"].ArrayValue()
	}

	for id := range oldRules {
		if _, ok := newRules[id]; !ok {
			if origRule, ok := origProperties[id]; ok {
				newRulesList = append(newRulesList, origRule)
			} else {
				logging.V(3).Infof("Warning: restoreDefaultsForDeletedRules: rule %s not found in original state", id)
			}
		}
	}

	sortRules(newRulesList)
	news["rules"] = resource.NewArrayProperty(newRulesList)
	return nil
}

func getOriginalState(state resource.PropertyMap) (resource.PropertyMap, error) {
	origStateRaw, ok := state[OriginalStateKey]
	if !ok {
		return nil, fmt.Errorf("[PIM] restoreDefaultsForDeletedRules: no original state to restore")
	}
	return origStateRaw.ObjectValue(), nil
}

func sortRules(rules []resource.PropertyValue) {
	slices.SortFunc(rules, func(a, b resource.PropertyValue) int {
		return cmp.Compare(
			a.ObjectValue()["id"].StringValue(),
			b.ObjectValue()["id"].StringValue())
	})
}

// Put each rule $r in `managementPolicy["rules"]` into a `map[$r.id]$r`.
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
