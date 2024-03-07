package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewAdministrator(ctx, "administrator", &dbforpostgresql.AdministratorArgs{
			ObjectId:          pulumi.String("oooooooo-oooo-oooo-oooo-oooooooooooo"),
			PrincipalName:     pulumi.String("testuser1@microsoft.com"),
			PrincipalType:     pulumi.String("User"),
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("testserver"),
			TenantId:          pulumi.String("tttttttt-tttt-tttt-tttt-tttttttttttt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
