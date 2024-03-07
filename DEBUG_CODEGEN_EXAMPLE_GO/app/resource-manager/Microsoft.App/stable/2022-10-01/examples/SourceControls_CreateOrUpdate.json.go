package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewContainerAppsSourceControl(ctx, "containerAppsSourceControl", &app.ContainerAppsSourceControlArgs{
			Branch:           pulumi.String("master"),
			ContainerAppName: pulumi.String("testcanadacentral"),
			GithubActionConfiguration: app.GithubActionConfigurationResponse{
				AzureCredentials: &app.AzureCredentialsArgs{
					ClientId:     pulumi.String("<clientid>"),
					ClientSecret: pulumi.String("<clientsecret>"),
					TenantId:     pulumi.String("<tenantid>"),
				},
				ContextPath: pulumi.String("./"),
				Image:       pulumi.String("image/tag"),
				RegistryInfo: &app.RegistryInfoArgs{
					RegistryPassword: pulumi.String("<registrypassword>"),
					RegistryUrl:      pulumi.String("xwang971reg.azurecr.io"),
					RegistryUserName: pulumi.String("xwang971reg"),
				},
			},
			RepoUrl:           pulumi.String("https://github.com/xwang971/ghatest"),
			ResourceGroupName: pulumi.String("workerapps-rg-xj"),
			SourceControlName: pulumi.String("current"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
