// Copyright 2021, Pulumi Corporation.  All rights reserved.

package tle

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type BasicTokenType int

const (
	leftCurlyBracket BasicTokenType = iota
	rightCurlyBracket
	leftSquareBracket
	rightSquareBracket
	leftParenthesis
	rightParenthesis
	underscore
	period
	dash
	plus
	comma
	colon
	singleQuote
	doubleQuote
	backslash
	forwardSlash
	asterisk
	space
	tab
	newLine
	carriageReturn
	carriageReturnNewLine
	letters
	digits
	unrecognized
)

var (
	tokenLeftCurlyBracket      = basicToken{"{", leftCurlyBracket}
	tokenRightCurlyBracket     = basicToken{"}", rightCurlyBracket}
	tokenLeftSquareBracket     = basicToken{"[", leftSquareBracket}
	tokenRightSquareBracket    = basicToken{"]", rightSquareBracket}
	tokenLeftParenthesis       = basicToken{"(", leftParenthesis}
	tokenRightParenthesis      = basicToken{")", rightParenthesis}
	tokenUnderscore            = basicToken{"_", underscore}
	tokenPeriod                = basicToken{".", period}
	tokenDash                  = basicToken{"-", dash}
	tokenPlus                  = basicToken{"+", plus}
	tokenComma                 = basicToken{",", comma}
	tokenColon                 = basicToken{":", colon}
	tokenSingleQuote           = basicToken{`'`, singleQuote}
	tokenDoubleQuote           = basicToken{`"`, doubleQuote}
	tokenBackslash             = basicToken{"\\", backslash}
	tokenForwardSlash          = basicToken{"/", forwardSlash}
	tokenAsterisk              = basicToken{"*", asterisk}
	tokenSpace                 = basicToken{" ", space}
	tokenTab                   = basicToken{"\t", tab}
	tokenNewLine               = basicToken{"\n", newLine}
	tokenCarriageReturn        = basicToken{"\r", carriageReturn}
	tokenCarriageReturnNewLine = basicToken{"\r\n", carriageReturnNewLine}
)

// lettersFromText creates a Token from the provided text.
func lettersFromText(text string) basicToken {
	return basicToken{text: text, typ: letters}
}

/**
 * Create a digitsFromText Token from the provided text.
 */
func digitsFromText(text string) basicToken {
	return basicToken{text: text, typ: digits}
}

/**
 * Create an unrecognizedFromText Token from the provided text.
 */
func unrecognizedFromText(text string) basicToken {
	return basicToken{text: text, typ: unrecognized}
}

/**
 * Get whether the provided character is a letter or not.
 */
