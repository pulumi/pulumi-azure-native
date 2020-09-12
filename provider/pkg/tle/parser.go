package tle

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

type Value interface {
	Span() *Span
}

func NewStringValue(token Token) StringValue {
	contract.Assert(token.GetType() == QuotedString)
	return StringValue{
		token: token,
	}
}

type StringValue struct {
	token Token
}

func (s StringValue) Span() *Span {
	return &s.token.span
}

// NewNumberValue creates a new NumberValue
func NewNumberValue(token Token) NumberValue {
	contract.Assert(token.GetType() == Number)
	return NumberValue{token}
}

type NumberValue struct {
	token Token
}

func (n NumberValue) Span() *Span {
	return &n.token.span
}

func NewPropertyAccess(source Value, periodToken Token, nameToken *Token) *PropertyAccess {
	return &PropertyAccess{
		source:      source,
		periodToken: periodToken,
		nameToken:   nameToken,
	}
}

type PropertyAccess struct {
	source      Value
	periodToken Token
	nameToken   *Token
}

func (p *PropertyAccess) Span() *Span {
	result := p.source.Span()

	if p.nameToken != nil {
		result = result.union(&p.nameToken.span)
	} else {
		result = result.union(&p.periodToken.span)
	}

	return result
}

func NewArrayAccessValue(source Value, leftSquareBracketToken Token, indexValue Value, rightSquareBracketToken *Token) *ArrayAccessValue {
	return &ArrayAccessValue{
		source:                  source,
		leftSquareBracketToken:  leftSquareBracketToken,
		indexValue:              indexValue,
		rightSquareBracketToken: rightSquareBracketToken,
	}
}

type ArrayAccessValue struct {
	source                  Value
	leftSquareBracketToken  Token
	indexValue              Value
	rightSquareBracketToken *Token
}

func (a *ArrayAccessValue) Span() *Span {
	result := a.source.Span()

	if a.rightSquareBracketToken != nil {
		result = union(result, &a.rightSquareBracketToken.span)
	} else if a.indexValue != nil {
		result = union(result, a.indexValue.Span())
	} else {
		result = union(result, &a.leftSquareBracketToken.span)
	}

	return result
}

func NewFunctionCallValue(namespaceToken, periodToken, leftParenthesisToken *Token, nameToken *Token, commaTokens []*Token, argumentExpressions []Value, rightParanthesis *Token) *FunctionCallValue {
	return &FunctionCallValue{
		namespaceToken:        namespaceToken,
		periodToken:           periodToken,
		nameToken:             nameToken,
		leftParenthesisToken:  leftParenthesisToken,
		commaTokens:           commaTokens,
		argumentExpressions:   argumentExpressions,
		rightParenthesisToken: rightParanthesis,
	}
}

type FunctionCallValue struct {
	namespaceToken        *Token
	periodToken           *Token
	nameToken             *Token
	leftParenthesisToken  *Token
	commaTokens           []*Token
	argumentExpressions   []Value
	rightParenthesisToken *Token
}

func (f FunctionCallValue) Span() *Span {
	return union(f.argumentListSpan(), f.fullNameSpan())
}

func (f FunctionCallValue) fullNameSpan() *Span {
	u := union(&f.namespaceToken.span, &f.periodToken.span)

	result :=
		union(u, &f.nameToken.span)

		//Should have had at least one of a namespace or a name, therefore span should be non-empty
	contract.Assert(result != nil)
	// tslint:disable-next-line: no-non-null-assertion // Asserted
	return result
}

func (f *FunctionCallValue) argumentListSpan() *Span {
	var result *Span

	if f.leftParenthesisToken != nil {
		result = &f.leftParenthesisToken.span

		if f.rightParenthesisToken != nil {
			result = result.union(&f.rightParenthesisToken.span)
		} else if len(f.argumentExpressions) > 0 || len(f.commaTokens) > 0 {
			for i := len(f.argumentExpressions) - 1; 0 <= i; i-- {
				arg := f.argumentExpressions[i]
				if arg != nil {
					span := arg.Span()
					result = result.union(span)
					break
				}
			}

			if 0 < len(f.commaTokens) {
				result = result.union(&f.commaTokens[len(f.commaTokens)-1].span)
			}
		}
	}

	return result
}

