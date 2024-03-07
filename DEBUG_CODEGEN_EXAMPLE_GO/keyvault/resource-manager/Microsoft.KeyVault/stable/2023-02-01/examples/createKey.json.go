package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewKey(ctx, "key", &keyvault.KeyArgs{
			KeyName: pulumi.String("sample-key-name"),
			Properties: &keyvault.KeyPropertiesArgs{
				Kty: pulumi.String("RSA"),
			},
			ResourceGroupName: pulumi.String("sample-group"),
			VaultName:         pulumi.String("sample-vault-name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
