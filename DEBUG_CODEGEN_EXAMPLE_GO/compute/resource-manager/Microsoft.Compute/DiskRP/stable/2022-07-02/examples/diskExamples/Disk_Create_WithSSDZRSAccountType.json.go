package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption: pulumi.String("Empty"),
			},
			DiskName:          pulumi.String("myDisk"),
			DiskSizeGB:        pulumi.Int(200),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &compute.DiskSkuArgs{
				Name: pulumi.String("Premium_ZRS"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
