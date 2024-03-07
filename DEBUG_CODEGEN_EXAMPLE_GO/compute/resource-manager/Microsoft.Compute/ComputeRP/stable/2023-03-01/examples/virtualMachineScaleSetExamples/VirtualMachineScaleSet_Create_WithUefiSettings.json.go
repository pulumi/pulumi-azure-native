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
				Name:     pulumi.String("Standard_D2s_v3"),
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
				SecurityProfile: &compute.SecurityProfileArgs{
					SecurityType: pulumi.String("TrustedLaunch"),
					UefiSettings: &compute.UefiSettingsArgs{
						SecureBootEnabled: pulumi.Bool(true),
						VTpmEnabled:       pulumi.Bool(true),
					},
				},
				StorageProfile: &compute.VirtualMachineScaleSetStorageProfileArgs{
					ImageReference: &compute.ImageReferenceArgs{
						Offer:     pulumi.String("windowsserver-gen2preview-preview"),
						Publisher: pulumi.String("MicrosoftWindowsServer"),
						Sku:       pulumi.String("windows10-tvm"),
						Version:   pulumi.String("18363.592.2001092016"),
					},
					OsDisk: &compute.VirtualMachineScaleSetOSDiskArgs{
						Caching:      compute.CachingTypesReadOnly,
						CreateOption: pulumi.String("FromImage"),
						ManagedDisk: &compute.VirtualMachineScaleSetManagedDiskParametersArgs{
							StorageAccountType: pulumi.String("StandardSSD_LRS"),
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
