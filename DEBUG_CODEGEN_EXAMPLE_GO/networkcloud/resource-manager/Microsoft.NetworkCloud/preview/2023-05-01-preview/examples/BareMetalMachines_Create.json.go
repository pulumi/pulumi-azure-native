package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewBareMetalMachine(ctx, "bareMetalMachine", &networkcloud.BareMetalMachineArgs{
			BareMetalMachineName: pulumi.String("bareMetalMachineName"),
			BmcConnectionString:  pulumi.String("bmcconnectionstring"),
			BmcCredentials: &networkcloud.AdministrativeCredentialsArgs{
				Password: pulumi.String("{password}"),
				Username: pulumi.String("bmcuser"),
			},
			BmcMacAddress:  pulumi.String("00:00:4f:00:57:00"),
			BootMacAddress: pulumi.String("00:00:4e:00:58:af"),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:          pulumi.String("location"),
			MachineDetails:    pulumi.String("User-provided machine details."),
			MachineName:       pulumi.String("r01c001"),
			MachineSkuId:      pulumi.String("684E-3B16-399E"),
			RackId:            pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
			RackSlot:          pulumi.Float64(1),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			SerialNumber:      pulumi.String("BM1219XXX"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
