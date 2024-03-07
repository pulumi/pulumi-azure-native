package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewRole(ctx, "role", &dbforpostgresql.RoleArgs{
			ClusterName:       pulumi.String("pgtestsvc4"),
			Password:          pulumi.String("password"),
			ResourceGroupName: pulumi.String("TestGroup"),
			RoleName:          pulumi.String("role1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
