package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aadiam/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aadiam.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &aadiam.PrivateEndpointConnectionArgs{
			PolicyName: pulumi.String("example-policy-5849"),
			PrivateEndpoint: &aadiam.PrivateEndpointArgs{
				Id: pulumi.String("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/microsoft.aadiam/privateLinkForAzureAD/ddb1/privateLinkConnections/{privateEndpointConnection name}"),
			},
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnection name}"),
			PrivateLinkServiceConnectionState: &aadiam.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("You may pass"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("resourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
