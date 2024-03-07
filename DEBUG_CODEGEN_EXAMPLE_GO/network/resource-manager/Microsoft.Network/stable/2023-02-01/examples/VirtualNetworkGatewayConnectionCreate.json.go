package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := network.NewVirtualNetworkGatewayConnection(ctx, "virtualNetworkGatewayConnection", &network.VirtualNetworkGatewayConnectionArgs{
ConnectionMode: pulumi.String("Default"),
ConnectionProtocol: pulumi.String("IKEv2"),
ConnectionType: pulumi.String("IPsec"),
DpdTimeoutSeconds: pulumi.Int(30),
EgressNatRules: []network.SubResourceArgs{
{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
},
},
EnableBgp: pulumi.Bool(false),
GatewayCustomBgpIpAddresses: []network.GatewayCustomBgpIpAddressIpConfigurationArgs{
{
CustomBgpIpAddress: pulumi.String("169.254.21.1"),
IpConfigurationId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
},
{
CustomBgpIpAddress: pulumi.String("169.254.21.3"),
IpConfigurationId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
},
},
IngressNatRules: []network.SubResourceArgs{
{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
},
},
IpsecPolicies: network.IpsecPolicyArray{
},
LocalNetworkGateway2: network.LocalNetworkGatewayResponse{
GatewayIpAddress: pulumi.String("x.x.x.x"),
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
LocalNetworkAddressSpace: &network.AddressSpaceArgs{
AddressPrefixes: pulumi.StringArray{
pulumi.String("10.1.0.0/16"),
},
},
Location: pulumi.String("centralus"),
Tags: nil,
},
Location: pulumi.String("centralus"),
ResourceGroupName: pulumi.String("rg1"),
RoutingWeight: pulumi.Int(0),
SharedKey: pulumi.String("Abc123"),
TrafficSelectorPolicies: network.TrafficSelectorPolicyArray{
},
UsePolicyBasedTrafficSelectors: pulumi.Bool(false),
VirtualNetworkGateway1: network.VirtualNetworkGatewayResponse{
ActiveActive: pulumi.Bool(false),
BgpSettings: &network.BgpSettingsArgs{
Asn: pulumi.Float64(65514),
BgpPeeringAddress: pulumi.String("10.0.1.30"),
PeerWeight: pulumi.Int(0),
},
EnableBgp: pulumi.Bool(false),
GatewayType: pulumi.String("Vpn"),
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
IpConfigurations: network.VirtualNetworkGatewayIPConfigurationArray{
interface{}{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
Name: pulumi.String("gwipconfig1"),
PrivateIPAllocationMethod: pulumi.String("Dynamic"),
PublicIPAddress: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
},
Subnet: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
},
},
},
Location: pulumi.String("centralus"),
Sku: &network.VirtualNetworkGatewaySkuArgs{
Name: pulumi.String("VpnGw1"),
Tier: pulumi.String("VpnGw1"),
},
Tags: nil,
VpnType: pulumi.String("RouteBased"),
},
VirtualNetworkGatewayConnectionName: pulumi.String("connS2S"),
})
if err != nil {
return err
}
return nil
})
}
