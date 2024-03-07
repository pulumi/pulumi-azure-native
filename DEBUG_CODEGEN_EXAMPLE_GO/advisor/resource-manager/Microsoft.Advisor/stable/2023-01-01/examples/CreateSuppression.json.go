package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/advisor/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := advisor.NewSuppression(ctx, "suppression", &advisor.SuppressionArgs{
			Name:             pulumi.String("suppressionName1"),
			RecommendationId: pulumi.String("recommendationId"),
			ResourceUri:      pulumi.String("resourceUri"),
			Ttl:              pulumi.String("07:00:00:00"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
