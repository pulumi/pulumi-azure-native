package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachineScaleSet(ctx, "virtualMachineScaleSet", &compute.VirtualMachineScaleSetArgs{
Location: pulumi.String("eastus2euap"),
Overprovision: pulumi.Bool(true),
ResourceGroupName: pulumi.String("myResourceGroup"),
Sku: &compute.SkuArgs{
Capacity: pulumi.Float64(3),
Name: pulumi.String("Standard_A1"),
Tier: pulumi.String("Standard"),
},
UpgradePolicy: compute.UpgradePolicyResponse{
AutomaticOSUpgradePolicy: &compute.AutomaticOSUpgradePolicyArgs{
EnableAutomaticOSUpgrade: pulumi.Bool(true),
},
Mode: compute.UpgradeModeAutomatic,
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
ServiceArtifactReference: &compute.ServiceArtifactReferenceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/serviceArtifacts/serviceArtifactName/vmArtifactsProfiles/vmArtifactsProfilesName"),
},
StorageProfile: interface{}{
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("2022-Datacenter"),
Version: pulumi.String("latest"),
},
OsDisk: &compute.VirtualMachineScaleSetOSDiskArgs{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
Name: pulumi.String("osDisk"),
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
