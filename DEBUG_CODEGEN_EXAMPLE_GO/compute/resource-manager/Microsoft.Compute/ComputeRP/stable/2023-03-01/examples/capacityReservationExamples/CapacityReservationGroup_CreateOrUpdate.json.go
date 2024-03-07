package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewCapacityReservationGroup(ctx, "capacityReservationGroup", &compute.CapacityReservationGroupArgs{
			CapacityReservationGroupName: pulumi.String("myCapacityReservationGroup"),
			Location:                     pulumi.String("westus"),
			ResourceGroupName:            pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"department": pulumi.String("finance"),
			},
			Zones: pulumi.StringArray{
				pulumi.String("1"),
				pulumi.String("2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
