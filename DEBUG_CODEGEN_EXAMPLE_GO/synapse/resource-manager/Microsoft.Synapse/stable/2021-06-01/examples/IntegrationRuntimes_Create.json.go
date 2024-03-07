package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewIntegrationRuntime(ctx, "integrationRuntime", &synapse.IntegrationRuntimeArgs{
			IntegrationRuntimeName: pulumi.String("exampleIntegrationRuntime"),
			Properties: synapse.SelfHostedIntegrationRuntime{
				Description: "A selfhosted integration runtime",
				Type:        "SelfHosted",
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
			WorkspaceName:     pulumi.String("exampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
