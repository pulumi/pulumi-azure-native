package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azure-nextgen:attestation/.*:AttestationProvider",

	"azure-nextgen:billing/.*:ReportByBillingAccount",
	"azure-nextgen:billing/.*:ReportByDepartment",

	"azure-nextgen:batch/.*:Certificate",
	"azure-nextgen:batch/.*:Pool",

	"azure-nextgen:batchai/v20180501:Workspace",
	"azure-nextgen:batchai/latest:Workspace",

	"azure-nextgen:costmanagement/v20180531:ReportConfig",
	"azure-nextgen:costmanagement/latest:ReportConfig",
	"azure-nextgen:costmanagement/.*:Export",
	"azure-nextgen:costmanagement/.*:ReportConfigByResourceGroupName",
	"azure-nextgen:costmanagement/.*:Export",
	"azure-nextgen:costmanagement/.*:ViewByScope",
	"azure-nextgen:costmanagement/.*:View",
	"azure-nextgen:costmanagement/.*:Report",
	"azure-nextgen:costmanagement/.*:Budget",

	"azure-nextgen:databox/.*:Job",

	"azure-nextgen:datamigration/.*:Task",
	"azure-nextgen:datamigration/.*:ServiceTask",

	"azure-nextgen:hybridcompute/.*:Machine",

	"azure-nextgen:machinelearning/.*:WebService",

	"azure-nextgen:media/.*:Job",

	"azure-nextgen:management/.*:DeploymentAtManagementGroupScope",
	"azure-nextgen:management/.*:ManagementGroup",

	"azure-nextgen:migrate/.*:MoveResource",

	"azure-nextgen:powerbidedicated/.*:CapacityDetails",

	"azure-nextgen:recoveryservices/.*:ReplicationFabric",
	"azure-nextgen:recoveryservices/.*:ReplicationProtectedItem",
	"azure-nextgen:recoveryservices/.*:ReplicationProtectionContainerMapping",

	"azure-nextgen:resources/.*:Deployment",
	"azure-nextgen:resources/.*:DeploymentAtScope",
	"azure-nextgen:resources/.*:DeploymentAtTenantScope",
	"azure-nextgen:resources/.*:DeploymentAtSubscriptionScope",
	// ^ recursive references in schema
	// TODO: Recompile this list after https://github.com/pulumi/pulumi/pull/5404 is merged.

	"azure-nextgen:botservice/v20200602:BotConnection", // malformed body

	"azure-nextgen:hybridcompute/v20181120:GuestConfigurationHCRPAssignment",
	"azure-nextgen:hybridcompute/v20200625:GuestConfigurationHCRPAssignment", // python name mismatch
	"azure-nextgen:hybridcompute/latest:GuestConfigurationHCRPAssignment",    // python name mismatch
}

// ShouldExclude checks if the given pulumi resource token matches known-broken
// resources with respect to program generation.
func ShouldExclude(pulumiResourceToken string) bool {
	var excludeRegexes []*regexp.Regexp
	for _, pattern := range excludeResourcePatterns {
		excludeRegexes = append(excludeRegexes, regexp.MustCompile(pattern))
	}
	for _, re := range excludeRegexes {
		if re.MatchString(pulumiResourceToken) {
			return true
		}
	}

	return false
}
