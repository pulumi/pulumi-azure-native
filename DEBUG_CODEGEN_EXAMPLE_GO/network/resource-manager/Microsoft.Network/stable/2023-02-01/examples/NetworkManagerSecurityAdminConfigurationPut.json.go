package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSecurityAdminConfiguration(ctx, "securityAdminConfiguration", &network.SecurityAdminConfigurationArgs{
			ApplyOnNetworkIntentPolicyBasedServices: pulumi.StringArray{
				pulumi.String("None"),
			},
			ConfigurationName:  pulumi.String("myTestSecurityConfig"),
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
