package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewDeploymentAtScope(ctx, "deploymentAtScope", &resources.DeploymentAtScopeArgs{
			DeploymentName: pulumi.String("my-deployment"),
			Location:       pulumi.String("eastus"),
			Properties: resources.DeploymentPropertiesExtendedResponse{
				Mode:       resources.DeploymentModeIncremental,
				Parameters: nil,
				TemplateLink: &resources.TemplateLinkArgs{
					Uri: pulumi.String("https://example.com/exampleTemplate.json"),
				},
			},
			Scope: pulumi.String("providers/Microsoft.Management/managementGroups/my-management-group-id"),
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("tag-value-1"),
				"tagKey2": pulumi.String("tag-value-2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
