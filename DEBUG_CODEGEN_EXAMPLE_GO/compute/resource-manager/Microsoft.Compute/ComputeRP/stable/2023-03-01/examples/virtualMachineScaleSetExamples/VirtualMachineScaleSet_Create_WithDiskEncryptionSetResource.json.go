package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachineScaleSet(ctx, "virtualMachineScaleSet", &compute.VirtualMachineScaleSetArgs{
Location: pulumi.String("westus"),
Overprovision: pulumi.Bool(true),
ResourceGroupName: pulumi.String("myResourceGroup"),
Sku: &compute.SkuArgs{
Capacity: pulumi.Float64(3),
Name: pulumi.String("Standard_DS1_v2"),
Tier: pulumi.String("Standard"),
},
UpgradePolicy: &compute.UpgradePolicyArgs{
Mode: compute.UpgradeModeManual,
},
VirtualMachineProfile: compute.VirtualMachineScaleSetVMProfileResponse{
NetworkProfile: interface{}{
NetworkInterfaceConfigurations: compute.VirtualMachineScaleSetNetworkConfigurationArray{
interface{}{
EnableIPForwarding: pulumi.Bool(true),
IpConfigurations: compute.VirtualMachineScaleSetIPConfigurationArray{
interface{}{
Name: pulumi.String("{vmss-name}"),
Subnet: &compute.ApiEntityReferenceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
},
},
},
Name: pulumi.String("{vmss-name}"),
Primary: pulumi.Bool(true),
},
},
},
OsProfile: &compute.VirtualMachineScaleSetOSProfileArgs{
AdminPassword: pulumi.String("{your-password}"),
AdminUsername: pulumi.String("{your-username}"),
ComputerNamePrefix: pulumi.String("{vmss-name}"),
},
StorageProfile: interface{}{
DataDisks: compute.VirtualMachineScaleSetDataDiskArray{
interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("Empty"),
DiskSizeGB: pulumi.Int(1023),
Lun: pulumi.Int(0),
ManagedDisk: interface{}{
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
},
StorageAccountType: pulumi.String("Standard_LRS"),
},
},
},
ImageReference: &compute.ImageReferenceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: interface{}{
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
},
StorageAccountType: pulumi.String("Standard_LRS"),
},
},
},
},
VmScaleSetName: pulumi.String("{vmss-name}"),
})
if err != nil {
return err
}
return nil
})
}
