package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/search/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := search.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &search.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546"),
			Properties: &search.PrivateEndpointConnectionPropertiesArgs{
				PrivateLinkServiceConnectionState: &search.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("Rejected for some reason"),
					Status:      search.PrivateLinkServiceConnectionStatusRejected,
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			SearchServiceName: pulumi.String("mysearchservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
