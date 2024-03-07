package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewSecret(ctx, "secret", &keyvault.SecretArgs{
			Properties: &keyvault.SecretPropertiesArgs{
				Value: pulumi.String("secret-value"),
			},
			ResourceGroupName: pulumi.String("sample-group"),
			SecretName:        pulumi.String("secret-name"),
			VaultName:         pulumi.String("sample-vault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