func isLetter(character byte) bool {
	return ('a' <= character && character <= 'z') || ('A' <= character && character <= 'Z')
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

type basicToken struct {
	text string
	typ  BasicTokenType
}

func (t *basicToken) toString() string {
	return t.text
}

func (t *basicToken) getType() BasicTokenType {
	return t.typ
}

func newBasicTokenizer(text string) *basicTokenizer {
	return &basicTokenizer{
		text:       text,
		textIndex:  -1,
		textLength: len(text),
	}
}

type basicTokenizer struct {
	textLength   int
	textIndex    int
	currentToken *basicToken
	text         string
}

type TokenIterator interface {
	HasStarted() bool
	Current() *basicToken
	MoveNext() bool
}

func (t *basicTokenizer) HasStarted() bool {
	return t.textIndex >= 0
}

func (t *basicTokenizer) Current() *basicToken {
	return t.currentToken
}

func (t *basicTokenizer) currentCharacter() (char byte, eof bool) {
	if 0 <= t.textIndex && t.textIndex < t.textLength {
		char = t.text[t.textIndex]
		return
	}
	eof = true
	return
}

func (t *basicTokenizer) nextCharacter() {
	t.textIndex++
}

func (t *basicTokenizer) MoveNext() bool {
	if !t.HasStarted() {
		t.nextCharacter()
	}

	c, eof := t.currentCharacter()
	if eof {
		t.currentToken = nil
	} else {
		switch c {
		case '{':
			t.currentToken = &tokenLeftCurlyBracket
			t.nextCharacter()
		case '}':
			t.currentToken = &tokenRightCurlyBracket
			t.nextCharacter()
		case '[':
			t.currentToken = &tokenLeftSquareBracket
			t.nextCharacter()
		case ']':
			t.currentToken = &tokenRightSquareBracket
			t.nextCharacter()
		case '(':
			t.currentToken = &tokenLeftParenthesis
			t.nextCharacter()
		case ')':
			t.currentToken = &tokenRightParenthesis
			t.nextCharacter()
		case '_':
			t.currentToken = &tokenUnderscore
			t.nextCharacter()
		case '.':
			t.currentToken = &tokenPeriod
			t.nextCharacter()
		case '-':
			t.currentToken = &tokenDash
			t.nextCharacter()
		case '+':
			t.currentToken = &tokenPlus
			t.nextCharacter()
		case ',':
			t.currentToken = &tokenComma
			t.nextCharacter()
		case ':':
			t.currentToken = &tokenColon
			t.nextCharacter()
		case '\'':
			t.currentToken = &tokenSingleQuote
			t.nextCharacter()
		case '"':
			t.currentToken = &tokenDoubleQuote
			t.nextCharacter()
		case '\\':
			t.currentToken = &tokenBackslash
			t.nextCharacter()
		case '/':
			t.currentToken = &tokenForwardSlash
			t.nextCharacter()
		case '*':
			t.currentToken = &tokenAsterisk
			t.nextCharacter()
		case '\n':
			t.currentToken = &tokenNewLine
			t.nextCharacter()
		case '\r':
			t.nextCharacter()
			var newC byte
			newC, eof = t.currentCharacter()
			if !eof && string(newC) == "\n" {
				t.currentToken = &tokenCarriageReturnNewLine
				t.nextCharacter()
			} else {
				t.currentToken = &tokenCarriageReturn
			}
		case ' ':
			t.currentToken = &tokenSpace
			t.nextCharacter()
		case '\t':
			t.currentToken = &tokenTab
			t.nextCharacter()
		default:
			if isLetter(c) {
				letters := lettersFromText(t.readWhile(isLetter))
				t.currentToken = &letters
				break
			}
			var cNew byte
			cNew, _ = t.currentCharacter()
			if isDigit(cNew) {
				digits := digitsFromText(t.readWhile(isDigit))
				t.currentToken = &digits
			} else {
				unrecognized := unrecognizedFromText(string(cNew))
				t.currentToken = &unrecognized
				t.nextCharacter()
			}
		}
	}

	return t.currentToken != nil
}

func (t *basicTokenizer) readWhile(condition func(character byte) bool) string {
	var result string
	res, _ := t.currentCharacter()
	result = string(res)
	t.nextCharacter()
	_, eof := t.currentCharacter()

	for !eof {
		res, eof = t.currentCharacter()
		if !condition(res) {
			break
		}
		result += string(res)
		t.nextCharacter()
	}

	return result
}

/**
 * Read a JSON whitespace string from the provided tokenizer.
 */
func readWhitespace(iterator TokenIterator) []*basicToken {
	var whitespaceTokens []*basicToken

	for {
		current := iterator.Current()
		if current == nil {
			return whitespaceTokens
		}

		switch current.getType() {
		case space, tab, carriageReturn, newLine, carriageReturnNewLine:
			whitespaceTokens = append(whitespaceTokens, current)
			iterator.MoveNext()
		default:
			return whitespaceTokens
		}
	}
}

/**
 * Read a JSON quoted string from the provided tokenizer. The tokenizer must be pointing at
 * either a SingleQuote or DoubleQuote Token.
 *
 * Note that the returned quoted string may not end with a quote (if EOD is reached)
 */
func readQuotedString(iterator TokenIterator) []*basicToken {
	startingToken := iterator.Current()
	contract.Assert(startingToken != nil &&
		(startingToken.getType() == singleQuote || startingToken.getType() == doubleQuote))

	startQuote := startingToken
	quotedStringTokens := []*basicToken{startQuote}
	iterator.MoveNext()

	escaped := false
	var endQuote *basicToken
	current := iterator.Current()
	for endQuote == nil && current != nil {
		quotedStringTokens = append(quotedStringTokens, current)

		if escaped {
			escaped = false
		} else {
			if current.getType() == backslash {
				escaped = true
			} else if current.getType() == startQuote.getType() {
				endQuote = iterator.Current()
			}
		}

		iterator.MoveNext()
		current = iterator.Current()
	}

	return quotedStringTokens
}

/**
 * Read a JSON number from the provided iterator. The iterator must be pointing at either a
 * Dash or digitsFromText Token when this function is called.
 */
func readNumber(iterator TokenIterator) []*basicToken {
	var numberBasicTokens []*basicToken

	dashOrDigitsToken := iterator.Current()
	if dashOrDigitsToken.getType() == dash {
		// Negative sign
		numberBasicTokens = append(numberBasicTokens, dashOrDigitsToken)
		iterator.MoveNext()
	}

	digitsI := iterator.Current()
	if digitsI != nil && digitsI.getType() == digits {
		// Whole number digits
		numberBasicTokens = append(numberBasicTokens, digitsI)
		iterator.MoveNext()
	}

	decimal := iterator.Current()
	if decimal != nil && decimal.getType() == period {
		// Decimal point
		numberBasicTokens = append(numberBasicTokens, decimal)
		iterator.MoveNext()

		fractionalDigits := iterator.Current()
		if fractionalDigits != nil && fractionalDigits.getType() == digits {
			// Fractional number digits
			numberBasicTokens = append(numberBasicTokens, fractionalDigits)
			iterator.MoveNext()
		}
	}

	exponentLetter := iterator.Current()
	if exponentLetter != nil {
		if exponentLetter.getType() == letters && strings.ToLower(exponentLetter.toString()) == "e" {
			// e
			numberBasicTokens = append(numberBasicTokens, exponentLetter)
			iterator.MoveNext()

			exponentSign := iterator.Current()
			if exponentSign != nil &&
				(exponentSign.getType() == dash || exponentSign.getType() == plus) {
				// Exponent number sign
				numberBasicTokens = append(numberBasicTokens, exponentSign)
				iterator.MoveNext()
			}

			exponentDigits := iterator.Current()
			if exponentDigits != nil && exponentDigits.getType() == digits {
				// Exponent number digits
				numberBasicTokens = append(numberBasicTokens, exponentDigits)
				iterator.MoveNext()
			}
		}
	}

	return numberBasicTokens
}
