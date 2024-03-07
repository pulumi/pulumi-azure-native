package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewWorkspaceManagerMember(ctx, "workspaceManagerMember", &securityinsights.WorkspaceManagerMemberArgs{
			ResourceGroupName:          pulumi.String("myRg"),
			TargetWorkspaceResourceId:  pulumi.String("/subscriptions/7aef9d48-814f-45ad-b644-b0343316e174/resourceGroups/otherRg/providers/Microsoft.OperationalInsights/workspaces/Example_Workspace"),
			TargetWorkspaceTenantId:    pulumi.String("f676d436-8d16-42db-81b7-ab578e110ccd"),
			WorkspaceManagerMemberName: pulumi.String("afbd324f-6c48-459c-8710-8d1e1cd03812"),
			WorkspaceName:              pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
