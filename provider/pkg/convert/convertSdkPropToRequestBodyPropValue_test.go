package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestNoChangeWithNoPropSchema(t *testing.T) {
	types := map[string]resources.AzureAPIType{}
	prop := resources.AzureAPIProperty{}
	t.Run("SDK to API", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			id := rapid.String().Draw(t, "id")
			value := propComplex().Draw(t, "value")
			converted := testSdkToApi(types, id, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
	t.Run("API to SDK", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			value := propComplex().Draw(t, "value")
			converted := testApiToSdk(types, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
	t.Run("Output to Input", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			value := propComplex().Draw(t, "value")
			converted := testOutputToInput(types, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
}

func TestMapWithKnownType(t *testing.T) {
	id := "id"
	types := map[string]resources.AzureAPIType{
		"myType": {
			Properties: map[string]resources.AzureAPIProperty{
				"prop1": {
					Type: "string",
				},
			},
		},
	}
	prop := resources.AzureAPIProperty{
		Type: "object",
		Ref:  "#/types/myType",
	}
	t.Run("SDK to API", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			value := map[string]interface{}{
				"prop1": propComplex().Draw(t, "prop1"),
			}
			converted := testSdkToApi(types, id, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
	t.Run("API to SDK", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			value := propComplex().Draw(t, "value")
			converted := testApiToSdk(types, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
	t.Run("Output to Input", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			value := propComplex().Draw(t, "value")
			converted := testOutputToInput(types, &prop, value)
			assert.Equal(t, value, converted)
		})
	})
}

func testSdkToApi(types map[string]resources.AzureAPIType, id string, prop *resources.AzureAPIProperty, value interface{}) interface{} {
	c := NewSdkShapeConverterFull(types)
	return c.convertSdkPropToRequestBodyPropValue(id, prop, value)
}

func testApiToSdk(types map[string]resources.AzureAPIType, prop *resources.AzureAPIProperty, value interface{}) interface{} {
	c := NewSdkShapeConverterFull(types)
	return c.convertBodyPropToSdkPropValue(prop, value)
}

func testOutputToInput(types map[string]resources.AzureAPIType, prop *resources.AzureAPIProperty, value interface{}) interface{} {
	c := NewSdkShapeConverterFull(types)
	return c.convertOutputToInputPropValue(prop, value)
}

func propSimpleValue() *rapid.Generator[interface{}] {
	return rapid.OneOf[interface{}](
		rapid.String().AsAny(),
		rapid.Int().AsAny(),
		rapid.Float64().AsAny(),
		rapid.Bool().AsAny(),
	)
}

func propNonRecursiveMap() *rapid.Generator[map[string]interface{}] {
	gen := rapid.MapOf(rapid.String(), propSimpleValue())
	return gen
}

func propNonRecursiveArray() *rapid.Generator[[]interface{}] {
	gen := rapid.SliceOf(propSimpleValue())
	return gen
}

func propComplex() *rapid.Generator[any] {
	return rapid.OneOf[interface{}](
		propSimpleValue(),
		propNonRecursiveMap().AsAny(),
		propNonRecursiveArray().AsAny(),
	)
}
