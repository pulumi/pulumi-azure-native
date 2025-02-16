// Copyright 2025, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

	scheduleToGet    *armauthorization.RoleEligibilitySchedule
	scheduleGetError error

	cancelError error

	createError    error
	capturedCreate struct {
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
	if c.scheduleGetError != nil {
		return nil, c.scheduleGetError
	}
	return c.scheduleToGet, nil
}

func (c *fakePimEligibilityScheduleClient) createSchedule(ctx context.Context, scope string, roleEligibilityScheduleRequestName string,
	parameters armauthorization.RoleEligibilityScheduleRequest,
) (armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse, error) {
	c.capturedCreate.scope = scope
	c.capturedCreate.name = roleEligibilityScheduleRequestName
	c.capturedCreate.parameters = parameters

	return armauthorization.RoleEligibilityScheduleRequestsClientCreateResponse{}, c.createError
}

func (c *fakePimEligibilityScheduleClient) cancelSchedule(ctx context.Context, scope, scheduleName string) (armauthorization.RoleEligibilityScheduleRequestsClientCancelResponse, error) {
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

	pimRoleEligibilityScheduleTickerInterval = 2 * time.Second

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
		assert.Contains(t, err.Error(), "timed out waiting")
	})
}

func TestDeletePimEligibilitySchedule(t *testing.T) {
	t.Parallel()

	pimRoleEligibilityScheduleTickerInterval = 2 * time.Second

	guid := uuid.New().String()

	status := armauthorization.StatusGranted
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

	t.Run("happy path", func(t *testing.T) {
		client := &fakePimEligibilityScheduleClient{
			scheduleToGet: schedule,
		}
		err := deletePimEligibilitySchedule(context.Background(), "/id/is/not/used", resource.PropertyMap{}, validInputs, client, 5*time.Second)
		require.NoError(t, err)
	})
}

// isValidGUID checks if the provided string is a valid GUID/UUID.
func isValidGUID(guid string) bool {
	_, err := uuid.Parse(guid)
	return err == nil
}
