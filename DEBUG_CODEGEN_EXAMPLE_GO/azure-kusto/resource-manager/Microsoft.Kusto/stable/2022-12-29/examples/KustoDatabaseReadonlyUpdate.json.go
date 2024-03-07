package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewReadWriteDatabase(ctx, "readWriteDatabase", &kusto.ReadWriteDatabaseArgs{
			ClusterName:       pulumi.String("kustoCluster"),
			DatabaseName:      pulumi.String("kustoReadOnlyDatabase"),
			ResourceGroupName: pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
