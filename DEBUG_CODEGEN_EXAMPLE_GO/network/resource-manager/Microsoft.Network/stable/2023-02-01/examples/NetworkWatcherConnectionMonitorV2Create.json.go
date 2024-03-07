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
					Name:       pulumi.String("vm1"),
					ResourceId: pulumi.String("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
				{
					Filter: {
						Items: network.ConnectionMonitorEndpointFilterItemArray{
							{
								Address: pulumi.String("npmuser"),
								Type:    pulumi.String("AgentAddress"),
							},
						},
						Type: pulumi.String("Include"),
					},
					Name:       pulumi.String("CanaryWorkspaceVamshi"),
					ResourceId: pulumi.String("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace"),
				},
				{
					Address: pulumi.String("bing.com"),
					Name:    pulumi.String("bing"),
				},
				{
					Address: pulumi.String("google.com"),
					Name:    pulumi.String("google"),
				},
			},
			NetworkWatcherName: pulumi.String("nw1"),
			Outputs:            network.ConnectionMonitorOutputTypeArray{},
			ResourceGroupName:  pulumi.String("rg1"),
			TestConfigurations: []network.ConnectionMonitorTestConfigurationArgs{
				{
					Name:     pulumi.String("testConfig1"),
					Protocol: pulumi.String("Tcp"),
					TcpConfiguration: {
						DisableTraceRoute: pulumi.Bool(false),
						Port:              pulumi.Int(80),
					},
					TestFrequencySec: pulumi.Int(60),
				},
			},
			TestGroups: []network.ConnectionMonitorTestGroupArgs{
				{
					Destinations: pulumi.StringArray{
						pulumi.String("bing"),
						pulumi.String("google"),
					},
					Disable: pulumi.Bool(false),
					Name:    pulumi.String("test1"),
					Sources: pulumi.StringArray{
						pulumi.String("vm1"),
						pulumi.String("CanaryWorkspaceVamshi"),
					},
					TestConfigurations: pulumi.StringArray{
						pulumi.String("testConfig1"),
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
