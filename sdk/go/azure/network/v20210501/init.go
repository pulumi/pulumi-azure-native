


package v20210501

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
	case "azure-native:network/v20210501:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20210501:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20210501:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20210501:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20210501:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20210501:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20210501:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20210501:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20210501:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20210501:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20210501:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20210501:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20210501:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20210501:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20210501:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20210501:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20210501:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20210501:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20210501:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20210501:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20210501:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20210501:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20210501:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20210501:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20210501:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20210501:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20210501:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20210501:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20210501:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20210501:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20210501:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20210501:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20210501:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20210501:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20210501:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20210501:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20210501:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20210501:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20210501:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20210501:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20210501:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20210501:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20210501:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20210501:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20210501:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20210501:Route":
		r = &Route{}
	case "azure-native:network/v20210501:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20210501:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20210501:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20210501:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20210501:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20210501:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20210501:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20210501:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20210501:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20210501:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20210501:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20210501:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20210501:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20210501:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20210501:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20210501:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20210501:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20210501:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20210501:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20210501:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20210501:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20210501:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20210501:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20210501:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20210501:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20210501:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20210501:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20210501:WebApplicationFirewallPolicy":
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
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network/v20210501",
		&module{version},
	)
}
