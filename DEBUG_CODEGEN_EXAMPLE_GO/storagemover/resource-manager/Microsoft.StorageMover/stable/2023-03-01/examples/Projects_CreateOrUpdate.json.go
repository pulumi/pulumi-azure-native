package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagemover/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagemover.NewProject(ctx, "project", &storagemover.ProjectArgs{
			Description:       pulumi.String("Example Project Description"),
			ProjectName:       pulumi.String("examples-projectName"),
			ResourceGroupName: pulumi.String("examples-rg"),
			StorageMoverName:  pulumi.String("examples-storageMoverName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
