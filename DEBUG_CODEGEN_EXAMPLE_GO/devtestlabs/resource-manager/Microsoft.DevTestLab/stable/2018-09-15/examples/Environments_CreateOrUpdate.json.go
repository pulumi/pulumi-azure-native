package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewEnvironment(ctx, "environment", &devtestlab.EnvironmentArgs{
			DeploymentProperties: &devtestlab.EnvironmentDeploymentPropertiesArgs{
				ArmTemplateId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/armTemplates/{armTemplateName}"),
				Parameters:    devtestlab.ArmTemplateParameterPropertiesArray{},
			},
			LabName:           pulumi.String("{labName}"),
			Name:              pulumi.String("{environmentName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			UserName:          pulumi.String("@me"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
