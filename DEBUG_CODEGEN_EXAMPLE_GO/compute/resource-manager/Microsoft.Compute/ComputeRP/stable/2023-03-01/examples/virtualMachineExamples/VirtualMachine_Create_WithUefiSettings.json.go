package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
			HardwareProfile: &compute.HardwareProfileArgs{
				VmSize: pulumi.String("Standard_D2s_v3"),
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
			SecurityProfile: &compute.SecurityProfileArgs{
				SecurityType: pulumi.String("TrustedLaunch"),
				UefiSettings: &compute.UefiSettingsArgs{
					SecureBootEnabled: pulumi.Bool(true),
					VTpmEnabled:       pulumi.Bool(true),
				},
			},
			StorageProfile: &compute.StorageProfileArgs{
				ImageReference: &compute.ImageReferenceArgs{
					Offer:     pulumi.String("windowsserver-gen2preview-preview"),
					Publisher: pulumi.String("MicrosoftWindowsServer"),
					Sku:       pulumi.String("windows10-tvm"),
					Version:   pulumi.String("18363.592.2001092016"),
				},
				OsDisk: &compute.OSDiskArgs{
					Caching:      compute.CachingTypesReadOnly,
					CreateOption: pulumi.String("FromImage"),
					ManagedDisk: &compute.ManagedDiskParametersArgs{
						StorageAccountType: pulumi.String("StandardSSD_LRS"),
					},
					Name: pulumi.String("myVMosdisk"),
				},
			},
			VmName: pulumi.String("myVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
