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
				Mode:       resources.DeploymentModeIncremental,
				Parameters: nil,
				TemplateLink: &resources.TemplateLinkArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"),
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
