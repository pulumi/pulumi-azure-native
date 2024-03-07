package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkTap(ctx, "networkTap", &managednetworkfabric.NetworkTapArgs{
			Annotation: pulumi.String("annotation"),
			Destinations: managednetworkfabric.NetworkTapPropertiesDestinationsArray{
				&managednetworkfabric.NetworkTapPropertiesDestinationsArgs{
					DestinationId:        pulumi.String("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsloationDomains/example-l3Domain/internalNetworks/example-internalNetwork"),
					DestinationTapRuleId: pulumi.String("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-destinationTapRule"),
					DestinationType:      pulumi.String("IsolationDomain"),
					IsolationDomainProperties: &managednetworkfabric.IsolationDomainPropertiesArgs{
						Encapsulation: pulumi.String("None"),
						NeighborGroupIds: pulumi.StringArray{
							pulumi.String("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
						},
					},
					Name: pulumi.String("example-destinaionName"),
				},
			},
			Location:              pulumi.String("eastuseuap"),
			NetworkPacketBrokerId: pulumi.String("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/example-networkPacketBroker"),
			NetworkTapName:        pulumi.String("example-networkTap"),
			PollingType:           pulumi.String("Pull"),
			ResourceGroupName:     pulumi.String("example-rg"),
			Tags: pulumi.StringMap{
				"key6024": pulumi.String("1234"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
