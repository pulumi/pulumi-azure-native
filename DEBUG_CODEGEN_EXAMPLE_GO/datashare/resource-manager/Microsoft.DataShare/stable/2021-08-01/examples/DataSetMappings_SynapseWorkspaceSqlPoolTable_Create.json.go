package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewSynapseWorkspaceSqlPoolTableDataSetMapping(ctx, "synapseWorkspaceSqlPoolTableDataSetMapping", &datashare.SynapseWorkspaceSqlPoolTableDataSetMappingArgs{
			AccountName:                            pulumi.String("consumerAccount"),
			DataSetId:                              pulumi.String("3dc64e49-1fc3-4186-b3dc-d388c4d3076a"),
			DataSetMappingName:                     pulumi.String("datasetMappingName1"),
			Kind:                                   pulumi.String("SynapseWorkspaceSqlPoolTable"),
			ResourceGroupName:                      pulumi.String("SampleResourceGroup"),
			ShareSubscriptionName:                  pulumi.String("ShareSubscription1"),
			SynapseWorkspaceSqlPoolTableResourceId: pulumi.String("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
