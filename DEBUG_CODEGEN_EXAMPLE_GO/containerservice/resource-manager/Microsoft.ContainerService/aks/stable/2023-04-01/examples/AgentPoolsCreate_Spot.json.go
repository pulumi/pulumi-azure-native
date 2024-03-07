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
			NodeLabels: pulumi.StringMap{
				"key1": pulumi.String("val1"),
			},
			NodeTaints: pulumi.StringArray{
				pulumi.String("Key1=Value1:NoSchedule"),
			},
			OrchestratorVersion:    pulumi.String(""),
			OsType:                 pulumi.String("Linux"),
			ResourceGroupName:      pulumi.String("rg1"),
			ResourceName:           pulumi.String("clustername1"),
			ScaleSetEvictionPolicy: pulumi.String("Delete"),
			ScaleSetPriority:       pulumi.String("Spot"),
			Tags: pulumi.StringMap{
				"name1": pulumi.String("val1"),
			},
			VmSize: pulumi.String("Standard_DS1_v2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
