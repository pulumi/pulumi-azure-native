package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachineScaleSetVM(ctx, "virtualMachineScaleSetVM", &compute.VirtualMachineScaleSetVMArgs{
AdditionalCapabilities: &compute.AdditionalCapabilitiesArgs{
HibernationEnabled: pulumi.Bool(true),
UltraSSDEnabled: pulumi.Bool(true),
},
AvailabilitySet: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
DiagnosticsProfile: compute.DiagnosticsProfileResponse{
BootDiagnostics: &compute.BootDiagnosticsArgs{
Enabled: pulumi.Bool(true),
StorageUri: pulumi.String("aaaaaaaaaaaaa"),
},
},
HardwareProfile: compute.HardwareProfileResponse{
VmSize: pulumi.String("Basic_A0"),
VmSizeProperties: &compute.VMSizePropertiesArgs{
VCPUsAvailable: pulumi.Int(9),
VCPUsPerCore: pulumi.Int(12),
},
},
InstanceId: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
LicenseType: pulumi.String("aaaaaaaaaa"),
Location: pulumi.String("westus"),
NetworkProfile: compute.NetworkProfileResponse{
NetworkApiVersion: pulumi.String("2020-11-01"),
NetworkInterfaceConfigurations: compute.VirtualMachineNetworkInterfaceConfigurationArray{
interface{}{
DeleteOption: pulumi.String("Delete"),
DnsSettings: &compute.VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs{
DnsServers: pulumi.StringArray{
pulumi.String("aaaaaa"),
},
},
DscpConfiguration: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
EnableAcceleratedNetworking: pulumi.Bool(true),
EnableFpga: pulumi.Bool(true),
EnableIPForwarding: pulumi.Bool(true),
IpConfigurations: compute.VirtualMachineNetworkInterfaceIPConfigurationArray{
interface{}{
ApplicationGatewayBackendAddressPools: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
ApplicationSecurityGroups: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
LoadBalancerBackendAddressPools: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
Name: pulumi.String("aa"),
Primary: pulumi.Bool(true),
PrivateIPAddressVersion: pulumi.String("IPv4"),
PublicIPAddressConfiguration: interface{}{
DeleteOption: pulumi.String("Delete"),
DnsSettings: &compute.VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs{
DomainNameLabel: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaa"),
},
IdleTimeoutInMinutes: pulumi.Int(2),
IpTags: compute.VirtualMachineIpTagArray{
&compute.VirtualMachineIpTagArgs{
IpTagType: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaa"),
Tag: pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
},
},
Name: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
PublicIPAddressVersion: pulumi.String("IPv4"),
PublicIPAllocationMethod: pulumi.String("Dynamic"),
PublicIPPrefix: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
Sku: &compute.PublicIPAddressSkuArgs{
Name: pulumi.String("Basic"),
Tier: pulumi.String("Regional"),
},
},
Subnet: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
},
Name: pulumi.String("aaaaaaaaaaa"),
NetworkSecurityGroup: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
Primary: pulumi.Bool(true),
},
},
NetworkInterfaces: compute.NetworkInterfaceReferenceArray{
&compute.NetworkInterfaceReferenceArgs{
DeleteOption: pulumi.String("Delete"),
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
Primary: pulumi.Bool(true),
},
},
},
NetworkProfileConfiguration: compute.VirtualMachineScaleSetVMNetworkProfileConfigurationResponse{
NetworkInterfaceConfigurations: compute.VirtualMachineScaleSetNetworkConfigurationArray{
interface{}{
DeleteOption: pulumi.String("Delete"),
DnsSettings: &compute.VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs{
DnsServers: pulumi.StringArray{
},
},
EnableAcceleratedNetworking: pulumi.Bool(true),
EnableFpga: pulumi.Bool(true),
EnableIPForwarding: pulumi.Bool(true),
IpConfigurations: compute.VirtualMachineScaleSetIPConfigurationArray{
interface{}{
ApplicationGatewayBackendAddressPools: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
ApplicationSecurityGroups: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
LoadBalancerBackendAddressPools: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
LoadBalancerInboundNatPools: compute.SubResourceArray{
&compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
Name: pulumi.String("vmsstestnetconfig9693"),
Primary: pulumi.Bool(true),
PrivateIPAddressVersion: pulumi.String("IPv4"),
PublicIPAddressConfiguration: interface{}{
DeleteOption: pulumi.String("Delete"),
DnsSettings: &compute.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs{
DomainNameLabel: pulumi.String("aaaaaaaaaaaaaaaaaa"),
},
IdleTimeoutInMinutes: pulumi.Int(18),
IpTags: compute.VirtualMachineScaleSetIpTagArray{
&compute.VirtualMachineScaleSetIpTagArgs{
IpTagType: pulumi.String("aaaaaaa"),
Tag: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
},
},
Name: pulumi.String("aaaaaaaaaaaaaaaaaa"),
PublicIPAddressVersion: pulumi.String("IPv4"),
PublicIPPrefix: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
Sku: &compute.PublicIPAddressSkuArgs{
Name: pulumi.String("Basic"),
Tier: pulumi.String("Regional"),
},
},
Subnet: &compute.ApiEntityReferenceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"),
},
},
},
Name: pulumi.String("vmsstestnetconfig5415"),
NetworkSecurityGroup: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
Primary: pulumi.Bool(true),
},
},
},
OsProfile: compute.OSProfileResponse{
AdminPassword: pulumi.String("aaaaaaaaaaaaaaaa"),
AdminUsername: pulumi.String("Foo12"),
AllowExtensionOperations: pulumi.Bool(true),
ComputerName: pulumi.String("test000000"),
CustomData: pulumi.String("aaaa"),
LinuxConfiguration: interface{}{
DisablePasswordAuthentication: pulumi.Bool(true),
PatchSettings: &compute.LinuxPatchSettingsArgs{
AssessmentMode: pulumi.String("ImageDefault"),
PatchMode: pulumi.String("ImageDefault"),
},
ProvisionVMAgent: pulumi.Bool(true),
Ssh: interface{}{
PublicKeys: compute.SshPublicKeyTypeArray{
&compute.SshPublicKeyTypeArgs{
KeyData: pulumi.String("aaaaaa"),
Path: pulumi.String("aaa"),
},
},
},
},
RequireGuestProvisionSignal: pulumi.Bool(true),
Secrets: compute.VaultSecretGroupArray{
},
WindowsConfiguration: interface{}{
AdditionalUnattendContent: compute.AdditionalUnattendContentArray{
&compute.AdditionalUnattendContentArgs{
ComponentName: compute.ComponentNames_Microsoft_Windows_Shell_Setup,
Content: pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
PassName: compute.PassNamesOobeSystem,
SettingName: compute.SettingNamesAutoLogon,
},
},
EnableAutomaticUpdates: pulumi.Bool(true),
PatchSettings: &compute.PatchSettingsArgs{
AssessmentMode: pulumi.String("ImageDefault"),
EnableHotpatching: pulumi.Bool(true),
PatchMode: pulumi.String("Manual"),
},
ProvisionVMAgent: pulumi.Bool(true),
TimeZone: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
WinRM: interface{}{
Listeners: compute.WinRMListenerArray{
&compute.WinRMListenerArgs{
CertificateUrl: pulumi.String("aaaaaaaaaaaaaaaaaaaaaa"),
Protocol: compute.ProtocolTypesHttp,
},
},
},
},
},
Plan: &compute.PlanArgs{
Name: pulumi.String("aaaaaaaaaa"),
Product: pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
PromotionCode: pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
Publisher: pulumi.String("aaaaaaaaaaaaaaaaaaaaaa"),
},
ProtectionPolicy: &compute.VirtualMachineScaleSetVMProtectionPolicyArgs{
ProtectFromScaleIn: pulumi.Bool(true),
ProtectFromScaleSetActions: pulumi.Bool(true),
},
ResourceGroupName: pulumi.String("rgcompute"),
SecurityProfile: compute.SecurityProfileResponse{
EncryptionAtHost: pulumi.Bool(true),
SecurityType: pulumi.String("TrustedLaunch"),
UefiSettings: &compute.UefiSettingsArgs{
SecureBootEnabled: pulumi.Bool(true),
VTpmEnabled: pulumi.Bool(true),
},
},
StorageProfile: compute.StorageProfileResponse{
DataDisks: compute.DataDiskArray{
interface{}{
Caching: compute.CachingTypesNone,
CreateOption: pulumi.String("Empty"),
DeleteOption: pulumi.String("Delete"),
DetachOption: pulumi.String("ForceDetach"),
DiskSizeGB: pulumi.Int(128),
Image: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
},
Lun: pulumi.Int(1),
ManagedDisk: interface{}{
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("aaaaaaaaaaaa"),
},
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
StorageAccountType: pulumi.String("Standard_LRS"),
},
Name: pulumi.String("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
ToBeDetached: pulumi.Bool(true),
Vhd: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
},
WriteAcceleratorEnabled: pulumi.Bool(true),
},
},
ImageReference: &compute.ImageReferenceArgs{
Id: pulumi.String("a"),
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
SharedGalleryImageId: pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
Sku: pulumi.String("2012-R2-Datacenter"),
Version: pulumi.String("4.127.20180315"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesNone,
CreateOption: pulumi.String("FromImage"),
DeleteOption: pulumi.String("Delete"),
DiffDiskSettings: &compute.DiffDiskSettingsArgs{
Option: pulumi.String("Local"),
Placement: pulumi.String("CacheDisk"),
},
DiskSizeGB: pulumi.Int(127),
EncryptionSettings: interface{}{
DiskEncryptionKey: interface{}{
SecretUrl: pulumi.String("aaaaaaaa"),
SourceVault: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
Enabled: pulumi.Bool(true),
KeyEncryptionKey: interface{}{
KeyUrl: pulumi.String("aaaaaaaaaaaaaa"),
SourceVault: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
},
},
},
Image: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
},
ManagedDisk: interface{}{
DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
Id: pulumi.String("aaaaaaaaaaaa"),
},
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
StorageAccountType: pulumi.String("Standard_LRS"),
},
Name: pulumi.String("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
OsType: compute.OperatingSystemTypesWindows,
Vhd: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
},
WriteAcceleratorEnabled: pulumi.Bool(true),
},
},
Tags: nil,
UserData: pulumi.String("RXhhbXBsZSBVc2VyRGF0YQ=="),
VmScaleSetName: pulumi.String("aaaaaaaaaaaaaa"),
})
if err != nil {
return err
}
return nil
})
}
