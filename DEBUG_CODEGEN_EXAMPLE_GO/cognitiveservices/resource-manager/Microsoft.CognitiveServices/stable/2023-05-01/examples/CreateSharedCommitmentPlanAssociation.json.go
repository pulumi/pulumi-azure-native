package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewCommitmentPlanAssociation(ctx, "commitmentPlanAssociation", &cognitiveservices.CommitmentPlanAssociationArgs{
			AccountId:                     pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
			CommitmentPlanAssociationName: pulumi.String("commitmentPlanAssociationName"),
			CommitmentPlanName:            pulumi.String("commitmentPlanName"),
			ResourceGroupName:             pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
