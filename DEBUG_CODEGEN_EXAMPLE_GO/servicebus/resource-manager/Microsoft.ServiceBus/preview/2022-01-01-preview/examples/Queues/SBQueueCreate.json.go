package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewQueue(ctx, "queue", &servicebus.QueueArgs{
			EnablePartitioning: pulumi.Bool(true),
			NamespaceName:      pulumi.String("sdk-Namespace-3174"),
			QueueName:          pulumi.String("sdk-Queues-5647"),
			ResourceGroupName:  pulumi.String("ArunMonocle"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
