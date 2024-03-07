package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSubscriptionNetworkManagerConnection(ctx, "subscriptionNetworkManagerConnection", &network.SubscriptionNetworkManagerConnectionArgs{
			NetworkManagerConnectionName: pulumi.String("TestNMConnection"),
			NetworkManagerId:             pulumi.String("/subscriptions/subscriptionC/resourceGroup/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
