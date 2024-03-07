package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewVideo(ctx, "video", &videoanalyzer.VideoArgs{
			AccountName:       pulumi.String("testaccount2"),
			Description:       pulumi.String("Sample Description 1"),
			ResourceGroupName: pulumi.String("testrg"),
			Title:             pulumi.String("Sample Title 1"),
			VideoName:         pulumi.String("video1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
