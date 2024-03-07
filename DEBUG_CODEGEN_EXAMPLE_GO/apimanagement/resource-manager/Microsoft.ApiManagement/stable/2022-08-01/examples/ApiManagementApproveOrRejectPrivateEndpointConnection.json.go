package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewPrivateEndpointConnectionByName(ctx, "privateEndpointConnectionByName", &apimanagement.PrivateEndpointConnectionByNameArgs{
			Id:                            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/connectionName"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			Properties: &apimanagement.PrivateEndpointConnectionRequestPropertiesArgs{
				PrivateLinkServiceConnectionState: &apimanagement.PrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("The Private Endpoint Connection is approved."),
					Status:      pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
