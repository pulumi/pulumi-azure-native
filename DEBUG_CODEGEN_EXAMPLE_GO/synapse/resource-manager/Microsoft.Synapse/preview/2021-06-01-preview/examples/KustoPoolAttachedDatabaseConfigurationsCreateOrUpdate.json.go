package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewKustoPoolAttachedDatabaseConfiguration(ctx, "kustoPoolAttachedDatabaseConfiguration", &synapse.KustoPoolAttachedDatabaseConfigurationArgs{
			AttachedDatabaseConfigurationName: pulumi.String("attachedDatabaseConfigurations1"),
			DatabaseName:                      pulumi.String("kustodatabase"),
			DefaultPrincipalsModificationKind: pulumi.String("Union"),
			KustoPoolName:                     pulumi.String("kustoclusterrptest4"),
			KustoPoolResourceId:               pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/kustoclusterrptest4"),
			Location:                          pulumi.String("westus"),
			ResourceGroupName:                 pulumi.String("kustorptest"),
			TableLevelSharingProperties: &synapse.TableLevelSharingPropertiesArgs{
				ExternalTablesToExclude: pulumi.StringArray{
					pulumi.String("ExternalTable2"),
				},
				ExternalTablesToInclude: pulumi.StringArray{
					pulumi.String("ExternalTable1"),
				},
				MaterializedViewsToExclude: pulumi.StringArray{
					pulumi.String("MaterializedViewTable2"),
				},
				MaterializedViewsToInclude: pulumi.StringArray{
					pulumi.String("MaterializedViewTable1"),
				},
				TablesToExclude: pulumi.StringArray{
					pulumi.String("Table2"),
				},
				TablesToInclude: pulumi.StringArray{
					pulumi.String("Table1"),
				},
			},
			WorkspaceName: pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
