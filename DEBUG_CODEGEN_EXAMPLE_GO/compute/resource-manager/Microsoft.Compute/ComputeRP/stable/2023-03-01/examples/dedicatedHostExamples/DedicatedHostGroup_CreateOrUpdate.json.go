package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDedicatedHostGroup(ctx, "dedicatedHostGroup", &compute.DedicatedHostGroupArgs{
			HostGroupName:             pulumi.String("myDedicatedHostGroup"),
			Location:                  pulumi.String("westus"),
			PlatformFaultDomainCount:  pulumi.Int(3),
			ResourceGroupName:         pulumi.String("myResourceGroup"),
			SupportAutomaticPlacement: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"department": pulumi.String("finance"),
			},
			Zones: pulumi.StringArray{
				pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
