package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewVirtualMachineInstance(ctx, "virtualMachineInstance", &azurestackhci.VirtualMachineInstanceArgs{
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			HardwareProfile: &azurestackhci.VirtualMachineInstancePropertiesHardwareProfileArgs{
				VmSize: pulumi.String("Default"),
			},
			NetworkProfile: &azurestackhci.VirtualMachineInstancePropertiesNetworkProfileArgs{
				NetworkInterfaces: azurestackhci.VirtualMachineInstancePropertiesNetworkInterfacesArray{
					&azurestackhci.VirtualMachineInstancePropertiesNetworkInterfacesArgs{
						Id: pulumi.String("test-nic"),
					},
				},
			},
			OsProfile: &azurestackhci.VirtualMachineInstancePropertiesOsProfileArgs{
				AdminPassword: pulumi.String("password"),
				AdminUsername: pulumi.String("localadmin"),
				ComputerName:  pulumi.String("luamaster"),
				WindowsConfiguration: &azurestackhci.VirtualMachineInstancePropertiesWindowsConfigurationArgs{
					ProvisionVMConfigAgent: pulumi.Bool(true),
				},
			},
			ResourceUri: pulumi.String("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM"),
			SecurityProfile: &azurestackhci.VirtualMachineInstancePropertiesSecurityProfileArgs{
				EnableTPM: pulumi.Bool(true),
				UefiSettings: &azurestackhci.VirtualMachineInstancePropertiesUefiSettingsArgs{
					SecureBootEnabled: pulumi.Bool(true),
				},
			},
			StorageProfile: &azurestackhci.VirtualMachineInstancePropertiesStorageProfileArgs{
				ImageReference: &azurestackhci.VirtualMachineInstancePropertiesImageReferenceArgs{
					Id: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
				},
				VmConfigStoragePathId: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
