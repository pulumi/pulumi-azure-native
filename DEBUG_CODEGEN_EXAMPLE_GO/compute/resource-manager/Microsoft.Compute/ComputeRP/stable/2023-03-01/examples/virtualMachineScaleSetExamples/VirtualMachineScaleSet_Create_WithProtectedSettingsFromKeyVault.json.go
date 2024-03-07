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
DiagnosticsProfile: interface{}{
BootDiagnostics: &compute.BootDiagnosticsArgs{
Enabled: pulumi.Bool(true),
StorageUri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net"),
},
},
ExtensionProfile: interface{}{
Extensions: compute.VirtualMachineScaleSetExtensionTypeArray{
interface{}{
AutoUpgradeMinorVersion: pulumi.Bool(false),
Name: pulumi.String("{extension-name}"),
ProtectedSettingsFromKeyVault: interface{}{
SecretUrl: pulumi.String("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
SourceVault: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
},
},
Publisher: pulumi.String("{extension-Publisher}"),
Settings: nil,
Type: pulumi.String("{extension-Type}"),
TypeHandlerVersion: pulumi.String("{handler-version}"),
},
},
},
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
},
VmScaleSetName: pulumi.String("{vmss-name}"),
})
if err != nil {
return err
}
return nil
})
}
