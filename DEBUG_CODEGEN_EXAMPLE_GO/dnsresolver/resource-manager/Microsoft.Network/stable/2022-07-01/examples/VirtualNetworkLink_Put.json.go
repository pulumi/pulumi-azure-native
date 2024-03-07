package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateResolverVirtualNetworkLink(ctx, "privateResolverVirtualNetworkLink", &network.PrivateResolverVirtualNetworkLinkArgs{
			DnsForwardingRulesetName: pulumi.String("sampleDnsForwardingRuleset"),
			Metadata: pulumi.StringMap{
				"additionalProp1": pulumi.String("value1"),
			},
			ResourceGroupName: pulumi.String("sampleResourceGroup"),
			VirtualNetwork: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"),
			},
			VirtualNetworkLinkName: pulumi.String("sampleVirtualNetworkLink"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
