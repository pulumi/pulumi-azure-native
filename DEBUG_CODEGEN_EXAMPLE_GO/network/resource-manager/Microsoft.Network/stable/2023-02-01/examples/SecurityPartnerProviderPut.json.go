package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSecurityPartnerProvider(ctx, "securityPartnerProvider", &network.SecurityPartnerProviderArgs{
			Location:                    pulumi.String("West US"),
			ResourceGroupName:           pulumi.String("rg1"),
			SecurityPartnerProviderName: pulumi.String("securityPartnerProvider"),
			SecurityProviderName:        pulumi.String("ZScaler"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualHub: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
