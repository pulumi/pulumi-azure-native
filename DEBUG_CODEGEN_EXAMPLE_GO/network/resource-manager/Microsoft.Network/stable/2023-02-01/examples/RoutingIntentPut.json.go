package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRoutingIntent(ctx, "routingIntent", &network.RoutingIntentArgs{
			ResourceGroupName: pulumi.String("rg1"),
			RoutingIntentName: pulumi.String("Intent1"),
			RoutingPolicies: []network.RoutingPolicyArgs{
				{
					Destinations: pulumi.StringArray{
						pulumi.String("Internet"),
					},
					Name:    pulumi.String("InternetTraffic"),
					NextHop: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
				},
				{
					Destinations: pulumi.StringArray{
						pulumi.String("PrivateTraffic"),
					},
					Name:    pulumi.String("PrivateTrafficPolicy"),
					NextHop: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
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
