package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewEventSubscription(ctx, "eventSubscription", &eventgrid.EventSubscriptionArgs{
			Destination: eventgrid.WebHookEventSubscriptionDestination{
				EndpointType: "WebHook",
				EndpointUrl:  "https://requestb.in/15ksip71",
			},
			EventSubscriptionName: pulumi.String("examplesubscription2"),
			Filter: &eventgrid.EventSubscriptionFilterArgs{
				IsSubjectCaseSensitive: pulumi.Bool(false),
				SubjectBeginsWith:      pulumi.String("ExamplePrefix"),
				SubjectEndsWith:        pulumi.String("ExampleSuffix"),
			},
			Scope: pulumi.String("subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
