package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewRoleAssignment(ctx, "roleAssignment", &customerinsights.RoleAssignmentArgs{
			AssignmentName: pulumi.String("assignmentName8976"),
			HubName:        pulumi.String("sdkTestHub"),
			Principals: customerinsights.AssignmentPrincipalArray{
				&customerinsights.AssignmentPrincipalArgs{
					PrincipalId:   pulumi.String("4c54c38ffa9b416ba5a6d6c8a20cbe7e"),
					PrincipalType: pulumi.String("User"),
				},
				&customerinsights.AssignmentPrincipalArgs{
					PrincipalId:   pulumi.String("93061d15a5054f2b9948ae25724cf9d5"),
					PrincipalType: pulumi.String("User"),
				},
			},
			ResourceGroupName: pulumi.String("TestHubRG"),
			Role:              customerinsights.RoleTypesAdmin,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
