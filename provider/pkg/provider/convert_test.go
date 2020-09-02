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

package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var p *azurermProvider
var sampleSdkProps map[string]interface{}
var sampleApiPackage map[string]interface{}

func init() {
	p = &azurermProvider{
		resourceMap: &AzureApiMetadata{
			Types: map[string]AzureApiType{
				"azurerm:testing:Structure": {
					Properties: map[string]AzureApiProperty{
						"v1": {},
						"v2": {},
						"v3-odd": {
							SdkName: "v3",
						},
						"v4-nested": {
							SdkName:   "v4",
							Container: "props",
						},
					},
				},
				"azurerm:testing:More": {
					Properties: map[string]AzureApiProperty{
						"items": {
							Items: &AzureApiProperty{
								Ref: "#/types/azurerm:testing:MoreItem",
							},
						},
					},
				},
				"azurerm:testing:MoreItem": {
					Properties: map[string]AzureApiProperty{
						"aaa": {
							SdkName: "Aaa",
						},
						"bbb": {
							Container: "ccc",
						},
					},
				},
			},
			Resources: map[string]AzureApiResource{
				"r1": {
					PutParameters: []AzureApiParameter{
						{
							Body: &AzureApiType{
								Properties: map[string]AzureApiProperty{
									"name": {},
									"x-threshold": {
										SdkName: "threshold",
									},
									"structure": {
										Ref: "#/types/azurerm:testing:Structure",
									},
									"p1": {
										Container: "properties",
									},
									"p2": {
										Container: "properties",
									},
									"more": {
										Container: "properties",
										Ref:       "#/types/azurerm:testing:More",
									},
									"tags": {},
								},
							},
						},
					},
					Response: map[string]AzureApiProperty{
						"name": {},
						"x-threshold": {
							SdkName: "threshold",
						},
						"structure": {
							Ref: "#/types/azurerm:testing:Structure",
						},
						"p1": {
							Container: "properties",
						},
						"p2": {
							Container: "properties",
						},
						"more": {
							Container: "properties",
							Ref:       "#/types/azurerm:testing:More",
						},
						"tags": {},
					},
				},
			},
		},
	}

	sampleApiPackage = map[string]interface{}{
		"name":        "MyResource",
		"x-threshold": 123,
		"structure": map[string]interface{}{
			"v1":     "value1",
			"v2":     2,
			"v3-odd": "odd-value",
			"props": map[string]interface{}{
				"v4-nested": true,
			},
		},
		"properties": map[string]interface{}{
			"p1": "prop1",
			"p2": "prop2",
			"more": map[string]interface{}{
				"items": []interface{}{
					map[string]interface{}{"aaa": "111", "ccc": map[string]interface{}{"bbb": "333"}},
					map[string]interface{}{"aaa": "222"},
				},
			},
		},
		"tags": map[string]interface{}{
			"createdBy":   "admin",
			"application": "dashboard",
		},
	}
	sampleSdkProps = map[string]interface{}{
		"name":      "MyResource",
		"threshold": 123,
		"structure": map[string]interface{}{
			"v1": "value1",
			"v2": 2,
			"v3": "odd-value",
			"v4": true,
		},
		"p1": "prop1",
		"p2": "prop2",
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"Aaa": "111", "bbb": "333"},
				map[string]interface{}{"Aaa": "222"},
			},
		},
		"tags": map[string]interface{}{
			"createdBy":   "admin",
			"application": "dashboard",
		},
	}
}

func TestResponseToSdkOutputs(t *testing.T) {
	outputs := p.responseToSdkOutputs(p.resourceMap.Resources["r1"].Response, sampleApiPackage)
	assert.Equal(t, outputs, sampleSdkProps)
}

func TestSdkPropertiesToRequest(t *testing.T) {
	bodyProperties := p.resourceMap.Resources["r1"].PutParameters[0].Body.Properties
	data := p.sdkPropertiesToRequest(bodyProperties, sampleSdkProps)
	assert.Equal(t, data, sampleApiPackage)
}
