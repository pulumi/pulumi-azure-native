package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewPublisher(ctx, "publisher", &hybridnetwork.PublisherArgs{
			Location: pulumi.String("eastus"),
			Properties: &hybridnetwork.PublisherPropertiesFormatArgs{
				Scope: pulumi.String("Public"),
			},
			PublisherName:     pulumi.String("TestPublisher"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
