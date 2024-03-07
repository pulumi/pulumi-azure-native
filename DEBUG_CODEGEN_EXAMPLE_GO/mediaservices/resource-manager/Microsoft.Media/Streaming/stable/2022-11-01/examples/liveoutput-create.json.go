package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewLiveOutput(ctx, "liveOutput", &media.LiveOutputArgs{
			AccountName:         pulumi.String("slitestmedia10"),
			ArchiveWindowLength: pulumi.String("PT5M"),
			AssetName:           pulumi.String("6f3264f5-a189-48b4-a29a-a40f22575212"),
			Description:         pulumi.String("test live output 1"),
			Hls: &media.HlsArgs{
				FragmentsPerTsSegment: pulumi.Int(5),
			},
			LiveEventName:      pulumi.String("myLiveEvent1"),
			LiveOutputName:     pulumi.String("myLiveOutput1"),
			ManifestName:       pulumi.String("testmanifest"),
			ResourceGroupName:  pulumi.String("mediaresources"),
			RewindWindowLength: pulumi.String("PT4M"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
