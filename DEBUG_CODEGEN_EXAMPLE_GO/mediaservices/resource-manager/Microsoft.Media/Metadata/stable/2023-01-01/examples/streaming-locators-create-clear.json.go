package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewStreamingLocator(ctx, "streamingLocator", &media.StreamingLocatorArgs{
			AccountName:          pulumi.String("contosomedia"),
			AssetName:            pulumi.String("ClimbingMountRainier"),
			ResourceGroupName:    pulumi.String("contosorg"),
			StreamingLocatorName: pulumi.String("UserCreatedClearStreamingLocator"),
			StreamingPolicyName:  pulumi.String("clearStreamingPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
