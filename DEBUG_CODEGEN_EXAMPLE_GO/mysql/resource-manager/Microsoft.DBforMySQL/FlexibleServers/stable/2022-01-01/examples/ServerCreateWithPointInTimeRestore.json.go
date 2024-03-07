package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
			CreateMode:         pulumi.String("PointInTimeRestore"),
			Location:           pulumi.String("SoutheastAsia"),
			ResourceGroupName:  pulumi.String("TargetResourceGroup"),
			RestorePointInTime: pulumi.String("2021-06-24T00:00:37.467Z"),
			ServerName:         pulumi.String("targetserver"),
			Sku: &dbformysql.SkuArgs{
				Name: pulumi.String("Standard_D14_v2"),
				Tier: pulumi.String("GeneralPurpose"),
			},
			SourceServerResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver"),
			Tags: pulumi.StringMap{
				"num": pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
