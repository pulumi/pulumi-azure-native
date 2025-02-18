// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPendingStatus(t *testing.T) {
	t.Parallel()

	pending := []armauthorization.Status{
		armauthorization.StatusGranted,
		armauthorization.StatusPendingAdminDecision,
		armauthorization.StatusPendingApproval,
		armauthorization.StatusPendingApprovalProvisioning,
		armauthorization.StatusPendingEvaluation,
		armauthorization.StatusPendingExternalProvisioning,
		armauthorization.StatusPendingProvisioning,
		armauthorization.StatusPendingRevocation,
		armauthorization.StatusPendingScheduleCreation,
	}

	t.Run("pending status", func(t *testing.T) {
		t.Parallel()

		for _, status := range pending {
			assert.True(t, statusIsPending(&status), status)
		}
	})

	t.Run("terminal status", func(t *testing.T) {
		t.Parallel()

		for _, status := range armauthorization.PossibleStatusValues() {
			if slices.Contains(pending, status) {
				continue
			}
			assert.False(t, statusIsPending(&status), status)
		}
	})
}

type fakePimEligibilityScheduleClient struct {
	scheduleToFind    *armauthorization.RoleEligibilitySchedule
	scheduleFindError error

	curScheduleIndex int
	schedule         func(numCall int) (*armauthorization.RoleEligibilitySchedule, error)

	cancelCalls int
	cancelError error

	createError    error
	capturedCreate *struct {
		scope      string
		name       string
		parameters armauthorization.RoleEligibilityScheduleRequest
	}
}

func (c *fakePimEligibilityScheduleClient) findSchedule(ctx context.Context, scope, principalID, roleDefinitionID string) (*armauthorization.RoleEligibilitySchedule, error) {
	if c.scheduleFindError != nil {
		return nil, c.scheduleFindError
	}
	return c.scheduleToFind, nil
}

func (c *fakePimEligibilityScheduleClient) getSchedule(ctx context.Context, scope, scheduleName string) (*armauthorization.RoleEligibilitySchedule, error) {
	s, err := c.schedule(c.curScheduleIndex)
	c.curScheduleIndex++
	return s, err
}

func (c *fakePimEligibilityScheduleClient) createSchedule(ctx context.Context, scope string, roleEligibilityScheduleRequestName string,
	parameters armauthorization.RoleEligibilityScheduleRequest,
) (armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse, error) {
	c.capturedCreate = &struct {
		scope      string
		name       string
		parameters armauthorization.RoleEligibilityScheduleRequest
	}{
		scope:      scope,
		name:       roleEligibilityScheduleRequestName,
		parameters: parameters,
	}

	return armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse{}, c.createError
}

func (c *fakePimEligibilityScheduleClient) cancelSchedule(ctx context.Context, scope, scheduleName string) (armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse, error) {
	c.cancelCalls++
	if c.cancelError != nil {
		return armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse{}, c.cancelError
	}
	return armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse{}, nil
}

func (c *fakePimEligibilityScheduleClient) mapScheduleToOutputs(schedule *armauthorization.RoleEligibilitySchedule) (map[string]any, error) {
	return map[string]any{}, nil
}

