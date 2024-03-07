package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewView(ctx, "view", &costmanagement.ViewArgs{
			Accumulated: pulumi.String("true"),
			Chart:       pulumi.String("Table"),
			DataSet: &costmanagement.ReportConfigDatasetArgs{
				Aggregation: costmanagement.ReportConfigAggregationMap{
					"totalCost": &costmanagement.ReportConfigAggregationArgs{
						Function: pulumi.String("Sum"),
						Name:     pulumi.String("PreTaxCost"),
					},
				},
				Granularity: pulumi.String("Daily"),
				Grouping:    costmanagement.ReportConfigGroupingArray{},
				Sorting: costmanagement.ReportConfigSortingArray{
					&costmanagement.ReportConfigSortingArgs{
						Direction: pulumi.String("Ascending"),
						Name:      pulumi.String("UsageDate"),
					},
				},
			},
			DisplayName: pulumi.String("swagger Example"),
			ETag:        pulumi.String("\"1d4ff9fe66f1d10\""),
			Kpis: costmanagement.KpiPropertiesArray{
				&costmanagement.KpiPropertiesArgs{
					Enabled: pulumi.Bool(true),
					Type:    pulumi.String("Forecast"),
				},
				&costmanagement.KpiPropertiesArgs{
					Enabled: pulumi.Bool(true),
					Id:      pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Consumption/budgets/swaggerDemo"),
					Type:    pulumi.String("Budget"),
				},
			},
			Metric: pulumi.String("ActualCost"),
			Pivots: costmanagement.PivotPropertiesArray{
				&costmanagement.PivotPropertiesArgs{
					Name: pulumi.String("ServiceName"),
					Type: pulumi.String("Dimension"),
				},
				&costmanagement.PivotPropertiesArgs{
					Name: pulumi.String("MeterCategory"),
					Type: pulumi.String("Dimension"),
				},
				&costmanagement.PivotPropertiesArgs{
					Name: pulumi.String("swaggerTagKey"),
					Type: pulumi.String("TagKey"),
				},
			},
			Timeframe: pulumi.String("MonthToDate"),
			Type:      pulumi.String("Usage"),
			ViewName:  pulumi.String("swaggerExample"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
