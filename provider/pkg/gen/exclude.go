package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azure-native:billing:ReportByBillingAccount",
	"azure-native:billing:ReportByDepartment",
	"azure-native:costmanagement:ReportConfig",
	"azure-native:costmanagement:Report",
	"azure-native:costmanagement:Budget",
	"azure-native:datamigration:ServiceTask",
	"azure-native:datamigration:Task",
	"azure-native:machinelearning:WebService",
	"azure-native:machinelearningservices:LabelingJob",
	"azure-native:media:Job",
	"azure-native:migrate:MoveCollection",
	"azure-native:migrate:MoveResource",
	"azure-native:portal:Dashboard", // go codegen stack overflow

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
