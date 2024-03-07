package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewWorkspaceManagerConfiguration(ctx, "workspaceManagerConfiguration", &securityinsights.WorkspaceManagerConfigurationArgs{
			Mode:                              pulumi.String("Enabled"),
			ResourceGroupName:                 pulumi.String("myRg"),
			WorkspaceManagerConfigurationName: pulumi.String("default"),
			WorkspaceName:                     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
