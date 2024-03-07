package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewCassandraResourceCassandraKeyspace(ctx, "cassandraResourceCassandraKeyspace", &documentdb.CassandraResourceCassandraKeyspaceArgs{
			AccountName:  pulumi.String("ddb1"),
			KeyspaceName: pulumi.String("keyspaceName"),
			Location:     pulumi.String("West US"),
			Options:      nil,
			Resource: &documentdb.CassandraKeyspaceResourceArgs{
				Id: pulumi.String("keyspaceName"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
