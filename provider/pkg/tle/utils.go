package tle

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

func isQuoteCharacter(character byte) bool {
	return character == '\'' ||
		character == '"'
}

type stringable interface {
	toString() string
}

/**
 * Get the combined text of the provided values.
 */
func getCombinedText(values ...stringable) string {
	var result strings.Builder
	for _, v := range values {
		_, _ = result.WriteString(v.toString())
	}
	return result.String()
}

// TODO move this back to tokenizer.go
/**
 * Handles reading a single-quoted string inside a JSON-encoded TLE string. Handles both JSON string
 * escape characters (e.g. \n, \") and escaped single quotes in TLE style (two single quotes together,
 * e.g. 'That''s all, folks!')
 */
func readQuotedTLEString(iterator TokenIterator) []*basicToken {
	current := iterator.Current()
	contract.Assert(current != nil)
	contract.Assert(current.getType() == singleQuote)
	quotedStringTokens := []*basicToken{current}
	iterator.MoveNext()

	escaped := false
	for iterator.Current() != nil {
		current = iterator.Current()
		quotedStringTokens = append(quotedStringTokens, current)

		if escaped {
			escaped = false
		} else {
			if current.getType() == backslash {
				escaped = true
			} else if current.getType() == singleQuote {
				// If the next token is also a single quote, it's escaped, otherwise it's the
				// end of the string.
				iterator.MoveNext()
				afterCurrent := iterator.Current()
				if afterCurrent == nil {
					break
				}

				if afterCurrent.getType() == singleQuote {
					escaped = true
					continue
				} else {
					break
				}
			}
		}

		iterator.MoveNext()
	}

	return quotedStringTokens
}

func unquote(value string) string {
	if value == "" {
		return ""
	}

	if strings.HasPrefix(value, "\"") {
		if strings.HasSuffix(value, "\"") {
			return value[1 : len(value)-1]
		} else {
			return value[1:]
		}
	} else if strings.HasPrefix(value, "'") {
		if strings.HasSuffix(value, "'") {
			return value[1 : len(value)-1]
		}
		return value[1:]
	}

	return value
}
