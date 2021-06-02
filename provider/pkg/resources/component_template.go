// Copyright 2016-2020, Pulumi Corporation.
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

package resources

import "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

func newArmTemplate() *CustomResource {
	return &CustomResource{
		tok: "azure-native:extensions:ArmTemplate",
		Schema: &schema.ResourceSpec{
			InputProperties: map[string]schema.PropertySpec{
				"resourceGroupName": {
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
				"content": {
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
			},
			RequiredInputs: []string{"resourceGroupName", "content"},
			IsComponent:    true,
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"result": {
						TypeSpec: schema.TypeSpec{Type: "string"},
					},
				},
				Required: []string{"result"},
			},
		},
	}
}
