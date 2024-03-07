package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewEncryptionProtector(ctx, "encryptionProtector", &sql.EncryptionProtectorArgs{
			AutoRotationEnabled:     pulumi.Bool(false),
			EncryptionProtectorName: pulumi.String("current"),
			ResourceGroupName:       pulumi.String("sqlcrudtest-7398"),
			ServerKeyName:           pulumi.String("someVault_someKey_01234567890123456789012345678901"),
			ServerKeyType:           pulumi.String("AzureKeyVault"),
			ServerName:              pulumi.String("sqlcrudtest-4645"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
