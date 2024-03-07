package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewLinkedWorkspace(ctx, "linkedWorkspace", &machinelearningservices.LinkedWorkspaceArgs{
			LinkName: pulumi.String("link-1"),
			Name:     pulumi.String("link-1"),
			Properties: &machinelearningservices.LinkedWorkspacePropsArgs{
				LinkedWorkspaceResourceId:      pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.Synapse/workspaces/linkedWorkspace-1"),
				UserAssignedIdentityResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai1"),
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
