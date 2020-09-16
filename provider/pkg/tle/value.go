package tle

import "github.com/pulumi/pulumi/sdk/go/common/util/contract"

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

func (n *StringValue) Token() Token {
	return n.token
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

func (n *NumberValue) Token() Token {
	return n.token
}

func (n *NumberValue) Accept(v Visitor) error {
	return v.VisitNumberValue(n)
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

func (p *PropertyAccessValue) Source() Value {
	return p.source
}

func (p *PropertyAccessValue) PropertyName() string {
	return p.nameToken.StringValue()
}

func (p *PropertyAccessValue) Span() *Span {
	result := p.source.Span()
	return result.union(&p.nameToken.span)
}

func (p *PropertyAccessValue) Accept(v Visitor) error {
	return v.VisitPropertyAccessValue(p)
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

func (a *ArrayAccessValue) Source() Value {
	return a.source
}

func (a *ArrayAccessValue) Index() Value {
	return a.indexValue
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
	NameToken             Token
	LeftParenthesisToken  *Token
	CommaTokens           []*Token
	ArgumentExpressions   []Value
	RightParenthesisToken *Token
}

func (f *FunctionCallValue) Span() *Span {
	return union(f.argumentListSpan(), f.fullNameSpan())
}

func (f *FunctionCallValue) Accept(visitor Visitor) error {
	return visitor.VisitFunctionCallValue(f)
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
