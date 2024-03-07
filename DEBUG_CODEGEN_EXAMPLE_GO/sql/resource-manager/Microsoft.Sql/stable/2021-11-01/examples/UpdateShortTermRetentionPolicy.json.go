package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewBackupShortTermRetentionPolicy(ctx, "backupShortTermRetentionPolicy", &sql.BackupShortTermRetentionPolicyArgs{
			DatabaseName:              pulumi.String("testdb"),
			DiffBackupIntervalInHours: pulumi.Int(24),
			PolicyName:                pulumi.String("default"),
			ResourceGroupName:         pulumi.String("resourceGroup"),
			RetentionDays:             pulumi.Int(7),
			ServerName:                pulumi.String("testsvr"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
