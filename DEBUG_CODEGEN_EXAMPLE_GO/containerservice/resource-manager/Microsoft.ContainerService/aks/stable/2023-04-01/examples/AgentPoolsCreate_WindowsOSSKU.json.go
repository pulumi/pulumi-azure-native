package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
			AgentPoolName:       pulumi.String("wnp2"),
			Count:               pulumi.Int(3),
			OrchestratorVersion: pulumi.String("1.23.3"),
			OsSKU:               pulumi.String("Windows2022"),
			OsType:              pulumi.String("Windows"),
			ResourceGroupName:   pulumi.String("rg1"),
			ResourceName:        pulumi.String("clustername1"),
			VmSize:              pulumi.String("Standard_D4s_v3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
