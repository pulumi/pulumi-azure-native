package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewGalleryImageVersion(ctx, "galleryImageVersion", &compute.GalleryImageVersionArgs{
			GalleryImageName:        pulumi.String("myGalleryImageName"),
			GalleryImageVersionName: pulumi.String("1.0.0"),
			GalleryName:             pulumi.String("myGalleryName"),
			Location:                pulumi.String("West US"),
			PublishingProfile: &compute.GalleryImageVersionPublishingProfileArgs{
				TargetRegions: compute.TargetRegionArray{
					&compute.TargetRegionArgs{
						Encryption: &compute.EncryptionImagesArgs{
							DataDiskImages: compute.DataDiskImageEncryptionArray{
								&compute.DataDiskImageEncryptionArgs{
									DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
									Lun:                 pulumi.Int(1),
								},
							},
							OsDiskImage: &compute.OSDiskImageEncryptionArgs{
								DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    pulumi.Bool(false),
						Name:                 pulumi.String("West US"),
						RegionalReplicaCount: pulumi.Int(1),
					},
					&compute.TargetRegionArgs{
						Encryption: &compute.EncryptionImagesArgs{
							DataDiskImages: compute.DataDiskImageEncryptionArray{
								&compute.DataDiskImageEncryptionArgs{
									DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
									Lun:                 pulumi.Int(1),
								},
							},
							OsDiskImage: &compute.OSDiskImageEncryptionArgs{
								DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    pulumi.Bool(false),
						Name:                 pulumi.String("East US"),
						RegionalReplicaCount: pulumi.Int(2),
						StorageAccountType:   pulumi.String("Standard_ZRS"),
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SafetyProfile: &compute.GalleryImageVersionSafetyProfileArgs{
				AllowDeletionOfReplicatedLocations: pulumi.Bool(false),
			},
			StorageProfile: &compute.GalleryImageVersionStorageProfileArgs{
				DataDiskImages: compute.GalleryDataDiskImageArray{
					&compute.GalleryDataDiskImageArgs{
						HostCaching: compute.HostCachingNone,
						Lun:         pulumi.Int(1),
						Source: &compute.GalleryDiskImageSourceArgs{
							Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}"),
						},
					},
				},
				OsDiskImage: &compute.GalleryOSDiskImageArgs{
					HostCaching: compute.HostCachingReadOnly,
					Source: &compute.GalleryDiskImageSourceArgs{
						Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
