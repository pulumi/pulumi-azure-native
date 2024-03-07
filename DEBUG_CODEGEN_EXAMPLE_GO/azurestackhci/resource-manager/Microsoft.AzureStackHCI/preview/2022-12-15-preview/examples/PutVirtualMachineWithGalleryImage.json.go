package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewVirtualMachine(ctx, "virtualMachine", &azurestackhci.VirtualMachineArgs{
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			HardwareProfile: &azurestackhci.VirtualMachinePropertiesHardwareProfileArgs{
				VmSize: pulumi.String("Default"),
			},
			Location: pulumi.String("West US2"),
			NetworkProfile: &azurestackhci.VirtualMachinePropertiesNetworkProfileArgs{
				NetworkInterfaces: azurestackhci.VirtualMachinePropertiesNetworkInterfacesArray{
					&azurestackhci.VirtualMachinePropertiesNetworkInterfacesArgs{
						Id: pulumi.String("test-nic"),
					},
				},
			},
			OsProfile: &azurestackhci.VirtualMachinePropertiesOsProfileArgs{
				AdminPassword: pulumi.String("password"),
				AdminUsername: pulumi.String("localadmin"),
				ComputerName:  pulumi.String("luamaster"),
			},
			ResourceGroupName: pulumi.String("test-rg"),
			SecurityProfile: &azurestackhci.VirtualMachinePropertiesSecurityProfileArgs{
				EnableTPM: pulumi.Bool(true),
				UefiSettings: &azurestackhci.VirtualMachinePropertiesUefiSettingsArgs{
					SecureBootEnabled: pulumi.Bool(true),
				},
			},
			StorageProfile: &azurestackhci.VirtualMachinePropertiesStorageProfileArgs{
				ImageReference: &azurestackhci.VirtualMachinePropertiesImageReferenceArgs{
					Id: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
				},
				VmConfigStoragePathId: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
			},
			VirtualMachineName: pulumi.String("test-vm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
