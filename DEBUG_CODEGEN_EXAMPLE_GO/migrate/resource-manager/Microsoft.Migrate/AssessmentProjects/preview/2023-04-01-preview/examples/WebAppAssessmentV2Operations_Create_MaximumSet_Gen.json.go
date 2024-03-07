package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewWebAppAssessmentV2Operation(ctx, "webAppAssessmentV2Operation", &migrate.WebAppAssessmentV2OperationArgs{
			AppSvcContainerSettings: &migrate.AppSvcContainerSettingsArgs{
				IsolationRequired: pulumi.Bool(true),
			},
			AppSvcNativeSettings: &migrate.AppSvcNativeSettingsArgs{
				IsolationRequired: pulumi.Bool(true),
			},
			AssessmentName:               pulumi.String("anraghun-selfhost-v2"),
			AssessmentType:               pulumi.String("WebAppAssessment"),
			AzureLocation:                pulumi.String("UkWest"),
			AzureOfferCode:               pulumi.String("MSAZR0003P"),
			AzureSecurityOfferingType:    pulumi.String("NO"),
			ConfidenceRatingInPercentage: pulumi.Float64(13),
			Currency:                     pulumi.String("USD"),
			DiscountPercentage:           pulumi.Float64(13),
			DiscoveredEntityLightSummary: &migrate.DiscoveredEntityLightSummaryArgs{
				NumberOfMachines: pulumi.Int(27),
				NumberOfServers:  pulumi.Int(5),
				NumberOfWebApps:  pulumi.Int(23),
			},
			EaSubscriptionId: pulumi.String(""),
			EntityUptime: &migrate.EntityUptimeArgs{
				DaysPerMonth: pulumi.Int(18),
				HoursPerDay:  pulumi.Int(13),
			},
			EnvironmentType:   pulumi.String("Production"),
			GroupName:         pulumi.String("anraghun-selfhost-v2"),
			GroupType:         pulumi.String("Default"),
			Percentile:        pulumi.String("Percentile50"),
			PerfDataEndTime:   pulumi.String("2023-11-03T05:42:45.496Z"),
			PerfDataStartTime: pulumi.String("2023-11-03T05:42:45.496Z"),
			ProjectName:       pulumi.String("sumukk-ccy-bcs4557project"),
			ReservedInstance:  pulumi.String("None"),
			ResourceGroupName: pulumi.String("rgopenapi"),
			ScalingFactor:     pulumi.Float64(17),
			SizingCriterion:   pulumi.String("PerformanceBased"),
			TimeRange:         pulumi.String("Day"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
