package gen

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// mergeTypes checks that two type specs are allowed to be represented as a single schema type.
// Refs are not followed, but are just checked for equality as the referred type will then be checked independently.
// If you face an error coming from this function for a new API change, consider changing the typeNameOverride map
// or adjusting the way the top-level resources are projected in versioner.go.
func mergeTypes(t1 schema.ComplexTypeSpec, t2 schema.ComplexTypeSpec, isOutput bool) (*schema.ComplexTypeSpec, error) {
	if t1.Type != t2.Type {
		return nil, errors.Errorf("types do not match: %s vs %s", t1.Type, t2.Type)
	}

	if !isOutput {
		// Check that every required property of T1 and T2 exists in T1 and has the same type (for intputs only).
		t1Required := codegen.NewStringSet(t1.Required...)
		t2Required := codegen.NewStringSet(t2.Required...)
		t1Only := t1Required.Subtract(t2Required)
		t2Only := t2Required.Subtract(t1Required)
		if len(t1Only) != 0 || len(t2Only) != 0 {
			var requiredErrors []string
			if len(t1Only) != 0 {
				t1Fmt := strings.Join(t1Only.SortedValues(), ",")
				requiredErrors = append(requiredErrors, fmt.Sprintf("only required in A: %s", t1Fmt))
			}
			if len(t2Only) != 0 {
				t2Fmt := strings.Join(t2Only.SortedValues(), ",")
				requiredErrors = append(requiredErrors, fmt.Sprintf("only required in B: %s", t2Fmt))
			}
			return nil, errors.Errorf("required properties do not match: %s", strings.Join(requiredErrors, "; "))
		}
	}

	if t1.Properties == nil && t2.Properties == nil {
		return &t1, nil
	}
	mergedProperties := map[string]pschema.PropertySpec{}
	for name, p := range t1.Properties {
		mergedProperties[name] = p
	}
	for name, p2 := range t2.Properties {
		p1, found := mergedProperties[name]
		if !found {
			mergedProperties[name] = p2
			continue
		}
		mergedProperty, err := mergePropertySpec(p1, p2)
		if err != nil {
			return nil, err
		}
		mergedProperties[name] = *mergedProperty
	}

	merged := t1
	merged.Properties = mergedProperties
	return &merged, nil
}

func mergePropertySpec(p1, p2 pschema.PropertySpec) (*pschema.PropertySpec, error) {
	mergedTypeSpec, err := mergeTypeSpec(p1.TypeSpec, p2.TypeSpec)
	if err != nil {
		return nil, err
	}
	p1.TypeSpec = *mergedTypeSpec
	return &p1, nil
}

func mergeTypeSpec(t1, t2 schema.TypeSpec) (*schema.TypeSpec, error) {
	// Simple case: both types are the same primitive.
	if isPrimitiveType(t1) && isPrimitiveType(t2) {
		if t1.Type == t2.Type {
			return &t1, nil
		}
		if t1.Type == "integer" && t2.Type == "number" || t1.Type == "number" && t2.Type == "integer" {
			// If one is an integer and the other is a number, the merged type is number.
			return &schema.TypeSpec{Type: "number"}, nil
		}
	}

	if t1.Ref != "" && t2.Ref != "" {
		// Simple case: both types are the same ref.
		if t1.Ref == t2.Ref {
			return &t1, nil
		}
		// Refs don't match - fail with error.
		return nil, errors.Errorf("refs do not match: %s vs %s", t1.Ref, t2.Ref)
	}

	if t1.Type == "array" && t2.Type == "array" {
		contract.Assertf(t1.Items != nil && t2.Items != nil, "Type %v is missing items (other: %v)", t1, t2)
		// Both are arrays - merge the element types.
		items, err := mergeTypeSpec(*t1.Items, *t2.Items)
		if err != nil {
			return nil, err
		}
		t1.Items = items
		return &t1, nil
	}

	if t1.Type == "object" && t2.Type == "object" {
		contract.Assertf(t1.AdditionalProperties != nil, "Type %v is missing additionalProperties (other: %v)", t1, t2)
		contract.Assertf(t2.AdditionalProperties != nil, "Type %v is missing additionalProperties (other: %v)", t2, t1)
		// Both are objects - merge the properties.
		merged, err := mergeTypeSpec(*t1.AdditionalProperties, *t2.AdditionalProperties)
		if err != nil {
			return nil, err
		}
		t1.AdditionalProperties = merged
		return &t1, nil
	}

	// One is a oneOf - combine the other into it.
	if t1.OneOf != nil || t2.OneOf != nil {
		oneOf, other := t1, t2
		if t1.OneOf == nil {
			oneOf, other = t2, t1
		}

		return unionOneOfTypes(oneOf, other)
	}
	// Fail with error
	return nil, errors.Errorf("types not mergable: %v vs %v", t1, t2)
}

