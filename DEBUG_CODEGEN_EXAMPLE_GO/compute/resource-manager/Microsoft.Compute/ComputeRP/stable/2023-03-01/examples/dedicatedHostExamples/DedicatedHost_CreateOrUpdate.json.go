package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDedicatedHost(ctx, "dedicatedHost", &compute.DedicatedHostArgs{
			HostGroupName:       pulumi.String("myDedicatedHostGroup"),
			HostName:            pulumi.String("myDedicatedHost"),
			Location:            pulumi.String("westus"),
			PlatformFaultDomain: pulumi.Int(1),
			ResourceGroupName:   pulumi.String("myResourceGroup"),
			Sku: &compute.SkuArgs{
				Name: pulumi.String("DSv3-Type1"),
			},
			Tags: pulumi.StringMap{
				"department": pulumi.String("HR"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
