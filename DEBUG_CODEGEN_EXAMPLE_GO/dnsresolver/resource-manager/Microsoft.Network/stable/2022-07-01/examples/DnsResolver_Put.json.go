package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewDnsResolver(ctx, "dnsResolver", &network.DnsResolverArgs{
			DnsResolverName:   pulumi.String("sampleDnsResolver"),
			Location:          pulumi.String("westus2"),
			ResourceGroupName: pulumi.String("sampleResourceGroup"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualNetwork: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/cbb1387e-4b03-44f2-ad41-58d4677b9873/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
