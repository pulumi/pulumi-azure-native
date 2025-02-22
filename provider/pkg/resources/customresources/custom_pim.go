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

type Rule struct {
	Id string
}

type PolicyWithRules struct {
	Rules []Rule
}

// roleManagementPolicyClient is a helper struct containing dependencies and methods for CRUD operations on PIM Role Management Policies.
type roleManagementPolicyClient struct {
	client           crud.ResourceCrudClient
	resourceMetadata resources.AzureAPIResource
}

func (c *roleManagementPolicyClient) create(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
	originalState, err := c.client.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	bodyParams, err := c.client.PrepareAzureRESTBody(id, inputs)
	if err != nil {
		return nil, err
	}
	queryParams := map[string]any{"api-version": c.client.ApiVersion()}

	// We could skip this if bodyParams = originalState, i.e., the user adds a policy
	// in its default configuration to their program, but we don't have a diff function.
	resp, _, err := c.client.Update(ctx, id, bodyParams, queryParams)
	if err != nil {
		return nil, err
	}

	outputs := c.client.ResponseBodyToSdkOutputs(resp)
	outputs[OriginalStateKey] = originalState
	return outputs, nil
}

func (c *roleManagementPolicyClient) update(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
	if !olds.HasValue(OriginalStateKey) {
		logging.V(3).Infof("Warning: no original state found for %s, cannot reset deleted rules", id)
	} else {
		restoreDefaultsForDeletedRules(olds, news)
	}

	bodyParams, err := c.client.PrepareAzureRESTBody(id, news)
	if err != nil {
		return nil, err
	}
	queryParams := map[string]any{"api-version": c.client.ApiVersion()}

	resp, _, err := c.client.Update(ctx, id, bodyParams, queryParams)
	if err != nil {
		return nil, err
	}

	outputs := c.client.ResponseBodyToSdkOutputs(resp)
	if olds.HasValue(OriginalStateKey) {
		outputs[OriginalStateKey] = olds[OriginalStateKey].Mappable()
	}
	return outputs, nil
}

func (c *roleManagementPolicyClient) delete(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
	if !state.HasValue(OriginalStateKey) {
		logging.V(3).Infof("Warning: no original state found for %s, cannot reset", id)
		return nil
	}

	origState := state[OriginalStateKey].Mappable().(map[string]any)
	pathItems, err := resources.ParseResourceID(id, c.resourceMetadata.Path)
	if err != nil {
		return err
	}
	origSdkInputs := c.client.ResponseToSdkInputs(pathItems, origState)
	origRequest, err := c.client.SdkInputsToRequestBody(origSdkInputs, id)
	if err != nil {
		return err
	}

	queryParams := map[string]any{"api-version": c.client.ApiVersion()}
	_, _, err = c.client.Update(ctx, id, origRequest, queryParams)
	return err
}

// pimRoleManagementPolicy creates a CustomResource for PIM Role Management Policies. See #2455 and
// https://learn.microsoft.com/en-us/rest/api/authorization/privileged-role-policy-rest-sample.
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

	policyClient := &roleManagementPolicyClient{
		client:           client,
		resourceMetadata: res,
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
		Create: policyClient.create,

		// Tricky because rules can be removed from the list of rules of the policy, but simply removing them from the
		// request will leave them in their current state, not the default state.
		Update: policyClient.update,

		// PIM Role Management Policies cannot be deleted. Instead, we reset the policy to its default.
		Delete: policyClient.delete,
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

	sortRules(newRulesList)
	news["rules"] = resource.NewArrayProperty(newRulesList)
}

func sortRules(rules []resource.PropertyValue) {
	slices.SortFunc(rules, func(a, b resource.PropertyValue) int {
		return cmp.Compare(
			a.ObjectValue()["id"].StringValue(),
			b.ObjectValue()["id"].StringValue())
	})
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
