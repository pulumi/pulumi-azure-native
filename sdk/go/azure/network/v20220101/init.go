


package v20220101

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:network/v20220101:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20220101:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20220101:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20220101:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20220101:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20220101:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20220101:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20220101:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network/v20220101:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20220101:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20220101:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20220101:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20220101:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20220101:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20220101:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20220101:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20220101:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20220101:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20220101:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20220101:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20220101:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20220101:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20220101:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20220101:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network/v20220101:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20220101:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20220101:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20220101:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20220101:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20220101:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20220101:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20220101:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20220101:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20220101:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20220101:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20220101:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network/v20220101:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20220101:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20220101:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20220101:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20220101:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20220101:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20220101:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20220101:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20220101:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20220101:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20220101:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20220101:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20220101:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20220101:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20220101:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20220101:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20220101:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20220101:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20220101:Route":
		r = &Route{}
	case "azure-native:network/v20220101:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20220101:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20220101:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20220101:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20220101:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network/v20220101:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20220101:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20220101:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20220101:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20220101:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20220101:StaticMember":
		r = &StaticMember{}
	case "azure-native:network/v20220101:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20220101:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network/v20220101:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20220101:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20220101:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20220101:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20220101:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20220101:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20220101:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20220101:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20220101:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20220101:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20220101:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20220101:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20220101:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20220101:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20220101:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20220101:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20220101:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20220101:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20220101:WebApplicationFirewallPolicy":
		r = &WebApplicationFirewallPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network/v20220101",
		&module{version},
	)
}
