package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewLoadBalancer(ctx, "loadBalancer", &network.LoadBalancerArgs{
			BackendAddressPools: network.BackendAddressPoolArray{
				&network.BackendAddressPoolArgs{
					Name: pulumi.String("be-lb"),
				},
			},
			FrontendIPConfigurations: network.FrontendIPConfigurationArray{
				&network.FrontendIPConfigurationArgs{
					GatewayLoadBalancer: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider"),
					},
					Name: pulumi.String("fe-lb"),
					Subnet: &network.SubnetTypeArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
					},
				},
			},
			InboundNatPools: network.InboundNatPoolArray{},
			InboundNatRules: network.InboundNatRuleTypeArray{
				&network.InboundNatRuleTypeArgs{
					BackendPort:      pulumi.Int(3389),
					EnableFloatingIP: pulumi.Bool(true),
					FrontendIPConfiguration: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					},
					FrontendPort:         pulumi.Int(3389),
					IdleTimeoutInMinutes: pulumi.Int(15),
					Name:                 pulumi.String("in-nat-rule"),
					Protocol:             pulumi.String("Tcp"),
				},
			},
			LoadBalancerName: pulumi.String("lb"),
			LoadBalancingRules: network.LoadBalancingRuleArray{
				&network.LoadBalancingRuleArgs{
					BackendAddressPool: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
					},
					BackendPort:      pulumi.Int(80),
					EnableFloatingIP: pulumi.Bool(true),
					FrontendIPConfiguration: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					},
					FrontendPort:         pulumi.Int(80),
					IdleTimeoutInMinutes: pulumi.Int(15),
					LoadDistribution:     pulumi.String("Default"),
					Name:                 pulumi.String("rulelb"),
					Probe: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
					},
					Protocol: pulumi.String("Tcp"),
				},
			},
			Location:      pulumi.String("eastus"),
			OutboundRules: network.OutboundRuleArray{},
			Probes: network.ProbeArray{
				&network.ProbeArgs{
					IntervalInSeconds: pulumi.Int(15),
					Name:              pulumi.String("probe-lb"),
					NumberOfProbes:    pulumi.Int(2),
					Port:              pulumi.Int(80),
					ProbeThreshold:    pulumi.Int(1),
					Protocol:          pulumi.String("Http"),
					RequestPath:       pulumi.String("healthcheck.aspx"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
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