func isPrimitiveType(t schema.TypeSpec) bool {
	return t.Type == "string" || t.Type == "integer" || t.Type == "number" || t.Type == "boolean"
}

// unionOneOfTypes checks that two oneOf types are allowed to be represented as a single schema type.
func unionOneOfTypes(oneOf, other schema.TypeSpec) (*schema.TypeSpec, error) {
	contract.Assertf(oneOf.OneOf != nil, "Type %v is not a oneOf", oneOf)
	if isPrimitiveType(other) {
		// Check if type exists in oneOf.
		for _, t := range oneOf.OneOf {
			if t.Type == other.Type {
				return &oneOf, nil
			}
		}
		// Not found - add it.
		oneOf.OneOf = append(oneOf.OneOf, other)
		return &oneOf, nil
	}
	if other.Ref != "" {
		// Check if ref exists in oneOf.
		for _, t := range oneOf.OneOf {
			if t.Ref == other.Ref {
				return &oneOf, nil
			}
		}
		// Not found - add it.
		oneOf.OneOf = append(oneOf.OneOf, other)
		return &oneOf, nil
	}
	if other.OneOf != nil {
		bothHaveDiscriminators := oneOf.Discriminator != nil && other.Discriminator != nil
		onlyOneHasDiscriminator := (oneOf.Discriminator == nil && other.Discriminator != nil) || (oneOf.Discriminator != nil && other.Discriminator == nil)
		if onlyOneHasDiscriminator || (bothHaveDiscriminators && oneOf.Discriminator.PropertyName != other.Discriminator.PropertyName) {
			return nil, errors.Errorf("cannot union oneOf with different discriminators: %v vs %v", oneOf, other)
		}
		if oneOf.Discriminator != nil {
			// Union the discriminators.
			for k, otherRef := range other.Discriminator.Mapping {
				if oneOfRef, found := oneOf.Discriminator.Mapping[k]; found && oneOfRef != otherRef {
					return nil, errors.Errorf("cannot union oneOf with different discriminators: %v vs %v", oneOf, other)
				}
				oneOf.Discriminator.Mapping[k] = otherRef
			}
		}
		// Union each type in other oneOf with the first oneOf.
		for _, t := range other.OneOf {
			merged, err := unionOneOfTypes(oneOf, t)
			if err != nil {
				return nil, err
			}
			oneOf = *merged
		}

		return &oneOf, nil
	}
	return nil, errors.Errorf("cannot union oneOf with specified type %v", other)
}

// func intersectOneOfTypes(oneOf, other schema.TypeSpec) (*schema.TypeSpec, error) {
// 	contract.Assert(oneOf.OneOf != nil)
// 	if isPrimitiveType(other) {
// 		// Check if type exists in oneOf.
// 		for _, t := range oneOf.OneOf {
// 			if t.Type == other.Type {
// 				return &other, nil
// 			}
// 		}
// 		// Not found - return error.
// 		return nil, errors.Errorf("cannot intersect oneOf with specified type %v", other)
// 	}
// 	if other.Ref != "" {
// 		// Check if ref exists in oneOf.
// 		for _, t := range oneOf.OneOf {
// 			if t.Ref == other.Ref {
// 				return &other, nil
// 			}
// 		}
// 		// Not found - return error.
// 		return nil, errors.Errorf("cannot intersect oneOf with specified type %v", other)
// 	}
// 	if other.OneOf != nil {
// 		// Find common types in other oneOf and the first oneOf.
// 		common := []schema.TypeSpec{}
// 		otherByKey := make(map[string]schema.TypeSpec)
// 		for _, t := range other.OneOf {
// 			otherByKey[oneOfTypeKey(t)] = t
// 		}
// 		for _, t := range oneOf.OneOf {
// 			if _, ok := otherByKey[oneOfTypeKey(t)]; ok {
// 				common = append(common, t)
// 			}
// 		}
// 		if len(common) == 0 {
// 			return nil, errors.Errorf("cannot intersect oneOf with specified type %v", other)
// 		}
// 		if len(common) == 1 {
// 			return &common[0], nil
// 		}
// 		oneOf.OneOf = common
// 		return &oneOf, nil
// 	}
// 	return nil, errors.Errorf("cannot union oneOf with specified type %v", other)
// }
//
// func oneOfTypeKey(t schema.TypeSpec) string {
// 	// convert to json
// 	b, err := json.Marshal(t)
// 	contract.Assert(err == nil)
// 	return string(b)
// }
