// Copyright 2016-2021, Pulumi Corporation.
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

package gen

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func Test_mergeTypes(t *testing.T) {
	type args struct {
		t1       schema.ComplexTypeSpec
		t2       schema.ComplexTypeSpec
		isOutput bool
	}
	tests := []struct {
		name    string
		args    args
		want    schema.ComplexTypeSpec
		wantErr bool
	}{
		{
			name: "object and enum",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: schema.TypeSpec{
							Type: "number",
						},
					},
				}),
				t2: schema.ComplexTypeSpec{
					ObjectTypeSpec: schema.ObjectTypeSpec{
						Type: "string",
					},
					Enum: []schema.EnumValueSpec{
						{Value: "foo"},
					},
				},
			},
			wantErr: true,
		},

		{
			name: "integer and number becomes number",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: schema.TypeSpec{
							Type: "number",
						},
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: schema.TypeSpec{
							Type: "integer",
						},
					},
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: schema.TypeSpec{
						Type: "number",
					},
				},
			}),
		},

		{
			name: "includes both distinct properties",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"bar": stringPropertySpec(),
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": stringPropertySpec(),
				"bar": stringPropertySpec(),
			}),
		},

		{
			name: "fails if requiredness differs for inputs",
			args: args{
				t1: requiredObjectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
				isOutput: false,
			},
			wantErr: true,
		},

		{
			name: "requiredness differs for outputs",
			args: args{
				t1: requiredObjectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
				isOutput: true,
			},
			want: requiredObjectComplexType(map[string]schema.PropertySpec{
				"foo": stringPropertySpec(),
			}),
		},

		{
			name: "matching refs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": refPropertySpec("#/definitions/ResourceId"),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": refPropertySpec("#/definitions/ResourceId"),
				}),
				isOutput: true,
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": refPropertySpec("#/definitions/ResourceId"),
			}),
		},

		{
			name: "non-matching refs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": refPropertySpec("#/definitions/A"),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": refPropertySpec("#/definitions/B"),
				}),
				isOutput: true,
			},
			wantErr: true,
		},

		{
			name: "output oneOf merged with ref",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": refPropertySpec("#/definitions/ResourceId"),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/ResourceId"), stringType()),
					},
				}),
				isOutput: true,
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: oneOfType(refType("#/definitions/ResourceId"), stringType()),
				},
			}),
		},

		{
			name: "union two oneOf outputs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/FooA"), stringType()),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/FooB"), stringType()),
					},
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: oneOfType(refType("#/definitions/FooA"), stringType(), refType("#/definitions/FooB")),
				},
			}),
		},

		{
			name: "union two identical oneOf inputs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: oneOfType(refType("#/definitions/FooA"), refType("#/definitions/FooB")),
				},
			}),
		},

		{
			name: "union two identical discriminated types",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB"}, refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB"}, refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB"}, refType("#/definitions/FooA"), refType("#/definitions/FooB")),
				},
			}),
		},

		{
			name: "union two different discriminated types",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB"}, refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "c": "#/definitions/FooC"}, refType("#/definitions/FooA"), refType("#/definitions/FooC")),
					},
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB", "c": "#/definitions/FooC"}, refType("#/definitions/FooA"), refType("#/definitions/FooB"), refType("#/definitions/FooC")),
				},
			}),
		},

		{
			name: "union discriminated and non-discriminated types fails",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("prop", map[string]string{"a": "#/definitions/FooA", "b": "#/definitions/FooB"}, refType("#/definitions/FooA"), refType("#/definitions/FooB")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/FooA"), refType("#/definitions/FooC")),
					},
				}),
			},
			wantErr: true,
		},

		{
			name: "array of refs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": arrayPropertySpec(refType("#/definitions/ResourceId")),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": arrayPropertySpec(refType("#/definitions/ResourceId")),
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": arrayPropertySpec(refType("#/definitions/ResourceId")),
			}),
		},

		{
			name: "array of incompatible refs",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": arrayPropertySpec(refType("#/definitions/A")),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": arrayPropertySpec(refType("#/definitions/B")),
				}),
			},
			wantErr: true,
		},

		{
			name: "map of strings",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": objectPropertySpec(stringType()),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": objectPropertySpec(stringType()),
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": objectPropertySpec(stringType()),
			}),
		},

		{
			name: "map of incompatible types",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": objectPropertySpec(refType("#/definitions/A")),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": objectPropertySpec(refType("#/definitions/B")),
				}),
			},
			wantErr: true,
		},

		{
			name: "non-mergable types",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": objectPropertySpec(refType("#/definitions/A")),
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": arrayPropertySpec(refType("#/definitions/A")),
				}),
			},
			wantErr: true,
		},

		{
			name: "merge oneOf with primative",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/A"), refType("#/definitions/B")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": stringPropertySpec(),
				}),
			},
			want: objectComplexType(map[string]schema.PropertySpec{
				"foo": {
					TypeSpec: oneOfType(refType("#/definitions/A"), refType("#/definitions/B"), stringType()),
				},
			}),
		},

		{
			name: "merge oneOfs with different discriminators",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("a", map[string]string{"a": "#/definitions/A", "b": "#/definitions/B"}, refType("#/definitions/A"), refType("#/definitions/B")),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: discriminatedType("b", map[string]string{"a": "#/definitions/A", "c": "#/definitions/C"}, refType("#/definitions/A"), refType("#/definitions/C")),
					},
				}),
			},
			wantErr: true,
		},

		{
			name: "merge oneOfs with array",
			args: args{
				t1: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: oneOfType(refType("#/definitions/A"), refType("#/definitions/B"), stringType()),
					},
				}),
				t2: objectComplexType(map[string]schema.PropertySpec{
					"foo": {
						TypeSpec: arrayType(refType("#/definitions/A")),
					},
				}),
			},
			// cannot union oneOf with specified type
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeTypes(tt.args.t1, tt.args.t2, tt.args.isOutput)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if !assert.NotNil(t, got) {
				return
			}
			assert.Equal(t, tt.want, *got)
		})
	}
}

