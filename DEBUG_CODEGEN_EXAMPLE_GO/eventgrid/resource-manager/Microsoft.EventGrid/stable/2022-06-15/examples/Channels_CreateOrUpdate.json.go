package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewChannel(ctx, "channel", &eventgrid.ChannelArgs{
			ChannelName:                     pulumi.String("exampleChannelName1"),
			ChannelType:                     pulumi.String("PartnerTopic"),
			ExpirationTimeIfNotActivatedUtc: pulumi.String("2021-10-21T22:50:25.410433Z"),
			MessageForActivation:            pulumi.String("Example message to approver"),
			PartnerNamespaceName:            pulumi.String("examplePartnerNamespaceName1"),
			PartnerTopicInfo: &eventgrid.PartnerTopicInfoArgs{
				AzureSubscriptionId: pulumi.String("5b4b650e-28b9-4790-b3ab-ddbd88d727c4"),
				Name:                pulumi.String("examplePartnerTopic1"),
				ResourceGroupName:   pulumi.String("examplerg2"),
				Source:              pulumi.String("ContosoCorp.Accounts.User1"),
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
