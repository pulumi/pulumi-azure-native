package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewAssessmentsOperation(ctx, "assessmentsOperation", &migrate.AssessmentsOperationArgs{
			AssessmentName: pulumi.String("asm1"),
			AzureDiskTypes: pulumi.StringArray{
				pulumi.String("Premium"),
				pulumi.String("StandardSSD"),
			},
			AzureHybridUseBenefit:  pulumi.String("Unknown"),
			AzureLocation:          pulumi.String("njxbwdtsxzhichsnk"),
			AzureOfferCode:         pulumi.String("Unknown"),
			AzurePricingTier:       pulumi.String("Standard"),
			AzureStorageRedundancy: pulumi.String("Unknown"),
			AzureVmFamilies: pulumi.StringArray{
				pulumi.String("D_series"),
				pulumi.String("Lsv2_series"),
				pulumi.String("M_series"),
				pulumi.String("Mdsv2_series"),
				pulumi.String("Msv2_series"),
				pulumi.String("Mv2_series"),
			},
			Currency:                   pulumi.String("Unknown"),
			DiscountPercentage:         pulumi.Float64(6),
			EaSubscriptionId:           pulumi.String("kwsu"),
			GroupName:                  pulumi.String("kuchatur-test"),
			LinuxAzureHybridUseBenefit: pulumi.String("Unknown"),
			Percentile:                 pulumi.String("Percentile50"),
			PerfDataEndTime:            pulumi.String("2023-09-26T09:36:48.491Z"),
			PerfDataStartTime:          pulumi.String("2023-09-26T09:36:48.491Z"),
			ProjectName:                pulumi.String("app18700project"),
			ProvisioningState:          pulumi.String("Succeeded"),
			ReservedInstance:           pulumi.String("None"),
			ResourceGroupName:          pulumi.String("ayagrawrg"),
			ScalingFactor:              pulumi.Float64(24),
			SizingCriterion:            pulumi.String("PerformanceBased"),
			TimeRange:                  pulumi.String("Day"),
			VmUptime: &migrate.VmUptimeArgs{
				DaysPerMonth: pulumi.Float64(13),
				HoursPerDay:  pulumi.Float64(26),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
