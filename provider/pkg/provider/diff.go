// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/davegardnerisme/deephash"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

const body = "body"

func ignoreKey(key resource.PropertyKey) bool {
	return strings.HasPrefix(string(key), "__")
}

// mapProperties traverses the given APIResource and creates a map from all properties, including
// in referenced types, to their property spec. The map key is a JSONPath to the property.
func mapProperties(res resources.AzureAPIResource, lookupType resources.TypeLookupFunc) map[string]resources.AzureAPIProperty {
	properties := map[string]resources.AzureAPIProperty{}
	for _, p := range res.PutParameters {
		if p.Location == body && p.Body != nil {
			resources.TraverseProperties(p.Body.Properties, lookupType, func(name string, prop resources.AzureAPIProperty, path []string) {
				path = append(path, name)
				jsonPath := strings.Join(path, ".")
				properties[jsonPath] = prop
			})
		}
	}
	return properties
}

// diff is the provider's main entry point to diffing resources.
func diff(lookupType resources.TypeLookupFunc,
	res resources.AzureAPIResource,
	oldInputs, newInputs resource.PropertyMap) map[string]*rpc.PropertyDiff {

	properties := mapProperties(res, lookupType)

	// This JSONPath will be updated as we traverse the inputs, to look up properties in the above
	// `properties` map.
	path := ""

	diff := objectDiff(properties, oldInputs, newInputs, path)
	if diff == nil {
		return nil
	}

	// Calculate the detailed diff object containing information about replacements.
	return calculateDetailedDiff(&res, lookupType, diff)
}

func objectDiff(properties map[string]resources.AzureAPIProperty,
	oldInputs, newInputs resource.PropertyMap,
	path string) *resource.ObjectDiff {

	adds := make(resource.PropertyMap)
	deletes := make(resource.PropertyMap)
	sames := make(resource.PropertyMap)
	updates := make(map[resource.PropertyKey]resource.ValueDiff)

	addToPath := func(p string, key resource.PropertyKey) string {
		if p == "" {
			return string(key)
		}
		return p + "." + string(key)
	}

	// First find any updates or deletes.
	for k, old := range oldInputs {
		if ignoreKey(k) {
			continue
		}

		if new, has := newInputs[k]; has {
			// If a new exists, use it; for output properties, however, ignore differences.
			if new.IsOutput() {
				sames[k] = old
			} else if diff := valueDiff(properties, old, new, addToPath(path, k)); diff != nil {
				if !old.HasValue() {
					adds[k] = new
				} else if !new.HasValue() {
					deletes[k] = old
				} else {
					updates[k] = *diff
				}
			} else {
				sames[k] = old
			}
		} else if old.HasValue() {
			// If there was no new property, it has been deleted.
			deletes[k] = old
		}
	}

	// Next find any additions not in the old map.
	for k, new := range newInputs {
		if ignoreKey(k) {
			continue
		}

		if _, has := oldInputs[k]; !has && new.HasValue() {
			adds[k] = new
		}
	}

	// If no diffs were found, return nil; else return a diff structure.
	if len(adds) == 0 && len(deletes) == 0 && len(updates) == 0 {
		return nil
	}

	return &resource.ObjectDiff{
		Adds:    adds,
		Deletes: deletes,
		Sames:   sames,
		Updates: updates,
	}
}

