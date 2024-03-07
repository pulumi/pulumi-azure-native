package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hardwaresecuritymodules/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hardwaresecuritymodules.NewDedicatedHsm(ctx, "dedicatedHsm", &hardwaresecuritymodules.DedicatedHsmArgs{
			Location: pulumi.String("westus"),
			Name:     pulumi.String("hsm1"),
			NetworkProfile: hardwaresecuritymodules.NetworkProfileResponse{
				NetworkInterfaces: hardwaresecuritymodules.NetworkInterfaceArray{
					&hardwaresecuritymodules.NetworkInterfaceArgs{
						PrivateIpAddress: pulumi.String("1.0.0.1"),
					},
				},
				Subnet: &hardwaresecuritymodules.ApiEntityReferenceArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
				},
			},
			ResourceGroupName: pulumi.String("hsm-group"),
			Sku: &hardwaresecuritymodules.SkuArgs{
				Name: pulumi.String("payShield10K_LMK1_CPS60"),
			},
			StampId: pulumi.String("stamp01"),
			Tags: pulumi.StringMap{
				"Dept":        pulumi.String("hsm"),
				"Environment": pulumi.String("dogfood"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
