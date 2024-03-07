package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewNamespaceTopic(ctx, "namespaceTopic", &eventgrid.NamespaceTopicArgs{
			EventRetentionInDays: pulumi.Int(1),
			InputSchema:          pulumi.String("CloudEventSchemaV1_0"),
			NamespaceName:        pulumi.String("examplenamespace2"),
			PublisherType:        pulumi.String("Custom"),
			ResourceGroupName:    pulumi.String("examplerg"),
			TopicName:            pulumi.String("examplenamespacetopic2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
