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

	"azure-nextgen:network/.*:ApplicationGateway",
	"azure-nextgen:network/.*:InboundNatRule",
	"azure-nextgen:network/.*:LoadBalancer",
	"azure-nextgen:network/.*:NetworkInterface",
	"azure-nextgen:network/.*:NetworkSecurityGroup",
	"azure-nextgen:network/.*:PublicIPAddress",
	"azure-nextgen:network/.*:PrivateEndpoint",
	"azure-nextgen:network/.*:RouteTable",
	"azure-nextgen:network/.*:Subnet",
	"azure-nextgen:network/.*:VirtualNetwork",
	"azure-nextgen:network/.*:RouteFilter",
	"azure-nextgen:network/.*:ExpressRouteCircuit",
	"azure-nextgen:network/.*:ExpressRouteCircuitPeering",
	"azure-nextgen:network/.*:ExpressRouteCrossConnectionPeering",
	"azure-nextgen:network/.*:InterfaceEndpoint",
	"azure-nextgen:network/.*:NetworkInterfaceTapConfiguration",
	"azure-nextgen:network/.*:NetworkProfile",
	"azure-nextgen:network/.*:ServiceEndpointPolicy",
	"azure-nextgen:network/.*:VirtualNetworkTap",
	"azure-nextgen:network/.*:WebApplicationFirewallPolicy",
	"azure-nextgen:network/.*:PrivateLinkService",
	"azure-nextgen:network/.*:PrivateLinkServicePrivateEndpointConnection",
	"azure-nextgen:network/.*:LoadBalancerBackendAddressPool",
	"azure-nextgen:network/.*:VirtualHubIpConfiguration",
	"azure-nextgen:network/.*:VpnSite",
	"azure-nextgen:network/.*:ApplicationGatewayPrivateEndpointConnection",
	"azure-nextgen:network/.*:DscpConfiguration",

	"azure-nextgen:powerbidedicated/.*:CapacityDetails",

	"azure-nextgen:recoveryservices/.*:ReplicationFabric",
	"azure-nextgen:recoveryservices/.*:ReplicationProtectedItem",
	"azure-nextgen:recoveryservices/.*:ReplicationProtectionContainerMapping",

	"azure-nextgen:resources/.*:Deployment",
	"azure-nextgen:resources/.*:DeploymentAtScope",
	"azure-nextgen:resources/.*:DeploymentAtTenantScope",
	"azure-nextgen:resources/.*:DeploymentAtSubscriptionScope",
	// ^ recursive references in schema

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
