package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewSourceControl(ctx, "sourceControl", &securityinsights.SourceControlArgs{
			ContentTypes: pulumi.StringArray{
				pulumi.String("AnalyticRules"),
				pulumi.String("Workbook"),
			},
			Description: pulumi.String("This is a source control"),
			DisplayName: pulumi.String("My Source Control"),
			RepoType:    pulumi.String("Github"),
			Repository: &securityinsights.RepositoryArgs{
				Branch:     pulumi.String("master"),
				DisplayUrl: pulumi.String("https://github.com/user/repo"),
				PathMapping: securityinsights.ContentPathMapArray{
					&securityinsights.ContentPathMapArgs{
						ContentType: pulumi.String("AnalyticRules"),
						Path:        pulumi.String("path/to/rules"),
					},
					&securityinsights.ContentPathMapArgs{
						ContentType: pulumi.String("Workbook"),
						Path:        pulumi.String("path/to/workbooks"),
					},
				},
				Url: pulumi.String("https://github.com/user/repo"),
			},
			ResourceGroupName: pulumi.String("myRg"),
			SourceControlId:   pulumi.String("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
