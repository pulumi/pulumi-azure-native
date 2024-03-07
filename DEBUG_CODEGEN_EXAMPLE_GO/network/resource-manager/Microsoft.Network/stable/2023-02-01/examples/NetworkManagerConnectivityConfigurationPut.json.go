package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewConnectivityConfiguration(ctx, "connectivityConfiguration", &network.ConnectivityConfigurationArgs{
			AppliesToGroups: network.ConnectivityGroupItemArray{
				&network.ConnectivityGroupItemArgs{
					GroupConnectivity: pulumi.String("None"),
					IsGlobal:          pulumi.String("False"),
					NetworkGroupId:    pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
					UseHubGateway:     pulumi.String("True"),
				},
			},
			ConfigurationName:     pulumi.String("myTestConnectivityConfig"),
			ConnectivityTopology:  pulumi.String("HubAndSpoke"),
			DeleteExistingPeering: pulumi.String("True"),
			Description:           pulumi.String("Sample Configuration"),
			Hubs: network.HubArray{
				&network.HubArgs{
					ResourceId:   pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig"),
					ResourceType: pulumi.String("Microsoft.Network/virtualNetworks"),
				},
			},
			IsGlobal:           pulumi.String("True"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			ResourceGroupName:  pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
