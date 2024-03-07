package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewShareSubscription(ctx, "shareSubscription", &datashare.ShareSubscriptionArgs{
			AccountName:           pulumi.String("Account1"),
			ExpirationDate:        pulumi.String("2020-08-26T22:33:24.5785265Z"),
			InvitationId:          pulumi.String("12345678-1234-1234-12345678abd"),
			ResourceGroupName:     pulumi.String("SampleResourceGroup"),
			ShareSubscriptionName: pulumi.String("ShareSubscription1"),
			SourceShareLocation:   pulumi.String("eastus2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
