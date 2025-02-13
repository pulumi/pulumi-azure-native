package customresources

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/google/uuid"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	pimRoleEligibilityScheduleTok  = "azure-native:authorization:RoleEligibilitySchedule"
	pimRoleEligibilitySchedulePath = "/{scope}/providers/Microsoft.Authorization/roleEligibilitySchedules/{roleEligibilityScheduleName}"

	// pimRoleEligibilityScheduleRequestTok  = "azure-native:authorization:RoleEligibilityScheduleRequest"
	// pimRoleEligibilityScheduleRequestPath = "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
)

// pimRoleEligibilitySchedule returns a custom resource for Role Eligibility Schedules, used to limit standing
// administrator access to privileged roles in Azure Privileged Identity Management (PIM). See
// https://learn.microsoft.com/en-us/rest/api/authorization/privileged-role-eligibility-rest-sample for details.
//
// Role Eligibility Schedules cannot be created directly. Instead, one submits a Role Eligibility Schedule Request,
// which can then be approved or denied. This resource handles the request internally.
//
// A Role Eligibility Schedule is defined by the triple of scope, principal, and role. At present, only one instance
// of this resource can exist for a given scope|principal|role triple.
//
// Note that this resource cannot be updated. Each change leads to a recreation.
//
// The APIs used are Microsoft.Authorization/roleEligibilitySchedules and Microsoft.Authorization/
// roleEligibilityScheduleRequests.
func pimRoleEligibilitySchedule(
	lookupResource resources.ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	token azcore.TokenCredential,
) (*CustomResource, error) {
	// A bit of a hack to initialize some resource. This func's parameters are all nil when the
	// function is called for the first time, for customresources.featureLookup, which is ok but
	// would break our initialization here.
	var crudClient crud.ResourceCrudClient
	var res resources.AzureAPIResource
	var schedulesClient *armauthorization.RoleEligibilitySchedulesClient
	var requestsClient *armauthorization.RoleEligibilityScheduleRequestsClient

	if lookupResource != nil && crudClientFactory != nil {
		var found bool
		var err error
		res, found, err = lookupResource(pimRoleEligibilityScheduleTok)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", pimRoleEligibilityScheduleTok)
		}
		crudClient = crudClientFactory(&res)

		schedulesClient, err = armauthorization.NewRoleEligibilitySchedulesClient(token, nil)
		if err != nil {
			return nil, err
		}
		requestsClient, err = armauthorization.NewRoleEligibilityScheduleRequestsClient(token, nil)
		if err != nil {
			return nil, err
		}
	}

	schedule := &pimEligibilityScheduleClient{
		crudClient:      crudClient,
		schedulesClient: schedulesClient,
		requestsClient:  requestsClient,
	}
	schedule.lookupScheduleFunc = schedule.lookupSchedule

	return &CustomResource{
		tok:    pimRoleEligibilityScheduleTok,
		path:   pimRoleEligibilitySchedulePath,
		Schema: genSchema,
		Read:   schedule.Read,
		Create: schedule.Create,
		Delete: schedule.Delete,
	}, nil
}

