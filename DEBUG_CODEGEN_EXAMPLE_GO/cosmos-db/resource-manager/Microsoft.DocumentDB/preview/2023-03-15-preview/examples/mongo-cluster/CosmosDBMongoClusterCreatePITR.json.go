package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoCluster(ctx, "mongoCluster", &documentdb.MongoClusterArgs{
			CreateMode:        pulumi.String("PointInTimeRestore"),
			Location:          pulumi.String("westus2"),
			MongoClusterName:  pulumi.String("myMongoCluster"),
			ResourceGroupName: pulumi.String("TestResourceGroup"),
			RestoreParameters: &documentdb.MongoClusterRestoreParametersArgs{
				PointInTimeUTC:   pulumi.String("2023-01-13T20:07:35Z"),
				SourceResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
