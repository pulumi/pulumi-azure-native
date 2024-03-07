package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformariadb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformariadb.NewServer(ctx, "server", &dbformariadb.ServerArgs{
			Location: pulumi.String("brazilsouth"),
			Properties: dbformariadb.ServerPropertiesForRestore{
				CreateMode:         "PointInTimeRestore",
				RestorePointInTime: "2017-12-14T00:00:37.467Z",
				SourceServerId:     "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMariaDB/servers/sourceserver",
			},
			ResourceGroupName: pulumi.String("TargetResourceGroup"),
			ServerName:        pulumi.String("targetserver"),
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
