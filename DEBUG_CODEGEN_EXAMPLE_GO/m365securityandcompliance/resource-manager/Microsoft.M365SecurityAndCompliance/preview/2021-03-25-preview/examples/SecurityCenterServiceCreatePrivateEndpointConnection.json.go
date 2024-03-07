package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/m365securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := m365securityandcompliance.NewPrivateEndpointConnectionsSec(ctx, "privateEndpointConnectionsSec", &m365securityandcompliance.PrivateEndpointConnectionsSecArgs{
			PrivateEndpointConnectionName: pulumi.String("myConnection"),
			PrivateLinkServiceConnectionState: &m365securityandcompliance.PrivateLinkServiceConnectionStateArgs{
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
