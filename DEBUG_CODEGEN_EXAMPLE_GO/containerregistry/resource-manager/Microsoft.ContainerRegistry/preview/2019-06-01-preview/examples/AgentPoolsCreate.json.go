package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewAgentPool(ctx, "agentPool", &containerregistry.AgentPoolArgs{
			AgentPoolName:     pulumi.String("myAgentPool"),
			Count:             pulumi.Int(1),
			Location:          pulumi.String("WESTUS"),
			Os:                pulumi.String("Linux"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			Tier: pulumi.String("S1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
