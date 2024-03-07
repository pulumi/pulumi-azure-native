package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerKey(ctx, "serverKey", &sql.ServerKeyArgs{
			KeyName:           pulumi.String("someVault_someKey_01234567890123456789012345678901"),
			ResourceGroupName: pulumi.String("sqlcrudtest-7398"),
			ServerKeyType:     pulumi.String("AzureKeyVault"),
			ServerName:        pulumi.String("sqlcrudtest-4645"),
			Uri:               pulumi.String("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
