package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewSqlPoolTransparentDataEncryption(ctx, "sqlPoolTransparentDataEncryption", &synapse.SqlPoolTransparentDataEncryptionArgs{
			ResourceGroupName:             pulumi.String("sqlcrudtest-6852"),
			SqlPoolName:                   pulumi.String("sqlcrudtest-9187"),
			Status:                        pulumi.String("Enabled"),
			TransparentDataEncryptionName: pulumi.String("current"),
			WorkspaceName:                 pulumi.String("sqlcrudtest-2080"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
