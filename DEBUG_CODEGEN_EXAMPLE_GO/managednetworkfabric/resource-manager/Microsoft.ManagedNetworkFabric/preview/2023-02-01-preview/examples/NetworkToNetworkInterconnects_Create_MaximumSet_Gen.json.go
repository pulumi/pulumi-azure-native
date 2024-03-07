package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkToNetworkInterconnect(ctx, "networkToNetworkInterconnect", &managednetworkfabric.NetworkToNetworkInterconnectArgs{
			IsManagementType: pulumi.String("True"),
			Layer2Configuration: &managednetworkfabric.Layer2ConfigurationArgs{
				Mtu:       pulumi.Int(1500),
				PortCount: pulumi.Int(10),
			},
			Layer3Configuration: &managednetworkfabric.Layer3ConfigurationArgs{
				ExportRoutePolicyId: pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
				ImportRoutePolicyId: pulumi.String("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
				PeerASN:             pulumi.Int(50272),
				PrimaryIpv4Prefix:   pulumi.String("172.31.0.0/31"),
				PrimaryIpv6Prefix:   pulumi.String("3FFE:FFFF:0:CD30::a0/126"),
				SecondaryIpv4Prefix: pulumi.String("172.31.0.20/31"),
				SecondaryIpv6Prefix: pulumi.String("3FFE:FFFF:0:CD30::a4/126"),
				VlanId:              pulumi.Int(2064),
			},
			NetworkFabricName:                pulumi.String("FabricName"),
			NetworkToNetworkInterconnectName: pulumi.String("DefaultNNI"),
			NniType:                          pulumi.String("CE"),
			ResourceGroupName:                pulumi.String("resourceGroupName"),
			UseOptionB:                       pulumi.String("False"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
