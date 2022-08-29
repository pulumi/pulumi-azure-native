


package v20210301

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
	case "azure-native:network/v20210301:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20210301:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20210301:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20210301:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20210301:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20210301:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20210301:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20210301:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20210301:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20210301:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20210301:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20210301:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20210301:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20210301:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20210301:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20210301:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20210301:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20210301:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20210301:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20210301:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20210301:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20210301:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20210301:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20210301:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20210301:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20210301:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20210301:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20210301:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20210301:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20210301:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20210301:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20210301:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20210301:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20210301:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20210301:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20210301:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20210301:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20210301:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20210301:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20210301:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20210301:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20210301:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20210301:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20210301:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20210301:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20210301:Route":
		r = &Route{}
	case "azure-native:network/v20210301:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20210301:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20210301:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20210301:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20210301:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20210301:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20210301:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20210301:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20210301:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20210301:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20210301:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20210301:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20210301:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20210301:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20210301:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20210301:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20210301:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20210301:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20210301:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20210301:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20210301:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20210301:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20210301:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20210301:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20210301:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20210301:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20210301:WebApplicationFirewallPolicy":
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
		"network/v20210301",
		&module{version},
	)
}
