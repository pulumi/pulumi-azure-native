package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/agfoodplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := agfoodplatform.NewDataConnector(ctx, "dataConnector", &agfoodplatform.DataConnectorArgs{
			DataConnectorName:                     pulumi.String("WeatherIBM"),
			DataManagerForAgricultureResourceName: pulumi.String("examples-dataManagerForAgricultureResourceName"),
			Properties: agfoodplatform.DataConnectorPropertiesResponse{
				Credentials: agfoodplatform.ApiKeyAuthCredentials{
					ApiKey: agfoodplatform.KeyVaultProperties{
						KeyName:     "abcApiKey",
						KeyVaultUri: "https://testKeyVault.vault.azure.net/",
						KeyVersion:  "239c0475c7d44f20b0fc27d3fe90a41d",
					},
					Kind: "ApiKeyAuthCredentials",
				},
			},
			ResourceGroupName: pulumi.String("examples-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
