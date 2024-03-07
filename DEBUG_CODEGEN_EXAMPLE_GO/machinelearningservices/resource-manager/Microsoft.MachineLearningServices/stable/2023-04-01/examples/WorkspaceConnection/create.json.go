package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewWorkspaceConnection(ctx, "workspaceConnection", &machinelearningservices.WorkspaceConnectionArgs{
			ConnectionName: pulumi.String("connection-1"),
			Properties: machinelearningservices.NoneAuthTypeWorkspaceConnectionProperties{
				AuthType: "None",
				Category: "ContainerRegistry",
				Target:   "www.facebook.com",
			},
			ResourceGroupName: pulumi.String("resourceGroup-1"),
			WorkspaceName:     pulumi.String("workspace-1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
