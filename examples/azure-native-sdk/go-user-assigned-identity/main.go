// Copyright 2021, Pulumi Corporation.  All rights reserved.

package main

import (
	managedidentity "github.com/pulumi/pulumi-azure-native-sdk/managedidentity"
	resources "github.com/pulumi/pulumi-azure-native-sdk/resources"
	storage "github.com/pulumi/pulumi-azure-native-sdk/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		userAssignedId, err := managedidentity.NewUserAssignedIdentity(ctx, "ua-id", &managedidentity.UserAssignedIdentityArgs{
			ResourceGroupName: resourceGroup.Name,
		})
		if err != nil {
			return err
		}

		// Create an Azure resource (Storage Account)
		account, err := storage.NewStorageAccount(ctx, "sa", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
			Kind: pulumi.String("StorageV2"),
			Identity: &storage.IdentityArgs{
				Type:                   pulumi.String("UserAssigned"),
				UserAssignedIdentities: pulumi.StringArray{userAssignedId.ID()},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
