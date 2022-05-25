// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

const body = "body"

// calculateDetailedDiff produced a property diff for a given object diff and a resource definition. It inspects
// the schema of the resource to find out if the requested diff can be performed in-place or requires a replacement.
func calculateDetailedDiff(resource *resources.AzureAPIResource, types resources.PartialMap[resources.AzureAPIType],
	diff *resource.ObjectDiff) map[string]*rpc.PropertyDiff {
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
		if p.Location == body && p.Body != nil {
			// Top-level properties.
			findForceNew("", p.Body.Properties, replaceKeys)
			for propName, prop := range p.Body.Properties {
				// Object types.
				if prop.Ref != "" {
					typName := strings.TrimPrefix(prop.Ref, "#/types/")
					if typ, has, err := types.Get(typName); has && err != nil {
						findForceNew(propName+".", typ.Properties, replaceKeys)
					}
				}
				// Arrays of objects.
				if prop.Items != nil && prop.Items.Ref != "" {
					typName := strings.TrimPrefix(prop.Items.Ref, "#/types/")
					if typ, has, err := types.Get(typName); has && err != nil {
						findForceNew(propName+"[].", typ.Properties, replaceKeys)
					}
				}
			}
		}
	}

	d := differ{replaceKeys: replaceKeys}
	return d.calculateObjectDiff(diff, "", "")
}

func findForceNew(base string, props map[string]resources.AzureAPIProperty, replaceKeys codegen.StringSet) {
	for propName, prop := range props {
		if prop.ForceNew {
			name := propName
			if prop.SdkName != "" {
				name = prop.SdkName
			}
			replaceKeys.Add(base + name)
		}
	}
}

type differ struct {
	replaceKeys codegen.StringSet
}

func (d *differ) calculateObjectDiff(diff *resource.ObjectDiff, diffBase, replaceBase string) map[string]*rpc.PropertyDiff {
	detailedDiff := map[string]*rpc.PropertyDiff{}
	for k, v := range diff.Updates {
		key := diffBase + string(k)
		subDiff := d.calculateValueDiff(&v, key, replaceBase+string(k))
		for sk, sv := range subDiff {
			detailedDiff[sk] = sv
		}
	}
	for k := range diff.Adds {
		diffKey := diffBase + string(k)
		replaceKey := replaceBase + string(k)
		kind := rpc.PropertyDiff_ADD
		if d.replaceKeys.Has(replaceKey) {
			kind = rpc.PropertyDiff_ADD_REPLACE
		}
		detailedDiff[diffKey] = &rpc.PropertyDiff{Kind: kind}
	}
	for k := range diff.Deletes {
		diffKey := diffBase + string(k)
		replaceKey := replaceBase + string(k)
		kind := rpc.PropertyDiff_DELETE
		if d.replaceKeys.Has(replaceKey) {
			kind = rpc.PropertyDiff_DELETE_REPLACE
		}
		detailedDiff[diffKey] = &rpc.PropertyDiff{Kind: kind}
	}

	return detailedDiff
}

func (d *differ) calculateValueDiff(v *resource.ValueDiff, diffBase, replaceBase string) map[string]*rpc.PropertyDiff {
	detailedDiff := map[string]*rpc.PropertyDiff{}
	switch {
	case v.Object != nil:
		subDiff := d.calculateObjectDiff(v.Object, diffBase+".", replaceBase+".")
		for sk, sv := range subDiff {
			if sv.Kind == rpc.PropertyDiff_UPDATE && d.replaceKeys.Has(replaceBase) {
				// If the parent property causes a replacement, all child properties cause a replacement.
				sv.Kind = rpc.PropertyDiff_UPDATE_REPLACE
			}
			detailedDiff[sk] = sv
		}
	case v.Array != nil:
		for idx, item := range v.Array.Updates {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			subDiff := d.calculateValueDiff(&item, key, replaceBase+"[]")
			for sk, sv := range subDiff {
				detailedDiff[sk] = sv
			}
		}
		for idx := range v.Array.Adds {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			detailedDiff[key] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_ADD}
		}
		for idx := range v.Array.Deletes {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			detailedDiff[key] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_DELETE}
		}
	default:
		kind := rpc.PropertyDiff_UPDATE
		if d.replaceKeys.Has(replaceBase) {
			kind = rpc.PropertyDiff_UPDATE_REPLACE
		}
		detailedDiff[diffBase] = &rpc.PropertyDiff{Kind: kind}
	}
	return detailedDiff
}

// applyDiff produces a new map as a merge of a calculated diff into an existing map of values.
func applyDiff(values resource.PropertyMap, diff *resource.ObjectDiff) resource.PropertyMap {
	if diff == nil {
		return values
	}

	result := resource.PropertyMap{}
	for name, value := range values {
		if !diff.Deleted(name) {
			result[name] = value
		}
	}
	for key, value := range diff.Adds {
		result[key] = value
	}
	for key, value := range diff.Updates {
		result[key] = value.New
	}
	return result
}
