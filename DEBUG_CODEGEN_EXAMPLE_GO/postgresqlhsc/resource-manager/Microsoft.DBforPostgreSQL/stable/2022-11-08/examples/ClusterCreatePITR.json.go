package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewCluster(ctx, "cluster", &dbforpostgresql.ClusterArgs{
			ClusterName:       pulumi.String("testcluster"),
			Location:          pulumi.String("westus"),
			PointInTimeUTC:    pulumi.String("2017-12-14T00:00:37.467Z"),
			ResourceGroupName: pulumi.String("TestGroup"),
			SourceLocation:    pulumi.String("westus"),
			SourceResourceId:  pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/source-cluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
