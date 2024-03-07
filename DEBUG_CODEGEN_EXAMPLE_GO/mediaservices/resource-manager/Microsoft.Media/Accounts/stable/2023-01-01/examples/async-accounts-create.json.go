package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewMediaService(ctx, "mediaService", &media.MediaServiceArgs{
			AccountName:       pulumi.String("contososports"),
			Location:          pulumi.String("South Central US"),
			ResourceGroupName: pulumi.String("contosorg"),
			StorageAccounts: media.StorageAccountArray{
				&media.StorageAccountArgs{
					Id:   pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosorg/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
					Type: pulumi.String("Primary"),
				},
			},
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
