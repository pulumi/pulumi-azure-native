package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewNamespaceNetworkRuleSet(ctx, "namespaceNetworkRuleSet", &eventhub.NamespaceNetworkRuleSetArgs{
			DefaultAction: pulumi.String("Deny"),
			IpRules: eventhub.NWRuleSetIpRulesArray{
				&eventhub.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.1"),
				},
				&eventhub.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.2"),
				},
				&eventhub.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.3"),
				},
				&eventhub.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.4"),
				},
				&eventhub.NWRuleSetIpRulesArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("1.1.1.5"),
				},
			},
			NamespaceName:     pulumi.String("sdk-Namespace-6019"),
			ResourceGroupName: pulumi.String("ResourceGroup"),
			VirtualNetworkRules: eventhub.NWRuleSetVirtualNetworkRulesArray{
				&eventhub.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(true),
					Subnet: &eventhub.SubnetArgs{
						Id: pulumi.String("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
					},
				},
				&eventhub.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
					Subnet: &eventhub.SubnetArgs{
						Id: pulumi.String("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
					},
				},
				&eventhub.NWRuleSetVirtualNetworkRulesArgs{
					IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
					Subnet: &eventhub.SubnetArgs{
						Id: pulumi.String("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
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
