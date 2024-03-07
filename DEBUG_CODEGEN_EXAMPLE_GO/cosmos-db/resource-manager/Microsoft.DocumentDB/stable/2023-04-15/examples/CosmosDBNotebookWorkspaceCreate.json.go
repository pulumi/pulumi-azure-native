package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewNotebookWorkspace(ctx, "notebookWorkspace", &documentdb.NotebookWorkspaceArgs{
			AccountName:           pulumi.String("ddb1"),
			NotebookWorkspaceName: pulumi.String("default"),
			ResourceGroupName:     pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
