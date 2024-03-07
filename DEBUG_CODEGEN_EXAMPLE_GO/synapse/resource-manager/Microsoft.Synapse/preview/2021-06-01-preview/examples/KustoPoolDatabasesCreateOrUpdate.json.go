package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewReadWriteDatabase(ctx, "readWriteDatabase", &synapse.ReadWriteDatabaseArgs{
			DatabaseName:      pulumi.String("KustoDatabase8"),
			Kind:              pulumi.String("ReadWrite"),
			KustoPoolName:     pulumi.String("kustoclusterrptest4"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("kustorptest"),
			SoftDeletePeriod:  pulumi.String("P1D"),
			WorkspaceName:     pulumi.String("synapseWorkspaceName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
