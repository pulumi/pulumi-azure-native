package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewLiveEvent(ctx, "liveEvent", &media.LiveEventArgs{
			AccountName: pulumi.String("slitestmedia10"),
			Description: pulumi.String("test event 1"),
			Input: &media.LiveEventInputTypeArgs{
				AccessControl: &media.LiveEventInputAccessControlArgs{
					Ip: &media.IPAccessControlArgs{
						Allow: media.IPRangeArray{
							&media.IPRangeArgs{
								Address:            pulumi.String("0.0.0.0"),
								Name:               pulumi.String("AllowAll"),
								SubnetPrefixLength: pulumi.Int(0),
							},
						},
					},
				},
				KeyFrameIntervalDuration: pulumi.String("PT6S"),
				StreamingProtocol:        pulumi.String("RTMP"),
			},
			LiveEventName: pulumi.String("myLiveEvent1"),
			Location:      pulumi.String("West US"),
			Preview: &media.LiveEventPreviewArgs{
				AccessControl: &media.LiveEventPreviewAccessControlArgs{
					Ip: &media.IPAccessControlArgs{
						Allow: media.IPRangeArray{
							&media.IPRangeArgs{
								Address:            pulumi.String("0.0.0.0"),
								Name:               pulumi.String("AllowAll"),
								SubnetPrefixLength: pulumi.Int(0),
							},
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("mediaresources"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
