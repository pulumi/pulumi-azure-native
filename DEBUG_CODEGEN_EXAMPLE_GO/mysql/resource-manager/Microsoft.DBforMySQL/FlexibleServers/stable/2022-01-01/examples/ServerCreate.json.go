package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
			AdministratorLogin:         pulumi.String("cloudsa"),
			AdministratorLoginPassword: pulumi.String("your_password"),
			AvailabilityZone:           pulumi.String("1"),
			Backup: &dbformysql.BackupArgs{
				BackupRetentionDays: pulumi.Int(7),
				GeoRedundantBackup:  pulumi.String("Disabled"),
			},
			CreateMode: pulumi.String("Default"),
			HighAvailability: &dbformysql.HighAvailabilityArgs{
				Mode:                    pulumi.String("ZoneRedundant"),
				StandbyAvailabilityZone: pulumi.String("3"),
			},
			Location:          pulumi.String("southeastasia"),
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("mysqltestserver"),
			Sku: &dbformysql.SkuArgs{
				Name: pulumi.String("Standard_D2ds_v4"),
				Tier: pulumi.String("GeneralPurpose"),
			},
			Storage: &dbformysql.StorageArgs{
				AutoGrow:      pulumi.String("Disabled"),
				Iops:          pulumi.Int(600),
				StorageSizeGB: pulumi.Int(100),
			},
			Tags: pulumi.StringMap{
				"num": pulumi.String("1"),
			},
			Version: pulumi.String("5.7"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