func ParseExpression(tokenizer *Tokenizer) (Value, error) {
	var expression Value
	if tokenizer.Current() != nil {
		var rootExpression Value // Initial expression

		token := tokenizer.Current()
		tokenType := token.GetType()
		if tokenType == Literal {
			var err error
			rootExpression, err = parseFunctionCall(tokenizer)
			if err != nil {
				return nil, err
			}
		} else if tokenType == QuotedString {
			if !strings.HasSuffix(token.stringValue, string(token.stringValue[0])) {
				return nil, fmt.Errorf("a constant string is missing an end quote: %d", token.span.getStartIndex())
			}
			rootExpression = NewStringValue(*token)
			tokenizer.Next()
		} else if tokenType == Number {
			rootExpression = NewNumberValue(*token)
			tokenizer.Next()
		} else if tokenType != RightSquareBracket && tokenType != Comma {
			return nil, fmt.Errorf("template language expressions must start with a function: %d", token.span.getStartIndex())
			tokenizer.Next()
		}

		if rootExpression == nil {
			return nil, nil
		}

		expression = rootExpression
	} else {
		return nil, nil
	}

	// Check for property or array accesses off of the root expression
	for tokenizer.Current() != nil {
		if tokenizer.Current().GetType() == Period {
			periodToken := tokenizer.Current()
			tokenizer.Next()

			var propertyNameToken *Token
			var errorSpan Span

			if tokenizer.Current() != nil {
				if tokenizer.Current().GetType() == Literal {
					propertyNameToken = tokenizer.Current()
					tokenizer.Next()
				} else {
					errorSpan = tokenizer.Current().Span()

					tokenType := tokenizer.Current().GetType()
					if tokenType != RightParenthesis &&
						tokenType != RightSquareBracket &&
						tokenType != Comma {
						tokenizer.Next()
					}
				}
			} else {
				errorSpan = periodToken.Span()
			}

			if propertyNameToken == nil {
				// tslint:disable-next-line: no-non-null-assertion // Asserted
				return nil, fmt.Errorf("expected a literal value in span at %d", errorSpan.getStartIndex())
			}

			// We go ahead and create a property access expression whether the property name
			//   was correctly given or not, so we can have proper intellisense/etc.
			expression = NewPropertyAccess(expression, *periodToken, propertyNameToken)
		} else if tokenizer.Current().GetType() == LeftSquareBracket {
			leftSquareBracketToken := tokenizer.Current()
			tokenizer.Next()

			indexValue, err := ParseExpression(tokenizer)
			if err != nil {
				return nil, err
			}

			var rightSquareBracketToken *Token
			if tokenizer.Current() != nil && tokenizer.Current().GetType() == RightSquareBracket {
				rightSquareBracketToken = tokenizer.Current()
				tokenizer.Next()
			}

			expression = NewArrayAccessValue(expression, *leftSquareBracketToken, indexValue, rightSquareBracketToken)
		} else {
			break
		}
	}

	return expression, nil
}

func parseFunctionCall(tokenizer *Tokenizer) (*FunctionCallValue, error) {
	contract.Assert(tokenizer != nil)
	contract.Assert(tokenizer.Current() != nil)
	contract.Assert(Literal == tokenizer.Current().GetType())

	var namespaceToken, nameToken, periodToken *Token

	firstToken := tokenizer.Current()
	tokenizer.Next()

	// Check for <namespace>.<functionname>
	if tokenizer.Current() != nil && tokenizer.Current().GetType() == Period {
		// It's a user-defined function because it has a namespace before the function name
		periodToken = tokenizer.Current()
		namespaceToken = firstToken
		tokenizer.Next()

		// Get the function name following the period
		if tokenizer.HasCurrent() && tokenizer.Current().GetType() == Literal {
			nameToken = tokenizer.Current()
			tokenizer.Next()
		} else {
			return nil, fmt.Errorf("Expected user-defined function name: %s", periodToken.Span().getStartIndex())
		}
	} else {
		nameToken = firstToken
	}

	var leftParenthesisToken, rightParanthesisToken *Token
	var commaTokens []*Token
	var argumentExpressions []Value

	// tslint:disable-next-line: strict-boolean-expressions
	if tokenizer.Current() != nil {
		// tslint:disable-next-line: strict-boolean-expressions // False positive
		for tokenizer.Current() != nil {
			if tokenizer.Current().GetType() == LeftParenthesis {
				leftParenthesisToken = tokenizer.Current()
				tokenizer.Next()
				break
			} else if tokenizer.Current().GetType() == RightSquareBracket {
				return nil, fmt.Errorf("Missing function argument list: %d", nameToken.Span().getStartIndex())
			} else {
				return nil, fmt.Errorf("Expected the end of the string: %d", nameToken.Span().getStartIndex())
				tokenizer.Next()
			}
		}
	} else {
		return nil, fmt.Errorf("Missing function argument list: %d", nameToken.Span().getStartIndex())
	}

	if tokenizer.HasCurrent() {
		expectingArgument := true

		// tslint:disable-next-line: strict-boolean-expressions
		for tokenizer.Current() != nil {
			if tokenizer.Current().GetType() == RightParenthesis || tokenizer.Current().GetType() == RightSquareBracket {
				break
			} else if expectingArgument {
				expression, err := ParseExpression(tokenizer)
				if err != nil {
					return nil, err
				}
				if expression == nil && tokenizer.HasCurrent() && tokenizer.Current().GetType() == Comma {
					return nil, fmt.Errorf("Expected a constant string, function, or property expression.")
				}
				argumentExpressions = append(argumentExpressions, expression)
				expectingArgument = false
			} else if tokenizer.Current().GetType() == Comma {
				expectingArgument = true
				commaTokens = append(commaTokens, tokenizer.Current())
				tokenizer.Next()
			} else {
				return nil, errors.New("Expected a comma (',').")
				tokenizer.Next()
			}
		}
	} else if leftParenthesisToken != nil {
		return nil, fmt.Errorf("expected a right parenthesis (')')")
	}

	// tslint:disable-next-line: strict-boolean-expressions
	if tokenizer.Current() != nil {
		switch tokenizer.current.GetType() {
		case RightParenthesis:
			rightParanthesisToken = tokenizer.Current()
			tokenizer.Next()
			break

		case RightSquareBracket:
			if leftParenthesisToken != nil {
				return nil, fmt.Errorf("expected a right parenthesis (')')")
			}
		}
	}

	return NewFunctionCallValue(namespaceToken, periodToken, nameToken, leftParenthesisToken, commaTokens, argumentExpressions, rightParanthesisToken), nil
}
