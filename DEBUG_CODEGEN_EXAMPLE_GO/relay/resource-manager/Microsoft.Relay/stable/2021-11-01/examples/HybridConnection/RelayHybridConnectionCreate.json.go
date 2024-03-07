package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/relay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := relay.NewHybridConnection(ctx, "hybridConnection", &relay.HybridConnectionArgs{
			HybridConnectionName:        pulumi.String("example-Relay-Hybrid-01"),
			NamespaceName:               pulumi.String("example-RelayNamespace-01"),
			RequiresClientAuthorization: pulumi.Bool(true),
			ResourceGroupName:           pulumi.String("resourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
