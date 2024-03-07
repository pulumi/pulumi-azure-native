package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewBackup(ctx, "backup", &netapp.BackupArgs{
			AccountName:       pulumi.String("account1"),
			BackupName:        pulumi.String("backup1"),
			Label:             pulumi.String("myLabel"),
			Location:          pulumi.String("eastus"),
			PoolName:          pulumi.String("pool1"),
			ResourceGroupName: pulumi.String("myRG"),
			VolumeName:        pulumi.String("volume1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