func TestCreatePimEligibilitySchedule(t *testing.T) {
	t.Parallel()

	pimRoleEligibilityScheduleTickerInterval = 1 * time.Second

	validInputs := resource.PropertyMap{
		"scope":            resource.NewStringProperty("scope"),
		"principalId":      resource.NewStringProperty("principalID"),
		"roleDefinitionId": resource.NewStringProperty("roleDefinitionID"),
	}

	t.Run("happy path", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{
			scheduleToFind: &armauthorization.RoleEligibilitySchedule{
				Properties: &armauthorization.RoleEligibilityScheduleProperties{
					PrincipalID:      pulumi.StringRef("principalID"),
					RoleDefinitionID: pulumi.StringRef("roleDefinitionID"),
				},
			},
		}
		_, err := createPimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 5*time.Second)
		require.NoError(t, err)
		require.NotNil(t, client.capturedCreate)
		assert.True(t, isValidGUID(client.capturedCreate.name), client.capturedCreate.name)
		assert.Equal(t, "scope", client.capturedCreate.scope)
		require.NotNil(t, client.capturedCreate.parameters.Properties.PrincipalID)
		assert.Equal(t, "principalID", *client.capturedCreate.parameters.Properties.PrincipalID)
		require.NotNil(t, client.capturedCreate.parameters.Properties.RoleDefinitionID)
		assert.Equal(t, "roleDefinitionID", *client.capturedCreate.parameters.Properties.RoleDefinitionID)
	})

	t.Run("invalid input", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{}
		_, err := createPimEligibilitySchedule(context.Background(), "/id/is/not/used", resource.PropertyMap{}, client, 4*time.Second)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "is required")
	})

	t.Run("error creating schedule", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{
			createError: errors.New("error creating schedule"),
		}
		_, err := createPimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 4*time.Second)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error creating schedule")
	})

	t.Run("error polling for schedule", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{
			scheduleFindError: errors.New("error finding schedule"),
		}
		_, err := createPimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 4*time.Second)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error finding schedule")
	})

	t.Run("timeout polling for schedule", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{}
		_, err := createPimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 1*time.Second)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "timed out")
	})
}

func TestDeletePimEligibilitySchedule(t *testing.T) {
	t.Parallel()

	pimRoleEligibilityScheduleTickerInterval = 1 * time.Second

	guid := uuid.New().String()

	createScheduleData := func(status armauthorization.Status) (resource.PropertyMap, *armauthorization.RoleEligibilitySchedule) {
		validInputs := resource.PropertyMap{
			"name":             resource.NewStringProperty(guid),
			"scope":            resource.NewStringProperty("scope"),
			"principalId":      resource.NewStringProperty("principalID"),
			"roleDefinitionId": resource.NewStringProperty("roleDefinitionID"),
			"status":           resource.NewStringProperty(string(status)),
		}

		schedule := &armauthorization.RoleEligibilitySchedule{
			Name: pulumi.StringRef(guid),
			Properties: &armauthorization.RoleEligibilityScheduleProperties{
				PrincipalID:      pulumi.StringRef("principalID"),
				RoleDefinitionID: pulumi.StringRef("roleDefinitionID"),
				Scope:            pulumi.StringRef("scope"),
				Status:           &status,
			},
		}
		return validInputs, schedule
	}

	t.Run("happy path, pending schedule", func(t *testing.T) {
		t.Parallel()

		validInputs, schedule := createScheduleData(armauthorization.StatusPendingApproval)
		client := &fakePimEligibilityScheduleClient{
			schedule: func(_ int) (*armauthorization.RoleEligibilitySchedule, error) { return schedule, nil },
		}
		err := deletePimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 5*time.Second)
		require.NoError(t, err)
		assert.Equal(t, 1, client.cancelCalls)
		assert.Nil(t, client.capturedCreate)
	})

	t.Run("happy path, active schedule", func(t *testing.T) {
		t.Parallel()

		validInputs, schedule := createScheduleData(armauthorization.StatusScheduleCreated)
		client := &fakePimEligibilityScheduleClient{
			// On the third polling request, the schedule is gone.
			schedule: func(numCall int) (*armauthorization.RoleEligibilitySchedule, error) {
				if numCall < 3 {
					return schedule, nil
				}
				return nil, &azcore.ResponseError{StatusCode: 404}
			},
		}
		err := deletePimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 5*time.Second)
		require.NoError(t, err)
		assert.Equal(t, 0, client.cancelCalls)
		require.NotNil(t, client.capturedCreate) // deletion is a new request with type=remove
		assert.Equal(t, armauthorization.RequestTypeAdminRemove, *client.capturedCreate.parameters.Properties.RequestType)
		assert.Equal(t, "scope", client.capturedCreate.scope)
		assert.True(t, isValidGUID(client.capturedCreate.name), client.capturedCreate.name)
	})

	t.Run("active schedule, deletion times out", func(t *testing.T) {
		t.Parallel()

		validInputs, schedule := createScheduleData(armauthorization.StatusScheduleCreated)
		client := &fakePimEligibilityScheduleClient{
			schedule: func(_ int) (*armauthorization.RoleEligibilitySchedule, error) { return schedule, nil },
		}
		err := deletePimEligibilitySchedule(context.Background(), "/id/is/not/used", validInputs, client, 3*time.Second)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "timed out")
		assert.Equal(t, 0, client.cancelCalls)
		require.NotNil(t, client.capturedCreate)
		assert.Equal(t, "scope", client.capturedCreate.scope)
		assert.True(t, isValidGUID(client.capturedCreate.name), client.capturedCreate.name)
		assert.Equal(t, armauthorization.RequestTypeAdminRemove, *client.capturedCreate.parameters.Properties.RequestType)
	})
}

