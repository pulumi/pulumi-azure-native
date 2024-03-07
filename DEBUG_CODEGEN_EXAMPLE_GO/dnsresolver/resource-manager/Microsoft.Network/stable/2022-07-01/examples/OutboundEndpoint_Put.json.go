package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewOutboundEndpoint(ctx, "outboundEndpoint", &network.OutboundEndpointArgs{
			DnsResolverName:      pulumi.String("sampleDnsResolver"),
			Location:             pulumi.String("westus2"),
			OutboundEndpointName: pulumi.String("sampleOutboundEndpoint"),
			ResourceGroupName:    pulumi.String("sampleResourceGroup"),
			Subnet: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
			},
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
