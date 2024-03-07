package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azureactivedirectory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azureactivedirectory.NewB2CTenant(ctx, "b2cTenant", &azureactivedirectory.B2CTenantArgs{
			CountryCode:       pulumi.String("US"),
			DisplayName:       pulumi.String("Contoso"),
			Location:          pulumi.String("United States"),
			ResourceGroupName: pulumi.String("contosoResourceGroup"),
			ResourceName:      pulumi.String("contoso.onmicrosoft.com"),
			Sku: &azureactivedirectory.B2CResourceSKUArgs{
				Name: pulumi.String("Standard"),
				Tier: pulumi.String("A0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
