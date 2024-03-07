package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewWorkspaceSetting(ctx, "workspaceSetting", &security.WorkspaceSettingArgs{
			Scope:                pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
			WorkspaceId:          pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace"),
			WorkspaceSettingName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
