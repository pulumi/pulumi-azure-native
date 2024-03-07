package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/purview/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := purview.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &purview.PrivateEndpointConnectionArgs{
			AccountName:                   pulumi.String("account1"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnection1"),
			PrivateLinkServiceConnectionState: &purview.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@company.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
