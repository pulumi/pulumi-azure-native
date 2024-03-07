package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/botservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := botservice.NewChannel(ctx, "channel", &botservice.ChannelArgs{
			ChannelName: pulumi.String("EmailChannel"),
			Location:    pulumi.String("global"),
			Properties: botservice.EmailChannel{
				ChannelName: "EmailChannel",
				Properties: botservice.EmailChannelProperties{
					EmailAddress: "a@b.com",
					IsEnabled:    true,
					Password:     "pwd",
				},
			},
			ResourceGroupName: pulumi.String("OneResourceGroupName"),
			ResourceName:      pulumi.String("samplebotname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
