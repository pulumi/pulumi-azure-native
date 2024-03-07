package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/timeseriesinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timeseriesinsights.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &timeseriesinsights.PrivateEndpointConnectionArgs{
			EnvironmentName:               pulumi.String("myEnvironment"),
			PrivateEndpointConnectionName: pulumi.String("myPrivateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &timeseriesinsights.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Rejected for some reason"),
				Status:      pulumi.String("Rejected"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
