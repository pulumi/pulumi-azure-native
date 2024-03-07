package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewHubRouteTable(ctx, "hubRouteTable", &network.HubRouteTableArgs{
			Labels: pulumi.StringArray{
				pulumi.String("label1"),
				pulumi.String("label2"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			RouteTableName:    pulumi.String("hubRouteTable1"),
			Routes: network.HubRouteArray{
				&network.HubRouteArgs{
					DestinationType: pulumi.String("CIDR"),
					Destinations: pulumi.StringArray{
						pulumi.String("10.0.0.0/8"),
						pulumi.String("20.0.0.0/8"),
						pulumi.String("30.0.0.0/8"),
					},
					Name:        pulumi.String("route1"),
					NextHop:     pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1"),
					NextHopType: pulumi.String("ResourceId"),
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
