package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/elasticsan/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsan.NewElasticSan(ctx, "elasticSan", &elasticsan.ElasticSanArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("aaaaaaaaaaaaaaaaa"),
			},
			BaseSizeTiB:             pulumi.Float64(26),
			ElasticSanName:          pulumi.String("ti7q-k952-1qB3J_5"),
			ExtendedCapacitySizeTiB: pulumi.Float64(7),
			Location:                pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			ResourceGroupName:       pulumi.String("rgelasticsan"),
			Sku: &elasticsan.SkuArgs{
				Name: pulumi.String("Premium_LRS"),
				Tier: pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"key896": pulumi.String("aaaaaaaaaaaaaaaaaa"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
