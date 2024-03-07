package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewConfigurationPolicyGroup(ctx, "configurationPolicyGroup", &network.ConfigurationPolicyGroupArgs{
			ConfigurationPolicyGroupName: pulumi.String("policyGroup1"),
			IsDefault:                    pulumi.Bool(true),
			PolicyMembers: network.VpnServerConfigurationPolicyGroupMemberArray{
				&network.VpnServerConfigurationPolicyGroupMemberArgs{
					AttributeType:  pulumi.String("RadiusAzureGroupId"),
					AttributeValue: pulumi.String("6ad1bd08"),
					Name:           pulumi.String("policy1"),
				},
				&network.VpnServerConfigurationPolicyGroupMemberArgs{
					AttributeType:  pulumi.String("CertificateGroupId"),
					AttributeValue: pulumi.String("red.com"),
					Name:           pulumi.String("policy2"),
				},
			},
			Priority:                   pulumi.Int(0),
			ResourceGroupName:          pulumi.String("rg1"),
			VpnServerConfigurationName: pulumi.String("vpnServerConfiguration1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
