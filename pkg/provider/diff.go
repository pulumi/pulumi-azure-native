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
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
	"strings"
)

// calculateDetailedDiff produced a property diff for a given diff and a resource definition. It inspects the schema
// of the resource to find out if the requested diff can be performed in-place or requires a replacement.
func calculateDetailedDiff(resource Resource, diff *resource.ObjectDiff) (map[string]*rpc.PropertyDiff, error) {
	spec, err := openapi.NewSpec(resource.swagggerSpecLocation)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load spec %s", resource.swagggerSpecLocation)
	}
	path := spec.Paths.Paths[resource.path]

	// If a property is a part of the PATCH payload, we consider it mutable.
	var patchSchema *openapi.Schema
	if path.Patch != nil {
		for _, param := range path.Patch.Parameters {
			if param.In == "body" {
				param, err := spec.ResolveParameter(param)
				if err != nil {
					return nil, err
				}

				if param.Schema == nil {
					break
				}

				patchSchema, err = param.ResolveSchema(param.Schema)
				if err != nil {
					return nil, err
				}
				break
			}
		}
	}

	d := differ { patchSchema: patchSchema}
	return d.calculateDetailedDiff(diff, "")
}

type differ struct {
	patchSchema *openapi.Schema
}

func (d *differ) calculateDetailedDiff(diff *resource.ObjectDiff, base string) (map[string]*rpc.PropertyDiff, error) {
	detailedDiff := map[string]*rpc.PropertyDiff{}
	for k, v := range diff.Updates {
		key := base + string(k)
		if v.Object == nil {
			kind := rpc.PropertyDiff_UPDATE_REPLACE
			pb, err := d.isPatchable(key)
			if err != nil {
				return nil, err
			}

			if pb {
				kind = rpc.PropertyDiff_UPDATE
			}
			detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
		} else {
			subDiff, err := d.calculateDetailedDiff(v.Object, key + ".")
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
		kind := rpc.PropertyDiff_ADD_REPLACE
		pb, err := d.isPatchable(key)
		if err != nil {
			return nil, err
		}

		if pb {
			kind = rpc.PropertyDiff_ADD
		}
		detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
	}
	for k, _ := range diff.Deletes {
		key := base + string(k)
		kind := rpc.PropertyDiff_DELETE_REPLACE
		pb, err := d.isPatchable(key)
		if err != nil {
			return nil, err
		}

		if pb {
			kind = rpc.PropertyDiff_DELETE
		}
		detailedDiff[key] = &rpc.PropertyDiff{Kind: kind}
	}

	return detailedDiff, nil
}

// isPatchable checks if a property with a given name belongs to the PATCH method payload.
func (d *differ) isPatchable(name string) (bool, error) {
	return d.existsInSchema(name, d.patchSchema)
}

// existsInSchema checks whether a property with the given path (e.g. `foo.bar`) exists in the given schema.
func (d *differ) existsInSchema(propertyPath string, schema *openapi.Schema) (bool, error) {
	if schema == nil {
		return false, nil
	}

	parts := strings.Split(propertyPath, ".")

	for k, v := range schema.Properties {
		if k == parts[0] {
			// If the path has no dots in it, we found the match.
			if len(parts) == 1 {
				return true, nil
			}

			// Otherwise, resolve the child schema.
			propertySchema, err := schema.ResolveSchema(&v)
			if err != nil {
				return false, err
			}

			// If child schema is not an object with properties (e.g. we are looking at a dictionary of tags),
			// we found the match.
			if len(propertySchema.Properties) == 0 {
				return true, nil
			}

			// Inspect the child properties recursively.
			newPath := strings.Join(parts[1:], ".")
			subPatchable, err := d.existsInSchema(newPath, propertySchema)
			if err != nil {
				return false, err
			}

			return subPatchable, nil
		}
	}

	// If the main schema doesn't contain the target property, check if it's in AllOf schemas.
	for _, s := range schema.AllOf {
		allOfSchema, err := schema.ResolveSchema(&s)
		if err != nil {
			return false, err
		}

		basePatchable, err := d.existsInSchema(propertyPath, allOfSchema)
		if err != nil {
			return false, err
		}

		if basePatchable {
			return true, nil
		}
	}

	return false, nil
}
