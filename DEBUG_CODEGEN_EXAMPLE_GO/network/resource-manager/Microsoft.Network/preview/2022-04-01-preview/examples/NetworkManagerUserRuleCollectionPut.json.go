package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewUserRuleCollection(ctx, "userRuleCollection", &network.UserRuleCollectionArgs{
			AppliesToGroups: network.NetworkManagerSecurityGroupItemArray{
				&network.NetworkManagerSecurityGroupItemArgs{
					NetworkGroupId: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"),
				},
			},
			ConfigurationName:  pulumi.String("myTestSecurityConfig"),
			Description:        pulumi.String("A sample policy"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			ResourceGroupName:  pulumi.String("rg1"),
			RuleCollectionName: pulumi.String("testRuleCollection"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
