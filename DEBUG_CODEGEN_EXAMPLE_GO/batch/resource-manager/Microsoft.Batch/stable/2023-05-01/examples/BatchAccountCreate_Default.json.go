package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewBatchAccount(ctx, "batchAccount", &batch.BatchAccountArgs{
			AccountName: pulumi.String("sampleacct"),
			AutoStorage: &batch.AutoStorageBasePropertiesArgs{
				StorageAccountId: pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
			},
			Location:          pulumi.String("japaneast"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
