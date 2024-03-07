package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/consumption/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := consumption.NewBudget(ctx, "budget", &consumption.BudgetArgs{
			Amount:     pulumi.Float64(100.65),
			BudgetName: pulumi.String("TestBudget"),
			Category:   pulumi.String("Cost"),
			ETag:       pulumi.String("\"1d34d016a593709\""),
			Filter: &consumption.BudgetFilterArgs{
				And: consumption.BudgetFilterPropertiesArray{
					&consumption.BudgetFilterPropertiesArgs{
						Dimensions: &consumption.BudgetComparisonExpressionArgs{
							Name:     pulumi.String("ResourceId"),
							Operator: pulumi.String("In"),
							Values: pulumi.StringArray{
								pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
								pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1"),
							},
						},
					},
					&consumption.BudgetFilterPropertiesArgs{
						Tags: &consumption.BudgetComparisonExpressionArgs{
							Name:     pulumi.String("category"),
							Operator: pulumi.String("In"),
							Values: pulumi.StringArray{
								pulumi.String("Dev"),
								pulumi.String("Prod"),
							},
						},
					},
					&consumption.BudgetFilterPropertiesArgs{
						Tags: &consumption.BudgetComparisonExpressionArgs{
							Name:     pulumi.String("department"),
							Operator: pulumi.String("In"),
							Values: pulumi.StringArray{
								pulumi.String("engineering"),
								pulumi.String("sales"),
							},
						},
					},
				},
			},
			Notifications: consumption.NotificationMap{
				"Actual_GreaterThan_80_Percent": &consumption.NotificationArgs{
					ContactEmails: pulumi.StringArray{
						pulumi.String("johndoe@contoso.com"),
						pulumi.String("janesmith@contoso.com"),
					},
					ContactGroups: pulumi.StringArray{
						pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"),
					},
					ContactRoles: pulumi.StringArray{
						pulumi.String("Contributor"),
						pulumi.String("Reader"),
					},
					Enabled:       pulumi.Bool(true),
					Locale:        pulumi.String("en-us"),
					Operator:      pulumi.String("GreaterThan"),
					Threshold:     pulumi.Float64(80),
					ThresholdType: pulumi.String("Actual"),
				},
			},
			Scope:     pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
			TimeGrain: pulumi.String("Monthly"),
			TimePeriod: &consumption.BudgetTimePeriodArgs{
				EndDate:   pulumi.String("2018-10-31T00:00:00Z"),
				StartDate: pulumi.String("2017-10-01T00:00:00Z"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