func genSchema(resource *ResourceDefinition) (*ResourceDefinition, error) {
	if resource != nil {
		return nil, fmt.Errorf("resource %q already exists", pimRoleEligibilityScheduleTok)
	}

	resource = &ResourceDefinition{
		Resource: schema.ResourceSpec{
			InputProperties: map[string]schema.PropertySpec{
				"scope": {
					Description:      "The scope of the role eligibility schedule request to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.",
					TypeSpec:         schema.TypeSpec{Type: "string"},
					ReplaceOnChanges: true,
				},
				"principalId": {
					Description:      "The principal ID",
					TypeSpec:         schema.TypeSpec{Type: "string"},
					ReplaceOnChanges: true,
				},
				"roleDefinitionId": {
					Description:      "The role definition ID",
					TypeSpec:         schema.TypeSpec{Type: "string"},
					ReplaceOnChanges: true,
				},
				"justification": {
					Description:      "Justification for the role eligibility schedule request",
					TypeSpec:         schema.TypeSpec{Type: "string"},
					ReplaceOnChanges: true,
				},
				"scheduleInfo": {
					Description: "The schedule information for the role eligibility schedule",
					TypeSpec: schema.TypeSpec{
						Type: "object",
						Ref:  "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesScheduleInfo",
					},
					ReplaceOnChanges: true,
				},
			},
			RequiredInputs: []string{"scope", "principalId", "roleDefinitionId"},
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"scope": {
						Description:      "The scope of the role eligibility schedule to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.",
						TypeSpec:         schema.TypeSpec{Type: "string"},
						ReplaceOnChanges: true,
					},
					"principalId": {
						Description:      "The principal ID",
						TypeSpec:         schema.TypeSpec{Type: "string"},
						ReplaceOnChanges: true,
					},
					"roleDefinitionId": {
						Description:      "The role definition ID",
						TypeSpec:         schema.TypeSpec{Type: "string"},
						ReplaceOnChanges: true,
					},
					"justification": {
						Description:      "Justification for the role eligibility schedule request",
						TypeSpec:         schema.TypeSpec{Type: "string"},
						ReplaceOnChanges: true,
					},
					"scheduleInfo": {
						Description: "The schedule information for the role eligibility schedule",
						TypeSpec: schema.TypeSpec{
							Type: "object",
							Ref:  "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesScheduleInfo",
						},
						ReplaceOnChanges: true,
					},
					"status": {
						Description: "The status of the role eligibility schedule request",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					"startDateTime": {
						Description: "The end date time of the role eligibility schedule",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					"endDateTime": {
						Description: "The end date time of the role eligibility schedule",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
				},
				Required: []string{"scope", "principalId", "roleDefinitionId"},
			},
		},
		MetaResource: resources.AzureAPIResource{
			APIVersion: "2020-10-01",
			Path:       pimRoleEligibilitySchedulePath,
			PutParameters: []resources.AzureAPIParameter{
				{
					Name:       "scope",
					Location:   "path",
					IsRequired: true,
					Value:      &resources.AzureAPIProperty{Type: "string", ForceNew: true},
				},
				{
					Name:       "properties",
					Location:   "body",
					IsRequired: true,
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							"principalId":      {Type: "string", ForceNew: true},
							"roleDefinitionId": {Type: "string", ForceNew: true},
							"justification":    {Type: "string", ForceNew: true},
						},
					},
				},
			},
		},
		Types: map[string]schema.ComplexTypeSpec{
			// "azure-native:authorization:RequestType": {
			// 	ObjectTypeSpec: schema.ObjectTypeSpec{
			// 		Description: "The type of the role assignment schedule request. Eg: SelfActivate, AdminAssign etc",
			// 		Type:        "string",
			// 	},
			// 	Enum: []schema.EnumValueSpec{
			// 		{Value: "Accepted"},
			// 		{Value: "PendingEvaluation"},
			// 		{Value: "Granted"},
			// 		{Value: "Denied"},
			// 		{Value: "PendingProvisioning"},
			// 		{Value: "Provisioned"},
			// 		{Value: "PendingRevocation"},
			// 		{Value: "Revoked"},
			// 		{Value: "Canceled"},
			// 		{Value: "Failed"},
			// 		{Value: "PendingApprovalProvisioning"},
			// 		{Value: "PendingApproval"},
			// 		{Value: "FailedAsResourceIsLocked"},
			// 		{Value: "PendingAdminDecision"},
			// 		{Value: "AdminApproved"},
			// 		{Value: "AdminDenied"},
			// 		{Value: "TimedOut"},
			// 		{Value: "ProvisioningStarted"},
			// 		{Value: "Invalid"},
			// 		{Value: "PendingScheduleCreation"},
			// 		{Value: "ScheduleCreated"},
			// 		{Value: "PendingExternalProvisioning"},
			// 	},
			// },
			// "azure-native:authorization:RoleEligibilityScheduleRequestStatus": {
			// 	ObjectTypeSpec: schema.ObjectTypeSpec{
			// 		Description: "The status of the role eligibility schedule request.",
			// 		Type:        "string",
			// 	},
			// 	Enum: []schema.EnumValueSpec{
			// 		{Value: "Accepted"},
			// 		{Value: "PendingEvaluation"},
			// 		{Value: "Granted"},
			// 		{Value: "Denied"},
			// 		{Value: "PendingProvisioning"},
			// 		{Value: "Provisioned"},
			// 		{Value: "PendingRevocation"},
			// 		{Value: "Revoked"},
			// 		{Value: "Canceled"},
			// 		{Value: "Failed"},
			// 		{Value: "PendingApprovalProvisioning"},
			// 		{Value: "PendingApproval"},
			// 		{Value: "FailedAsResourceIsLocked"},
			// 		{Value: "PendingAdminDecision"},
			// 		{Value: "AdminApproved"},
			// 		{Value: "AdminDenied"},
			// 		{Value: "TimedOut"},
			// 		{Value: "ProvisioningStarted"},
			// 		{Value: "Invalid"},
			// 		{Value: "PendingScheduleCreation"},
			// 		{Value: "ScheduleCreated"},
			// 		{Value: "PendingExternalProvisioning"},
			// 	},
			// },
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpiration": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Expiration of the role eligibility schedule",
					Properties: map[string]schema.PropertySpec{
						"duration":    {TypeSpec: schema.TypeSpec{Type: "string"}},
						"endDateTime": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"type": {
							TypeSpec:    schema.TypeSpec{Type: "string", Ref: "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpirationType"},
							Description: "Type of the role eligibility schedule expiration",
						},
					},
				},
			},
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpirationType": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Type of the role eligibility schedule expiration",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{Value: "AfterDuration"},
					{Value: "AfterDateTime"},
					{Value: "NoExpiration"},
				},
			},
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesScheduleInfo": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Schedule info of the role eligibility schedule",
					Properties: map[string]schema.PropertySpec{
						"startDateTime": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"expiration": {
							TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpiration"},
							Description: "Expiration of the role eligibility schedule",
						},
					},
				},
			},
		},
		MetaTypes: map[string]resources.AzureAPIType{
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpiration": {
				Properties: map[string]resources.AzureAPIProperty{
					"duration":    {Type: "string"},
					"endDateTime": {Type: "string"},
					"type":        {Type: "string", Ref: "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpirationType"},
				},
			},
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpirationType": {},
			"azure-native:authorization:RoleEligibilityScheduleRequestPropertiesScheduleInfo": {
				Properties: map[string]resources.AzureAPIProperty{
					"startDateTime": {Type: "string"},
					"expiration":    {Type: "object", Ref: "#/types/azure-native:authorization:RoleEligibilityScheduleRequestPropertiesExpiration"},
				},
			},
		},
	}
	return resource, nil
}

