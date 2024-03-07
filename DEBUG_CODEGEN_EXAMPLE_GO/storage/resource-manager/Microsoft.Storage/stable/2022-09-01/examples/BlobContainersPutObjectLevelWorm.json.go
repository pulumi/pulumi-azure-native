package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewBlobContainer(ctx, "blobContainer", &storage.BlobContainerArgs{
			AccountName:   pulumi.String("sto328"),
			ContainerName: pulumi.String("container6185"),
			ImmutableStorageWithVersioning: &storage.ImmutableStorageWithVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
			ResourceGroupName: pulumi.String("res3376"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
