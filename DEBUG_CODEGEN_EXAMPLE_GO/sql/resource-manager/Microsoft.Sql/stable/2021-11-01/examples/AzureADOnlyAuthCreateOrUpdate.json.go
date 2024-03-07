package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerAzureADOnlyAuthentication(ctx, "serverAzureADOnlyAuthentication", &sql.ServerAzureADOnlyAuthenticationArgs{
			AuthenticationName:        pulumi.String("Default"),
			AzureADOnlyAuthentication: pulumi.Bool(false),
			ResourceGroupName:         pulumi.String("sqlcrudtest-4799"),
			ServerName:                pulumi.String("sqlcrudtest-6440"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
