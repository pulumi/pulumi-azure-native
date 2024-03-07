package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewDeployment(ctx, "deployment", &apicenter.DeploymentArgs{
			ApiName:           pulumi.String("echo-api"),
			DefinitionId:      pulumi.String("/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
			DeploymentName:    pulumi.String("production"),
			Description:       pulumi.String("Public cloud production deployment."),
			EnvironmentId:     pulumi.String("/workspaces/default/environments/production"),
			ResourceGroupName: pulumi.String("contoso-resources"),
			Server: &apicenter.DeploymentServerArgs{
				RuntimeUri: pulumi.StringArray{
					pulumi.String("https://api.contoso.com"),
				},
			},
			ServiceName:   pulumi.String("contoso"),
			State:         pulumi.String("active"),
			Title:         pulumi.String("Production deployment"),
			WorkspaceName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
