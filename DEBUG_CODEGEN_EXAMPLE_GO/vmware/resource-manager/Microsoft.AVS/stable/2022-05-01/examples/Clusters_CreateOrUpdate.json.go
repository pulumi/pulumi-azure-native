package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewCluster(ctx, "cluster", &avs.ClusterArgs{
			ClusterName:       pulumi.String("cluster1"),
			ClusterSize:       pulumi.Int(3),
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Sku: &avs.SkuArgs{
				Name: pulumi.String("AV20"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
