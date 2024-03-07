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
					LoadBalancerBackendAddresses: network.LoadBalancerBackendAddressArray{
						{
							LoadBalancerFrontendIPConfiguration: {
								Id: pulumi.String("/subscriptions/subid/resourceGroups/regional-lb-rg1/providers/Microsoft.Network/loadBalancers/regional-lb/frontendIPConfigurations/fe-rlb"),
							},
							Name: pulumi.String("regional-lb1-address"),
						},
					},
					Name: pulumi.String("be-lb"),
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
			LoadBalancerName: pulumi.String("lb"),
			LoadBalancingRules: []network.LoadBalancingRuleArgs{
				{
					BackendAddressPool: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
					},
					BackendPort:      pulumi.Int(80),
					EnableFloatingIP: pulumi.Bool(false),
					FrontendIPConfiguration: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					},
					FrontendPort:         pulumi.Int(80),
					IdleTimeoutInMinutes: pulumi.Int(15),
					LoadDistribution:     pulumi.String("Default"),
					Name:                 pulumi.String("rulelb"),
					Probe: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
					},
					Protocol: pulumi.String("Tcp"),
				},
			},
			Location: pulumi.String("eastus"),
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
				Name: pulumi.String("Standard"),
				Tier: pulumi.String("Global"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
