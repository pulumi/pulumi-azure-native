package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewEndpointDeployment(ctx, "endpointDeployment", &machinelearningservices.EndpointDeploymentArgs{
			DeploymentName: pulumi.String("text-davinci-003"),
			EndpointName:   pulumi.String("Azure.OpenAI"),
			Properties: machinelearningservices.OpenAIEndpointDeploymentResourceProperties{
				Model: machinelearningservices.EndpointDeploymentModel{
					Format:  "OpenAI",
					Name:    "text-davinci-003",
					Version: "1",
				},
				Type:                 "Azure.OpenAI",
				VersionUpgradeOption: "OnceNewDefaultVersionAvailable",
			},
			ResourceGroupName: pulumi.String("resourceGroup-1"),
			WorkspaceName:     pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
