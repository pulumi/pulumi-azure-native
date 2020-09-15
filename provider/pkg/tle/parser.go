package tle

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

type Visitor interface {
	VisitArrayAccessValue(*ArrayAccessValue) error
	VisitFunctionCall(*FunctionCallValue) error
	VisitNumber(*NumberValue) error
	VisitPropertyAccess(*PropertyAccessValue) error
	VisitStringValue(*StringValue) error
}

type Value interface {
	Span() *Span
	Accept(visitor Visitor) error
}

func NewStringValue(token Token) *StringValue {
	contract.Assert(token.GetType() == QuotedString)
	return &StringValue{
		token: token,
	}
}

type StringValue struct {
	token Token
}

func (s *StringValue) Accept(visitor Visitor) error {
	return visitor.VisitStringValue(s)
}

func (s *StringValue) Span() *Span {
	return &s.token.span
}

// NewNumberValue creates a new NumberValue
func NewNumberValue(token Token) *NumberValue {
	contract.Assert(token.GetType() == Number)
	return &NumberValue{token}
}

type NumberValue struct {
	token Token
}

func (n *NumberValue) Span() *Span {
	return &n.token.span
}

func (n *NumberValue) Accept(v Visitor) error {
	return v.VisitNumber(n)
}

func NewPropertyAccessValue(source Value, periodToken Token, nameToken Token) *PropertyAccessValue {
	return &PropertyAccessValue{
		source:      source,
		periodToken: periodToken,
		nameToken:   nameToken,
	}
}

type PropertyAccessValue struct {
	source      Value
	periodToken Token
	nameToken   Token
}

func (p *PropertyAccessValue) Span() *Span {
	result := p.source.Span()
	return result.union(&p.nameToken.span)
}

func (p *PropertyAccessValue) Accept(v Visitor) error {
	return v.VisitPropertyAccess(p)
}

func NewArrayAccessValue(source Value, leftSquareBracketToken Token, indexValue Value, rightSquareBracketToken Token) *ArrayAccessValue {
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
	rightSquareBracketToken Token
}

func (a *ArrayAccessValue) Span() *Span {
	result := a.source.Span()
	return union(result, &a.rightSquareBracketToken.span)
}

func (a *ArrayAccessValue) Accept(v Visitor) error {
	return v.VisitArrayAccessValue(a)
}

type FunctionCallValue struct {
	NamespaceToken        *Token
	PeriodToken           *Token
	NameToken             *Token
	LeftParenthesisToken  *Token
	CommaTokens           []*Token
	ArgumentExpressions   []Value
	RightParenthesisToken *Token
}

func (f *FunctionCallValue) Span() *Span {
	return union(f.argumentListSpan(), f.fullNameSpan())
}

func (f *FunctionCallValue) Accept(visitor Visitor) error {
	return visitor.VisitFunctionCall(f)
}

func (f *FunctionCallValue) fullNameSpan() *Span {
	u := union(&f.NamespaceToken.span, &f.PeriodToken.span)

	result :=
		union(u, &f.NameToken.span)

		//Should have had at least one of a namespace or a name, therefore span should be non-empty
	contract.Assert(result != nil)
	// tslint:disable-next-line: no-non-null-assertion // Asserted
	return result
}

func (f *FunctionCallValue) argumentListSpan() *Span {
	var result *Span

	if f.LeftParenthesisToken != nil {
		result = &f.LeftParenthesisToken.span

		if f.RightParenthesisToken != nil {
			result = result.union(&f.RightParenthesisToken.span)
		} else if len(f.ArgumentExpressions) > 0 || len(f.CommaTokens) > 0 {
			for i := len(f.ArgumentExpressions) - 1; 0 <= i; i-- {
				arg := f.ArgumentExpressions[i]
				if arg != nil {
					span := arg.Span()
					result = result.union(span)
					break
				}
			}

			if 0 < len(f.CommaTokens) {
				result = result.union(&f.CommaTokens[len(f.CommaTokens)-1].span)
			}
		}
	}

	return result
}

type TLEParseResult struct {
	LeftSquareBracketToken  *Token
	Expression              Value
	RightSquareBracketToken *Token
}

