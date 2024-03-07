package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
			HardwareProfile: &compute.HardwareProfileArgs{
				VmSize: pulumi.String("Standard_D1_v2"),
			},
			Location: pulumi.String("westus"),
			NetworkProfile: &compute.NetworkProfileArgs{
				NetworkApiVersion: pulumi.String("2020-11-01"),
				NetworkInterfaceConfigurations: compute.VirtualMachineNetworkInterfaceConfigurationArray{
					&compute.VirtualMachineNetworkInterfaceConfigurationArgs{
						DeleteOption: pulumi.String("Delete"),
						IpConfigurations: compute.VirtualMachineNetworkInterfaceIPConfigurationArray{
							&compute.VirtualMachineNetworkInterfaceIPConfigurationArgs{
								Name:    pulumi.String("{ip-config-name}"),
								Primary: pulumi.Bool(true),
								PublicIPAddressConfiguration: &compute.VirtualMachinePublicIPAddressConfigurationArgs{
									DeleteOption:             pulumi.String("Detach"),
									Name:                     pulumi.String("{publicIP-config-name}"),
									PublicIPAllocationMethod: pulumi.String("Static"),
									Sku: &compute.PublicIPAddressSkuArgs{
										Name: pulumi.String("Basic"),
										Tier: pulumi.String("Global"),
									},
								},
							},
						},
						Name:    pulumi.String("{nic-config-name}"),
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
			VmName: pulumi.String("myVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
