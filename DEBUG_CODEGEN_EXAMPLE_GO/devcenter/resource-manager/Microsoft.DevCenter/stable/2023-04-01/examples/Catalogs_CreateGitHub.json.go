package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewCatalog(ctx, "catalog", &devcenter.CatalogArgs{
			CatalogName:   pulumi.String("CentralCatalog"),
			DevCenterName: pulumi.String("Contoso"),
			GitHub: &devcenter.GitCatalogArgs{
				Branch:           pulumi.String("main"),
				Path:             pulumi.String("/templates"),
				SecretIdentifier: pulumi.String("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
				Uri:              pulumi.String("https://github.com/Contoso/centralrepo-fake.git"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
