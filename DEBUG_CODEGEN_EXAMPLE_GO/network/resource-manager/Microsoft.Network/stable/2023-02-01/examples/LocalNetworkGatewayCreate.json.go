package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewLocalNetworkGateway(ctx, "localNetworkGateway", &network.LocalNetworkGatewayArgs{
			Fqdn:             pulumi.String("site1.contoso.com"),
			GatewayIpAddress: pulumi.String("11.12.13.14"),
			LocalNetworkAddressSpace: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String("10.1.0.0/16"),
				},
			},
			LocalNetworkGatewayName: pulumi.String("localgw"),
			Location:                pulumi.String("Central US"),
			ResourceGroupName:       pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
