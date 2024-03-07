package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewConnectionMonitor(ctx, "connectionMonitor", &network.ConnectionMonitorArgs{
			ConnectionMonitorName: pulumi.String("cm1"),
			Endpoints: []network.ConnectionMonitorEndpointArgs{
				{
					Name:       pulumi.String("source"),
					ResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
				},
				{
					Address: pulumi.String("bing.com"),
					Name:    pulumi.String("destination"),
				},
			},
			Location:           pulumi.String("eastus"),
			NetworkWatcherName: pulumi.String("nw1"),
			ResourceGroupName:  pulumi.String("rg1"),
			TestConfigurations: []network.ConnectionMonitorTestConfigurationArgs{
				{
					Name:     pulumi.String("tcp"),
					Protocol: pulumi.String("Tcp"),
					TcpConfiguration: {
						Port: pulumi.Int(80),
					},
					TestFrequencySec: pulumi.Int(60),
				},
			},
			TestGroups: []network.ConnectionMonitorTestGroupArgs{
				{
					Destinations: pulumi.StringArray{
						pulumi.String("destination"),
					},
					Name: pulumi.String("tg"),
					Sources: pulumi.StringArray{
						pulumi.String("source"),
					},
					TestConfigurations: pulumi.StringArray{
						pulumi.String("tcp"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
