package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewRoleAssignment(ctx, "roleAssignment", &authorization.RoleAssignmentArgs{
			PrincipalId:        pulumi.String("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:      pulumi.String("User"),
			RoleAssignmentName: pulumi.String("05c5a614-a7d6-4502-b150-c2fb455033ff"),
			RoleDefinitionId:   pulumi.String("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
			Scope:              pulumi.String("subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
