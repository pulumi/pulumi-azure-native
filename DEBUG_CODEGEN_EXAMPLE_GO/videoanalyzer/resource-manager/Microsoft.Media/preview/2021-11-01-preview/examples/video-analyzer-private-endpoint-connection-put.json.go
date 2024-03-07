package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &videoanalyzer.PrivateEndpointConnectionArgs{
			AccountName: pulumi.String("contososports"),
			Name:        pulumi.String("10000000-0000-0000-0000-000000000000"),
			PrivateLinkServiceConnectionState: &videoanalyzer.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Test description."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("contoso"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
