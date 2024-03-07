package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewCatalog(ctx, "catalog", &devcenter.CatalogArgs{
			AdoGit: &devcenter.GitCatalogArgs{
				Branch:           pulumi.String("main"),
				Path:             pulumi.String("/templates"),
				SecretIdentifier: pulumi.String("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
				Uri:              pulumi.String("https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso"),
			},
			CatalogName:       pulumi.String("CentralCatalog"),
			DevCenterName:     pulumi.String("Contoso"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
