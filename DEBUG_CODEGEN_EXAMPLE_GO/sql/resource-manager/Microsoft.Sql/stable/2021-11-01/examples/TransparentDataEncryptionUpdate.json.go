package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewTransparentDataEncryption(ctx, "transparentDataEncryption", &sql.TransparentDataEncryptionArgs{
			DatabaseName:      pulumi.String("testdb"),
			ResourceGroupName: pulumi.String("securitytde-42-rg"),
			ServerName:        pulumi.String("securitytde-42"),
			State:             sql.TransparentDataEncryptionStateEnabled,
			TdeName:           pulumi.String("current"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
