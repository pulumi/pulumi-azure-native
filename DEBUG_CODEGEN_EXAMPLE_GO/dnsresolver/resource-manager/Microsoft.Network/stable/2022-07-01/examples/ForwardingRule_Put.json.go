package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewForwardingRule(ctx, "forwardingRule", &network.ForwardingRuleArgs{
			DnsForwardingRulesetName: pulumi.String("sampleDnsForwardingRuleset"),
			DomainName:               pulumi.String("contoso.com."),
			ForwardingRuleName:       pulumi.String("sampleForwardingRule"),
			ForwardingRuleState:      pulumi.String("Enabled"),
			Metadata: pulumi.StringMap{
				"additionalProp1": pulumi.String("value1"),
			},
			ResourceGroupName: pulumi.String("sampleResourceGroup"),
			TargetDnsServers: []network.TargetDnsServerArgs{
				{
					IpAddress: pulumi.String("10.0.0.1"),
					Port:      pulumi.Int(53),
				},
				{
					IpAddress: pulumi.String("10.0.0.2"),
					Port:      pulumi.Int(53),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
