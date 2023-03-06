// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

func TestShouldFlatten(t *testing.T) {
	type testCase struct {
		additionalProperties *spec.SchemaOrBool
		clientFlatten        bool
		isRequired           bool
		expected             bool
	}
	cases := []testCase{
		testCase{
			additionalProperties: nil,
			clientFlatten:        false,
			isRequired:           false,
			expected:             false,
		},
		testCase{
			additionalProperties: nil,
			clientFlatten:        true,
			isRequired:           false,
			expected:             true,
		},
		testCase{
			additionalProperties: nil,
			clientFlatten:        true,
			isRequired:           true,
			expected:             false,
		},
		testCase{
			additionalProperties: &spec.SchemaOrBool{},
			clientFlatten:        true,
			isRequired:           false,
			expected:             false,
		},
		testCase{
			additionalProperties: &spec.SchemaOrBool{},
			clientFlatten:        true,
			isRequired:           true,
			expected:             false,
		},
	}

	for _, tc := range cases {
		assert.Equal(t,
			tc.expected,
			shouldFlatten("#/dummy", tc.additionalProperties, tc.clientFlatten, tc.isRequired),
			tc)
	}
}
