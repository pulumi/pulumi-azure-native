package customresources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func pimEligibleRoleAssignmentResource(client *armauthorization.RoleEligibilityScheduleInstancesClient) *CustomResource {
	a := &pimEligibleRoleAssignment{}
	return &CustomResource{
		path:   "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{roleEligibilityScheduleInstanceName}",
		tok:    "azure-native:authorization:PimEligibleRoleAssignment",
		Create: a.createOrUpdate,
		Update: a.createOrUpdate,
		Read:   a.read,
		Delete: a.delete,
		Schema: &schema.ResourceSpec{},
		Meta:   &AzureAPIResource{},
	}
}

type pimEligibleRoleAssignment struct {
	client         *armauthorization.RoleEligibilityScheduleInstancesClient
	requestsClient *armauthorization.RoleEligibilityScheduleRequestsClient
}

func (a *pimEligibleRoleAssignment) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	// TODO TF retrieves first the instance (and fails if it doesn't exist), but then uses it only to get the GUID to
	// retrieve the request, from which it takes the data. Why not use the instance directly?

	scope := properties["scope"].StringValue()
	principalId := properties["principalId"].StringValue()
	roleDefinitionId := properties["roleDefinitionId"].StringValue()

	filterExpr := fmt.Sprintf("(principalId eq '%s' and roleDefinitionId eq '%s')", principalId, roleDefinitionId)
	p := a.client.NewListForScopePager(scope, &armauthorization.RoleEligibilityScheduleInstancesClientListForScopeOptions{
		Filter: &filterExpr,
	})

	// TODO there can be more pages
	next, err := p.NextPage(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("listing role assignments on scope %s: %+v", id, err)
	}

	var instance *armauthorization.RoleEligibilityScheduleInstance
	for _, item := range next.Value {
		if *item.Properties.MemberType == armauthorization.MemberTypeDirect &&
			strings.EqualFold(*item.Properties.Scope, scope) {
			instance = item
			break
		}
	}
	if instance == nil {
		return nil, false, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	jsonInst, err := instance.MarshalJSON()
	if err != nil {
		return nil, false, fmt.Errorf("serializing RoleEligibilityScheduleInstance: %+v", err)
	}
	result := map[string]any{}
	json.Unmarshal(jsonInst, &result)
	return result, true, nil
}
