package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewElasticPool(ctx, "elasticPool", &sql.ElasticPoolArgs{
			ElasticPoolName: pulumi.String("sqlcrudtest-8102"),
			Location:        pulumi.String("Japan East"),
			PerDatabaseSettings: &sql.ElasticPoolPerDatabaseSettingsArgs{
				MaxCapacity: pulumi.Float64(2),
				MinCapacity: pulumi.Float64(0.25),
			},
			ResourceGroupName: pulumi.String("sqlcrudtest-2369"),
			ServerName:        pulumi.String("sqlcrudtest-8069"),
			Sku: &sql.SkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("GP_Gen4_2"),
				Tier:     pulumi.String("GeneralPurpose"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