type pimEligibilityScheduleClient struct {
	crudClient         crud.ResourceCrudClient
	schedulesClient    *armauthorization.RoleEligibilitySchedulesClient
	requestsClient     *armauthorization.RoleEligibilityScheduleRequestsClient
	lookupScheduleFunc lookupScheduleFunc
}

type lookupScheduleFunc func(ctx context.Context, scope, principalId, roleDefinitionId string) (*armauthorization.RoleEligibilitySchedule, error)

func (c *pimEligibilityScheduleClient) lookupSchedule(ctx context.Context, scope, principalId, roleDefinitionId string) (*armauthorization.RoleEligibilitySchedule, error) {
	pager := c.schedulesClient.NewListForScopePager(scope, &armauthorization.RoleEligibilitySchedulesClientListForScopeOptions{
		Filter: pulumi.StringRef(fmt.Sprintf("principalId eq '%s'", principalId)),
		// Filter: pulumi.StringRef("asTarget()"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("listing role eligibility schedules: %w", err)
		}
		for _, schedule := range page.Value {
			if strings.EqualFold(*schedule.Properties.PrincipalID, principalId) &&
				strings.EqualFold(*schedule.Properties.RoleDefinitionID, roleDefinitionId) &&
				*schedule.Properties.MemberType == armauthorization.MemberTypeDirect {
				return schedule, nil
			}
		}
	}
	return nil, nil
}

