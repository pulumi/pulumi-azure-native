package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewIntegrationRuntime(ctx, "integrationRuntime", &datafactory.IntegrationRuntimeArgs{
			FactoryName:            pulumi.String("exampleFactoryName"),
			IntegrationRuntimeName: pulumi.String("exampleIntegrationRuntime"),
			Properties: datafactory.SelfHostedIntegrationRuntime{
				Description: "A selfhosted integration runtime",
				Type:        "SelfHosted",
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
