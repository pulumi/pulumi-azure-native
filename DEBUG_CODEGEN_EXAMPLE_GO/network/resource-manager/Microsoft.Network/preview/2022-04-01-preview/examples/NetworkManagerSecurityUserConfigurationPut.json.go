package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSecurityUserConfiguration(ctx, "securityUserConfiguration", &network.SecurityUserConfigurationArgs{
			ConfigurationName:  pulumi.String("myTestSecurityConfig"),
			DeleteExistingNSGs: pulumi.String("True"),
			Description:        pulumi.String("A sample policy"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			ResourceGroupName:  pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
