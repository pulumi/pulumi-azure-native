package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewAppServiceEnvironment(ctx, "appServiceEnvironment", &web.AppServiceEnvironmentArgs{
			Kind:              pulumi.String("Asev3"),
			Location:          pulumi.String("South Central US"),
			Name:              pulumi.String("test-ase"),
			ResourceGroupName: pulumi.String("test-rg"),
			VirtualNetwork: &web.VirtualNetworkProfileArgs{
				Id: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/delegated"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
