package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewNetworkFunctionDefinitionGroup(ctx, "networkFunctionDefinitionGroup", &hybridnetwork.NetworkFunctionDefinitionGroupArgs{
			Location:                           pulumi.String("eastus"),
			NetworkFunctionDefinitionGroupName: pulumi.String("TestNetworkFunctionDefinitionGroupName"),
			PublisherName:                      pulumi.String("TestPublisher"),
			ResourceGroupName:                  pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
