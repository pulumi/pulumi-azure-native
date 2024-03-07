package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicelinker/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicelinker.NewConnector(ctx, "connector", &servicelinker.ConnectorArgs{
			AuthInfo: servicelinker.SecretAuthInfo{
				AuthType: "secret",
			},
			ConnectorName:     pulumi.String("connectorName"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("test-rg"),
			SecretStore: &servicelinker.SecretStoreArgs{
				KeyVaultId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/test-kv"),
			},
			TargetService: servicelinker.AzureResource{
				Id:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
				Type: "AzureResource",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
