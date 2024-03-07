package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
HardwareProfile: &compute.HardwareProfileArgs{
VmSize: pulumi.String("Standard_D2_v2"),
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
DataDisks: compute.DataDiskArray{
interface{}{
CreateOption: pulumi.String("Empty"),
DiskSizeGB: pulumi.Int(1023),
Lun: pulumi.Int(0),
Vhd: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd"),
},
},
interface{}{
CreateOption: pulumi.String("Empty"),
DiskSizeGB: pulumi.Int(1023),
Lun: pulumi.Int(1),
Vhd: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd"),
},
},
},
ImageReference: &compute.ImageReferenceArgs{
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("2016-Datacenter"),
Version: pulumi.String("latest"),
},
OsDisk: interface{}{
Caching: compute.CachingTypesReadWrite,
CreateOption: pulumi.String("FromImage"),
Name: pulumi.String("myVMosdisk"),
Vhd: &compute.VirtualHardDiskArgs{
Uri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
},
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
