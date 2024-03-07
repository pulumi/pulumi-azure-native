package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewStorageAccount(ctx, "storageAccount", &storage.StorageAccountArgs{
			AccountName:          pulumi.String("sto4445"),
			AllowSharedKeyAccess: pulumi.Bool(true),
			Encryption: &storage.EncryptionArgs{
				KeySource:                       pulumi.String("Microsoft.Storage"),
				RequireInfrastructureEncryption: pulumi.Bool(false),
				Services: &storage.EncryptionServicesArgs{
					Blob: &storage.EncryptionServiceArgs{
						Enabled: pulumi.Bool(true),
						KeyType: pulumi.String("Account"),
					},
					File: &storage.EncryptionServiceArgs{
						Enabled: pulumi.Bool(true),
						KeyType: pulumi.String("Account"),
					},
				},
			},
			Kind:              pulumi.String("BlockBlobStorage"),
			Location:          pulumi.String("eastus"),
			MinimumTlsVersion: pulumi.String("TLS1_2"),
			ResourceGroupName: pulumi.String("res9101"),
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Premium_LRS"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
