package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewTable(ctx, "table", &operationalinsights.TableArgs{
			ResourceGroupName: pulumi.String("oiautorest6685"),
			RetentionInDays:   pulumi.Int(45),
			Schema: &operationalinsights.SchemaArgs{
				Columns: operationalinsights.ColumnArray{
					&operationalinsights.ColumnArgs{
						Name: pulumi.String("MyNewColumn"),
						Type: pulumi.String("guid"),
					},
				},
				Name: pulumi.String("AzureNetworkFlow"),
			},
			TableName:            pulumi.String("AzureNetworkFlow"),
			TotalRetentionInDays: pulumi.Int(70),
			WorkspaceName:        pulumi.String("oiautorest6685"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
