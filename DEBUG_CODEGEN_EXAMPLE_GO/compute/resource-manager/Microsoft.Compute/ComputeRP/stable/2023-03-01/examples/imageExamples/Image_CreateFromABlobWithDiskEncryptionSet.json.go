package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
ImageName: pulumi.String("myImage"),
Location: pulumi.String("West US"),
ResourceGroupName: pulumi.String("myResourceGroup"),
StorageProfile: compute.ImageStorageProfileResponse{
OsDisk: interface{}{
BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
},
OsState: compute.OperatingSystemStateTypesGeneralized,
OsType: compute.OperatingSystemTypesLinux,
},
},
})
if err != nil {
return err
}
return nil
})
}
