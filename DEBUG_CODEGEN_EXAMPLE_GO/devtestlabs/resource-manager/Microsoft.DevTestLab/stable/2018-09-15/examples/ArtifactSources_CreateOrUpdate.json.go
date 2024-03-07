package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewArtifactSource(ctx, "artifactSource", &devtestlab.ArtifactSourceArgs{
			ArmTemplateFolderPath: pulumi.String("{armTemplateFolderPath}"),
			BranchRef:             pulumi.String("{branchRef}"),
			DisplayName:           pulumi.String("{displayName}"),
			FolderPath:            pulumi.String("{folderPath}"),
			LabName:               pulumi.String("{labName}"),
			Name:                  pulumi.String("{artifactSourceName}"),
			ResourceGroupName:     pulumi.String("resourceGroupName"),
			SecurityToken:         pulumi.String("{securityToken}"),
			SourceType:            pulumi.String("{VsoGit|GitHub|StorageAccount}"),
			Status:                pulumi.String("{Enabled|Disabled}"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			Uri: pulumi.String("{artifactSourceUri}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
