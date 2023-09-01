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
