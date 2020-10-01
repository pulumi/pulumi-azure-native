package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azure-nextgen:billing/.*:ReportByBillingAccount",
	"azure-nextgen:billing/.*:ReportByDepartment",
	"azure-nextgen:costmanagement/.*:ReportConfig",
	"azure-nextgen:costmanagement/.*:Report",
	"azure-nextgen:costmanagement/.*:Budget",
	"azure-nextgen:datamigration/.*:Task",
	"azure-nextgen:machinelearning/.*:WebService",
	"azure-nextgen:media/latest:Job",
	"azure-nextgen:migrate/.*:MoveResource",
	"azure-nextgen:portal/v20200901preview:Dashboard", // go codegen stack overflow

	"azure-nextgen:hybridcompute/v20181120:GuestConfigurationHCRPAssignment",
	"azure-nextgen:hybridcompute/v20200625:GuestConfigurationHCRPAssignment", // python name mismatch
	"azure-nextgen:hybridcompute/latest:GuestConfigurationHCRPAssignment",    // python name mismatch
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
