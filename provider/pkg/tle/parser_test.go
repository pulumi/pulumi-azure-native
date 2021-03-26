// Copyright 2021, Pulumi Corporation.  All rights reserved.

package tle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseExpression(t *testing.T) {
	for _, test := range []struct {
		name       string
		expression string
		expected   Value
		err        error
	}{
		{name: "Literal unquoted",
			expression: "hello there",
			expected: &StringValue{
				token: Token{typ: "QuotedString",
					span:        Span{startIndex: 0, length: 13},
					stringValue: "\"hello there\""}},
		},
		{name: "Literal double quoted",
			expression: "\"hello there\"",
			expected: &StringValue{
				token: Token{typ: "QuotedString",
					span:        Span{startIndex: 0, length: 13},
					stringValue: "\"hello there\""}},
		},
		{name: "Literal single quoted",
			expression: "'hello there'",
			expected: &StringValue{
				token: Token{typ: "QuotedString",
					span:        Span{startIndex: 0, length: 13},
					stringValue: "'hello there'"}},
		},
		{name: "Literal unquoted with internal single quote",
			expression: "hello 'there'",
			expected: &StringValue{
				token: Token{typ: "QuotedString",
					span:        Span{startIndex: 0, length: 15},
					stringValue: "\"hello 'there'\""}},
		},
		{name: "Literal unquoted with internal double quote",
			expression: "hello \"there\"",
			expected: &StringValue{
				token: Token{typ: "QuotedString",
					span:        Span{startIndex: 0, length: 15},
					stringValue: "'hello \"there\"'"}},
		},
		{
			name:       "Simple Function Invocation",
			expression: "\"[parameters('dnsPrefix')]\"",
			expected: &FunctionCallValue{
				NameToken: Token{
					typ: "Literal",
					span: Span{
						startIndex: 1,
						length:     10,
					},
					stringValue: "parameters",
				},
				LeftParenthesisToken: &Token{
					typ: LeftParenthesis,
					span: Span{
						startIndex: 11,
						length:     1,
					},
					stringValue: "(",
				},
				ArgumentExpressions: []Value{
					&StringValue{
						token: Token{
							typ: QuotedString,
							span: Span{
								startIndex: 12,
								length:     11,
							},
							stringValue: "'dnsPrefix'",
						},
					},
				},
				RightParenthesisToken: &Token{
					typ: RightParenthesis,
					span: Span{
						startIndex: 23,
						length:     1,
					},
					stringValue: ")",
				},
			},
		},
		{
			name:       "Nested Functions",
			expression: "\"[reference(parameters('clusterName')).fqdn]\"",
			expected: NewPropertyAccessValue(&FunctionCallValue{
				NameToken: Token{
					typ:         Literal,
					span:        Span{startIndex: 1, length: 9},
					stringValue: "reference",
				},
				LeftParenthesisToken: &Token{
					typ: LeftParenthesis,
					span: Span{
						startIndex: 10,
						length:     1,
					},
					stringValue: "(",
				},
				ArgumentExpressions: []Value{
					&FunctionCallValue{
						NameToken: Token{
							typ:         Literal,
							span:        Span{startIndex: 11, length: 10},
							stringValue: "parameters",
						},
						LeftParenthesisToken: &Token{
							typ:         LeftParenthesis,
							span:        Span{startIndex: 21, length: 1},
							stringValue: "(",
						},
						ArgumentExpressions: []Value{
							&StringValue{
								token: Token{
									typ: QuotedString,
									span: Span{
										startIndex: 22,
										length:     13,
									},
									stringValue: "'clusterName'",
								},
							},
						},
						RightParenthesisToken: &Token{
							typ: RightParenthesis,
							span: Span{
								startIndex: 35,
								length:     1,
							},
							stringValue: ")",
						},
					},
				},
				RightParenthesisToken: &Token{
					typ: RightParenthesis,
					span: Span{
						startIndex: 36,
						length:     1,
					},
					stringValue: ")",
				},
			},
				Token{typ: Period,
					span:        Span{startIndex: 37, length: 1},
					stringValue: "."},
				Token{
					typ:         Literal,
					span:        Span{startIndex: 38, length: 4},
					stringValue: "fqdn",
				},
			),
		},
		{
			name:       "Multiple Arguments",
			expression: "\"[resourceId('Microsoft.Network/networkInterfaces','nic1')]\"",
			expected: &FunctionCallValue{
				NameToken: Token{
					typ: "Literal",
					span: Span{
						startIndex: 1,
						length:     10,
					},
					stringValue: "resourceId",
				},
				LeftParenthesisToken: &Token{
					typ: LeftParenthesis,
					span: Span{
						startIndex: 11,
						length:     1,
					},
					stringValue: "(",
				},
				ArgumentExpressions: []Value{
					&StringValue{
						token: Token{
							typ: QuotedString,
							span: Span{
								startIndex: 12,
								length:     37,
							},
							stringValue: "'Microsoft.Network/networkInterfaces'",
						},
					},
					&StringValue{
						token: Token{
							typ: QuotedString,
							span: Span{
								startIndex: 50,
								length:     6,
							},
							stringValue: "'nic1'",
						},
					},
				},
				CommaTokens: []*Token{
					{
						typ: Comma,
						span: Span{
							startIndex: 49,
							length:     1,
						},
						stringValue: ",",
					},
				},
				RightParenthesisToken: &Token{
					typ: RightParenthesis,
					span: Span{
						startIndex: 56,
						length:     1,
					},
					stringValue: ")",
				},
			},
		},
		{
			name:       "Function Invocation With Array Traversal",
			expression: "\"[parameters('dnsPrefix')[0]]\"",
			expected: NewArrayAccessValue(
				&FunctionCallValue{
					NameToken: Token{
						typ:         Literal,
						span:        Span{startIndex: 1, length: 10},
						stringValue: "parameters",
					},
					LeftParenthesisToken: &Token{
						typ:         LeftParenthesis,
						span:        Span{startIndex: 11, length: 1},
						stringValue: "(",
					},
					ArgumentExpressions: []Value{
						&StringValue{
							token: Token{
								typ: QuotedString,
								span: Span{
									startIndex: 12,
									length:     11,
								},
								stringValue: "'dnsPrefix'",
							},
						},
					},
					RightParenthesisToken: &Token{
						typ: RightParenthesis,
						span: Span{
							startIndex: 23,
							length:     1,
						},
						stringValue: ")",
					},
				},
				Token{
					typ:         LeftSquareBracket,
					span:        Span{startIndex: 24, length: 1},
					stringValue: "[",
				},
				NewNumberValue(Token{typ: Number, span: Span{startIndex: 25, length: 1}, stringValue: "0"}),
				Token{typ: RightSquareBracket, span: Span{startIndex: 26, length: 1}, stringValue: "]"},
			),
		},
		{
			name:       "Function Invocation With Array Traversal And Property Access",
			expression: "\"[parameters('dns')[0].ttl.value]\"",
			expected: NewPropertyAccessValue(
				NewPropertyAccessValue(
					NewArrayAccessValue(
						&FunctionCallValue{
							NameToken: Token{
								typ:         Literal,
								span:        Span{startIndex: 1, length: 10},
								stringValue: "parameters",
							},
							LeftParenthesisToken: &Token{
								typ:         LeftParenthesis,
								span:        Span{startIndex: 11, length: 1},
								stringValue: "(",
							},
							ArgumentExpressions: []Value{
								&StringValue{
									token: Token{
										typ: QuotedString,
										span: Span{
											startIndex: 12,
											length:     5,
										},
										stringValue: "'dns'",
									},
								},
							},
							RightParenthesisToken: &Token{
								typ: RightParenthesis,
								span: Span{
									startIndex: 17,
									length:     1,
								},
								stringValue: ")",
							},
						},
						Token{
							typ:         LeftSquareBracket,
							span:        Span{startIndex: 18, length: 1},
							stringValue: "[",
						},
						NewNumberValue(Token{typ: Number, span: Span{startIndex: 19, length: 1}, stringValue: "0"}),
						Token{typ: RightSquareBracket, span: Span{startIndex: 20, length: 1}, stringValue: "]"},
					),
					Token{typ: Period, stringValue: ".", span: Span{startIndex: 21, length: 1}},
					Token{typ: Literal, stringValue: "ttl", span: Span{startIndex: 22, length: 3}}),
				CreatePeriod(25),
				CreateLiteral(26, "value")),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			res, err := Parse(test.expression)
			if test.err != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.expected, res.Expression, "%#v", res.Expression)
		})

	}
}
