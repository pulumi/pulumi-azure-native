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
					TunnelInterfaces: network.GatewayLoadBalancerTunnelInterfaceArray{
						&network.GatewayLoadBalancerTunnelInterfaceArgs{
							Identifier: pulumi.Int(900),
							Port:       pulumi.Int(15000),
							Protocol:   pulumi.String("VXLAN"),
							Type:       pulumi.String("Internal"),
						},
						&network.GatewayLoadBalancerTunnelInterfaceArgs{
							Identifier: pulumi.Int(901),
							Port:       pulumi.Int(15001),
							Protocol:   pulumi.String("VXLAN"),
							Type:       pulumi.String("Internal"),
						},
					},
				},
			},
			FrontendIPConfigurations: network.FrontendIPConfigurationArray{
				&network.FrontendIPConfigurationArgs{
					Name: pulumi.String("fe-lb"),
					Subnet: &network.SubnetTypeArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
					},
				},
			},
			InboundNatPools:  network.InboundNatPoolArray{},
			LoadBalancerName: pulumi.String("lb"),
			LoadBalancingRules: network.LoadBalancingRuleArray{
				&network.LoadBalancingRuleArgs{
					BackendAddressPools: network.SubResourceArray{
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
						},
					},
					BackendPort:      pulumi.Int(0),
					EnableFloatingIP: pulumi.Bool(true),
					FrontendIPConfiguration: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					},
					FrontendPort:         pulumi.Int(0),
					IdleTimeoutInMinutes: pulumi.Int(15),
					LoadDistribution:     pulumi.String("Default"),
					Name:                 pulumi.String("rulelb"),
					Probe: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
					},
					Protocol: pulumi.String("All"),
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
				Name: pulumi.String("Premium"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
