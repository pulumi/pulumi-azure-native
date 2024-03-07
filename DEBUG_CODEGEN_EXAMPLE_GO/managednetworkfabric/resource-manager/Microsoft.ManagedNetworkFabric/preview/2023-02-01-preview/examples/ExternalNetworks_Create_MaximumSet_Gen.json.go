package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewExternalNetwork(ctx, "externalNetwork", &managednetworkfabric.ExternalNetworkArgs{
			ExportRoutePolicyId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
			ExternalNetworkName:   pulumi.String("example-externalnetwork"),
			ImportRoutePolicyId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
			L3IsolationDomainName: pulumi.String("example-l3domain"),
			OptionAProperties: &managednetworkfabric.ExternalNetworkPropertiesOptionAPropertiesArgs{
				Mtu:                 pulumi.Int(1500),
				PeerASN:             pulumi.Int(65047),
				PrimaryIpv4Prefix:   pulumi.String("10.1.1.0/30"),
				PrimaryIpv6Prefix:   pulumi.String("3FFE:FFFF:0:CD30::a0/126"),
				SecondaryIpv4Prefix: pulumi.String("10.1.1.4/30"),
				SecondaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::a4/126"),
				VlanId:              pulumi.Int(1001),
			},
			OptionBProperties: &managednetworkfabric.OptionBPropertiesArgs{
				ExportRouteTargets: pulumi.StringArray{
					pulumi.String("65046:10039"),
				},
				ImportRouteTargets: pulumi.StringArray{
					pulumi.String("65046:10039"),
				},
			},
			PeeringOption:     pulumi.String("OptionA"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
