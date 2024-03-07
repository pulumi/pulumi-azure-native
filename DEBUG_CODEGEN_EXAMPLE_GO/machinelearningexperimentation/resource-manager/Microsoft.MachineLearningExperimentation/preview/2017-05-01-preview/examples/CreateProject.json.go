package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningexperimentation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningexperimentation.NewProject(ctx, "project", &machinelearningexperimentation.ProjectArgs{
			AccountName:       pulumi.String("testaccount"),
			FriendlyName:      pulumi.String("testName"),
			Gitrepo:           pulumi.String("https://github/abc"),
			Location:          pulumi.String("East US"),
			ProjectName:       pulumi.String("testProject"),
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
