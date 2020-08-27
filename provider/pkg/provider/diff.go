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
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
)

// calculateDetailedDiff produced a property diff for a given object diff and a resource definition. It inspects
// the schema of the resource to find out if the requested diff can be performed in-place or requires a replacement.
func calculateDetailedDiff(resource *AzureApiResource, diff *resource.ObjectDiff) (map[string]*rpc.PropertyDiff, error) {
	replaceKeys := codegen.NewStringSet()

	for _, p := range resource.PutParameters {
		// All the parameters that are part of the resource path cause a replacement.
		if p.Location == "path" {
			name := p.Name
			if p.Value != nil && p.Value.SdkName != "" {
				name = p.Value.SdkName
			}
			replaceKeys.Add(name)
		}

		// Force New on resource properties also cause a replacement.
		if p.Location == "body" {
			for propName, prop := range p.Body.Properties {
				if prop.ForceNew {
					name := propName
					if prop.SdkName != "" {
						name = prop.SdkName
					}
					replaceKeys.Add(name)
				}
			}

		}
	}

	d := differ{replaceKeys: replaceKeys}
	return d.calculateDetailedDiff(diff, "")
}

type differ struct {
	replaceKeys codegen.StringSet
}

func (d *differ) calculateDetailedDiff(diff *resource.ObjectDiff, base string) (map[string]*rpc.PropertyDiff, error) {
	detailedDiff := map[string]*rpc.PropertyDiff{}
	for k, v := range diff.Updates {
		key := base + string(k)
		if v.Object == nil {
			kind := rpc.PropertyDiff_UPDATE
			if d.replaceKeys.Has(key) {
				kind = rpc.PropertyDiff_UPDATE_REPLACE
			}
			detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
		} else {
			subDiff, err := d.calculateDetailedDiff(v.Object, key+".")
			if err != nil {
				return nil, err
			}

			for sk, sv := range subDiff {
				detailedDiff[sk] = sv
			}
		}
	}
	for k, _ := range diff.Adds {
		key := base + string(k)
		kind := rpc.PropertyDiff_ADD
		if d.replaceKeys.Has(key) {
			kind = rpc.PropertyDiff_ADD_REPLACE
		}
		detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
	}
	for k, _ := range diff.Deletes {
		key := base + string(k)
		kind := rpc.PropertyDiff_DELETE
		if d.replaceKeys.Has(key) {
			kind = rpc.PropertyDiff_DELETE_REPLACE
		}
		detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
	}

	return detailedDiff, nil
}
