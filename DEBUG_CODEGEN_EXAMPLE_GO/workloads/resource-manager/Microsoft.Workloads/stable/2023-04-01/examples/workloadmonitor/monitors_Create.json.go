package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewMonitor(ctx, "monitor", &workloads.MonitorArgs{
			AppLocation:                pulumi.String("westus"),
			Location:                   pulumi.String("westus"),
			LogAnalyticsWorkspaceArmId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
			ManagedResourceGroupConfiguration: &workloads.ManagedRGConfigurationArgs{
				Name: pulumi.String("myManagedRg"),
			},
			MonitorName:       pulumi.String("mySapMonitor"),
			MonitorSubnet:     pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RoutingPreference: pulumi.String("RouteAll"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
