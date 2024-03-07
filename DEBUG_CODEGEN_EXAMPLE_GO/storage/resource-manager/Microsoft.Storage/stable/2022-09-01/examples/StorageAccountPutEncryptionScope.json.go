package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewEncryptionScope(ctx, "encryptionScope", &storage.EncryptionScopeArgs{
			AccountName:         pulumi.String("{storage-account-name}"),
			EncryptionScopeName: pulumi.String("{encryption-scope-name}"),
			ResourceGroupName:   pulumi.String("resource-group-name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
