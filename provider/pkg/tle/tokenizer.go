package tle

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

func NewTokenizer(basicTokenizer *basicTokenizer, text string) *Tokenizer {
	return &Tokenizer{
		basicTokenizer: basicTokenizer,
		text:           text,
	}
}

type Tokenizer struct {
	current                *Token
	currentTokenStartIndex int
	text                   string
	basicTokenizer         *basicTokenizer
}

func fromString(stringValue string) *Tokenizer {
	contract.Assert(stringValue != "")
	contract.Assert(1 <= len(stringValue))
	contract.Assert(isQuoteCharacter(stringValue[0]))

	initialQuoteCharacter := stringValue[0]
	trimmedLength := len(stringValue)
	if strings.HasSuffix(stringValue, string(initialQuoteCharacter)) {
		trimmedLength = len(stringValue) - 1
	}
	trimmedString := stringValue[1:trimmedLength]

	return NewTokenizer(NewBasicTokenizer(trimmedString), stringValue)
}

func (t *Tokenizer) HasStarted() bool {
	return t.basicTokenizer.HasStarted()
}

func (t *Tokenizer) HasCurrent() bool {
	return t.current != nil
}

func (t *Tokenizer) Current() *Token {
	return t.current
}

func (t *Tokenizer) nextBasicToken() {
	t.basicTokenizer.MoveNext()
}

func (t *Tokenizer) currentBasicToken() *basicToken {
	return t.basicTokenizer.Current()
}

func (t *Tokenizer) ReadToken() *Token {
	if !t.HasStarted() {
		t.nextBasicToken()
	} else if t.current != nil {
		t.currentTokenStartIndex += t.current.Length()
	}

	t.current = nil
	currentBasicToken := t.currentBasicToken()
	if currentBasicToken != nil {
		switch currentBasicToken.getType() {
		case leftParenthesis:
			curr := CreateLeftParenthesis(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case rightParenthesis:
			curr := CreateRightParenthesis(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case leftSquareBracket:
			curr := CreateLeftSquareBracket(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case rightSquareBracket:
			curr := CreateRightSquareBracket(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case comma:
			curr := CreateComma(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case period:
			curr := CreatePeriod(t.currentTokenStartIndex)
			t.current = &curr
			t.nextBasicToken()
			break

		case space, tab, carriageReturn, newLine, carriageReturnNewLine:
			ws := readWhitespace(t.basicTokenizer)
			var strs []stringable
			for _, w := range ws {
				strs = append(strs, w)
			}
			curr := CreateWhitespace(
				t.currentTokenStartIndex, getCombinedText(strs...))
			t.current = &curr
			break

		case doubleQuote:
			qs := readQuotedString(t.basicTokenizer)
			var strs []stringable
			for _, q := range qs {
				strs = append(strs, q)
			}
			curr := CreateQuotedString(
				t.currentTokenStartIndex,
				getCombinedText(strs...))
			t.current = &curr
			break

		case singleQuote:
			qts := readQuotedTLEString(t.basicTokenizer)
			var strs []stringable
			for _, q := range qts {
				strs = append(strs, q)
			}
			curr := CreateQuotedString(
				t.currentTokenStartIndex,
				getCombinedText(strs...))
			t.current = &curr
			break

		case dash, digits:
			ns := readNumber(t.basicTokenizer)
			var strs []stringable
			for _, n := range ns {
				strs = append(strs, n)
			}
			curr := CreateNumber(
				t.currentTokenStartIndex,
				getCombinedText(strs...))
			t.current = &curr
			break

		default:
			literalTokens := []*basicToken{currentBasicToken}
			t.nextBasicToken()

			for t.currentBasicToken() != nil &&
				(t.currentBasicToken().getType() == letters ||
					t.currentBasicToken().getType() == digits ||
					t.currentBasicToken().getType() == underscore) {
				literalTokens = append(literalTokens, t.currentBasicToken())
				t.nextBasicToken()
			}

			var ls []stringable
			for _, l := range literalTokens {
				ls = append(ls, l)
			}
			curr := CreateLiteral(t.currentTokenStartIndex, getCombinedText(ls...))
			t.current = &curr
			break
		}
	}

	return t.current
}

func (t *Tokenizer) Next() bool {
	result := t.ReadToken() != nil
	t.skipWhitespace()
	return result
}

func (t *Tokenizer) skipWhitespace() {
	for t.current != nil && t.current.GetType() == Whitespace {
		t.Next()
	}
}

type TokenType string

const (
	LeftSquareBracket  TokenType = "LeftSquareBracket"
	RightSquareBracket TokenType = "RightSquareBracket"
	LeftParenthesis    TokenType = "LeftParenthesis"
	RightParenthesis   TokenType = "RightParenthesis"
	QuotedString       TokenType = "QuotedString"
	Comma              TokenType = "Comma"
	Whitespace         TokenType = "Whitespave"
	Literal            TokenType = "Literal"
	Period             TokenType = "Period"
	Number             TokenType = "Number"
)

func NewToken(tokenType TokenType, startIndex int, stringValue string) Token {
	return Token{
		typ:         tokenType,
		span:        Span{startIndex: startIndex, length: len(stringValue)},
		stringValue: stringValue,
	}
}

type Token struct {
	typ         TokenType
	span        Span
	stringValue string
}

func (t Token) GetType() TokenType {
	return t.typ
}

func (t Token) Span() Span {
	return t.span
}

func (t Token) Length() int {
	return t.span.getLength()
}

func (t Token) StringValue() string {
	return t.stringValue
}

func CreateLeftParenthesis(startIndex int) Token {
	return NewToken(LeftParenthesis, startIndex, "(")
}

func CreateRightParenthesis(startIndex int) Token {
	return NewToken(RightParenthesis, startIndex, ")")
}

func CreateLeftSquareBracket(startIndex int) Token {
	return NewToken(LeftSquareBracket, startIndex, "[")
}

func CreateRightSquareBracket(startIndex int) Token {
	return NewToken(RightSquareBracket, startIndex, "]")
}

func CreateComma(startIndex int) Token {
	return NewToken(Comma, startIndex, ",")
}

func CreatePeriod(startIndex int) Token {
	return NewToken(Period, startIndex, ".")
}

func CreateWhitespace(startIndex int, stringValue string) Token {
	return NewToken(Whitespace, startIndex, stringValue)
}

func CreateQuotedString(startIndex int, stringValue string) Token {
	return NewToken(QuotedString, startIndex, stringValue)
}

func CreateNumber(startIndex int, stringValue string) Token {
	return NewToken(Number, startIndex, stringValue)
}

func CreateLiteral(startIndex int, stringValue string) Token {
	return NewToken(Literal, startIndex, stringValue)
}

func CreateEmptyLiteral(startIndex int) Token {
	return NewToken(Literal, startIndex, "")
}
