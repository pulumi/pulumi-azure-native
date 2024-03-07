package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := managednetworkfabric.NewNetworkFabric(ctx, "networkFabric", &managednetworkfabric.NetworkFabricArgs{
Annotation: pulumi.String("annotationValue"),
FabricASN: pulumi.Int(29249),
Ipv4Prefix: pulumi.String("10.18.0.0/19"),
Ipv6Prefix: pulumi.String("3FFE:FFFF:0:CD40::/59"),
Location: pulumi.String("eastuseuap"),
ManagementNetworkConfiguration: managednetworkfabric.ManagementNetworkConfigurationResponse{
InfrastructureVpnConfiguration: interface{}{
OptionAProperties: &managednetworkfabric.OptionAPropertiesArgs{
Mtu: pulumi.Int(5892),
PeerASN: pulumi.Int(42666),
PrimaryIpv4Prefix: pulumi.String("20.0.0.12/30"),
PrimaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::a8/126"),
SecondaryIpv4Prefix: pulumi.String("20.0.0.13/30"),
SecondaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::ac/126"),
VlanId: pulumi.Int(2724),
},
OptionBProperties: &managednetworkfabric.FabricOptionBPropertiesArgs{
ExportRouteTargets: pulumi.StringArray{
pulumi.String("65046:10039"),
},
ImportRouteTargets: pulumi.StringArray{
pulumi.String("65046:10039"),
},
},
PeeringOption: pulumi.String("OptionA"),
},
WorkloadVpnConfiguration: interface{}{
OptionAProperties: &managednetworkfabric.OptionAPropertiesArgs{
Mtu: pulumi.Int(5892),
PeerASN: pulumi.Int(42666),
PrimaryIpv4Prefix: pulumi.String("10.0.0.14/30"),
PrimaryIpv6Prefix: pulumi.String("2FFE:FFFF:0:CD30::a7/126"),
SecondaryIpv4Prefix: pulumi.String("10.0.0.15/30"),
SecondaryIpv6Prefix: pulumi.String("2FFE:FFFF:0:CD30::ac/126"),
VlanId: pulumi.Int(2724),
},
OptionBProperties: &managednetworkfabric.FabricOptionBPropertiesArgs{
ExportRouteTargets: pulumi.StringArray{
pulumi.String("65046:10050"),
},
ImportRouteTargets: pulumi.StringArray{
pulumi.String("65046:10050"),
},
},
PeeringOption: pulumi.String("OptionA"),
},
},
NetworkFabricControllerId: pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
NetworkFabricName: pulumi.String("FabricName"),
NetworkFabricSku: pulumi.String("M4-A400-A100-C16-aa"),
RackCount: pulumi.Int(4),
ResourceGroupName: pulumi.String("resourceGroupName"),
ServerCountPerRack: pulumi.Int(8),
Tags: pulumi.StringMap{
"key6468": pulumi.String(""),
},
TerminalServerConfiguration: &managednetworkfabric.TerminalServerConfigurationArgs{
Password: pulumi.String("xxxx"),
PrimaryIpv4Prefix: pulumi.String("20.0.0.12/30"),
PrimaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::a8/126"),
SecondaryIpv4Prefix: pulumi.String("20.0.0.13/30"),
SecondaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::ac/126"),
SerialNumber: pulumi.String("123456"),
Username: pulumi.String("username"),
},
})
if err != nil {
return err
}
return nil
})
}
