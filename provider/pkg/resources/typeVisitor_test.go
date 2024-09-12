// Copyright 2024, Pulumi Corporation.  All rights reserved.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestVisitPackageSpecTypes(t *testing.T) {
	pkg := &schema.PackageSpec{
		Resources: map[string]schema.ResourceSpec{
			"azure-native:test:MyResource": {
				InputProperties: map[string]schema.PropertySpec{
					"input1": {
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/azure-native:test:MyType1",
						},
					},
				},
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						"output1": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: "#/types/azure-native:test:MyType2",
								},
							},
						},
					},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{
			"azure-native:test:MyFunction": {
				Inputs: &schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						"input1": {
							TypeSpec: schema.TypeSpec{
								Type: "object",
								AdditionalProperties: &schema.TypeSpec{
									Ref: "#/types/azure-native:test:MyType3",
								},
							},
						},
					},
				},
				Outputs: &schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						"output1": {
							TypeSpec: schema.TypeSpec{
								OneOf: []schema.TypeSpec{
									{
										Ref: "#/types/azure-native:test:MyType4",
									},
									{
										Ref: "#/types/azure-native:test:MyType5",
									},
								},
							},
						},
					},
				},
			},
		},
		Types: map[string]schema.ComplexTypeSpec{
			"azure-native:test:MyType1":             {},
			"azure-native:test:MyType2":             {},
			"azure-native:test:MyType3":             {},
			"azure-native:test:MyType4":             {},
			"azure-native:test:MyType5":             {},
			"azure-native:test:MyTypeNotReferenced": {},
		},
	}

	visitedTypes := make(map[string]bool)
	visitor := func(tok string, t schema.ComplexTypeSpec) {
		visitedTypes[tok] = true
	}

	VisitPackageSpecTypes(pkg, visitor)

	assert.Len(t, visitedTypes, 5)
	for tok := range visitedTypes {
		assert.Contains(t, pkg.Types, tok)
	}
	assert.NotContains(t, visitedTypes, "azure-native:test:MyTypeNotReferenced")
}
