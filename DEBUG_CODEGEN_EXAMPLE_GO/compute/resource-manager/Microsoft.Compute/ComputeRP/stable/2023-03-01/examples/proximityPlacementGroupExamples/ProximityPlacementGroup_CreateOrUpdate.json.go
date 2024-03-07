package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewProximityPlacementGroup(ctx, "proximityPlacementGroup", &compute.ProximityPlacementGroupArgs{
			Intent: &compute.ProximityPlacementGroupPropertiesIntentArgs{
				VmSizes: pulumi.StringArray{
					pulumi.String("Basic_A0"),
					pulumi.String("Basic_A2"),
				},
			},
			Location:                    pulumi.String("westus"),
			ProximityPlacementGroupName: pulumi.String("myProximityPlacementGroup"),
			ProximityPlacementGroupType: pulumi.String("Standard"),
			ResourceGroupName:           pulumi.String("myResourceGroup"),
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
