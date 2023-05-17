package main

import (
	resources "github.com/pulumi/pulumi-azure-native-sdk/resources"
	storage "github.com/pulumi/pulumi-azure-native-sdk/storage/v20210201"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		resourceGroupNameParam := cfg.Require("resourceGroupNameParam")
		resourceGroupVar, err := resources.LookupResourceGroup(ctx, &resources.LookupResourceGroupArgs{
			ResourceGroupName: resourceGroupNameParam,
		}, nil)
		if err != nil {
			return err
		}
		locationParam := resourceGroupVar.Location
		if param := cfg.Get("locationParam"); param != "" {
			locationParam = param
		}
		storageAccountNameParam := cfg.Require("storageAccountNameParam")
		storageAccountTypeParam := "Standard_LRS"
		if param := cfg.Get("storageAccountTypeParam"); param != "" {
			storageAccountTypeParam = param
		}
		storageAccountResource, err := storage.NewStorageAccount(ctx, "storageAccountResource", &storage.StorageAccountArgs{
			AccountName:       pulumi.String(storageAccountNameParam),
			Kind:              pulumi.String("StorageV2"),
			Location:          pulumi.String(locationParam),
			ResourceGroupName: pulumi.String(resourceGroupNameParam),
			Sku: &storage.SkuArgs{
				Name: pulumi.String(storageAccountTypeParam),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("storageAccountNameOut", storageAccountResource.Name)
		return nil
	})
}