func valueDiff(properties map[string]resources.AzureAPIProperty,
	old, new resource.PropertyValue,
	path string) *resource.ValueDiff {
	if old.IsArray() && new.IsArray() {
		oldArr := old.ArrayValue()
		newArr := new.ArrayValue()

		prop, ok := properties[path]
		if ok && len(prop.ArrayIdentifiers) > 0 {
			return diffKeyedArrays(properties, prop.ArrayIdentifiers, oldArr, newArr, path)
		}

		// If any elements exist in the new array but not the old, track them as adds.
		adds := make(map[int]resource.PropertyValue)
		for i := len(oldArr); i < len(newArr); i++ {
			adds[i] = newArr[i]
		}
		// If any elements exist in the old array but not the new, track them as adds.
		deletes := make(map[int]resource.PropertyValue)
		for i := len(newArr); i < len(oldArr); i++ {
			deletes[i] = oldArr[i]
		}
		// Now if elements exist in both, track them as sames or updates.
		sames := make(map[int]resource.PropertyValue)
		updates := make(map[int]resource.ValueDiff)
		for i := 0; i < len(oldArr) && i < len(newArr); i++ {
			if diff := valueDiff(properties, oldArr[i], newArr[i], path); diff != nil {
				updates[i] = *diff
			} else {
				sames[i] = oldArr[i]
			}
		}

		if len(adds) == 0 && len(deletes) == 0 && len(updates) == 0 {
			return nil
		}
		return &resource.ValueDiff{
			Old: old,
			New: new,
			Array: &resource.ArrayDiff{
				Adds:    adds,
				Deletes: deletes,
				Sames:   sames,
				Updates: updates,
			},
		}
	}
	if old.IsObject() && new.IsObject() {
		oldObj := old.ObjectValue()
		newObj := new.ObjectValue()
		if diff := objectDiff(properties, oldObj, newObj, path); diff != nil {
			return &resource.ValueDiff{
				Old:    old,
				New:    new,
				Object: diff,
			}
		}
		return nil
	}

	// If we got here, either the values are primitives, or they weren't the same type; do a simple diff.
	if old.DeepEquals(new) {
		return nil
	}
	return &resource.ValueDiff{Old: old, New: new}
}

// calculateDetailedDiff produced a property diff for a given object diff and a resource definition. It inspects
// the schema of the resource to find out if the requested diff can be performed in-place or requires a replacement.
func calculateDetailedDiff(resource *resources.AzureAPIResource, lookupType resources.TypeLookupFunc,
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
					if typ, has, err := lookupType(prop.Ref); has && err == nil {
						findForceNew(propName+".", typ.Properties, replaceKeys)
					}
				}
				// Arrays of objects.
				if prop.Items != nil && prop.Items.Ref != "" {
					if typ, has, err := lookupType(prop.Items.Ref); has && err == nil {
						findForceNew(propName+"[].", typ.Properties, replaceKeys)
					}
				}
			}
		}
	}

	applyAzureSpecificDiff(diff)

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
// Note that the use case for it is rather specific to the particular Read implementation:
// - values are inputs currently stored in state before Read was called
// - diff is a calculated difference between old and new _outputs_
// So, diff values are independent of the values map. Practically, diff values may contain
// more properties in object maps, and that doesn't necessarily mean we need to copy all of
// them. Instead, we want to copy only the changed properties, and do so recursively.
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
	for key, update := range diff.Updates {
		if value, has := values[key]; has {
			result[key] = applyValueDiff(value, update.Old, update.New)
		} else {
			result[key] = update.New
		}
	}
	return result
}

func applyValueDiff(baseValue resource.PropertyValue, oldValue resource.PropertyValue,
	newValue resource.PropertyValue) resource.PropertyValue {

	// Objects are compared property-by-property recursively and we only copy the changed sub-properties back.
	if baseValue.IsObject() && oldValue.IsObject() && newValue.IsObject() {
		subDiff := oldValue.ObjectValue().Diff(newValue.ObjectValue())
		result := applyDiff(baseValue.ObjectValue(), subDiff)
		return resource.NewObjectProperty(result)
	}

	// Arrays are compared element-by-element recursively.
	if baseValue.IsArray() && oldValue.IsArray() && newValue.IsArray() {
		baseValueArray := baseValue.ArrayValue()
		oldValueArray := oldValue.ArrayValue()
		newValueArray := newValue.ArrayValue()
		length := len(newValueArray)
		result := make([]resource.PropertyValue, length)

		for i, el := range newValueArray {
			if i >= len(baseValueArray) || i >= len(oldValueArray) {
				// A new element was added, copy as-is.
				result[i] = el
				continue
			}
			// Calculate the diff recursively for each element.
			result[i] = applyValueDiff(baseValueArray[i], oldValueArray[i], newValueArray[i])
		}
		return resource.NewArrayProperty(result)
	}

	// For plain values, just use the new baseValue.
	return newValue
}

