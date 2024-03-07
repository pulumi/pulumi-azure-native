package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewAttachedDatabaseConfiguration(ctx, "attachedDatabaseConfiguration", &kusto.AttachedDatabaseConfigurationArgs{
			AttachedDatabaseConfigurationName: pulumi.String("attachedDatabaseConfigurationsTest"),
			ClusterName:                       pulumi.String("kustoCluster2"),
			ClusterResourceId:                 pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
			DatabaseName:                      pulumi.String("kustodatabase"),
			DatabaseNameOverride:              pulumi.String("overridekustodatabase"),
			DefaultPrincipalsModificationKind: pulumi.String("Union"),
			Location:                          pulumi.String("westus"),
			ResourceGroupName:                 pulumi.String("kustorptest"),
			TableLevelSharingProperties: &kusto.TableLevelSharingPropertiesArgs{
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
		})
		if err != nil {
			return err
		}
		return nil
	})
}
