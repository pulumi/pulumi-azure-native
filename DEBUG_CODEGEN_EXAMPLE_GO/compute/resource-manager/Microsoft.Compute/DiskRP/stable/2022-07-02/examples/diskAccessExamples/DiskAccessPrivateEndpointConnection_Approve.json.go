package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDiskAccessAPrivateEndpointConnection(ctx, "diskAccessAPrivateEndpointConnection", &compute.DiskAccessAPrivateEndpointConnectionArgs{
			DiskAccessName:                pulumi.String("myDiskAccess"),
			PrivateEndpointConnectionName: pulumi.String("myPrivateEndpointConnection"),
			PrivateLinkServiceConnectionState: &compute.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approving myPrivateEndpointConnection"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
