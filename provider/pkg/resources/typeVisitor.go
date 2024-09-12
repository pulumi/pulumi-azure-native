// Copyright 2024, Pulumi Corporation.  All rights reserved.

package resources

import (
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// VisitPackageSpecTypes navigates all resources and functions, searching for all schema types that they reference.
// It then calls the visitor callback for each type found.
func VisitPackageSpecTypes(pkg *pschema.PackageSpec, visitor func(tok string, t pschema.ComplexTypeSpec)) {
	seen := codegen.Set{}
	for _, r := range pkg.Resources {
		for _, p := range r.InputProperties {
			visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
		}
		for _, p := range r.Properties {
			visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
		}
	}
	for _, f := range pkg.Functions {
		if f.Inputs != nil {
			for _, p := range f.Inputs.Properties {
				visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
			}
		}
		if f.Outputs != nil {
			for _, p := range f.Outputs.Properties {
				visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
			}
			if f.ReturnType != nil && f.ReturnType.TypeSpec != nil {
				if f.ReturnType.TypeSpec != nil {
					visitComplexTypes(pkg.Types, *f.ReturnType.TypeSpec, visitor, seen)
				} else if f.ReturnType.ObjectTypeSpec != nil {
					for _, p := range f.ReturnType.ObjectTypeSpec.Properties {
						visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
					}
				}
			}
		}
	}
}

func VisitResourceTypes(pkg *pschema.PackageSpec, resourceToken string, visitor func(tok string, t pschema.ComplexTypeSpec)) {
	seen := codegen.Set{}
	resource, ok := pkg.Resources[resourceToken]
	if !ok {
		return
	}
	for _, p := range resource.InputProperties {
		visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
	}
	for _, p := range resource.Properties {
		visitComplexTypes(pkg.Types, p.TypeSpec, visitor, seen)
	}
}

func visitComplexTypes(types map[string]pschema.ComplexTypeSpec, t pschema.TypeSpec, visitor func(tok string, t pschema.ComplexTypeSpec), seen codegen.Set) {
	if strings.HasPrefix(t.Ref, "#/types/azure-native:") {
		typeName := strings.TrimPrefix(t.Ref, "#/types/")

		other, ok := types[typeName]
		if ok {
			visitor(typeName, other)
			if !seen.Has(typeName) {
				seen.Add(typeName)
				for _, p := range other.ObjectTypeSpec.Properties {
					visitComplexTypes(types, p.TypeSpec, visitor, seen)
				}
			}
		}
	}

	if t.AdditionalProperties != nil {
		visitComplexTypes(types, *t.AdditionalProperties, visitor, seen)
	}

	if t.Items != nil {
		visitComplexTypes(types, *t.Items, visitor, seen)
	}

	for _, other := range t.OneOf {
		visitComplexTypes(types, other, visitor, seen)
	}
}
