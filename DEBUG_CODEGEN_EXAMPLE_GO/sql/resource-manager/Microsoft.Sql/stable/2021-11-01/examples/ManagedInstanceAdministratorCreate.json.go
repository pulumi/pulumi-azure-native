package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedInstanceAdministrator(ctx, "managedInstanceAdministrator", &sql.ManagedInstanceAdministratorArgs{
			AdministratorName:   pulumi.String("ActiveDirectory"),
			AdministratorType:   pulumi.String("ActiveDirectory"),
			Login:               pulumi.String("bob@contoso.com"),
			ManagedInstanceName: pulumi.String("managedInstance"),
			ResourceGroupName:   pulumi.String("Default-SQL-SouthEastAsia"),
			Sid:                 pulumi.String("44444444-3333-2222-1111-000000000000"),
			TenantId:            pulumi.String("55555555-4444-3333-2222-111111111111"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
