package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewConfigurationService(ctx, "configurationService", &appplatform.ConfigurationServiceArgs{
			ConfigurationServiceName: pulumi.String("default"),
			Properties: &appplatform.ConfigurationServicePropertiesArgs{
				Settings: &appplatform.ConfigurationServiceSettingsArgs{
					GitProperty: &appplatform.ConfigurationServiceGitPropertyArgs{
						Repositories: appplatform.ConfigurationServiceGitRepositoryArray{
							&appplatform.ConfigurationServiceGitRepositoryArgs{
								Label: pulumi.String("master"),
								Name:  pulumi.String("fake"),
								Patterns: pulumi.StringArray{
									pulumi.String("app/dev"),
								},
								Uri: pulumi.String("https://github.com/fake-user/fake-repository"),
							},
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
