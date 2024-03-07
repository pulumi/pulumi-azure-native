package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkFabricController(ctx, "networkFabricController", &managednetworkfabric.NetworkFabricControllerArgs{
			Annotation: pulumi.String("lab 1"),
			InfrastructureExpressRouteConnections: []managednetworkfabric.ExpressRouteConnectionInformationArgs{
				{
					ExpressRouteAuthorizationKey: pulumi.String("xxxxxxx"),
					ExpressRouteCircuitId:        pulumi.String("/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
				},
			},
			Ipv4AddressSpace: pulumi.String("172.253.0.0/19"),
			Location:         pulumi.String("eastus"),
			ManagedResourceGroupConfiguration: &managednetworkfabric.ManagedResourceGroupConfigurationArgs{
				Location: pulumi.String("eastus"),
				Name:     pulumi.String("managedResourceGroupName"),
			},
			NetworkFabricControllerName: pulumi.String("NetworkControllerName"),
			ResourceGroupName:           pulumi.String("resourceGroupName"),
			WorkloadExpressRouteConnections: []managednetworkfabric.ExpressRouteConnectionInformationArgs{
				{
					ExpressRouteAuthorizationKey: pulumi.String("xxxxx"),
					ExpressRouteCircuitId:        pulumi.String("/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
