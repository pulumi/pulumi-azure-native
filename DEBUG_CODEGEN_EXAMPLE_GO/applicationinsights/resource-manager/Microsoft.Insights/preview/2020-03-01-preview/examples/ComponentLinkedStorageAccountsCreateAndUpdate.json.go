package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewComponentLinkedStorageAccount(ctx, "componentLinkedStorageAccount", &insights.ComponentLinkedStorageAccountArgs{
			LinkedStorageAccount: pulumi.String("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname"),
			ResourceGroupName:    pulumi.String("someResourceGroupName"),
			ResourceName:         pulumi.String("myComponent"),
			StorageType:          pulumi.String("ServiceProfiler"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
