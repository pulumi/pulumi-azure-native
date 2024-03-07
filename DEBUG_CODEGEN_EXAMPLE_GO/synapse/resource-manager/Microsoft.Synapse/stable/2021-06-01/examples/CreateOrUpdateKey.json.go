package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewKey(ctx, "key", &synapse.KeyArgs{
			IsActiveCMK:       pulumi.Bool(true),
			KeyName:           pulumi.String("somekey"),
			KeyVaultUrl:       pulumi.String("https://vault.azure.net/keys/somesecret"),
			ResourceGroupName: pulumi.String("ExampleResourceGroup"),
			WorkspaceName:     pulumi.String("ExampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
