package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkWatcher(ctx, "networkWatcher", &network.NetworkWatcherArgs{
			Location:           pulumi.String("eastus"),
			NetworkWatcherName: pulumi.String("nw1"),
			ResourceGroupName:  pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
