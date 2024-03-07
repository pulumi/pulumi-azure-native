package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewL2IsolationDomain(ctx, "l2IsolationDomain", &managednetworkfabric.L2IsolationDomainArgs{
			L2IsolationDomainName: pulumi.String("example-l2domain"),
			Location:              pulumi.String("eastus"),
			Mtu:                   pulumi.Int(1500),
			NetworkFabricId:       pulumi.String("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName"),
			ResourceGroupName:     pulumi.String("resourceGroupName"),
			VlanId:                pulumi.Int(501),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
