package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerstorage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerstorage.NewPool(ctx, "pool", &containerstorage.PoolArgs{
			Assignments: []containerstorage.AssignmentArgs{
				{
					Id: pulumi.String("/subscriptions/bb4d87a2-4273-466c-a6ba-61d818061b3a/resourceGroups/test-rg/providers/Microsoft.ContainerService/managedClusters/containerstoragetest"),
				},
			},
			Location: pulumi.String("eastus"),
			PoolName: pulumi.String("test-pool"),
			PoolType: containerstorage.PoolTypeResponse{
				EphemeralDisk: &containerstorage.EphemeralDiskArgs{
					Replicas: pulumi.Float64(3),
				},
			},
			ReclaimPolicy:     pulumi.String("Delete"),
			ResourceGroupName: pulumi.String("test-rg"),
			Resources: containerstorage.ResourcesResponse{
				Requests: &containerstorage.RequestsArgs{
					Storage: pulumi.Float64(15578),
				},
			},
			Tags: pulumi.StringMap{
				"key1888": pulumi.String("value1888"),
			},
			Zones: pulumi.StringArray{
				pulumi.String("1"),
				pulumi.String("2"),
				pulumi.String("3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
