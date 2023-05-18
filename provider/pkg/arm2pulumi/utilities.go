// Copyright 2021, Pulumi Corporation.  All rights reserved.

package arm2pulumi

import (
	"errors"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

func unquote(value string) string {
	if value == "" {
		return ""
	}

	if strings.HasPrefix(value, "\"") {
		if strings.HasSuffix(value, "\"") {
			return value[1 : len(value)-1]
		}
		return value[1:]
	} else if strings.HasPrefix(value, "'") {
		if strings.HasSuffix(value, "'") {
			return value[1 : len(value)-1]
		}
		return value[1:]
	}

	return value
}

func extractResourceNameParameter(res *resources.AzureAPIResource) (*resources.AzureAPIParameter, error) {
	path := res.Path
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 4; i-- {
		part := parts[i]
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name := part[1 : len(part)-1]
			for _, v := range res.PutParameters {
				if v.Name == name {
					prop := v.Value
					// We expect names to be always a required string.
					if prop.Type != "string" {
						break
					}
					if !v.IsRequired {
						break
					}
					return &v, nil
				}
			}
		}
	}
	return nil, errors.New("no resource name parameter found")
}

var rgPathRegex = regexp.MustCompile(`(?i)/.*/resourceGroups/\{(resourceGroupName)\}/providers/.*`)

func extractResourceGroupNameParameter(res *resources.AzureAPIResource) *resources.AzureAPIParameter {
	path := res.Path
	var resouceGroupParamName string
	if rgPathRegex.MatchString(path) {
		subMatches := rgPathRegex.FindStringSubmatch(path)
		if len(subMatches) > 1 {
			resouceGroupParamName = subMatches[1]
		}
	}

	if resouceGroupParamName == "" {
		return nil
	}

	for i, param := range res.PutParameters {
		if strings.EqualFold(param.Name, resouceGroupParamName) {
			return &res.PutParameters[i]
		}
	}
	return nil
}
