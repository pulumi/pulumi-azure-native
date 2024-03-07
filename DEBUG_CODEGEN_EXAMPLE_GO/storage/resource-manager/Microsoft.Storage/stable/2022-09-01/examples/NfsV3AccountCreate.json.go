package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewStorageAccount(ctx, "storageAccount", &storage.StorageAccountArgs{
			AccountName:            pulumi.String("sto4445"),
			EnableHttpsTrafficOnly: pulumi.Bool(false),
			EnableNfsV3:            pulumi.Bool(true),
			IsHnsEnabled:           pulumi.Bool(true),
			Kind:                   pulumi.String("BlockBlobStorage"),
			Location:               pulumi.String("eastus"),
			NetworkRuleSet: storage.NetworkRuleSetResponse{
				Bypass:        pulumi.String("AzureServices"),
				DefaultAction: storage.DefaultActionAllow,
				IpRules:       storage.IPRuleArray{},
				VirtualNetworkRules: storage.VirtualNetworkRuleArray{
					&storage.VirtualNetworkRuleArgs{
						VirtualNetworkResourceId: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Network/virtualNetworks/net123/subnets/subnet12"),
					},
				},
			},
			ResourceGroupName: pulumi.String("res9101"),
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Premium_LRS"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
