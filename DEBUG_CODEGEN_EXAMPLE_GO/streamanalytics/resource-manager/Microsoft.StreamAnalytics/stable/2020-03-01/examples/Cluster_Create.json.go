package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewCluster(ctx, "cluster", &streamanalytics.ClusterArgs{
			ClusterName:       pulumi.String("An Example Cluster"),
			Location:          pulumi.String("North US"),
			ResourceGroupName: pulumi.String("sjrg"),
			Sku: &streamanalytics.ClusterSkuArgs{
				Capacity: pulumi.Int(48),
				Name:     pulumi.String("Default"),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
