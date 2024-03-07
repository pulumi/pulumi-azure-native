package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/notificationhubs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := notificationhubs.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &notificationhubs.PrivateEndpointConnectionArgs{
			NamespaceName:                 pulumi.String("nh-sdk-ns"),
			PrivateEndpointConnectionName: pulumi.String("nh-sdk-ns.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
			Properties: &notificationhubs.PrivateEndpointConnectionPropertiesArgs{
				PrivateLinkServiceConnectionState: &notificationhubs.RemotePrivateLinkServiceConnectionStateArgs{
					Status: pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("5ktrial"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
