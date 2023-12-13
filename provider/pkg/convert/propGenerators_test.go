package convert

import (
	"pgregory.net/rapid"
)

func propSimpleValue() *rapid.Generator[interface{}] {
	return rapid.OneOf[interface{}](
		rapid.String().AsAny(),
		rapid.Int().AsAny(),
		rapid.Float64().AsAny(),
		rapid.Bool().AsAny(),
	)
}

func propNonRecursiveMap() *rapid.Generator[map[string]interface{}] {
	gen := rapid.MapOfN(rapid.String(), propSimpleValue(), 1, -1)
	return gen
}

func propNonRecursiveArray() *rapid.Generator[[]interface{}] {
	gen := rapid.SliceOfN(propSimpleValue(), 1, -1)
	return gen
}

func propComplex() *rapid.Generator[any] {
	return rapid.OneOf[interface{}](
		propSimpleValue(),
		propNonRecursiveMap().AsAny(),
		propNonRecursiveArray().AsAny(),
	)
}

func propComplexArray() *rapid.Generator[[]interface{}] {
	return rapid.SliceOfN(propComplex(), 1, -1)
}

func propComplexMap() *rapid.Generator[map[string]interface{}] {
	return rapid.MapOfN(rapid.String(), propComplex(), 1, -1)
}

func propNestedComplex() *rapid.Generator[any] {
	return rapid.OneOf[interface{}](
		propComplex(),
		propComplexArray().AsAny(),
		propComplexMap().AsAny(),
	)
}
