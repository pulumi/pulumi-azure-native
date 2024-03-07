package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewGalleryImageVersion(ctx, "galleryImageVersion", &compute.GalleryImageVersionArgs{
GalleryImageName: pulumi.String("myGalleryImageName"),
GalleryImageVersionName: pulumi.String("1.0.0"),
GalleryName: pulumi.String("myGalleryName"),
Location: pulumi.String("West US"),
PublishingProfile: compute.GalleryImageVersionPublishingProfileResponse{
TargetRegions: compute.TargetRegionArray{
interface{}{
Encryption: interface{}{
DataDiskImages: compute.DataDiskImageEncryptionArray{
&compute.DataDiskImageEncryptionArgs{
DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
Lun: pulumi.Int(1),
},
},
OsDiskImage: &compute.OSDiskImageEncryptionArgs{
DiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
},
},
ExcludeFromLatest: pulumi.Bool(false),
Name: pulumi.String("West US"),
RegionalReplicaCount: pulumi.Int(1),
},
&compute.TargetRegionArgs{
ExcludeFromLatest: pulumi.Bool(false),
Name: pulumi.String("East US"),
RegionalReplicaCount: pulumi.Int(2),
StorageAccountType: pulumi.String("Standard_ZRS"),
},
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
SafetyProfile: &compute.GalleryImageVersionSafetyProfileArgs{
AllowDeletionOfReplicatedLocations: pulumi.Bool(false),
},
StorageProfile: compute.GalleryImageVersionStorageProfileResponse{
DataDiskImages: compute.GalleryDataDiskImageArray{
interface{}{
HostCaching: compute.HostCachingNone,
Lun: pulumi.Int(1),
Source: &compute.GalleryDiskImageSourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
Uri: pulumi.String("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
},
},
},
OsDiskImage: interface{}{
HostCaching: compute.HostCachingReadOnly,
Source: &compute.GalleryDiskImageSourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
Uri: pulumi.String("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
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
