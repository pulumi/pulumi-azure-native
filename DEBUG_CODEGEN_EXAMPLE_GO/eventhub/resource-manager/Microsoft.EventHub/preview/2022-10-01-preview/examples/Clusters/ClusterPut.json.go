package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewCluster(ctx, "cluster", &eventhub.ClusterArgs{
			ClusterName:       pulumi.String("testCluster"),
			Location:          pulumi.String("South Central US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &eventhub.ClusterSkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Dedicated"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
