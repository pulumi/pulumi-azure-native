package customresources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/google/uuid"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryboe/q"
)

const (
	pimRoleEligibilityScheduleRequestTok  = "azure-native:authorization/v20201001:RoleEligibilityScheduleRequest"
	pimRoleEligibilityScheduleRequestPath = "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
)

// pimRoleEligibilityScheduleRequest returns a custom resource for limiting standing administrator access to privileged roles in
// Azure Privileged Identity Management (PIM). See
// https://learn.microsoft.com/en-us/rest/api/authorization/privileged-role-eligibility-rest-sample for details.
//
// The APIs used are Microsoft.Authorization/roleEligibilitySchedules and Microsoft.Authorization/
// roleEligibilityScheduleInstances. The difference between schedules and schedule instances is that while schedule
// instances only include assignments that are active at the current time, schedules also include assignments that will
// become active in the future.
func pimRoleEligibilityScheduleRequest(lookupResource resources.ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	token azcore.TokenCredential,
) (*CustomResource, error) {
	// A bit of a hack to initialize some resource. This func's parameters are all nil when the
	// function is called for the first time, for customresources.featureLookup, which is ok but
	// would break our initialization here.
	var client crud.ResourceCrudClient
	var res resources.AzureAPIResource
	if lookupResource != nil && crudClientFactory != nil {
		var found bool
		var err error
		res, found, err = lookupResource(pimRoleEligibilityScheduleRequestTok)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", pimRoleEligibilityScheduleRequestTok)
		}

		client = crudClientFactory(&res)
	}

	schedulesClient, err := armauthorization.NewRoleEligibilitySchedulesClient(token, nil)
	if err != nil {
		return nil, err
	}
	requestsClient, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(token, nil)
	if err != nil {
		return nil, err
	}

	return &CustomResource{
		path: pimRoleEligibilityScheduleRequestPath,
		tok:  pimRoleEligibilityScheduleRequestTok,

		Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
			inputsPlain := inputs.Mappable()
			scope := getScope(inputsPlain)
			if scope == "" {
				return nil, false, fmt.Errorf("scope is required")
			}

			roleEligibilityScheduleRequestName, ok := inputsPlain["roleEligibilityScheduleRequestName"]
			if !ok {
				// This is the first Pulumi read. Find the schedule by scope and principalId.
				principalId := inputsPlain["principalId"].(string)
				roleDefinitionId := inputsPlain["roleDefinitionId"].(string)

				// TODO does this ever happen?
				filterScope := scope
				if !strings.HasPrefix(scope, "/") {
					filterScope = "/" + scope
				}

				filter := fmt.Sprintf("principalId eq %s", principalId)
				pager := schedulesClient.NewListForScopePager(scope, &armauthorization.RoleEligibilitySchedulesClientListForScopeOptions{
					Filter: &filter,
				})

				for pager.More() {
					page, err := pager.NextPage(ctx)
					if err != nil {
						return nil, false, err
					}
					for _, schedule := range page.Value {
						props := schedule.Properties
						if props.RoleDefinitionID != nil && strings.EqualFold(*props.RoleDefinitionID, roleDefinitionId) &&
							props.Scope != nil && strings.EqualFold(*props.Scope, filterScope) &&
							props.PrincipalID != nil && strings.EqualFold(*props.PrincipalID, principalId) &&
							props.MemberType != nil && *props.MemberType == armauthorization.MemberTypeDirect {
							// TODO can we take the whole schedule and skip the following read?
							roleEligibilityScheduleRequestName = *schedule.Name
							break
						}
					}
				}
			}

			if roleEligibilityScheduleRequestName == nil {
				return nil, false, nil
			}

			// Get schedule using schedulesClient. Response shape is
			// type RoleEligibilitySchedule struct {
			// // role eligibility schedule properties.
			// Properties *RoleEligibilityScheduleProperties
			// // READ-ONLY; The role eligibility schedule Id.
			// ID *string
			// // READ-ONLY; The role eligibility schedule name.
			// Name *string
			// // READ-ONLY; The role eligibility schedule type.
			// Type *string
			schedule, err := schedulesClient.Get(ctx, scope, roleEligibilityScheduleRequestName.(string), nil)
			if err != nil {
				if azure.IsNotFound(err) {
					return nil, false, nil
				}
				return nil, false, err
			}
			q.Q("Read", id, schedule)

			resultMap, err := structToMap(schedule)
			if err != nil {
				return nil, false, err
			}
			outputs := client.ResponseBodyToSdkOutputs(resultMap)
			q.Q("Read converted", id, outputs)

			return outputs, true, nil
		},

		GetIdAndQuery: func(ctx context.Context, inputs resource.PropertyMap, crudClient crud.ResourceCrudClient) (string, map[string]any, error) {
			return getIdAndQueryForRoleEligibilityScheduleRequest(inputs, crudClient)
		},

		Create: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, error) {
			bodyParams, err := client.PrepareAzureRESTBody(id, inputs)
			if err != nil {
				return nil, err
			}
			q.Q("Create", id, bodyParams)
			queryParams := map[string]any{"api-version": client.ApiVersion()}

			// We could skip the Create if bodyParams = originalState, i.e., the user adds a policy in its default
			// configuration, but we don't have a diff function.
			resp, _, err := client.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				return nil, err
			}

			// TODO Check for OData.Error.Code != "SubjectNotFound"
			// Poll schedules "list for scope" with filter assignedTo=principalId until
			//   item.Properties.RoleDefinitionId == roleDefinitionId && *item.Properties.MemberType == roleeligibilityschedules.MemberTypeDirect
			q.Q("Create", id, resp)

			outputs := client.ResponseBodyToSdkOutputs(resp)
			return outputs, nil
		},

		Update: func(ctx context.Context, id string, news, olds resource.PropertyMap) (map[string]any, error) {
			return nil, nil
		},

		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			inputsPlain := properties.Mappable()
			scope := getScope(inputsPlain)
			if scope == "" {
				return fmt.Errorf("unexpected: scope not found in Delete inputs")
			}
			roleEligibilityScheduleRequestName, ok := inputsPlain["roleEligibilityScheduleRequestName"]
			if !ok {
				return fmt.Errorf("unexpected: roleEligibilityScheduleRequestName not found in Delete inputs")
			}

			// Get schedule request using schedulesClient.
			scheduleRequest, err := requestsClient.Get(ctx, scope, roleEligibilityScheduleRequestName.(string), nil)
			if err != nil {
				if azure.IsNotFound(err) {
					return nil
				}
				return err
			}

			// If the schedule request is still active, cancel it.
			if *scheduleRequest.Properties.Status == armauthorization.StatusPendingAdminDecision ||
				*scheduleRequest.Properties.Status == armauthorization.StatusPendingApproval ||
				*scheduleRequest.Properties.Status == armauthorization.StatusPendingApprovalProvisioning ||
				*scheduleRequest.Properties.Status == armauthorization.StatusPendingEvaluation ||
				*scheduleRequest.Properties.Status == armauthorization.StatusPendingProvisioning {
				_, err = requestsClient.Cancel(ctx, scope, roleEligibilityScheduleRequestName.(string), nil)
				if err != nil {
					err = fmt.Errorf("delete failed to cancel schedule request: %w", err)
				}
				return err
			}

			// If the schedule request is already granted:
			// Create a new Scoped RoleEligibilityScheduleRequest with new GUID and requestType = AdminRemove
			principalId := inputsPlain["principalId"].(string)
			roleDefinitionId := inputsPlain["roleDefinitionId"].(string)

			adminRemove := armauthorization.RequestTypeAdminRemove
			removeRequestUuid := uuid.New().String()
			removeRequest := armauthorization.RoleEligibilityScheduleRequest{
				Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
					RequestType:      &adminRemove,
					PrincipalID:      &principalId,
					RoleDefinitionID: &roleDefinitionId,
					Justification:    pulumi.StringRef("Removed by Pulumi"),
				},
			}
			// TODO repeat request while error code is "ActiveDurationTooShort", until err == nil || error code is "RoleAssignmentDoesNotExist"
			createResp, err := requestsClient.Create(ctx, scope, removeRequestUuid, removeRequest, nil)
			if err != nil {
				q.Q("Delete error", id, err)
				return err
			}
			q.Q("Delete", id, createResp)

			// Done, do not wait for schedule to be removed
			return nil
		},
	}, nil
}