// isValidGUID checks if the provided string is a valid GUID/UUID.
func isValidGUID(guid string) bool {
	_, err := uuid.Parse(guid)
	return err == nil
}

func TestUpdateResourceDescription(t *testing.T) {
	t.Parallel()

	t.Run("no existing description", func(t *testing.T) {
		t.Parallel()

		resource := &ResourceDefinition{
			Resource: schema.ResourceSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "",
				},
			},
		}

		updateResourceDescription(resource, "foo bar")
		assert.Equal(t, "foo bar", resource.Resource.Description)
	})

	t.Run("existing description without API version", func(t *testing.T) {
		t.Parallel()

		resource := &ResourceDefinition{
			Resource: schema.ResourceSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "existing description",
				},
			},
		}

		updateResourceDescription(resource, "foo bar")
		assert.Equal(t, "foo bar", resource.Resource.Description)
	})

	t.Run("existing description with API versions", func(t *testing.T) {
		t.Parallel()

		resource := &ResourceDefinition{
			Resource: schema.ResourceSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: `Role Eligibility schedule request
Azure REST API version: 2020-10-01.

Other available API versions: 2020-10-01-preview, 2022-04-01-preview, 2024-02-01-preview, 2024-09-01-preview.
## Import
...`,
				},
			},
		}

		updateResourceDescription(resource, "foo bar.\n\nmore info.")
		assert.Equal(t, `foo bar.

more info.

Azure REST API version: 2020-10-01.

Other available API versions: 2020-10-01-preview, 2022-04-01-preview, 2024-02-01-preview, 2024-09-01-preview.
## Import
...`, resource.Resource.Description)
	})
}

func TestInputsToSdk(t *testing.T) {
	t.Parallel()

	validInputs := resource.PropertyMap{
		"name":             resource.NewStringProperty(uuid.New().String()),
		"scope":            resource.NewStringProperty("scope"),
		"principalId":      resource.NewStringProperty("principalID"),
		"roleDefinitionId": resource.NewStringProperty("roleDefinitionID"),
	}

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		sdk, err := inputsToSdk(validInputs)
		require.NoError(t, err)
		assert.Equal(t, "scope", *sdk.Properties.Scope)
		assert.Equal(t, "principalID", *sdk.Properties.PrincipalID)
		assert.Equal(t, "roleDefinitionID", *sdk.Properties.RoleDefinitionID)
	})

	t.Run("required fields missing", func(t *testing.T) {
		t.Parallel()

		for _, prop := range []string{"scope", "principalId", "roleDefinitionId"} {
			inputs := validInputs.Copy()
			delete(inputs, resource.PropertyKey(prop))
			_, err := inputsToSdk(inputs)
			require.Error(t, err)
			assert.Contains(t, err.Error(), prop+" is required")
		}
	})
}
