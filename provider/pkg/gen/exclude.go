package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azure-nextgen:billing:ReportByBillingAccount",
	"azure-nextgen:billing:ReportByDepartment",
	"azure-nextgen:costmanagement:ReportConfig",
	"azure-nextgen:costmanagement:Report",
	"azure-nextgen:costmanagement:Budget",
	"azure-nextgen:datamigration:ServiceTask",
	"azure-nextgen:datamigration:Task",
	"azure-nextgen:machinelearning:WebService",
	"azure-nextgen:machinelearningservices:LabelingJob",
	"azure-nextgen:media:Job",
	"azure-nextgen:migrate:MoveCollection",
	"azure-nextgen:migrate:MoveResource",
	"azure-nextgen:portal:Dashboard", // go codegen stack overflow

	"azure-nextgen:datafactory:Pipeline", // go codegen goes full CPU and doesn't return

	"azure-nextgen:hybridcompute:GuestConfigurationHCRPAssignment",    // python name mismatch
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
