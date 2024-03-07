package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
			AgentPoolName:             pulumi.String("agentpool1"),
			Count:                     pulumi.Int(3),
			OrchestratorVersion:       pulumi.String(""),
			OsType:                    pulumi.String("Linux"),
			ProximityPlacementGroupID: pulumi.String("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
			ResourceGroupName:         pulumi.String("rg1"),
			ResourceName:              pulumi.String("clustername1"),
			VmSize:                    pulumi.String("Standard_DS2_v2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
