package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewPrivateCloud(ctx, "privateCloud", &avs.PrivateCloudArgs{
			Identity: &avs.PrivateCloudIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location: pulumi.String("eastus2"),
			ManagementCluster: &avs.ManagementClusterArgs{
				ClusterSize: pulumi.Int(4),
			},
			NetworkBlock:      pulumi.String("192.168.48.0/22"),
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Sku: &avs.SkuArgs{
				Name: pulumi.String("AV36"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
