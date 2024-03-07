package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkRack(ctx, "networkRack", &managednetworkfabric.NetworkRackArgs{
			Annotation:        pulumi.String("null"),
			Location:          pulumi.String("eastus"),
			NetworkFabricId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabrics/networkFabricName"),
			NetworkRackName:   pulumi.String("networkRackName"),
			NetworkRackSku:    pulumi.String("RackSKU"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"keyID": pulumi.String("keyValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
