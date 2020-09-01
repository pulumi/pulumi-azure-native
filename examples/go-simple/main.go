package main


import (
	resources "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/resources/latest"
	storage "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/storage/latest"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", &resources.ResourceGroupArgs{
			ResourceGroupName: pulumi.String("azurerm-go"),
			Location:          pulumi.String("WestUS"),
		})
		if err != nil {
			return err
		}

		_, err = storage.NewStorageAccount(ctx, "sa", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			AccountName: pulumi.String("pulumi14345sago"),
			Location: pulumi.String("westus2"),
			Sku: &storage.SkuArgs {
				Name: pulumi.String("Standard_LRS"),
				Tier: pulumi.String("Standard"),
			},
			Kind: pulumi.String("StorageV2"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
