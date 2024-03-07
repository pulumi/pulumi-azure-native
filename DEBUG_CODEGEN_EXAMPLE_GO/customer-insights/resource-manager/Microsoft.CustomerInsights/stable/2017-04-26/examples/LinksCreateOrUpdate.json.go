package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewLink(ctx, "link", &customerinsights.LinkArgs{
			Description: pulumi.StringMap{
				"en-us": pulumi.String("Link Description"),
			},
			DisplayName: pulumi.StringMap{
				"en-us": pulumi.String("Link DisplayName"),
			},
			HubName:  pulumi.String("sdkTestHub"),
			LinkName: pulumi.String("linkTest4806"),
			Mappings: customerinsights.TypePropertiesMappingArray{
				&customerinsights.TypePropertiesMappingArgs{
					LinkType:           customerinsights.LinkTypesUpdateAlways,
					SourcePropertyName: pulumi.String("testInteraction1949"),
					TargetPropertyName: pulumi.String("testProfile1446"),
				},
			},
			ParticipantPropertyReferences: customerinsights.ParticipantPropertyReferenceArray{
				&customerinsights.ParticipantPropertyReferenceArgs{
					SourcePropertyName: pulumi.String("testInteraction1949"),
					TargetPropertyName: pulumi.String("ProfileId"),
				},
			},
			ResourceGroupName:    pulumi.String("TestHubRG"),
			SourceEntityType:     customerinsights.EntityTypeInteraction,
			SourceEntityTypeName: pulumi.String("testInteraction1949"),
			TargetEntityType:     customerinsights.EntityTypeProfile,
			TargetEntityTypeName: pulumi.String("testProfile1446"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