func Parse(quotedStringValue string) (*TLEParseResult, error) {
	contract.Assert(1 <= len(quotedStringValue)) // TLE strings must be at least 1 character
	// The first character in the TLE string to parse must be a quote character
	contract.Assert(isQuoteCharacter(quotedStringValue[0]))

	var leftSquareBracketToken *Token
	var expression Value
	var rightSquareBracketToken *Token
	if 3 <= len(quotedStringValue) && quotedStringValue[1:2] == "[[" {
		expression = NewStringValue(CreateQuotedString(0, quotedStringValue))
	} else {
		tokenizer := fromString(quotedStringValue)
		tokenizer.Next()

		if tokenizer.Current() == nil || tokenizer.Current().GetType() != LeftSquareBracket {
			// This is just a plain old string (no brackets). Mark its expression as being
			// the string value.
			expression = NewStringValue(CreateQuotedString(0, quotedStringValue))
		} else {
			leftSquareBracketToken = tokenizer.Current()
			tokenizer.Next()

			if tokenizer.Current() != nil &&
				tokenizer.Current().GetType() != Literal &&
				tokenizer.Current().GetType() != RightSquareBracket {
				return nil, fmt.Errorf("Expected a literal value: %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
			}

			var err error
			expression, err = parseExpression(tokenizer)
			if err != nil {
				return nil, fmt.Errorf("failed to parse expression '%s': %w", quotedStringValue, err)
			}

			for tokenizer.Current() != nil {
				if tokenizer.Current().GetType() == RightSquareBracket {
					rightSquareBracketToken = tokenizer.Current()
					tokenizer.Next()
					break
				} else {
					return nil, fmt.Errorf("nxpected end of string: %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
				}
			}

			if rightSquareBracketToken != nil {
				if tokenizer.Current() != nil {
					return nil, fmt.Errorf("nothing should exist after the closing ']' except for whitespace, %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
				}
			} else {
				return nil, fmt.Errorf("expected a right square bracket (']'), %s:%v", quotedStringValue, len(quotedStringValue)-1)
			}

			if expression == nil {
				errorSpan := &leftSquareBracketToken.span
				if rightSquareBracketToken != nil {
					errorSpan = union(errorSpan, &rightSquareBracketToken.span)
				}
				return nil, fmt.Errorf("expected a function of property expression, %s:%v", quotedStringValue, errorSpan.getStartIndex())
			}
		}
	}

	return &TLEParseResult{
		LeftSquareBracketToken:  leftSquareBracketToken,
		Expression:              expression,
		RightSquareBracketToken: rightSquareBracketToken,
	}, nil
}

func parseExpression(tokenizer *Tokenizer) (Value, error) {
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
				return nil, fmt.Errorf("expected a literal value in span at %d", errorSpan.getStartIndex())
			}

			expression = NewPropertyAccessValue(expression, *periodToken, *propertyNameToken)
		} else if tokenizer.Current().GetType() == LeftSquareBracket {
			leftSquareBracketToken := tokenizer.Current()
			tokenizer.Next()

			indexValue, err := parseExpression(tokenizer)
			if err != nil {
				return nil, err
			}

			var rightSquareBracketToken *Token
			if tokenizer.Current() != nil && tokenizer.Current().GetType() == RightSquareBracket {
				rightSquareBracketToken = tokenizer.Current()
				tokenizer.Next()
			}

			if leftSquareBracketToken == nil || rightSquareBracketToken == nil {
				return nil, fmt.Errorf("malformed array offset specification near: %d", indexValue.Span().getStartIndex())
			}
			expression = NewArrayAccessValue(expression, *leftSquareBracketToken, indexValue, *rightSquareBracketToken)
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
			return nil, fmt.Errorf("Expected user-defined function name: %v", periodToken.Span().getStartIndex())
		}
	} else {
		nameToken = firstToken
	}

	var leftParenthesisToken, rightParanthesisToken *Token
	var commaTokens []*Token
	var argumentExpressions []Value

	if tokenizer.Current() != nil {
		for tokenizer.Current() != nil {
			if tokenizer.Current().GetType() == LeftParenthesis {
				leftParenthesisToken = tokenizer.Current()
				tokenizer.Next()
				break
			} else if tokenizer.Current().GetType() == RightSquareBracket {
				return nil, fmt.Errorf("Missing function argument list: %d", nameToken.Span().getStartIndex())
			} else {
				return nil, fmt.Errorf("Expected the end of the string: %d", nameToken.Span().getStartIndex())
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
				expression, err := parseExpression(tokenizer)
				if err != nil {
					return nil, err
				}
				if expression == nil && tokenizer.HasCurrent() && tokenizer.Current().GetType() == Comma {
					return nil, errors.New("expected a constant string, function, or property expression")
				}
				argumentExpressions = append(argumentExpressions, expression)
				expectingArgument = false
			} else if tokenizer.Current().GetType() == Comma {
				expectingArgument = true
				commaTokens = append(commaTokens, tokenizer.Current())
				tokenizer.Next()
			} else {
				return nil, errors.New("expected a comma (',')")
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

	return &FunctionCallValue{
		NamespaceToken:        namespaceToken,
		PeriodToken:           periodToken,
		NameToken:             nameToken,
		LeftParenthesisToken:  leftParenthesisToken,
		CommaTokens:           commaTokens,
		ArgumentExpressions:   argumentExpressions,
		RightParenthesisToken: rightParanthesisToken,
	}, nil
}
