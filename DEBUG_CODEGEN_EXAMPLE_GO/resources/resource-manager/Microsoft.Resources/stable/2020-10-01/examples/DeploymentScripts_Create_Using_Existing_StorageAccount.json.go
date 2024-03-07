package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewAzureCliScript(ctx, "azureCliScript", &resources.AzureCliScriptArgs{
			ResourceGroupName: pulumi.String("script-rg"),
			ScriptName:        pulumi.String("MyDeploymentScript"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
