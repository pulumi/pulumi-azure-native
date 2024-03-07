package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewMarkupRule(ctx, "markupRule", &costmanagement.MarkupRuleArgs{
			BillingAccountId: pulumi.String("2af90bea-080c-438c-8977-17cddd5f115a:ef5ce3cf-f5af-4fcb-a5ed-c376e1d6d2b6"),
			BillingProfileId: pulumi.String("cbf78278-f4b8-43d9-8f13-47112da1c63e"),
			CustomerDetails: &costmanagement.CustomerMetadataArgs{
				BillingAccountId: pulumi.String("cff9aa6d-941c-43f2-b6cb-1d2bb34a02b4:780237f3-1aa6-4159-943b-498e0d647dd9"),
				BillingProfileId: pulumi.String("08eeecee-efb2-40d5-817c-0a254d2e042c"),
			},
			Description: pulumi.String("Markup rule for year 2022"),
			EndDate:     pulumi.String("2022-12-31T00:00:00Z"),
			Name:        pulumi.String("markup-2022"),
			Percentage:  pulumi.Float64(5),
			StartDate:   pulumi.String("2022-01-01T00:00:00Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
