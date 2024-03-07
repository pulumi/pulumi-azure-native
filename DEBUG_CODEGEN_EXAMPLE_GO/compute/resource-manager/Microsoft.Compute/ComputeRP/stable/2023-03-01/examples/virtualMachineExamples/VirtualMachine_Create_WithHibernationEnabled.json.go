package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
AdditionalCapabilities: &compute.AdditionalCapabilitiesArgs{
HibernationEnabled: pulumi.Bool(true),
},
DiagnosticsProfile: compute.DiagnosticsProfileResponse{
BootDiagnostics: &compute.BootDiagnosticsArgs{
Enabled: pulumi.Bool(true),
StorageUri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net"),
},
},
HardwareProfile: &compute.HardwareProfileArgs{
VmSize: pulumi.String("Standard_D2s_v3"),
},
Location: pulumi.String("eastus2euap"),
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
ComputerName: pulumi.String("{vm-name}"),
},
ResourceGroupName: pulumi.String("myResourceGroup"),
StorageProfile: compute.StorageProfileResponse{
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("2019-Datacenter"),
Version: pulumi.String("latest"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
ManagedDisk: &compute.ManagedDiskParametersArgs{
StorageAccountType: pulumi.String("Standard_LRS"),
},
Name: pulumi.String("vmOSdisk"),
},
},
VmName: pulumi.String("{vm-name}"),
})
if err != nil {
return err
}
return nil
})
}
