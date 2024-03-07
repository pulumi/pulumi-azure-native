package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServer(ctx, "server", &sql.ServerArgs{
			AdministratorLogin:         pulumi.String("dummylogin"),
			AdministratorLoginPassword: pulumi.String("PLACEHOLDER"),
			Administrators: &sql.ServerExternalAdministratorArgs{
				AzureADOnlyAuthentication: pulumi.Bool(true),
				Login:                     pulumi.String("bob@contoso.com"),
				PrincipalType:             pulumi.String("User"),
				Sid:                       pulumi.String("00000011-1111-2222-2222-123456789111"),
				TenantId:                  pulumi.String("00000011-1111-2222-2222-123456789111"),
			},
			Location:                      pulumi.String("Japan East"),
			PublicNetworkAccess:           pulumi.String("Enabled"),
			ResourceGroupName:             pulumi.String("sqlcrudtest-7398"),
			RestrictOutboundNetworkAccess: pulumi.String("Enabled"),
			ServerName:                    pulumi.String("sqlcrudtest-4645"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
