package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewPrediction(ctx, "prediction", &customerinsights.PredictionArgs{
			AutoAnalyze: pulumi.Bool(true),
			Description: pulumi.StringMap{
				"en-us": pulumi.String("sdktest"),
			},
			DisplayName: pulumi.StringMap{
				"en-us": pulumi.String("sdktest"),
			},
			Grades:                   customerinsights.PredictionGradesArray{},
			HubName:                  pulumi.String("sdkTestHub"),
			InvolvedInteractionTypes: pulumi.StringArray{},
			InvolvedKpiTypes:         pulumi.StringArray{},
			InvolvedRelationships:    pulumi.StringArray{},
			Mappings: &customerinsights.PredictionMappingsArgs{
				Grade:  pulumi.String("sdktest_Grade"),
				Reason: pulumi.String("sdktest_Reason"),
				Score:  pulumi.String("sdktest_Score"),
			},
			NegativeOutcomeExpression: pulumi.String("Customers.FirstName = 'Mike'"),
			PositiveOutcomeExpression: pulumi.String("Customers.FirstName = 'David'"),
			PredictionName:            pulumi.String("sdktest"),
			PrimaryProfileType:        pulumi.String("Customers"),
			ResourceGroupName:         pulumi.String("TestHubRG"),
			ScopeExpression:           pulumi.String("*"),
			ScoreLabel:                pulumi.String("score label"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
