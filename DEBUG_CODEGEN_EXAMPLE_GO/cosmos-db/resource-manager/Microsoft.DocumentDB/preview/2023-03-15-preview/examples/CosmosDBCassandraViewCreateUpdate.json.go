package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewCassandraResourceCassandraView(ctx, "cassandraResourceCassandraView", &documentdb.CassandraResourceCassandraViewArgs{
			AccountName:  pulumi.String("ddb1"),
			KeyspaceName: pulumi.String("keyspacename"),
			Options:      nil,
			Resource: &documentdb.CassandraViewResourceArgs{
				Id:             pulumi.String("viewname"),
				ViewDefinition: pulumi.String("SELECT columna, columnb, columnc FROM keyspacename.srctablename WHERE columna IS NOT NULL AND columnc IS NOT NULL PRIMARY (columnc, columna)"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags:              nil,
			ViewName:          pulumi.String("viewname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
