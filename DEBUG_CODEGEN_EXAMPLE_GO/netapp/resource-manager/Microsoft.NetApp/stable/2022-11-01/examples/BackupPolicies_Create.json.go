package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewBackupPolicy(ctx, "backupPolicy", &netapp.BackupPolicyArgs{
			AccountName:          pulumi.String("account1"),
			BackupPolicyName:     pulumi.String("backupPolicyName"),
			DailyBackupsToKeep:   pulumi.Int(10),
			Enabled:              pulumi.Bool(true),
			Location:             pulumi.String("westus"),
			MonthlyBackupsToKeep: pulumi.Int(10),
			ResourceGroupName:    pulumi.String("myRG"),
			WeeklyBackupsToKeep:  pulumi.Int(10),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
