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
