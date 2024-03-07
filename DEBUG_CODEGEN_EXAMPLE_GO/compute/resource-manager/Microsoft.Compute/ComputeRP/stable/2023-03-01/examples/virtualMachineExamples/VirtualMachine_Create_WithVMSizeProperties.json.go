package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
			DiagnosticsProfile: &compute.DiagnosticsProfileArgs{
				BootDiagnostics: &compute.BootDiagnosticsArgs{
					Enabled:    pulumi.Bool(true),
					StorageUri: pulumi.String("http://{existing-storage-account-name}.blob.core.windows.net"),
				},
			},
			HardwareProfile: &compute.HardwareProfileArgs{
				VmSize: pulumi.String("Standard_D4_v3"),
				VmSizeProperties: &compute.VMSizePropertiesArgs{
					VCPUsAvailable: pulumi.Int(1),
					VCPUsPerCore:   pulumi.Int(1),
				},
			},
			Location: pulumi.String("westus"),
			NetworkProfile: &compute.NetworkProfileArgs{
				NetworkInterfaces: compute.NetworkInterfaceReferenceArray{
					&compute.NetworkInterfaceReferenceArgs{
						Id:      pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Primary: pulumi.Bool(true),
					},
				},
			},
			OsProfile: &compute.OSProfileArgs{
				AdminPassword: pulumi.String("{your-password}"),
				AdminUsername: pulumi.String("{your-username}"),
				ComputerName:  pulumi.String("myVM"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			StorageProfile: &compute.StorageProfileArgs{
				ImageReference: &compute.ImageReferenceArgs{
					Offer:     pulumi.String("WindowsServer"),
					Publisher: pulumi.String("MicrosoftWindowsServer"),
					Sku:       pulumi.String("2016-Datacenter"),
					Version:   pulumi.String("latest"),
				},
				OsDisk: &compute.OSDiskArgs{
					Caching:      compute.CachingTypesReadWrite,
					CreateOption: pulumi.String("FromImage"),
					ManagedDisk: &compute.ManagedDiskParametersArgs{
						StorageAccountType: pulumi.String("Standard_LRS"),
					},
					Name: pulumi.String("myVMosdisk"),
				},
			},
			UserData: pulumi.String("U29tZSBDdXN0b20gRGF0YQ=="),
			VmName:   pulumi.String("myVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
