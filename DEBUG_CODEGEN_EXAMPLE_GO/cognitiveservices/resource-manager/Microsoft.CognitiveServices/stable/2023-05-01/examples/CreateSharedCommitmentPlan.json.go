package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewSharedCommitmentPlan(ctx, "sharedCommitmentPlan", &cognitiveservices.SharedCommitmentPlanArgs{
			CommitmentPlanName: pulumi.String("commitmentPlanName"),
			Kind:               pulumi.String("SpeechServices"),
			Location:           pulumi.String("West US"),
			Properties: cognitiveservices.CommitmentPlanPropertiesResponse{
				AutoRenew: pulumi.Bool(true),
				Current: &cognitiveservices.CommitmentPeriodArgs{
					Tier: pulumi.String("T1"),
				},
				HostingModel: pulumi.String("Web"),
				PlanType:     pulumi.String("STT"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Sku: &cognitiveservices.SkuArgs{
				Name: pulumi.String("S0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
