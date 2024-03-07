package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := workloads.NewSAPVirtualInstance(ctx, "sapVirtualInstance", &workloads.SAPVirtualInstanceArgs{
Configuration: workloads.DeploymentWithOSConfiguration{
AppLocation: "eastus",
ConfigurationType: "DeploymentWithOSConfig",
InfrastructureConfiguration: workloads.SingleServerConfiguration{
AppResourceGroup: "X00-RG",
DatabaseType: "HANA",
DbDiskConfiguration: workloads.DiskConfiguration{
DiskVolumeConfigurations: interface{}{
Backup: workloads.DiskVolumeConfiguration{
Count: 2,
SizeGB: 256,
Sku: workloads.DiskSku{
Name: "StandardSSD_LRS",
},
},
Hana/data: workloads.DiskVolumeConfiguration{
Count: 4,
SizeGB: 128,
Sku: workloads.DiskSku{
Name: "Premium_LRS",
},
},
Hana/log: workloads.DiskVolumeConfiguration{
Count: 3,
SizeGB: 128,
Sku: workloads.DiskSku{
Name: "Premium_LRS",
},
},
Hana/shared: workloads.DiskVolumeConfiguration{
Count: 1,
SizeGB: 256,
Sku: workloads.DiskSku{
Name: "StandardSSD_LRS",
},
},
Os: workloads.DiskVolumeConfiguration{
Count: 1,
SizeGB: 64,
Sku: workloads.DiskSku{
Name: "StandardSSD_LRS",
},
},
Usr/sap: workloads.DiskVolumeConfiguration{
Count: 1,
SizeGB: 128,
Sku: workloads.DiskSku{
Name: "Premium_LRS",
},
},
},
},
DeploymentType: "SingleServer",
NetworkConfiguration: workloads.NetworkConfiguration{
IsSecondaryIpEnabled: true,
},
SubnetId: "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
ImageReference: workloads.ImageReference{
Offer: "RHEL-SAP",
Publisher: "RedHat",
Sku: "84sapha-gen2",
Version: "latest",
},
OsProfile: workloads.OSProfile{
AdminUsername: "{your-username}",
OsConfiguration: workloads.LinuxConfiguration{
DisablePasswordAuthentication: true,
OsType: "Linux",
SshKeyPair: workloads.SshKeyPair{
PrivateKey: "xyz",
PublicKey: "abc",
},
},
},
VmSize: "Standard_E32ds_v4",
},
},
OsSapConfiguration: workloads.OsSapConfiguration{
SapFqdn: "xyz.test.com",
},
},
Environment: pulumi.String("NonProd"),
Location: pulumi.String("westcentralus"),
ResourceGroupName: pulumi.String("test-rg"),
SapProduct: pulumi.String("S4HANA"),
SapVirtualInstanceName: pulumi.String("X00"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
