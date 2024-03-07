package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/delegatednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := delegatednetwork.NewDelegatedSubnetServiceDetails(ctx, "delegatedSubnetServiceDetails", &delegatednetwork.DelegatedSubnetServiceDetailsArgs{
			ControllerDetails: &delegatednetwork.ControllerDetailsTypeArgs{
				Id: pulumi.String("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/dnctestcontroller"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestRG"),
			ResourceName:      pulumi.String("delegated1"),
			SubnetDetails: &delegatednetwork.SubnetDetailsArgs{
				Id: pulumi.String("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
