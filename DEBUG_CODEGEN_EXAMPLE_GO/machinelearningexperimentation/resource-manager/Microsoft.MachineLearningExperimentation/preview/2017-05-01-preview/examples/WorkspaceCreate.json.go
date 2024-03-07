package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningexperimentation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningexperimentation.NewWorkspace(ctx, "workspace", &machinelearningexperimentation.WorkspaceArgs{
			AccountName:       pulumi.String("testaccount"),
			FriendlyName:      pulumi.String("testName"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("TagValue1"),
			},
			WorkspaceName: pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
