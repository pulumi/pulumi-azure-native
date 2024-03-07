package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datacatalog/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datacatalog.NewADCCatalog(ctx, "adcCatalog", &datacatalog.ADCCatalogArgs{
			Admins: []datacatalog.PrincipalsArgs{
				{
					ObjectId: pulumi.String("99999999-9999-9999-999999999999"),
					Upn:      pulumi.String("myupn@microsoft.com"),
				},
			},
			CatalogName:                   pulumi.String("exampleCatalog"),
			EnableAutomaticUnitAdjustment: pulumi.Bool(false),
			Location:                      pulumi.String("North US"),
			ResourceGroupName:             pulumi.String("exampleResourceGroup"),
			Sku:                           pulumi.String("Standard"),
			Tags: pulumi.StringMap{
				"mykey":  pulumi.String("myvalue"),
				"mykey2": pulumi.String("myvalue2"),
			},
			Units: pulumi.Int(1),
			Users: []datacatalog.PrincipalsArgs{
				{
					ObjectId: pulumi.String("99999999-9999-9999-999999999999"),
					Upn:      pulumi.String("myupn@microsoft.com"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
