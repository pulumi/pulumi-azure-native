package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningexperimentation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningexperimentation.NewAccount(ctx, "account", &machinelearningexperimentation.AccountArgs{
			AccountName:       pulumi.String("accountcrud5678"),
			KeyVaultId:        pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.KeyVault/vaults/testkv"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("accountcrud-1234"),
			StorageAccount: &machinelearningexperimentation.StorageAccountPropertiesArgs{
				AccessKey:        pulumi.String("key"),
				StorageAccountId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
			},
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("TagValue1"),
			},
			VsoAccountId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/microsoft.visualstudio/account/vsotest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
