package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewStreamingLocator(ctx, "streamingLocator", &media.StreamingLocatorArgs{
			AccountName: pulumi.String("contosomedia"),
			AssetName:   pulumi.String("ClimbingMountRainier"),
			ContentKeys: []media.StreamingLocatorContentKeyArgs{
				{
					Id:                              pulumi.String("60000000-0000-0000-0000-000000000001"),
					LabelReferenceInStreamingPolicy: pulumi.String("aesDefaultKey"),
					Value:                           pulumi.String("1UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					Id:                              pulumi.String("60000000-0000-0000-0000-000000000004"),
					LabelReferenceInStreamingPolicy: pulumi.String("cencDefaultKey"),
					Value:                           pulumi.String("4UqLohAfWsEGkULYxHjYZg=="),
				},
				{
					Id:                              pulumi.String("60000000-0000-0000-0000-000000000007"),
					LabelReferenceInStreamingPolicy: pulumi.String("cbcsDefaultKey"),
					Value:                           pulumi.String("7UqLohAfWsEGkULYxHjYZg=="),
				},
			},
			ResourceGroupName:    pulumi.String("contosorg"),
			StreamingLocatorId:   pulumi.String("90000000-0000-0000-0000-00000000000A"),
			StreamingLocatorName: pulumi.String("UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys"),
			StreamingPolicyName:  pulumi.String("secureStreamingPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
