package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuresphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuresphere.NewProduct(ctx, "product", &azuresphere.ProductArgs{
			CatalogName:       pulumi.String("MyCatalog1"),
			ProductName:       pulumi.String("MyProduct1"),
			ResourceGroupName: pulumi.String("MyResourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
