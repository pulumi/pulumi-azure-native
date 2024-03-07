package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewAsset(ctx, "asset", &media.AssetArgs{
			AccountName:        pulumi.String("contosomedia"),
			AssetName:          pulumi.String("ClimbingMountLogan"),
			Description:        pulumi.String("A documentary showing the ascent of Mount Logan"),
			EncryptionScope:    pulumi.String("encryptionScope1"),
			ResourceGroupName:  pulumi.String("contosorg"),
			StorageAccountName: pulumi.String("storage0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
