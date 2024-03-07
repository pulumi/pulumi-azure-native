package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerbidedicated/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerbidedicated.NewCapacityDetails(ctx, "capacityDetails", &powerbidedicated.CapacityDetailsArgs{
			Administration: &powerbidedicated.DedicatedCapacityAdministratorsArgs{
				Members: pulumi.StringArray{
					pulumi.String("azsdktest@microsoft.com"),
					pulumi.String("azsdktest2@microsoft.com"),
				},
			},
			DedicatedCapacityName: pulumi.String("azsdktest"),
			Location:              pulumi.String("West US"),
			ResourceGroupName:     pulumi.String("TestRG"),
			Sku: &powerbidedicated.CapacitySkuArgs{
				Name: pulumi.String("A1"),
				Tier: pulumi.String("PBIE_Azure"),
			},
			Tags: pulumi.StringMap{
				"testKey": pulumi.String("testValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
