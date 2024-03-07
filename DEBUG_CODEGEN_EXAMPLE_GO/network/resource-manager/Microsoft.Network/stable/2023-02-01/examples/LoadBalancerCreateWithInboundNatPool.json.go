package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewLoadBalancer(ctx, "loadBalancer", &network.LoadBalancerArgs{
			BackendAddressPools: network.BackendAddressPoolArray{},
			FrontendIPConfigurations: []network.FrontendIPConfigurationArgs{
				{
					Id:                        pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
					Name:                      pulumi.String("test"),
					PrivateIPAllocationMethod: pulumi.String("Dynamic"),
					Subnet: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet"),
					},
					Zones: pulumi.StringArray{},
				},
			},
			InboundNatPools: []network.InboundNatPoolArgs{
				{
					BackendPort:      pulumi.Int(8888),
					EnableFloatingIP: pulumi.Bool(true),
					EnableTcpReset:   pulumi.Bool(true),
					FrontendIPConfiguration: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
					},
					FrontendPortRangeEnd:   pulumi.Int(8085),
					FrontendPortRangeStart: pulumi.Int(8080),
					Id:                     pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test"),
					IdleTimeoutInMinutes:   pulumi.Int(10),
					Name:                   pulumi.String("test"),
					Protocol:               pulumi.String("Tcp"),
				},
			},
			InboundNatRules:    network.InboundNatRuleTypeArray{},
			LoadBalancerName:   pulumi.String("lb"),
			LoadBalancingRules: network.LoadBalancingRuleArray{},
			Location:           pulumi.String("eastus"),
			OutboundRules:      network.OutboundRuleArray{},
			Probes:             network.ProbeArray{},
			ResourceGroupName:  pulumi.String("rg1"),
			Sku: &network.LoadBalancerSkuArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
