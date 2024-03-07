package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
			CapacityReservation: &compute.CapacityReservationProfileArgs{
				CapacityReservationGroup: &compute.SubResourceArgs{
					Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/{crgName}"),
				},
			},
			HardwareProfile: &compute.HardwareProfileArgs{
				VmSize: pulumi.String("Standard_DS1_v2"),
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
			Plan: &compute.PlanArgs{
				Name:      pulumi.String("windows2016"),
				Product:   pulumi.String("windows-data-science-vm"),
				Publisher: pulumi.String("microsoft-ads"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			StorageProfile: &compute.StorageProfileArgs{
				ImageReference: &compute.ImageReferenceArgs{
					Offer:     pulumi.String("windows-data-science-vm"),
					Publisher: pulumi.String("microsoft-ads"),
					Sku:       pulumi.String("windows2016"),
					Version:   pulumi.String("latest"),
				},
				OsDisk: &compute.OSDiskArgs{
					Caching:      compute.CachingTypesReadOnly,
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
