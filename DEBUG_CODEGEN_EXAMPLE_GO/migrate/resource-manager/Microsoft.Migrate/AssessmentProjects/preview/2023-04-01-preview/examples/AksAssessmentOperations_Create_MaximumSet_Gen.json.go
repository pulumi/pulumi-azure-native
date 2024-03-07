package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewAksAssessmentOperation(ctx, "aksAssessmentOperation", &migrate.AksAssessmentOperationArgs{
			AssessmentName:    pulumi.String("testaksassessment"),
			ProjectName:       pulumi.String("testproject"),
			ResourceGroupName: pulumi.String("rgaksswagger"),
			Scope: &migrate.AssessmentScopeParametersArgs{
				ServerGroupId: pulumi.String("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentProjects/testproject/groups/testgrp"),
			},
			Settings: &migrate.AKSAssessmentSettingsArgs{
				AzureLocation:      pulumi.String("Unknown"),
				Category:           pulumi.String("All"),
				Consolidation:      pulumi.String("Full"),
				Currency:           pulumi.String("Unknown"),
				DiscountPercentage: pulumi.Float64(15),
				EnvironmentType:    pulumi.String("Unknown"),
				LicensingProgram:   pulumi.String("Default"),
				PerformanceData: &migrate.PerfDataSettingsArgs{
					Percentile:        pulumi.String("Percentile50"),
					PerfDataEndTime:   pulumi.String("2023-11-07T06:51:24.320Z"),
					PerfDataStartTime: pulumi.String("2023-11-07T06:51:24.320Z"),
					TimeRange:         pulumi.String("Day"),
				},
				PricingTier:    pulumi.String("Standard"),
				SavingsOptions: pulumi.String("None"),
				ScalingFactor:  pulumi.Float64(3),
				SizingCriteria: pulumi.String("PerformanceBased"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
