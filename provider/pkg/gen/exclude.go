// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azure-native:chaos:Experiment",
	"azure-native:costmanagement:Budget",
	"azure-native:costmanagement:Report",
	"azure-native:datafactory:Pipeline", // go codegen goes full CPU and doesn't return

	"azure-native:hybridcompute:GuestConfigurationHCRPAssignment",    // python name mismatch
}
var excludeRegexes []*regexp.Regexp

func init() {
	for _, pattern := range excludeResourcePatterns {
		excludeRegexes = append(excludeRegexes, regexp.MustCompile(pattern))
	}
}

func ShouldExclude(pulumiToken string) bool {
	for _, re := range excludeRegexes {
		if re.MatchString(pulumiToken) {
			return true
		}
	}

	return false
}
