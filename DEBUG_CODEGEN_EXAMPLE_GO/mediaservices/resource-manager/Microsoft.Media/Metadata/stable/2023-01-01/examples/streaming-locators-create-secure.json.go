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
			EndTime:              pulumi.String("2028-12-31T23:59:59.9999999Z"),
			ResourceGroupName:    pulumi.String("contosorg"),
			StartTime:            pulumi.String("2018-03-01T00:00:00Z"),
			StreamingLocatorName: pulumi.String("UserCreatedSecureStreamingLocator"),
			StreamingPolicyName:  pulumi.String("UserCreatedSecureStreamingPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
