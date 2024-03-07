package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewExport(ctx, "export", &costmanagement.ExportArgs{
			Definition: &costmanagement.ExportDefinitionArgs{
				DataSet: &costmanagement.ExportDatasetArgs{
					Configuration: &costmanagement.ExportDatasetConfigurationArgs{
						Columns: pulumi.StringArray{
							pulumi.String("Date"),
							pulumi.String("MeterId"),
							pulumi.String("ResourceId"),
							pulumi.String("ResourceLocation"),
							pulumi.String("Quantity"),
						},
					},
					Granularity: pulumi.String("Daily"),
				},
				Timeframe: pulumi.String("MonthToDate"),
				Type:      pulumi.String("ActualCost"),
			},
			DeliveryInfo: &costmanagement.ExportDeliveryInfoArgs{
				Destination: &costmanagement.ExportDeliveryDestinationArgs{
					Container:      pulumi.String("exports"),
					ResourceId:     pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182"),
					RootFolderPath: pulumi.String("ad-hoc"),
				},
			},
			ExportName: pulumi.String("TestExport"),
			Format:     pulumi.String("Csv"),
			Schedule: &costmanagement.ExportScheduleArgs{
				Recurrence: pulumi.String("Weekly"),
				RecurrencePeriod: &costmanagement.ExportRecurrencePeriodArgs{
					From: pulumi.String("2020-06-01T00:00:00Z"),
					To:   pulumi.String("2020-10-31T00:00:00Z"),
				},
				Status: pulumi.String("Active"),
			},
			Scope: pulumi.String("providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
