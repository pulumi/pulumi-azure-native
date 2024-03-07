package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/deviceupdate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := deviceupdate.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &deviceupdate.PrivateEndpointConnectionArgs{
			AccountName:                   pulumi.String("contoso"),
			PrivateEndpointConnectionName: pulumi.String("peexample01"),
			PrivateLinkServiceConnectionState: &deviceupdate.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
