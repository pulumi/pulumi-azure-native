package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hanaonazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hanaonazure.NewSapMonitor(ctx, "sapMonitor", &hanaonazure.SapMonitorArgs{
			EnableCustomerAnalytics:        pulumi.Bool(true),
			Location:                       pulumi.String("westus"),
			LogAnalyticsWorkspaceArmId:     pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
			LogAnalyticsWorkspaceId:        pulumi.String("00000000-0000-0000-0000-000000000000"),
			LogAnalyticsWorkspaceSharedKey: pulumi.String("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000=="),
			MonitorSubnet:                  pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			ResourceGroupName:              pulumi.String("myResourceGroup"),
			SapMonitorName:                 pulumi.String("mySapMonitor"),
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
