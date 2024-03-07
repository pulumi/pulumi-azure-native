package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewInventoryItem(ctx, "inventoryItem", &scvmm.InventoryItemArgs{
			InventoryItemName: pulumi.String("12345678-1234-1234-1234-123456789abc"),
			InventoryType:     pulumi.String("Cloud"),
			ResourceGroupName: pulumi.String("testrg"),
			VmmServerName:     pulumi.String("ContosoVMMServer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
