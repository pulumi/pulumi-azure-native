package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewDeploymentAtManagementGroupScope(ctx, "deploymentAtManagementGroupScope", &resources.DeploymentAtManagementGroupScopeArgs{
			DeploymentName: pulumi.String("my-deployment"),
			GroupId:        pulumi.String("my-management-group-id"),
			Location:       pulumi.String("eastus"),
			Properties: &resources.DeploymentPropertiesArgs{
				Mode:       resources.DeploymentModeIncremental,
				Parameters: nil,
				TemplateLink: &resources.TemplateLinkArgs{
					Uri: pulumi.String("https://example.com/exampleTemplate.json"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
