package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewDnsForwardingRuleset(ctx, "dnsForwardingRuleset", &network.DnsForwardingRulesetArgs{
			DnsForwardingRulesetName: pulumi.String("samplednsForwardingRuleset"),
			DnsResolverOutboundEndpoints: []network.SubResourceArgs{
				{
					Id: pulumi.String("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint0"),
				},
				{
					Id: pulumi.String("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint1"),
				},
			},
			Location:          pulumi.String("westus2"),
			ResourceGroupName: pulumi.String("sampleResourceGroup"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
