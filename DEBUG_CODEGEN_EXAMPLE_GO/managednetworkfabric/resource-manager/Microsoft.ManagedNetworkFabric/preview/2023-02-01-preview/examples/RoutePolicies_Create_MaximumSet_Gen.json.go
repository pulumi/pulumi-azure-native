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
			Statements: []managednetworkfabric.RoutePolicyStatementPropertiesArgs{
				{
					Action: {
						ActionType: pulumi.String("Permit"),
						IpCommunityProperties: {
							Add: {
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
							Delete: {
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
							Set: {
								IpCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName"),
								},
							},
						},
						IpExtendedCommunityProperties: {
							Add: {
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
							Delete: {
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
							Set: {
								IpExtendedCommunityIds: pulumi.StringArray{
									pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName"),
								},
							},
						},
						LocalPreference: pulumi.Float64(20),
					},
					Annotation: pulumi.String("annotationValue"),
					Condition: {
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
