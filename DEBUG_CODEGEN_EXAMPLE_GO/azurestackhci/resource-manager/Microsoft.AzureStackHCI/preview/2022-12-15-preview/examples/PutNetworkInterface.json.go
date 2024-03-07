package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewNetworkInterface(ctx, "networkInterface", &azurestackhci.NetworkInterfaceArgs{
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			IpConfigurations: azurestackhci.IPConfigurationArray{
				&azurestackhci.IPConfigurationArgs{
					Name: pulumi.String("ipconfig-sample"),
					Properties: &azurestackhci.IPConfigurationPropertiesArgs{
						Subnet: &azurestackhci.IPConfigurationSubnetArgs{
							Id: pulumi.String("test-vnet"),
						},
					},
				},
			},
			Location:             pulumi.String("West US2"),
			NetworkInterfaceName: pulumi.String("test-nic"),
			ResourceGroupName:    pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
