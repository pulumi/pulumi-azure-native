package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewPrivateEndpointConnectionController(ctx, "privateEndpointConnectionController", &offazure.PrivateEndpointConnectionControllerArgs{
			PeConnectionName:  pulumi.String("-p-U2"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("1--tOY-47Q-X95WC3d-U"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
