package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerRegistration(ctx, "partnerRegistration", &eventgrid.PartnerRegistrationArgs{
			Location:                pulumi.String("global"),
			PartnerRegistrationName: pulumi.String("examplePartnerRegistrationName1"),
			ResourceGroupName:       pulumi.String("examplerg"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("Value2"),
				"key3": pulumi.String("Value3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
