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
OsState: compute.OperatingSystemStateTypesGeneralized,
OsType: compute.OperatingSystemTypesLinux,
Snapshot: &compute.SubResourceArgs{
Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
},
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
