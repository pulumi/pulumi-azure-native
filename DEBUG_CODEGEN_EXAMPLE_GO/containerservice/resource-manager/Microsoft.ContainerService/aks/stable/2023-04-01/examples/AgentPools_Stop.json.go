package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
			AgentPoolName: pulumi.String("agentpool1"),
			PowerState: &containerservice.PowerStateArgs{
				Code: pulumi.String("Stopped"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("clustername1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
