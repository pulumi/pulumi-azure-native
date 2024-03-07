package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/analysisservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := analysisservices.NewServerDetails(ctx, "serverDetails", &analysisservices.ServerDetailsArgs{
			AsAdministrators: &analysisservices.ServerAdministratorsArgs{
				Members: pulumi.StringArray{
					pulumi.String("azsdktest@microsoft.com"),
					pulumi.String("azsdktest2@microsoft.com"),
				},
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestRG"),
			ServerName:        pulumi.String("azsdktest"),
			Sku: &analysisservices.ResourceSkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("S1"),
				Tier:     pulumi.String("Standard"),
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
