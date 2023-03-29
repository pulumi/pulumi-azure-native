// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"strings"
	"unicode"
)

var legalIdentifierReplacer *strings.Replacer

// MakeLegalIdentifier removes characters that are not allowed in identifiers.
func MakeLegalIdentifier(s string) string {
	if legalIdentifierReplacer == nil {
		legalIdentifierReplacer = strings.NewReplacer("-", "", "[", "", "]", "")
	}
	return legalIdentifierReplacer.Replace(s)
}

// firstToLower returns a string with the first character lowercased (`HelloWorld` => `helloWorld`).
func firstToLower(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	return string(append([]rune{unicode.ToLower(runes[0])}, runes[1:]...))
}

var commonPrefixReplacements = map[rune]string{
	'*': "asterisk",
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
	'_': "",
}

// ToLowerCamel converts a string to lowerCamelCase.
// The code is adopted from https://github.com/iancoleman/strcase but changed in several ways to handle
// all the cases that are found in Azure in a most user-friendly way.
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	r := []rune(s)
	if prefix, found := commonPrefixReplacements[r[0]]; found {
		if len(r) > 1 {
			s = prefix + strings.ToUpper(string(r[1])) + string(r[2:])
		} else {
			s = prefix
		}
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

// ToUpperCamel converts a string to UpperCamelCase.
func ToUpperCamel(s string) string {
	return toCamelInitCase(s, true)
}

func toCamelInitCase(s string, initCase bool) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

var uppercaseAcronym = map[string]bool{
	"ID":  true,
	"TTL": true,
}
