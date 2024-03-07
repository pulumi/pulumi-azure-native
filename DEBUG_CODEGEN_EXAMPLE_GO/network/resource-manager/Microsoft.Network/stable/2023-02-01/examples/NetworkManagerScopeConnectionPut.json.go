package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewScopeConnection(ctx, "scopeConnection", &network.ScopeConnectionArgs{
			Description:         pulumi.String("This is a scope connection to a cross tenant subscription."),
			NetworkManagerName:  pulumi.String("testNetworkManager"),
			ResourceGroupName:   pulumi.String("rg1"),
			ResourceId:          pulumi.String("subscriptions/f0dc2b34-dfad-40e4-83e0-2309fed8d00b"),
			ScopeConnectionName: pulumi.String("TestScopeConnection"),
			TenantId:            pulumi.String("6babcaad-604b-40ac-a9d7-9fd97c0b779f"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
