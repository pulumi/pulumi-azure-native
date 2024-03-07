package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewConsumerGroup(ctx, "consumerGroup", &eventhub.ConsumerGroupArgs{
			ConsumerGroupName: pulumi.String("sdk-ConsumerGroup-5563"),
			EventHubName:      pulumi.String("sdk-EventHub-6681"),
			NamespaceName:     pulumi.String("sdk-Namespace-2661"),
			ResourceGroupName: pulumi.String("ArunMonocle"),
			UserMetadata:      pulumi.String("New consumergroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
