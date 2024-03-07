package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewCommitmentPlan(ctx, "commitmentPlan", &cognitiveservices.CommitmentPlanArgs{
			AccountName:        pulumi.String("accountName"),
			CommitmentPlanName: pulumi.String("commitmentPlanName"),
			Properties: &cognitiveservices.CommitmentPlanPropertiesArgs{
				AutoRenew: pulumi.Bool(true),
				Current: &cognitiveservices.CommitmentPeriodArgs{
					Tier: pulumi.String("T1"),
				},
				HostingModel: pulumi.String("Web"),
				PlanType:     pulumi.String("Speech2Text"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
