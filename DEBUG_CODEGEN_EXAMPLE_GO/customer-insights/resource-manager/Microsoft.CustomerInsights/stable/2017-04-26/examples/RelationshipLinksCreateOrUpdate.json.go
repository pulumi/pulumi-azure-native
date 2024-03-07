package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewRelationshipLink(ctx, "relationshipLink", &customerinsights.RelationshipLinkArgs{
			Description: pulumi.StringMap{
				"en-us": pulumi.String("Link Description"),
			},
			DisplayName: pulumi.StringMap{
				"en-us": pulumi.String("Link DisplayName"),
			},
			HubName:         pulumi.String("sdkTestHub"),
			InteractionType: pulumi.String("testInteraction4332"),
			ProfilePropertyReferences: customerinsights.ParticipantProfilePropertyReferenceArray{
				&customerinsights.ParticipantProfilePropertyReferenceArgs{
					InteractionPropertyName: pulumi.String("profile1"),
					ProfilePropertyName:     pulumi.String("ProfileId"),
				},
			},
			RelatedProfilePropertyReferences: customerinsights.ParticipantProfilePropertyReferenceArray{
				&customerinsights.ParticipantProfilePropertyReferenceArgs{
					InteractionPropertyName: pulumi.String("profile1"),
					ProfilePropertyName:     pulumi.String("ProfileId"),
				},
			},
			RelationshipLinkName: pulumi.String("Somelink"),
			RelationshipName:     pulumi.String("testProfile2326994"),
			ResourceGroupName:    pulumi.String("TestHubRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
