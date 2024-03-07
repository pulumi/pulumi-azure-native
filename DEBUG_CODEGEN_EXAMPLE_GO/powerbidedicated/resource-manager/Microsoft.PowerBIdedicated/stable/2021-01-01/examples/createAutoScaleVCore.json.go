package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerbidedicated/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerbidedicated.NewAutoScaleVCore(ctx, "autoScaleVCore", &powerbidedicated.AutoScaleVCoreArgs{
			CapacityLimit:     pulumi.Int(10),
			CapacityObjectId:  pulumi.String("a28f00bd-5330-4572-88f1-fa883e074785"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestRG"),
			Sku: &powerbidedicated.AutoScaleVCoreSkuArgs{
				Capacity: pulumi.Int(0),
				Name:     pulumi.String("AutoScale"),
				Tier:     pulumi.String("AutoScale"),
			},
			Tags: pulumi.StringMap{
				"testKey": pulumi.String("testValue"),
			},
			VcoreName: pulumi.String("testvcore"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
