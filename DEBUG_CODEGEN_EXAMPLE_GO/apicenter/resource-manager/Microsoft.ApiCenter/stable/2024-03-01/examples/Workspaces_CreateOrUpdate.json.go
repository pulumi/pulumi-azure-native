package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewWorkspace(ctx, "workspace", &apicenter.WorkspaceArgs{
			ResourceGroupName: pulumi.String("contoso-resources"),
			ServiceName:       pulumi.String("contoso"),
			Title:             pulumi.String("default"),
			WorkspaceName:     pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
