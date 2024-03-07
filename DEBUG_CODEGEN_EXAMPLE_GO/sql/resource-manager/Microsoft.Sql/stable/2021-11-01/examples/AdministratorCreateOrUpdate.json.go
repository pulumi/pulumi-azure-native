package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerAzureADAdministrator(ctx, "serverAzureADAdministrator", &sql.ServerAzureADAdministratorArgs{
			AdministratorName: pulumi.String("ActiveDirectory"),
			AdministratorType: pulumi.String("ActiveDirectory"),
			Login:             pulumi.String("bob@contoso.com"),
			ResourceGroupName: pulumi.String("sqlcrudtest-4799"),
			ServerName:        pulumi.String("sqlcrudtest-6440"),
			Sid:               pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			TenantId:          pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
