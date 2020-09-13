package gen

import (
	"regexp"
)

// excludeResourcePatterns lists resources being skipped due to known codegen issues.
var excludeResourcePatterns = []string{
	"azurerm:attestation/.*:AttestationProvider",

	"azurerm:billing/.*:ReportByBillingAccount",
	"azurerm:billing/.*:ReportByDepartment",

	"azurerm:batch/.*:Certificate",
	"azurerm:batch/.*:Pool",

	"azurerm:batchai/v20180501:Workspace",
	"azurerm:batchai/latest:Workspace",

	"azurerm:costmanagement/v20180531:ReportConfig",
	"azurerm:costmanagement/latest:ReportConfig",
	"azurerm:costmanagement/.*:Export",
	"azurerm:costmanagement/.*:ReportConfigByResourceGroupName",
	"azurerm:costmanagement/.*:Export",
	"azurerm:costmanagement/.*:ViewByScope",
	"azurerm:costmanagement/.*:View",
	"azurerm:costmanagement/.*:Report",
	"azurerm:costmanagement/.*:Budget",

	"azurerm:databox/.*:Job",

	"azurerm:datamigration/.*:Task",
	"azurerm:datamigration/.*:ServiceTask",

	"azurerm:hybridcompute/.*:Machine",

	"azurerm:machinelearning/.*:WebService",

	"azurerm:media/.*:Job",

	"azurerm:management/.*:DeploymentAtManagementGroupScope",
	"azurerm:management/.*:ManagementGroup",

	"azurerm:migrate/.*:MoveResource",

	"azurerm:network/.*:ApplicationGateway",
	"azurerm:network/.*:InboundNatRule",
	"azurerm:network/.*:LoadBalancer",
	"azurerm:network/.*:NetworkInterface",
	"azurerm:network/.*:NetworkSecurityGroup",
	"azurerm:network/.*:PublicIPAddress",
	"azurerm:network/.*:PrivateEndpoint",
	"azurerm:network/.*:RouteTable",
	"azurerm:network/.*:Subnet",
	"azurerm:network/.*:VirtualNetwork",
	"azurerm:network/.*:RouteFilter",
	"azurerm:network/.*:ExpressRouteCircuit",
	"azurerm:network/.*:ExpressRouteCircuitPeering",
	"azurerm:network/.*:ExpressRouteCrossConnectionPeering",
	"azurerm:network/.*:InterfaceEndpoint",
	"azurerm:network/.*:NetworkInterfaceTapConfiguration",
	"azurerm:network/.*:NetworkProfile",
	"azurerm:network/.*:ServiceEndpointPolicy",
	"azurerm:network/.*:VirtualNetworkTap",
	"azurerm:network/.*:WebApplicationFirewallPolicy",
	"azurerm:network/.*:PrivateLinkService",
	"azurerm:network/.*:PrivateLinkServicePrivateEndpointConnection",
	"azurerm:network/.*:LoadBalancerBackendAddressPool",
	"azurerm:network/.*:VirtualHubIpConfiguration",
	"azurerm:network/.*:VpnSite",
	"azurerm:network/.*:ApplicationGatewayPrivateEndpointConnection",
	"azurerm:network/.*:DscpConfiguration",

	"azurerm:powerbidedicated/.*:CapacityDetails",

	"azurerm:recoveryservices/.*:ReplicationFabric",
	"azurerm:recoveryservices/.*:ReplicationProtectedItem",
	"azurerm:recoveryservices/.*:ReplicationProtectionContainerMapping",

	"azurerm:resources/.*:Deployment",
	"azurerm:resources/.*:DeploymentAtScope",
	"azurerm:resources/.*:DeploymentAtTenantScope",
	"azurerm:resources/.*:DeploymentAtSubscriptionScope",
	// ^ recursive references in schema

	"azurerm:botservice/v20200602:BotConnection", // malformed body

	"azurerm:hybridcompute/v20181120:GuestConfigurationHCRPAssignment",
	"azurerm:hybridcompute/v20200625:GuestConfigurationHCRPAssignment", // python name mismatch
	"azurerm:hybridcompute/latest:GuestConfigurationHCRPAssignment",    // python name mismatch
}
var excludeRegexes []*regexp.Regexp

func init() {
	for _, pattern := range excludeResourcePatterns {
		excludeRegexes = append(excludeRegexes, regexp.MustCompile(pattern))
	}
}

func shouldExclude(pulumiToken string) bool {
	for _, re := range excludeRegexes {
		if re.MatchString(pulumiToken) {
			return true
		}
	}

	return false
}
