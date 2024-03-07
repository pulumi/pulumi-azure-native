package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewInboundNatRule(ctx, "inboundNatRule", &network.InboundNatRuleArgs{
			BackendPort:      pulumi.Int(3389),
			EnableFloatingIP: pulumi.Bool(false),
			EnableTcpReset:   pulumi.Bool(false),
			FrontendIPConfiguration: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
			},
			FrontendPort:         pulumi.Int(3390),
			IdleTimeoutInMinutes: pulumi.Int(4),
			InboundNatRuleName:   pulumi.String("natRule1.1"),
			LoadBalancerName:     pulumi.String("lb1"),
			Protocol:             pulumi.String("Tcp"),
			ResourceGroupName:    pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
