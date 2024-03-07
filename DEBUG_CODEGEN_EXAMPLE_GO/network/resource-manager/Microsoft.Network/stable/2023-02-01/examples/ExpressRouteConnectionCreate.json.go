package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := network.NewExpressRouteConnection(ctx, "expressRouteConnection", &network.ExpressRouteConnectionArgs{
AuthorizationKey: pulumi.String("authorizationKey"),
ConnectionName: pulumi.String("connectionName"),
ExpressRouteCircuitPeering: &network.ExpressRouteCircuitPeeringIdArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
},
ExpressRouteGatewayName: pulumi.String("gateway-2"),
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName"),
Name: pulumi.String("connectionName"),
ResourceGroupName: pulumi.String("resourceGroupName"),
RoutingConfiguration: network.RoutingConfigurationResponse{
AssociatedRouteTable: &network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
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
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
},
&network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
},
&network.SubResourceArgs{
Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
},
},
Labels: pulumi.StringArray{
pulumi.String("label1"),
pulumi.String("label2"),
},
},
},
RoutingWeight: pulumi.Int(2),
})
if err != nil {
return err
}
return nil
})
}
