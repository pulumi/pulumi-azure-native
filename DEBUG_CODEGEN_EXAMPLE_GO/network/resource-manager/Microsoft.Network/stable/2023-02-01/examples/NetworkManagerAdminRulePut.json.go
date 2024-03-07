package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewDefaultAdminRule(ctx, "defaultAdminRule", &network.DefaultAdminRuleArgs{
			ConfigurationName:  pulumi.String("myTestSecurityConfig"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			ResourceGroupName:  pulumi.String("rg1"),
			RuleCollectionName: pulumi.String("testRuleCollection"),
			RuleName:           pulumi.String("SampleAdminRule"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
