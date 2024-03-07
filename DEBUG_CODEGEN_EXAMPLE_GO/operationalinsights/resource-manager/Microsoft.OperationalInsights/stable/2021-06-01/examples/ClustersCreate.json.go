package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewCluster(ctx, "cluster", &operationalinsights.ClusterArgs{
			ClusterName:       pulumi.String("oiautorest6685"),
			Location:          pulumi.String("australiasoutheast"),
			ResourceGroupName: pulumi.String("oiautorest6685"),
			Sku: &operationalinsights.ClusterSkuArgs{
				Capacity: pulumi.Float64(1000),
				Name:     pulumi.String("CapacityReservation"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("val1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
