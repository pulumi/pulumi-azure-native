package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/botservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := botservice.NewBotConnection(ctx, "botConnection", &botservice.BotConnectionArgs{
			ConnectionName: pulumi.String("sampleConnection"),
			Location:       pulumi.String("West US"),
			Properties: botservice.ConnectionSettingPropertiesResponse{
				ClientId:     pulumi.String("sampleclientid"),
				ClientSecret: pulumi.String("samplesecret"),
				Parameters: botservice.ConnectionSettingParameterArray{
					&botservice.ConnectionSettingParameterArgs{
						Key:   pulumi.String("key1"),
						Value: pulumi.String("value1"),
					},
					&botservice.ConnectionSettingParameterArgs{
						Key:   pulumi.String("key2"),
						Value: pulumi.String("value2"),
					},
				},
				Scopes:            pulumi.String("samplescope"),
				ServiceProviderId: pulumi.String("serviceproviderid"),
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
