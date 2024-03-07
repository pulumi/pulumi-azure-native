package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewEventSubscription(ctx, "eventSubscription", &eventgrid.EventSubscriptionArgs{
			DeadLetterDestination: &eventgrid.StorageBlobDeadLetterDestinationArgs{
				BlobContainerName: pulumi.String("contosocontainer"),
				EndpointType:      pulumi.String("StorageBlob"),
				ResourceId:        pulumi.String("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
			},
			Destination: eventgrid.HybridConnectionEventSubscriptionDestination{
				EndpointType: "HybridConnection",
				ResourceId:   "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1",
			},
			EventSubscriptionName: pulumi.String("examplesubscription1"),
			Filter: &eventgrid.EventSubscriptionFilterArgs{
				IsSubjectCaseSensitive: pulumi.Bool(false),
				SubjectBeginsWith:      pulumi.String("ExamplePrefix"),
				SubjectEndsWith:        pulumi.String("ExampleSuffix"),
			},
			Scope: pulumi.String("subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
