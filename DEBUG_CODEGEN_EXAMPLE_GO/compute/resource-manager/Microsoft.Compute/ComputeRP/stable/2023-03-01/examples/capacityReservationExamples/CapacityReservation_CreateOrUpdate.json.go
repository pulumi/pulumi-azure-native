package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewCapacityReservation(ctx, "capacityReservation", &compute.CapacityReservationArgs{
			CapacityReservationGroupName: pulumi.String("myCapacityReservationGroup"),
			CapacityReservationName:      pulumi.String("myCapacityReservation"),
			Location:                     pulumi.String("westus"),
			ResourceGroupName:            pulumi.String("myResourceGroup"),
			Sku: &compute.SkuArgs{
				Capacity: pulumi.Float64(4),
				Name:     pulumi.String("Standard_DS1_v2"),
			},
			Tags: pulumi.StringMap{
				"department": pulumi.String("HR"),
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
