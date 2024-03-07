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
Name: pulumi.String("Standard_DC2as_v5"),
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
SecurityProfile: interface{}{
SecurityType: pulumi.String("ConfidentialVM"),
UefiSettings: &compute.UefiSettingsArgs{
SecureBootEnabled: pulumi.Bool(true),
VTpmEnabled: pulumi.Bool(true),
},
},
StorageProfile: interface{}{
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
SecurityProfile: &compute.VMDiskSecurityProfileArgs{
SecurityEncryptionType: pulumi.String("VMGuestStateOnly"),
},
StorageAccountType: pulumi.String("StandardSSD_LRS"),
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
