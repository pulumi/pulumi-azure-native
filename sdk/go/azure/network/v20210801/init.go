


package v20210801

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
	case "azure-native:network/v20210801:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20210801:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20210801:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20210801:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20210801:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20210801:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network/v20210801:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20210801:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20210801:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20210801:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20210801:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20210801:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20210801:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20210801:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20210801:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20210801:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20210801:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20210801:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20210801:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20210801:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network/v20210801:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20210801:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20210801:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20210801:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20210801:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20210801:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20210801:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20210801:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20210801:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20210801:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20210801:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20210801:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20210801:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20210801:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20210801:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20210801:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20210801:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20210801:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20210801:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20210801:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20210801:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20210801:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20210801:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20210801:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20210801:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20210801:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20210801:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20210801:Route":
		r = &Route{}
	case "azure-native:network/v20210801:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20210801:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20210801:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20210801:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20210801:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20210801:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20210801:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20210801:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20210801:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20210801:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20210801:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20210801:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20210801:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20210801:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20210801:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20210801:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20210801:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20210801:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20210801:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20210801:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20210801:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20210801:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20210801:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20210801:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20210801:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20210801:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20210801:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20210801:WebApplicationFirewallPolicy":
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
		"network/v20210801",
		&module{version},
	)
}
