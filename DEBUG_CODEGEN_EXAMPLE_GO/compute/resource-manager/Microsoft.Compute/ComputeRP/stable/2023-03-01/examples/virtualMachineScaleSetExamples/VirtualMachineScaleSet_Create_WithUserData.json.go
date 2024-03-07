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
Name: pulumi.String("Standard_D1_v2"),
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
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("2016-Datacenter"),
Version: pulumi.String("latest"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: &compute.VirtualMachineScaleSetManagedDiskParametersArgs{
StorageAccountType: pulumi.String("Standard_LRS"),
},
},
},
UserData: pulumi.String("RXhhbXBsZSBVc2VyRGF0YQ=="),
},
VmScaleSetName: pulumi.String("{vmss-name}"),
})
if err != nil {
return err
}
return nil
})
}
