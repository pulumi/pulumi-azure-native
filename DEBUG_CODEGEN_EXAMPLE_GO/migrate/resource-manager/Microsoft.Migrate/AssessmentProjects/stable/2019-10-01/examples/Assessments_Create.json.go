package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewAssessment(ctx, "assessment", &migrate.AssessmentArgs{
			AssessmentName: pulumi.String("assessment_5_14_2019_16_48_47"),
			ETag:           pulumi.String("\"1e000c2c-0000-0d00-0000-5cdaa4190000\""),
			GroupName:      pulumi.String("Group2"),
			ProjectName:    pulumi.String("abgoyalWEselfhostb72bproject"),
			Properties: migrate.AssessmentPropertiesResponse{
				AzureDiskType:          pulumi.String("StandardOrPremium"),
				AzureHybridUseBenefit:  pulumi.String("Yes"),
				AzureLocation:          pulumi.String("NorthEurope"),
				AzureOfferCode:         pulumi.String("MSAZR0003P"),
				AzurePricingTier:       pulumi.String("Standard"),
				AzureStorageRedundancy: pulumi.String("LocallyRedundant"),
				AzureVmFamilies: pulumi.StringArray{
					pulumi.String("Dv2_series"),
					pulumi.String("F_series"),
					pulumi.String("Dv3_series"),
					pulumi.String("DS_series"),
					pulumi.String("DSv2_series"),
					pulumi.String("Fs_series"),
					pulumi.String("Dsv3_series"),
					pulumi.String("Ev3_series"),
					pulumi.String("Esv3_series"),
					pulumi.String("D_series"),
					pulumi.String("M_series"),
					pulumi.String("Fsv2_series"),
					pulumi.String("H_series"),
				},
				Currency:           pulumi.String("USD"),
				DiscountPercentage: pulumi.Float64(100),
				Percentile:         pulumi.String("Percentile95"),
				ReservedInstance:   pulumi.String("RI3Year"),
				ScalingFactor:      pulumi.Float64(1),
				SizingCriterion:    pulumi.String("PerformanceBased"),
				Stage:              pulumi.String("InProgress"),
				TimeRange:          pulumi.String("Day"),
				VmUptime: &migrate.VmUptimeArgs{
					DaysPerMonth: pulumi.Float64(31),
					HoursPerDay:  pulumi.Float64(24),
				},
			},
			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
