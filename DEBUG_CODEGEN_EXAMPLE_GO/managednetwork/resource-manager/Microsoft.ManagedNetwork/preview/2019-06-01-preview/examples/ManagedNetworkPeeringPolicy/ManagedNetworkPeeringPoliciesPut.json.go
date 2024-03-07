package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetwork.NewManagedNetworkPeeringPolicy(ctx, "managedNetworkPeeringPolicy", &managednetwork.ManagedNetworkPeeringPolicyArgs{
			ManagedNetworkName:              pulumi.String("myManagedNetwork"),
			ManagedNetworkPeeringPolicyName: pulumi.String("myHubAndSpoke"),
			Properties: managednetwork.ManagedNetworkPeeringPolicyPropertiesResponse{
				Hub: &managednetwork.ResourceIdArgs{
					Id: pulumi.String("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myHubVnet"),
				},
				Spokes: managednetwork.ResourceIdArray{
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork/managedNetworkGroups/myManagedNetworkGroup1"),
					},
				},
				Type: pulumi.String("HubAndSpokeTopology"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
