package gen

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestEnumExtensionCaseCollapsing(t *testing.T) {
	generator := moduleGenerator{
		pkg: &schema.PackageSpec{
			Name:  "azure-native",
			Types: map[string]schema.ComplexTypeSpec{},
		},
		module: "MyModule",
	}
	// Doesn't return the enum - it's added to the Types map
	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{},
	}
	enumExtensions := map[string]interface{}{"name": "myProp", "values": []interface{}{
		map[string]interface{}{"value": "foo", "name": "Foo"},
		map[string]interface{}{"value": "Foo", "name": "Foo"},
		map[string]interface{}{"value": "bar", "name": "Bar"},
	}}
	_, err := generator.genEnumType(schema, &openapi.ReferenceContext{}, enumExtensions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, generator.pkg.Types, 1)
	assert.Len(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum, 2)
	assert.Equal(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum[0].Value, "foo")
	assert.Equal(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum[1].Value, "bar")
}

func TestEnumNonExtensionCaseCollapsing(t *testing.T) {
	generator := moduleGenerator{
		pkg: &schema.PackageSpec{
			Name:  "azure-native",
			Types: map[string]schema.ComplexTypeSpec{},
		},
		module: "MyModule",
	}
	// Doesn't return the enum - it's added to the Types map
	_, err := generator.genEnumType(&spec.Schema{
		SchemaProps: spec.SchemaProps{
			Enum: []interface{}{
				"foo",
				"Foo",
				"bar",
			},
		},
	}, &openapi.ReferenceContext{}, map[string]interface{}{"name": "myProp"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, generator.pkg.Types, 1)
	assert.Len(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum, 2)
	assert.Equal(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum[0].Value, "foo")
	assert.Equal(t, generator.pkg.Types["azure-native:MyModule:MyProp"].Enum[1].Value, "bar")
}
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
