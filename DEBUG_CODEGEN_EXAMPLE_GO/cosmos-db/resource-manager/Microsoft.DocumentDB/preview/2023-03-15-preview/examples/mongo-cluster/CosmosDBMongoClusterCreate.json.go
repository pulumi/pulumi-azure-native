package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoCluster(ctx, "mongoCluster", &documentdb.MongoClusterArgs{
			AdministratorLogin:         pulumi.String("mongoAdmin"),
			AdministratorLoginPassword: pulumi.String("password"),
			Location:                   pulumi.String("westus2"),
			MongoClusterName:           pulumi.String("myMongoCluster"),
			NodeGroupSpecs: documentdb.NodeGroupSpecArray{
				&documentdb.NodeGroupSpecArgs{
					DiskSizeGB: pulumi.Float64(128),
					EnableHa:   pulumi.Bool(true),
					Kind:       pulumi.String("Shard"),
					NodeCount:  pulumi.Int(3),
					Sku:        pulumi.String("M30"),
				},
			},
			ResourceGroupName: pulumi.String("TestResourceGroup"),
			ServerVersion:     pulumi.String("5.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