// TODO with a client for RoleEligibilitySchedule, could we just use c.crudClient.ResponseBodyToSdkOutputs()?
func mapScheduleToOutputs(schedule *armauthorization.RoleEligibilitySchedule, inputs resource.PropertyMap) (map[string]any, error) {
	return map[string]any{
		"scope":            schedule.Properties.Scope,
		"principalId":      schedule.Properties.PrincipalID,
		"roleDefinitionId": schedule.Properties.RoleDefinitionID,
		"startDateTime":    schedule.Properties.StartDateTime,
		"endDateTime":      schedule.Properties.EndDateTime,
		"status":           schedule.Properties.Status,
	}, nil
}

func (c *pimEligibilityScheduleClient) Read(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
	scope := inputs["scope"].StringValue()
	principalId := inputs["principalId"].StringValue()
	roleDefinitionId := inputs["roleDefinitionId"].StringValue()

	existingSchedule, err := c.lookupScheduleFunc(ctx, scope, principalId, roleDefinitionId)
	if err != nil {
		return nil, false, fmt.Errorf("looking up role eligibility schedule: %w", err)
	}

	if existingSchedule == nil {
		return nil, false, nil
	}

	result, err := mapScheduleToOutputs(existingSchedule, inputs)
	if err != nil {
		return nil, true, fmt.Errorf("mapping role eligibility schedule to outputs: %w", err)
	}
	return result, true, nil
}

// Create submits a Role Eligibility Schedule Request to the PIM service. It then waits and polls for the request to be
// in a terminal state like approved or denied. All possible states are defined
// [here](https://learn.microsoft.com/en-us/rest/api/authorization/role-eligibility-schedules/get?view=rest-authorization-2020-10-01&tabs=HTTP#status).
func (c *pimEligibilityScheduleClient) Create(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
	scope := inputs["scope"].StringValue()
	principalId := inputs["principalId"].StringValue()
	roleDefinitionId := inputs["roleDefinitionId"].StringValue()

	// Look for an existing schedule for the same scope|principal|role triple.
	// TODO this is what TF does but it's not clear why there couldn't be multiple schedules.
	schedule, err := c.lookupScheduleFunc(ctx, scope, principalId, roleDefinitionId)
	if err != nil {
		return nil, fmt.Errorf("looking up role eligibility schedule: %w", err)
	}
	if schedule != nil {
		return nil, fmt.Errorf("a role eligibility schedule already exists")
	}

	// Generate a new GUID for the schedule request name
	scheduleRequestName := uuid.New().String()

	// Prepare the payload for the role eligibility schedule request
	// TODO could this just be c.crudClient.PrepareAzureRESTBody() if we had a client for RoleEligibilityScheduleRequest?
	typeAssign := armauthorization.RequestTypeAdminAssign
	scheduleInfo := inputs["scheduleInfo"].ObjectValue()
	startTime, err := time.Parse(time.RFC3339, scheduleInfo["startDateTime"].StringValue())
	if err != nil {
		return nil, fmt.Errorf("invalid start time: %w", err)
	}
	expiration := scheduleInfo["expiration"].ObjectValue()
	expirationType := armauthorization.Type(expiration["type"].StringValue())
	payload := armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Justification:    pulumi.StringRef(inputs["justification"].StringValue()),
			PrincipalID:      pulumi.StringRef(principalId),
			RoleDefinitionID: pulumi.StringRef(roleDefinitionId),
			RequestType:      &typeAssign,
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				StartDateTime: &startTime,
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Duration: pulumi.StringRef(expiration["duration"].StringValue()),
					Type:     &expirationType,
				},
			},
		},
	}

	// Create the request
	_, err = c.requestsClient.Create(ctx, scope, scheduleRequestName, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("creating role eligibility schedule request: %w", err)
	}

	// Poll for completion. Success is when we find a schedule with the matching scope|principal|role triple.
	maxWait := 10 * time.Minute
	timeout := time.After(maxWait)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("timed out waiting %s for role eligibility schedule to be created", maxWait)
		case <-ticker.C:
			// Use lookupSchedule to check if the schedule exists
			schedule, err := c.lookupSchedule(ctx, scope, principalId, roleDefinitionId)
			if err != nil {
				return nil, fmt.Errorf("looking up role eligibility schedule: %w", err)
			}

			if schedule != nil {
				// Schedule found, exit the loop
				return mapScheduleToOutputs(schedule, inputs)
			}
		}
	}
}

