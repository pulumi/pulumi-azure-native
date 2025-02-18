// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/google/uuid"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	pimRoleEligibilityScheduleResourceDescription = `A PIM (Privileged Identity Management) Role Eligibility Schedule.

Role Eligibility Schedules are used to limit standing administrator access to privileged roles in Azure PIM. See
[here](https://learn.microsoft.com/en-us/rest/api/authorization/privileged-role-eligibility-rest-sample) for details.

A Role Eligibility Schedule is uniquely defined by scope, principal, and role. At present, only one instance of this
resource can exist for a given scope|principal|role tuple.

Note that this resource cannot be updated. Each change leads to a recreation.

Internally, this resource uses the
[Role Eligibility Schedule Requests](https://learn.microsoft.com/en-us/rest/api/authorization/role-eligibility-schedule-requests?view=rest-authorization-2020-10-01)
API to create and delete the schedules.
`

	pimRoleEligibilityScheduleResourceName = "PimRoleEligibilitySchedule"
	pimRoleEligibilityScheduleTok          = "azure-native:authorization:" + pimRoleEligibilityScheduleResourceName

	pimRoleEligibilityScheduleRequestResourceName = "RoleEligibilityScheduleRequest"
	pimRoleEligibilityScheduleRequestTok          = "azure-native:authorization:" + pimRoleEligibilityScheduleRequestResourceName
	PimRoleEligibilityScheduleRequestPath         = "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"

	pimRoleEligibilityScheduleMaxWaitForCreate        = 10 * time.Minute
	pimRoleEligibilityScheduleMaxWaitForDelete        = 10 * time.Minute
	pimRoleEligibilityScheduleMaxWaitForDeleteRequest = 2 * time.Minute
)

// Tests are allowed to change this to a smaller value.
var pimRoleEligibilityScheduleTickerInterval = 10 * time.Second

// pimRoleEligibilitySchedule returns a custom resource for Role Eligibility Schedules.
//
// See the docs in the constant `pimRoleEligibilityScheduleResourceDescription` for details.
//
// The APIs used are Microsoft.Authorization/roleEligibilitySchedules and Microsoft.Authorization/
// roleEligibilityScheduleRequests.
func pimRoleEligibilitySchedule(
	lookupResource resources.ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	token azcore.TokenCredential,
) (*CustomResource, error) {
	// A bit of a hack to initialize some resource. This func's parameters are all nil when the function is called for
	// the first time, for customresources.featureLookup, which is ok but would break our initialization here.
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

	client := &pimEligibilityScheduleClientImpl{
		crudClient:      crudClient,
		schedulesClient: schedulesClient,
		requestsClient:  requestsClient,
	}

	return &CustomResource{
		tok:                pimRoleEligibilityScheduleRequestTok,
		CustomResourceName: pimRoleEligibilityScheduleResourceName,
		path:               PimRoleEligibilityScheduleRequestPath,
		Schema:             genSchema,
		Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
			return read(ctx, id, inputs, client)
		},
		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			return createPimEligibilitySchedule(ctx, id, inputs, client, pimRoleEligibilityScheduleMaxWaitForCreate)
		},
		Delete: func(ctx context.Context, id string, _, state resource.PropertyMap) error {
			return deletePimEligibilitySchedule(ctx, id, state, client, pimRoleEligibilityScheduleMaxWaitForDelete)
		},
	}, nil
}

// genSchema modifies the schema extracted from the Azure spec.
func genSchema(resource *ResourceDefinition) (*ResourceDefinition, error) {
	if resource == nil {
		return nil, fmt.Errorf("resource %q not found in the spec", pimRoleEligibilityScheduleRequestTok)
	}

	// remove requestType, we're only supporting one so far (like TF)
	delete(resource.Resource.InputProperties, "requestType")
	resource.makePropertyNotRequired("requestType")

	// remove the name, it's a random GUID so we generate it ourselves
	delete(resource.Resource.InputProperties, "roleEligibilityScheduleRequestName")

	// This resource doesn't support direct updates, all changes need to go through a new schedule request.
	for name, p := range resource.Resource.InputProperties {
		p.ReplaceOnChanges = true
		resource.Resource.InputProperties[name] = p
	}

	updateResourceDescription(resource, pimRoleEligibilityScheduleResourceDescription)

	return resource, nil
}

// updateResourceDescription replaces the resource description with the given one, but preserves API version
// documentation if present.
func updateResourceDescription(resource *ResourceDefinition, newDescription string) {
	if idx := strings.Index(resource.Resource.Description, "Azure REST API version:"); idx != -1 {
		newDescription += "\n\n" + resource.Resource.Description[idx:]
	}
	resource.Resource.Description = newDescription
}

// pimEligibilityScheduleClient is a client for the PIM Role Eligibility Schedule API. The interface allows to use fake
// clients in tests.
type pimEligibilityScheduleClient interface {
	findSchedule(ctx context.Context, scope, principalId, roleDefinitionId string) (*armauthorization.RoleEligibilitySchedule, error)
	getSchedule(ctx context.Context, scope, scheduleName string) (*armauthorization.RoleEligibilitySchedule, error)
	// Matches armauthorization.RoleEligibilityScheduleRequestsClient.Create
	createSchedule(ctx context.Context, scope string, roleEligibilityScheduleRequestName string,
		parameters armauthorization.RoleEligibilityScheduleRequest,
	) (armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse, error)
	cancelSchedule(ctx context.Context, scope, scheduleName string) (armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse, error)
	// Map an Azure SDK schedule object to Pulumi SDK shape output.
	mapScheduleToOutputs(schedule *armauthorization.RoleEligibilitySchedule) (map[string]any, error)
}

type pimEligibilityScheduleClientImpl struct {
	crudClient      crud.ResourceCrudClient
	schedulesClient *armauthorization.RoleEligibilitySchedulesClient
	requestsClient  *armauthorization.RoleEligibilityScheduleRequestsClient
}

func (c *pimEligibilityScheduleClientImpl) findSchedule(ctx context.Context, scope, principalId, roleDefinitionId string,
) (*armauthorization.RoleEligibilitySchedule, error) {
	pager := c.schedulesClient.NewListForScopePager(scope, &armauthorization.RoleEligibilitySchedulesClientListForScopeOptions{
		Filter: pulumi.StringRef(fmt.Sprintf("principalId eq '%s'", principalId)),
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

func (c *pimEligibilityScheduleClientImpl) getSchedule(ctx context.Context, scope, scheduleName string) (*armauthorization.RoleEligibilitySchedule, error) {
	s, err := c.schedulesClient.Get(ctx, scope, scheduleName, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving role eligibility schedule: %w", err)
	}
	return &s.RoleEligibilitySchedule, nil
}

func (c *pimEligibilityScheduleClientImpl) createSchedule(ctx context.Context, scope string, roleEligibilityScheduleRequestName string,
	parameters armauthorization.RoleEligibilityScheduleRequest,
) (armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse, error) {
	return c.requestsClient.Create(ctx, scope, roleEligibilityScheduleRequestName, parameters, nil)
}

func (c *pimEligibilityScheduleClientImpl) cancelSchedule(ctx context.Context, scope, scheduleName string) (armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse, error) {
	return c.requestsClient.Cancel(ctx, scope, scheduleName, nil)
}

// Map an Azure SDK schedule object to Pulumi SDK shape output. Since the SDK uses the same API, it's almost in the
// right shape, except the SDK ignores the `"x-ms-client-flatten": true` annotation in the API spec, so we run it
// through our converter to remove the extra layer of nesting.
func (c *pimEligibilityScheduleClientImpl) mapScheduleToOutputs(schedule *armauthorization.RoleEligibilitySchedule) (map[string]any, error) {
	j, err := json.Marshal(schedule)
	if err != nil {
		return nil, fmt.Errorf("converting role eligibility schedule from SDK to JSON: %w", err)
	}
	var js map[string]any
	err = json.Unmarshal(j, &js)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling role eligibility schedule from JSON: %w", err)
	}

	// We don't need the "Microsoft.Authorization/roleEligibilitySchedules" type that the SDK adds.
	delete(js, "type")

	result := c.crudClient.ResponseBodyToSdkOutputs(js)
	return result, nil
}

func read(ctx context.Context, id string, inputs resource.PropertyMap, client pimEligibilityScheduleClient) (map[string]any, bool, error) {
	sdkInputs, err := inputsToSdk(inputs)
	if err != nil {
		return nil, false, fmt.Errorf("converting inputs to SDK shape: %w", err)
	}

	existingSchedule, err := client.findSchedule(ctx, *sdkInputs.Properties.Scope,
		*sdkInputs.Properties.PrincipalID, *sdkInputs.Properties.RoleDefinitionID)
	if err != nil {
		return nil, false, fmt.Errorf("looking up role eligibility schedule: %w", err)
	}

	if existingSchedule == nil {
		return nil, false, nil
	}

	result, err := client.mapScheduleToOutputs(existingSchedule)
	if err != nil {
		return nil, true, fmt.Errorf("mapping role eligibility schedule to outputs: %w", err)
	}
	return result, true, nil
}

// Create submits a Role Eligibility Schedule Request to the PIM service. It then waits and polls for the request to be
// in a terminal state like approved or denied. All possible states are defined
// [here](https://learn.microsoft.com/en-us/rest/api/authorization/role-eligibility-schedules/get?view=rest-authorization-2020-10-01&tabs=HTTP#status).
func createPimEligibilitySchedule(ctx context.Context, id string, inputs resource.PropertyMap,
	client pimEligibilityScheduleClient, maxWait time.Duration,
) (map[string]any, error) {
	payload, err := inputsToSdk(inputs)
	if err != nil {
		return nil, fmt.Errorf("converting inputs to SDK shape: %w", err)
	}

	// Generate a new GUID for the schedule request name
	scheduleRequestName := uuid.New().String()
	payload.Name = nil

	typeAssign := armauthorization.RequestTypeAdminAssign
	payload.Properties.RequestType = &typeAssign

	// Create the schedule request
	_, err = client.createSchedule(ctx, *payload.Properties.Scope, scheduleRequestName, *payload)
	if err != nil {
		return nil, fmt.Errorf("creating role eligibility schedule request: %w", err)
	}

	// Poll for completion. Success is when we find a schedule with the matching scope|principal|role tuple.
	var schedule *armauthorization.RoleEligibilitySchedule
	err = util.RetryOperation(maxWait, pimRoleEligibilityScheduleTickerInterval,
		"waiting for role eligibility schedule to be created",
		func() (bool, error) {
			var err error
			schedule, err = client.findSchedule(ctx, *payload.Properties.Scope,
				*payload.Properties.PrincipalID, *payload.Properties.RoleDefinitionID)
			if err != nil {
				return true, fmt.Errorf("looking up role eligibility schedule: %w", err)
			}
			return schedule != nil, nil
		})
	if err != nil {
		return nil, err
	}

	return client.mapScheduleToOutputs(schedule)
}

// Delete checks if the Role Eligibility Schedule Request for this schedule is still active, i.e., in a non-terminal
// state. If so, it cancels the request. Otherwise, it submits a new Schedule Request with RequestType=AdminRemove to
// delete the schedule.
func deletePimEligibilitySchedule(ctx context.Context,
	id string,
	state resource.PropertyMap,
	client pimEligibilityScheduleClient,
	maxWait time.Duration,
) error {
	sdkState, err := inputsToSdk(state)
	if err != nil {
		return fmt.Errorf("converting inputs to SDK shape: %w", err)
	}

	scheduleName := *sdkState.Name
	scope := *sdkState.Properties.Scope

	// Check the current status of the schedule
	schedule, err := client.getSchedule(ctx, scope, scheduleName)
	if err != nil {
		var responseErr *azcore.ResponseError
		if errors.As(err, &responseErr) && responseErr.StatusCode == 404 {
			// Schedule not found, nothing to delete
			return nil
		}
		return fmt.Errorf("retrieving role eligibility schedule: %w", err)
	}

	// If the schedule is active, cancel it
	if statusIsPending(schedule.Properties.Status) {
		_, err := client.cancelSchedule(ctx, *sdkState.Properties.Scope, id)
		if err != nil {
			return fmt.Errorf("canceling role eligibility schedule: %w", err)
		}
		return nil
	}

	// If the schedule is not active, submit a new request with RequestType=AdminRemove.
	// Generate a new GUID for the removal request.
	removeRequestName := uuid.New().String()

	typeRemove := armauthorization.RequestTypeAdminRemove
	payload := armauthorization.RoleEligibilityScheduleRequest{
		Name: pulumi.StringRef(removeRequestName),
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			PrincipalID:      sdkState.Properties.PrincipalID,
			RoleDefinitionID: sdkState.Properties.RoleDefinitionID,
			RequestType:      &typeRemove,
			Justification:    pulumi.StringRef("Removed by Pulumi"),
		},
	}

	err = util.RetryOperation(2*time.Minute, pimRoleEligibilityScheduleTickerInterval, "requesting role eligibility schedule removal", func() (bool, error) {
		_, err = client.createSchedule(ctx, scope, removeRequestName, payload)
		if err != nil {
			var responseErr *azcore.ResponseError
			if errors.As(err, &responseErr) {
				// occasional intermittent error, retry
				if responseErr.ErrorCode == "ActiveDurationTooShort" {
					return false, nil
				}
				// already deleted, we're done
				if responseErr.ErrorCode == "RoleAssignmentDoesNotExist" {
					return true, nil
				}
			}
			return true, fmt.Errorf("creating role eligibility schedule removal request: %w", err)
		}
		return true, nil
	})
	if err != nil {
		return err
	}

	// Poll for the schedule to be gone.
	err = util.RetryOperation(maxWait, pimRoleEligibilityScheduleTickerInterval, "role eligibility schedule deletion", func() (bool, error) {
		_, err := client.getSchedule(ctx, scope, scheduleName)
		if err != nil {
			var responseErr *azcore.ResponseError
			if errors.As(err, &responseErr) && responseErr.StatusCode == 404 {
				return true, nil
			}
			return false, fmt.Errorf("retrieving role eligibility schedule: %w", err)
		}
		return false, nil
	})
	return err
}

func statusIsPending(status *armauthorization.Status) bool {
	return *status == armauthorization.StatusGranted ||
		strings.HasPrefix(string(*status), "Pending")
}

// inputsToSdk converts the Pulumi SDK shape inputs to the Azure SDK shape.
func inputsToSdk(inputs resource.PropertyMap) (*armauthorization.RoleEligibilityScheduleRequest, error) {
	result := &armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{},
	}

	if name, ok := inputs["roleEligibilityScheduleRequestName"]; ok {
		result.Name = pulumi.StringRef(name.StringValue())
	} else if name, ok = inputs["name"]; ok {
		result.Name = pulumi.StringRef(name.StringValue())
	}

	if scope, ok := inputs["scope"]; ok {
		result.Properties.Scope = pulumi.StringRef(scope.StringValue())
	} else {
		return nil, fmt.Errorf("scope is required")
	}

	if principalId, ok := inputs["principalId"]; ok {
		result.Properties.PrincipalID = pulumi.StringRef(principalId.StringValue())
	} else {
		return nil, fmt.Errorf("principalId is required")
	}

	if roleDefinitionId, ok := inputs["roleDefinitionId"]; ok {
		result.Properties.RoleDefinitionID = pulumi.StringRef(roleDefinitionId.StringValue())
	} else {
		return nil, fmt.Errorf("roleDefinitionId is required")
	}

	if justification, ok := inputs["justification"]; ok {
		result.Properties.Justification = pulumi.StringRef(justification.StringValue())
	}

	if status, ok := inputs["status"]; ok {
		statusVal := armauthorization.Status(status.StringValue())
		result.Properties.Status = &statusVal
	}

	if scheduleInfo, ok := inputs["scheduleInfo"]; ok {
		info := scheduleInfo.ObjectValue()
		result.Properties.ScheduleInfo = &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
			Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{},
		}

		if startDateTime, ok := info["startDateTime"]; ok {
			startTime, err := time.Parse(time.RFC3339, startDateTime.StringValue())
			if err != nil {
				return nil, fmt.Errorf("invalid start time: %w", err)
			}
			result.Properties.ScheduleInfo.StartDateTime = &startTime
		}

		if expiration, ok := info["expiration"]; ok {
			exp := expiration.ObjectValue()
			result.Properties.ScheduleInfo.Expiration.Duration = pulumi.StringRef(exp["duration"].StringValue())
			expirationType := armauthorization.Type(exp["type"].StringValue())
			result.Properties.ScheduleInfo.Expiration.Type = &expirationType
		}
	}

	return result, nil
}
