package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerTopicEventSubscription(ctx, "partnerTopicEventSubscription", &eventgrid.PartnerTopicEventSubscriptionArgs{
			Destination: eventgrid.WebHookEventSubscriptionDestination{
				EndpointType: "WebHook",
				EndpointUrl:  "https://requestb.in/15ksip71",
			},
			EventSubscriptionName: pulumi.String("exampleEventSubscriptionName1"),
			Filter: &eventgrid.EventSubscriptionFilterArgs{
				IsSubjectCaseSensitive: pulumi.Bool(false),
				SubjectBeginsWith:      pulumi.String("ExamplePrefix"),
				SubjectEndsWith:        pulumi.String("ExampleSuffix"),
			},
			PartnerTopicName:  pulumi.String("examplePartnerTopic1"),
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
