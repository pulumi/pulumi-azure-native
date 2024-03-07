package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewNamespaceTopicEventSubscription(ctx, "namespaceTopicEventSubscription", &eventgrid.NamespaceTopicEventSubscriptionArgs{
			DeliveryConfiguration: eventgrid.DeliveryConfigurationResponse{
				DeliveryMode: pulumi.String("Queue"),
				Queue: &eventgrid.QueueInfoArgs{
					EventTimeToLive:              pulumi.String("P1D"),
					MaxDeliveryCount:             pulumi.Int(4),
					ReceiveLockDurationInSeconds: pulumi.Int(60),
				},
			},
			EventDeliverySchema:   pulumi.String("CloudEventSchemaV1_0"),
			EventSubscriptionName: pulumi.String("examplenamespacetopicEventSub2"),
			NamespaceName:         pulumi.String("examplenamespace2"),
			ResourceGroupName:     pulumi.String("examplerg"),
			TopicName:             pulumi.String("examplenamespacetopic2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
