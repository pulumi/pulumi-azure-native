package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
ApplicationProfile: compute.ApplicationProfileResponse{
GalleryApplications: compute.VMGalleryApplicationArray{
&compute.VMGalleryApplicationArgs{
ConfigurationReference: pulumi.String("https://mystorageaccount.blob.core.windows.net/configurations/settings.config"),
EnableAutomaticUpgrade: pulumi.Bool(false),
Order: pulumi.Int(1),
PackageReferenceId: pulumi.String("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0"),
Tags: pulumi.String("myTag1"),
TreatFailureAsDeploymentFailure: pulumi.Bool(false),
},
&compute.VMGalleryApplicationArgs{
PackageReferenceId: pulumi.String("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1"),
},
},
},
HardwareProfile: &compute.HardwareProfileArgs{
VmSize: pulumi.String("Standard_D1_v2"),
},
Location: pulumi.String("westus"),
NetworkProfile: compute.NetworkProfileResponse{
NetworkInterfaces: compute.NetworkInterfaceReferenceArray{
&compute.NetworkInterfaceReferenceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
Primary: pulumi.Bool(true),
},
},
},
OsProfile: &compute.OSProfileArgs{
AdminPassword: pulumi.String("{your-password}"),
AdminUsername: pulumi.String("{your-username}"),
ComputerName: pulumi.String("myVM"),
},
ResourceGroupName: pulumi.String("myResourceGroup"),
StorageProfile: compute.StorageProfileResponse{
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("{image_offer}"),
Publisher: pulumi.String("{image_publisher}"),
Sku: pulumi.String("{image_sku}"),
Version: pulumi.String("latest"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: &compute.ManagedDiskParametersArgs{
StorageAccountType: pulumi.String("Standard_LRS"),
},
Name: pulumi.String("myVMosdisk"),
},
},
VmName: pulumi.String("myVM"),
})
if err != nil {
return err
}
return nil
})
}
