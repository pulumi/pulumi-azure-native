


package v20220501

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
	case "azure-native:network/v20220501:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20220501:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20220501:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20220501:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20220501:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20220501:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20220501:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20220501:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network/v20220501:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20220501:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20220501:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20220501:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20220501:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20220501:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20220501:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20220501:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20220501:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20220501:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20220501:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20220501:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20220501:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20220501:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20220501:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20220501:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network/v20220501:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20220501:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20220501:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20220501:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20220501:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20220501:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20220501:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20220501:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20220501:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20220501:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20220501:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20220501:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network/v20220501:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20220501:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20220501:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20220501:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20220501:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20220501:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20220501:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20220501:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20220501:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20220501:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20220501:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20220501:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20220501:Policy":
		r = &Policy{}
	case "azure-native:network/v20220501:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20220501:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20220501:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20220501:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20220501:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20220501:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20220501:Route":
		r = &Route{}
	case "azure-native:network/v20220501:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20220501:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20220501:RouteMap":
		r = &RouteMap{}
	case "azure-native:network/v20220501:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20220501:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20220501:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network/v20220501:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20220501:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20220501:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20220501:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20220501:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20220501:StaticMember":
		r = &StaticMember{}
	case "azure-native:network/v20220501:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20220501:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network/v20220501:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20220501:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20220501:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20220501:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20220501:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20220501:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20220501:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20220501:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20220501:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20220501:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20220501:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20220501:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20220501:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20220501:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20220501:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20220501:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20220501:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20220501:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20220501:WebApplicationFirewallPolicy":
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
		"network/v20220501",
		&module{version},
	)
}
