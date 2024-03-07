package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewUserRule(ctx, "userRule", &network.UserRuleArgs{
			ConfigurationName: pulumi.String("myTestSecurityConfig"),
			Description:       pulumi.String("Sample User Rule"),
			DestinationPortRanges: pulumi.StringArray{
				pulumi.String("22"),
			},
			Destinations: network.AddressPrefixItemArray{
				&network.AddressPrefixItemArgs{
					AddressPrefix:     pulumi.String("*"),
					AddressPrefixType: pulumi.String("IPPrefix"),
				},
			},
			Direction:          pulumi.String("Inbound"),
			Kind:               pulumi.String("Custom"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			Protocol:           pulumi.String("Tcp"),
			ResourceGroupName:  pulumi.String("rg1"),
			RuleCollectionName: pulumi.String("testRuleCollection"),
			RuleName:           pulumi.String("SampleUserRule"),
			SourcePortRanges: pulumi.StringArray{
				pulumi.String("0-65535"),
			},
			Sources: network.AddressPrefixItemArray{
				&network.AddressPrefixItemArgs{
					AddressPrefix:     pulumi.String("*"),
					AddressPrefixType: pulumi.String("IPPrefix"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
