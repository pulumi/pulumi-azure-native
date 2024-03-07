package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewWorkspaceManagerAssignment(ctx, "workspaceManagerAssignment", &securityinsights.WorkspaceManagerAssignmentArgs{
			Items: securityinsights.AssignmentItemArray{
				&securityinsights.AssignmentItemArgs{
					ResourceId: pulumi.String("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleOne"),
				},
				&securityinsights.AssignmentItemArgs{
					ResourceId: pulumi.String("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleTwo"),
				},
			},
			ResourceGroupName:              pulumi.String("myRg"),
			TargetResourceName:             pulumi.String("37207a7a-3b8a-438f-a559-c7df400e1b96"),
			WorkspaceManagerAssignmentName: pulumi.String("47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58"),
			WorkspaceName:                  pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
