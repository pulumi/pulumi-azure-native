package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkanalytics.NewDataProduct(ctx, "dataProduct", &networkanalytics.DataProductArgs{
			DataProductName:   pulumi.String("dataproduct01"),
			Location:          pulumi.String("eastus"),
			MajorVersion:      pulumi.String("1.0.0"),
			Product:           pulumi.String("MCC"),
			Publisher:         pulumi.String("Microsoft"),
			ResourceGroupName: pulumi.String("aoiresourceGroupName"),
			Tags: pulumi.StringMap{
				"userSpecifiedKeyName": pulumi.String("userSpecifiedKeyValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
