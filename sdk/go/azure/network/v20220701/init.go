


package v20220701

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
	case "azure-native:network/v20220701:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20220701:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20220701:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20220701:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20220701:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20220701:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20220701:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20220701:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network/v20220701:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20220701:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20220701:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20220701:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20220701:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20220701:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20220701:DnsForwardingRuleset":
		r = &DnsForwardingRuleset{}
	case "azure-native:network/v20220701:DnsResolver":
		r = &DnsResolver{}
	case "azure-native:network/v20220701:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20220701:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20220701:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20220701:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20220701:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20220701:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20220701:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20220701:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20220701:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20220701:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network/v20220701:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20220701:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20220701:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20220701:ForwardingRule":
		r = &ForwardingRule{}
	case "azure-native:network/v20220701:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20220701:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20220701:InboundEndpoint":
		r = &InboundEndpoint{}
	case "azure-native:network/v20220701:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20220701:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20220701:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20220701:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20220701:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20220701:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20220701:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network/v20220701:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20220701:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20220701:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20220701:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20220701:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20220701:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20220701:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20220701:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20220701:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20220701:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20220701:OutboundEndpoint":
		r = &OutboundEndpoint{}
	case "azure-native:network/v20220701:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20220701:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20220701:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20220701:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20220701:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20220701:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20220701:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20220701:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20220701:Route":
		r = &Route{}
	case "azure-native:network/v20220701:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20220701:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20220701:RouteMap":
		r = &RouteMap{}
	case "azure-native:network/v20220701:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20220701:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20220701:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network/v20220701:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20220701:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20220701:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20220701:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20220701:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20220701:StaticMember":
		r = &StaticMember{}
	case "azure-native:network/v20220701:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20220701:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network/v20220701:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20220701:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20220701:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20220701:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20220701:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20220701:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20220701:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20220701:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20220701:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20220701:VirtualNetworkLink":
		r = &VirtualNetworkLink{}
	case "azure-native:network/v20220701:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20220701:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20220701:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20220701:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20220701:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20220701:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20220701:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20220701:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20220701:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20220701:WebApplicationFirewallPolicy":
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
		"network/v20220701",
		&module{version},
	)
}