func objectComplexType(properties map[string]schema.PropertySpec) schema.ComplexTypeSpec {
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Type:       "object",
			Properties: properties,
		},
	}
}

func requiredObjectComplexType(properties map[string]schema.PropertySpec) schema.ComplexTypeSpec {
	required := make([]string, 0, len(properties))
	for k := range properties {
		required = append(required, k)
	}
	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Type:       "object",
			Properties: properties,
			Required:   required,
		},
	}
}

func stringPropertySpec() schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: stringType(),
	}
}

func refPropertySpec(ref string) schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: refType(ref),
	}
}

func arrayPropertySpec(items schema.TypeSpec) schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: arrayType(items),
	}
}

func objectPropertySpec(additionalProperties schema.TypeSpec) schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: objectType(additionalProperties),
	}
}

func oneOfType(types ...schema.TypeSpec) schema.TypeSpec {
	return schema.TypeSpec{
		OneOf: types,
	}
}

func discriminatedType(propertyName string, mapping map[string]string, types ...schema.TypeSpec) schema.TypeSpec {
	return schema.TypeSpec{
		OneOf: types,
		Discriminator: &schema.DiscriminatorSpec{
			PropertyName: propertyName,
			Mapping:      mapping,
		},
	}
}

func refType(ref string) schema.TypeSpec {
	return schema.TypeSpec{
		Ref: ref,
	}
}

func stringType() schema.TypeSpec {
	return schema.TypeSpec{
		Type: "string",
	}
}

func arrayType(items schema.TypeSpec) schema.TypeSpec {
	return schema.TypeSpec{
		Type:  "array",
		Items: &items,
	}
}

func objectType(additionalProperties schema.TypeSpec) schema.TypeSpec {
	return schema.TypeSpec{
		Type:                 "object",
		AdditionalProperties: &additionalProperties,
	}
}
