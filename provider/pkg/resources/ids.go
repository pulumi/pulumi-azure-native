// Copyright 2024, Pulumi Corporation.

package resources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// ParseResourceID extracts templated values from the given resource ID based on the names of those templated
// values in an HTTP path. The structure of id and path must match: we validate it by building a regular
// expression based on the path parameters and matching the id.
func ParseResourceID(id, path string) (map[string]string, error) {
	pathParts := strings.Split(path, "/")
	regexParts := make([]string, len(pathParts))
	// Track the names of the regex matches so we can map them back to the original names.
	regexNames := make(map[string]string, len(pathParts))
	regexMatchGroupNameReplacer := regexp.MustCompile(`[^a-zA-Z0-9]`)
	for i, s := range pathParts {
		if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
			name := s[1 : len(s)-1]
			regexName := regexMatchGroupNameReplacer.ReplaceAllString(name, "")
			if clashingPathName, ok := regexNames[regexName]; ok {
				return nil, errors.Errorf("clashing path names: '%s' and '%s' while parsing resource ID for path '%s", clashingPathName, name, path)
			}
			regexNames[regexName] = name
			regexParts[i] = fmt.Sprintf("(?P<%s>.*?)", regexName)
		} else {
			regexParts[i] = pathParts[i]
		}
	}

	expr := fmt.Sprintf("(?i)^%s$", strings.Join(regexParts, "/"))
	pattern, err := regexp.Compile(expr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compile expression '%s' for path '%s'", expr, path)
	}

	match := pattern.FindStringSubmatch(id)
	if len(match) < len(pattern.SubexpNames()) {
		return nil, errors.Errorf("failed to parse '%s' against the path '%s'", id, path)
	}

	result := map[string]string{}
	for i, regexpGroupName := range pattern.SubexpNames() {
		if i > 0 && regexpGroupName != "" {
			originalName := regexNames[regexpGroupName]
			result[originalName] = match[i]
		}
	}
	return result, nil
}

// ParseToken extracts the module, version, and resource/type name from a resource token.
func ParseToken(tok string) (string, string, string, error) {
	parts := strings.Split(tok, ":")
	if len(parts) != 3 {
		return "", "", "", errors.Errorf("malformed token '%s'", tok)
	}
	provider, module, resource := parts[0], parts[1], parts[2]
	if provider == "azure-native" {
		return module, "", resource, nil
	}
	providerParts := strings.Split(provider, "_")
	if len(providerParts) != 3 {
		return "", "", "", errors.Errorf("malformed token '%s'", tok)
	}
	return module, providerParts[2], resource, nil
}

// BuildToken constructs a resource token from the given module, API version, and resource name.
func BuildToken(mod string, apiVersion string, name string) string {
	providerName := "azure-native"
	if apiVersion != "" {
		providerName = parameterizedProviderName(mod, apiVersion)
	}
	return strings.Join([]string{providerName, mod, name}, ":")
}

func parameterizedProviderName(targetModule, targetApiVersion string) string {
	return strings.Join([]string{"azure-native", targetModule, targetApiVersion}, "_")
}
