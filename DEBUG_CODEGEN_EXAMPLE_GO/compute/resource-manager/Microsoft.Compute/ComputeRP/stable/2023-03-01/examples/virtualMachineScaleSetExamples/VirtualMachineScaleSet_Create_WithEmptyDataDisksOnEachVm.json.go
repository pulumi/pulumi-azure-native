package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSet(ctx, "virtualMachineScaleSet", &compute.VirtualMachineScaleSetArgs{
			Location:          pulumi.String("westus"),
			Overprovision:     pulumi.Bool(true),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &compute.SkuArgs{
				Capacity: pulumi.Float64(3),
				Name:     pulumi.String("Standard_D2_v2"),
				Tier:     pulumi.String("Standard"),
			},
			UpgradePolicy: &compute.UpgradePolicyArgs{
				Mode: compute.UpgradeModeManual,
			},
			VirtualMachineProfile: &compute.VirtualMachineScaleSetVMProfileArgs{
				NetworkProfile: &compute.VirtualMachineScaleSetNetworkProfileArgs{
					NetworkInterfaceConfigurations: compute.VirtualMachineScaleSetNetworkConfigurationArray{
						&compute.VirtualMachineScaleSetNetworkConfigurationArgs{
							EnableIPForwarding: pulumi.Bool(true),
							IpConfigurations: compute.VirtualMachineScaleSetIPConfigurationArray{
								&compute.VirtualMachineScaleSetIPConfigurationArgs{
									Name: pulumi.String("{vmss-name}"),
									Subnet: &compute.ApiEntityReferenceArgs{
										Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
									},
								},
							},
							Name:    pulumi.String("{vmss-name}"),
							Primary: pulumi.Bool(true),
						},
					},
				},
				OsProfile: &compute.VirtualMachineScaleSetOSProfileArgs{
					AdminPassword:      pulumi.String("{your-password}"),
					AdminUsername:      pulumi.String("{your-username}"),
					ComputerNamePrefix: pulumi.String("{vmss-name}"),
				},
				StorageProfile: &compute.VirtualMachineScaleSetStorageProfileArgs{
					DataDisks: compute.VirtualMachineScaleSetDataDiskArray{
						&compute.VirtualMachineScaleSetDataDiskArgs{
							CreateOption: pulumi.String("Empty"),
							DiskSizeGB:   pulumi.Int(1023),
							Lun:          pulumi.Int(0),
						},
						&compute.VirtualMachineScaleSetDataDiskArgs{
							CreateOption: pulumi.String("Empty"),
							DiskSizeGB:   pulumi.Int(1023),
							Lun:          pulumi.Int(1),
						},
					},
					ImageReference: &compute.ImageReferenceArgs{
						Offer:     pulumi.String("WindowsServer"),
						Publisher: pulumi.String("MicrosoftWindowsServer"),
						Sku:       pulumi.String("2016-Datacenter"),
						Version:   pulumi.String("latest"),
					},
					OsDisk: &compute.VirtualMachineScaleSetOSDiskArgs{
						Caching:      compute.CachingTypesReadWrite,
						CreateOption: pulumi.String("FromImage"),
						DiskSizeGB:   pulumi.Int(512),
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
