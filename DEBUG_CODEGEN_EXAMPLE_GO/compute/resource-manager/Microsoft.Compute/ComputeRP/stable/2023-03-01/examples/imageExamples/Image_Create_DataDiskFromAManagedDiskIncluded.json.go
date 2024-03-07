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
DataDisks: compute.ImageDataDiskArray{
interface{}{
Lun: pulumi.Int(1),
ManagedDisk: &compute.SubResourceArgs{
Id: pulumi.String("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
},
},
},
OsDisk: interface{}{
ManagedDisk: &compute.SubResourceArgs{
Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
},
OsState: compute.OperatingSystemStateTypesGeneralized,
OsType: compute.OperatingSystemTypesLinux,
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
