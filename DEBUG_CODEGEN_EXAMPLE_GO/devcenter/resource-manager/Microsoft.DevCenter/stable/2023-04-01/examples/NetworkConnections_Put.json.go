package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewNetworkConnection(ctx, "networkConnection", &devcenter.NetworkConnectionArgs{
			DomainJoinType:              pulumi.String("HybridAzureADJoin"),
			DomainName:                  pulumi.String("mydomaincontroller.local"),
			DomainPassword:              pulumi.String("Password value for user"),
			DomainUsername:              pulumi.String("testuser@mydomaincontroller.local"),
			Location:                    pulumi.String("centralus"),
			NetworkConnectionName:       pulumi.String("uswest3network"),
			NetworkingResourceGroupName: pulumi.String("NetworkInterfaces"),
			ResourceGroupName:           pulumi.String("rg1"),
			SubnetId:                    pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ExampleRG/providers/Microsoft.Network/virtualNetworks/ExampleVNet/subnets/default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
