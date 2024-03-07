package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewElasticPool(ctx, "elasticPool", &sql.ElasticPoolArgs{
			ElasticPoolName:              pulumi.String("sqlcrudtest-8102"),
			HighAvailabilityReplicaCount: pulumi.Int(2),
			Location:                     pulumi.String("Japan East"),
			ResourceGroupName:            pulumi.String("sqlcrudtest-2369"),
			ServerName:                   pulumi.String("sqlcrudtest-8069"),
			Sku: &sql.SkuArgs{
				Name: pulumi.String("HS_Gen5_4"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
