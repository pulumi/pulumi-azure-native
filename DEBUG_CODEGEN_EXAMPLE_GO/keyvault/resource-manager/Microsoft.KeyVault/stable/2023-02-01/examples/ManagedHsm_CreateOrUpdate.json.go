package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewManagedHsm(ctx, "managedHsm", &keyvault.ManagedHsmArgs{
			Location: pulumi.String("westus"),
			Name:     pulumi.String("hsm1"),
			Properties: &keyvault.ManagedHsmPropertiesArgs{
				EnablePurgeProtection: pulumi.Bool(false),
				EnableSoftDelete:      pulumi.Bool(true),
				InitialAdminObjectIds: pulumi.StringArray{
					pulumi.String("00000000-0000-0000-0000-000000000000"),
				},
				SoftDeleteRetentionInDays: pulumi.Int(90),
				TenantId:                  pulumi.String("00000000-0000-0000-0000-000000000000"),
			},
			ResourceGroupName: pulumi.String("hsm-group"),
			Sku: &keyvault.ManagedHsmSkuArgs{
				Family: pulumi.String("B"),
				Name:   keyvault.ManagedHsmSkuName_Standard_B1,
			},
			Tags: pulumi.StringMap{
				"Dept":        pulumi.String("hsm"),
				"Environment": pulumi.String("dogfood"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
