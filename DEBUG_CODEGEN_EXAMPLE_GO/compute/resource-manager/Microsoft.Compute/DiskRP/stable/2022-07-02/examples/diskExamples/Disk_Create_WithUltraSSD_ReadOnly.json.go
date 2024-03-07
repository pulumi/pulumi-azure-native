package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption:      pulumi.String("Empty"),
				LogicalSectorSize: pulumi.Int(4096),
			},
			DiskIOPSReadWrite: pulumi.Float64(125),
			DiskMBpsReadWrite: pulumi.Float64(3000),
			DiskName:          pulumi.String("myUltraReadOnlyDisk"),
			DiskSizeGB:        pulumi.Int(200),
			Encryption: &compute.EncryptionArgs{
				Type: pulumi.String("EncryptionAtRestWithPlatformKey"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &compute.DiskSkuArgs{
				Name: pulumi.String("UltraSSD_LRS"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
