package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualHubRouteTableV2(ctx, "virtualHubRouteTableV2", &network.VirtualHubRouteTableV2Args{
			AttachedConnections: pulumi.StringArray{
				pulumi.String("All_Vnets"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			RouteTableName:    pulumi.String("virtualHubRouteTable1a"),
			Routes: network.VirtualHubRouteV2Array{
				&network.VirtualHubRouteV2Args{
					DestinationType: pulumi.String("CIDR"),
					Destinations: pulumi.StringArray{
						pulumi.String("20.10.0.0/16"),
						pulumi.String("20.20.0.0/16"),
					},
					NextHopType: pulumi.String("IPAddress"),
					NextHops: pulumi.StringArray{
						pulumi.String("10.0.0.68"),
					},
				},
				&network.VirtualHubRouteV2Args{
					DestinationType: pulumi.String("CIDR"),
					Destinations: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
					NextHopType: pulumi.String("IPAddress"),
					NextHops: pulumi.StringArray{
						pulumi.String("10.0.0.68"),
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
