package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagemover/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagemover.NewStorageMover(ctx, "storageMover", &storagemover.StorageMoverArgs{
			Description:       pulumi.String("Example Storage Mover Description"),
			Location:          pulumi.String("eastus2"),
			ResourceGroupName: pulumi.String("examples-rg"),
			StorageMoverName:  pulumi.String("examples-storageMoverName"),
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