func getScope(inputs map[string]any) string {
	scopeRaw, ok := inputs["scope"]
	if !ok {
		return ""
	}
	return scopeRaw.(string)
}

// getIdAndQueryForRoleEligibilityScheduleRequest replaces the default implementation of crud.ResourceCrudClient.PrepareAzureRESTIdAndQuery.
// It doesn't change queryParams, only the id, to correctly build and encode the "scope" of the request.
func getIdAndQueryForRoleEligibilityScheduleRequest(inputs resource.PropertyMap, crudClient crud.ResourceCrudClient) (string, map[string]any, error) {
	// We want the standard queryParams but will make our own id.
	_, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(inputs)
	if err != nil {
		return "", nil, fmt.Errorf("failed to prepare Azure REST ID and query: %w", err)
	}

	if !inputs.HasValue("scope") {
		return "", nil, fmt.Errorf("GetIdAndQuery: scope is required")
	}
	if !inputs.HasValue("roleEligibilityScheduleRequestName") {
		return "", nil, fmt.Errorf("GetIdAndQuery: roleEligibilityScheduleRequestName is required")
	}

	scope := inputs["scope"].StringValue()
	roleEligibilityScheduleRequestName := inputs["roleEligibilityScheduleRequestName"].StringValue()

	// If the user uses resource.id as an input to scope, which is natural, the scope as required by the Authentication API will be missing
	// the leading "/providers/Microsoft.Subscription". See
	// https://github.com/Azure/azure-rest-api-specs/blob/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/RoleEligibilityScheduleRequest.json#L50.
	// scope = strings.TrimPrefix(scope, "/")
	// scope = strings.TrimSuffix(scope, "/")
	// if strings.HasPrefix(scope, "subscriptions") {
	// 	scope = "providers/Microsoft.Subscription/" + scope
	// }

	id := pimRoleEligibilityScheduleRequestPath
	id = strings.Replace(id, "{scope}", scope, -1)
	id = strings.Replace(id, "{roleEligibilityScheduleRequestName}", roleEligibilityScheduleRequestName, -1)

	return id, queryParams, nil
}

func structToMap(obj any) (map[string]any, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(data, &result)
	return result, err
}
