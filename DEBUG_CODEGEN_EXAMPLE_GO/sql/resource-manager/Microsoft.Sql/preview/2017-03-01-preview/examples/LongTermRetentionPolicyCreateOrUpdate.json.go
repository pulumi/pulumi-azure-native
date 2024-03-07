package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewBackupLongTermRetentionPolicy(ctx, "backupLongTermRetentionPolicy", &sql.BackupLongTermRetentionPolicyArgs{
			DatabaseName:      pulumi.String("testDatabase"),
			MonthlyRetention:  pulumi.String("P1Y"),
			PolicyName:        pulumi.String("default"),
			ResourceGroupName: pulumi.String("resourceGroup"),
			ServerName:        pulumi.String("testserver"),
			WeekOfYear:        pulumi.Int(5),
			WeeklyRetention:   pulumi.String("P1M"),
			YearlyRetention:   pulumi.String("P5Y"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
