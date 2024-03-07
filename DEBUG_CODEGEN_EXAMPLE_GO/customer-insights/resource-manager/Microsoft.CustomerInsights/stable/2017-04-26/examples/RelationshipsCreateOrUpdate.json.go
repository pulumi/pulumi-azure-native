package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewRelationship(ctx, "relationship", &customerinsights.RelationshipArgs{
			Cardinality: customerinsights.CardinalityTypesOneToOne,
			Description: pulumi.StringMap{
				"en-us": pulumi.String("Relationship Description"),
			},
			DisplayName: pulumi.StringMap{
				"en-us": pulumi.String("Relationship DisplayName"),
			},
			Fields:             customerinsights.PropertyDefinitionArray{},
			HubName:            pulumi.String("sdkTestHub"),
			ProfileType:        pulumi.String("testProfile2326994"),
			RelatedProfileType: pulumi.String("testProfile2326994"),
			RelationshipName:   pulumi.String("SomeRelationship"),
			ResourceGroupName:  pulumi.String("TestHubRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
