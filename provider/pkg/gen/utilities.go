// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"strings"
	"unicode"
)

// makeLegalIdentifier removes characters that are not allowed in identifiers.
func makeLegalIdentifier(s string) string {
	replacer := strings.NewReplacer("-", "", "[", "", "]", "")
	return replacer.Replace(s)
}

// firstToLower returns a string with the first character lowercased (`HelloWorld` => `helloWorld`).
func firstToLower(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	return string(append([]rune{unicode.ToLower(runes[0])}, runes[1:]...))
}

// toLowerCamel converts a string to lowerCamelCase.
// The code is adopted from https://github.com/iancoleman/strcase but changed in several ways to handle
// all the cases that are found in Azure in a most user-friendly way.
func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	if r := rune(s[0]); r == '_' {
		s = s[1:]
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

// toUpperCamel converts a string to UpperCamelCase.
func toUpperCamel(s string) string {
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
