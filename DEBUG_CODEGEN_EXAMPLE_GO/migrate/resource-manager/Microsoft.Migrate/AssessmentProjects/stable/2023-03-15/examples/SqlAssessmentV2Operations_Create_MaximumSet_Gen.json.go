package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewSqlAssessmentV2Operation(ctx, "sqlAssessmentV2Operation", &migrate.SqlAssessmentV2OperationArgs{
			AssessmentName:        pulumi.String("test_swagger_1"),
			AsyncCommitModeIntent: pulumi.String("DisasterRecovery"),
			AzureLocation:         pulumi.String("SoutheastAsia"),
			AzureOfferCode:        pulumi.String("MSAZR0003P"),
			AzureOfferCodeForVm:   pulumi.String("MSAZR0003P"),
			AzureSqlDatabaseSettings: &migrate.SqlDbSettingsArgs{
				AzureSqlComputeTier:   pulumi.String("Automatic"),
				AzureSqlDataBaseType:  pulumi.String("SingleDatabase"),
				AzureSqlPurchaseModel: pulumi.String("VCore"),
				AzureSqlServiceTier:   pulumi.String("Automatic"),
			},
			AzureSqlManagedInstanceSettings: &migrate.SqlMiSettingsArgs{
				AzureSqlInstanceType: pulumi.String("SingleInstance"),
				AzureSqlServiceTier:  pulumi.String("Automatic"),
			},
			AzureSqlVmSettings: &migrate.SqlVmSettingsArgs{
				InstanceSeries: pulumi.StringArray{
					pulumi.String("Eadsv5_series"),
				},
			},
			Currency:                 pulumi.String("USD"),
			DisasterRecoveryLocation: pulumi.String("EastAsia"),
			DiscountPercentage:       pulumi.Float64(0),
			EnableHadrAssessment:     pulumi.Bool(true),
			EntityUptime: &migrate.EntityUptimeArgs{
				DaysPerMonth: pulumi.Int(30),
				HoursPerDay:  pulumi.Int(24),
			},
			EnvironmentType:       pulumi.String("Production"),
			GroupName:             pulumi.String("test_fci_hadr"),
			MultiSubnetIntent:     pulumi.String("DisasterRecovery"),
			OptimizationLogic:     pulumi.String("MinimizeCost"),
			OsLicense:             pulumi.String("Unknown"),
			Percentile:            pulumi.String("Percentile95"),
			ProjectName:           pulumi.String("fci-test6904project"),
			ReservedInstance:      pulumi.String("None"),
			ReservedInstanceForVm: pulumi.String("None"),
			ResourceGroupName:     pulumi.String("rgmigrate"),
			ScalingFactor:         pulumi.Float64(1),
			SizingCriterion:       pulumi.String("PerformanceBased"),
			SqlServerLicense:      pulumi.String("Unknown"),
			TimeRange:             pulumi.String("Day"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
