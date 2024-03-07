package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearning/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearning.NewWorkspace(ctx, "workspace", &machinelearning.WorkspaceArgs{
			Location:          pulumi.String("West Europe"),
			OwnerEmail:        pulumi.String("abc@microsoft.com"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &machinelearning.SkuArgs{
				Name: pulumi.String("Enterprise"),
				Tier: pulumi.String("Enterprise"),
			},
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("TagValue1"),
			},
			UserStorageAccountId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/teststorage"),
			WorkspaceName:        pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
