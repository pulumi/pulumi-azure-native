package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewRoleDefinition(ctx, "roleDefinition", &authorization.RoleDefinitionArgs{
			RoleDefinitionId: pulumi.String("roleDefinitionId"),
			Scope:            pulumi.String("scope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
