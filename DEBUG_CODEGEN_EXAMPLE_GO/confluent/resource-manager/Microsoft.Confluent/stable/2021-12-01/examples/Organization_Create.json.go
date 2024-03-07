package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/confluent/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := confluent.NewOrganization(ctx, "organization", &confluent.OrganizationArgs{
			Location: pulumi.String("West US"),
			OfferDetail: &confluent.OfferDetailArgs{
				Id:          pulumi.String("string"),
				PlanId:      pulumi.String("string"),
				PlanName:    pulumi.String("string"),
				PublisherId: pulumi.String("string"),
				TermUnit:    pulumi.String("string"),
			},
			OrganizationName:  pulumi.String("myOrganization"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
			UserDetail: &confluent.UserDetailArgs{
				EmailAddress: pulumi.String("contoso@microsoft.com"),
				FirstName:    pulumi.String("string"),
				LastName:     pulumi.String("string"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
