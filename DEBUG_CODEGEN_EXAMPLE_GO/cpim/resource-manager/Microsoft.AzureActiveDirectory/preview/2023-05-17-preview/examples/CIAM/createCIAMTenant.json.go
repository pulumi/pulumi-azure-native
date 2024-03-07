package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azureactivedirectory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azureactivedirectory.NewCIAMTenant(ctx, "ciamTenant", &azureactivedirectory.CIAMTenantArgs{
			CreateTenantProperties: &azureactivedirectory.CreateCIAMTenantPropertiesArgs{
				CountryCode: pulumi.String("US"),
				DisplayName: pulumi.String("Contoso"),
			},
			Location:          pulumi.String("United States"),
			ResourceGroupName: pulumi.String("contosoResourceGroup"),
			ResourceName:      pulumi.String("contoso"),
			Sku: &azureactivedirectory.CIAMResourceSKUArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
