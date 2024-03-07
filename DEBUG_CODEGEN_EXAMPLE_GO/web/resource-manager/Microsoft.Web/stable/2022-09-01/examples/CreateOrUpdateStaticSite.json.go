package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewStaticSite(ctx, "staticSite", &web.StaticSiteArgs{
			Branch: pulumi.String("master"),
			BuildProperties: &web.StaticSiteBuildPropertiesArgs{
				ApiLocation:         pulumi.String("api"),
				AppArtifactLocation: pulumi.String("build"),
				AppLocation:         pulumi.String("app"),
			},
			Location:          pulumi.String("West US 2"),
			Name:              pulumi.String("testStaticSite0"),
			RepositoryToken:   pulumi.String("repoToken123"),
			RepositoryUrl:     pulumi.String("https://github.com/username/RepoName"),
			ResourceGroupName: pulumi.String("rg"),
			Sku: &web.SkuDescriptionArgs{
				Name: pulumi.String("Basic"),
				Tier: pulumi.String("Basic"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
