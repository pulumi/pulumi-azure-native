package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewReadWriteDatabase(ctx, "readWriteDatabase", &kusto.ReadWriteDatabaseArgs{
			CallerRole:        pulumi.String("Admin"),
			ClusterName:       pulumi.String("kustoCluster"),
			DatabaseName:      pulumi.String("KustoDatabase8"),
			Kind:              pulumi.String("ReadWrite"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("kustorptest"),
			SoftDeletePeriod:  pulumi.String("P1D"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
