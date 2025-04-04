// Copyright 2025, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		_, err = resources.NewResource(ctx, "acc", &resources.ResourceArgs{
			ResourceProviderNamespace: pulumi.String("Microsoft.Storage"),
			ResourceType:              pulumi.String("storageAccounts"),
			ApiVersion:                pulumi.String("2024-01-01"),
			ParentResourcePath:        pulumi.String(""),
			ResourceGroupName:         resourceGroup.Name,
			Kind:                      pulumi.String("StorageV2"),
			Sku: &resources.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
