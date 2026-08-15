// Copyright 2026, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeDeleteAzureClient wraps MockAzureClient to allow injecting a configurable error from Delete,
// which MockAzureClient doesn't support.
type fakeDeleteAzureClient struct {
	azure.MockAzureClient
	deleteErr error

	deleteId         string
	deleteApiVersion string
	deleteAsyncStyle string
}

func (c *fakeDeleteAzureClient) Delete(ctx context.Context, id, apiVersion, asyncStyle string, queryParams map[string]any) error {
	c.deleteId = id
	c.deleteApiVersion = apiVersion
	c.deleteAsyncStyle = asyncStyle
	return c.deleteErr
}

func TestDefaultRuleSchemaHasNoRuleNameInput(t *testing.T) {
	custom, err := defaultRule(nil, nil, &azure.MockAzureClient{})
	require.NoError(t, err)

	require.NotNil(t, custom.Schema)
	def, err := custom.Schema(nil)
	require.NoError(t, err)
	require.NotNil(t, def)

	_, hasRuleName := def.Resource.InputProperties["ruleName"]
	assert.False(t, hasRuleName, "DefaultRule should not accept a ruleName input")

	for _, required := range def.Resource.RequiredInputs {
		assert.NotEqual(t, "ruleName", required)
	}

	assert.True(t, def.MetaResource.Singleton, "the $Default rule always exists once its subscription does")
	assert.Contains(t, def.MetaResource.Path, "/rules/$Default")
	for _, param := range def.MetaResource.PutParameters {
		assert.NotEqual(t, "ruleName", param.Name)
	}
}

func TestDefaultRuleDeleteToleratesNotFound(t *testing.T) {
	azureClient := &fakeDeleteAzureClient{
		deleteErr: &azcore.ResponseError{StatusCode: http.StatusNotFound, ErrorCode: "EntityNotFound"},
	}
	custom, err := defaultRule(nil, nil, azureClient)
	require.NoError(t, err)

	id := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.ServiceBus/namespaces/ns/topics/topic/subscriptions/sub1/rules/$Default"
	deleteErr := custom.Delete(context.Background(), id, nil, nil)

	assert.NoError(t, deleteErr, "deleting an already-gone $Default rule should not be an error")
	assert.Equal(t, id, azureClient.deleteId)
}

func TestDefaultRuleDeletePropagatesOtherErrors(t *testing.T) {
	azureClient := &fakeDeleteAzureClient{
		deleteErr: errors.New("some other error"),
	}
	custom, err := defaultRule(nil, nil, azureClient)
	require.NoError(t, err)

	deleteErr := custom.Delete(context.Background(), "id", nil, nil)
	assert.EqualError(t, deleteErr, "some other error")
}

func TestDefaultRuleDeleteSucceeds(t *testing.T) {
	azureClient := &fakeDeleteAzureClient{}
	custom, err := defaultRule(nil, nil, azureClient)
	require.NoError(t, err)

	deleteErr := custom.Delete(context.Background(), "id", nil, nil)
	assert.NoError(t, deleteErr)
}

func TestReadDefaultRuleWithRetry(t *testing.T) {
	origDelay := defaultRuleReadRetryDelay
	defaultRuleReadRetryDelay = time.Millisecond
	defer func() { defaultRuleReadRetryDelay = origDelay }()

	notFoundErr := &azcore.ResponseError{StatusCode: http.StatusNotFound, ErrorCode: "SubscriptionNotFound"}

	t.Run("succeeds immediately", func(t *testing.T) {
		calls := 0
		response, found, err := readDefaultRuleWithRetry(context.Background(), func(ctx context.Context) (map[string]any, error) {
			calls++
			return map[string]any{"filterType": "SqlFilter"}, nil
		})
		require.NoError(t, err)
		assert.True(t, found)
		assert.Equal(t, map[string]any{"filterType": "SqlFilter"}, response)
		assert.Equal(t, 1, calls)
	})

	t.Run("succeeds after transient not-found", func(t *testing.T) {
		calls := 0
		_, found, err := readDefaultRuleWithRetry(context.Background(), func(ctx context.Context) (map[string]any, error) {
			calls++
			if calls < defaultRuleReadMaxAttempts {
				return nil, notFoundErr
			}
			return map[string]any{}, nil
		})
		require.NoError(t, err)
		assert.True(t, found)
		assert.Equal(t, defaultRuleReadMaxAttempts, calls)
	})

	t.Run("gives up after persistent not-found, no error", func(t *testing.T) {
		calls := 0
		_, found, err := readDefaultRuleWithRetry(context.Background(), func(ctx context.Context) (map[string]any, error) {
			calls++
			return nil, notFoundErr
		})
		require.NoError(t, err)
		assert.False(t, found)
		assert.Equal(t, defaultRuleReadMaxAttempts, calls)
	})

	t.Run("propagates non-404 errors immediately, no retry", func(t *testing.T) {
		calls := 0
		otherErr := errors.New("boom")
		_, found, err := readDefaultRuleWithRetry(context.Background(), func(ctx context.Context) (map[string]any, error) {
			calls++
			return nil, otherErr
		})
		assert.False(t, found)
		assert.Equal(t, otherErr, err)
		assert.Equal(t, 1, calls)
	})
}

// fakeGetAzureClient wraps MockAzureClient to allow injecting a configurable error from Get,
// which MockAzureClient doesn't support, and to record the id it was called with.
type fakeGetAzureClient struct {
	azure.MockAzureClient
	getErr error
	getID  string
}

func (c *fakeGetAzureClient) Get(ctx context.Context, id string, apiVersion string, queryParams map[string]any) (map[string]any, error) {
	c.getID = id
	if c.getErr != nil {
		return nil, c.getErr
	}
	return map[string]any{}, nil
}

func TestDefaultRuleNotFoundOutcome(t *testing.T) {
	ruleID := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.ServiceBus/namespaces/ns/topics/topic/subscriptions/sub1/rules/$Default"
	subscriptionID := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.ServiceBus/namespaces/ns/topics/topic/subscriptions/sub1"

	t.Run("subscription still exists: errors instead of reporting deletion", func(t *testing.T) {
		azureClient := &fakeGetAzureClient{}
		_, found, err := defaultRuleNotFoundOutcome(context.Background(), azureClient, "2024-01-01", ruleID)
		assert.False(t, found)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "parent subscription still exists")
		assert.Equal(t, subscriptionID, azureClient.getID)
	})

	t.Run("subscription also gone: reports deletion with no error", func(t *testing.T) {
		azureClient := &fakeGetAzureClient{
			getErr: &azcore.ResponseError{StatusCode: http.StatusNotFound, ErrorCode: "ResourceNotFound"},
		}
		_, found, err := defaultRuleNotFoundOutcome(context.Background(), azureClient, "2024-01-01", ruleID)
		assert.NoError(t, err)
		assert.False(t, found)
	})

	t.Run("checking the subscription errors: propagates that error", func(t *testing.T) {
		otherErr := errors.New("boom")
		azureClient := &fakeGetAzureClient{getErr: otherErr}
		_, found, err := defaultRuleNotFoundOutcome(context.Background(), azureClient, "2024-01-01", ruleID)
		assert.False(t, found)
		assert.Equal(t, otherErr, err)
	})
}
