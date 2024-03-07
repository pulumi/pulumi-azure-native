package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagemover/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagemover.NewAgent(ctx, "agent", &storagemover.AgentArgs{
			AgentName:         pulumi.String("examples-agentName"),
			ArcResourceId:     pulumi.String("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
			ArcVmUuid:         pulumi.String("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
			Description:       pulumi.String("Example Agent Description"),
			ResourceGroupName: pulumi.String("examples-rg"),
			StorageMoverName:  pulumi.String("examples-storageMoverName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
