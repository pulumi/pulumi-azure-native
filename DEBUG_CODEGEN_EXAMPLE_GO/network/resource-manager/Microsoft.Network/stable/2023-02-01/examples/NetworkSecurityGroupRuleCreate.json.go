package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSecurityRule(ctx, "securityRule", &network.SecurityRuleArgs{
			Access:                   pulumi.String("Deny"),
			DestinationAddressPrefix: pulumi.String("11.0.0.0/8"),
			DestinationPortRange:     pulumi.String("8080"),
			Direction:                pulumi.String("Outbound"),
			NetworkSecurityGroupName: pulumi.String("testnsg"),
			Priority:                 pulumi.Int(100),
			Protocol:                 pulumi.String("*"),
			ResourceGroupName:        pulumi.String("rg1"),
			SecurityRuleName:         pulumi.String("rule1"),
			SourceAddressPrefix:      pulumi.String("10.0.0.0/8"),
			SourcePortRange:          pulumi.String("*"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
