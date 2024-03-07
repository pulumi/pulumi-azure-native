package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewRoutePolicy(ctx, "routePolicy", &managednetworkfabric.RoutePolicyArgs{
			Annotation:        pulumi.String("annotationValue"),
			Location:          pulumi.String("EastUS"),
			ResourceGroupName: pulumi.String("rgRoutePolicies"),
			RoutePolicyName:   pulumi.String("routePolicyName"),
			Statements: managednetworkfabric.RoutePolicyStatementPropertiesArray{
				&managednetworkfabric.RoutePolicyStatementPropertiesArgs{
					Action: &managednetworkfabric.StatementActionPropertiesArgs{
						ActionType: pulumi.String("Permit"),
						IpCommunityProperties: &managednetworkfabric.ActionIpCommunityPropertiesArgs{
							Add: &managednetworkfabric.IpCommunityIdListArgs{
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
							Delete: &managednetworkfabric.IpCommunityIdListArgs{
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
							Set: &managednetworkfabric.IpCommunityIdListArgs{
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
						},
						IpExtendedCommunityProperties: &managednetworkfabric.ActionIpExtendedCommunityPropertiesArgs{
							Add: &managednetworkfabric.IpExtendedCommunityIdListArgs{
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
							Delete: &managednetworkfabric.IpExtendedCommunityIdListArgs{
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
							Set: &managednetworkfabric.IpExtendedCommunityIdListArgs{
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
						},
						LocalPreference: pulumi.Float64(20),
					},
					Annotation: pulumi.String("annotationValue"),
					Condition: &managednetworkfabric.StatementConditionPropertiesArgs{
						IpCommunityIds: pulumi.StringArray{
							pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
						},
						IpExtendedCommunityIds: pulumi.StringArray{
							pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
						},
						IpPrefixId: pulumi.String("subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
					},
					SequenceNumber: pulumi.Float64(7),
				},
			},
			Tags: pulumi.StringMap{
				"key8254": pulumi.String(""),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
