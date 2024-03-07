package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
			AgentPoolName: pulumi.String("agentpool1"),
			Count:         pulumi.Int(3),
			CreationData: &containerservice.CreationDataArgs{
				SourceResourceId: pulumi.String("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
			},
			EnableFIPS:          pulumi.Bool(true),
			OrchestratorVersion: pulumi.String(""),
			OsType:              pulumi.String("Linux"),
			ResourceGroupName:   pulumi.String("rg1"),
			ResourceName:        pulumi.String("clustername1"),
			VmSize:              pulumi.String("Standard_DS2_v2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
