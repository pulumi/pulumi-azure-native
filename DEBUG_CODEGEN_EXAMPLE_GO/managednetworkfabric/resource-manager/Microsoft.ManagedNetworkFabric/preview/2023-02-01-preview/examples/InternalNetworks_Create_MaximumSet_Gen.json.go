package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewInternalNetwork(ctx, "internalNetwork", &managednetworkfabric.InternalNetworkArgs{
			BgpConfiguration: managednetworkfabric.BgpConfigurationResponse{
				AllowAS:               pulumi.Int(1),
				AllowASOverride:       pulumi.String("Enable"),
				DefaultRouteOriginate: pulumi.String("True"),
				Ipv4ListenRangePrefixes: pulumi.StringArray{
					pulumi.String("10.1.0.0/25"),
				},
				Ipv4NeighborAddress: managednetworkfabric.NeighborAddressArray{
					&managednetworkfabric.NeighborAddressArgs{
						Address: pulumi.String("10.1.0.0"),
					},
				},
				Ipv6ListenRangePrefixes: pulumi.StringArray{
					pulumi.String("2fff::/66"),
				},
				Ipv6NeighborAddress: managednetworkfabric.NeighborAddressArray{
					&managednetworkfabric.NeighborAddressArgs{
						Address: pulumi.String("2fff::"),
					},
				},
				PeerASN: pulumi.Int(6),
			},
			ConnectedIPv4Subnets: []managednetworkfabric.ConnectedSubnetArgs{
				{
					Prefix: pulumi.String("10.0.0.0/24"),
				},
			},
			ConnectedIPv6Subnets: []managednetworkfabric.ConnectedSubnetArgs{
				{
					Prefix: pulumi.String("3FFE:FFFF:0:CD30::a0/29"),
				},
			},
			ExportRoutePolicyId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
			ImportRoutePolicyId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
			InternalNetworkName:   pulumi.String("example-internalnetwork"),
			L3IsolationDomainName: pulumi.String("example-l3domain"),
			Mtu:                   pulumi.Int(1500),
			ResourceGroupName:     pulumi.String("resourceGroupName"),
			StaticRouteConfiguration: managednetworkfabric.StaticRouteConfigurationResponse{
				Ipv4Routes: managednetworkfabric.StaticRoutePropertiesArray{
					&managednetworkfabric.StaticRoutePropertiesArgs{
						NextHop: pulumi.StringArray{
							pulumi.String("10.0.0.1"),
						},
						Prefix: pulumi.String("10.1.0.0/24"),
					},
				},
				Ipv6Routes: managednetworkfabric.StaticRoutePropertiesArray{
					&managednetworkfabric.StaticRoutePropertiesArgs{
						NextHop: pulumi.StringArray{
							pulumi.String("2ffe::1"),
						},
						Prefix: pulumi.String("2fff::/64"),
					},
				},
			},
			VlanId: pulumi.Int(501),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
