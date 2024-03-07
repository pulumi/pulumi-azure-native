package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewLoadBalancer(ctx, "loadBalancer", &network.LoadBalancerArgs{
			BackendAddressPools: []network.BackendAddressPoolArgs{
				{
					Name: pulumi.String("be-lb1"),
				},
				{
					Name: pulumi.String("be-lb2"),
				},
			},
			FrontendIPConfigurations: []network.FrontendIPConfigurationArgs{
				{
					Name: pulumi.String("fe-lb"),
					Subnet: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
					},
				},
			},
			InboundNatPools:  network.InboundNatPoolArray{},
			LoadBalancerName: pulumi.String("lb"),
			LoadBalancingRules: []network.LoadBalancingRuleArgs{
				{
					BackendAddressPool: nil,
					BackendAddressPools: network.SubResourceArray{
						{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb1"),
						},
						{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb2"),
						},
					},
					BackendPort:      pulumi.Int(0),
					EnableFloatingIP: pulumi.Bool(true),
					FrontendIPConfiguration: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					},
					FrontendPort:         pulumi.Int(0),
					IdleTimeoutInMinutes: pulumi.Int(15),
					LoadDistribution:     pulumi.String("Default"),
					Name:                 pulumi.String("rulelb"),
					Probe: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
					},
					Protocol: pulumi.String("All"),
				},
			},
			Location:      pulumi.String("eastus"),
			OutboundRules: network.OutboundRuleArray{},
			Probes: []network.ProbeArgs{
				{
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
				Name: pulumi.String("Premium"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