// applyAzureSpecificDiff modifies a generic diff calculated by the Platform with any
// Azure-specific diffing adjustments.
func applyAzureSpecificDiff(diff *resource.ObjectDiff) {
	updates := map[resource.PropertyKey]resource.ValueDiff{}
	for k, v := range diff.Updates {
		// Apply special diffing logic to some well-known top-level properties.
		switch string(k) {
		case "resourceGroupName":
			// "resourceGroupName" is case-insensitive.
			if v.Old.IsString() && v.New.IsString() {
				if strings.EqualFold(strings.ToLower(v.Old.StringValue()), v.New.StringValue()) {
					continue
				}
			}
		case "location":
			// "location" is case- and spaces-insensitive.
			if v.Old.IsString() && v.New.IsString() {
				if normalizedLocation(v.Old.StringValue()) == normalizedLocation(v.New.StringValue()) {
					continue
				}
			}
		case "sku":
			// Another special case is "sku" in AKS clusters.
			if v.Old.IsObject() && v.New.IsObject() {
				if sameManagedClusterSku(v.Old.ObjectValue(), v.New.ObjectValue()) {
					continue
				}
			}
		}
		updates[k] = v
	}
	diff.Updates = updates
}

// normalizedLocation converts Azure location values of a format like "West US 2" and
// "WestUS2" to the lowercase and no-space format of "westus2".
func normalizedLocation(location string) string {
	return strings.ToLower(strings.ReplaceAll(location, " ", ""))
}

// sameManagedClusterSku checks whether two property maps representing a SKU are
// equivalent in terms of AKS nomenclature.
// See https://github.com/pulumi/pulumi-azure-native/issues/2600
func sameManagedClusterSku(oldMap resource.PropertyMap, newMap resource.PropertyMap) bool {
	// Expect exactly two keys: name and tier - in both maps.
	if len(oldMap) != 2 || len(newMap) != 2 {
		return false
	}

	// Check that 'name' exists in both.
	oldName, hasOld := oldMap["name"]
	newName, hasNew := newMap["name"]
	if !hasOld || !hasNew || !oldName.IsString() || !newName.IsString() {
		return false
	}

	// Check that 'tier' exists in both.
	oldTier, hasOld := oldMap["tier"]
	newTier, hasNew := newMap["tier"]
	if !hasOld || !hasNew || !oldTier.IsString() || !newTier.IsString() {
		return false
	}

	// Check that name is exactly the same or both are Basic or Base.
	sameName := oldName.StringValue() == newName.StringValue() ||
		(oldName.StringValue() == "Basic" && newName.StringValue() == "Base") ||
		(oldName.StringValue() == "Base" && newName.StringValue() == "Basic")
	// Check that tier is exactly the same or both are Paid or Standard.
	sameTier := oldTier.StringValue() == newTier.StringValue() ||
		(oldTier.StringValue() == "Paid" && newTier.StringValue() == "Standard") ||
		(oldTier.StringValue() == "Standard" && newTier.StringValue() == "Paid")
	// Return true if both of the above hold.
	return sameName && sameTier
}

