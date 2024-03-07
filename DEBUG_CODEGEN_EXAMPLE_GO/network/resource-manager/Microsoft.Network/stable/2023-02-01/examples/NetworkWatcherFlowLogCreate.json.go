package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFlowLog(ctx, "flowLog", &network.FlowLogArgs{
			Enabled:     pulumi.Bool(true),
			FlowLogName: pulumi.String("fl"),
			Format: &network.FlowLogFormatParametersArgs{
				Type:    pulumi.String("JSON"),
				Version: pulumi.Int(1),
			},
			Location:           pulumi.String("centraluseuap"),
			NetworkWatcherName: pulumi.String("nw1"),
			ResourceGroupName:  pulumi.String("rg1"),
			StorageId:          pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
			TargetResourceId:   pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
