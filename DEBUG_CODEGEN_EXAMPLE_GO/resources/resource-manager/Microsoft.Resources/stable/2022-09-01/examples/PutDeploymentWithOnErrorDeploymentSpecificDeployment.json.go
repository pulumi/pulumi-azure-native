package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewDeployment(ctx, "deployment", &resources.DeploymentArgs{
			DeploymentName: pulumi.String("my-deployment"),
			Properties: resources.DeploymentPropertiesExtendedResponse{
				Mode: resources.DeploymentModeComplete,
				OnErrorDeployment: &resources.OnErrorDeploymentArgs{
					DeploymentName: pulumi.String("name-of-deployment-to-use"),
					Type:           resources.OnErrorDeploymentTypeSpecificDeployment,
				},
				Parameters: nil,
				TemplateLink: &resources.TemplateLinkArgs{
					Uri: pulumi.String("https://example.com/exampleTemplate.json"),
				},
			},
			ResourceGroupName: pulumi.String("my-resource-group"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
