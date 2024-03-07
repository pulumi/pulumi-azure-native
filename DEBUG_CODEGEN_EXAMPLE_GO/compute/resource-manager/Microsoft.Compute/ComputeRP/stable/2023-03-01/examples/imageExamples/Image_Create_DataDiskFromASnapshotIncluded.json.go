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
						Lun: pulumi.Int(1),
						Snapshot: &compute.SubResourceArgs{
							Id: pulumi.String("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
						},
					},
				},
				OsDisk: &compute.ImageOSDiskArgs{
					OsState: compute.OperatingSystemStateTypesGeneralized,
					OsType:  compute.OperatingSystemTypesLinux,
					Snapshot: &compute.SubResourceArgs{
						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
					},
				},
				ZoneResilient: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
