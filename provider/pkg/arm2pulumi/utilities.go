package arm2pulumi

import "strings"

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

