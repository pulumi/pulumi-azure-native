package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/portal/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := portal.NewTenantConfiguration(ctx, "tenantConfiguration", &portal.TenantConfigurationArgs{
			ConfigurationName:             pulumi.String("default"),
			EnforcePrivateMarkdownStorage: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
