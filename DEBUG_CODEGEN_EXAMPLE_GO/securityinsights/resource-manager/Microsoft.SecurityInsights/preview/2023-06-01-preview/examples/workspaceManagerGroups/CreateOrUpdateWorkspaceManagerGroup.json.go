package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewWorkspaceManagerGroup(ctx, "workspaceManagerGroup", &securityinsights.WorkspaceManagerGroupArgs{
			Description: pulumi.String("Group of all financial and banking institutions"),
			DisplayName: pulumi.String("Banks"),
			MemberResourceNames: pulumi.StringArray{
				pulumi.String("afbd324f-6c48-459c-8710-8d1e1cd03812"),
				pulumi.String("f5fa104e-c0e3-4747-9182-d342dc048a9e"),
			},
			ResourceGroupName:         pulumi.String("myRg"),
			WorkspaceManagerGroupName: pulumi.String("37207a7a-3b8a-438f-a559-c7df400e1b96"),
			WorkspaceName:             pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
