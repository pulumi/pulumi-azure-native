package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewBackupVault(ctx, "backupVault", &netapp.BackupVaultArgs{
			AccountName:       pulumi.String("account1"),
			BackupVaultName:   pulumi.String("backupVault1"),
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("myRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
