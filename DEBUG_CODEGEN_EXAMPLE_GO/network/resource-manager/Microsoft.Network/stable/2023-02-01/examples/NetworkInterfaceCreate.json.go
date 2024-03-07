package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := network.NewNetworkInterface(ctx, "networkInterface", &network.NetworkInterfaceArgs{
DisableTcpStateTracking: pulumi.Bool(true),
EnableAcceleratedNetworking: pulumi.Bool(true),
IpConfigurations: network.NetworkInterfaceIPConfigurationArray{
interface{}{
Name: pulumi.String("ipconfig1"),
PublicIPAddress: &network.PublicIPAddressTypeArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
},
Subnet: &network.SubnetTypeArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
},
},
},
Location: pulumi.String("eastus"),
NetworkInterfaceName: pulumi.String("test-nic"),
ResourceGroupName: pulumi.String("rg1"),
})
if err != nil {
return err
}
return nil
})
}