// Delete checks if the Role Eligibility Schedule Request for this schedule is still active, i.e., in a non-terminal
// state. If so, it cancels the request. Otherwise, it submits a new Schedule Request with RequestType=AdminRemove to
// delete the schedule.
func (c *pimEligibilityScheduleClient) Delete(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
	// Extract necessary information from the state
	scope := state["scope"].StringValue()
	principalId := state["principalId"].StringValue()
	roleDefinitionId := state["roleDefinitionId"].StringValue()

	// Check the current status of the schedule
	schedule, err := c.schedulesClient.Get(ctx, scope, id, nil)
	if err != nil {
		var responseErr *azcore.ResponseError
		if errors.As(err, &responseErr) && responseErr.StatusCode == 404 {
			// Schedule not found, nothing to delete
			return nil
		}
		return fmt.Errorf("retrieving role eligibility schedule: %w", err)
	}

	// If the schedule is active, cancel it
	if *schedule.Properties.Status == armauthorization.StatusPendingApproval || *schedule.Properties.Status == armauthorization.StatusGranted {
		_, err := c.requestsClient.Cancel(ctx, scope, id, nil)
		if err != nil {
			return fmt.Errorf("canceling role eligibility schedule: %w", err)
		}
		return nil
	}

	// If the schedule is not active, submit a new request with RequestType=AdminRemove
	typeRemove := armauthorization.RequestTypeAdminRemove
	payload := armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			PrincipalID:      pulumi.StringRef(principalId),
			RoleDefinitionID: pulumi.StringRef(roleDefinitionId),
			RequestType:      &typeRemove,
			Justification:    pulumi.StringRef("Removed by Pulumi"),
		},
	}

	// Generate a new GUID for the removal request
	removeRequestName := uuid.New().String()

	_, err = c.requestsClient.Create(ctx, scope, removeRequestName, payload, nil)
	if err != nil {
		return fmt.Errorf("creating role eligibility schedule removal request: %w", err)
	}

	// Poll for the removal request to be processed.
	maxWait := 10 * time.Minute
	timeout := time.After(maxWait)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting %s for role eligibility schedule to be deleted", maxWait)
		case <-ticker.C:
			request, err := c.requestsClient.Get(ctx, scope, removeRequestName, nil)
			if err != nil {
				return fmt.Errorf("retrieving role eligibility schedule removal request: %w", err)
			}

			status := request.Properties.Status
			if *status == armauthorization.StatusAdminApproved || *status == armauthorization.StatusGranted {
				break
			}

			if *status == armauthorization.StatusAdminDenied {
				return fmt.Errorf("role eligibility schedule removal request denied")
			}

			time.Sleep(10 * time.Second) // Polling interval
		}
	}
}
