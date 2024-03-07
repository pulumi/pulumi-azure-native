package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewClusterManager(ctx, "clusterManager", &networkcloud.ClusterManagerArgs{
			AnalyticsWorkspaceId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
			ClusterManagerName:   pulumi.String("clusterManagerName"),
			FabricControllerId:   pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/fabricControllerName"),
			Location:             pulumi.String("location"),
			ManagedResourceGroupConfiguration: &networkcloud.ManagedResourceGroupConfigurationArgs{
				Location: pulumi.String("East US"),
				Name:     pulumi.String("my-managed-rg"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
