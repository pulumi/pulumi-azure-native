// Copyright 2016-2024, Pulumi Corporation.

package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersionToVersionPart(t *testing.T) {
	tests := []struct {
		name       string
		apiVersion string
		expected   string
	}{
		{
			name:       "stable version",
			apiVersion: "2024-01-02",
			expected:   "v20240102",
		},
		{
			name:       "preview version",
			apiVersion: "2024-01-02-preview",
			expected:   "v20240102preview",
		},
		{
			name:       "already has v prefix",
			apiVersion: "v2024-01-02",
			expected:   "v20240102",
		},
		{
			name:       "preview with v prefix",
			apiVersion: "v2024-01-02-preview",
			expected:   "v20240102preview",
		},
		{
			name:       "real example from issue",
			apiVersion: "2024-10-02-preview",
			expected:   "v20241002preview",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := apiVersionToVersionPart(tt.apiVersion)
			assert.Equal(t, tt.expected, result)
		})
	}
}
