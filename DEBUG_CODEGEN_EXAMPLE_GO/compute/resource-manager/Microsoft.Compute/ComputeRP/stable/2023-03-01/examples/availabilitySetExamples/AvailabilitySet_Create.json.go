package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewAvailabilitySet(ctx, "availabilitySet", &compute.AvailabilitySetArgs{
			AvailabilitySetName:       pulumi.String("myAvailabilitySet"),
			Location:                  pulumi.String("westus"),
			PlatformFaultDomainCount:  pulumi.Int(2),
			PlatformUpdateDomainCount: pulumi.Int(20),
			ResourceGroupName:         pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
