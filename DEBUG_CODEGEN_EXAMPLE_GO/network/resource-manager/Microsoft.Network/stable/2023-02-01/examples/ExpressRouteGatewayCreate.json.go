package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteGateway(ctx, "expressRouteGateway", &network.ExpressRouteGatewayArgs{
			AllowNonVirtualWanTraffic: pulumi.Bool(false),
			AutoScaleConfiguration: &network.ExpressRouteGatewayPropertiesAutoScaleConfigurationArgs{
				Bounds: &network.ExpressRouteGatewayPropertiesBoundsArgs{
					Min: pulumi.Int(3),
				},
			},
			ExpressRouteGatewayName: pulumi.String("gateway-2"),
			Location:                pulumi.String("westus"),
			ResourceGroupName:       pulumi.String("resourceGroupName"),
			VirtualHub: &network.VirtualHubIdArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupId/providers/Microsoft.Network/virtualHubs/virtualHubName"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
