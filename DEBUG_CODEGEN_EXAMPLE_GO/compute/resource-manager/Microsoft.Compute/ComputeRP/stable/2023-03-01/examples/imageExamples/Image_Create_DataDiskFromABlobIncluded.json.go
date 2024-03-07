package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
			ImageName:         pulumi.String("myImage"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			StorageProfile: &compute.ImageStorageProfileArgs{
				DataDisks: compute.ImageDataDiskArray{
					&compute.ImageDataDiskArgs{
						BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
						Lun:     pulumi.Int(1),
					},
				},
				OsDisk: &compute.ImageOSDiskArgs{
					BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
					OsState: compute.OperatingSystemStateTypesGeneralized,
					OsType:  compute.OperatingSystemTypesLinux,
				},
				ZoneResilient: pulumi.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
