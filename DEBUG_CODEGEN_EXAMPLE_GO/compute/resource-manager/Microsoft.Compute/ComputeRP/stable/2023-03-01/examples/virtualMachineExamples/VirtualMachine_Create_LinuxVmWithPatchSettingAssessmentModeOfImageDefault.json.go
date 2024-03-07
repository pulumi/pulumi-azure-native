package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
HardwareProfile: &compute.HardwareProfileArgs{
VmSize: pulumi.String("Standard_D2s_v3"),
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
OsProfile: compute.OSProfileResponse{
AdminPassword: pulumi.String("{your-password}"),
AdminUsername: pulumi.String("{your-username}"),
ComputerName: pulumi.String("myVM"),
LinuxConfiguration: interface{}{
PatchSettings: &compute.LinuxPatchSettingsArgs{
AssessmentMode: pulumi.String("ImageDefault"),
},
ProvisionVMAgent: pulumi.Bool(true),
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
StorageProfile: compute.StorageProfileResponse{
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("UbuntuServer"),
Publisher: pulumi.String("Canonical"),
Sku: pulumi.String("16.04-LTS"),
Version: pulumi.String("latest"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: &compute.ManagedDiskParametersArgs{
StorageAccountType: pulumi.String("Premium_LRS"),
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
