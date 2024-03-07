package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewDeployment(ctx, "deployment", &cognitiveservices.DeploymentArgs{
			AccountName:    pulumi.String("accountName"),
			DeploymentName: pulumi.String("deploymentName"),
			Properties: cognitiveservices.DeploymentPropertiesResponse{
				Model: &cognitiveservices.DeploymentModelArgs{
					Format:  pulumi.String("OpenAI"),
					Name:    pulumi.String("ada"),
					Version: pulumi.String("1"),
				},
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Sku: &cognitiveservices.SkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
