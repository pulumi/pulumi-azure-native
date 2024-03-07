package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/engagementfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := engagementfabric.NewChannel(ctx, "channel", &engagementfabric.ChannelArgs{
			AccountName: pulumi.String("ExampleAccount"),
			ChannelFunctions: pulumi.StringArray{
				pulumi.String("MockFunction1"),
				pulumi.String("MockFunction2"),
			},
			ChannelName: pulumi.String("ExampleChannel"),
			ChannelType: pulumi.String("MockChannel"),
			Credentials: pulumi.StringMap{
				"AppId":  pulumi.String("exampleApp"),
				"AppKey": pulumi.String("exampleAppKey"),
			},
			ResourceGroupName: pulumi.String("ExampleRg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
