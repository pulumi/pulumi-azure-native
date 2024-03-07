package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewSynapseWorkspaceSqlPoolTableDataSet(ctx, "synapseWorkspaceSqlPoolTableDataSet", &datashare.SynapseWorkspaceSqlPoolTableDataSetArgs{
			AccountName:                            pulumi.String("sourceAccount"),
			DataSetName:                            pulumi.String("dataset1"),
			Kind:                                   pulumi.String("SynapseWorkspaceSqlPoolTable"),
			ResourceGroupName:                      pulumi.String("SampleResourceGroup"),
			ShareName:                              pulumi.String("share1"),
			SynapseWorkspaceSqlPoolTableResourceId: pulumi.String("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
