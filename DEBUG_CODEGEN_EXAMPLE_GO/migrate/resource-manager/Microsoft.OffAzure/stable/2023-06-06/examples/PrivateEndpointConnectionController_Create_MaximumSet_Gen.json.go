package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewPrivateEndpointConnectionController(ctx, "privateEndpointConnectionController", &offazure.PrivateEndpointConnectionControllerArgs{
			PeConnectionName: pulumi.String("-KF3-86-D3L-8M"),
			PrivateLinkServiceConnectionState: &offazure.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("swxghzuajrasojk"),
				Description:     pulumi.String("ljdykrrhmzi"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("LuBO2I1-B85pJ1BZ251"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
