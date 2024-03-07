package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewWorkspace(ctx, "workspace", &desktopvirtualization.WorkspaceArgs{
			Description:       pulumi.String("des1"),
			FriendlyName:      pulumi.String("friendly"),
			Location:          pulumi.String("centralus"),
			ResourceGroupName: pulumi.String("resourceGroup1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			WorkspaceName: pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
