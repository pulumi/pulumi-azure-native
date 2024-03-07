package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewLogicalNetwork(ctx, "logicalNetwork", &azurestackhci.LogicalNetworkArgs{
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:           pulumi.String("West US2"),
			LogicalNetworkName: pulumi.String("test-lnet"),
			ResourceGroupName:  pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
