package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := network.NewHubVirtualNetworkConnection(ctx, "hubVirtualNetworkConnection", &network.HubVirtualNetworkConnectionArgs{
ConnectionName: pulumi.String("connection1"),
EnableInternetSecurity: pulumi.Bool(false),
RemoteVirtualNetwork: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
},
ResourceGroupName: pulumi.String("rg1"),
RoutingConfiguration: network.RoutingConfigurationResponse{
AssociatedRouteTable: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
},
InboundRouteMap: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
},
OutboundRouteMap: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
},
PropagatedRouteTables: interface{}{
Ids: network.SubResourceArray{
&network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
},
},
Labels: pulumi.StringArray{
pulumi.String("label1"),
pulumi.String("label2"),
},
},
VnetRoutes: interface{}{
StaticRoutes: network.StaticRouteArray{
&network.StaticRouteArgs{
AddressPrefixes: pulumi.StringArray{
pulumi.String("10.1.0.0/16"),
pulumi.String("10.2.0.0/16"),
},
Name: pulumi.String("route1"),
NextHopIpAddress: pulumi.String("10.0.0.68"),
},
&network.StaticRouteArgs{
AddressPrefixes: pulumi.StringArray{
pulumi.String("10.3.0.0/16"),
pulumi.String("10.4.0.0/16"),
},
Name: pulumi.String("route2"),
NextHopIpAddress: pulumi.String("10.0.0.65"),
},
},
StaticRoutesConfig: &network.StaticRoutesConfigArgs{
VnetLocalRouteOverrideCriteria: pulumi.String("Equal"),
},
},
},
VirtualHubName: pulumi.String("virtualHub1"),
})
if err != nil {
return err
}
return nil
})
}
