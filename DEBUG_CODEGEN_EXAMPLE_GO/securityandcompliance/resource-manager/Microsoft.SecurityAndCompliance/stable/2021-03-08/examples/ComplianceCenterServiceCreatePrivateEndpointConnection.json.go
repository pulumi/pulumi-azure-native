package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityandcompliance.NewPrivateEndpointConnectionsComp(ctx, "privateEndpointConnectionsComp", &securityandcompliance.PrivateEndpointConnectionsCompArgs{
			PrivateEndpointConnectionName: pulumi.String("myConnection"),
			PrivateLinkServiceConnectionState: &securityandcompliance.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rgname"),
			ResourceName:      pulumi.String("service1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
