package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewLinkedStorageAccount(ctx, "linkedStorageAccount", &operationalinsights.LinkedStorageAccountArgs{
			DataSourceType:    pulumi.String("CustomLogs"),
			ResourceGroupName: pulumi.String("mms-eus"),
			StorageAccountIds: pulumi.StringArray{
				pulumi.String("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
				pulumi.String("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageB"),
			},
			WorkspaceName: pulumi.String("testLinkStorageAccountsWS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
