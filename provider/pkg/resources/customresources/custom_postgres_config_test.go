// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetValueFromAzureResponse(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		resp := map[string]any{
			"properties": map[string]any{
				"defaultValue": "foo",
			},
		}

		defaultValue := getValueFromAzureResponse(resp, "defaultValue")
		require.NotNil(t, defaultValue)
		require.Equal(t, "foo", defaultValue)
	})

	t.Run("empty response", func(t *testing.T) {
		require.Nil(t, getValueFromAzureResponse(map[string]any{}, "defaultValue"))
	})

	t.Run("nil response", func(t *testing.T) {
		require.Nil(t, getValueFromAzureResponse(nil, "defaultValue"))
	})

	t.Run("missing property", func(t *testing.T) {
		require.Nil(t, getValueFromAzureResponse(map[string]any{
			"properties": map[string]any{},
		}, "defaultValue"))
	})
}
