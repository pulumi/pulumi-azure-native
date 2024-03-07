package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewKpi(ctx, "kpi", &customerinsights.KpiArgs{
			Aliases: customerinsights.KpiAliasArray{
				&customerinsights.KpiAliasArgs{
					AliasName:  pulumi.String("alias"),
					Expression: pulumi.String("Id+4"),
				},
			},
			CalculationWindow: customerinsights.CalculationWindowTypesDay,
			Description: pulumi.StringMap{
				"en-us": pulumi.String("Kpi Description"),
			},
			DisplayName: pulumi.StringMap{
				"en-us": pulumi.String("Kpi DisplayName"),
			},
			EntityType:     customerinsights.EntityTypesProfile,
			EntityTypeName: pulumi.String("testProfile2327128"),
			Expression:     pulumi.String("SavingAccountBalance"),
			Function:       customerinsights.KpiFunctionsSum,
			GroupBy: pulumi.StringArray{
				pulumi.String("SavingAccountBalance"),
			},
			HubName:           pulumi.String("sdkTestHub"),
			KpiName:           pulumi.String("kpiTest45453647"),
			ResourceGroupName: pulumi.String("TestHubRG"),
			ThresHolds: &customerinsights.KpiThresholdsArgs{
				IncreasingKpi: pulumi.Bool(true),
				LowerLimit:    pulumi.Float64(5),
				UpperLimit:    pulumi.Float64(50),
			},
			Unit: pulumi.String("unit"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
