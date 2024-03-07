package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkSecurityGroup(ctx, "networkSecurityGroup", &network.NetworkSecurityGroupArgs{
			Location:                 pulumi.String("eastus"),
			NetworkSecurityGroupName: pulumi.String("testnsg"),
			ResourceGroupName:        pulumi.String("rg1"),
			SecurityRules: network.SecurityRuleTypeArray{
				&network.SecurityRuleTypeArgs{
					Access:                   pulumi.String("Allow"),
					DestinationAddressPrefix: pulumi.String("*"),
					DestinationPortRange:     pulumi.String("80"),
					Direction:                pulumi.String("Inbound"),
					Name:                     pulumi.String("rule1"),
					Priority:                 pulumi.Int(130),
					Protocol:                 pulumi.String("*"),
					SourceAddressPrefix:      pulumi.String("*"),
					SourcePortRange:          pulumi.String("*"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
