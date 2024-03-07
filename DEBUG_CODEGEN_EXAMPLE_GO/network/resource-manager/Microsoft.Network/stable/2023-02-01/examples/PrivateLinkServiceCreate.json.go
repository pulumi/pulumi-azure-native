package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateLinkService(ctx, "privateLinkService", &network.PrivateLinkServiceArgs{
			AutoApproval: &network.PrivateLinkServicePropertiesAutoApprovalArgs{
				Subscriptions: pulumi.StringArray{
					pulumi.String("subscription1"),
					pulumi.String("subscription2"),
				},
			},
			Fqdns: pulumi.StringArray{
				pulumi.String("fqdn1"),
				pulumi.String("fqdn2"),
				pulumi.String("fqdn3"),
			},
			IpConfigurations: []network.PrivateLinkServiceIpConfigurationArgs{
				{
					Name:                      pulumi.String("fe-lb"),
					PrivateIPAddress:          pulumi.String("10.0.1.4"),
					PrivateIPAddressVersion:   pulumi.String("IPv4"),
					PrivateIPAllocationMethod: pulumi.String("Static"),
					Subnet: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
					},
				},
			},
			LoadBalancerFrontendIpConfigurations: []network.FrontendIPConfigurationArgs{
				{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
				},
			},
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("testPls"),
			Visibility: &network.PrivateLinkServicePropertiesVisibilityArgs{
				Subscriptions: pulumi.StringArray{
					pulumi.String("subscription1"),
					pulumi.String("subscription2"),
					pulumi.String("subscription3"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
