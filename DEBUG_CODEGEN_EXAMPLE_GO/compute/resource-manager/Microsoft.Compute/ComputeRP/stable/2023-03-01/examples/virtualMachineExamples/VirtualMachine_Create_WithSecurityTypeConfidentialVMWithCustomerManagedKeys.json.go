package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
HardwareProfile: &compute.HardwareProfileArgs{
VmSize: pulumi.String("Standard_DC2as_v5"),
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
SecurityProfile: compute.SecurityProfileResponse{
SecurityType: pulumi.String("ConfidentialVM"),
UefiSettings: &compute.UefiSettingsArgs{
SecureBootEnabled: pulumi.Bool(true),
VTpmEnabled: pulumi.Bool(true),
},
},
StorageProfile: compute.StorageProfileResponse{
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("2019-datacenter-cvm"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("windows-cvm"),
Version: pulumi.String("17763.2183.2109130127"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadOnly,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: interface{}{
SecurityProfile: interface{}{
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
},
SecurityEncryptionType: pulumi.String("DiskWithVMGuestState"),
},
StorageAccountType: pulumi.String("StandardSSD_LRS"),
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
