package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewNamespaceNetworkRuleSet(ctx, "namespaceNetworkRuleSet", &servicebus.NamespaceNetworkRuleSetArgs{
			DefaultAction: pulumi.String("Deny"),
			IpRules: servicebus.NWRuleSetIpRulesArray{
				&servicebus.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.1"),
				},
				&servicebus.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.2"),
				},
				&servicebus.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.3"),
				},
				&servicebus.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.4"),
				},
				&servicebus.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.5"),
				},
			},
			NamespaceName:     pulumi.String("sdk-Namespace-6019"),
			ResourceGroupName: pulumi.String("ResourceGroup"),
			VirtualNetworkRules: servicebus.NWRuleSetVirtualNetworkRulesArray{
				&servicebus.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(true),
					Subnet: &servicebus.SubnetArgs{
						Id: pulumi.String("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
					},
				},
				&servicebus.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
					Subnet: &servicebus.SubnetArgs{
						Id: pulumi.String("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
					},
				},
				&servicebus.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
					Subnet: &servicebus.SubnetArgs{
						Id: pulumi.String("/subscriptions/854d368f-1828-428f-8f3c-f2affa9b2f7d/resourcegroups/alitest/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
