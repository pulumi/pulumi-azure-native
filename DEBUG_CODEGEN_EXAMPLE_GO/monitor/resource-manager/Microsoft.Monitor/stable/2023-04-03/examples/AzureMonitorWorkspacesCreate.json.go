package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/monitor/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := monitor.NewAzureMonitorWorkspace(ctx, "azureMonitorWorkspace", &monitor.AzureMonitorWorkspaceArgs{
			AzureMonitorWorkspaceName: pulumi.String("myAzureMonitorWorkspace"),
			Location:                  pulumi.String("eastus"),
			ResourceGroupName:         pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
