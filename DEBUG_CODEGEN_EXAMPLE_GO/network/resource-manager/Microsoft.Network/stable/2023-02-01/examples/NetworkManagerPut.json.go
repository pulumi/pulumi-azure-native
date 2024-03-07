package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkManager(ctx, "networkManager", &network.NetworkManagerArgs{
			Description:        pulumi.String("My Test Network Manager"),
			NetworkManagerName: pulumi.String("TestNetworkManager"),
			NetworkManagerScopeAccesses: pulumi.StringArray{
				pulumi.String("Connectivity"),
			},
			NetworkManagerScopes: &network.NetworkManagerPropertiesNetworkManagerScopesArgs{
				ManagementGroups: pulumi.StringArray{
					pulumi.String("/providers/Microsoft.Management/managementGroups/sampleMG"),
				},
				Subscriptions: pulumi.StringArray{
					pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
