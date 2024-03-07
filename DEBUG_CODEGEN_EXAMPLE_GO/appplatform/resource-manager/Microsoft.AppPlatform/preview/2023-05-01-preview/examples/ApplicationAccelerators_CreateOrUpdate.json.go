package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewApplicationAccelerator(ctx, "applicationAccelerator", &appplatform.ApplicationAcceleratorArgs{
			ApplicationAcceleratorName: pulumi.String("default"),
			ResourceGroupName:          pulumi.String("myResourceGroup"),
			ServiceName:                pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("E0"),
				Tier:     pulumi.String("Enterprise"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
