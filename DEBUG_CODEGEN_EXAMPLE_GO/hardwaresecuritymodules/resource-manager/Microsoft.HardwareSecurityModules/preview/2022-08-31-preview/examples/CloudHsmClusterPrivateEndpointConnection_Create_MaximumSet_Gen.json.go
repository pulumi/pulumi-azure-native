package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hardwaresecuritymodules/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hardwaresecuritymodules.NewCloudHsmClusterPrivateEndpointConnection(ctx, "cloudHsmClusterPrivateEndpointConnection", &hardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionArgs{
			CloudHsmClusterName: pulumi.String("chsm1"),
			PeConnectionName:    pulumi.String("sample-pec"),
			PrivateLinkServiceConnectionState: &hardwaresecuritymodules.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("My name is Joe and I'm approving this."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rgcloudhsm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
