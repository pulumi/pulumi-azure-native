package main

import (
	"github.com/pulumi/pulumi-azurerm/sdk/go/azurerm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		_, err := azurerm.NewResourceGroup(ctx, "resourceGroup", &azurerm.ResourceGroupArgs{
			Name:     pulumi.String("azurerm-go"),
			Location: pulumi.String("WestUS"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
