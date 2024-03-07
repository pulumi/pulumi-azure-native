package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewInventoryItem(ctx, "inventoryItem", &connectedvmwarevsphere.InventoryItemArgs{
			InventoryItemName: pulumi.String("testItem"),
			InventoryType:     pulumi.String("ResourcePool"),
			ResourceGroupName: pulumi.String("testrg"),
			VcenterName:       pulumi.String("ContosoVCenter"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
