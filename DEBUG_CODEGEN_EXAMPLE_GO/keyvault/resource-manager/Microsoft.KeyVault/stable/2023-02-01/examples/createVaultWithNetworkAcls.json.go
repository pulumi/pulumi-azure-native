package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewVault(ctx, "vault", &keyvault.VaultArgs{
			Location: pulumi.String("westus"),
			Properties: &keyvault.VaultPropertiesArgs{
				EnabledForDeployment:         pulumi.Bool(true),
				EnabledForDiskEncryption:     pulumi.Bool(true),
				EnabledForTemplateDeployment: pulumi.Bool(true),
				NetworkAcls: &keyvault.NetworkRuleSetArgs{
					Bypass:        pulumi.String("AzureServices"),
					DefaultAction: pulumi.String("Deny"),
					IpRules: keyvault.IPRuleArray{
						&keyvault.IPRuleArgs{
							Value: pulumi.String("124.56.78.91"),
						},
						&keyvault.IPRuleArgs{
							Value: pulumi.String("'10.91.4.0/24'"),
						},
					},
					VirtualNetworkRules: keyvault.VirtualNetworkRuleArray{
						&keyvault.VirtualNetworkRuleArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
						},
					},
				},
				Sku: &keyvault.SkuArgs{
					Family: pulumi.String("A"),
					Name:   keyvault.SkuNameStandard,
				},
				TenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
			},
			ResourceGroupName: pulumi.String("sample-resource-group"),
			VaultName:         pulumi.String("sample-vault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
