package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerTopic(ctx, "partnerTopic", &eventgrid.PartnerTopicArgs{
			ExpirationTimeIfNotActivatedUtc: pulumi.String("2022-03-23T23:06:13.109Z"),
			Location:                        pulumi.String("westus2"),
			MessageForActivation:            pulumi.String("Example message for activation"),
			PartnerRegistrationImmutableId:  pulumi.String("6f541064-031d-4cc8-9ec3-a3b4fc0f7185"),
			PartnerTopicFriendlyDescription: pulumi.String("Example description"),
			PartnerTopicName:                pulumi.String("examplePartnerTopicName1"),
			ResourceGroupName:               pulumi.String("examplerg"),
			Source:                          pulumi.String("ContosoCorp.Accounts.User1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
