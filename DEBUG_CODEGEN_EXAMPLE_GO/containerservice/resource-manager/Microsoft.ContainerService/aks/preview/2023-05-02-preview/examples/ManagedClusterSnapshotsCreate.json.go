package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewManagedClusterSnapshot(ctx, "managedClusterSnapshot", &containerservice.ManagedClusterSnapshotArgs{
			CreationData: &containerservice.CreationDataArgs{
				SourceResourceId: pulumi.String("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("snapshot1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("val1"),
				"key2": pulumi.String("val2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
