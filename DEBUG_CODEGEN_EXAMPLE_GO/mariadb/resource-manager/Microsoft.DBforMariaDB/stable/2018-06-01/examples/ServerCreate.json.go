package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformariadb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformariadb.NewServer(ctx, "server", &dbformariadb.ServerArgs{
			Location: pulumi.String("westus"),
			Properties: dbformariadb.ServerPropertiesForDefaultCreate{
				AdministratorLogin:         "cloudsa",
				AdministratorLoginPassword: "<administratorLoginPassword>",
				CreateMode:                 "Default",
				MinimalTlsVersion:          "TLS1_2",
				SslEnforcement:             dbformariadb.SslEnforcementEnumEnabled,
				StorageProfile: dbformariadb.StorageProfile{
					BackupRetentionDays: 7,
					GeoRedundantBackup:  "Enabled",
					StorageMB:           128000,
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("mariadbtestsvc4"),
			Sku: &dbformariadb.SkuArgs{
				Capacity: pulumi.Int(2),
				Family:   pulumi.String("Gen5"),
				Name:     pulumi.String("GP_Gen5_2"),
				Tier:     pulumi.String("GeneralPurpose"),
			},
			Tags: pulumi.StringMap{
				"ElasticServer": pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
