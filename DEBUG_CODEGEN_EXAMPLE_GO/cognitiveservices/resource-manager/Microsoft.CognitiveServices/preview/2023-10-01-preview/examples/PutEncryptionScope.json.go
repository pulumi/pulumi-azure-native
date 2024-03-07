package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewEncryptionScope(ctx, "encryptionScope", &cognitiveservices.EncryptionScopeArgs{
			AccountName:         pulumi.String("accountName"),
			EncryptionScopeName: pulumi.String("encryptionScopeName"),
			Properties: cognitiveservices.EncryptionScopePropertiesResponse{
				KeySource: pulumi.String("Microsoft.KeyVault"),
				KeyVaultProperties: &cognitiveservices.KeyVaultPropertiesArgs{
					IdentityClientId: pulumi.String("00000000-0000-0000-0000-000000000000"),
					KeyName:          pulumi.String("DevKeyWestUS2"),
					KeyVaultUri:      pulumi.String("https://devkvwestus2.vault.azure.net/"),
					KeyVersion:       pulumi.String("9f85549d7bf14ff4bf178c10d3bdca95"),
				},
				State: pulumi.String("Enabled"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
