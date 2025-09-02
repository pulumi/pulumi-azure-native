// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByName(t *testing.T) {
	for _, tc := range []struct {
		name     string
		expected Configuration
		ok       bool
	}{
		{name: "public", expected: AzurePublic, ok: true},
		{name: "AzureCloud", expected: AzurePublic, ok: true},
		{name: "china", expected: AzureChina, ok: true},
		{name: "azurechinacloud", expected: AzureChina, ok: true},
		{name: "AzureChinaCloud", expected: AzureChina, ok: true},
		{name: "usgov", expected: AzureGovernment, ok: true},
		{name: "usgovernment", expected: AzureGovernment, ok: true},
		{name: "azureusgovernment", expected: AzureGovernment, ok: true},
		{name: "AzureUSGovernment", expected: AzureGovernment, ok: true},
		{name: "azureusgovernmentcloud", expected: AzureGovernment, ok: true},
		{name: "AzureUSGovernmentCloud", expected: AzureGovernment, ok: true},
		{name: "MyCloud", expected: AzurePublic, ok: false},
	} {
		actual, ok := FromName(tc.name)
		assert.Equal(t, tc.ok, ok, tc.name)
		assert.Equal(t, tc.expected, actual, tc.name)
	}
}
