package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewRoleManagementPolicyAssignment(ctx, "roleManagementPolicyAssignment", &authorization.RoleManagementPolicyAssignmentArgs{
			PolicyId:                           pulumi.String("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
			RoleDefinitionId:                   pulumi.String("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
			RoleManagementPolicyAssignmentName: pulumi.String("b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
			Scope:                              pulumi.String("providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