// calculateChangesAndReplacements compares a property diff with the old and new inputs and the
// previous state. It returns a list of properties that will change, and a list of properties
// that will be replaced. In some cases, entries of the diff may not lead to any change or
// replacement, e.g., new properties set to their default value are not considered changes.
func calculateChangesAndReplacements(
	detailedDiff map[string]*rpc.PropertyDiff,
	oldInputs, newInputs, oldState resource.PropertyMap,
	res resources.AzureAPIResource) ([]string, []string) {

	var changes, replaces []string
	for k, v := range detailedDiff {
		switch v.Kind {
		case rpc.PropertyDiff_ADD_REPLACE:
			// Special case: previously, the property input had no value but is now set to X.
			// If the output contains this property and it's X, that's not a replacement.
			// Workaround for https://github.com/pulumi/pulumi-azure-nextgen/issues/238
			// TODO: remove this block before GA.
			key := resource.PropertyKey(k)
			_, hasOldInput := oldInputs[key]
			newInputValue, hasNewInput := newInputs[key]
			outputValue, hasOutput := oldState[key]
			if !hasOldInput && hasNewInput && hasOutput && reflect.DeepEqual(newInputValue, outputValue) {
				v.Kind = rpc.PropertyDiff_ADD
			} else {
				replaces = append(replaces, k)
			}

		case rpc.PropertyDiff_ADD:
			key := resource.PropertyKey(k)
			_, hasOldInput := oldInputs[key]
			newInputValue, hasNewInput := newInputs[key]
			if !hasOldInput && hasNewInput {
				// If this is a new property that is merely being initialized with its default
				// value, we don't need to show a noisy diff.
				prop, found := res.LookupProperty(k)
				if found && prop.Default != nil && reflect.DeepEqual(newInputValue.V, prop.Default) {
					logging.V(9).Infof("Skipping diff for %s, property with default value %v is added", k, newInputValue)
					continue
				}
			}

		case rpc.PropertyDiff_DELETE_REPLACE, rpc.PropertyDiff_UPDATE_REPLACE:
			replaces = append(replaces, k)
		}

		parts := strings.Split(k, ".")
		changes = append(changes, parts[0])
		v.InputDiff = true
	}
	return changes, replaces
}

// arrayElement is a helper struct to track the index of a property value in an array.
type arrayElement struct {
	element resource.PropertyValue
	index   int
}

// diffKeyedArrays compares two arrays that are annotated in the API spec with the extension
// "x-ms-identifiers" and behave more like sets. The combined `keys` form a unique identifier for
// the array elements, so order doesn't matter.
func diffKeyedArrays(properties map[string]resources.AzureAPIProperty,
	keys []string,
	old, new []resource.PropertyValue,
	path string) *resource.ValueDiff {

	adds := make(map[int]resource.PropertyValue)
	deletes := make(map[int]resource.PropertyValue)
	sames := make(map[int]resource.PropertyValue)
	updates := make(map[int]resource.ValueDiff)

	sortedKeys := sort.StringSlice(keys) // for stable map keys

	oldIdValues := map[string]arrayElement{}
	for i, oldItem := range old {
		hash := hashObject(oldItem, sortedKeys)
		if hash != "" {
			oldIdValues[hash] = arrayElement{element: oldItem, index: i}
		}
	}

	newSeen := map[string]struct{}{}
	for i, newItem := range new {
		hash := hashObject(newItem, sortedKeys)
		if hash == "" {
			continue
		}

		newSeen[hash] = struct{}{}

		oldItem, ok := oldIdValues[hash]
		if !ok {
			adds[i] = newItem
		} else {
			diff := valueDiff(properties, oldItem.element, newItem, path)
			if diff == nil {
				sames[i] = newItem
			} else {
				updates[i] = *diff
			}
		}
	}

	// Check for olds that are gone
	for oldHash, oldEntry := range oldIdValues {
		if _, ok := newSeen[oldHash]; !ok {
			deletes[oldEntry.index] = oldEntry.element
		}
	}

	return &resource.ValueDiff{
		Old: resource.NewPropertyValue(old),
		New: resource.NewPropertyValue(new),
		Array: &resource.ArrayDiff{
			Adds:    adds,
			Deletes: deletes,
			Sames:   sames,
			Updates: updates,
		},
	}
}

// hashObject computes a deep hash of an object value using the object's property values indexed by
// the given keys. Returns the empty string if the value is not an object.
func hashObject(val resource.PropertyValue, sortedKeys []string) string {
	if !val.IsObject() {
		logging.V(5).Infof("WARNING: diffKeyedArray: item %v is not an object\n", val)
		return ""
	}

	obj := val.ObjectValue()
	idValues := make([]any, len(sortedKeys))
	for i, key := range sortedKeys {
		idValues[i] = obj[resource.PropertyKey(key)]
	}

	return string(deephash.Hash(idValues)[:])
}
