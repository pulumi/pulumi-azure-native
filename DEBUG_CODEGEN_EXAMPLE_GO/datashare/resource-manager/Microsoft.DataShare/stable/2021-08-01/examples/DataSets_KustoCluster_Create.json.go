package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewSynapseWorkspaceSqlPoolTableDataSet(ctx, "synapseWorkspaceSqlPoolTableDataSet", &datashare.SynapseWorkspaceSqlPoolTableDataSetArgs{
			AccountName:       pulumi.String("Account1"),
			DataSetName:       pulumi.String("Dataset1"),
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
			ShareName:         pulumi.String("Share1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
