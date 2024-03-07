package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewSqlPool(ctx, "sqlPool", &synapse.SqlPoolArgs{
			Collation:             pulumi.String(""),
			CreateMode:            pulumi.String(""),
			Location:              pulumi.String("Southeast Asia"),
			MaxSizeBytes:          pulumi.Float64(0),
			RecoverableDatabaseId: pulumi.String(""),
			ResourceGroupName:     pulumi.String("ExampleResourceGroup"),
			Sku: &synapse.SkuArgs{
				Name: pulumi.String(""),
				Tier: pulumi.String(""),
			},
			SourceDatabaseId:   pulumi.String(""),
			SqlPoolName:        pulumi.String("ExampleSqlPool"),
			StorageAccountType: pulumi.String("LRS"),
			Tags:               nil,
			WorkspaceName:      pulumi.String("ExampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
