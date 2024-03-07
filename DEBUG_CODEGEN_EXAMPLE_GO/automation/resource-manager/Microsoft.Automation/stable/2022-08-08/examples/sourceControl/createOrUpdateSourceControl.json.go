package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewSourceControl(ctx, "sourceControl", &automation.SourceControlArgs{
			AutoSync:              pulumi.Bool(true),
			AutomationAccountName: pulumi.String("sampleAccount9"),
			Branch:                pulumi.String("master"),
			Description:           pulumi.String("my description"),
			FolderPath:            pulumi.String("/folderOne/folderTwo"),
			PublishRunbook:        pulumi.Bool(true),
			RepoUrl:               pulumi.String("https://sampleUser.visualstudio.com/myProject/_git/myRepository"),
			ResourceGroupName:     pulumi.String("rg"),
			SecurityToken: &automation.SourceControlSecurityTokenPropertiesArgs{
				AccessToken: pulumi.String("******"),
				TokenType:   pulumi.String("PersonalAccessToken"),
			},
			SourceControlName: pulumi.String("sampleSourceControl"),
			SourceType:        pulumi.String("VsoGit"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
