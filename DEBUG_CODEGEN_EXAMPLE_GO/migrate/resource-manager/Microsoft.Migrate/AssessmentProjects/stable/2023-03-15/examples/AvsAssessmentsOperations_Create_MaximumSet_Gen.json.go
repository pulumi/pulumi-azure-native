package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewAvsAssessmentsOperation(ctx, "avsAssessmentsOperation", &migrate.AvsAssessmentsOperationArgs{
			AssessmentName:                 pulumi.String("asm2"),
			AzureLocation:                  pulumi.String("EastUs"),
			AzureOfferCode:                 pulumi.String("MSAZR0003P"),
			Currency:                       pulumi.String("USD"),
			DedupeCompression:              pulumi.Float64(1.5),
			DiscountPercentage:             pulumi.Float64(0),
			FailuresToTolerateAndRaidLevel: pulumi.String("Ftt1Raid1"),
			GroupName:                      pulumi.String("kuchatur-test"),
			IsStretchClusterEnabled:        pulumi.Bool(true),
			MemOvercommit:                  pulumi.Float64(1),
			NodeType:                       pulumi.String("AV36"),
			Percentile:                     pulumi.String("Percentile95"),
			PerfDataEndTime:                pulumi.String("2023-09-26T13:35:56.5671462Z"),
			PerfDataStartTime:              pulumi.String("2023-09-25T13:35:56.5671462Z"),
			ProjectName:                    pulumi.String("app18700project"),
			ProvisioningState:              pulumi.String("Succeeded"),
			ReservedInstance:               pulumi.String("RI3Year"),
			ResourceGroupName:              pulumi.String("ayagrawrg"),
			ScalingFactor:                  pulumi.Float64(1),
			SizingCriterion:                pulumi.String("AsOnPremises"),
			TimeRange:                      pulumi.String("Day"),
			VcpuOversubscription:           pulumi.Float64(4),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
