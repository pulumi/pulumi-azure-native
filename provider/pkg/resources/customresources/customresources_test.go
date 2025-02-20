// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsSingleton(t *testing.T) {
	assert.False(t, IsSingleton(""))
	assert.False(t, IsSingleton("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}"))

	postgresResource, err := postgresFlexibleServerConfiguration(nil, nil)
	require.NoError(t, err)
	assert.True(t, postgresResource.isSingleton)
	assert.True(t, IsSingleton(postgresResource.path))
}

func TestIncludeCustomResource(t *testing.T) {
	t.Run("Include any API version for WebApp", func(t *testing.T) {
		isCustom, include := IncludeCustomResource(webAppPath, "1888-11-11")
		assert.True(t, isCustom)
		assert.True(t, include)
	})

	t.Run("Include specific API version for Role Eligibility Schedule Request", func(t *testing.T) {
		isCustom, include := IncludeCustomResource(PimRoleEligibilityScheduleRequestPath, "2020-10-01")
		assert.True(t, isCustom)
		assert.True(t, include)
	})

	t.Run("Don't include other API versions for Role Eligibility Schedule Request", func(t *testing.T) {
		isCustom, include := IncludeCustomResource(PimRoleEligibilityScheduleRequestPath, "2024-07-01")
		assert.True(t, isCustom)
		assert.False(t, include)
	})

	t.Run("Non-custom resource", func(t *testing.T) {
		isCustom, include := IncludeCustomResource(
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}",
			"2020-10-01")
		assert.False(t, isCustom)
		assert.False(t, include)
	})
}
