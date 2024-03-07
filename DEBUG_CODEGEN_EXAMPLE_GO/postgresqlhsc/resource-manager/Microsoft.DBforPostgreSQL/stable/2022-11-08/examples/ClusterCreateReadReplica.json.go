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
			ResourceGroupName: pulumi.String("TestGroup"),
			SourceLocation:    pulumi.String("westus"),
			SourceResourceId:  pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/sourcecluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
