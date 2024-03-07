package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewAppServiceEnvironmentPrivateEndpointConnection(ctx, "appServiceEnvironmentPrivateEndpointConnection", &web.AppServiceEnvironmentPrivateEndpointConnectionArgs{
			Name:                          pulumi.String("test-ase"),
			PrivateEndpointConnectionName: pulumi.String("fa38656c-034e-43d8-adce-fe06ce039c98"),
			PrivateLinkServiceConnectionState: &web.PrivateLinkConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@company.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
