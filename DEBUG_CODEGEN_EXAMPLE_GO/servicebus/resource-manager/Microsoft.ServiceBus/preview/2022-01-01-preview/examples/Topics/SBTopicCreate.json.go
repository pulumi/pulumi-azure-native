package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewTopic(ctx, "topic", &servicebus.TopicArgs{
			EnableExpress:     pulumi.Bool(true),
			NamespaceName:     pulumi.String("sdk-Namespace-1617"),
			ResourceGroupName: pulumi.String("ArunMonocle"),
			TopicName:         pulumi.String("sdk-Topics-5488"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
