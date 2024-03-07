package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewL3IsolationDomain(ctx, "l3IsolationDomain", &managednetworkfabric.L3IsolationDomainArgs{
			AggregateRouteConfiguration: &managednetworkfabric.AggregateRouteConfigurationArgs{
				Ipv4Routes: managednetworkfabric.AggregateRouteArray{
					&managednetworkfabric.AggregateRouteArgs{
						Prefix: pulumi.String("10.0.0.0/24"),
					},
				},
				Ipv6Routes: managednetworkfabric.AggregateRouteArray{
					&managednetworkfabric.AggregateRouteArgs{
						Prefix: pulumi.String("10.0.0.1"),
					},
				},
			},
			ConnectedSubnetRoutePolicy: &managednetworkfabric.L3IsolationDomainPatchPropertiesConnectedSubnetRoutePolicyArgs{
				ExportRoutePolicyId: pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
			},
			Description:                  pulumi.String("creating L3 isolation domain"),
			L3IsolationDomainName:        pulumi.String("example-l3domain"),
			Location:                     pulumi.String("eastus"),
			NetworkFabricId:              pulumi.String("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName"),
			RedistributeConnectedSubnets: pulumi.String("True"),
			RedistributeStaticRoutes:     pulumi.String("False"),
			ResourceGroupName:            pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
