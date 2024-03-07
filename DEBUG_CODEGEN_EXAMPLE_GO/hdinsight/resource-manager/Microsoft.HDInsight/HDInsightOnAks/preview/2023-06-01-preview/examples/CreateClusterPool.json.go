package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hdinsight/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hdinsight.NewClusterPool(ctx, "clusterPool", &hdinsight.ClusterPoolArgs{
			ClusterPoolName: pulumi.String("clusterpool1"),
			ClusterPoolProfile: &hdinsight.ClusterPoolResourcePropertiesClusterPoolProfileArgs{
				ClusterPoolVersion: pulumi.String("1.2"),
			},
			ComputeProfile: &hdinsight.ClusterPoolResourcePropertiesComputeProfileArgs{
				VmSize: pulumi.String("Standard_D3_v2"),
			},
			Location:          pulumi.String("West US 2"),
			ResourceGroupName: pulumi.String("hiloResourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
