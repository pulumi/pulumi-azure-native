package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &containerservice.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("privateendpointconnection1"),
			PrivateLinkServiceConnectionState: &containerservice.PrivateLinkServiceConnectionStateArgs{
				Status: pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("clustername1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
