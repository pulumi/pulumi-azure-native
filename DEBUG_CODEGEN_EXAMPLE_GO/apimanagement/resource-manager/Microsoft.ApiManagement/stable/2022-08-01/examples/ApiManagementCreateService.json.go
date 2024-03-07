package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
			Location:          pulumi.String("South Central US"),
			PublisherEmail:    pulumi.String("foo@contoso.com"),
			PublisherName:     pulumi.String("foo"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Developer"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Contoso"),
				"Test": pulumi.String("User"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
