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
					QueryString: pulumi.String("sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=xxxxxxxx0xxxxxxxxxxxxx%2bxxxxxxxxxxxxxxxxxxxx%3d"),
					Uri:         pulumi.String("https://example.com/exampleTemplate.json"),
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
