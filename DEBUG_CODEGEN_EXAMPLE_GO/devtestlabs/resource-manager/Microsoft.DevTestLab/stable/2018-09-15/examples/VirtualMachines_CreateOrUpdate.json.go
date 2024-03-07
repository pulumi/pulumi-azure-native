package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewVirtualMachine(ctx, "virtualMachine", &devtestlab.VirtualMachineArgs{
			AllowClaim:              pulumi.Bool(true),
			DisallowPublicIpAddress: pulumi.Bool(true),
			GalleryImageReference: &devtestlab.GalleryImageReferenceArgs{
				Offer:     pulumi.String("UbuntuServer"),
				OsType:    pulumi.String("Linux"),
				Publisher: pulumi.String("Canonical"),
				Sku:       pulumi.String("16.04-LTS"),
				Version:   pulumi.String("Latest"),
			},
			LabName:             pulumi.String("{labName}"),
			LabSubnetName:       pulumi.String("{virtualNetworkName}Subnet"),
			LabVirtualNetworkId: pulumi.String("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
			Location:            pulumi.String("{location}"),
			Name:                pulumi.String("{vmName}"),
			Password:            pulumi.String("{userPassword}"),
			ResourceGroupName:   pulumi.String("resourceGroupName"),
			Size:                pulumi.String("Standard_A2_v2"),
			StorageType:         pulumi.String("Standard"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			UserName: pulumi.String("{userName}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
